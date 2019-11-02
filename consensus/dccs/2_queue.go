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
	"math/big"
	"sort"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"golang.org/x/crypto/sha3"
)

type SealingQueue struct {
	hash       common.Hash                 // hash of the header
	sealer     common.Address              // sealer address of the header
	seed       []byte                      // random seed
	active     map[common.Address]struct{} // active sealers for the next block
	recent     map[common.Address]struct{} // recently signed sealers
	sorted     []common.Address            // sorted queue ([active]-[recent]+sealer)
	digest     common.Hash                 // hash(sort(active))
	sortedOnce sync.Once
	digestOnce sync.Once
}

func addressesHash(adrs []common.Address) common.Hash {
	hasher := sha3.NewLegacyKeccak256()
	for _, adr := range adrs {
		hasher.Write(adr[:])
	}
	return common.BytesToHash(hasher.Sum(nil))
}

func (q *SealingQueue) sealersDigest() common.Hash {
	q.digestOnce.Do(func() {
		active := make([]common.Address, 0, len(q.active))
		for adr := range q.active {
			active = append(active, adr)
		}
		sort.Sort(signersAscending(active))
		q.digest = addressesHash(active)
	})
	return q.digest
}

// return the common least ratio, and whether the continuity is broken
func (q *SealingQueue) commonRatio(r *SealingQueue) (*big.Rat, bool) {
	qLen := len(q.active)
	rLen := len(r.active)
	var smaller, larger map[common.Address]struct{}
	if qLen < rLen {
		smaller, larger = q.active, r.active
	} else {
		smaller, larger = r.active, q.active
	}
	common := 0
	for a := range smaller {
		if _, ok := larger[a]; ok {
			common++
		}
	}
	ratio := big.NewRat(int64(common), int64(len(larger)))
	// common must be more than super majority of both queues
	broken := common*3 <= qLen*2 || common*3 <= rLen*2
	return ratio, broken
}

// sealerShuffling implements the sort interface to allow sorting a list of addresses
type seedShuffle struct {
	seed  []byte
	queue []common.Address
}

func (s seedShuffle) Len() int      { return len(s.queue) }
func (s seedShuffle) Swap(i, j int) { s.queue[i], s.queue[j] = s.queue[j], s.queue[i] }
func (s seedShuffle) Less(i, j int) bool {
	// TODO cache these hashes
	sha := func(address common.Address, seed []byte) []byte {
		hasher := sha3.NewLegacyKeccak256()
		hasher.Write(address[:])
		hasher.Write(seed)
		return hasher.Sum(nil)
	}
	si := sha(s.queue[i], s.seed)
	sj := sha(s.queue[j], s.seed)
	return bytes.Compare(si, sj) < 0
}

// include the last sealer in for positioning even when the last sealer is just left/leaked
func (q *SealingQueue) sortedQueue() []common.Address {
	q.sortedOnce.Do(func() {
		size := len(q.active) - len(q.recent) + 1
		s := seedShuffle{
			seed:  q.seed,
			queue: make([]common.Address, 1, size),
		}
		s.queue[0] = q.sealer
		for adr := range q.active {
			_, recentlySigned := q.recent[adr]
			if !recentlySigned {
				s.queue = append(s.queue, adr)
				// TODO: pre-calculate the shuffling hash here
			}
		}
		sort.Sort(s)
		q.sorted = s.queue
	})
	return q.sorted
}

func (q *SealingQueue) isRecentlySigned(address common.Address) bool {
	_, recentlySigned := q.recent[address]
	return recentlySigned
}

func (q *SealingQueue) isActive(address common.Address) bool {
	_, active := q.active[address]
	return active
}

// position returns the possition of the signer in the active queue,
// or -1 if not in the active queue
func (q *SealingQueue) position(signer common.Address) int {
	for i, sig := range q.sortedQueue() {
		if sig == signer {
			return i
		}
	}
	return -1
}

