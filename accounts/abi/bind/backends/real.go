// Copyright 2019 The gonex Authors
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

package backends

import (
	"context"
	"math/big"
	"sync"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/eth/filters"
	"github.com/ethereum/go-ethereum/event"
	"github.com/ethereum/go-ethereum/params"
)

// This nil assignment ensures compile time that RealBackend implements bind.ContractBackend.
var _ bind.ContractBackend = (*RealBackend)(nil)

// RealBackend implements bind.ContractBackend, directly modifying the blockchain state.
// Its main purpose is to allow easily inter-operating between consensus and contract code.
type RealBackend struct {
	blockchain consensus.ChainReader
	header     *types.Header
	caller     *common.Address

	mu    sync.Mutex
	state *state.StateDB // The current mutating state.

	events *filters.EventSystem // Event system for filtering log events live

	config *params.ChainConfig
}

// NewRealBackend creates a new binding backend for modifying the real blockchain state.
func NewRealBackend(chain consensus.ChainReader, header *types.Header, state *state.StateDB, caller *common.Address) *RealBackend {
	if caller == nil {
		// zero sender to make sure no one has any special permission
		caller = &common.Address{}
	}
	backend := &RealBackend{
		blockchain: chain,
		header:     header,
		caller:     caller,
		state:      state,
		config:     chain.Config(),
	}
	return backend
}

// CodeAt returns the code associated with a certain account in the blockchain.
func (b *RealBackend) CodeAt(ctx context.Context, contract common.Address, blockNumber *big.Int) ([]byte, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if blockNumber != nil && blockNumber.Cmp(b.header.Number) != 0 {
		return nil, errBlockNumberUnsupported
	}
	return b.state.GetCode(contract), nil
}

// BalanceAt returns the wei balance of a certain account in the blockchain.
func (b *RealBackend) BalanceAt(ctx context.Context, contract common.Address, blockNumber *big.Int) (*big.Int, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if blockNumber != nil && blockNumber.Cmp(b.header.Number) != 0 {
		return nil, errBlockNumberUnsupported
	}
	return b.state.GetBalance(contract), nil
}

// NonceAt returns the nonce of a certain account in the blockchain.
func (b *RealBackend) NonceAt(ctx context.Context, contract common.Address, blockNumber *big.Int) (uint64, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if blockNumber != nil && blockNumber.Cmp(b.header.Number) != 0 {
		return 0, errBlockNumberUnsupported
	}
	return b.state.GetNonce(contract), nil
}

// MRUNumberAt returns the nonce of a certain account in the blockchain.
func (b *RealBackend) MRUNumberAt(ctx context.Context, contract common.Address, blockNumber *big.Int) (uint64, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if blockNumber != nil && blockNumber.Cmp(b.header.Number) != 0 {
		return 0, errBlockNumberUnsupported
	}
	return b.state.GetMRUNumber(contract), nil
}

// ForEachStorageAt returns func to read all keys, values in the storage
func (b *RealBackend) ForEachStorageAt(ctx context.Context, contract common.Address, blockNumber *big.Int, f func(key, val common.Hash) bool) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if blockNumber != nil && blockNumber.Cmp(b.header.Number) != 0 {
		return errBlockNumberUnsupported
	}
	b.state.ForEachStorage(contract, f)
	return nil
}

// StorageAt returns the value of key in the storage of an account in the blockchain.
func (b *RealBackend) StorageAt(ctx context.Context, contract common.Address, key common.Hash, blockNumber *big.Int) ([]byte, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if blockNumber != nil && blockNumber.Cmp(b.header.Number) != 0 {
		return nil, errBlockNumberUnsupported
	}
	val := b.state.GetState(contract, key)
	return val[:], nil
}

// PendingCodeAt returns the code associated with an account in the pending state.
func (b *RealBackend) PendingCodeAt(ctx context.Context, contract common.Address) ([]byte, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.state.GetCode(contract), nil
}

// CallContract executes a contract call.
func (b *RealBackend) CallContract(ctx context.Context, call ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if blockNumber != nil && blockNumber.Cmp(b.header.Number) != 0 {
		return nil, errBlockNumberUnsupported
	}
	// make sure pure/view contract call does not modify the state
	defer b.state.RevertToSnapshot(b.state.Snapshot())
	rval, _, _, err := b.callContract(ctx, call, b.header, b.state)
	return rval, err
}

