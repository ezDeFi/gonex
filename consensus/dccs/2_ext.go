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

	"github.com/ethereum/go-ethereum/rlp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
)

var (
	errInvalidCrosslinkLength     = errors.New("invalid crosslink digest length")
	errInvalidSealerDigestLength  = errors.New("invalid sealer digest length")
	errInvalidSealerAddressLength = errors.New("invalid sealer address length")
	errInvalidExtraLinkLength     = errors.New("invalid extra link data length")
)

// A single byte right after extraVanity indicates the DataType, allow multiple
// structures and/or versions of RLP can be decoded from the extra bytes.
const (
	ExtendedDataTypeNone        byte = 0x00
	ExtendedDataTypeVDF         byte = 0x01
	ExtendedDataTypePrice       byte = 0x02
	ExtendedDataTypeSealerJoin  byte = 0xF0
	ExtendedDataTypeSealerLeave byte = 0xF1
	ExtendedDataTypeAnchor      byte = 0xFF
)

// ExtendedData is the data encoded in header.Extra[extraVanity:-extraSeal]
type ExtendedData struct {
	anchor *AnchorData // always comes first
	random RandomData
	price  *Price
}

func (e *ExtendedData) toExtra() []byte {
	if e == nil {
		return []byte{}
	}
	var bytes []byte
	bytes = append(bytes, e.anchor.toExtra()...)
	bytes = append(bytes, e.random.toExtra()...)
	bytes = append(bytes, e.price.toExtra()...)
	return bytes
}

func extDataFrom(extBytes []byte) (*ExtendedData, error) {
	anchorData, n, err := anchorDataFrom(extBytes)
	if err != nil {
		return nil, err
	}
	extBytes = extBytes[n:]
	randomData, n := randomDataFrom(extBytes)
	extBytes = extBytes[n:]
	price, _ := priceFrom(extBytes)
	extData := ExtendedData{
		anchor: anchorData,
		random: randomData,
		price:  price,
	}
	return &extData, nil
}

func (c *Context) getExtData(header *types.Header) (*ExtendedData, error) {
	if len(header.Extra) <= extraVanity+extraSeal {
		return nil, nil
	}
	hash := header.Hash()
	if e, ok := c.engine.extDataCache.Get(hash); ok {
		// in-memory SealingQueue found
		ed := e.(*ExtendedData)
		return ed, nil
	}
	ext, err := extDataFrom(header.Extra[extraVanity : len(header.Extra)-extraSeal])
	if err != nil || ext == nil {
		return nil, err
	}
	c.engine.extDataCache.Add(hash, ext)
	return ext, nil
}

// RandomData is the seed for sealer shuffleing.
type RandomData []byte

func randomDataFrom(extra []byte) (RandomData, int) {
	size := len(extra)
	if size < 1+randomSeedSize {
		return nil, 0
	}
	if extra[0] != ExtendedDataTypeVDF {
		return nil, 0
	}
	output := extra[1 : 1+randomSeedSize]
	return append(output[:0:0], output...), 1 + randomSeedSize
}

func (r RandomData) toExtra() []byte {
	if len(r) > 0 {
		return append([]byte{ExtendedDataTypeVDF}, r...)
	}
	return nil
}

func (c *Context) getRandomData(header *types.Header) (RandomData, error) {
	extData, err := c.getExtData(header)
	if err != nil {
		return nil, err
	}
	if extData == nil {
		return nil, nil
	}
	return extData.random, nil
}

// SealerApplication packs the sealer address and its application action.
type SealerApplication struct {
	action byte
	sealer common.Address
}

func (a *SealerApplication) isJoined() bool {
	return a.action == ExtendedDataTypeSealerJoin
}

type AnchorData struct {
	destHash      common.Hash         // hash of the anchor destination block
	sealersDigest common.Hash         // digest of the ordered active sealer for this block
	applications  []SealerApplication // sealer applications confirmed in the parent block
}