func (q *SealingQueue) offset(signer common.Address) (int, error) {
	activeLen := len(q.active)

	if activeLen-len(q.recent) <= 1 {
		// no competition
		return 0, nil
	}

	if signer == q.sealer {
		return activeLen, errRecentlySigned
	}

	pos := q.position(signer)
	if pos < 0 {
		// unable to find the signer position
		return activeLen, errUnauthorizedSigner
	}

	prevPos := q.position(q.sealer)
	if prevPos < 0 {
		// unable to find the previous signer position: should never happen
		return activeLen, errUnknownPreviousSealer
	}

	queue := q.sortedQueue()

	offset := pos - prevPos - 1
	if offset < 0 {
		offset += len(queue)
	}

	log.Trace("SealingQueue.offset",
		"offset", offset,
		"signer position", pos,
		"previous signer position", prevPos,
		"len(queue)", len(queue),
		"len(active)", len(q.active),
		"len(recent)", len(q.recent))

	return offset, nil
}

func (q *SealingQueue) difficulty(address common.Address) uint64 {
	offset, err := q.offset(address)
	if err != nil {
		return 0
	}
	return uint64(len(q.active) - offset)
}

func (q *SealingQueue) sealerFromDifficulty(difficulty uint64) common.Address {
	queue := q.sortedQueue()
	prevPos := q.position(q.sealer)
	if prevPos < 0 {
		// unable to find the previous signer position: should never happen
		log.Error("Unknown previous signer", "sealer", q.sealer)
		return common.Address{}
	}

	offset := len(q.active) - int(difficulty)
	pos := offset + prevPos + 1
	if pos >= len(queue) {
		pos %= len(queue)
	}
	return queue[pos]
}

// recents len is MIN(lastActiveLen,activeLen)*2/3
func (c *Context) getSealingQueue(parentHash common.Hash) (*SealingQueue, error) {
	if q, ok := c.engine.sealingQueueCache.Get(parentHash); ok {
		// in-memory SealingQueue found
		queue := q.(*SealingQueue)
		return queue, nil
	}
	log.Trace("getSealingQueue", "parentHash", parentHash)
	parent := c.getHeaderByHash(parentHash)
	if parent == nil {
		return nil, errUnknownPreviousSealer
	}
	sealer, err := c.ecrecover(parent)
	if err != nil {
		return nil, err
	}
	randomData, err := c.getChainRandomSeed(parent)
	if err != nil {
		return nil, err
	}
	queue := SealingQueue{
		hash:   parentHash,
		sealer: sealer,
		seed:   randomData,
		active: map[common.Address]struct{}{},
		recent: map[common.Address]struct{}{},
	}

	// temporary queue for recents
	var recents []common.Address

	addRecent := func(sealer common.Address) {
		if _, exists := queue.recent[sealer]; !exists {
			queue.recent[sealer] = struct{}{}
			recents = append(recents, sealer)
		}
	}
	addActive := func(sealer common.Address) {
		if _, exists := queue.active[sealer]; !exists {
			queue.active[sealer] = struct{}{}
		}
	}
	remActive := func(sealer common.Address) {
		if _, exists := queue.active[sealer]; exists {
			delete(queue.active, sealer)
		}
	}

	maxDiff := parent.Difficulty.Uint64()
	n := parent.Number.Uint64()
	startLimit := n - c.engine.config.LeakDuration
	if startLimit > n { // overflown
		startLimit = 0
	}
	// scan backward from parent number for recents and difficulty
	// somewhat probabilistically optimized, fairly safe nonetheless
	for hash := parentHash; uint64(len(recents)) < maxDiff*2/3 && n > startLimit; n-- {
		log.Trace("------------", "n", n)
		header := c.getHeader(hash, n)
		if header == nil {
			log.Error("Header not found", "number", n, "hash", hash, "len(parents)", len(c.parents))
			return nil, errUnknownBlock
		}
		sealer, err := c.ecrecover(header)
		if err != nil {
			return nil, err
		}

		addActive(sealer)
		addRecent(sealer)

		// use the difficulty for total number of recently active sealer count
		diff := header.Difficulty.Uint64()
		if diff > maxDiff {
			maxDiff = diff
		}

		// next parent in the hash chain
		hash = header.ParentHash
	}

	// scan forward to collect the rest of the sealers
	endLimit := n
	for n := startLimit + 1; uint64(len(queue.active)) < maxDiff && n < endLimit; n++ {
		log.Trace("++++++++++++", "n", n)
		header := c.getHeaderByNumber(n)
		if header == nil {
			log.Error("Header not found", "number", n, "len(parents)", len(c.parents))
			return nil, errUnknownBlock
		}
		sealer, err := c.ecrecover(header)
		if err != nil {
			return nil, err
		}

		addActive(sealer)
	}

	apps, err := c.crawlSealerApplications(parent)
	if err != nil {
		return nil, err
	}
	var b strings.Builder
	for _, app := range apps {
		if app.isJoined() {
			addActive(app.sealer)
			b.WriteRune('+')
		} else {
			remActive(app.sealer)
			b.WriteRune('-')
		}
		b.WriteString(common.Bytes2Hex(app.sealer.Bytes()[:4]))
	}
	log.Trace("Sealer applications", "apps", b.String())

	// truncate the extra recents
	for i := len(queue.active) * 2 / 3; i < len(recents); i++ {
		delete(queue.recent, recents[i])
	}

	// Store found snapshot into mem-cache
	c.engine.sealingQueueCache.Add(queue.hash, &queue)
	return &queue, nil
}

