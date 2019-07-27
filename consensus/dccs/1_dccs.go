// Copyright 2019 The gonex Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// Package dccs implements the proof-of-foundation consensus engine.
package dccs

import (
	"bytes"
	"errors"
	"math"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/deployer"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/consensus/misc"
	"github.com/ethereum/go-ethereum/contracts/nexty/contract"
	"github.com/ethereum/go-ethereum/contracts/nexty/token"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	lru "github.com/hashicorp/golang-lru"
)

const (
	wiggleTime = 500 * time.Millisecond // Random delay (per signer) to allow concurrent signers
)

var (
	rewards = []*big.Int{
		big.NewInt(1e+4),
		big.NewInt(5e+3),
		big.NewInt(25e+2),
		big.NewInt(1250),
		big.NewInt(625),
		big.NewInt(500),
	} // rewards per year in percent of current total supply
	initialSupply = big.NewInt(18e+10)   // initial total supply in NTY
	blockPerYear  = big.NewInt(15778476) // Number of blocks per year with blocktime = 2s
)

// Init the first hardfork of DCCS consensus
func (d *Dccs) init1() *Dccs {
	// ensure that the hard-fork block must be divisible by both the old and new epoch value
	tlBlock := d.config.ThangLongBlock.Uint64()
	if tlBlock%d.config.Epoch != 0 {
		log.Crit("Unable to create DCCS consensus engine", "ThangLong block", tlBlock, "Epoch", d.config.Epoch)
		return nil
	}
	if tlBlock%d.config.ThangLongEpoch != 0 {
		log.Crit("Unable to create DCCS consensus engine", "ThangLong block", tlBlock, "ThangLongEpoch", d.config.ThangLongEpoch)
		return nil
	}
	return d
}

// verifyHeader1 checks whether a header conforms to the consensus rules.The
// caller may optionally pass in a batch of parents (ascending order) to avoid
// looking those up from the database. This is useful for concurrently verifying
// a batch of new headers.
func (d *Dccs) verifyHeader1(chain consensus.ChainReader, header *types.Header, parents []*types.Header) error {
	if header.Number == nil {
		return errUnknownBlock
	}
	number := header.Number.Uint64()

	// Don't waste time checking blocks from the future
	if header.Time > uint64(time.Now().Unix()) {
		return consensus.ErrFutureBlock
	}

	checkpoint := d.config.IsCheckpoint(number)
	// Check that the extra-data contains both the vanity and signature
	if len(header.Extra) < extraVanity {
		return errMissingVanity
	}
	if len(header.Extra) < extraVanity+extraSeal {
		return errMissingSignature
	}
	// Ensure that the extra-data contains a signer list on checkpoint, but none otherwise
	signersBytes := len(header.Extra) - extraVanity - extraSeal
	if !checkpoint && signersBytes != 0 {
		return errExtraSigners
	}
	if checkpoint && signersBytes%common.AddressLength != 0 {
		return errInvalidCheckpointSigners
	}
	// Ensure that the mix digest is zero as we don't have fork protection currently
	if header.MixDigest != (common.Hash{}) {
		return errInvalidMixDigest
	}
	// Ensure that the block doesn't contain any uncles which are meaningless in PoA
	if header.UncleHash != uncleHash {
		return errInvalidUncleHash
	}
	// Ensure that the block's difficulty is meaningful (may not be correct at this point)
	if number > 0 {
		if header.Difficulty == nil {
			return errInvalidDifficulty
		}
	}
	// If all checks passed, validate any special fields for hard forks
	if err := misc.VerifyForkHashes(chain.Config(), header, false); err != nil {
		return err
	}
	// All basic checks passed, verify cascading fields
	return d.verifyCascadingFields1(chain, header, parents)
}

