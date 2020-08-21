// Copyright 2015 The Nexty Authors
// This file is part of the gonex library.
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

package params

import (
	"bytes"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const (
	FeePrefix = "fee: " // prefix for consensus errors happen while token fee payment

	ErrorLogRevertUnknown            = "revert with no reason"
	ErrorLogInvalidOpCode            = "invalid opcode 0x%x"
	ErrorLogStackUnderflow           = "stack underflow (%d <=> %d)"
	ErrorLogStackLimitReached        = "stack limit reached %d (%d)"
	ErrorLogWriteProtection          = "write protection"
	ErrorLogOutOfGas                 = "out of gas"
	ErrorLogGasUintOverflow          = "gas overflow unsigned 64 bit integer"
	ErrorLogReturnDataOutOfBounds    = "return data out of bounds"
	ErrorLogMaxCodeSizeExceeded      = "max code size exceeded"
	ErrorLogInvalidJump              = "invalid jump destination"
	ErrorLogCodeStoreOutOfGas        = "contract creation code storage out of gas"
	ErrorLogDepth                    = "max call depth exceeded"
	ErrorLogInsufficientBalance      = "insufficient balance for transfer"
	ErrorLogContractAddressCollision = "contract address collision"
	ErrorLogTxCodeOverspent          = "tx code value limit overspent"
)

var (
	// SolidityErrorSignature is Keccak("Error(string)")
	SolidityErrorSignature = []byte{0x08, 0xc3, 0x79, 0xa0}

	// GonexErrorSignature is Keccak("Error")
	GonexErrorSignature = []byte{0xe3, 0x42, 0xda, 0xa4}
)

// GetSolidityRevertMessage handles Solidity revert and require message.
func GetSolidityRevertMessage(res []byte) string {
	if len(res) < 4+32+32 {
		return string(res)
	}
	if bytes.Compare(res[:4], SolidityErrorSignature) != 0 {
		return string(res)
	}
	res = res[4:]
	offset := int(new(big.Int).SetBytes(res[:32]).Uint64())
	res = res[32:]
	// TODO: offset is actually the size field length?
	size := int(new(big.Int).SetBytes(res[:32]).Uint64())
	if len(res) < offset+size {
		return string(res)
	}
	msg := string(res[offset : offset+size])
	return msg
}

// FromSolidityRevertMessage construct the call result from an revert message
func FromSolidityRevertMessage(msg string) []byte {
	len := len(msg)
	res := make([]byte, 0, len+4+32+32)
	res = append(res, SolidityErrorSignature...)

	offset := common.BigToHash(common.Big32)
	res = append(res, offset.Bytes()...)

	size := common.BigToHash(big.NewInt(int64(len)))
	res = append(res, size.Bytes()...)

	res = append(res, []byte(msg)...)
	return res
}
