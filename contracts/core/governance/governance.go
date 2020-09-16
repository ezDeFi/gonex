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

// GovernanceABI is the input ABI used to generate the binding from.
const GovernanceABI = "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakeRequire\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stakeLockHeight\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_signers\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sealer\",\"type\":\"address\"}],\"name\":\"Banned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sealer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Deposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sealer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"Joined\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sealer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"Left\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sealer\",\"type\":\"address\"}],\"name\":\"Unbanned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_sealer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"account\",\"outputs\":[{\"internalType\":\"enumGovernance.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"unlockHeight\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getCoinbase\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getUnlockHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isBanned\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isWithdrawable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signer\",\"type\":\"address\"}],\"name\":\"join\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"leave\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"signerCoinbase\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"signers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeLockHeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeRequire\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// GovernanceFuncSigs maps the 4-byte function signature to its string representation.
var GovernanceFuncSigs = map[string]string{
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

// GovernanceBin is the compiled bytecode used for deploying new contracts.
var GovernanceBin = "0x608060405234801561001057600080fd5b50604051620010fa380380620010fa8339818101604052606081101561003557600080fd5b8151602083015160408085018051915193959294830192918464010000000082111561006057600080fd5b90830190602082018581111561007557600080fd5b825186602082028301116401000000008211171561009257600080fd5b82525081516020918201928201910280838360005b838110156100bf5781810151838201526020016100a7565b5050505091909101604052505050600384905550600482905560005b81518110156102685760008282815181106100f257fe5b60209081029190910181015182546001810184556000938452919092200180546001600160a01b0319166001600160a01b03909216919091179055815182908290811061013b57fe5b60200260200101516001600084848151811061015357fe5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a8154816001600160a01b0302191690836001600160a01b031602179055508181815181106101ab57fe5b6020026020010151600260008484815181106101c357fe5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060020160006101000a8154816001600160a01b0302191690836001600160a01b0316021790555060016002600084848151811061022457fe5b6020908102919091018101516001600160a01b03168252810191909152604001600020805460ff1916600183600481111561025b57fe5b02179055506001016100db565b50505050610e7e806200027c6000396000f3fe6080604052600436106100ec5760003560e01c806373b9aa911161008a578063c690873911610059578063c69087391461037a578063d0e30db01461038f578063d66d9e1914610397578063f8b2cb4f146103ac576100fc565b806373b9aa91146102815780638943b62c146102ff57806397f735d5146103325780639b212d0114610365576100fc565b806328ffe6c8116100c657806328ffe6c8146101c157806330ccebb5146101f45780633ccfd60b1461023957806355235d471461024e576100fc565b806302a18484146101015780630a9a3f07146101485780632079fb9a14610197576100fc565b366100fc576100fa346103df565b005b600080fd5b34801561010d57600080fd5b506101346004803603602081101561012457600080fd5b50356001600160a01b0316610457565b604080519115158252519081900360200190f35b34801561015457600080fd5b5061017b6004803603602081101561016b57600080fd5b50356001600160a01b03166104e5565b604080516001600160a01b039092168252519081900360200190f35b3480156101a357600080fd5b5061017b600480360360208110156101ba57600080fd5b5035610507565b3480156101cd57600080fd5b50610134600480360360208110156101e457600080fd5b50356001600160a01b031661052e565b34801561020057600080fd5b506102276004803603602081101561021757600080fd5b50356001600160a01b031661086c565b60408051918252519081900360200190f35b34801561024557600080fd5b50610134610891565b34801561025a57600080fd5b506102276004803603602081101561027157600080fd5b50356001600160a01b03166109df565b34801561028d57600080fd5b506102b4600480360360208110156102a457600080fd5b50356001600160a01b03166109fd565b604051808560048111156102c457fe5b60ff168152602001848152602001836001600160a01b03166001600160a01b0316815260200182815260200194505050505060405180910390f35b34801561030b57600080fd5b5061017b6004803603602081101561032257600080fd5b50356001600160a01b0316610a33565b34801561033e57600080fd5b506101346004803603602081101561035557600080fd5b50356001600160a01b0316610a4e565b34801561037157600080fd5b50610227610a80565b34801561038657600080fd5b50610227610a86565b6100fa610a8c565b3480156103a357600080fd5b50610134610a97565b3480156103b857600080fd5b50610227600480360360208110156103cf57600080fd5b50356001600160a01b0316610c2b565b33600090815260026020526040902060010154610402908263ffffffff610c4916565b3360008181526002602090815260409182902060010193909355805191825291810183905281517f2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4929181900390910190a150565b600060016001600160a01b03831660009081526002602052604090205460ff16600481111561048257fe5b141580156104b7575060046001600160a01b03831660009081526002602052604090205460ff1660048111156104b457fe5b14155b80156104dd57506001600160a01b03821660009081526002602052604090206003015443115b90505b919050565b6001600160a01b03908116600090815260026020819052604090912001541690565b6000818154811061051457fe5b6000918252602090912001546001600160a01b0316905081565b600060043360009081526002602052604090205460ff16600481111561055057fe5b141561058d576040805162461bcd60e51b815260206004820152600760248201526603130b73732b2160cd1b604482015290519081900360640190fd5b60013360009081526002602052604090205460ff1660048111156105ad57fe5b14156105f2576040805162461bcd60e51b815260206004820152600f60248201526e030b63932b0b23c903537b4b732b21608d1b604482015290519081900360640190fd5b60035433600090815260026020526040902060010154101561064c576040805162461bcd60e51b815260206004820152600e60248201526d3737ba1032b737bab3b410373a3360911b604482015290519081900360640190fd5b6001600160a01b03808316600090815260016020526040902054839116156106b3576040805162461bcd60e51b815260206004820152601560248201527418dbda5b98985cd948185b1c9958591e481d5cd959605a1b604482015290519081900360640190fd5b6001600160a01b0381166106fc576040805162461bcd60e51b815260206004820152600b60248201526a7369676e6572207a65726f60a81b604482015290519081900360640190fd5b6001600160a01b03811630141561075a576040805162461bcd60e51b815260206004820152601760248201527f73616d6520636f6e747261637427732061646472657373000000000000000000604482015290519081900360640190fd5b6001600160a01b0381163314156107b0576040805162461bcd60e51b815260206004820152601560248201527473616d652073656e6465722773206164647265737360581b604482015290519081900360640190fd5b33600090815260026020819052604090912090810180546001600160a01b0319166001600160a01b03861617905580546001919060ff1916828002179055506001600160a01b038316600090815260016020526040902080546001600160a01b0319163317905561082083610caa565b604080513381526001600160a01b038516602082015281517f7702dccda75540ad1dca8d5276c048f4a5c0e4203f6da4be214bfb1901b203ea929181900390910190a150600192915050565b6001600160a01b0381166000908152600260205260408120546104dd9060ff16610cf9565b600060043360009081526002602052604090205460ff1660048111156108b357fe5b14156108f0576040805162461bcd60e51b815260206004820152600760248201526603130b73732b2160cd1b604482015290519081900360640190fd5b6108f933610457565b61094a576040805162461bcd60e51b815260206004820181905260248201527f756e61626c6520746f20776974686472617720617420746865206d6f6d656e74604482015290519081900360640190fd5b3360008181526002602052604080822060018101805490849055815460ff191660031790915590519092916108fc841502918491818181858888f1935050505015801561099b573d6000803e3d6000fd5b50604080513381526020810183905281517f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5929181900390910190a1600191505090565b6001600160a01b031660009081526002602052604090206003015490565b6002602081905260009182526040909120805460018201549282015460039092015460ff90911692916001600160a01b03169084565b6001602052600090815260409020546001600160a01b031681565b600060046001600160a01b03831660009081526002602052604090205460ff166004811115610a7957fe5b1492915050565b60045481565b60035481565b610a95346103df565b565b600060043360009081526002602052604090205460ff166004811115610ab957fe5b1415610af6576040805162461bcd60e51b815260206004820152600760248201526603130b73732b2160cd1b604482015290519081900360640190fd5b60013360009081526002602052604090205460ff166004811115610b1657fe5b14610b56576040805162461bcd60e51b815260206004820152600b60248201526a03737ba103537b4b732b2160ad1b604482015290519081900360640190fd5b33600090815260026020819052604090912080820180546001600160a01b03198116909155815460ff191690921790556004546001600160a01b0390911690610b9f9043610c49565b336000908152600260209081526040808320600301939093556001600160a01b0384168252600190522080546001600160a01b0319169055610be081610d72565b604080513381526001600160a01b038316602082015281517f4b9ee4dd061ba088b22898a02491f3896a4a580c6cda8783ca579ee159f8e8c5929181900390910190a1600191505090565b6001600160a01b031660009081526002602052604090206001015490565b600082820183811015610ca3576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b600080546001810182559080527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5630180546001600160a01b0319166001600160a01b0392909216919091179055565b600080826004811115610d0857fe5b1415610d16575060006104e0565b6001826004811115610d2457fe5b1415610d32575060016104e0565b6002826004811115610d4057fe5b1415610d4e575060026104e0565b6003826004811115610d5c57fe5b1415610d6a575060036104e0565b50607f919050565b60005b600054811015610e435760008181548110610d8c57fe5b6000918252602090912001546001600160a01b0383811691161415610e3b57600080546000198101908110610dbd57fe5b600091825260208220015481546001600160a01b03909116919083908110610de157fe5b6000918252602082200180546001600160a01b0319166001600160a01b039390931692909217909155805480610e1357fe5b600082815260209020810160001990810180546001600160a01b031916905501905550610e45565b600101610d75565b505b5056fea264697066735822122076e2fe905cdac8fcd2560da54ee91a4aff8088d468109510c6da6a1c74c0510464736f6c63430006080033"

// DeployGovernance deploys a new Ethereum contract, binding an instance of Governance to it.
func DeployGovernance(auth *bind.TransactOpts, backend bind.ContractBackend, _stakeRequire *big.Int, _stakeLockHeight *big.Int, _signers []common.Address) (common.Address, *types.Transaction, *Governance, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernanceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GovernanceBin), backend, _stakeRequire, _stakeLockHeight, _signers)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Governance{GovernanceCaller: GovernanceCaller{contract: contract}, GovernanceTransactor: GovernanceTransactor{contract: contract}, GovernanceFilterer: GovernanceFilterer{contract: contract}}, nil
}