// verifyCascadingFields1 verifies all the header fields that are not standalone,
// rather depend on a batch of previous headers. The caller may optionally pass
// in a batch of parents (ascending order) to avoid looking those up from the
// database. This is useful for concurrently verifying a batch of new headers.
func (d *Dccs) verifyCascadingFields1(chain consensus.ChainReader, header *types.Header, parents []*types.Header) error {
	// The genesis block is the always valid dead-end
	number := header.Number.Uint64()
	if number == 0 {
		return nil
	}
	// Ensure that the block's timestamp isn't too close to it's parent
	var parent *types.Header
	if len(parents) > 0 {
		parent = parents[len(parents)-1]
	} else {
		parent = chain.GetHeader(header.ParentHash, number-1)
	}
	if parent == nil || parent.Number.Uint64() != number-1 || parent.Hash() != header.ParentHash {
		return consensus.ErrUnknownAncestor
	}
	if parent.Time+d.config.Period > header.Time {
		return ErrInvalidTimestamp
	}
	// Retrieve the snapshot needed to verify this header and cache it
	snap, err := d.snapshot1(chain, number-1, header.ParentHash, parents)
	if err != nil {
		return err
	}
	// If the block is a checkpoint block, verify the signer list
	if d.config.IsCheckpoint(number) {
		signers := make([]byte, len(snap.Signers)*common.AddressLength)
		for i, signer := range snap.signers1() {
			copy(signers[i*common.AddressLength:], signer.Address[:])
		}
		extraSuffix := len(header.Extra) - extraSeal
		if !bytes.Equal(header.Extra[extraVanity:extraSuffix], signers) {
			return errInvalidCheckpointSigners
		}
	}
	// All basic checks passed, verify the seal and return
	return d.verifySeal1(chain, header, parents)
}

// snapshot1 retrieves the authorization snapshot at a given point in time.
func (d *Dccs) snapshot1(chain consensus.ChainReader, number uint64, hash common.Hash, parents []*types.Header) (*Snapshot, error) {
	// Search for a snapshot in memory or on disk for checkpoints
	var (
		snap *Snapshot
	)
	cp := d.config.Snapshot(number + 1)
	// looping until data/state are available on local
	for snap == nil {
		// Get signers from Nexty staking smart contract at the latest epoch checkpoint from block number
		checkpoint := chain.GetHeaderByNumber(cp)
		if checkpoint == nil {
			log.Trace("snapshot header not available", "number", cp)
			continue
		}
		hash := checkpoint.Hash()
		log.Trace("Reading signers from epoch checkpoint", "number", cp, "hash", hash)
		// If an in-memory snapshot was found, use that
		if s, ok := d.recents.Get(hash); ok && number+1 != d.config.ThangLongBlock.Uint64() {
			snap = s.(*Snapshot)
			log.Trace("Loading snapshot from mem-cache", "hash", snap.Hash, "length", len(snap.signers()))
			break
		}
		state, err := chain.StateAt(checkpoint.Root)
		if state == nil || err != nil {
			log.Trace("snapshot state not available", "number", cp, "err", err)
			continue
		}
		size := state.GetCodeSize(chain.Config().Dccs.Contract)
		if size <= 0 || state.Error() != nil {
			log.Trace("snapshot contract state not available", "number", cp, "err", state.Error())
			continue
		}
		index := common.BigToHash(common.Big0)
		result := state.GetState(chain.Config().Dccs.Contract, index)
		var length int64
		if (result == common.Hash{}) {
			length = 0
		} else {
			length = result.Big().Int64()
		}
		log.Trace("Total number of signer from staking smart contract", "length", length)
		signers := make([]common.Address, length)
		key := crypto.Keccak256Hash(hexutil.MustDecode(index.String()))
		for i := 0; i < len(signers); i++ {
			log.Trace("key hash", "key", key)
			singer := state.GetState(chain.Config().Dccs.Contract, key)
			signers[i] = common.HexToAddress(singer.Hex())
			key = key.Plus()
		}
		snap = newSnapshot(d.config, d.signatures, number, hash, signers)
		// Store found snapshot into mem-cache
		d.recents.Add(snap.Hash, snap)
		break
	}

	// Set current block number for snapshot to calculate the inturn & difficulty
	snap.Number = number
	return snap, nil
}