func (e *AnchorData) toExtra() []byte {
	if e == nil {
		return []byte{}
	}
	size := 1 + 2*common.HashLength
	size += (1 + common.AddressLength) * len(e.applications) // sealer applications
	extra := make([]byte, 0, size)
	extra = append(extra, ExtendedDataTypeAnchor)
	extra = append(extra, e.destHash.Bytes()...)
	extra = append(extra, e.sealersDigest.Bytes()...)
	for _, app := range e.applications {
		extra = append(extra, app.action)
		extra = append(extra, app.sealer.Bytes()...)
	}
	return extra
}

func anchorDataFrom(extra []byte) (*AnchorData, int, error) {
	size := len(extra)
	if size == 0 {
		return nil, 0, nil
	}
	if extra[0] != ExtendedDataTypeAnchor {
		return nil, 0, nil
	}
	if size < 1+2*common.HashLength {
		return nil, 0, errInvalidExtraLinkLength
	}
	anchorData := AnchorData{
		destHash:      common.BytesToHash(extra[1 : 1+common.HashLength]),
		sealersDigest: common.BytesToHash(extra[1+common.HashLength : 1+2*common.HashLength]),
	}
	for i := 1 + 2*common.HashLength; i < size; {
		kind := extra[i]
		i++
		switch kind {
		case ExtendedDataTypeSealerJoin, ExtendedDataTypeSealerLeave:
			if i+common.AddressLength > size {
				log.Error("bytesToLinkData", "extra", common.Bytes2Hex(extra), "i", i)
				return nil, i, errInvalidSealerAddressLength
			}
			anchorData.applications = append(anchorData.applications, SealerApplication{
				sealer: common.BytesToAddress(extra[i : i+common.AddressLength]),
				action: kind,
			})
			i += common.AddressLength
		default:
			return &anchorData, i - 1, nil
		}
	}
	return &anchorData, size, nil
}

func verifyAnchorBytes(extBytes []byte, expectedAnchorBytes []byte) error {
	n := len(expectedAnchorBytes)
	if len(extBytes) < n {
		return errInvalidExtendedDataLength
	}
	if n > 0 {
		// anchor data always comes first in the extended extra
		if bytes.Compare(extBytes[:n], expectedAnchorBytes) != 0 {
			return errInvalidExtendedData
		}
	}
	return nil
}

func (c *Context) getAnchorData(header *types.Header) (*AnchorData, error) {
	extData, err := c.getExtData(header)
	if err != nil {
		return nil, err
	}
	if extData == nil {
		return nil, nil
	}
	return extData.anchor, nil
}

func hasAnchorData(header *types.Header) bool {
	if len(header.Extra) <= extraVanity+extraSeal {
		return false
	}
	// link data always at the start of the extended data
	return header.Extra[extraVanity] == ExtendedDataTypeAnchor
}

