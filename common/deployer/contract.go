package deployer

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/params"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/log"
)

type DeployCallbackFn = func(*backends.SimulatedBackend, *bind.TransactOpts) (common.Address, error)

// DeployContract deploy a smart contract to simulated chain to the target address in the StateDB
func DeployContract(deployCallback DeployCallbackFn) (code []byte, storage map[common.Hash]common.Hash, err error) {
	// Generate a new random account and a funded simulator
	prvKey, _ := crypto.GenerateKey()
	auth := bind.NewKeyedTransactor(prvKey)
	auth.GasLimit = 12345678
	sim := backends.NewSimulatedBackend(
		params.ConsensusDeploymentConfig,
		core.GenesisAlloc{
			auth.From: {
				Balance: new(big.Int).Lsh(big.NewInt(1), 256-7),
			},
		},
		auth.GasLimit)
	address, err := deployCallback(sim, auth)
	if err != nil {
		log.Error("Unable to deploy consensus contract", "error", err)
		return nil, nil, err
	}
	sim.Commit()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	code, _ = sim.CodeAt(ctx, address, nil)
	storage = make(map[common.Hash]common.Hash)
	sim.ForEachStorageAt(ctx, address, nil, func(key, val common.Hash) bool {
		log.Trace("DecodeBytes", "key", key, "val", val.Hex(), "storage value", storage[key].Hex())
		storage[key] = val
		return true
	})
	return code, storage, err
}

func CopyContractToAddress(state *state.StateDB, address common.Address, code []byte, storage map[common.Hash]common.Hash, overwrite bool) {
	// Ensure there's no existing contract already at the designated address
	contractHash := state.GetCodeHash(address)
	// this is an consensus upgrade
	exist := state.GetNonce(address) != 0 || (contractHash != (common.Hash{}) && contractHash != vm.EmptyCodeHash)
	if !exist {
		// Create a new account on the state
		state.CreateAccount(address)
		// Assuming chainConfig.IsEIP158(BlockNumber)
		state.SetNonce(address, 1)
	} else if !overwrite {
		// disable overwrite flag to prevent unintentional contract upgrade
		return
	}

	// Transfer the code and state from simulated backend to the real state db
	state.SetCode(address, code)
	for key, value := range storage {
		state.SetState(address, key, value)
	}
	state.Commit(false)
}