// PendingCallContract executes a contract call on the pending state.
func (b *RealBackend) PendingCallContract(ctx context.Context, call ethereum.CallMsg) ([]byte, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	rval, _, _, err := b.callContract(ctx, call, b.header, b.state)
	return rval, err
}

// PendingNonceAt implements PendingStateReader.PendingNonceAt, retrieving
// the nonce currently pending for the account.
func (b *RealBackend) PendingNonceAt(ctx context.Context, account common.Address) (uint64, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	return b.state.GetOrNewStateObject(account).Nonce(), nil
}

// SuggestGasPrice implements ContractTransactor.SuggestGasPrice.
func (b *RealBackend) SuggestGasPrice(ctx context.Context) (*big.Int, error) {
	return common.Big0, nil
}

// EstimateGas executes the requested code against the currently pending block/state and
// returns the used amount of gas.
func (b *RealBackend) EstimateGas(ctx context.Context, call ethereum.CallMsg) (uint64, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Determine the lowest and highest possible gas limits to binary search in between
	var (
		lo  uint64 = params.TxGas - 1
		hi  uint64
		cap uint64
	)
	if call.Gas >= params.TxGas {
		hi = call.Gas
	} else {
		hi = b.header.GasLimit
	}
	cap = hi

	// Create a helper to check if a gas allowance results in an executable transaction
	executable := func(gas uint64) bool {
		call.Gas = gas

		snapshot := b.state.Snapshot()
		_, _, failed, err := b.callContract(ctx, call, b.header, b.state)
		b.state.RevertToSnapshot(snapshot)

		if err != nil || failed {
			return false
		}
		return true
	}
	// Execute the binary search and hone in on an executable gas limit
	for lo+1 < hi {
		mid := (hi + lo) / 2
		if !executable(mid) {
			lo = mid
		} else {
			hi = mid
		}
	}
	// Reject the transaction as invalid if it still fails at the highest allowance
	if hi == cap {
		if !executable(hi) {
			return 0, errGasEstimationFailed
		}
	}
	return hi, nil
}

// callContract implements common code between normal and pending contract calls.
// state is modified during execution, make sure to copy it if necessary.
func (b *RealBackend) callContract(ctx context.Context, call ethereum.CallMsg, header *types.Header, statedb *state.StateDB) ([]byte, uint64, bool, error) {
	// Ensure message is initialized properly.
	if call.GasPrice == nil {
		call.GasPrice = common.Big0
	}
	if call.Gas == 0 {
		call.Gas = math.MaxUint64
	}
	if call.Value == nil {
		call.Value = common.Big0
	}
	// Execute the call.
	msg := callmsg{call}

	evmContext := core.NewEVMContext(msg, header, b.blockchain, nil)
	// Create a new environment which holds all relevant information
	// about the transaction and calling mechanisms.
	vmenv := vm.NewEVM(evmContext, statedb, b.config, vm.Config{})
	gaspool := new(core.GasPool).AddGas(math.MaxUint64)

	return core.NewStateTransition(vmenv, msg, gaspool).TransitionDb()
}

// SendTransaction updates the pending block to include the given transaction.
// It panics if the transaction is invalid.
func (b *RealBackend) SendTransaction(ctx context.Context, tx *types.Transaction) error {
	_, err := b.PendingCallContract(ctx, ethereum.CallMsg{
		From:     *b.caller,
		To:       tx.To(),
		Gas:      tx.Gas(),
		GasPrice: tx.GasPrice(),
		Value:    tx.Value(),
		Data:     tx.Data(),
	})
	return err
}

// FilterLogs executes a log filter operation, blocking during execution and
// returning all the results in one batch.
//
// TODO(karalabe): Deprecate when the subscription one can return past data too.
func (b *RealBackend) FilterLogs(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	return nil, nil
}

// SubscribeFilterLogs creates a background log filtering operation, returning a
// subscription immediately, which can be used to stream the found events.
func (b *RealBackend) SubscribeFilterLogs(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	// Subscribe to contract events
	sink := make(chan []*types.Log)

	sub, err := b.events.SubscribeLogs(query, sink)
	if err != nil {
		return nil, err
	}
	// Since we're getting logs in batches, we need to flatten them into a plain stream
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case logs := <-sink:
				for _, log := range logs {
					select {
					case ch <- *log:
					case err := <-sub.Err():
						return err
					case <-quit:
						return nil
					}
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