// verifySeal1 checks whether the signature contained in the header satisfies the
// consensus protocol requirements. The method accepts an optional list of parent
// headers that aren't yet part of the local blockchain to generate the snapshots
// from.
func (d *Dccs) verifySeal1(chain consensus.ChainReader, header *types.Header, parents []*types.Header) error {
	// Verifying the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return errUnknownBlock
	}
	// Retrieve the snapshot needed to verify this header and cache it
	snap, err := d.snapshot1(chain, number-1, header.ParentHash, parents)
	if err != nil {
		return err
	}

	// Resolve the authorization key and check against signers
	signer, err := ecrecover(header, d.signatures)
	if err != nil {
		return err
	}
	if _, ok := snap.Signers[signer]; !ok {
		return errUnauthorizedSigner
	}

	headers, err := d.GetRecentHeaders(snap, chain, header, parents)
	if err != nil {
		return err
	}
	for _, h := range headers {
		sig, err := ecrecover(h, d.signatures)
		if err != nil {
			return err
		}
		if signer == sig {
			// Signer is among recents, only fail if the current block doesn't shift it out
			return errRecentlySigned
		}
	}

	var parent *types.Header
	if len(headers) > 0 {
		parent = headers[0]
	}

	// Ensure that the difficulty corresponds to the turn-ness of the signer
	signerDifficulty := snap.difficulty(signer, parent)
	if header.Difficulty.Uint64() != signerDifficulty {
		return errInvalidDifficulty
	}
	return nil
}

// prepare1 implements consensus.Engine, preparing all the consensus fields of the
// header for running the transactions on top.
func (d *Dccs) prepare1(chain consensus.ChainReader, header *types.Header) error {
	header.Nonce = types.BlockNonce{}
	// Get the beneficiary of signer from smart contract and set to header's coinbase to give sealing reward later
	number := header.Number.Uint64()
	cp := d.config.Snapshot(number)
	checkpoint := chain.GetHeaderByNumber(cp)
	if checkpoint != nil {
		root, _ := chain.StateAt(checkpoint.Root)
		index := common.BigToHash(common.Big1).String()[2:]
		coinbase := "0x000000000000000000000000" + header.Coinbase.String()[2:]
		key := crypto.Keccak256Hash(hexutil.MustDecode(coinbase + index))
		result := root.GetState(chain.Config().Dccs.Contract, key)
		beneficiary := common.HexToAddress(result.Hex())
		header.Coinbase = beneficiary
	} else {
		log.Error("state is not available at the block number", "number", cp)
		return errors.New("state is not available at the block number")
	}

	// Set the correct difficulty
	snap, err := d.snapshot1(chain, number-1, header.ParentHash, nil)
	if err != nil {
		return err
	}
	parent := chain.GetHeader(header.ParentHash, number-1)
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}
	header.Difficulty = CalcDifficulty1(snap, d.signer, parent)
	log.Trace("header.Difficulty", "difficulty", header.Difficulty)

	// Ensure the extra data has all it's components
	if len(header.Extra) < extraVanity {
		header.Extra = append(header.Extra, bytes.Repeat([]byte{0x00}, extraVanity-len(header.Extra))...)
	}
	header.Extra = header.Extra[:extraVanity]

	if d.config.IsCheckpoint(number) {
		for _, signer := range snap.signers1() {
			header.Extra = append(header.Extra, signer.Address[:]...)
		}
	}
	header.Extra = append(header.Extra, make([]byte, extraSeal)...)

	// Mix digest is reserved for now, set to empty
	header.MixDigest = common.Hash{}

	// Ensure the timestamp has the correct delay
	header.Time = parent.Time + d.config.Period
	if header.Time < uint64(time.Now().Unix()) {
		header.Time = uint64(time.Now().Unix())
	}
	return nil
}

