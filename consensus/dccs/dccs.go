// Copyright 2017 The go-ethereum Authors
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
	"io"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethdb"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	lru "github.com/hashicorp/golang-lru"
	"golang.org/x/crypto/sha3"
)

const (
	inmemorySnapshots  = 128  // Number of recent vote snapshots to keep in memory
	inmemorySignatures = 4096 // Number of recent block signatures to keep in memory
)

// Dccs proof-of-foundation protocol constants.
var (
	epochLength = uint64(30000) // Default number of blocks after which to checkpoint and reset the pending votes

	extraVanity = 32 // Fixed number of extra-data prefix bytes reserved for signer vanity
	extraSeal   = 65 // Fixed number of extra-data suffix bytes reserved for signer seal

	uncleHash = types.CalcUncleHash(nil) // Always Keccak256(RLP([])) as uncles are meaningless outside of PoW.
)

// Various error messages to mark blocks invalid. These should be private to
// prevent engine specific errors from being referenced in the remainder of the
// codebase, inherently breaking if the engine is swapped out. Please put common
// error types into the consensus package.
var (
	// errUnknownBlock is returned when the list of signers is requested for a block
	// that is not part of the local blockchain.
	errUnknownBlock = errors.New("unknown block")

	// errInvalidCheckpointBeneficiary is returned if a checkpoint/epoch transition
	// block has a beneficiary set to non-zeroes.
	errInvalidCheckpointBeneficiary = errors.New("beneficiary in checkpoint block non-zero")

	// errInvalidVote is returned if a nonce value is something else that the two
	// allowed constants of 0x00..0 or 0xff..f.
	errInvalidVote = errors.New("vote nonce not 0x00..0 or 0xff..f")

	// errInvalidCheckpointVote is returned if a checkpoint/epoch transition block
	// has a vote nonce set to non-zeroes.
	errInvalidCheckpointVote = errors.New("vote nonce in checkpoint block non-zero")

	// errMissingVanity is returned if a block's extra-data section is shorter than
	// 32 bytes, which is required to store the signer vanity.
	errMissingVanity = errors.New("extra-data 32 byte vanity prefix missing")

	// errMissingSignature is returned if a block's extra-data section doesn't seem
	// to contain a 65 byte secp256k1 signature.
	errMissingSignature = errors.New("extra-data 65 byte suffix signature missing")

	// errExtraSigners is returned if non-checkpoint block contain signer data in
	// their extra-data fields.
	errExtraSigners = errors.New("non-checkpoint block contains extra signer list")

	// errInvalidCheckpointSigners is returned if a checkpoint block contains an
	// invalid list of signers (i.e. non divisible by 20 bytes, or not the correct
	// ones).
	errInvalidCheckpointSigners = errors.New("invalid signer list on checkpoint block")

	// errInvalidMixDigest is returned if a block's mix digest is non-zero.
	errInvalidMixDigest = errors.New("non-zero mix digest")

	// errInvalidUncleHash is returned if a block contains an non-empty uncle list.
	errInvalidUncleHash = errors.New("non empty uncle hash")

	// errInvalidDifficulty is returned if the difficulty of a block is not either
	// of 1 or 2, or if the value does not match the turn of the signer.
	errInvalidDifficulty = errors.New("invalid difficulty")

	// ErrInvalidTimestamp is returned if the timestamp of a block is lower than
	// the previous block's timestamp + the minimum block period.
	ErrInvalidTimestamp = errors.New("invalid timestamp")

	// errInvalidVotingChain is returned if an authorization list is attempted to
	// be modified via out-of-range or non-contiguous headers.
	errInvalidVotingChain = errors.New("invalid voting chain")

	// errWaitTransactions is returned if an empty block is attempted to be sealed
	// on an instant chain (0 second period). It's important to refuse these as the
	// block reward is zero, so an empty block just bloats the chain... fast.
	errWaitTransactions = errors.New("waiting for transactions")

	// errUnauthorizedSigner is returned if a header is signed by a non-authorized entity.
	errUnauthorizedSigner = errors.New("unauthorized signer")

	// errRecentlySigned is returned if a header is signed by an authorized entity
	// that already signed a header recently, thus is temporarily not allowed to.
	errRecentlySigned = errors.New("recently signed")
)

// SignerFn is a signer callback function to request a header to be signed by a
// backing account.
type SignerFn func(accounts.Account, string, []byte) ([]byte, error)

