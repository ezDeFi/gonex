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
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/vm"
)

// IPayerABI is the input ABI used to generate the binding from.
const IPayerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"pay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"txGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"txMethodSig\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32[]\",\"name\":\"txMethodParams\",\"type\":\"bytes32[]\"}],\"name\":\"payment\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payGasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

var IPayerFuncSigPay = []byte{0xb3, 0xd7, 0x61, 0x88}     // pay(address,address,uint256)
var IPayerFuncSigPayment = []byte{0xdd, 0x50, 0x29, 0x5e} // payment(address,uint256,bytes4,bytes32[])
var abiIPayerFuncPay, _ = abiPayer.MethodById(IPayerFuncSigPay)
var abiIPayerFuncPayment, _ = abiPayer.MethodById(IPayerFuncSigPayment)

var abiPayer, _ = abi.JSON(strings.NewReader(IPayerABI))

type PaymentContext struct {
	evm           *vm.EVM
	payerContract common.Address
	msg           Message
}

func NewPaymentContext(evm *vm.EVM, msg Message) *PaymentContext {
	key := common.BytesToHash([]byte("TokenPayerCode"))
	payerContract := common.BytesToAddress(evm.StateDB.GetState(msg.From(), key).Bytes())
	if payerContract == params.ZeroAddress {
		// use the default TokenPayer contract deployed by the consensus
		payerContract = params.TokenPayerAddress
	}
	return &PaymentContext{
		evm:           evm,
		payerContract: payerContract,
		msg:           msg,
	}
}

// TODO: if failed, rollback to snapshot, and pay all gasLimit
func (context *PaymentContext) Prepaid() error {
	msg := context.msg
	msgGas := msg.Gas()

	// use the msg gas limit for payment gas limit
	token, price, payGasLimit, remainGas, vmerr := context.payment(msgGas)
	if vmerr != nil {
		log.Debug("VM returned with error for token payment query", "err", vmerr)
		return vmerr
	}
	if token == nil || params.ZeroAddress == *token || price.Sign() <= 0 {
		return fmt.Errorf("token payment unavailable")
	}
	used := msgGas - remainGas
	gasToPay := new(big.Int).SetUint64(msgGas)                      // safe
	gasToPay = gasToPay.Add(gasToPay, new(big.Int).SetUint64(used)) // safe
	gasToPay = gasToPay.Add(gasToPay, payGasLimit)                  // safe
	feeToPay := gasToPay.Mul(gasToPay, msg.GasPrice())
	log.Error("============= context.Prepaid", "used", used, "gas", gasToPay, "fee", feeToPay)

	tokenToPay := new(big.Int).Lsh(feeToPay, 128)
	tokenToPay = tokenToPay.Div(tokenToPay, price)
	if tokenToPay.Sign() <= 0 {
		tokenToPay.SetUint64(1) // zero token fee is not allowed
	}
	log.Error("**************** token pre-paid", "coinbase", context.evm.Coinbase.Hex(), "token", token.Hex(), "amount", tokenToPay)
	_, vmerr = context.pay(&context.evm.Coinbase, token, tokenToPay, payGasLimit.Uint64())
	return vmerr
}

func (context *PaymentContext) payment(gas uint64) (*common.Address, *big.Int, *big.Int, uint64, error) {
	msg := context.msg
	evm := context.evm

	input, err := abiIPayerFuncPayment.Inputs.Pack(
		msg.To(),
		new(big.Int).SetUint64(msg.Gas()),
		[4]byte{},       // TODO
		[]common.Hash{}, // TODO
	)
	if err != nil {
		return nil, nil, nil, gas, err
	}

	input = append(IPayerFuncSigPayment, input...)

	ret, gas, err := evm.CallCode(vm.AccountRef(msg.From()), context.payerContract, input, gas, common.Big0)
	if err != nil {
		if len(evm.FailureReason) > 0 {
			return nil, nil, nil, gas, errors.New(evm.FailureReason)
		}
		return nil, nil, nil, gas, err
	}
	if len(ret) != 32*3 {
		return nil, nil, nil, gas, errors.New("invalid payer payment method outputs size")
	}
	outputs := map[string]interface{}{}
	if err := abiIPayerFuncPayment.Outputs.UnpackIntoMap(outputs, ret); err != nil {
		return nil, nil, nil, gas, err
	}
	token := outputs["token"].(common.Address)
	price := outputs["price"].(*big.Int)
	payGasLimit := outputs["payGasLimit"].(*big.Int)
	return &token, price, payGasLimit, gas, nil
}

func (context *PaymentContext) pay(coinbase *common.Address, token *common.Address, fee *big.Int, gas uint64) (uint64, error) {
	input, err := abiIPayerFuncPay.Inputs.Pack(
		coinbase,
		token,
		fee,
	)
	input = append(IPayerFuncSigPay, input...)
	if err != nil {
		return gas, err
	}

	evm := context.evm

	_, gas, err = evm.CallCode(vm.AccountRef(context.msg.From()), context.payerContract, input, gas, common.Big0)
	if err != nil {
		if len(context.evm.FailureReason) > 0 {
			return gas, errors.New(evm.FailureReason)
		}
		return gas, err
	}
	return gas, nil
}