// crawl back the sealer applications skip-list
func (c *Context) crawlSealerApplications(header *types.Header) ([]SealerApplication, error) {
	number := header.Number.Uint64()
	apps := []SealerApplication{}
	for header := c.getHeaderByHash(header.MixDigest); header != nil; header = c.getHeaderByHash(header.MixDigest) {
		log.Trace("crawling", "appNumber", header.Number, "appNumber.Hash", header.Hash(), "cross-link", header.MixDigest)
		if (header.MixDigest == common.Hash{}) {
			// reach the CoLoa hardfork (new genesis)
			break
		}
		appConfirmedNumber := header.Number.Uint64() + c.engine.config.ApplicationConfirmation
		// condition: appConfirmedNumber in (number-LeakDuration;number]
		if appConfirmedNumber+c.engine.config.LeakDuration <= number {
			// any applications from this would be too old
			break
		}
		if appConfirmedNumber > number {
			// not enough confirmation
			continue
		}
		if len(header.Extra) <= extraVanity+extraSeal {
			log.Error("no sealer application data in header extra", "app number", header.Number, "number", number)
			return nil, errors.New("no sealer application data in header extra")
		}
		link, _, err := anchorDataFrom(header.Extra[extraVanity : len(header.Extra)-extraSeal])
		if err != nil {
			return nil, err
		}
		if len(link.applications) > 0 {
			apps = append(link.applications, apps...)
		}
	}
	return apps, nil
}

// nextSealer returns the next account to seal for the given header
func (c *Context) nextSealer(header *types.Header) (sealer *common.Address, diff uint64, err error) {
	q, err := c.getSealingQueue(header.ParentHash)
	if err != nil {
		return nil, 0, err
	}

	c.engine.lock.RLock()
	accManager := c.engine.accManager
	c.engine.lock.RUnlock()

	if accManager == nil {
		return nil, 0, errNoAccountManager
	}
	sealers := accManager.UnlockedAccounts()
	if len(sealers) == 0 {
		return nil, 0, errNoAccountUnlocked
	}

	// short-circuit for single sealer node
	if len(sealers) == 1 {
		sealer := sealers[0]
		if !q.isActive(sealer) {
			return nil, 0, errUnauthorizedSigner
		}
		if q.isRecentlySigned(sealer) {
			return nil, 0, errRecentlySigned
		}
		return &sealer, q.difficulty(sealer), nil
	}

	prevPos := q.position(q.sealer)
	if prevPos < 0 {
		// unable to find the previous signer position: should never happen
		return nil, 0, errUnknownPreviousSealer
	}

	queue := q.sortedQueue()
	queueLen := len(queue)
	for offset := 0; offset < queueLen; offset++ {
		pos := offset + prevPos + 1
		if pos >= queueLen {
			pos %= queueLen
		}
		sealer := queue[pos]
		for _, s := range sealers {
			if sealer == s {
				diff := uint64(len(q.active) - offset)
				return &sealer, diff, nil
			}
		}
	}

	return nil, 0, errNoEligibleAccount
}