// Dccs is the proof-of-foundation consensus engine proposed to support the
// Ethereum testnet following the Ropsten attacks.
type Dccs struct {
	config *params.DccsConfig // Consensus engine configuration parameters
	db     ethdb.Database     // Database to store and retrieve snapshot checkpoints

	recents    *lru.ARCCache // Snapshots for recent block to speed up reorgs
	signatures *lru.ARCCache // Signatures of recent blocks to speed up mining

	proposals map[common.Address]bool // Current list of proposals we are pushing

	signer common.Address // Ethereum address of the signing key
	signFn SignerFn       // Signer function to authorize hashes with
	lock   sync.RWMutex   // Protects the signer fields
}

// New creates a Dccs proof-of-foundation consensus engine with the initial
// signers set to the ones provided by the user.
func New(config *params.DccsConfig, db ethdb.Database) *Dccs {
	// Set any missing consensus parameters to their defaults
	conf := *config
	if conf.Epoch == 0 {
		conf.Epoch = epochLength
	}
	// Allocate the snapshot caches and create the engine
	recents, _ := lru.NewARC(inmemorySnapshots)
	signatures, _ := lru.NewARC(inmemorySignatures)

	dccs := &Dccs{
		config:     &conf,
		db:         db,
		recents:    recents,
		signatures: signatures,
	}

	return dccs.init().init1()
}

// VerifyHeader checks whether a header conforms to the consensus rules.
func (d *Dccs) VerifyHeader(chain consensus.ChainReader, header *types.Header, seal bool) error {
	if chain.Config().IsThangLong(header.Number) {
		return d.verifyHeader1(chain, header, nil)
	}
	return d.verifyHeader(chain, header, nil)
}

// VerifyHeaders is similar to VerifyHeader, but verifies a batch of headers. The
// method returns a quit channel to abort the operations and a results channel to
// retrieve the async verifications (the order is that of the input slice).
func (d *Dccs) VerifyHeaders(chain consensus.ChainReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	abort := make(chan struct{})
	results := make(chan error, len(headers))

	go func() {
		for i, header := range headers {
			var err error
			if chain.Config().IsThangLong(header.Number) {
				err = d.verifyHeader1(chain, header, headers[:i])
			} else {
				err = d.verifyHeader(chain, header, headers[:i])
			}

			select {
			case <-abort:
				return
			case results <- err:
			}
		}
	}()
	return abort, results
}

// VerifyUncles implements consensus.Engine, always returning an error for any
// uncles as this consensus mechanism doesn't permit uncles.
func (d *Dccs) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {
	if len(block.Uncles()) > 0 {
		return errors.New("uncles not allowed")
	}
	return nil
}

// VerifySeal implements consensus.Engine, checking whether the signature contained
// in the header satisfies the consensus protocol requirements.
func (d *Dccs) VerifySeal(chain consensus.ChainReader, header *types.Header) error {
	if chain.Config().IsThangLong(header.Number) {
		return d.verifySeal1(chain, header, nil)
	}
	return d.verifySeal(chain, header, nil)
}

// Prepare implements consensus.Engine, preparing all the consensus fields of the
// header for running the transactions on top.
func (d *Dccs) Prepare(chain consensus.ChainReader, header *types.Header) error {
	if chain.Config().IsThangLong(header.Number) {
		return d.prepare1(chain, header)
	}
	return d.prepare(chain, header)
}

// Initialize implements the consensus.Engine
func (d *Dccs) Initialize(chain consensus.ChainReader, header *types.Header, state *state.StateDB) (types.Transactions, types.Receipts, error) {
	return nil, nil, nil
}

// Finalize implements consensus.Engine, ensuring no uncles are set, nor block
// rewards given, and returns the final block.
func (d *Dccs) Finalize(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header) {
	if chain.Config().IsThangLong(header.Number) {
		d.finalize1(chain, header, state, txs, uncles)
	}
	if chain.Config().IsThangLongPreparationBlock(header.Number) {
		// Retrieve the pre-fork signers list
		s, err := d.snapshot(chain, header.Number.Uint64()-1, header.ParentHash, nil)
		if err != nil {
			return
		}
		sigs := s.signers()
		// Deploy the contract and ininitalize it with pre-fork signers
		if deployConsensusContracts(state, chain.Config(), sigs) != nil {
			return
		}
		log.Info("Successfully deploy Nexty governance contract", "Number of sealers", len(sigs))
	}
	d.finalize(chain, header, state, txs, uncles)
}

