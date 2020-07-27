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
	"fmt"
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
	"github.com/ethereum/go-ethereum/contracts/nexty/endurio"
	"github.com/ethereum/go-ethereum/contracts/nexty/endurio/stable"
	"github.com/ethereum/go-ethereum/contracts/nexty/endurio/volatile"
	"github.com/ethereum/go-ethereum/contracts/nexty/governance"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vdf"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	lru "github.com/hashicorp/golang-lru"
)

const (
	inmemorySealingQueues = 16
	inmemoryExtDatas      = 16
	inmemoryAnchorExtras  = 16
	randomSeedSize        = 32
)

var (
	errInvalidNonce              = errors.New("Invalid block nonce as distant from the last seed block")
	errUnknownPreviousSealer     = errors.New("Unknown previous block sealer")
	errInvalidExtendedDataLength = errors.New("Invalid extended data length")
	errInvalidExtendedData       = errors.New("Extended data does not matches")
	errInvalidRandomData         = errors.New("Invalid random data in extra data")
	errInvalidRandomDataSize     = errors.New("Invalid random data size from relayer")
	errRandomSeedHeaderMissing   = errors.New("Random seed header missing")
	errLinkHeaderMissing         = errors.New("Cross-linked header missing")

	errInvalidPriceData  = errors.New("price block contains invalid price value")
	errUnexpectPriceData = errors.New("non-price block contains price value")
	errNoAccountManager  = errors.New("no account manager")
	errNoSignFn          = errors.New("no SignData function for account")
	errNoAccountUnlocked = errors.New("no account unlocked")
	errNoEligibleAccount = errors.New("no eligible account for sealing")
)

// Init the second hardfork of DCCS consensus
func (d *Dccs) init2() *Dccs {
	// d.init1()
	d.sealingQueueCache, _ = lru.NewARC(inmemorySealingQueues)
	d.extDataCache, _ = lru.NewARC(inmemoryExtDatas)
	d.anchorExtraCache, _ = lru.NewARC(inmemoryAnchorExtras)
	d.queueShuffler = vdf.NewDelayer(d.vdfGen, randomSeedSize)
	return d
}

// Context represents the context of a consensus request.
type Context struct {
	head    *types.Header         // the verifying header, nil for preparation
	parents []*types.Header       // the previous headers are being parallel verified, empty for preparation
	chain   consensus.ChainReader // the underlining chain
	engine  *Dccs                 // shared config and caches
}

func NewContext(engine *Dccs, chain consensus.ChainReader) *Context {
	context := Context{
		chain:  chain,
		engine: engine,
	}
	return &context
}

// PriceEngine creates and returns the PriceEngine singleton instance
func (d *Dccs) PriceEngine() *PriceEngine {
	d.priceEngineOnce.Do(func() {
		d.priceEngine = newPriceEngine(d.config, d.priceURL)
	})
	return d.priceEngine
}

// verifyHeader2 checks whether a header conforms to the consensus rules.The
// caller may optionally pass in a batch of parents (ascending order) to avoid
// looking those up from the database. This is useful for concurrently verifying
// a batch of new headers.
func (c *Context) verifyHeader2(seal bool) error {
	header := c.head
	if header.Number == nil {
		return errUnknownBlock
	}
	number := header.Number.Uint64()

	// Don't waste time checking blocks from the future
	if header.Time > uint64(time.Now().Unix()) {
		return consensus.ErrFutureBlock
	}

	// Check that the extra-data contains both the vanity and signature
	if len(header.Extra) < extraVanity {
		return errMissingVanity
	}
	if len(header.Extra) < extraVanity+extraSeal {
		return errMissingSignature
	}

	// Don't verify the UncleHash to allow possible velvet upgrade later
	// // Ensure that the block doesn't contain any uncles which are meaningless in Dccs
	// if header.UncleHash != types.EmptyUncleHash {
	// 	return errInvalidUncleHash
	// }

	// Ensure that the block's difficulty is meaningful (may not be correct at this point)
	if number > 0 {
		if header.Difficulty == nil {
			return errInvalidDifficulty
		}
	}
	// If all checks passed, validate any special fields for hard forks
	if err := misc.VerifyForkHashes(c.chain.Config(), header, false); err != nil {
		return err
	}

	// All basic checks passed, verify cascading fields
	err := c.verifyCascadingFields2(seal)

	// Ignore missing random seed error while fast sync until the next random output available
	if len(c.parents) > 0 && err == errRandomSeedHeaderMissing {
		log.Warn("☢ Random seed header missing, continue at your own risk")
		return nil
	}
	return err
}

