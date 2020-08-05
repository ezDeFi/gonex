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
	"errors"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// IPayerABI is the input ABI used to generate the binding from.
const IPayerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"txTo\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"txGas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"txGasPrice\",\"type\":\"uint256\"}],\"name\":\"pay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

var IPayerFuncSigPay = []byte{0xdd, 0x9a, 0x76, 0xff} // pay(address,address,uint256,uint256)

var abiPayer, _ = abi.JSON(strings.NewReader(IPayerABI))
var abiIPayerFuncPay, _ = abiPayer.MethodById(IPayerFuncSigPay)

var tokenPayerKey = common.BytesToHash([]byte("TokenPayer"))

type PaymentContext struct {
	evm      *vm.EVM
	contract common.Address
	msg      Message
}

func NewPaymentContext(evm *vm.EVM, msg Message) *PaymentContext {
	payer := evm.StateDB.GetState(msg.From(), tokenPayerKey)
	contract := common.BytesToAddress(payer[:20])
	if contract == params.ZeroAddress {
		contract = params.TokenPayerAddress
	}
	return &PaymentContext{
		evm:      evm,
		contract: contract,
		msg:      msg,
	}
}

func (context *PaymentContext) Pay(gas uint64) (uint64, error) {
	msg := context.msg
	evm := context.evm

	input, err := abiIPayerFuncPay.Inputs.Pack(
		evm.Coinbase,
		msg.To(),
		new(big.Int).SetUint64(msg.Gas()),
		msg.GasPrice())
	if err != nil {
		return gas, err
	}
	input = append(IPayerFuncSigPay, input...)

	log.Error("++Pay", "before", gas)

	ret, gas, err := evm.CallCode(
		vm.AccountRef(msg.From()),
		context.contract,
		input,
		gas,
		common.Big0)
	log.Error("++Pay", "after", gas)
	if err == nil {
		// log.Error("================", "payment gas used", context.gas.Uint64()-remain)
		return gas, nil
	}

	if len(evm.FailureReason) > 0 {
		return gas, errors.New(evm.FailureReason)
	}
	// provide extra user friendly revert error
	reason := params.GetSolidityRevertMessage(ret)
	if len(reason) > 0 {
		return gas, errors.New(reason)
	}
	return gas, err
}
