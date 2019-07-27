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

package dccs

import (
	"bytes"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rlp"
	"golang.org/x/crypto/sha3"
)

// Signer contain data of signer's address and checkpoint block hash
type Signer struct {
	Hash    common.Hash
	Address common.Address
}

// signersAscendingByHash implements the sort interface to allow sorting a list of signers by hash
type signersAscendingByHash []Signer

func (sn signersAscendingByHash) Len() int { return len(sn) }
func (sn signersAscendingByHash) Less(i, j int) bool {
	h1 := rlpHash(sn[i])
	h2 := rlpHash(sn[j])
	return bytes.Compare(h1[:], h2[:]) < 0
}
func (sn signersAscendingByHash) Swap(i, j int) { sn[i], sn[j] = sn[j], sn[i] }

// init the Snapshot for the first hardfork of DCCS consensus
func (s *Snapshot) init1() *Snapshot {
	return s
}

// signers1 retrieves the list of authorized signers in hash ascending order.
func (s *Snapshot) signers1() []Signer {
	sigs := make([]Signer, 0, len(s.Signers))
	for sig := range s.Signers {
		sigs = append(sigs, Signer{Hash: s.Hash, Address: sig})
	}
	sort.Sort(signersAscendingByHash(sigs))
	return sigs
}

// inturn2 returns if a signer at a given block height is in-turn or not.
func (s *Snapshot) inturn1(signer common.Address, parent *types.Header) bool {
	offset, err := s.offset(signer, parent)
	if err != nil {
		return false
	}
	log.Trace("inturn", "offset", offset)
	return offset == 0
}

func signerPosition(signer common.Address, signers []Signer) (int, bool) {
	for i, sig := range signers {
		if sig.Address == signer {
			return i, true
		}
	}
	return -1, false
}

func (s *Snapshot) offset(signer common.Address, parent *types.Header) (int, error) {
	signers := s.signers1()
	n := len(signers)
	if n <= 1 {
		// no competition
		return 0, nil
	}

	pos, ok := signerPosition(signer, signers)
	if !ok {
		// unable to find the signer position
		return n, errUnauthorizedSigner
	}

	if parent == nil || s.config.IsCheckpoint(parent.Number.Uint64()+1) {
		// first block of an epoch, just return the rightful order
		log.Debug("the first block of an epoch")
		return pos, nil
	}

	// Resolve the last authorization key and check against signer
	prevSigner, err := ecrecover(parent, s.sigcache)
	if err != nil {
		return 0, err
	}
	prevPos, ok := signerPosition(prevSigner, signers)
	if !ok {
		// unable to find the previous signer position
		return n, errUnauthorizedSigner
	}

	offset := pos - prevPos - 1
	if offset < 0 {
		offset += n
	}

	log.Debug("offset", "signer position", pos, "previous signer position", prevPos, "len(signers)", n, "offset", offset)

	return offset, nil
}

// difficulty returns the block weight at a given block height for a signer.
// Turn-ness is the directional distant from a signer to the previous one,
// following a circular order of the signers list.
// @return maximum value = len(signers) if signer is right after the prevSigner (circularly)
// @return minimum value = 1 if the signer is right before the prevSigner (circularly)
// @return invalid value = 0 if the signer or parent signer is not on the sealer list
func (s *Snapshot) difficulty(signer common.Address, parent *types.Header) uint64 {
	offset, err := s.offset(signer, parent)
	if err != nil {
		return 0
	}

	signers := s.signers1()
	n := len(signers)

	return uint64(n - offset)
}

// rlpHash return hash of an input
func rlpHash(s Signer) (h common.Hash) {
	hasher := sha3.NewLegacyKeccak256()
	rlp.Encode(hasher, []interface{}{
		s.Address,
		s.Hash,
	})
	hasher.Sum(h[:0])
	return h
}