func (c *Context) getBlockNonce(parent *types.Header) types.BlockNonce {
	if parent.Number.Sign() == 0 {
		return types.BlockNonce{}
	}
	return types.EncodeNonce(parent.Nonce.Uint64() + 1)
}

// verifyCascadingFields2 verifies all the header fields that are not standalone,
// rather depend on a batch of previous headers. The caller may optionally pass
// in a batch of parents (ascending order) to avoid looking those up from the
// database. This is useful for concurrently verifying a batch of new headers.
func (c *Context) verifyCascadingFields2(seal bool) error {
	header := c.head
	// The genesis block is the always valid dead-end
	number := header.Number.Uint64()
	if number == 0 {
		return nil
	}

	// Ensure that the block's timestamp isn't too close to it's parent
	parent := c.getHeader(header.ParentHash, number-1)
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}
	if parent.Time+c.engine.config.Period > header.Time {
		return ErrInvalidTimestamp
	}

	// verify the cross-link reference to the last sealer application block
	var expectedMixDigest common.Hash
	if common.Big1.Cmp(header.Number) == 0 {
		expectedMixDigest = common.Hash{}
	} else if hasAnchorData(parent) {
		expectedMixDigest = parent.Hash()
	} else {
		expectedMixDigest = parent.MixDigest
	}
	if header.MixDigest != expectedMixDigest {
		return fmt.Errorf("invalid cross-link digest: number=%v, want=%v, have=%v", header.Number, expectedMixDigest, header.MixDigest)
	}

	if !seal && !hasAnchorData(header) {
		// only do full verfication for anchor and cross-link block while fast-syncing
		return nil
	}

	// Verify the extended extra data in header.Extra
	expectedAnchorBytes, err := c.assembleAnchorExtra(parent)
	if err != nil {
		return err
	}
	extBytes := header.Extra[extraVanity : len(header.Extra)-extraSeal]
	err = verifyAnchorBytes(extBytes, expectedAnchorBytes)
	if err != nil {
		return err
	}
	extBytes = extBytes[len(expectedAnchorBytes):]
	randomData, n := randomDataFrom(extBytes)

	// reset the seed block ref on valid vdf output
	nonce := types.BlockNonce{}
	if len(randomData) > 0 {
		if len(randomData) != randomSeedSize {
			return errInvalidRandomDataSize
		}
		// Verify VDF ouput here
		seedHeader := c.getChainRandomHeader(parent)
		if seedHeader == nil {
			// accept the random output if the input header is missing (our fault, not their)
			log.Error(errRandomSeedHeaderMissing.Error())
		} else {
			input := seedHeader.Hash()
			if !c.engine.queueShuffler.Verify(input[:], randomData, c.engine.config.RandomSeedIteration) {
				return errInvalidRandomData
			}
		}
		log.Info("New random data received", "random data", common.Bytes2Hex(randomData))
	} else {
		nonce = c.getBlockNonce(parent)
	}
	// verify the seed block ref
	if header.Nonce != nonce {
		log.Error("invalid nonce as seed ref", "want", nonce, "have", header.Nonce)
		return errInvalidNonce
	}

	extBytes = extBytes[n:]
	price, _ := priceFrom(extBytes)

	if c.engine.config.IsPriceBlock(number) {
		if price == nil {
			log.Warn("Missing price data in block", "number", number)
		} else if price.Rat().Cmp(common.Rat0) <= 0 {
			log.Error("Invalid price data in block", "number", number, "price", price.Rat().RatString())
			return errInvalidPriceData
		} else {
			log.Info("Block price data found", "number", number, "price", price.Rat().RatString())
		}
	} else if price != nil {
		return errUnexpectPriceData
	}

	// TODO: cache the extData here

	// All basic checks passed, verify the seal and return
	return c.verifySeal2()
}

// verifySeal2 checks whether the signature contained in the header satisfies the
// consensus protocol requirements. The method accepts an optional list of parent
// headers that aren't yet part of the local blockchain to generate the snapshots
// from.
func (c *Context) verifySeal2() error {
	header := c.head

	// Verifying the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return errUnknownBlock
	}

	// Resolve the authorization key and check against signers
	signer, err := c.ecrecover(header)
	if err != nil {
		return err
	}

	signerDifficulty, err := c.difficulty2(signer)
	if err != nil {
		return err
	}

	// Ensure that the difficulty corresponds to the turn-ness of the signer
	if header.Difficulty.Uint64() != signerDifficulty {
		return errInvalidDifficulty
	}
	return nil
}

