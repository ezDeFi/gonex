// Copyright 2017 The go-ethereum Authors
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

package dccs

import (
	"context"
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/consensus"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/params"
)

// API is a user facing RPC API to allow controlling the signer and voting
// mechanisms of the proof-of-foundation scheme.
type API struct {
	chain consensus.ChainReader
	dccs  *Dccs
}

// Proposals returns the current proposals the node tries to uphold and vote on.
func (api *API) Proposals() map[common.Address]bool {
	api.dccs.lock.RLock()
	defer api.dccs.lock.RUnlock()

	proposals := make(map[common.Address]bool)
	for address, auth := range api.dccs.proposals {
		proposals[address] = auth
	}
	return proposals
}

// Propose injects a new authorization proposal that the signer will attempt to
// push through.
func (api *API) Propose(address common.Address, auth bool) {
	api.dccs.lock.Lock()
	defer api.dccs.lock.Unlock()

	api.dccs.proposals[address] = auth
}

// Discard drops a currently running proposal, stopping the signer from casting
// further votes (either for or against).
func (api *API) Discard(address common.Address) {
	api.dccs.lock.Lock()
	defer api.dccs.lock.Unlock()

	delete(api.dccs.proposals, address)
}

func (api *API) Queue(ctx context.Context, kind string) ([]string, error) {
	switch kind {
	case "leaked":
		return api.dccs.getLeakedSealers(api.chain)
	case "joined":
		return api.dccs.getJoinedSealers(api.chain)
	case "active":
		return api.dccs.getActiveSealers(api.chain)
	case "ready":
		return api.dccs.getReadySealers(api.chain)
	default:
	}
	return []string{"Unimplemented"}, nil
}

func (d *Dccs) getLeakedSealers(chain consensus.ChainReader) ([]string, error) {
	c := Context{
		head:   chain.CurrentHeader(),
		engine: d,
		chain:  chain,
	}
	// Retrieve the sealing queue verify this header
	queue, err := c.getSealingQueue(c.head.ParentHash)
	if err != nil {
		log.Error("Unable to get the sealing queue", "err", err)
		return nil, err
	}

	signers, err := d.getJoinedSigners(chain)
	if err != nil {
		log.Error("Unable to get joined signers", "err", err)
		return nil, err
	}

	sealers := []string{}
	for _, signer := range signers {
		if _, ok := queue.active[signer]; !ok {
			sealers = append(sealers, signer.String())
		}
	}
	return sealers, nil
}

func (d *Dccs) getJoinedSealers(chain consensus.ChainReader) ([]string, error) {
	signers, err := d.getJoinedSigners(chain)
	if err != nil {
		log.Error("Unable to get joined signers", "err", err)
		return nil, err
	}

	sealers := []string{}
	for _, signer := range signers {
		sealers = append(sealers, signer.String())
	}
	return sealers, nil
}

func (d *Dccs) getActiveSealers(chain consensus.ChainReader) ([]string, error) {
	c := Context{
		head:   chain.CurrentHeader(),
		engine: d,
		chain:  chain,
	}
	// Retrieve the sealing queue verify this header
	queue, err := c.getSealingQueue(c.head.ParentHash)
	if err != nil {
		log.Error("Unable to get the sealing queue", "err", err)
		return nil, err
	}

	sealers := make([]string, 0, len(queue.active))
	for signer := range queue.active {
		sealers = append(sealers, signer.String())
	}
	return sealers, nil
}

func (d *Dccs) getReadySealers(chain consensus.ChainReader) ([]string, error) {
	c := Context{
		head:   chain.CurrentHeader(),
		engine: d,
		chain:  chain,
	}
	// Retrieve the sealing queue verify this header
	queue, err := c.getSealingQueue(c.head.ParentHash)
	if err != nil {
		log.Error("Unable to get the sealing queue", "err", err)
		return nil, err
	}

	sealers := make([]string, 0, len(queue.active))
	for signer := range queue.active {
		if _, ok := queue.recent[signer]; !ok {
			sealers = append(sealers, signer.String())
		}
	}
	return sealers, nil
}

func (d *Dccs) getJoinedSigners(chain consensus.ChainReader) ([]common.Address, error) {
	state, err := chain.State()
	if state == nil || err != nil {
		log.Trace("Head state not available", "err", err)
		return nil, errors.New("Head state not available")
	}
	size := state.GetCodeSize(params.GovernanceAddress)
	if size <= 0 || state.Error() != nil {
		log.Trace("Snapshot contract state not available", "err", state.Error())
		return nil, errors.New("Governance contract not available")
	}
	index := common.BigToHash(common.Big0)
	result := state.GetState(params.GovernanceAddress, index)
	var length int64
	if (result == common.Hash{}) {
		length = 0
	} else {
		length = result.Big().Int64()
	}
	log.Trace("Total number of signer from staking smart contract", "length", length)
	signers := make([]common.Address, length)
	key := crypto.Keccak256Hash(hexutil.MustDecode(index.String()))
	for i := 0; i < len(signers); i++ {
		log.Trace("key hash", "key", key)
		singer := state.GetState(params.GovernanceAddress, key)
		signers[i] = common.HexToAddress(singer.Hex())
		key = key.Plus()
	}
	return signers, nil
}
