// Copyright 2015 The gonex Authors
// This file is part of the gonex library.
//
// The gonex library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The gonex library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the gonex library. If not, see <http://www.gnu.org/licenses/>.

package dccs

import (
	"encoding/json"
	"errors"
	"io"
	"math/big"
	"sort"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
	"github.com/ethereum/go-ethereum/rlp"
	lru "github.com/hashicorp/golang-lru"
)

const (
	medianPriceCacheSize = 6
)

// PriceData represents the external price feeded from outside
type PriceData struct {
	Value     json.Number `json:"price"`
	Timestamp int64       `json:"timestamp"`
	Exchange  string      `json:"exchange"`
}

// PriceEngine is the price feeding and managing engine
type PriceEngine struct {
	feeder       *Feeder
	serviceURL   string
	ticker       *time.Ticker
	headerPrices *lru.Cache // header price: hash -> Price
	medianPrices *lru.Cache // calculated median price: hash -> Price
	ttl          time.Duration
}

func newPriceEngine(conf *params.DccsConfig, priceServiceURL string) *PriceEngine {
	priceSamplingInterval := time.Duration(conf.PriceSamplingInterval*conf.Period) * time.Second

	// the longest time for a price to stay valid = max(blocktime, priceSamplingInterval / 2)
	ttl := priceSamplingInterval / 2
	if ttl < time.Duration(conf.Period) {
		ttl = time.Duration(conf.Period)
	}

	e := &PriceEngine{
		feeder:     &Feeder{},
		serviceURL: priceServiceURL,
		ticker:     time.NewTicker(priceSamplingInterval / 3),
		ttl:        ttl,
	}

	var err error

	maxPriceCount := int(conf.PriceSamplingDuration / conf.PriceSamplingInterval)
	e.headerPrices, err = lru.New(maxPriceCount * 3 / 2) // add some extra buffer for values in forks
	if err != nil {
		log.Crit("Unable to create the header price cache", "pricesCount", (conf.PriceSamplingDuration / conf.PriceSamplingInterval), "error", err)
		return nil
	}

	e.medianPrices, err = lru.New(medianPriceCacheSize)
	if err != nil {
		log.Crit("Unable to create the median price cache", "medianPriceCacheSize", medianPriceCacheSize, "error", err)
	}

	go e.fetchingLoop()
	return e
}

func (e *PriceEngine) fetchingLoop() {
	for range e.ticker.C {
		e.feeder.data.Range(func(key interface{}, _ interface{}) bool {
			e.feeder.requestUpdate(key.(string), parsePriceFn)
			return true
		})
	}
}

// ByPrice sorts the price list by value
type ByPrice []*Price