// Governance is an auto generated Go binding around an Ethereum contract.
type Governance struct {
	GovernanceCaller     // Read-only binding to the contract
	GovernanceTransactor // Write-only binding to the contract
	GovernanceFilterer   // Log filterer for contract events
}

// GovernanceCaller is an auto generated read-only Go binding around an Ethereum contract.
type GovernanceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GovernanceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GovernanceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GovernanceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GovernanceSession struct {
	Contract     *Governance       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GovernanceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GovernanceCallerSession struct {
	Contract *GovernanceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// GovernanceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GovernanceTransactorSession struct {
	Contract     *GovernanceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GovernanceRaw is an auto generated low-level Go binding around an Ethereum contract.
type GovernanceRaw struct {
	Contract *Governance // Generic contract binding to access the raw methods on
}

// GovernanceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GovernanceCallerRaw struct {
	Contract *GovernanceCaller // Generic read-only contract binding to access the raw methods on
}

// GovernanceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GovernanceTransactorRaw struct {
	Contract *GovernanceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGovernance creates a new instance of Governance, bound to a specific deployed contract.
func NewGovernance(address common.Address, backend bind.ContractBackend) (*Governance, error) {
	contract, err := bindGovernance(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Governance{GovernanceCaller: GovernanceCaller{contract: contract}, GovernanceTransactor: GovernanceTransactor{contract: contract}, GovernanceFilterer: GovernanceFilterer{contract: contract}}, nil
}

// NewGovernanceCaller creates a new read-only instance of Governance, bound to a specific deployed contract.
func NewGovernanceCaller(address common.Address, caller bind.ContractCaller) (*GovernanceCaller, error) {
	contract, err := bindGovernance(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceCaller{contract: contract}, nil
}

// NewGovernanceTransactor creates a new write-only instance of Governance, bound to a specific deployed contract.
func NewGovernanceTransactor(address common.Address, transactor bind.ContractTransactor) (*GovernanceTransactor, error) {
	contract, err := bindGovernance(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GovernanceTransactor{contract: contract}, nil
}

// NewGovernanceFilterer creates a new log filterer instance of Governance, bound to a specific deployed contract.
func NewGovernanceFilterer(address common.Address, filterer bind.ContractFilterer) (*GovernanceFilterer, error) {
	contract, err := bindGovernance(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GovernanceFilterer{contract: contract}, nil
}

// bindGovernance binds a generic wrapper to an already deployed contract.
func bindGovernance(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GovernanceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Governance *GovernanceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Governance.Contract.GovernanceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Governance *GovernanceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.Contract.GovernanceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Governance *GovernanceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Governance.Contract.GovernanceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Governance *GovernanceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Governance.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Governance *GovernanceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Governance *GovernanceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Governance.Contract.contract.Transact(opts, method, params...)
}

// Account is a paid mutator transaction binding the contract method 0x73b9aa91.
//
// Solidity: function account(address ) returns(uint8 status, uint256 balance, address signer, uint256 unlockHeight)
func (_Governance *GovernanceTransactor) Account(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "account", arg0)
}

// Account is a paid mutator transaction binding the contract method 0x73b9aa91.
//
// Solidity: function account(address ) returns(uint8 status, uint256 balance, address signer, uint256 unlockHeight)
func (_Governance *GovernanceSession) Account(arg0 common.Address) (*types.Transaction, error) {
	return _Governance.Contract.Account(&_Governance.TransactOpts, arg0)
}

// Account is a paid mutator transaction binding the contract method 0x73b9aa91.
//
// Solidity: function account(address ) returns(uint8 status, uint256 balance, address signer, uint256 unlockHeight)
func (_Governance *GovernanceTransactorSession) Account(arg0 common.Address) (*types.Transaction, error) {
	return _Governance.Contract.Account(&_Governance.TransactOpts, arg0)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Governance *GovernanceTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Governance *GovernanceSession) Deposit() (*types.Transaction, error) {
	return _Governance.Contract.Deposit(&_Governance.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns()
func (_Governance *GovernanceTransactorSession) Deposit() (*types.Transaction, error) {
	return _Governance.Contract.Deposit(&_Governance.TransactOpts)
}

// GetBalance is a paid mutator transaction binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address _address) returns(uint256)
func (_Governance *GovernanceTransactor) GetBalance(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "getBalance", _address)
}

// GetBalance is a paid mutator transaction binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address _address) returns(uint256)
func (_Governance *GovernanceSession) GetBalance(_address common.Address) (*types.Transaction, error) {
	return _Governance.Contract.GetBalance(&_Governance.TransactOpts, _address)
}

// GetBalance is a paid mutator transaction binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address _address) returns(uint256)
func (_Governance *GovernanceTransactorSession) GetBalance(_address common.Address) (*types.Transaction, error) {
	return _Governance.Contract.GetBalance(&_Governance.TransactOpts, _address)
}

// GetCoinbase is a paid mutator transaction binding the contract method 0x0a9a3f07.
//
// Solidity: function getCoinbase(address _address) returns(address)
func (_Governance *GovernanceTransactor) GetCoinbase(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "getCoinbase", _address)
}

// GetCoinbase is a paid mutator transaction binding the contract method 0x0a9a3f07.
//
// Solidity: function getCoinbase(address _address) returns(address)
func (_Governance *GovernanceSession) GetCoinbase(_address common.Address) (*types.Transaction, error) {
	return _Governance.Contract.GetCoinbase(&_Governance.TransactOpts, _address)
}

// GetCoinbase is a paid mutator transaction binding the contract method 0x0a9a3f07.
//
// Solidity: function getCoinbase(address _address) returns(address)
func (_Governance *GovernanceTransactorSession) GetCoinbase(_address common.Address) (*types.Transaction, error) {
	return _Governance.Contract.GetCoinbase(&_Governance.TransactOpts, _address)
}

// GetStatus is a paid mutator transaction binding the contract method 0x30ccebb5.
//
// Solidity: function getStatus(address _address) returns(uint256)
func (_Governance *GovernanceTransactor) GetStatus(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "getStatus", _address)
}

// GetStatus is a paid mutator transaction binding the contract method 0x30ccebb5.
//
// Solidity: function getStatus(address _address) returns(uint256)
func (_Governance *GovernanceSession) GetStatus(_address common.Address) (*types.Transaction, error) {
	return _Governance.Contract.GetStatus(&_Governance.TransactOpts, _address)
}

// GetStatus is a paid mutator transaction binding the contract method 0x30ccebb5.
//
// Solidity: function getStatus(address _address) returns(uint256)
func (_Governance *GovernanceTransactorSession) GetStatus(_address common.Address) (*types.Transaction, error) {
	return _Governance.Contract.GetStatus(&_Governance.TransactOpts, _address)
}

// GetUnlockHeight is a paid mutator transaction binding the contract method 0x55235d47.
//
// Solidity: function getUnlockHeight(address _address) returns(uint256)
func (_Governance *GovernanceTransactor) GetUnlockHeight(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "getUnlockHeight", _address)
}

// GetUnlockHeight is a paid mutator transaction binding the contract method 0x55235d47.
//
// Solidity: function getUnlockHeight(address _address) returns(uint256)
func (_Governance *GovernanceSession) GetUnlockHeight(_address common.Address) (*types.Transaction, error) {
	return _Governance.Contract.GetUnlockHeight(&_Governance.TransactOpts, _address)
}

// GetUnlockHeight is a paid mutator transaction binding the contract method 0x55235d47.
//
// Solidity: function getUnlockHeight(address _address) returns(uint256)
func (_Governance *GovernanceTransactorSession) GetUnlockHeight(_address common.Address) (*types.Transaction, error) {
	return _Governance.Contract.GetUnlockHeight(&_Governance.TransactOpts, _address)
}

// IsBanned is a paid mutator transaction binding the contract method 0x97f735d5.
//
// Solidity: function isBanned(address _address) returns(bool)
func (_Governance *GovernanceTransactor) IsBanned(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "isBanned", _address)
}

// IsBanned is a paid mutator transaction binding the contract method 0x97f735d5.
//
// Solidity: function isBanned(address _address) returns(bool)
func (_Governance *GovernanceSession) IsBanned(_address common.Address) (*types.Transaction, error) {
	return _Governance.Contract.IsBanned(&_Governance.TransactOpts, _address)
}

// IsBanned is a paid mutator transaction binding the contract method 0x97f735d5.
//
// Solidity: function isBanned(address _address) returns(bool)
func (_Governance *GovernanceTransactorSession) IsBanned(_address common.Address) (*types.Transaction, error) {
	return _Governance.Contract.IsBanned(&_Governance.TransactOpts, _address)
}

// IsWithdrawable is a paid mutator transaction binding the contract method 0x02a18484.
//
// Solidity: function isWithdrawable(address _address) returns(bool)
func (_Governance *GovernanceTransactor) IsWithdrawable(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "isWithdrawable", _address)
}

