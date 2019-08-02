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

package vdf

import (
	"fmt"
	"os/exec"
	"sync"

	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/log"
	"github.com/harmony-one/vdf/src/vdf_go"
)

var (
	engine     *Engine
	engineOnce sync.Once
)

// Engine defines the structure of the engine
type Engine struct {
	cli string
}

// Instance returns the singleton instance of the VDF Engine
func Instance() *Engine {
	engineOnce.Do(func() {
		engine = newEngine()
	})
	return engine
}

func newEngine() *Engine {
	cli, err := exec.LookPath("vdf-cli")
	if err != nil {
		log.Warn("vdf.newEngine", "vdf-cli", "not found")
	}
	return &Engine{cli}
}

// Verify verifies the generated output against the seed
func (e *Engine) Verify(seed, output []byte, iteration uint64, bitSize uint64) (valid bool) {
	defer func() {
		if x := recover(); x != nil {
			log.Error("vdf.Verify: verification process panic", "reason", x)
			valid = false
		}
	}()
	return vdf_go.VerifyVDF(seed, output, int(iteration), int(bitSize))
}

// Generate generates the vdf output = (y, proof)
func (e *Engine) Generate(seed []byte, iteration uint64, bitSize uint64) (output []byte, err error) {
	if len(e.cli) == 0 {
		defer func() {
			if x := recover(); x != nil {
				log.Error("vdf.Generate: generation process panic", "reason", x)
				err = fmt.Errorf("%v", x)
			}
		}()
		y, proof := vdf_go.GenerateVDF(seed, int(iteration), int(bitSize))
		return append(y, proof...), nil
	}

	cmd := exec.Command("vdf-cli",
		"-l"+string(bitSize),
		common.Bytes2Hex(seed),
		string(iteration))
	log.Trace("vdf.Generate", "seed", seed, "iteration", iteration, "output", output)
	output, err = cmd.Output()
	if err, ok := err.(*exec.ExitError); ok {
		// verification failed
		log.Trace("vdf.Generate", "error", err.Error())
		return nil, err
	}
	if err != nil {
		// sum ting wong
		log.Error("vdf.Generate", "error", err)
		return nil, err
	}
	log.Trace("vdf.Generate", "output", output)
	return common.Hex2Bytes(string(output)), nil
}
