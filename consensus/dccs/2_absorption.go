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
	"fmt"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/core/vm"

	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/contracts/core/endurio"
	"github.com/ethereum/go-ethereum/contracts/core/endurio/stable"
	"github.com/ethereum/go-ethereum/contracts/core/endurio/volatile"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
)

var (
	emptyHash = common.Hash{}

	onBlockInitializedSignature = crypto.Keccak256([]byte("onBlockInitialized(uint256)"))[:4]
)

// OnBlockInitialized handles supply absorption on each block initialization
func (c *Context) OnBlockInitialized(header *types.Header, state *state.StateDB, medianPrice *Price) (types.Transactions, types.Receipts, error) {
	backend := backends.NewRealBackend(c.chain, header, state, nil) // consensus only
	target := common.Big0

	if medianPrice != nil {
		stableToken, err := stable.NewStableTokenCaller(params.StableTokenAddress, backend)
		if err != nil {
			log.Error("Failed to create StableToken contract caller", "err", err)
			return nil, nil, err
		}
		target, err = stableToken.TotalSupply(nil)
		if err != nil {
			log.Error("Failed to get stable token supply", "err", err)
			return nil, nil, err
		}

		target.Mul(target, medianPrice.Rat().Num())
		target.Div(target, medianPrice.Rat().Denom())
	}

	volatileToken, err := volatile.NewVolatileTokenCaller(params.VolatileTokenAddress, backend)
	if err != nil {
		log.Error("Failed to create VolatileToken contract caller", "err", err)
		return nil, nil, err
	}
	supply, err := volatileToken.TotalSupply(nil)
	if err != nil {
		log.Error("Failed to get volatile token supply", "err", err)
		return nil, nil, err
	}

	log.Trace("VolatileToken supply before", "supply", supply)

	tx := types.NewTransaction(
		state.GetNonce(common.Address{}), // nonce
		params.SeigniorageAddress,        // to
		common.Big0,                      // amount
		math.MaxUint64,                   // gasLimit
		common.Big0,                      // gasPrice
		append(onBlockInitializedSignature, common.BigToHash(target).Bytes()...), // data
	)

	gasPool := core.GasPool(math.MaxUint64)
	var gasUsed uint64

	snap := state.Snapshot()

	state.Prepare(tx.Hash(), emptyHash, 0)
	receipt, err := core.ApplyTransaction(c.chain.Config(), c.chain, &header.Coinbase, &gasPool, state, header, tx, &gasUsed,
		vm.Config{
			Signer: types.ConsensusSigner{
				Address: common.Address{},
			}})

	if err != nil {
		state.RevertToSnapshot(snap)
		log.Error("Failed to execute Seigniorage.OnBlockInitialized", "err", err)
		return nil, nil, err
	}

	newSupply, _ := volatileToken.TotalSupply(nil)
	if newSupply.Cmp(supply) != 0 {
		log.Trace("VolatileToken supply after", "new supply", newSupply, "change", new(big.Int).Sub(newSupply, supply))
		stateObject := state.GetOrNewStateObject(params.VolatileTokenAddress)
		stateObject.SetBalance(newSupply)
		stateObject.CommitTrie(state.Database())
	}

	return types.Transactions{tx}, types.Receipts{receipt}, nil
}

// AbsorbedStat returns ethstats data for stablecoin supply absorbed by the block
func AbsorbedStat(chain consensus.ChainReader, number uint64) string {
	if number <= 0 {
		return ""
	}
	header := chain.GetHeaderByNumber(number - 1)
	if header == nil {
		return "No header at " + fmt.Sprint(number)
	}
	state, err := chain.StateAt(header.Root)
	if err != nil {
		return err.Error()
	}
	if state == nil {
		return "No state at " + fmt.Sprint(number)
	}
	oldSupply, err := GetStableTokenSupply(chain, header, state)
	if err != nil {
		return err.Error()
	}
	if oldSupply == nil {
		return "No Old Supply"
	}

	header = chain.GetHeaderByNumber(number)
	state, err = chain.StateAt(header.Root)
	if err != nil {
		return err.Error()
	}
	if state == nil {
		return "No state at " + fmt.Sprint(number)
	}
	supply, err := GetStableTokenSupply(chain, header, state)
	if err != nil {
		return err.Error()
	}
	if supply == nil {
		return "No New Supply"
	}
	absorbed := supply.Sub(supply, oldSupply)
	f, _ := new(big.Rat).SetFrac(absorbed, common.Big1e18).Float64()
	return fmt.Sprint(f)
}

// RemainToAbsorbStat returns ethstats data for stablecoin supply remain to absorb
func RemainToAbsorbStat(chain consensus.ChainReader, number uint64) string {
	header := chain.GetHeaderByNumber(number)
	if header == nil {
		return "No Header"
	}
	state, err := chain.StateAt(chain.GetHeaderByNumber(number).Root)
	if err != nil {
		return "Failed to get chain state: " + err.Error()
	}
	if state == nil {
		return "No state at " + fmt.Sprint(number)
	}
	// Random key to make sure no one has any special right
	backend := backends.NewRealBackend(chain, header, state, nil)

	caller, err := endurio.NewSeigniorageCaller(params.SeigniorageAddress, backend)
	if err != nil {
		return "Failed to create Seigniorage caller: " + err.Error()
	}
	hasAbsorption, remain, err := caller.GetRemainToAbsorb(nil)
	if err != nil {
		return "Failed to call Seigniorage.GetRemainToAbsorb: " + err.Error()
	}
	if !hasAbsorption || remain == nil {
		return "0"
	}
	f, _ := new(big.Rat).SetFrac(remain, common.Big1e18).Float64()
	return fmt.Sprint(f)
}

// StableSupplyStat returns ethstats data for stablecoin supply
func StableSupplyStat(chain consensus.ChainReader, number uint64) string {
	header := chain.GetHeaderByNumber(number)
	if header == nil {
		return "No Header"
	}
	state, err := chain.StateAt(chain.GetHeaderByNumber(number).Root)
	if err != nil {
		return err.Error()
	}
	if state == nil {
		return "No state at " + fmt.Sprint(number)
	}
	supply, err := GetStableTokenSupply(chain, header, state)
	if err != nil {
		return err.Error()
	}
	if supply == nil {
		return "No Supply"
	}
	f, _ := new(big.Rat).SetFrac(supply, common.Big1e18).Float64()
	return fmt.Sprint(f)
}

// GetStableTokenSupply returns the current supply of the stable token in the stateDB
func GetStableTokenSupply(chain consensus.ChainReader, header *types.Header, state *state.StateDB) (*big.Int, error) {
	backend := backends.NewRealBackend(chain, header, state, nil)
	caller, err := stable.NewStableTokenCaller(params.StableTokenAddress, backend)
	if err != nil {
		return nil, err
	}
	return caller.TotalSupply(nil)
}