// IsWithdrawable is a paid mutator transaction binding the contract method 0x02a18484.
//
// Solidity: function isWithdrawable(address _address) returns(bool)
func (_Governance *GovernanceSession) IsWithdrawable(_address common.Address) (*types.Transaction, error) {
	return _Governance.Contract.IsWithdrawable(&_Governance.TransactOpts, _address)
}

// IsWithdrawable is a paid mutator transaction binding the contract method 0x02a18484.
//
// Solidity: function isWithdrawable(address _address) returns(bool)
func (_Governance *GovernanceTransactorSession) IsWithdrawable(_address common.Address) (*types.Transaction, error) {
	return _Governance.Contract.IsWithdrawable(&_Governance.TransactOpts, _address)
}

// Join is a paid mutator transaction binding the contract method 0x28ffe6c8.
//
// Solidity: function join(address _signer) returns(bool)
func (_Governance *GovernanceTransactor) Join(opts *bind.TransactOpts, _signer common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "join", _signer)
}

// Join is a paid mutator transaction binding the contract method 0x28ffe6c8.
//
// Solidity: function join(address _signer) returns(bool)
func (_Governance *GovernanceSession) Join(_signer common.Address) (*types.Transaction, error) {
	return _Governance.Contract.Join(&_Governance.TransactOpts, _signer)
}