// finalize1 implements consensus.Engine, ensuring no uncles are set, nor block
// rewards given, and returns the final block.
func (d *Dccs) finalize1(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header) {
	// Calculate any block reward for the sealer and commit the final state root
	d.calculateRewards(chain, state, header)
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	header.UncleHash = types.CalcUncleHash(nil)
}

// finalizeAndAssemble1 implements consensus.Engine, ensuring no uncles are set, nor block
// rewards given, and returns the final block.
func (d *Dccs) finalizeAndAssemble1(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {
	// Calculate any block reward for the sealer and commit the final state root
	d.calculateRewards(chain, state, header)
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
	header.UncleHash = types.CalcUncleHash(nil)

	// Assemble and return the final block for sealing
	return types.NewBlock(header, txs, nil, receipts), nil
}

// seal1 implements consensus.Engine, attempting to create a sealed block using
// the local signing credentials.
func (d *Dccs) seal1(chain consensus.ChainReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {
	header := block.Header()

	// Sealing the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return errUnknownBlock
	}
	// For 0-period chains, refuse to seal empty blocks (no reward but would spin sealing)
	if d.config.Period == 0 && len(block.Transactions()) == 0 {
		return errWaitTransactions
	}
	// Don't hold the signer fields for the entire sealing procedure
	d.lock.RLock()
	signer, signFn := d.signer, d.signFn
	d.lock.RUnlock()

	// Bail out if we're unauthorized to sign a block
	snap, err := d.snapshot1(chain, number-1, header.ParentHash, nil)
	if err != nil {
		return err
	}
	if _, authorized := snap.Signers[signer]; !authorized {
		return errUnauthorizedSigner
	}
	// If we're amongst the recent signers, wait for the next block
	headers, err := d.GetRecentHeaders(snap, chain, header, nil)
	if err != nil {
		return err
	}
	for _, h := range headers {
		sig, err := ecrecover(h, d.signatures)
		if err != nil {
			return err
		}
		if signer == sig {
			// Signer is among recents
			log.Info("Signed recently, must wait for others")
			return nil
		}
	}
	var parent *types.Header
	if len(headers) > 0 {
		parent = headers[0]
	}
	// Sweet, the protocol permits us to sign the block, wait for our time
	delay := time.Unix(int64(header.Time), 0).Sub(time.Now()) // nolint: gosimple
	if !snap.inturn1(signer, parent) {
		// It's not our turn explicitly to sign, delay it a bit
		offset, err := snap.offset(signer, parent)
		if err != nil {
			return err
		}
		wiggle := d.calcDelayTimeForOffset(offset)
		delay += wiggle
		log.Trace("Out-of-turn signing requested", "wiggle", common.PrettyDuration(wiggle))
	}
	// Sign all the things!
	sighash, err := signFn(accounts.Account{Address: signer}, accounts.MimetypeClique, DccsRLP(header))
	if err != nil {
		return err
	}
	copy(header.Extra[len(header.Extra)-extraSeal:], sighash)
	// Wait until sealing is terminated or delay timeout.
	log.Trace("Waiting for slot to sign and propagate", "delay", common.PrettyDuration(delay))
	go func() {
		select {
		case <-stop:
			return
		case <-time.After(delay):
		}

		select {
		case results <- block.WithSeal(header):
		default:
			log.Warn("Sealing result is not read by miner", "sealhash", SealHash(header))
		}
	}()

	return nil
}

// CalcDifficulty1 is the difficulty adjustment algorithm. It returns the difficulty
// that a new block should have based on the previous blocks in the chain and the
// current signer.
func CalcDifficulty1(snap *Snapshot, signer common.Address, parent *types.Header) *big.Int {
	return new(big.Int).SetUint64(snap.difficulty(signer, parent))
}

// Author implements consensus.Engine, returning the Ethereum address recovered
// from the signature in the header's extra-data section.
func (d *Dccs) Author(header *types.Header) (common.Address, error) {
	return ecrecover(header, d.signatures)
}

// GetRecentHeaders get some recent headers back from the current header.
// Return empty header list at checkpoint, a.k.a reset the recent headers at every checkpoint.
func (d *Dccs) GetRecentHeaders(snap *Snapshot, chain consensus.ChainReader, header *types.Header, parents []*types.Header) ([]*types.Header, error) {
	var headers []*types.Header
	number := header.Number.Uint64()
	// Reset the recent headers at checkpoint to ensure the in-turn signer
	if d.config.IsCheckpoint(number) {
		return headers, nil
	}
	limit := len(snap.Signers) / 2
	num, hash := number-1, header.ParentHash
	for i := 1; i <= limit; i++ {
		// shortcut for genesis block because it has no signature
		if num == 0 {
			break
		}
		var h *types.Header
		if len(parents) > 0 {
			// If we have explicit parents, pick from there (enforced)
			h = parents[len(parents)-1]
			if h.Hash() != hash || h.Number.Uint64() != num {
				return nil, consensus.ErrUnknownAncestor
			}
			parents = parents[:len(parents)-1]
		} else {
			// No explicit parents (or no more left), reach out to the database
			h = chain.GetHeader(hash, num)
			if header == nil {
				return nil, consensus.ErrUnknownAncestor
			}
		}
		headers = append(headers, h)
		if d.config.IsCheckpoint(num) {
			break
		}
		num, hash = num-1, h.ParentHash
	}
	return headers, nil
}

// calculateRewards calculate reward for block sealer
func (d *Dccs) calculateRewards(chain consensus.ChainReader, state *state.StateDB, header *types.Header) {
	number := header.Number.Uint64()
	yo := (number - chain.Config().Dccs.ThangLongBlock.Uint64()) / blockPerYear.Uint64()
	per := yo
	if per > 5 {
		per = 5
	}
	totalSupply := new(big.Int).Mul(initialSupply, big.NewInt(1e+18)) // total supply in Wei
	for i := uint64(1); i <= yo; i++ {
		r := i
		if r > 5 {
			r = 5
		}
		totalReward := new(big.Int).Mul(totalSupply, rewards[r])
		totalReward = totalReward.Div(totalReward, big.NewInt(1e+5))
		totalSupply = totalSupply.Add(totalSupply, totalReward)
	}
	totalYearReward := new(big.Int).Mul(totalSupply, rewards[per])
	totalYearReward = totalYearReward.Div(totalYearReward, big.NewInt(1e+5))
	log.Trace("Total reward for current year", "reward", totalYearReward, "total sypply", totalSupply)
	blockReward := new(big.Int).Div(totalYearReward, blockPerYear)
	log.Trace("Give reward for sealer", "beneficiary", header.Coinbase, "reward", blockReward, "number", number, "hash", header.Hash)
	state.AddBalance(header.Coinbase, blockReward)
}

// calcDelayTime calculate delay time for current sealing node
func (d *Dccs) calcDelayTime(snap *Snapshot, block *types.Block, signer common.Address) time.Duration {
	header := block.Header()
	number := header.Number.Uint64()
	sigs := snap.signers1()
	pos := 0
	for seen, sig := range sigs {
		if sig.Address == signer {
			pos = seen
		}
	}
	cp := d.config.Checkpoint(number)
	total := uint64(len(sigs))
	offset := (number - cp) - (number-cp)/total*total
	log.Trace("calcDelayTime", "number", number, "checkpoint", cp, "len", uint64(len(sigs)), "offset", offset)
	if pos >= int(offset) {
		pos -= int(offset)
	} else {
		pos += len(sigs) - int(offset)
	}
	return d.calcDelayTimeForOffset(pos)
}

// calcDelayTime calculate delay time for current sealing node
func (d *Dccs) calcDelayTimeForOffset(pos int) time.Duration {
	delay := time.Duration(0)
	wiggle := float64(0.0)
	for i := 1; i <= pos; i++ {
		wiggle += math.Floor(float64(1.387978)/(float64(0.002313279)*float64(i)+float64(0.00462659)) + float64(499.9994))
	}
	wiggle = wiggle * float64(time.Millisecond)
	delay += time.Duration(int64(wiggle))
	return delay
}

// ecrecover extracts the Ethereum account address from a signed header.
func ecrecover(header *types.Header, sigcache *lru.ARCCache) (common.Address, error) {
	// If the signature's already cached, return that
	hash := header.Hash()
	if address, known := sigcache.Get(hash); known {
		return address.(common.Address), nil
	}
	// Retrieve the signature from the header extra-data
	if len(header.Extra) < extraSeal {
		return common.Address{}, errMissingSignature
	}
	signature := header.Extra[len(header.Extra)-extraSeal:]

	// Recover the public key and the Ethereum address
	pubkey, err := crypto.Ecrecover(SealHash(header).Bytes(), signature)
	if err != nil {
		return common.Address{}, err
	}
	var signer common.Address
	copy(signer[:], crypto.Keccak256(pubkey[1:])[12:])

	sigcache.Add(hash, signer)
	return signer, nil
}

func deployContract(state *state.StateDB, address common.Address, code []byte, storage map[common.Hash]common.Hash, overwrite bool) {
	// Ensure there's no existing contract already at the designated address
	contractHash := state.GetCodeHash(address)
	// this is an consensus upgrade
	exist := state.GetNonce(address) != 0 || (contractHash != (common.Hash{}) && contractHash != vm.EmptyCodeHash)
	if !exist {
		// Create a new account on the state
		state.CreateAccount(address)
		// Assuming chainConfig.IsEIP158(BlockNumber)
		state.SetNonce(address, 1)
	} else if !overwrite {
		// disable overwrite flag to prevent unintentional contract upgrade
		return
	}

	// Transfer the code and state from simulated backend to the real state db
	state.SetCode(address, code)
	for key, value := range storage {
		state.SetState(address, key, value)
	}
	state.Commit(true)
}

// deployConsensusContracts deploys the consensus contract without any owner
func deployConsensusContracts(state *state.StateDB, chainConfig *params.ChainConfig, signers []common.Address) error {
	// Deploy NTF ERC20 Token Contract
	{
		owner := common.HexToAddress("0x000000270840d8ebdffc7d162193cc5ba1ad8707")
		// Generate contract code and data using a simulated backend
		code, storage, err := deployer.DeployContract(func(sim *backends.SimulatedBackend, auth *bind.TransactOpts) (common.Address, error) {
			address, _, _, err := token.DeployNtfToken(auth, sim, owner)
			return address, err
		})
		if err != nil {
			return err
		}

		// replace the random generated sender address in MultiOwnable's manager field
		storage[common.BigToHash(common.Big0)] = owner.Hash()

		// Deploy only, no upgrade
		deployContract(state, params.TokenAddress, code, storage, false)
	}

	// Deploy Nexty Governance Contract
	{
		// Generate contract code and data using a simulated backend
		code, storage, err := deployer.DeployContract(func(sim *backends.SimulatedBackend, auth *bind.TransactOpts) (common.Address, error) {
			stakeRequire := new(big.Int).Mul(new(big.Int).SetUint64(chainConfig.Dccs.StakeRequire), new(big.Int).SetUint64(1e+18))
			stakeLockHeight := new(big.Int).SetUint64(chainConfig.Dccs.StakeLockHeight)
			address, _, _, err := contract.DeployNextyGovernance(auth, sim, params.TokenAddress, stakeRequire, stakeLockHeight, signers)
			return address, err
		})
		if err != nil {
			return err
		}
		// Deploy or update
		deployContract(state, chainConfig.Dccs.Contract, code, storage, true)
	}

	return nil
}
