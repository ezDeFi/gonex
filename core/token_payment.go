// Copyright 2015 The go-ethereum Authors
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

package core

import (
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/params"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// iPayerABI is the input ABI used to generate the binding from.
const iPayerABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"txTo\",\"type\":\"address\"}],\"name\":\"pay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

var abiPayer, _ = abi.JSON(strings.NewReader(iPayerABI))

// "31cbf5e3": "pay(uint256,address)"
var sigFuncPay = []byte{0x31, 0xcb, 0xf5, 0xe3}
var abiFuncPay, _ = abiPayer.MethodById(sigFuncPay)

var tokenPayerKey = common.BytesToHash([]byte("TokenPayer"))
var tokenPriceKeyPrefix = []byte("TokenPrice-")

// keccak('Transfer(address,address,uint256)')
var FuncHashTransfer = common.HexToHash("ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")

// PaymentContext stores the context for token payment
type PaymentContext struct {
	evm        *vm.EVM
	contract   common.Address
	input      []byte
	msg        Message
	paymentGas uint64 // paymentGas
}

// NewPaymentContext creates new PaymentContext
func NewPaymentContext(evm *vm.EVM, msg Message) (*PaymentContext, error) {
	payer := evm.StateDB.GetState(msg.From(), tokenPayerKey)
	contract := common.BytesToAddress(payer[:20])
	if contract == (common.Address{}) {
		contract = params.TokenPayementAddress
	}
	// payer[20:24] unused
	paymentGas := new(big.Int).SetBytes(payer[24:]).Uint64()
	if paymentGas == 0 {
		paymentGas = params.TokenPayementGas
	}

	gasToPay := new(big.Int).SetUint64(paymentGas)
	gasToPay = gasToPay.Add(gasToPay, new(big.Int).SetUint64(msg.Gas()))
	feeToPay := gasToPay.Mul(gasToPay, msg.GasPrice())

	input, err := abiFuncPay.Inputs.Pack(feeToPay, msg.To())
	if err != nil {
		return nil, err
	}
	input = append(sigFuncPay, input...)

	return &PaymentContext{
		evm:        evm,
		contract:   contract,
		input:      input,
		msg:        msg,
		paymentGas: paymentGas,
	}, nil
}

// Pay performs the token payment of the payment context
func (context *PaymentContext) Pay() ([]byte, error) {
	ret, _, err := context.evm.CallCode(
		vm.AccountRef(context.msg.From()),
		context.contract,
		context.input,
		context.paymentGas,
		common.Big0)

	return ret, err
}

// Prepare prepares the state for evm to capture the event logs
// (for tx pool)
func (context *PaymentContext) Prepare(txHash, blockHash common.Hash, txIndex int) {
	context.evm.StateDB.Prepare(txHash, blockHash, txIndex)
}

// EffectiveGasPrice returns the effective gas price (in ZD) of the pay by token tx
// (for tx pool)
func (context *PaymentContext) EffectiveGasPrice(txHash common.Hash) (*big.Int, error) {
	fee, err := context.extractPayment(txHash)
	if err != nil {
		return nil, err
	}
	// TODO: store this in context
	gasToPay := new(big.Int).SetUint64(context.paymentGas)
	gasToPay = gasToPay.Add(gasToPay, new(big.Int).SetUint64(context.msg.Gas()))

	// effective gas price = ceil (fee / gasToPay)
	price := fee.Add(fee, gasToPay)
	price = price.Sub(price, common.Big1) // + gasToPay-1 to round up the result
	price = price.Div(price, gasToPay)
	return price, nil
}

// extractPayment returns the actual paid fee (in wei) of the pay by token tx
func (context *PaymentContext) extractPayment(txHash common.Hash) (*big.Int, error) {
	events := context.evm.StateDB.GetLogs(txHash)
	for _, event := range events {
		if len(event.Topics) != 3 {
			continue
		}
		if event.Topics[0] != FuncHashTransfer {
			continue
		}
		// from := common.BytesToAddress(event.Topics[1][12:])
		to := common.BytesToAddress(event.Topics[2][12:])
		if to != context.evm.Coinbase {
			continue
		}

		amount := new(big.Int).SetBytes(event.Data)
		if amount.Sign() == 0 {
			continue
		}

		tokenPrice := context.getPrice(event.Address)
		if tokenPrice.Sign() == 0 {
			continue
		}

		// fee = amount * tokenPrice
		fee := amount.Mul(amount, tokenPrice)
		fee = fee.Div(fee, common.Big1e18)

		return fee, nil
	}
	return nil, fmt.Errorf("%v%v", params.FeePrefix, "no accepted payment log")
}

// return the tokenPrice from the state (with decimals = 18)
func (context *PaymentContext) getPrice(token common.Address) *big.Int {
	state := context.evm.StateDB
	var key [32]byte
	copy(key[:], tokenPriceKeyPrefix)
	copy(key[11:], token.Bytes())
	// key[31] = 0x0

	// TODO: check the personal price settings from local config and coinbase storage variable
	priceHash := state.GetState(params.TokenPayementAddress, common.BytesToHash(key[:]))
	tokenPrice := new(big.Int).SetBytes(priceHash.Bytes())
	return tokenPrice
}