// Join is a paid mutator transaction binding the contract method 0x28ffe6c8.
//
// Solidity: function join(address _signer) returns(bool)
func (_Governance *GovernanceTransactorSession) Join(_signer common.Address) (*types.Transaction, error) {
	return _Governance.Contract.Join(&_Governance.TransactOpts, _signer)
}

// Leave is a paid mutator transaction binding the contract method 0xd66d9e19.
//
// Solidity: function leave() returns(bool)
func (_Governance *GovernanceTransactor) Leave(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "leave")
}

// Leave is a paid mutator transaction binding the contract method 0xd66d9e19.
//
// Solidity: function leave() returns(bool)
func (_Governance *GovernanceSession) Leave() (*types.Transaction, error) {
	return _Governance.Contract.Leave(&_Governance.TransactOpts)
}

// Leave is a paid mutator transaction binding the contract method 0xd66d9e19.
//
// Solidity: function leave() returns(bool)
func (_Governance *GovernanceTransactorSession) Leave() (*types.Transaction, error) {
	return _Governance.Contract.Leave(&_Governance.TransactOpts)
}

// SignerCoinbase is a paid mutator transaction binding the contract method 0x8943b62c.
//
// Solidity: function signerCoinbase(address ) returns(address)
func (_Governance *GovernanceTransactor) SignerCoinbase(opts *bind.TransactOpts, arg0 common.Address) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "signerCoinbase", arg0)
}