func (a ByPrice) Len() int           { return len(a) }
func (a ByPrice) Less(i, j int) bool { return a[i].Rat().Cmp(a[j].Rat()) < 0 }
func (a ByPrice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

// CalcMedianPrice calculates the median price of a price block and cache it.
// TODO: optimize rolling median calculation
func (c *Context) CalcMedianPrice(number uint64) (*Price, error) {
	if len(c.engine.priceURL) == 0 {
		// price server URL not provided
		return nil, nil
	}
	if !c.engine.config.IsPriceBlock(number) {
		// not a price block
		return nil, nil
	}
	if number > c.chain.CurrentHeader().Number.Uint64() {
		return nil, errors.New("Block number too high")
	}
	header := c.getHeaderByNumber(number)
	e := c.engine.PriceEngine()
	if median, ok := e.medianPrices.Get(header.Hash()); ok {
		// cache found
		return median.(*Price), nil
	}
	samplingDuration := c.engine.config.PriceSamplingDuration
	samplingInterval := c.engine.config.PriceSamplingInterval
	cap := int(samplingDuration / samplingInterval)
	prices := make([]*Price, 0, cap)
	for n := number; n > number-samplingDuration; n -= samplingInterval {
		price := c.GetBlockPrice(n)
		if price != nil {
			prices = append(prices, price)
		}
	}
	// require atleast 2/3 of maximum price feed
	// TODO: make this configurable
	median, err := medianPrice(prices, cap*2/3)
	if err == nil && median != nil {
		// cache it for the header hash
		e.medianPrices.Add(header.Hash(), median)
	}
	return median, err
}

func medianPrice(prices []*Price, minValues int) (*Price, error) {
	count := len(prices)
	if count < minValues {
		return nil, errors.New("Not enough block with price to come to a consensus")
	}
	sort.Sort(ByPrice(prices))
	if count&1 == 1 {
		// count is odd, return the middle item
		return prices[count/2], nil
	}
	// count is even, return the average of the 2 middle items
	median := new(big.Rat).Add(prices[count/2-1].Rat(), prices[count/2].Rat())
	median.Mul(median, common.Rat1_2)
	return (*Price)(median), nil
}

// GetBlockPrice returns the price encoded in a block header extra data
func (c *Context) GetBlockPrice(number uint64) *Price {
	if len(c.engine.priceURL) == 0 {
		// price server URL not provided
		return nil
	}
	if !c.engine.config.IsPriceBlock(number) {
		// not a price block
		return nil
	}
	header := c.getHeaderByNumber(number)
	if header == nil {
		log.Error("failed to get header by number ", "number", number)
		return nil
	}
	hash := header.Hash()
	e := c.engine.PriceEngine()
	if price, ok := e.headerPrices.Get(hash); ok {
		// cache found
		return price.(*Price)
	}
	price, err := c.getPrice(header)
	if err != nil {
		log.Error("failed to get price from header", "number", number, "extra", common.Bytes2Hex(header.Extra), "err", err)
		return nil
	}
	if price != nil {
		log.Trace("Header block price", "number", number, "price", price.Rat().RatString())
	}
	e.headerPrices.Add(hash, price)
	return price
}

// CurrentPrice returns the current un-expired data fed from price service
func (e *PriceEngine) CurrentPrice() *Price {
	data := e.feeder.getCurrent(e.serviceURL)
	if data == nil {
		e.feeder.requestUpdate(e.serviceURL, parsePriceFn)
		return nil
	}
	if time.Now().Sub(data.ResponseTimestamp) > e.ttl {
		// expired data
		return nil
	}
	return data.Value.(*Price)
}

func parsePriceFn(body []byte) (*Data, error) {
	var priceData PriceData

	if err := json.Unmarshal(body, &priceData); err != nil {
		log.Error("Failed to unmarshal price json", "error", err, "body", string(body))
		return nil, err
	}

	log.Trace("PriceData", "priceData", priceData)

	price := PriceFromString(priceData.Value.String())
	if price == nil || common.Rat0.Cmp(price.Rat()) == 0 {
		log.Error("Failed to parse price value", "priceData.Value", priceData.Value)
		return nil, errors.New("Not a price value")
	}

	return &Data{
		Value:             price,
		DataTimestamp:     time.Unix(priceData.Timestamp, 0),
		ResponseTimestamp: time.Now(),
		Source:            priceData.Exchange,
	}, nil
}

// Price encoded in Rat
type Price big.Rat

// Rat returns the Price in big.Rat pointer type
func (p *Price) Rat() *big.Rat {
	return (*big.Rat)(p)
}

// EncodeRLP implements the rlp.Encoder interface.
func (p *Price) EncodeRLP(w io.Writer) error {
	a, err := rlp.EncodeToBytes(p.Rat().Num())
	if err != nil {
		return err
	}
	b, err := rlp.EncodeToBytes(p.Rat().Denom())
	if err != nil {
		return err
	}
	w.Write(a)
	w.Write(b)
	return nil
}

// DecodeRLP implements the rlp.Decoder interface.
func (p *Price) DecodeRLP(s *rlp.Stream) error {
	var a, b big.Int
	if err := s.Decode(&a); err != nil {
		return err
	}
	if err := s.Decode(&b); err != nil {
		return err
	}
	if b.Sign() == 0 {
		return errInvalidPriceData
	}
	p.Rat().SetFrac(&a, &b)
	return nil
}

// PriceFromString decodes the price string fed from feeder service
func PriceFromString(s string) *Price {
	price, ok := new(big.Rat).SetString(s)
	if !ok {
		return nil
	}
	return (*Price)(price)
}

// BlockPriceStat returns ethstats data for block price
func (c *Context) BlockPriceStat(number uint64) string {
	if !c.engine.config.IsPriceBlock(number) {
		return ""
	}
	price := c.GetBlockPrice(number)
	if price == nil {
		return "0"
	}
	return price.Rat().FloatString(4)
}

// MedianPriceStat returns ethstats data for median price
func (c *Context) MedianPriceStat(number uint64) string {
	if !c.engine.config.IsPriceBlock(number) {
		return ""
	}
	price, _ := c.CalcMedianPrice(number)
	if price == nil {
		return "0"
	}
	return price.Rat().FloatString(4)
}
