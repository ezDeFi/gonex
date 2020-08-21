// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package governance

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// NextyGovernanceABI is the input ABI used to generate the binding from.
const NextyGovernanceABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakeRequire\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stakeLockHeight\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sealer\",\"type\":\"address\"}],\"name\":\"Banned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sealer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sealer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"Joined\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sealer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"Left\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sealer\",\"type\":\"address\"}],\"name\":\"Unbanned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sealer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"account\",\"outputs\":[{\"internalType\":\"enumNextyGovernance.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"unlockHeight\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getCoinbase\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getUnlockHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isBanned\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isWithdrawable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"join\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"leave\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"signerCoinbase\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"signers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeLockHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeRequire\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// NextyGovernanceFuncSigs maps the 4-byte function signature to its string representation.
var NextyGovernanceFuncSigs = map[string]string{
	"73b9aa91": "account(address)",
	"d0e30db0": "deposit()",
	"f8b2cb4f": "getBalance(address)",
	"0a9a3f07": "getCoinbase(address)",
	"30ccebb5": "getStatus(address)",
	"55235d47": "getUnlockHeight(address)",
	"97f735d5": "isBanned(address)",
	"02a18484": "isWithdrawable(address)",
	"28ffe6c8": "join(address)",
	"d66d9e19": "leave()",
	"8943b62c": "signerCoinbase(address)",
	"2079fb9a": "signers(uint256)",
	"9b212d01": "stakeLockHeight()",
	"c6908739": "stakeRequire()",
	"3ccfd60b": "withdraw()",
}