// SignerCoinbase is a paid mutator transaction binding the contract method 0x8943b62c.
//
// Solidity: function signerCoinbase(address ) returns(address)
func (_Governance *GovernanceSession) SignerCoinbase(arg0 common.Address) (*types.Transaction, error) {
	return _Governance.Contract.SignerCoinbase(&_Governance.TransactOpts, arg0)
}

// SignerCoinbase is a paid mutator transaction binding the contract method 0x8943b62c.
//
// Solidity: function signerCoinbase(address ) returns(address)
func (_Governance *GovernanceTransactorSession) SignerCoinbase(arg0 common.Address) (*types.Transaction, error) {
	return _Governance.Contract.SignerCoinbase(&_Governance.TransactOpts, arg0)
}

// Signers is a paid mutator transaction binding the contract method 0x2079fb9a.
//
// Solidity: function signers(uint256 ) returns(address)
func (_Governance *GovernanceTransactor) Signers(opts *bind.TransactOpts, arg0 *big.Int) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "signers", arg0)
}

// Signers is a paid mutator transaction binding the contract method 0x2079fb9a.
//
// Solidity: function signers(uint256 ) returns(address)
func (_Governance *GovernanceSession) Signers(arg0 *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Signers(&_Governance.TransactOpts, arg0)
}