func (c *Context) difficulty2(signer common.Address) (uint64, error) {
	// Retrieve the sealing queue verify this header
	queue, err := c.getSealingQueue(c.head.ParentHash)
	if err != nil {
		return 0, err
	}

	if !queue.isActive(signer) {
		return 0, errUnauthorizedSigner
	}
	if queue.isRecentlySigned(signer) {
		return 0, errRecentlySigned
	}

	// Ensure that the difficulty corresponds to the turn-ness of the signer
	return queue.difficulty(signer), nil
}

// Author implements consensus.Engine, returning the Ethereum address recovered
// from the signature in the header's extra-data section.
func (d *Dccs) Author(header *types.Header) (common.Address, error) {
	return ecrecover(header, d.signatures)
}

func (c *Context) getHeader(hash common.Hash, number uint64) *types.Header {
	header := c.getHeaderByNumber(number)
	if header != nil && header.Hash() == hash {
		return header
	}
	header = c.getHeaderByHash(hash)
	if header != nil && header.Number.Uint64() == number {
		return header
	}
	return nil
}

// getHeaderByNumber returns either:
// + the context head, if number == head.Number
// + the header in parents if available (nessesary for batch headers processing)
// + chain.GetHeaderByNumber(number), if all else fail
func (c *Context) getHeaderByNumber(number uint64) *types.Header {
	var headerNumber uint64
	if c.head != nil {
		headerNumber = c.head.Number.Uint64()
		if number == headerNumber {
			return c.head
		}
		if number > headerNumber {
			return c.chain.GetHeaderByNumber(number)
		}
	} else {
		if len(c.parents) == 0 {
			return c.chain.GetHeaderByNumber(number)
		}
		headerNumber = c.parents[len(c.parents)-1].Number.Uint64() + 1
	}
	idx := len(c.parents) - int(headerNumber) + int(number)
	if idx >= 0 {
		header := c.parents[idx]
		if header.Number.Uint64() == number {
			return header
		}
		log.Error("invalid parrent number", "expected", number, "actual", header.Number)
	}
	return c.chain.GetHeaderByNumber(number)
}

// getHeaderByHash returns either:
// + the input header, if hash == header.Hash()
// + chain.GetHeaderByHash(hash) if available
// + the header in parents if available (nessesary for batch headers processing)
func (c *Context) getHeaderByHash(hash common.Hash) *types.Header {
	if c.head != nil && c.head.Hash() == hash {
		return c.head
	}
	if h := c.chain.GetHeaderByHash(hash); h != nil {
		return h
	}
	for _, parent := range c.parents {
		if parent.Hash() == hash {
			return parent
		}
	}
	return nil
}

func (c *Context) getChainRandomHeader(parent *types.Header) *types.Header {
	distance := parent.Nonce.Uint64()
	if distance >= params.CanonicalDepth*3 {
		// optimization for canonical chain
		seedNumber := parent.Number.Uint64() - distance
		return c.getHeaderByNumber(seedNumber)
	}
	// properly crawl back the non-canonical chain
	for i := distance; i > 0 && parent != nil; i-- {
		parent = c.getHeaderByHash(parent.ParentHash)
	}
	return parent
}

func (c *Context) getChainRandomSeed(parent *types.Header) (RandomData, error) {
	seedHeader := c.getChainRandomHeader(parent)
	if seedHeader == nil {
		return nil, errRandomSeedHeaderMissing
	}
	if common.Big1.Cmp(seedHeader.Number) == 0 {
		// use the sealer digest for hardfork block
		anchorData, err := c.getAnchorData(seedHeader)
		if err != nil {
			return nil, err
		}
		return anchorData.sealersDigest[:], nil
	}
	return c.getRandomData(seedHeader)
}