// NextyGovernanceBin is the compiled bytecode used for deploying new contracts.
var NextyGovernanceBin = "0x608060405234801561001057600080fd5b506040516200111f3803806200111f8339818101604052606081101561003557600080fd5b8151602083015160408085018051915193959294830192918464010000000082111561006057600080fd5b90830190602082018581111561007557600080fd5b825186602082028301116401000000008211171561009257600080fd5b82525081516020918201928201910280838360005b838110156100bf5781810151838201526020016100a7565b5050505091909101604052505050600384905550600482905560005b81518110156102685760008282815181106100f257fe5b60209081029190910181015182546001810184556000938452919092200180546001600160a01b0319166001600160a01b03909216919091179055815182908290811061013b57fe5b60200260200101516001600084848151811061015357fe5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b031602179055508181815181106101ab57fe5b6020026020010151600260008484815181106101c357fe5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060020160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060016002600084848151811061022457fe5b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff1916600183600481111561025b57fe5b02179055506001016100db565b50505050610ea3806200027c6000396000f3fe6080604052600436106100ec5760003560e01c806373b9aa911161008a578063c690873911610059578063c690873914610379578063d0e30db01461038e578063d66d9e1914610396578063f8b2cb4f146103ab576100fb565b806373b9aa91146102805780638943b62c146102fe57806397f735d5146103315780639b212d0114610364576100fb565b806328ffe6c8116100c657806328ffe6c8146101c057806330ccebb5146101f35780633ccfd60b1461023857806355235d471461024d576100fb565b806302a18484146101005780630a9a3f07146101475780632079fb9a14610196576100fb565b366100fb576100f96103de565b005b600080fd5b34801561010c57600080fd5b506101336004803603602081101561012357600080fd5b50356001600160a01b0316610458565b604080519115158252519081900360200190f35b34801561015357600080fd5b5061017a6004803603602081101561016a57600080fd5b50356001600160a01b03166104e6565b604080516001600160a01b039092168252519081900360200190f35b3480156101a257600080fd5b5061017a600480360360208110156101b957600080fd5b5035610508565b3480156101cc57600080fd5b50610133600480360360208110156101e357600080fd5b50356001600160a01b031661052f565b3480156101ff57600080fd5b506102266004803603602081101561021657600080fd5b50356001600160a01b031661086d565b60408051918252519081900360200190f35b34801561024457600080fd5b50610133610892565b34801561025957600080fd5b506102266004803603602081101561027057600080fd5b50356001600160a01b03166109e0565b34801561028c57600080fd5b506102b3600480360360208110156102a357600080fd5b50356001600160a01b03166109fe565b604051808560048111156102c357fe5b60ff168152602001848152602001836001600160a01b03166001600160a01b0316815260200182815260200194505050505060405180910390f35b34801561030a57600080fd5b5061017a6004803603602081101561032157600080fd5b50356001600160a01b0316610a34565b34801561033d57600080fd5b506101336004803603602081101561035457600080fd5b50356001600160a01b0316610a4f565b34801561037057600080fd5b50610226610a81565b34801561038557600080fd5b50610226610a87565b6100f9610a8d565b3480156103a257600080fd5b50610133610a97565b3480156103b757600080fd5b50610226600480360360208110156103ce57600080fd5b50356001600160a01b0316610c2b565b336000908152600260205260409020600101543490610403908263ffffffff610c4916565b3360008181526002602090815260409182902060010193909355805191825291810183905281517f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4929181900390910190a150565b600060016001600160a01b03831660009081526002602052604090205460ff16600481111561048357fe5b141580156104b8575060046001600160a01b03831660009081526002602052604090205460ff1660048111156104b557fe5b14155b80156104de57506001600160a01b03821660009081526002602052604090206003015443115b90505b919050565b6001600160a01b03908116600090815260026020819052604090912001541690565b6000818154811061051557fe5b6000918252602090912001546001600160a01b0316905081565b600060043360009081526002602052604090205460ff16600481111561055157fe5b141561058e576040805162461bcd60e51b815260206004820152600760248201526603130b73732b2160cd1b604482015290519081900360640190fd5b60013360009081526002602052604090205460ff1660048111156105ae57fe5b14156105f3576040805162461bcd60e51b815260206004820152600f60248201526e030b63932b0b23c903537b4b732b21608d1b604482015290519081900360640190fd5b60035433600090815260026020526040902060010154101561064d576040805162461bcd60e51b815260206004820152600e60248201526d3737ba1032b737bab3b410373a3360911b604482015290519081900360640190fd5b6001600160a01b03808316600090815260016020526040902054839116156106b4576040805162461bcd60e51b815260206004820152601560248201527418dbda5b98985cd948185b1c9958591e481d5cd959605a1b604482015290519081900360640190fd5b6001600160a01b0381166106fd576040805162461bcd60e51b815260206004820152600b60248201526a7369676e6572207a65726f60a81b604482015290519081900360640190fd5b6001600160a01b03811630141561075b576040805162461bcd60e51b815260206004820152601760248201527f73616d6520636f6e747261637427732061646472657373000000000000000000604482015290519081900360640190fd5b6001600160a01b0381163314156107b1576040805162461bcd60e51b815260206004820152601560248201527473616d652073656e6465722773206164647265737360581b604482015290519081900360640190fd5b33600090815260026020819052604090912090810180546001600160a01b0319166001600160a01b03861617905580546001919060ff1916828002179055506001600160a01b038316600090815260016020526040902080546001600160a01b0319163317905561082183610caa565b604080513381526001600160a01b038516602082015281517f7702dccda75540ad1dca8d5276c048f4a5c0e4203f6da4be214bfb1901b203ea929181900390910190a150600192915050565b6001600160a01b0381166000908152600260205260408120546104de9060ff16610cf9565b600060043360009081526002602052604090205460ff1660048111156108b457fe5b14156108f1576040805162461bcd60e51b815260206004820152600760248201526603130b73732b2160cd1b604482015290519081900360640190fd5b6108fa33610458565b61094b576040805162461bcd60e51b815260206004820181905260248201527f756e61626c6520746f20776974686472617720617420746865206d6f6d656e74604482015290519081900360640190fd5b3360008181526002602052604080822060018101805490849055815460ff191660031790915590519092916108fc841502918491818181858888f1935050505015801561099c573d6000803e3d6000fd5b50604080513381526020810183905281517f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5929181900390910190a1600191505090565b6001600160a01b031660009081526002602052604090206003015490565b6002602081905260009182526040909120805460018201549282015460039092015460ff90911692916001600160a01b03169084565b6001602052600090815260409020546001600160a01b031681565b600060046001600160a01b03831660009081526002602052604090205460ff166004811115610a7a57fe5b1492915050565b60045481565b60035481565b610a956103de565b565b600060043360009081526002602052604090205460ff166004811115610ab957fe5b1415610af6576040805162461bcd60e51b815260206004820152600760248201526603130b73732b2160cd1b604482015290519081900360640190fd5b60013360009081526002602052604090205460ff166004811115610b1657fe5b14610b56576040805162461bcd60e51b815260206004820152600b60248201526a03737ba103537b4b732b2160ad1b604482015290519081900360640190fd5b33600090815260026020819052604090912080820180546001600160a01b03198116909155815460ff191690921790556004546001600160a01b0390911690610b9f9043610c49565b336000908152600260209081526040808320600301939093556001600160a01b0384168252600190522080546001600160a01b0319169055610be081610d72565b604080513381526001600160a01b038316602082015281517f4b9ee4dd061ba088b22898a02491f3896a4a580c6cda8783ca579ee159f8e8c5929181900390910190a1600191505090565b6001600160a01b031660009081526002602052604090206001015490565b600082820183811015610ca3576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b600080546001810182559080527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5630180546001600160a01b0319166001600160a01b0392909216919091179055565b600080826004811115610d0857fe5b1415610d16575060006104e1565b6001826004811115610d2457fe5b1415610d32575060016104e1565b6002826004811115610d4057fe5b1415610d4e575060026104e1565b6003826004811115610d5c57fe5b1415610d6a575060036104e1565b50607f919050565b60005b600054811015610e435760008181548110610d8c57fe5b6000918252602090912001546001600160a01b0383811691161415610e3b57600080546000198101908110610dbd57fe5b600091825260208220015481546001600160a01b03909116919083908110610de157fe5b6000918252602082200180546001600160a01b0319166001600160a01b039390931692909217909155805480610e1357fe5b600082815260209020810160001990810180546001600160a01b031916905501905550610e45565b600101610d75565b505b5056fea2646970667358221220a3793a23d3d057ce34fd04d2fd7c67b25f113851e06cc846748f878a0874369c64736f6c637827302e362e392d646576656c6f702e323032302e352e32372b636f6d6d69742e39663430376665300058"