// AnchorData is recorded when the block is a cross-link, which happens when either:
// + the parent block has sealer application tx(s), or
// + the new SealingQueue super-majority continuity is broken
func (c *Context) assembleAnchorExtra(parent *types.Header) ([]byte, error) {
	parentHash := parent.Hash()
	if a, ok := c.engine.anchorExtraCache.Get(parentHash); ok {
		// in-memory SealingQueue found
		ab := a.([]byte)
		return ab, nil
	}
	queue, err := c.getSealingQueue(parentHash)
	if err != nil {
		return nil, err
	}

	if parent.Number.Sign() == 0 {
		// special handling for hardfork block
		anchorData := AnchorData{
			destHash:      common.Hash{},
			sealersDigest: queue.sealersDigest(),
		}
		ext := ExtendedData{
			anchor: &anchorData,
		}
		return ext.toExtra(), nil
	}

	anchorHeader, err := c.getAnchorDest(parent)
	if err != nil {
		return nil, err
	}

	anchorData := AnchorData{
		destHash: anchorHeader.Hash(),
	}
	newAnchor := false

	apps, err := c.fetchSealerApplications(parent)
	if err != nil {
		log.Error("failed to get changed sealer from log", "parent", parent.Number, "err", err)
		return nil, err
	}
	// parent block has sealer application tx(s)
	if len(apps) > 0 {
		anchorData.applications = apps
		newAnchor = true
	}

	anchorQueue, err := c.getSealingQueue(anchorHeader.ParentHash)
	if err != nil {
		return nil, err
	}

	anchorRatio, broken := queue.commonRatio(anchorQueue)
	if broken {
		log.Info("Anchor continuity is broken", "anchor ratio", anchorRatio.RatString(), "anchor number", anchorHeader.Number)
		// anchor continuity is broken, compare the current link to anchor
		linkHeader := c.getLinkDest(parent)
		linkQueue, err := c.getSealingQueue(linkHeader.ParentHash)
		if err != nil {
			return nil, err
		}
		linkRatio, broken := queue.commonRatio(linkQueue)
		if !broken || linkRatio.Cmp(anchorRatio) > 0 {
			// link continuity is preserved, or atleast better than anchor continuity
			log.Info("Link is promoted to Anchor", "link ratio", linkRatio.RatString(), "link number", linkHeader.Number)
			anchorData.destHash = parent.MixDigest
			newAnchor = true
		} else {
			// link continuity is also broken and worse than anchor continuity
			log.Warn("Broken anchor has no better alternative", "link ratio", linkRatio.RatString())
		}
	}

	if !newAnchor {
		return nil, nil
	}

	// only calculate sealersDigest when nessesary
	anchorData.sealersDigest = queue.sealersDigest()

	anchorExtra := anchorData.toExtra()
	c.engine.anchorExtraCache.Add(parentHash, anchorExtra)
	return anchorExtra, nil
}

func (c *Context) getLinkDest(header *types.Header) *types.Header {
	if header.MixDigest == (common.Hash{}) {
		// hardfork block is linked to itself
		return header
	}
	return c.getHeaderByHash(header.MixDigest)
}

func (c *Context) getAnchorDest(header *types.Header) (*types.Header, error) {
	linkHeader := header
	if !hasAnchorData(header) {
		linkHeader = c.getLinkDest(header)
		if linkHeader == nil {
			log.Error("getLinkDest returns nil", "number", header.Number, "digest", header.MixDigest)
			return nil, errLinkHeaderMissing
		}
	}
	anchorData, err := c.getAnchorData(linkHeader)
	if err != nil {
		return nil, err
	}
	if anchorData == nil {
		// should never happen
		log.Error("getAnchorHeader returns nil", "number", header.Number)
		return nil, errors.New("getAnchorHeader returns nil")
	}
	if anchorData.destHash == (common.Hash{}) {
		// hardfork block is anchored to itself
		return linkHeader, nil
	}
	return c.getHeaderByHash(anchorData.destHash), nil
}

func (c *Context) getPrice(header *types.Header) (*Price, error) {
	extData, err := c.getExtData(header)
	if err != nil {
		return nil, err
	}
	if extData == nil {
		return nil, nil
	}
	return extData.price, nil
}

func priceFrom(extra []byte) (*Price, int) {
	size := len(extra)
	if size < 1 {
		return nil, 0
	}
	if extra[0] != ExtendedDataTypePrice {
		return nil, 0
	}
	r := bytes.NewReader(extra[1:])
	var price Price
	if err := rlp.Decode(r, &price); err != nil {
		return nil, 0
	}
	return &price, size - r.Len()
}

func (p *Price) toExtra() []byte {
	if p == nil {
		return nil
	}
	bytes, err := rlp.EncodeToBytes(p)
	if err != nil {
		log.Error("Failed to serialize price data", "err", err, "price", p.Rat().RatString())
		return nil
	}
	return append([]byte{ExtendedDataTypePrice}, bytes...)
}