// prepareBeneficiary2 gets the beneficiary of signer from smart contract and
// set to header's coinbase to give sealing reward later.
// + check the contract of current state first
// + trace back the previous header for 'just left sealer'
// + if all else fails, the sealer address is kept as reward beneficiary
func (c *Context) prepareBeneficiary2(header *types.Header) {
	index := common.BigToHash(common.Big1).String()[2:]
	sealer := "0x000000000000000000000000" + header.Coinbase.String()[2:]
	key := crypto.Keccak256Hash(hexutil.MustDecode(sealer + index))

	number := header.Number.Uint64()

	// try the current active state first
	state, err := c.chain.State()
	if err == nil && state != nil {
		hash := state.GetState(params.GovernanceAddress, key)
		if (hash != common.Hash{}) {
			header.Coinbase = common.HexToAddress(hash.Hex())
			return
		}
	}

	// scan the previous signed blocks
	for n := number - 1; n >= number-c.engine.config.LeakDuration; n-- {
		if n < 1 {
			break
		}
		h := c.chain.GetHeaderByNumber(n)
		if h == nil {
			break
		}
		s, err := c.ecrecover(h)
		if err != nil {
			log.Error("Unable to recover signature", "err", err)
			return
		}
		if s == header.Coinbase {
			// found the previous sealed block
			header.Coinbase = h.Coinbase
			return
		}
	}
}

// prepare2 implements consensus.Engine, preparing all the consensus fields of the
// header for running the transactions on top.
func (c *Context) prepare2(header *types.Header) error {
	number := header.Number.Uint64()
	parent := c.chain.GetHeader(header.ParentHash, number-1)
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}

	// request the new vdf calculation after each zero block nonce
	if parent.Nonce == (types.BlockNonce{}) {
		input := parent.Hash()
		log.Info("Requesting new random seed", "input", common.Bytes2Hex(input[:]))
		c.engine.queueShuffler.Request(input[:], c.engine.config.RandomSeedIteration)
	} else {
		// request the first VDF task after the node started
		c.engine.queueShufflerOnce.Do(func() {
			seedHeader := c.getChainRandomHeader(parent)
			if seedHeader == nil {
				log.Error("Failed to request the first VDF task", "err", errRandomSeedHeaderMissing)
				return
			}
			input := seedHeader.Hash()
			log.Info("First request for random seed", "input", common.Bytes2Hex(input[:]))
			c.engine.queueShuffler.Request(input[:], c.engine.config.RandomSeedIteration)
		})
	}

	if header.Coinbase == (common.Address{}) {
		// non-mining preparation
		header.Difficulty = common.Big0
	} else {
		sealer, difficulty, err := c.nextSealer(header)
		if err != nil {
			return err
		}
		if sealer == nil {
			return errNoEligibleAccount
		}
		header.Coinbase = *sealer
		c.prepareBeneficiary2(header)
		header.Difficulty = new(big.Int).SetUint64(difficulty)
	}

	// Ensure the timestamp has the correct delay
	header.Time = parent.Time + c.engine.config.Period
	if header.Time < uint64(time.Now().Unix()) {
		header.Time = uint64(time.Now().Unix())
	}

	// set the cross-link reference to the last block with anchor data
	if common.Big1.Cmp(header.Number) == 0 {
		// special handling for hardfork block
		header.MixDigest = common.Hash{}
	} else if hasAnchorData(parent) {
		header.MixDigest = parent.Hash()
	} else {
		header.MixDigest = parent.MixDigest
	}

	anchorBytes, err := c.assembleAnchorExtra(parent)
	if err != nil {
		return err
	}

	seedHeader := c.getChainRandomHeader(parent)
	if seedHeader == nil {
		return errRandomSeedHeaderMissing
	}
	input := seedHeader.Hash()
	randomData := RandomData(c.engine.queueShuffler.Peek(input[:], c.engine.config.RandomSeedIteration))
	if len(randomData) > 0 {
		log.Trace("prepare2", "vdfOutput", common.Bytes2Hex(randomData))
		if len(randomData) != randomSeedSize {
			return errInvalidRandomDataSize
		}
		header.Nonce = types.BlockNonce{}
	} else {
		// record the distant from the last sealer application block
		header.Nonce = c.getBlockNonce(parent)
	}

	var price *Price
	if c.engine.config.IsPriceBlock(number) && len(c.engine.priceURL) > 0 {
		price = c.engine.PriceEngine().CurrentPrice()
		if price == nil {
			log.Warn("No price to record in block", "number", number)
		} else if price.Rat().Cmp(common.Rat0) <= 0 {
			log.Error("Skip recording invalid price data", "price", price.Rat().RatString())
		} else {
			log.Info("Encode price to block extra", "price", price.Rat().RatString())
		}
	}

	// Prepare the start of the header.Extra
	if len(header.Extra) < extraVanity {
		header.Extra = append(header.Extra, bytes.Repeat([]byte{0x00}, extraVanity-len(header.Extra))...)
	}
	header.Extra = header.Extra[:extraVanity]
	header.Extra = append(header.Extra, anchorBytes...)
	header.Extra = append(header.Extra, randomData.toExtra()...)
	header.Extra = append(header.Extra, price.toExtra()...)
	header.Extra = append(header.Extra, make([]byte, extraSeal)...)
	return nil
}