// DeployNextyGovernance deploys a new Ethereum contract, binding an instance of NextyGovernance to it.
func DeployNextyGovernance(auth *bind.TransactOpts, backend bind.ContractBackend, _stakeRequire *big.Int, _stakeLockHeight *big.Int, _signers []common.Address) (common.Address, *types.Transaction, *NextyGovernance, error) {
	parsed, err := abi.JSON(strings.NewReader(NextyGovernanceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NextyGovernanceBin), backend, _stakeRequire, _stakeLockHeight, _signers)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NextyGovernance{NextyGovernanceCaller: NextyGovernanceCaller{contract: contract}, NextyGovernanceTransactor: NextyGovernanceTransactor{contract: contract}, NextyGovernanceFilterer: NextyGovernanceFilterer{contract: contract}}, nil
}

// NextyGovernance is an auto generated Go binding around an Ethereum contract.
type NextyGovernance struct {
	NextyGovernanceCaller     // Read-only binding to the contract
	NextyGovernanceTransactor // Write-only binding to the contract
	NextyGovernanceFilterer   // Log filterer for contract events
}

// NextyGovernanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type NextyGovernanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NextyGovernanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NextyGovernanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NextyGovernanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NextyGovernanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NextyGovernanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NextyGovernanceSession struct {
	Contract     *NextyGovernance  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NextyGovernanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NextyGovernanceCallerSession struct {
	Contract *NextyGovernanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// NextyGovernanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NextyGovernanceTransactorSession struct {
	Contract     *NextyGovernanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// NextyGovernanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type NextyGovernanceRaw struct {
	Contract *NextyGovernance // Generic contract binding to access the raw methods on
}

// NextyGovernanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NextyGovernanceCallerRaw struct {
	Contract *NextyGovernanceCaller // Generic read-only contract binding to access the raw methods on
}

// NextyGovernanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NextyGovernanceTransactorRaw struct {
	Contract *NextyGovernanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNextyGovernance creates a new instance of NextyGovernance, bound to a specific deployed contract.
func NewNextyGovernance(address common.Address, backend bind.ContractBackend) (*NextyGovernance, error) {
	contract, err := bindNextyGovernance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NextyGovernance{NextyGovernanceCaller: NextyGovernanceCaller{contract: contract}, NextyGovernanceTransactor: NextyGovernanceTransactor{contract: contract}, NextyGovernanceFilterer: NextyGovernanceFilterer{contract: contract}}, nil
}

// NewNextyGovernanceCaller creates a new read-only instance of NextyGovernance, bound to a specific deployed contract.
func NewNextyGovernanceCaller(address common.Address, caller bind.ContractCaller) (*NextyGovernanceCaller, error) {
	contract, err := bindNextyGovernance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NextyGovernanceCaller{contract: contract}, nil
}

// NewNextyGovernanceTransactor creates a new write-only instance of NextyGovernance, bound to a specific deployed contract.
func NewNextyGovernanceTransactor(address common.Address, transactor bind.ContractTransactor) (*NextyGovernanceTransactor, error) {
	contract, err := bindNextyGovernance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NextyGovernanceTransactor{contract: contract}, nil
}

// NewNextyGovernanceFilterer creates a new log filterer instance of NextyGovernance, bound to a specific deployed contract.
func NewNextyGovernanceFilterer(address common.Address, filterer bind.ContractFilterer) (*NextyGovernanceFilterer, error) {
	contract, err := bindNextyGovernance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NextyGovernanceFilterer{contract: contract}, nil
}

// bindNextyGovernance binds a generic wrapper to an already deployed contract.
func bindNextyGovernance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NextyGovernanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NextyGovernance *NextyGovernanceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NextyGovernance.Contract.NextyGovernanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NextyGovernance *NextyGovernanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NextyGovernance.Contract.NextyGovernanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NextyGovernance *NextyGovernanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NextyGovernance.Contract.NextyGovernanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NextyGovernance *NextyGovernanceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _NextyGovernance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NextyGovernance *NextyGovernanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NextyGovernance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NextyGovernance *NextyGovernanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NextyGovernance.Contract.contract.Transact(opts, method, params...)
}

// Account is a paid mutator transaction binding the contract method 0x73b9aa91.
//
// Solidity: function account(address ) returns(uint8 status, uint256 balance, address signer, uint256 unlockHeight)
func (_NextyGovernance *NextyGovernanceTransactor) Account(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _NextyGovernance.contract.Transact(opts, "account", arg0)
}

// Account is a paid mutator transaction binding the contract method 0x73b9aa91.
//
// Solidity: function account(address ) returns(uint8 status, uint256 balance, address signer, uint256 unlockHeight)
func (_NextyGovernance *NextyGovernanceSession) Account(arg0 common.Address) (*types.Transaction, error) {
	return _NextyGovernance.Contract.Account(&_NextyGovernance.TransactOpts, arg0)
}