// Signers is a paid mutator transaction binding the contract method 0x2079fb9a.
//
// Solidity: function signers(uint256 ) returns(address)
func (_Governance *GovernanceTransactorSession) Signers(arg0 *big.Int) (*types.Transaction, error) {
	return _Governance.Contract.Signers(&_Governance.TransactOpts, arg0)
}

// StakeLockHeight is a paid mutator transaction binding the contract method 0x9b212d01.
//
// Solidity: function stakeLockHeight() returns(uint256)
func (_Governance *GovernanceTransactor) StakeLockHeight(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "stakeLockHeight")
}

// StakeLockHeight is a paid mutator transaction binding the contract method 0x9b212d01.
//
// Solidity: function stakeLockHeight() returns(uint256)
func (_Governance *GovernanceSession) StakeLockHeight() (*types.Transaction, error) {
	return _Governance.Contract.StakeLockHeight(&_Governance.TransactOpts)
}

// StakeLockHeight is a paid mutator transaction binding the contract method 0x9b212d01.
//
// Solidity: function stakeLockHeight() returns(uint256)
func (_Governance *GovernanceTransactorSession) StakeLockHeight() (*types.Transaction, error) {
	return _Governance.Contract.StakeLockHeight(&_Governance.TransactOpts)
}

// StakeRequire is a paid mutator transaction binding the contract method 0xc6908739.
//
// Solidity: function stakeRequire() returns(uint256)
func (_Governance *GovernanceTransactor) StakeRequire(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "stakeRequire")
}

// StakeRequire is a paid mutator transaction binding the contract method 0xc6908739.
//
// Solidity: function stakeRequire() returns(uint256)
func (_Governance *GovernanceSession) StakeRequire() (*types.Transaction, error) {
	return _Governance.Contract.StakeRequire(&_Governance.TransactOpts)
}