// initialize implements the consensus.Engine
func (c *Context) initialize2(header *types.Header, state *state.StateDB) (types.Transactions, types.Receipts, error) {
	if header.Number.Cmp(common.Big1) == 0 {
		// Deploy the contract and ininitalize it with empty signer list
		if err := deployConsensusContracts(state, c.chain.Config(), []common.Address{}); err != nil {
			log.Error("Failed to deploy governance contract", "err", err)
			return nil, nil, err
		}
		log.Info("⚙ Successfully deploy governance contract")

		if err := deployCoLoaContracts(c.chain, header, state); err != nil {
			log.Error("Failed to deploy stablecoin contracts", "err", err)
			return nil, nil, err
		}
		log.Info("⚙ Successfully deploy stablecoin contracts")

		header.Root = state.IntermediateRoot(c.chain.Config().IsEIP158(header.Number))
		return nil, nil, nil
	}

	medianPrice, err := c.CalcMedianPrice(header.Number.Uint64() - 1)
	if err != nil {
		log.Trace("Failed to calculate canonical median price", "err", err, "number", header.Number)
	}

	txs, receipts, err := c.OnBlockInitialized(header, state, medianPrice)
	if err != nil {
		return nil, nil, err
	}
	header.Root = state.IntermediateRoot(c.chain.Config().IsEIP158(header.Number))
	return txs, receipts, nil
}

// finalize2 implements consensus.Engine, ensuring no uncles are set, nor block
// rewards given, and returns the final block.
func (c *Context) finalize2(header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header) {
	// Calculate any block reward for the sealer and commit the final state root
	c.engine.calculateRewards(c.chain, state, header)
	header.Root = state.IntermediateRoot(c.chain.Config().IsEIP158(header.Number))
	header.UncleHash = types.EmptyUncleHash
}