// Account is a paid mutator transaction binding the contract method 0x73b9aa91.
//
// Solidity: function account(address ) returns(uint8 status, uint256 balance, address signer, uint256 unlockHeight)
func (_NextyGovernance *NextyGovernanceTransactorSession) Account(arg0 common.Address) (*types.Transaction, error) {
	return _NextyGovernance.Contract.Account(&_NextyGovernance.TransactOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_NextyGovernance *NextyGovernanceTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NextyGovernance.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_NextyGovernance *NextyGovernanceSession) Deposit() (*types.Transaction, error) {
	return _NextyGovernance.Contract.Deposit(&_NextyGovernance.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_NextyGovernance *NextyGovernanceTransactorSession) Deposit() (*types.Transaction, error) {
	return _NextyGovernance.Contract.Deposit(&_NextyGovernance.TransactOpts)
}

// GetBalance is a paid mutator transaction binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address _address) returns(uint256)
func (_NextyGovernance *NextyGovernanceTransactor) GetBalance(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _NextyGovernance.contract.Transact(opts, "getBalance", _address)
}

// GetBalance is a paid mutator transaction binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address _address) returns(uint256)
func (_NextyGovernance *NextyGovernanceSession) GetBalance(_address common.Address) (*types.Transaction, error) {
	return _NextyGovernance.Contract.GetBalance(&_NextyGovernance.TransactOpts, _address)
}

// GetBalance is a paid mutator transaction binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address _address) returns(uint256)
func (_NextyGovernance *NextyGovernanceTransactorSession) GetBalance(_address common.Address) (*types.Transaction, error) {
	return _NextyGovernance.Contract.GetBalance(&_NextyGovernance.TransactOpts, _address)
}

// GetCoinbase is a paid mutator transaction binding the contract method 0x0a9a3f07.
//
// Solidity: function getCoinbase(address _address) returns(address)
func (_NextyGovernance *NextyGovernanceTransactor) GetCoinbase(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _NextyGovernance.contract.Transact(opts, "getCoinbase", _address)
}

// GetCoinbase is a paid mutator transaction binding the contract method 0x0a9a3f07.
//
// Solidity: function getCoinbase(address _address) returns(address)
func (_NextyGovernance *NextyGovernanceSession) GetCoinbase(_address common.Address) (*types.Transaction, error) {
	return _NextyGovernance.Contract.GetCoinbase(&_NextyGovernance.TransactOpts, _address)
}

// GetCoinbase is a paid mutator transaction binding the contract method 0x0a9a3f07.
//
// Solidity: function getCoinbase(address _address) returns(address)
func (_NextyGovernance *NextyGovernanceTransactorSession) GetCoinbase(_address common.Address) (*types.Transaction, error) {
	return _NextyGovernance.Contract.GetCoinbase(&_NextyGovernance.TransactOpts, _address)
}

// GetStatus is a paid mutator transaction binding the contract method 0x30ccebb5.
//
// Solidity: function getStatus(address _address) returns(uint256)
func (_NextyGovernance *NextyGovernanceTransactor) GetStatus(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _NextyGovernance.contract.Transact(opts, "getStatus", _address)
}

// GetStatus is a paid mutator transaction binding the contract method 0x30ccebb5.
//
// Solidity: function getStatus(address _address) returns(uint256)
func (_NextyGovernance *NextyGovernanceSession) GetStatus(_address common.Address) (*types.Transaction, error) {
	return _NextyGovernance.Contract.GetStatus(&_NextyGovernance.TransactOpts, _address)
}

// GetStatus is a paid mutator transaction binding the contract method 0x30ccebb5.
//
// Solidity: function getStatus(address _address) returns(uint256)
func (_NextyGovernance *NextyGovernanceTransactorSession) GetStatus(_address common.Address) (*types.Transaction, error) {
	return _NextyGovernance.Contract.GetStatus(&_NextyGovernance.TransactOpts, _address)
}

// GetUnlockHeight is a paid mutator transaction binding the contract method 0x55235d47.
//
// Solidity: function getUnlockHeight(address _address) returns(uint256)
func (_NextyGovernance *NextyGovernanceTransactor) GetUnlockHeight(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _NextyGovernance.contract.Transact(opts, "getUnlockHeight", _address)
}

// GetUnlockHeight is a paid mutator transaction binding the contract method 0x55235d47.
//
// Solidity: function getUnlockHeight(address _address) returns(uint256)
func (_NextyGovernance *NextyGovernanceSession) GetUnlockHeight(_address common.Address) (*types.Transaction, error) {
	return _NextyGovernance.Contract.GetUnlockHeight(&_NextyGovernance.TransactOpts, _address)
}

// GetUnlockHeight is a paid mutator transaction binding the contract method 0x55235d47.
//
// Solidity: function getUnlockHeight(address _address) returns(uint256)
func (_NextyGovernance *NextyGovernanceTransactorSession) GetUnlockHeight(_address common.Address) (*types.Transaction, error) {
	return _NextyGovernance.Contract.GetUnlockHeight(&_NextyGovernance.TransactOpts, _address)
}

// IsBanned is a paid mutator transaction binding the contract method 0x97f735d5.
//
// Solidity: function isBanned(address _address) returns(bool)
func (_NextyGovernance *NextyGovernanceTransactor) IsBanned(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _NextyGovernance.contract.Transact(opts, "isBanned", _address)
}

// IsBanned is a paid mutator transaction binding the contract method 0x97f735d5.
//
// Solidity: function isBanned(address _address) returns(bool)
func (_NextyGovernance *NextyGovernanceSession) IsBanned(_address common.Address) (*types.Transaction, error) {
	return _NextyGovernance.Contract.IsBanned(&_NextyGovernance.TransactOpts, _address)
}

// IsBanned is a paid mutator transaction binding the contract method 0x97f735d5.
//
// Solidity: function isBanned(address _address) returns(bool)
func (_NextyGovernance *NextyGovernanceTransactorSession) IsBanned(_address common.Address) (*types.Transaction, error) {
	return _NextyGovernance.Contract.IsBanned(&_NextyGovernance.TransactOpts, _address)
}

// IsWithdrawable is a paid mutator transaction binding the contract method 0x02a18484.
//
// Solidity: function isWithdrawable(address _address) returns(bool)
func (_NextyGovernance *NextyGovernanceTransactor) IsWithdrawable(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _NextyGovernance.contract.Transact(opts, "isWithdrawable", _address)
}

// IsWithdrawable is a paid mutator transaction binding the contract method 0x02a18484.
//
// Solidity: function isWithdrawable(address _address) returns(bool)
func (_NextyGovernance *NextyGovernanceSession) IsWithdrawable(_address common.Address) (*types.Transaction, error) {
	return _NextyGovernance.Contract.IsWithdrawable(&_NextyGovernance.TransactOpts, _address)
}

// IsWithdrawable is a paid mutator transaction binding the contract method 0x02a18484.
//
// Solidity: function isWithdrawable(address _address) returns(bool)
func (_NextyGovernance *NextyGovernanceTransactorSession) IsWithdrawable(_address common.Address) (*types.Transaction, error) {
	return _NextyGovernance.Contract.IsWithdrawable(&_NextyGovernance.TransactOpts, _address)
}

// Join is a paid mutator transaction binding the contract method 0x28ffe6c8.
//
// Solidity: function join(address _signer) returns(bool)
func (_NextyGovernance *NextyGovernanceTransactor) Join(opts *bind.TransactOpts, _signer common.Address) (*types.Transaction, error) {
	return _NextyGovernance.contract.Transact(opts, "join", _signer)
}

// Join is a paid mutator transaction binding the contract method 0x28ffe6c8.
//
// Solidity: function join(address _signer) returns(bool)
func (_NextyGovernance *NextyGovernanceSession) Join(_signer common.Address) (*types.Transaction, error) {
	return _NextyGovernance.Contract.Join(&_NextyGovernance.TransactOpts, _signer)
}

// Join is a paid mutator transaction binding the contract method 0x28ffe6c8.
//
// Solidity: function join(address _signer) returns(bool)
func (_NextyGovernance *NextyGovernanceTransactorSession) Join(_signer common.Address) (*types.Transaction, error) {
	return _NextyGovernance.Contract.Join(&_NextyGovernance.TransactOpts, _signer)
}

// Leave is a paid mutator transaction binding the contract method 0xd66d9e19.
//
// Solidity: function leave() returns(bool)
func (_NextyGovernance *NextyGovernanceTransactor) Leave(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NextyGovernance.contract.Transact(opts, "leave")
}

// Leave is a paid mutator transaction binding the contract method 0xd66d9e19.
//
// Solidity: function leave() returns(bool)
func (_NextyGovernance *NextyGovernanceSession) Leave() (*types.Transaction, error) {
	return _NextyGovernance.Contract.Leave(&_NextyGovernance.TransactOpts)
}

// Leave is a paid mutator transaction binding the contract method 0xd66d9e19.
//
// Solidity: function leave() returns(bool)
func (_NextyGovernance *NextyGovernanceTransactorSession) Leave() (*types.Transaction, error) {
	return _NextyGovernance.Contract.Leave(&_NextyGovernance.TransactOpts)
}

// SignerCoinbase is a paid mutator transaction binding the contract method 0x8943b62c.
//
// Solidity: function signerCoinbase(address ) returns(address)
func (_NextyGovernance *NextyGovernanceTransactor) SignerCoinbase(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _NextyGovernance.contract.Transact(opts, "signerCoinbase", arg0)
}

// SignerCoinbase is a paid mutator transaction binding the contract method 0x8943b62c.
//
// Solidity: function signerCoinbase(address ) returns(address)
func (_NextyGovernance *NextyGovernanceSession) SignerCoinbase(arg0 common.Address) (*types.Transaction, error) {
	return _NextyGovernance.Contract.SignerCoinbase(&_NextyGovernance.TransactOpts, arg0)
}

// SignerCoinbase is a paid mutator transaction binding the contract method 0x8943b62c.
//
// Solidity: function signerCoinbase(address ) returns(address)
func (_NextyGovernance *NextyGovernanceTransactorSession) SignerCoinbase(arg0 common.Address) (*types.Transaction, error) {
	return _NextyGovernance.Contract.SignerCoinbase(&_NextyGovernance.TransactOpts, arg0)
}

// Signers is a paid mutator transaction binding the contract method 0x2079fb9a.
//
// Solidity: function signers(uint256 ) returns(address)
func (_NextyGovernance *NextyGovernanceTransactor) Signers(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _NextyGovernance.contract.Transact(opts, "signers", arg0)
}

// Signers is a paid mutator transaction binding the contract method 0x2079fb9a.
//
// Solidity: function signers(uint256 ) returns(address)
func (_NextyGovernance *NextyGovernanceSession) Signers(arg0 *big.Int) (*types.Transaction, error) {
	return _NextyGovernance.Contract.Signers(&_NextyGovernance.TransactOpts, arg0)
}

// Signers is a paid mutator transaction binding the contract method 0x2079fb9a.
//
// Solidity: function signers(uint256 ) returns(address)
func (_NextyGovernance *NextyGovernanceTransactorSession) Signers(arg0 *big.Int) (*types.Transaction, error) {
	return _NextyGovernance.Contract.Signers(&_NextyGovernance.TransactOpts, arg0)
}

// StakeLockHeight is a paid mutator transaction binding the contract method 0x9b212d01.
//
// Solidity: function stakeLockHeight() returns(uint256)
func (_NextyGovernance *NextyGovernanceTransactor) StakeLockHeight(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NextyGovernance.contract.Transact(opts, "stakeLockHeight")
}

// StakeLockHeight is a paid mutator transaction binding the contract method 0x9b212d01.
//
// Solidity: function stakeLockHeight() returns(uint256)
func (_NextyGovernance *NextyGovernanceSession) StakeLockHeight() (*types.Transaction, error) {
	return _NextyGovernance.Contract.StakeLockHeight(&_NextyGovernance.TransactOpts)
}

// StakeLockHeight is a paid mutator transaction binding the contract method 0x9b212d01.
//
// Solidity: function stakeLockHeight() returns(uint256)
func (_NextyGovernance *NextyGovernanceTransactorSession) StakeLockHeight() (*types.Transaction, error) {
	return _NextyGovernance.Contract.StakeLockHeight(&_NextyGovernance.TransactOpts)
}

// StakeRequire is a paid mutator transaction binding the contract method 0xc6908739.
//
// Solidity: function stakeRequire() returns(uint256)
func (_NextyGovernance *NextyGovernanceTransactor) StakeRequire(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NextyGovernance.contract.Transact(opts, "stakeRequire")
}

// StakeRequire is a paid mutator transaction binding the contract method 0xc6908739.
//
// Solidity: function stakeRequire() returns(uint256)
func (_NextyGovernance *NextyGovernanceSession) StakeRequire() (*types.Transaction, error) {
	return _NextyGovernance.Contract.StakeRequire(&_NextyGovernance.TransactOpts)
}

// StakeRequire is a paid mutator transaction binding the contract method 0xc6908739.
//
// Solidity: function stakeRequire() returns(uint256)
func (_NextyGovernance *NextyGovernanceTransactorSession) StakeRequire() (*types.Transaction, error) {
	return _NextyGovernance.Contract.StakeRequire(&_NextyGovernance.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_NextyGovernance *NextyGovernanceTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NextyGovernance.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_NextyGovernance *NextyGovernanceSession) Withdraw() (*types.Transaction, error) {
	return _NextyGovernance.Contract.Withdraw(&_NextyGovernance.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_NextyGovernance *NextyGovernanceTransactorSession) Withdraw() (*types.Transaction, error) {
	return _NextyGovernance.Contract.Withdraw(&_NextyGovernance.TransactOpts)
}

// NextyGovernanceBannedIterator is returned from FilterBanned and is used to iterate over the raw logs and unpacked data for Banned events raised by the NextyGovernance contract.
type NextyGovernanceBannedIterator struct {
	Event *NextyGovernanceBanned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NextyGovernanceBannedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NextyGovernanceBanned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NextyGovernanceBanned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NextyGovernanceBannedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NextyGovernanceBannedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NextyGovernanceBanned represents a Banned event raised by the NextyGovernance contract.
type NextyGovernanceBanned struct {
	Sealer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBanned is a free log retrieval operation binding the contract event 0x30d1df1214d91553408ca5384ce29e10e5866af8423c628be22860e41fb81005.
//
// Solidity: event Banned(address _sealer)
func (_NextyGovernance *NextyGovernanceFilterer) FilterBanned(opts *bind.FilterOpts) (*NextyGovernanceBannedIterator, error) {

	logs, sub, err := _NextyGovernance.contract.FilterLogs(opts, "Banned")
	if err != nil {
		return nil, err
	}
	return &NextyGovernanceBannedIterator{contract: _NextyGovernance.contract, event: "Banned", logs: logs, sub: sub}, nil
}

// WatchBanned is a free log subscription operation binding the contract event 0x30d1df1214d91553408ca5384ce29e10e5866af8423c628be22860e41fb81005.
//
// Solidity: event Banned(address _sealer)
func (_NextyGovernance *NextyGovernanceFilterer) WatchBanned(opts *bind.WatchOpts, sink chan<- *NextyGovernanceBanned) (event.Subscription, error) {

	logs, sub, err := _NextyGovernance.contract.WatchLogs(opts, "Banned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NextyGovernanceBanned)
				if err := _NextyGovernance.contract.UnpackLog(event, "Banned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBanned is a log parse operation binding the contract event 0x30d1df1214d91553408ca5384ce29e10e5866af8423c628be22860e41fb81005.
//
// Solidity: event Banned(address _sealer)
func (_NextyGovernance *NextyGovernanceFilterer) ParseBanned(log types.Log) (*NextyGovernanceBanned, error) {
	event := new(NextyGovernanceBanned)
	if err := _NextyGovernance.contract.UnpackLog(event, "Banned", log); err != nil {
		return nil, err
	}
	return event, nil
}

// NextyGovernanceDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the NextyGovernance contract.
type NextyGovernanceDepositedIterator struct {
	Event *NextyGovernanceDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NextyGovernanceDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NextyGovernanceDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NextyGovernanceDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NextyGovernanceDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NextyGovernanceDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NextyGovernanceDeposited represents a Deposited event raised by the NextyGovernance contract.
type NextyGovernanceDeposited struct {
	Sealer common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address _sealer, uint256 _amount)
func (_NextyGovernance *NextyGovernanceFilterer) FilterDeposited(opts *bind.FilterOpts) (*NextyGovernanceDepositedIterator, error) {

	logs, sub, err := _NextyGovernance.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &NextyGovernanceDepositedIterator{contract: _NextyGovernance.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address _sealer, uint256 _amount)
func (_NextyGovernance *NextyGovernanceFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *NextyGovernanceDeposited) (event.Subscription, error) {

	logs, sub, err := _NextyGovernance.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NextyGovernanceDeposited)
				if err := _NextyGovernance.contract.UnpackLog(event, "Deposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposited is a log parse operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address _sealer, uint256 _amount)
func (_NextyGovernance *NextyGovernanceFilterer) ParseDeposited(log types.Log) (*NextyGovernanceDeposited, error) {
	event := new(NextyGovernanceDeposited)
	if err := _NextyGovernance.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	return event, nil
}

// NextyGovernanceJoinedIterator is returned from FilterJoined and is used to iterate over the raw logs and unpacked data for Joined events raised by the NextyGovernance contract.
type NextyGovernanceJoinedIterator struct {
	Event *NextyGovernanceJoined // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NextyGovernanceJoinedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NextyGovernanceJoined)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NextyGovernanceJoined)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NextyGovernanceJoinedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NextyGovernanceJoinedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NextyGovernanceJoined represents a Joined event raised by the NextyGovernance contract.
type NextyGovernanceJoined struct {
	Sealer common.Address
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterJoined is a free log retrieval operation binding the contract event 0x7702dccda75540ad1dca8d5276c048f4a5c0e4203f6da4be214bfb1901b203ea.
//
// Solidity: event Joined(address _sealer, address _signer)
func (_NextyGovernance *NextyGovernanceFilterer) FilterJoined(opts *bind.FilterOpts) (*NextyGovernanceJoinedIterator, error) {

	logs, sub, err := _NextyGovernance.contract.FilterLogs(opts, "Joined")
	if err != nil {
		return nil, err
	}
	return &NextyGovernanceJoinedIterator{contract: _NextyGovernance.contract, event: "Joined", logs: logs, sub: sub}, nil
}

// WatchJoined is a free log subscription operation binding the contract event 0x7702dccda75540ad1dca8d5276c048f4a5c0e4203f6da4be214bfb1901b203ea.
//
// Solidity: event Joined(address _sealer, address _signer)
func (_NextyGovernance *NextyGovernanceFilterer) WatchJoined(opts *bind.WatchOpts, sink chan<- *NextyGovernanceJoined) (event.Subscription, error) {

	logs, sub, err := _NextyGovernance.contract.WatchLogs(opts, "Joined")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NextyGovernanceJoined)
				if err := _NextyGovernance.contract.UnpackLog(event, "Joined", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseJoined is a log parse operation binding the contract event 0x7702dccda75540ad1dca8d5276c048f4a5c0e4203f6da4be214bfb1901b203ea.
//
// Solidity: event Joined(address _sealer, address _signer)
func (_NextyGovernance *NextyGovernanceFilterer) ParseJoined(log types.Log) (*NextyGovernanceJoined, error) {
	event := new(NextyGovernanceJoined)
	if err := _NextyGovernance.contract.UnpackLog(event, "Joined", log); err != nil {
		return nil, err
	}
	return event, nil
}

// NextyGovernanceLeftIterator is returned from FilterLeft and is used to iterate over the raw logs and unpacked data for Left events raised by the NextyGovernance contract.
type NextyGovernanceLeftIterator struct {
	Event *NextyGovernanceLeft // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NextyGovernanceLeftIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NextyGovernanceLeft)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NextyGovernanceLeft)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NextyGovernanceLeftIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NextyGovernanceLeftIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NextyGovernanceLeft represents a Left event raised by the NextyGovernance contract.
type NextyGovernanceLeft struct {
	Sealer common.Address
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLeft is a free log retrieval operation binding the contract event 0x4b9ee4dd061ba088b22898a02491f3896a4a580c6cda8783ca579ee159f8e8c5.
//
// Solidity: event Left(address _sealer, address _signer)
func (_NextyGovernance *NextyGovernanceFilterer) FilterLeft(opts *bind.FilterOpts) (*NextyGovernanceLeftIterator, error) {

	logs, sub, err := _NextyGovernance.contract.FilterLogs(opts, "Left")
	if err != nil {
		return nil, err
	}
	return &NextyGovernanceLeftIterator{contract: _NextyGovernance.contract, event: "Left", logs: logs, sub: sub}, nil
}

// WatchLeft is a free log subscription operation binding the contract event 0x4b9ee4dd061ba088b22898a02491f3896a4a580c6cda8783ca579ee159f8e8c5.
//
// Solidity: event Left(address _sealer, address _signer)
func (_NextyGovernance *NextyGovernanceFilterer) WatchLeft(opts *bind.WatchOpts, sink chan<- *NextyGovernanceLeft) (event.Subscription, error) {

	logs, sub, err := _NextyGovernance.contract.WatchLogs(opts, "Left")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NextyGovernanceLeft)
				if err := _NextyGovernance.contract.UnpackLog(event, "Left", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLeft is a log parse operation binding the contract event 0x4b9ee4dd061ba088b22898a02491f3896a4a580c6cda8783ca579ee159f8e8c5.
//
// Solidity: event Left(address _sealer, address _signer)
func (_NextyGovernance *NextyGovernanceFilterer) ParseLeft(log types.Log) (*NextyGovernanceLeft, error) {
	event := new(NextyGovernanceLeft)
	if err := _NextyGovernance.contract.UnpackLog(event, "Left", log); err != nil {
		return nil, err
	}
	return event, nil
}

// NextyGovernanceUnbannedIterator is returned from FilterUnbanned and is used to iterate over the raw logs and unpacked data for Unbanned events raised by the NextyGovernance contract.
type NextyGovernanceUnbannedIterator struct {
	Event *NextyGovernanceUnbanned // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NextyGovernanceUnbannedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NextyGovernanceUnbanned)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NextyGovernanceUnbanned)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NextyGovernanceUnbannedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NextyGovernanceUnbannedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NextyGovernanceUnbanned represents a Unbanned event raised by the NextyGovernance contract.
type NextyGovernanceUnbanned struct {
	Sealer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnbanned is a free log retrieval operation binding the contract event 0x2ab91b53354938415bb6962c4322231cd4cb2c84930f1a4b9abbedc2fe8abe72.
//
// Solidity: event Unbanned(address _sealer)
func (_NextyGovernance *NextyGovernanceFilterer) FilterUnbanned(opts *bind.FilterOpts) (*NextyGovernanceUnbannedIterator, error) {

	logs, sub, err := _NextyGovernance.contract.FilterLogs(opts, "Unbanned")
	if err != nil {
		return nil, err
	}
	return &NextyGovernanceUnbannedIterator{contract: _NextyGovernance.contract, event: "Unbanned", logs: logs, sub: sub}, nil
}

// WatchUnbanned is a free log subscription operation binding the contract event 0x2ab91b53354938415bb6962c4322231cd4cb2c84930f1a4b9abbedc2fe8abe72.
//
// Solidity: event Unbanned(address _sealer)
func (_NextyGovernance *NextyGovernanceFilterer) WatchUnbanned(opts *bind.WatchOpts, sink chan<- *NextyGovernanceUnbanned) (event.Subscription, error) {

	logs, sub, err := _NextyGovernance.contract.WatchLogs(opts, "Unbanned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NextyGovernanceUnbanned)
				if err := _NextyGovernance.contract.UnpackLog(event, "Unbanned", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnbanned is a log parse operation binding the contract event 0x2ab91b53354938415bb6962c4322231cd4cb2c84930f1a4b9abbedc2fe8abe72.
//
// Solidity: event Unbanned(address _sealer)
func (_NextyGovernance *NextyGovernanceFilterer) ParseUnbanned(log types.Log) (*NextyGovernanceUnbanned, error) {
	event := new(NextyGovernanceUnbanned)
	if err := _NextyGovernance.contract.UnpackLog(event, "Unbanned", log); err != nil {
		return nil, err
	}
	return event, nil
}

// NextyGovernanceWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the NextyGovernance contract.
type NextyGovernanceWithdrawnIterator struct {
	Event *NextyGovernanceWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *NextyGovernanceWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NextyGovernanceWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(NextyGovernanceWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *NextyGovernanceWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NextyGovernanceWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NextyGovernanceWithdrawn represents a Withdrawn event raised by the NextyGovernance contract.
type NextyGovernanceWithdrawn struct {
	Sealer common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address _sealer, uint256 _amount)
func (_NextyGovernance *NextyGovernanceFilterer) FilterWithdrawn(opts *bind.FilterOpts) (*NextyGovernanceWithdrawnIterator, error) {

	logs, sub, err := _NextyGovernance.contract.FilterLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return &NextyGovernanceWithdrawnIterator{contract: _NextyGovernance.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address _sealer, uint256 _amount)
func (_NextyGovernance *NextyGovernanceFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *NextyGovernanceWithdrawn) (event.Subscription, error) {

	logs, sub, err := _NextyGovernance.contract.WatchLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NextyGovernanceWithdrawn)
				if err := _NextyGovernance.contract.UnpackLog(event, "Withdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address _sealer, uint256 _amount)
func (_NextyGovernance *NextyGovernanceFilterer) ParseWithdrawn(log types.Log) (*NextyGovernanceWithdrawn, error) {
	event := new(NextyGovernanceWithdrawn)
	if err := _NextyGovernance.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x607b6023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220436ff7c3075f257c72e37a50670b3183fc755a3fc6595ae79ce1f2c71885b9ec64736f6c637827302e362e392d646576656c6f702e323032302e352e32372b636f6d6d69742e39663430376665300058"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}