// StakeRequire is a paid mutator transaction binding the contract method 0xc6908739.
//
// Solidity: function stakeRequire() returns(uint256)
func (_Governance *GovernanceTransactorSession) StakeRequire() (*types.Transaction, error) {
	return _Governance.Contract.StakeRequire(&_Governance.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_Governance *GovernanceTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Governance.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_Governance *GovernanceSession) Withdraw() (*types.Transaction, error) {
	return _Governance.Contract.Withdraw(&_Governance.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns(bool)
func (_Governance *GovernanceTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Governance.Contract.Withdraw(&_Governance.TransactOpts)
}

// GovernanceBannedIterator is returned from FilterBanned and is used to iterate over the raw logs and unpacked data for Banned events raised by the Governance contract.
type GovernanceBannedIterator struct {
	Event *GovernanceBanned // Event containing the contract specifics and raw log

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
func (it *GovernanceBannedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceBanned)
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
		it.Event = new(GovernanceBanned)
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
func (it *GovernanceBannedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceBannedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceBanned represents a Banned event raised by the Governance contract.
type GovernanceBanned struct {
	Sealer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBanned is a free log retrieval operation binding the contract event 0x30d1df1214d91553408ca5384ce29e10e5866af8423c628be22860e41fb81005.
//
// Solidity: event Banned(address _sealer)
func (_Governance *GovernanceFilterer) FilterBanned(opts *bind.FilterOpts) (*GovernanceBannedIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "Banned")
	if err != nil {
		return nil, err
	}
	return &GovernanceBannedIterator{contract: _Governance.contract, event: "Banned", logs: logs, sub: sub}, nil
}

// WatchBanned is a free log subscription operation binding the contract event 0x30d1df1214d91553408ca5384ce29e10e5866af8423c628be22860e41fb81005.
//
// Solidity: event Banned(address _sealer)
func (_Governance *GovernanceFilterer) WatchBanned(opts *bind.WatchOpts, sink chan<- *GovernanceBanned) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "Banned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceBanned)
				if err := _Governance.contract.UnpackLog(event, "Banned", log); err != nil {
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
func (_Governance *GovernanceFilterer) ParseBanned(log types.Log) (*GovernanceBanned, error) {
	event := new(GovernanceBanned)
	if err := _Governance.contract.UnpackLog(event, "Banned", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceDepositedIterator is returned from FilterDeposited and is used to iterate over the raw logs and unpacked data for Deposited events raised by the Governance contract.
type GovernanceDepositedIterator struct {
	Event *GovernanceDeposited // Event containing the contract specifics and raw log

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
func (it *GovernanceDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceDeposited)
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
		it.Event = new(GovernanceDeposited)
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
func (it *GovernanceDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceDeposited represents a Deposited event raised by the Governance contract.
type GovernanceDeposited struct {
	Sealer common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposited is a free log retrieval operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address _sealer, uint256 _amount)
func (_Governance *GovernanceFilterer) FilterDeposited(opts *bind.FilterOpts) (*GovernanceDepositedIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return &GovernanceDepositedIterator{contract: _Governance.contract, event: "Deposited", logs: logs, sub: sub}, nil
}

// WatchDeposited is a free log subscription operation binding the contract event 0x2da466a7b24304f47e87fa2e1e5a81b9831ce54fec19055ce277ca2f39ba42c4.
//
// Solidity: event Deposited(address _sealer, uint256 _amount)
func (_Governance *GovernanceFilterer) WatchDeposited(opts *bind.WatchOpts, sink chan<- *GovernanceDeposited) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "Deposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceDeposited)
				if err := _Governance.contract.UnpackLog(event, "Deposited", log); err != nil {
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
func (_Governance *GovernanceFilterer) ParseDeposited(log types.Log) (*GovernanceDeposited, error) {
	event := new(GovernanceDeposited)
	if err := _Governance.contract.UnpackLog(event, "Deposited", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceJoinedIterator is returned from FilterJoined and is used to iterate over the raw logs and unpacked data for Joined events raised by the Governance contract.
type GovernanceJoinedIterator struct {
	Event *GovernanceJoined // Event containing the contract specifics and raw log

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
func (it *GovernanceJoinedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceJoined)
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
		it.Event = new(GovernanceJoined)
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
func (it *GovernanceJoinedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceJoinedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceJoined represents a Joined event raised by the Governance contract.
type GovernanceJoined struct {
	Sealer common.Address
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterJoined is a free log retrieval operation binding the contract event 0x7702dccda75540ad1dca8d5276c048f4a5c0e4203f6da4be214bfb1901b203ea.
//
// Solidity: event Joined(address _sealer, address _signer)
func (_Governance *GovernanceFilterer) FilterJoined(opts *bind.FilterOpts) (*GovernanceJoinedIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "Joined")
	if err != nil {
		return nil, err
	}
	return &GovernanceJoinedIterator{contract: _Governance.contract, event: "Joined", logs: logs, sub: sub}, nil
}

// WatchJoined is a free log subscription operation binding the contract event 0x7702dccda75540ad1dca8d5276c048f4a5c0e4203f6da4be214bfb1901b203ea.
//
// Solidity: event Joined(address _sealer, address _signer)
func (_Governance *GovernanceFilterer) WatchJoined(opts *bind.WatchOpts, sink chan<- *GovernanceJoined) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "Joined")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceJoined)
				if err := _Governance.contract.UnpackLog(event, "Joined", log); err != nil {
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
func (_Governance *GovernanceFilterer) ParseJoined(log types.Log) (*GovernanceJoined, error) {
	event := new(GovernanceJoined)
	if err := _Governance.contract.UnpackLog(event, "Joined", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceLeftIterator is returned from FilterLeft and is used to iterate over the raw logs and unpacked data for Left events raised by the Governance contract.
type GovernanceLeftIterator struct {
	Event *GovernanceLeft // Event containing the contract specifics and raw log

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
func (it *GovernanceLeftIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceLeft)
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
		it.Event = new(GovernanceLeft)
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
func (it *GovernanceLeftIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceLeftIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceLeft represents a Left event raised by the Governance contract.
type GovernanceLeft struct {
	Sealer common.Address
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLeft is a free log retrieval operation binding the contract event 0x4b9ee4dd061ba088b22898a02491f3896a4a580c6cda8783ca579ee159f8e8c5.
//
// Solidity: event Left(address _sealer, address _signer)
func (_Governance *GovernanceFilterer) FilterLeft(opts *bind.FilterOpts) (*GovernanceLeftIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "Left")
	if err != nil {
		return nil, err
	}
	return &GovernanceLeftIterator{contract: _Governance.contract, event: "Left", logs: logs, sub: sub}, nil
}

// WatchLeft is a free log subscription operation binding the contract event 0x4b9ee4dd061ba088b22898a02491f3896a4a580c6cda8783ca579ee159f8e8c5.
//
// Solidity: event Left(address _sealer, address _signer)
func (_Governance *GovernanceFilterer) WatchLeft(opts *bind.WatchOpts, sink chan<- *GovernanceLeft) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "Left")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceLeft)
				if err := _Governance.contract.UnpackLog(event, "Left", log); err != nil {
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
func (_Governance *GovernanceFilterer) ParseLeft(log types.Log) (*GovernanceLeft, error) {
	event := new(GovernanceLeft)
	if err := _Governance.contract.UnpackLog(event, "Left", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceUnbannedIterator is returned from FilterUnbanned and is used to iterate over the raw logs and unpacked data for Unbanned events raised by the Governance contract.
type GovernanceUnbannedIterator struct {
	Event *GovernanceUnbanned // Event containing the contract specifics and raw log

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
func (it *GovernanceUnbannedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceUnbanned)
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
		it.Event = new(GovernanceUnbanned)
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
func (it *GovernanceUnbannedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceUnbannedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceUnbanned represents a Unbanned event raised by the Governance contract.
type GovernanceUnbanned struct {
	Sealer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnbanned is a free log retrieval operation binding the contract event 0x2ab91b53354938415bb6962c4322231cd4cb2c84930f1a4b9abbedc2fe8abe72.
//
// Solidity: event Unbanned(address _sealer)
func (_Governance *GovernanceFilterer) FilterUnbanned(opts *bind.FilterOpts) (*GovernanceUnbannedIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "Unbanned")
	if err != nil {
		return nil, err
	}
	return &GovernanceUnbannedIterator{contract: _Governance.contract, event: "Unbanned", logs: logs, sub: sub}, nil
}

// WatchUnbanned is a free log subscription operation binding the contract event 0x2ab91b53354938415bb6962c4322231cd4cb2c84930f1a4b9abbedc2fe8abe72.
//
// Solidity: event Unbanned(address _sealer)
func (_Governance *GovernanceFilterer) WatchUnbanned(opts *bind.WatchOpts, sink chan<- *GovernanceUnbanned) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "Unbanned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceUnbanned)
				if err := _Governance.contract.UnpackLog(event, "Unbanned", log); err != nil {
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
func (_Governance *GovernanceFilterer) ParseUnbanned(log types.Log) (*GovernanceUnbanned, error) {
	event := new(GovernanceUnbanned)
	if err := _Governance.contract.UnpackLog(event, "Unbanned", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GovernanceWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Governance contract.
type GovernanceWithdrawnIterator struct {
	Event *GovernanceWithdrawn // Event containing the contract specifics and raw log

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
func (it *GovernanceWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GovernanceWithdrawn)
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
		it.Event = new(GovernanceWithdrawn)
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
func (it *GovernanceWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GovernanceWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GovernanceWithdrawn represents a Withdrawn event raised by the Governance contract.
type GovernanceWithdrawn struct {
	Sealer common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address _sealer, uint256 _amount)
func (_Governance *GovernanceFilterer) FilterWithdrawn(opts *bind.FilterOpts) (*GovernanceWithdrawnIterator, error) {

	logs, sub, err := _Governance.contract.FilterLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return &GovernanceWithdrawnIterator{contract: _Governance.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address _sealer, uint256 _amount)
func (_Governance *GovernanceFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *GovernanceWithdrawn) (event.Subscription, error) {

	logs, sub, err := _Governance.contract.WatchLogs(opts, "Withdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GovernanceWithdrawn)
				if err := _Governance.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
func (_Governance *GovernanceFilterer) ParseWithdrawn(log types.Log) (*GovernanceWithdrawn, error) {
	event := new(GovernanceWithdrawn)
	if err := _Governance.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212203ac7f9d8cacaf5cdc3ed3d09bc6d450bc3895e2315d9b676ffe418b80d701f0764736f6c63430006080033"

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