// finalizeAndAssemble2 implements consensus.Engine, ensuring no uncles are set, nor block
// rewards given, and returns the final block.
func (c *Context) finalizeAndAssemble2(header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (block *types.Block, err error) {
	// Calculate any block reward for the sealer and commit the final state root
	c.engine.calculateRewards(c.chain, state, header)
	header.Root = state.IntermediateRoot(c.chain.Config().IsEIP158(header.Number))
	header.UncleHash = types.EmptyUncleHash

	// Assemble and return the final block for sealing
	return types.NewBlock(header, txs, nil, receipts), nil
}

// deployConsensusContracts deploys the consensus contract without any owner
func deployConsensusContracts(state *state.StateDB, chainConfig *params.ChainConfig, signers []common.Address) error {
	// Deploy Nexty Governance Contract
	{
		// Generate contract code and data using a simulated backend
		code, storage, err := deployer.DeployContract(func(sim *backends.SimulatedBackend, auth *bind.TransactOpts) (common.Address, error) {
			stakeRequire := new(big.Int).Mul(new(big.Int).SetUint64(chainConfig.Dccs.StakeRequire), new(big.Int).SetUint64(1e+18))
			stakeLockHeight := new(big.Int).SetUint64(chainConfig.Dccs.StakeLockHeight)
			address, _, _, err := governance.DeployNextyGovernance(auth, sim, stakeRequire, stakeLockHeight, signers)
			return address, err
		})
		if err != nil {
			return err
		}
		// Deploy or update
		deployContract(state, params.GovernanceAddress, code, storage, true)
	}

	return nil
}

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

// calculateRewards calculate reward for block sealer
func (d *Dccs) calculateRewards(chain consensus.ChainReader, state *state.StateDB, header *types.Header) {
	number := header.Number.Uint64()
	yo := number / blockPerYear.Uint64()
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

// seal2 implements consensus.Engine, attempting to create a sealed block using
// the local signing credentials.
func (c *Context) seal2(block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {
	header := block.Header()

	// Sealing the genesis block is not supported
	number := header.Number.Uint64()
	if number == 0 {
		return errUnknownBlock
	}
	// For 0-period chains, refuse to seal empty blocks (no reward but would spin sealing)
	if c.engine.config.Period == 0 && len(block.Transactions()) == 0 {
		return errWaitTransactions
	}

	queue, err := c.getSealingQueue(header.ParentHash)
	if err != nil {
		return err
	}

	signer := queue.sealerFromDifficulty(header.Difficulty.Uint64())
	log.Error("sealerFromDifficulty", "address", signer)

	if !queue.isActive(signer) {
		return errUnauthorizedSigner
	}
	// If we're amongst the recent signers, wait for the next block
	if queue.isRecentlySigned(signer) {
		return errRecentlySigned
	}
	// Sweet, the protocol permits us to sign the block, wait for our time
	delay := time.Unix(int64(header.Time), 0).Sub(time.Now()) // nolint: gosimple
	// Find the signer offset
	offset, err := queue.offset(signer)
	if err != nil {
		return err
	}
	if offset > 0 {
		// It's not our turn explicitly to sign, delay it a bit
		wiggle := c.engine.calcDelayTimeForOffset(offset)
		delay += wiggle
		log.Trace("Out-of-turn signing requested", "wiggle", common.PrettyDuration(wiggle))
	}

	c.engine.lock.RLock()
	accManager := c.engine.accManager
	c.engine.lock.RUnlock()
	if accManager == nil {
		return errNoAccountManager
	}
	signFn := accManager.SignFunction(signer)
	if signFn == nil {
		return errNoSignFn
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

func (c *Context) ecrecover(header *types.Header) (common.Address, error) {
	return ecrecover(header, c.engine.signatures)
}

func deployCoLoaContracts(chain consensus.ChainReader, header *types.Header, state *state.StateDB) error {
	// Deploy Seigniorage Contract
	{
		// Generate contract code and data using a simulated backend
		code, storage, err := deployer.DeployContract(func(sim *backends.SimulatedBackend, auth *bind.TransactOpts) (common.Address, error) {
			address, _, _, err := endurio.DeploySeigniorage(auth, sim,
				new(big.Int).SetUint64(chain.Config().Dccs.AbsorptionDuration),
				new(big.Int).SetUint64(chain.Config().Dccs.AbsorptionExpiration),
				new(big.Int).SetUint64(chain.Config().Dccs.SlashingRate),
				new(big.Int).SetUint64(chain.Config().Dccs.LockdownExpiration),
			)
			return address, err
		})
		if err != nil {
			return err
		}

		// Deploy only, no upgrade
		deployer.CopyContractToAddress(state, params.SeigniorageAddress, code, storage, false)
		log.Info("⚙ Contract deployed successful", "contract", "Seigniorage")
	}

	// Deploy VolatileToken Contract
	{
		// Generate contract code and data using a simulated backend
		code, storage, err := deployer.DeployContract(func(sim *backends.SimulatedBackend, auth *bind.TransactOpts) (common.Address, error) {
			address, _, _, err := volatile.DeployVolatileToken(auth, sim, params.SeigniorageAddress, common.Address{}, common.Big0)
			return address, err
		})
		if err != nil {
			return err
		}

		// Deploy only, no upgrade
		deployer.CopyContractToAddress(state, params.VolatileTokenAddress, code, storage, false)
		log.Info("⚙ Contract deployed successful", "contract", "VolatileToken")
	}

	// Deploy StableToken Contract
	{
		// Generate contract code and data using a simulated backend
		code, storage, err := deployer.DeployContract(func(sim *backends.SimulatedBackend, auth *bind.TransactOpts) (common.Address, error) {
			address, _, _, err := stable.DeployStableToken(auth, sim, params.SeigniorageAddress, common.Address{}, common.Big0)
			return address, err
		})
		if err != nil {
			return err
		}

		// Deploy only, no upgrade
		deployer.CopyContractToAddress(state, params.StableTokenAddress, code, storage, false)
		log.Info("⚙ Contract deployed successful", "contract", "StableToken")
	}

	// Link them together
	{
		backend := backends.NewRealBackend(chain, header, state, nil)
		seign, err := endurio.NewSeigniorage(params.SeigniorageAddress, backend)
		if err != nil {
			log.Error("Failed to create new Seigniorage contract executor", "err", err)
			return err
		}

		consensusTransactOpts := &bind.TransactOpts{
			GasLimit: math.MaxUint64, // it's over 9000
			Signer: func(_ types.Signer, _ common.Address, tx *types.Transaction) (*types.Transaction, error) {
				return tx, nil
			},
		}

		_, err = seign.RegisterTokens(consensusTransactOpts, params.VolatileTokenAddress, params.StableTokenAddress)
		if err != nil {
			log.Error("Failed to execute Seigniorage.RegisterTokens", "err", err)
			return err
		}
		state.Commit(false)
	}
	return nil
}