// FinalizeAndAssemble implements consensus.Engine, ensuring no uncles are set, nor block
// rewards given, and returns the final block.
func (d *Dccs) FinalizeAndAssemble(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {
	if chain.Config().IsThangLong(header.Number) {
		return d.finalizeAndAssemble1(chain, header, state, txs, uncles, receipts)
	}
	if chain.Config().IsThangLongPreparationBlock(header.Number) {
		// Retrieve the pre-fork signers list
		s, err := d.snapshot(chain, header.Number.Uint64()-1, header.ParentHash, nil)
		if err != nil {
			return nil, err
		}
		sigs := s.signers()
		// Deploy the contract and ininitalize it with pre-fork signers
		if deployConsensusContracts(state, chain.Config(), sigs) != nil {
			return nil, errors.New("Unable to deploy Nexty governance smart contract")
		}
		log.Info("Successfully deploy Nexty governance contract", "Number of sealers", len(sigs))
	}
	return d.finalizeAndAssemble(chain, header, state, txs, uncles, receipts)
}

// Authorize injects a private key into the consensus engine to mint new blocks
// with.
func (d *Dccs) Authorize(signer common.Address, signFn SignerFn, state *state.StateDB, header *types.Header) {
	if d.config.IsThangLong(header.Number) {
		size := state.GetCodeSize(d.config.Contract)
		log.Info("smart contract size", "size", size)
		if size > 0 && state.Error() == nil {
			// Get token holder from coinbase
			index := common.BigToHash(common.Big1).String()[2:]
			coinbase := "0x000000000000000000000000" + signer.String()[2:]
			key := crypto.Keccak256Hash(hexutil.MustDecode(coinbase + index))
			result := state.GetState(d.config.Contract, key)

			if (result == common.Hash{}) {
				log.Warn("Validator is not in activation sealer set")
			}
		}
	}

	d.lock.Lock()
	defer d.lock.Unlock()

	d.signer = signer
	d.signFn = signFn
}

// Seal implements consensus.Engine, attempting to create a sealed block using
// the local signing credentials.
func (d *Dccs) Seal(chain consensus.ChainReader, block *types.Block, results chan<- *types.Block, stop <-chan struct{}) error {
	header := block.Header()
	if chain.Config().IsThangLong(header.Number) {
		return d.seal1(chain, block, results, stop)
	}
	return d.seal(chain, block, results, stop)
}

// CalcDifficulty is the difficulty adjustment algorithm. It returns the difficulty
// that a new block should have based on the previous blocks in the chain and the
// current signer.
func (d *Dccs) CalcDifficulty(chain consensus.ChainReader, time uint64, parent *types.Header) *big.Int {
	if chain.Config().IsThangLong(parent.Number) {
		snap, err := d.snapshot1(chain, parent.Number.Uint64(), parent.Hash(), nil)
		if err != nil {
			return nil
		}
		return CalcDifficulty1(snap, d.signer, parent)
	}
	snap, err := d.snapshot(chain, parent.Number.Uint64(), parent.Hash(), nil)
	if err != nil {
		return nil
	}
	return CalcDifficulty(snap, d.signer)
}

// SealHash returns the hash of a block prior to it being sealed.
func (d *Dccs) SealHash(header *types.Header) common.Hash {
	return SealHash(header)
}

// Close implements consensus.Engine. It's a noop for clique as there is are no background threads.
func (d *Dccs) Close() error {
	return nil
}

// APIs implements consensus.Engine, returning the user facing RPC API to allow
// controlling the signer voting.
func (d *Dccs) APIs(chain consensus.ChainReader) []rpc.API {
	return []rpc.API{{
		Namespace: "dccs",
		Version:   "1.1",
		Service:   &API{chain: chain, dccs: d},
		Public:    false,
	}}
}

// SealHash returns the hash of a block prior to it being sealed.
func SealHash(header *types.Header) (hash common.Hash) {
	hasher := sha3.NewLegacyKeccak256()
	encodeSigHeader(hasher, header)
	hasher.Sum(hash[:0])
	return hash
}

// DccsRLP returns the rlp bytes which needs to be signed for the proof-of-foundation
// sealing. The RLP to sign consists of the entire header apart from the 65 byte signature
// contained at the end of the extra data.
//
// Note, the method requires the extra data to be at least 65 bytes, otherwise it
// panics. This is done to avoid accidentally using both forms (signature present
// or not), which could be abused to produce different hashes for the same header.
func DccsRLP(header *types.Header) []byte {
	b := new(bytes.Buffer)
	encodeSigHeader(b, header)
	return b.Bytes()
}

func encodeSigHeader(w io.Writer, header *types.Header) {
	err := rlp.Encode(w, []interface{}{
		header.ParentHash,
		header.UncleHash,
		header.Coinbase,
		header.Root,
		header.TxHash,
		header.ReceiptHash,
		header.Bloom,
		header.Difficulty,
		header.Number,
		header.GasLimit,
		header.GasUsed,
		header.Time,
		header.Extra[:len(header.Extra)-65], // Yes, this will panic if extra is too short
		header.MixDigest,
		header.Nonce,
	})
	if err != nil {
		panic("can't encode: " + err.Error())
	}
}
