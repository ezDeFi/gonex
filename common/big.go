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

package common

import "math/big"

// Common big integers often used
var (
	Big1   = big.NewInt(1)
	Big2   = big.NewInt(2)
	Big3   = big.NewInt(3)
	Big0   = big.NewInt(0)
	Big32  = big.NewInt(32)
	Big256 = big.NewInt(256)
	Big257 = big.NewInt(257)

	Big1000 = big.NewInt(1000)
	Big1e12 = big.NewInt(1e12)
	Big1e18 = big.NewInt(1e18)
	Big1e24 = new(big.Int).Mul(Big1e12, Big1e12)

	Big2e14 = big.NewInt(2e14)
)

var (
	Rat0   = big.NewRat(0, 1)
	Rat1   = big.NewRat(1, 1)
	Rat2   = big.NewRat(2, 1)
	Rat1_2 = big.NewRat(1, 2)
	Rat1_4 = big.NewRat(1, 4)

	RatNeg1_2 = big.NewRat(-1, 2)
)
