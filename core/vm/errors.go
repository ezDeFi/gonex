// Copyright 2014 The go-ethereum Authors
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

package vm

import (
	"errors"

	"github.com/ethereum/go-ethereum/params"
)

// List execution errors
var (
	ErrOutOfGas                 = errors.New(params.ErrorLogOutOfGas)
	ErrCodeStoreOutOfGas        = errors.New(params.ErrorLogCodeStoreOutOfGas)
	ErrDepth                    = errors.New(params.ErrorLogDepth)
	ErrInsufficientBalance      = errors.New(params.ErrorLogInsufficientBalance)
	ErrContractAddressCollision = errors.New(params.ErrorLogContractAddressCollision)
	ErrTraceLimitReached        = errors.New("the number of logs reached the specified limit")
	ErrNoCompatibleInterpreter  = errors.New("no compatible interpreter")

	ErrTxConfigInvalidDataLength = errors.New(params.ErrorLogTxConfigInvalidDataLength)
)
