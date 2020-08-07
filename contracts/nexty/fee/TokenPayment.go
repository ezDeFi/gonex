// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fee

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

// IConfigABI is the input ABI used to generate the binding from.
const IConfigABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"getTokenFor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IConfigFuncSigs maps the 4-byte function signature to its string representation.
var IConfigFuncSigs = map[string]string{
	"41976e09": "getPrice(address)",
	"c0e04395": "getTokenFor(address)",
}

// IConfig is an auto generated Go binding around an Ethereum contract.
type IConfig struct {
	IConfigCaller     // Read-only binding to the contract
	IConfigTransactor // Write-only binding to the contract
	IConfigFilterer   // Log filterer for contract events
}

// IConfigCaller is an auto generated read-only Go binding around an Ethereum contract.
type IConfigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IConfigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IConfigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IConfigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IConfigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IConfigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IConfigSession struct {
	Contract     *IConfig          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IConfigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IConfigCallerSession struct {
	Contract *IConfigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IConfigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IConfigTransactorSession struct {
	Contract     *IConfigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IConfigRaw is an auto generated low-level Go binding around an Ethereum contract.
type IConfigRaw struct {
	Contract *IConfig // Generic contract binding to access the raw methods on
}

// IConfigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IConfigCallerRaw struct {
	Contract *IConfigCaller // Generic read-only contract binding to access the raw methods on
}

// IConfigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IConfigTransactorRaw struct {
	Contract *IConfigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIConfig creates a new instance of IConfig, bound to a specific deployed contract.
func NewIConfig(address common.Address, backend bind.ContractBackend) (*IConfig, error) {
	contract, err := bindIConfig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IConfig{IConfigCaller: IConfigCaller{contract: contract}, IConfigTransactor: IConfigTransactor{contract: contract}, IConfigFilterer: IConfigFilterer{contract: contract}}, nil
}

// NewIConfigCaller creates a new read-only instance of IConfig, bound to a specific deployed contract.
func NewIConfigCaller(address common.Address, caller bind.ContractCaller) (*IConfigCaller, error) {
	contract, err := bindIConfig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IConfigCaller{contract: contract}, nil
}

// NewIConfigTransactor creates a new write-only instance of IConfig, bound to a specific deployed contract.
func NewIConfigTransactor(address common.Address, transactor bind.ContractTransactor) (*IConfigTransactor, error) {
	contract, err := bindIConfig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IConfigTransactor{contract: contract}, nil
}

// NewIConfigFilterer creates a new log filterer instance of IConfig, bound to a specific deployed contract.
func NewIConfigFilterer(address common.Address, filterer bind.ContractFilterer) (*IConfigFilterer, error) {
	contract, err := bindIConfig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IConfigFilterer{contract: contract}, nil
}

// bindIConfig binds a generic wrapper to an already deployed contract.
func bindIConfig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IConfigABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IConfig *IConfigRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IConfig.Contract.IConfigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IConfig *IConfigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IConfig.Contract.IConfigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IConfig *IConfigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IConfig.Contract.IConfigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IConfig *IConfigCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IConfig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IConfig *IConfigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IConfig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IConfig *IConfigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IConfig.Contract.contract.Transact(opts, method, params...)
}

// GetPrice is a paid mutator transaction binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) returns(uint256 price)
func (_IConfig *IConfigTransactor) GetPrice(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _IConfig.contract.Transact(opts, "getPrice", token)
}

// GetPrice is a paid mutator transaction binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) returns(uint256 price)
func (_IConfig *IConfigSession) GetPrice(token common.Address) (*types.Transaction, error) {
	return _IConfig.Contract.GetPrice(&_IConfig.TransactOpts, token)
}

// GetPrice is a paid mutator transaction binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) returns(uint256 price)
func (_IConfig *IConfigTransactorSession) GetPrice(token common.Address) (*types.Transaction, error) {
	return _IConfig.Contract.GetPrice(&_IConfig.TransactOpts, token)
}

// GetTokenFor is a paid mutator transaction binding the contract method 0xc0e04395.
//
// Solidity: function getTokenFor(address to) returns(address token)
func (_IConfig *IConfigTransactor) GetTokenFor(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _IConfig.contract.Transact(opts, "getTokenFor", to)
}

// GetTokenFor is a paid mutator transaction binding the contract method 0xc0e04395.
//
// Solidity: function getTokenFor(address to) returns(address token)
func (_IConfig *IConfigSession) GetTokenFor(to common.Address) (*types.Transaction, error) {
	return _IConfig.Contract.GetTokenFor(&_IConfig.TransactOpts, to)
}

// GetTokenFor is a paid mutator transaction binding the contract method 0xc0e04395.
//
// Solidity: function getTokenFor(address to) returns(address token)
func (_IConfig *IConfigTransactorSession) GetTokenFor(to common.Address) (*types.Transaction, error) {
	return _IConfig.Contract.GetTokenFor(&_IConfig.TransactOpts, to)
}

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a paid mutator transaction binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) returns(uint256)
func (_IERC20 *IERC20Transactor) Allowance(opts *bind.TransactOpts, owner common.Address, spender common.Address) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "allowance", owner, spender)
}

// Allowance is a paid mutator transaction binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*types.Transaction, error) {
	return _IERC20.Contract.Allowance(&_IERC20.TransactOpts, owner, spender)
}

// Allowance is a paid mutator transaction binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) returns(uint256)
func (_IERC20 *IERC20TransactorSession) Allowance(owner common.Address, spender common.Address) (*types.Transaction, error) {
	return _IERC20.Contract.Allowance(&_IERC20.TransactOpts, owner, spender)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, value)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) returns(uint256)
func (_IERC20 *IERC20Transactor) BalanceOf(opts *bind.TransactOpts, who common.Address) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "balanceOf", who)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(who common.Address) (*types.Transaction, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.TransactOpts, who)
}

// BalanceOf is a paid mutator transaction binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address who) returns(uint256)
func (_IERC20 *IERC20TransactorSession) BalanceOf(who common.Address) (*types.Transaction, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.TransactOpts, who)
}

// TotalSupply is a paid mutator transaction binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_IERC20 *IERC20Transactor) TotalSupply(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "totalSupply")
}

// TotalSupply is a paid mutator transaction binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*types.Transaction, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.TransactOpts)
}

// TotalSupply is a paid mutator transaction binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() returns(uint256)
func (_IERC20 *IERC20TransactorSession) TotalSupply() (*types.Transaction, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, from, to, value)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IPayerABI is the input ABI used to generate the binding from.
const IPayerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"txTo\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"txGasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasToPay\",\"type\":\"uint256\"}],\"name\":\"pay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IPayerFuncSigs maps the 4-byte function signature to its string representation.
var IPayerFuncSigs = map[string]string{
	"dd9a76ff": "pay(address,address,uint256,uint256)",
}

// IPayer is an auto generated Go binding around an Ethereum contract.
type IPayer struct {
	IPayerCaller     // Read-only binding to the contract
	IPayerTransactor // Write-only binding to the contract
	IPayerFilterer   // Log filterer for contract events
}

// IPayerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IPayerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPayerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IPayerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPayerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IPayerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IPayerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IPayerSession struct {
	Contract     *IPayer           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPayerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IPayerCallerSession struct {
	Contract *IPayerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IPayerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IPayerTransactorSession struct {
	Contract     *IPayerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IPayerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IPayerRaw struct {
	Contract *IPayer // Generic contract binding to access the raw methods on
}

// IPayerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IPayerCallerRaw struct {
	Contract *IPayerCaller // Generic read-only contract binding to access the raw methods on
}

// IPayerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IPayerTransactorRaw struct {
	Contract *IPayerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIPayer creates a new instance of IPayer, bound to a specific deployed contract.
func NewIPayer(address common.Address, backend bind.ContractBackend) (*IPayer, error) {
	contract, err := bindIPayer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IPayer{IPayerCaller: IPayerCaller{contract: contract}, IPayerTransactor: IPayerTransactor{contract: contract}, IPayerFilterer: IPayerFilterer{contract: contract}}, nil
}

// NewIPayerCaller creates a new read-only instance of IPayer, bound to a specific deployed contract.
func NewIPayerCaller(address common.Address, caller bind.ContractCaller) (*IPayerCaller, error) {
	contract, err := bindIPayer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IPayerCaller{contract: contract}, nil
}

// NewIPayerTransactor creates a new write-only instance of IPayer, bound to a specific deployed contract.
func NewIPayerTransactor(address common.Address, transactor bind.ContractTransactor) (*IPayerTransactor, error) {
	contract, err := bindIPayer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IPayerTransactor{contract: contract}, nil
}

// NewIPayerFilterer creates a new log filterer instance of IPayer, bound to a specific deployed contract.
func NewIPayerFilterer(address common.Address, filterer bind.ContractFilterer) (*IPayerFilterer, error) {
	contract, err := bindIPayer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IPayerFilterer{contract: contract}, nil
}

// bindIPayer binds a generic wrapper to an already deployed contract.
func bindIPayer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IPayerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPayer *IPayerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IPayer.Contract.IPayerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPayer *IPayerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPayer.Contract.IPayerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPayer *IPayerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPayer.Contract.IPayerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IPayer *IPayerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IPayer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IPayer *IPayerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IPayer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IPayer *IPayerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IPayer.Contract.contract.Transact(opts, method, params...)
}

// Pay is a paid mutator transaction binding the contract method 0xdd9a76ff.
//
// Solidity: function pay(address coinbase, address txTo, uint256 txGasPrice, uint256 gasToPay) returns()
func (_IPayer *IPayerTransactor) Pay(opts *bind.TransactOpts, coinbase common.Address, txTo common.Address, txGasPrice *big.Int, gasToPay *big.Int) (*types.Transaction, error) {
	return _IPayer.contract.Transact(opts, "pay", coinbase, txTo, txGasPrice, gasToPay)
}

// Pay is a paid mutator transaction binding the contract method 0xdd9a76ff.
//
// Solidity: function pay(address coinbase, address txTo, uint256 txGasPrice, uint256 gasToPay) returns()
func (_IPayer *IPayerSession) Pay(coinbase common.Address, txTo common.Address, txGasPrice *big.Int, gasToPay *big.Int) (*types.Transaction, error) {
	return _IPayer.Contract.Pay(&_IPayer.TransactOpts, coinbase, txTo, txGasPrice, gasToPay)
}

// Pay is a paid mutator transaction binding the contract method 0xdd9a76ff.
//
// Solidity: function pay(address coinbase, address txTo, uint256 txGasPrice, uint256 gasToPay) returns()
func (_IPayer *IPayerTransactorSession) Pay(coinbase common.Address, txTo common.Address, txGasPrice *big.Int, gasToPay *big.Int) (*types.Transaction, error) {
	return _IPayer.Contract.Pay(&_IPayer.TransactOpts, coinbase, txTo, txGasPrice, gasToPay)
}

// PayerCodeABI is the input ABI used to generate the binding from.
const PayerCodeABI = "[{\"inputs\":[],\"name\":\"FeeTokenFallbackPath\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FeeTokenPath\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TokenPayment\",\"outputs\":[{\"internalType\":\"contractIConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TokenPricePath\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"txTo\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"txGasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasToPay\",\"type\":\"uint256\"}],\"name\":\"pay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PayerCodeFuncSigs maps the 4-byte function signature to its string representation.
var PayerCodeFuncSigs = map[string]string{
	"5e8f4d74": "FeeTokenFallbackPath()",
	"c8c1956b": "FeeTokenPath()",
	"f218dd97": "TokenPayment()",
	"c9969bcf": "TokenPricePath()",
	"dd9a76ff": "pay(address,address,uint256,uint256)",
}

// PayerCode is an auto generated Go binding around an Ethereum contract.
type PayerCode struct {
	PayerCodeCaller     // Read-only binding to the contract
	PayerCodeTransactor // Write-only binding to the contract
	PayerCodeFilterer   // Log filterer for contract events
}

// PayerCodeCaller is an auto generated read-only Go binding around an Ethereum contract.
type PayerCodeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PayerCodeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PayerCodeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PayerCodeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PayerCodeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PayerCodeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PayerCodeSession struct {
	Contract     *PayerCode        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PayerCodeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PayerCodeCallerSession struct {
	Contract *PayerCodeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// PayerCodeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PayerCodeTransactorSession struct {
	Contract     *PayerCodeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// PayerCodeRaw is an auto generated low-level Go binding around an Ethereum contract.
type PayerCodeRaw struct {
	Contract *PayerCode // Generic contract binding to access the raw methods on
}

// PayerCodeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PayerCodeCallerRaw struct {
	Contract *PayerCodeCaller // Generic read-only contract binding to access the raw methods on
}

// PayerCodeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PayerCodeTransactorRaw struct {
	Contract *PayerCodeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPayerCode creates a new instance of PayerCode, bound to a specific deployed contract.
func NewPayerCode(address common.Address, backend bind.ContractBackend) (*PayerCode, error) {
	contract, err := bindPayerCode(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PayerCode{PayerCodeCaller: PayerCodeCaller{contract: contract}, PayerCodeTransactor: PayerCodeTransactor{contract: contract}, PayerCodeFilterer: PayerCodeFilterer{contract: contract}}, nil
}

// NewPayerCodeCaller creates a new read-only instance of PayerCode, bound to a specific deployed contract.
func NewPayerCodeCaller(address common.Address, caller bind.ContractCaller) (*PayerCodeCaller, error) {
	contract, err := bindPayerCode(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PayerCodeCaller{contract: contract}, nil
}

// NewPayerCodeTransactor creates a new write-only instance of PayerCode, bound to a specific deployed contract.
func NewPayerCodeTransactor(address common.Address, transactor bind.ContractTransactor) (*PayerCodeTransactor, error) {
	contract, err := bindPayerCode(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PayerCodeTransactor{contract: contract}, nil
}

// NewPayerCodeFilterer creates a new log filterer instance of PayerCode, bound to a specific deployed contract.
func NewPayerCodeFilterer(address common.Address, filterer bind.ContractFilterer) (*PayerCodeFilterer, error) {
	contract, err := bindPayerCode(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PayerCodeFilterer{contract: contract}, nil
}

// bindPayerCode binds a generic wrapper to an already deployed contract.
func bindPayerCode(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PayerCodeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PayerCode *PayerCodeRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PayerCode.Contract.PayerCodeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PayerCode *PayerCodeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PayerCode.Contract.PayerCodeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PayerCode *PayerCodeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PayerCode.Contract.PayerCodeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PayerCode *PayerCodeCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PayerCode.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PayerCode *PayerCodeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PayerCode.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PayerCode *PayerCodeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PayerCode.Contract.contract.Transact(opts, method, params...)
}

// FeeTokenFallbackPath is a paid mutator transaction binding the contract method 0x5e8f4d74.
//
// Solidity: function FeeTokenFallbackPath() returns(bytes32)
func (_PayerCode *PayerCodeTransactor) FeeTokenFallbackPath(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PayerCode.contract.Transact(opts, "FeeTokenFallbackPath")
}

// FeeTokenFallbackPath is a paid mutator transaction binding the contract method 0x5e8f4d74.
//
// Solidity: function FeeTokenFallbackPath() returns(bytes32)
func (_PayerCode *PayerCodeSession) FeeTokenFallbackPath() (*types.Transaction, error) {
	return _PayerCode.Contract.FeeTokenFallbackPath(&_PayerCode.TransactOpts)
}

// FeeTokenFallbackPath is a paid mutator transaction binding the contract method 0x5e8f4d74.
//
// Solidity: function FeeTokenFallbackPath() returns(bytes32)
func (_PayerCode *PayerCodeTransactorSession) FeeTokenFallbackPath() (*types.Transaction, error) {
	return _PayerCode.Contract.FeeTokenFallbackPath(&_PayerCode.TransactOpts)
}

// FeeTokenPath is a paid mutator transaction binding the contract method 0xc8c1956b.
//
// Solidity: function FeeTokenPath() returns(bytes32)
func (_PayerCode *PayerCodeTransactor) FeeTokenPath(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PayerCode.contract.Transact(opts, "FeeTokenPath")
}

// FeeTokenPath is a paid mutator transaction binding the contract method 0xc8c1956b.
//
// Solidity: function FeeTokenPath() returns(bytes32)
func (_PayerCode *PayerCodeSession) FeeTokenPath() (*types.Transaction, error) {
	return _PayerCode.Contract.FeeTokenPath(&_PayerCode.TransactOpts)
}

// FeeTokenPath is a paid mutator transaction binding the contract method 0xc8c1956b.
//
// Solidity: function FeeTokenPath() returns(bytes32)
func (_PayerCode *PayerCodeTransactorSession) FeeTokenPath() (*types.Transaction, error) {
	return _PayerCode.Contract.FeeTokenPath(&_PayerCode.TransactOpts)
}

// TokenPayment is a paid mutator transaction binding the contract method 0xf218dd97.
//
// Solidity: function TokenPayment() returns(address)
func (_PayerCode *PayerCodeTransactor) TokenPayment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PayerCode.contract.Transact(opts, "TokenPayment")
}

// TokenPayment is a paid mutator transaction binding the contract method 0xf218dd97.
//
// Solidity: function TokenPayment() returns(address)
func (_PayerCode *PayerCodeSession) TokenPayment() (*types.Transaction, error) {
	return _PayerCode.Contract.TokenPayment(&_PayerCode.TransactOpts)
}

// TokenPayment is a paid mutator transaction binding the contract method 0xf218dd97.
//
// Solidity: function TokenPayment() returns(address)
func (_PayerCode *PayerCodeTransactorSession) TokenPayment() (*types.Transaction, error) {
	return _PayerCode.Contract.TokenPayment(&_PayerCode.TransactOpts)
}

// TokenPricePath is a paid mutator transaction binding the contract method 0xc9969bcf.
//
// Solidity: function TokenPricePath() returns(bytes32)
func (_PayerCode *PayerCodeTransactor) TokenPricePath(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PayerCode.contract.Transact(opts, "TokenPricePath")
}

// TokenPricePath is a paid mutator transaction binding the contract method 0xc9969bcf.
//
// Solidity: function TokenPricePath() returns(bytes32)
func (_PayerCode *PayerCodeSession) TokenPricePath() (*types.Transaction, error) {
	return _PayerCode.Contract.TokenPricePath(&_PayerCode.TransactOpts)
}

// TokenPricePath is a paid mutator transaction binding the contract method 0xc9969bcf.
//
// Solidity: function TokenPricePath() returns(bytes32)
func (_PayerCode *PayerCodeTransactorSession) TokenPricePath() (*types.Transaction, error) {
	return _PayerCode.Contract.TokenPricePath(&_PayerCode.TransactOpts)
}

// Pay is a paid mutator transaction binding the contract method 0xdd9a76ff.
//
// Solidity: function pay(address coinbase, address txTo, uint256 txGasPrice, uint256 gasToPay) returns()
func (_PayerCode *PayerCodeTransactor) Pay(opts *bind.TransactOpts, coinbase common.Address, txTo common.Address, txGasPrice *big.Int, gasToPay *big.Int) (*types.Transaction, error) {
	return _PayerCode.contract.Transact(opts, "pay", coinbase, txTo, txGasPrice, gasToPay)
}

// Pay is a paid mutator transaction binding the contract method 0xdd9a76ff.
//
// Solidity: function pay(address coinbase, address txTo, uint256 txGasPrice, uint256 gasToPay) returns()
func (_PayerCode *PayerCodeSession) Pay(coinbase common.Address, txTo common.Address, txGasPrice *big.Int, gasToPay *big.Int) (*types.Transaction, error) {
	return _PayerCode.Contract.Pay(&_PayerCode.TransactOpts, coinbase, txTo, txGasPrice, gasToPay)
}

// Pay is a paid mutator transaction binding the contract method 0xdd9a76ff.
//
// Solidity: function pay(address coinbase, address txTo, uint256 txGasPrice, uint256 gasToPay) returns()
func (_PayerCode *PayerCodeTransactorSession) Pay(coinbase common.Address, txTo common.Address, txGasPrice *big.Int, gasToPay *big.Int) (*types.Transaction, error) {
	return _PayerCode.Contract.Pay(&_PayerCode.TransactOpts, coinbase, txTo, txGasPrice, gasToPay)
}

// TokenPaymentABI is the input ABI used to generate the binding from.
const TokenPaymentABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"team\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"FeeTokenFallbackPath\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FeeTokenPath\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TRUE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TeamMemberPath\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TokenForPath\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TokenPayment\",\"outputs\":[{\"internalType\":\"contractIConfig\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TokenPricePath\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"getTokenFor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"txTo\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"txGasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasToPay\",\"type\":\"uint256\"}],\"name\":\"pay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"setTokenFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TokenPaymentFuncSigs maps the 4-byte function signature to its string representation.
var TokenPaymentFuncSigs = map[string]string{
	"5e8f4d74": "FeeTokenFallbackPath()",
	"c8c1956b": "FeeTokenPath()",
	"5994d984": "TRUE()",
	"654b8c2a": "TeamMemberPath()",
	"efc0a0a9": "TokenForPath()",
	"f218dd97": "TokenPayment()",
	"c9969bcf": "TokenPricePath()",
	"41976e09": "getPrice(address)",
	"c0e04395": "getTokenFor(address)",
	"dd9a76ff": "pay(address,address,uint256,uint256)",
	"00e4768b": "setPrice(address,uint256)",
	"90b63f21": "setTokenFor(address,address)",
}

// TokenPaymentBin is the compiled bytecode used for deploying new contracts.
var TokenPaymentBin = "0x60806040523480156200001157600080fd5b5060405162000a6138038062000a61833981810160405260608110156200003757600080fd5b81019080805160405193929190846401000000008211156200005857600080fd5b9083019060208201858111156200006e57600080fd5b82518660208202830111640100000000821117156200008c57600080fd5b82525081516020918201928201910280838360005b83811015620000bb578181015183820152602001620000a1565b5050505090500160405260200180516040519392919084640100000000821115620000e557600080fd5b908301906020820185811115620000fb57600080fd5b82518660208202830111640100000000821117156200011957600080fd5b82525081516020918201928201910280838360005b83811015620001485781810151838201526020016200012e565b50505050905001604052602001805160405193929190846401000000008211156200017257600080fd5b9083019060208201858111156200018857600080fd5b8251866020820283011164010000000082111715620001a657600080fd5b82525081516020918201928201910280838360005b83811015620001d5578181015183820152602001620001bb565b5050505090500160405250505060008090505b83518110156200022057620002178482815181106200020357fe5b60200260200101516200027b60201b60201c565b600101620001e8565b5060005b82518110156200027157620002688382815181106200023f57fe5b60200260200101518383815181106200025457fe5b6020026020010151620002be60201b60201c565b60010162000224565b5050505062000315565b620002bb620002a46a5465616d4d656d6265722d60a81b836200030060201b620004a11760201c565b60001960001b6200031160201b620004b21760201c565b50565b620002fc620002e76a546f6b656e50726963652d60a81b846200030060201b620004a11760201c565b8260001b6200031160201b620004b21760201c565b5050565b60081b610100600160a81b03161790565b9055565b61073c80620003256000396000f3fe608060405234801561001057600080fd5b50600436106100b35760003560e01c8063c0e0439511610071578063c0e0439514610164578063c8c1956b146101a6578063c9969bcf146101ae578063dd9a76ff146101b6578063efc0a0a9146101f2578063f218dd97146101fa576100b3565b8062e4768b146100b857806341976e09146100e65780635994d9841461011e5780635e8f4d7414610126578063654b8c2a1461012e57806390b63f2114610136575b600080fd5b6100e4600480360360408110156100ce57600080fd5b506001600160a01b038135169060200135610202565b005b61010c600480360360208110156100fc57600080fd5b50356001600160a01b0316610261565b60408051918252519081900360200190f35b61010c61028b565b61010c610291565b61010c6102a8565b6100e46004803603604081101561014c57600080fd5b506001600160a01b03813581169160200135166102ba565b61018a6004803603602081101561017a57600080fd5b50356001600160a01b031661032c565b604080516001600160a01b039092168252519081900360200190f35b61010c610351565b61010c610360565b6100e4600480360360808110156101cc57600080fd5b506001600160a01b03813581169160208101359091169060408101359060600135610372565b61010c610488565b61018a61049a565b61020b336104b6565b610253576040805162461bcd60e51b8152602060048201526014602482015273666f72207465616d206d656d626572206f6e6c7960601b604482015290519081900360640190fd5b61025d82826104db565b5050565b600061028361027e6a546f6b656e50726963652d60a81b846104a1565b6104f6565b90505b919050565b60001981565b6f466565546f6b656e46616c6c6261636b60801b81565b6a5465616d4d656d6265722d60a81b81565b6102c3336104b6565b61030b576040805162461bcd60e51b8152602060048201526014602482015273666f72207465616d206d656d626572206f6e6c7960601b604482015290519081900360640190fd5b61025d6103266a546f6b656e466f722d2d2d60a81b846104a1565b826104b2565b6000610337826104fa565b90506001600160a01b0381161561034d57610286565b5090565b672332b2aa37b5b2b760c11b81565b6a546f6b656e50726963652d60a81b81565b60008061037e85610517565b91509150600081116103d7576040805162461bcd60e51b815260206004820152601b60248201527f7061796d656e7420746f6b656e207072696365206e6f74207365740000000000604482015290519081900360640190fd5b828402600082670de0b6b3a76400008302816103ef57fe5b049050836001600160a01b031663a9059cbb89836040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050602060405180830381600087803b15801561045257600080fd5b505af1158015610466573d6000803e3d6000fd5b505050506040513d602081101561047c57600080fd5b50505050505050505050565b6a546f6b656e466f722d2d2d60a81b81565b6205678981565b60081b610100600160a81b03161790565b9055565b60006104d361027e6a5465616d4d656d6265722d60a81b846104a1565b151592915050565b61025d6103266a546f6b656e50726963652d60a81b846104a1565b5490565b600061028361027e6a546f6b656e466f722d2d2d60a81b846104a1565b6000806105238361058a565b91506001600160a01b038216610577576040805162461bcd60e51b81526020600482015260146024820152731b9bc81d1bdad95b88199bdc881c185e5b595b9d60621b604482015290519081900360640190fd5b816105818361065d565b91509150915091565b60006105a0672332b2aa37b5b2b760c11b6104f6565b90506001600160a01b038116156105b657610286565b6040805163c0e0439560e01b81526001600160a01b03841660048201529051620567899163c0e04395916024808301926020929190829003018186803b1580156105ff57600080fd5b505afa158015610613573d6000803e3d6000fd5b505050506040513d602081101561062957600080fd5b505190506001600160a01b0381161561064157610286565b6102836f466565546f6b656e46616c6c6261636b60801b6104f6565b60008061067b61027e6a546f6b656e50726963652d60a81b856104a1565b9050801561068a579050610286565b604080516341976e0960e01b81526001600160a01b0385166004820152905162056789916341976e09916024808301926020929190829003018186803b1580156106d357600080fd5b505afa1580156106e7573d6000803e3d6000fd5b505050506040513d60208110156106fd57600080fd5b5051939250505056fea2646970667358221220cd765c25a8c4ef2b90a13ab01c29dbd6e1862909449c271081de0189820dfd7d64736f6c63430006080033"

// DeployTokenPayment deploys a new Ethereum contract, binding an instance of TokenPayment to it.
func DeployTokenPayment(auth *bind.TransactOpts, backend bind.ContractBackend, team []common.Address, tokens []common.Address, prices []*big.Int) (common.Address, *types.Transaction, *TokenPayment, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenPaymentABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenPaymentBin), backend, team, tokens, prices)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenPayment{TokenPaymentCaller: TokenPaymentCaller{contract: contract}, TokenPaymentTransactor: TokenPaymentTransactor{contract: contract}, TokenPaymentFilterer: TokenPaymentFilterer{contract: contract}}, nil
}

// TokenPayment is an auto generated Go binding around an Ethereum contract.
type TokenPayment struct {
	TokenPaymentCaller     // Read-only binding to the contract
	TokenPaymentTransactor // Write-only binding to the contract
	TokenPaymentFilterer   // Log filterer for contract events
}

// TokenPaymentCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenPaymentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenPaymentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenPaymentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenPaymentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenPaymentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenPaymentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenPaymentSession struct {
	Contract     *TokenPayment     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenPaymentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenPaymentCallerSession struct {
	Contract *TokenPaymentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// TokenPaymentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenPaymentTransactorSession struct {
	Contract     *TokenPaymentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// TokenPaymentRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenPaymentRaw struct {
	Contract *TokenPayment // Generic contract binding to access the raw methods on
}

// TokenPaymentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenPaymentCallerRaw struct {
	Contract *TokenPaymentCaller // Generic read-only contract binding to access the raw methods on
}

// TokenPaymentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenPaymentTransactorRaw struct {
	Contract *TokenPaymentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenPayment creates a new instance of TokenPayment, bound to a specific deployed contract.
func NewTokenPayment(address common.Address, backend bind.ContractBackend) (*TokenPayment, error) {
	contract, err := bindTokenPayment(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenPayment{TokenPaymentCaller: TokenPaymentCaller{contract: contract}, TokenPaymentTransactor: TokenPaymentTransactor{contract: contract}, TokenPaymentFilterer: TokenPaymentFilterer{contract: contract}}, nil
}

// NewTokenPaymentCaller creates a new read-only instance of TokenPayment, bound to a specific deployed contract.
func NewTokenPaymentCaller(address common.Address, caller bind.ContractCaller) (*TokenPaymentCaller, error) {
	contract, err := bindTokenPayment(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenPaymentCaller{contract: contract}, nil
}

// NewTokenPaymentTransactor creates a new write-only instance of TokenPayment, bound to a specific deployed contract.
func NewTokenPaymentTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenPaymentTransactor, error) {
	contract, err := bindTokenPayment(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenPaymentTransactor{contract: contract}, nil
}

// NewTokenPaymentFilterer creates a new log filterer instance of TokenPayment, bound to a specific deployed contract.
func NewTokenPaymentFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenPaymentFilterer, error) {
	contract, err := bindTokenPayment(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenPaymentFilterer{contract: contract}, nil
}

// bindTokenPayment binds a generic wrapper to an already deployed contract.
func bindTokenPayment(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenPaymentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenPayment *TokenPaymentRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenPayment.Contract.TokenPaymentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenPayment *TokenPaymentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPayment.Contract.TokenPaymentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenPayment *TokenPaymentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenPayment.Contract.TokenPaymentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenPayment *TokenPaymentCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenPayment.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenPayment *TokenPaymentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPayment.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenPayment *TokenPaymentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenPayment.Contract.contract.Transact(opts, method, params...)
}

// FeeTokenFallbackPath is a paid mutator transaction binding the contract method 0x5e8f4d74.
//
// Solidity: function FeeTokenFallbackPath() returns(bytes32)
func (_TokenPayment *TokenPaymentTransactor) FeeTokenFallbackPath(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPayment.contract.Transact(opts, "FeeTokenFallbackPath")
}

// FeeTokenFallbackPath is a paid mutator transaction binding the contract method 0x5e8f4d74.
//
// Solidity: function FeeTokenFallbackPath() returns(bytes32)
func (_TokenPayment *TokenPaymentSession) FeeTokenFallbackPath() (*types.Transaction, error) {
	return _TokenPayment.Contract.FeeTokenFallbackPath(&_TokenPayment.TransactOpts)
}

// FeeTokenFallbackPath is a paid mutator transaction binding the contract method 0x5e8f4d74.
//
// Solidity: function FeeTokenFallbackPath() returns(bytes32)
func (_TokenPayment *TokenPaymentTransactorSession) FeeTokenFallbackPath() (*types.Transaction, error) {
	return _TokenPayment.Contract.FeeTokenFallbackPath(&_TokenPayment.TransactOpts)
}

// FeeTokenPath is a paid mutator transaction binding the contract method 0xc8c1956b.
//
// Solidity: function FeeTokenPath() returns(bytes32)
func (_TokenPayment *TokenPaymentTransactor) FeeTokenPath(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPayment.contract.Transact(opts, "FeeTokenPath")
}

// FeeTokenPath is a paid mutator transaction binding the contract method 0xc8c1956b.
//
// Solidity: function FeeTokenPath() returns(bytes32)
func (_TokenPayment *TokenPaymentSession) FeeTokenPath() (*types.Transaction, error) {
	return _TokenPayment.Contract.FeeTokenPath(&_TokenPayment.TransactOpts)
}

// FeeTokenPath is a paid mutator transaction binding the contract method 0xc8c1956b.
//
// Solidity: function FeeTokenPath() returns(bytes32)
func (_TokenPayment *TokenPaymentTransactorSession) FeeTokenPath() (*types.Transaction, error) {
	return _TokenPayment.Contract.FeeTokenPath(&_TokenPayment.TransactOpts)
}

// TRUE is a paid mutator transaction binding the contract method 0x5994d984.
//
// Solidity: function TRUE() returns(bytes32)
func (_TokenPayment *TokenPaymentTransactor) TRUE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPayment.contract.Transact(opts, "TRUE")
}

// TRUE is a paid mutator transaction binding the contract method 0x5994d984.
//
// Solidity: function TRUE() returns(bytes32)
func (_TokenPayment *TokenPaymentSession) TRUE() (*types.Transaction, error) {
	return _TokenPayment.Contract.TRUE(&_TokenPayment.TransactOpts)
}

// TRUE is a paid mutator transaction binding the contract method 0x5994d984.
//
// Solidity: function TRUE() returns(bytes32)
func (_TokenPayment *TokenPaymentTransactorSession) TRUE() (*types.Transaction, error) {
	return _TokenPayment.Contract.TRUE(&_TokenPayment.TransactOpts)
}

// TeamMemberPath is a paid mutator transaction binding the contract method 0x654b8c2a.
//
// Solidity: function TeamMemberPath() returns(bytes32)
func (_TokenPayment *TokenPaymentTransactor) TeamMemberPath(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPayment.contract.Transact(opts, "TeamMemberPath")
}

// TeamMemberPath is a paid mutator transaction binding the contract method 0x654b8c2a.
//
// Solidity: function TeamMemberPath() returns(bytes32)
func (_TokenPayment *TokenPaymentSession) TeamMemberPath() (*types.Transaction, error) {
	return _TokenPayment.Contract.TeamMemberPath(&_TokenPayment.TransactOpts)
}

// TeamMemberPath is a paid mutator transaction binding the contract method 0x654b8c2a.
//
// Solidity: function TeamMemberPath() returns(bytes32)
func (_TokenPayment *TokenPaymentTransactorSession) TeamMemberPath() (*types.Transaction, error) {
	return _TokenPayment.Contract.TeamMemberPath(&_TokenPayment.TransactOpts)
}

// TokenForPath is a paid mutator transaction binding the contract method 0xefc0a0a9.
//
// Solidity: function TokenForPath() returns(bytes32)
func (_TokenPayment *TokenPaymentTransactor) TokenForPath(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPayment.contract.Transact(opts, "TokenForPath")
}

// TokenForPath is a paid mutator transaction binding the contract method 0xefc0a0a9.
//
// Solidity: function TokenForPath() returns(bytes32)
func (_TokenPayment *TokenPaymentSession) TokenForPath() (*types.Transaction, error) {
	return _TokenPayment.Contract.TokenForPath(&_TokenPayment.TransactOpts)
}

// TokenForPath is a paid mutator transaction binding the contract method 0xefc0a0a9.
//
// Solidity: function TokenForPath() returns(bytes32)
func (_TokenPayment *TokenPaymentTransactorSession) TokenForPath() (*types.Transaction, error) {
	return _TokenPayment.Contract.TokenForPath(&_TokenPayment.TransactOpts)
}

// TokenPayment is a paid mutator transaction binding the contract method 0xf218dd97.
//
// Solidity: function TokenPayment() returns(address)
func (_TokenPayment *TokenPaymentTransactor) TokenPayment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPayment.contract.Transact(opts, "TokenPayment")
}

// TokenPayment is a paid mutator transaction binding the contract method 0xf218dd97.
//
// Solidity: function TokenPayment() returns(address)
func (_TokenPayment *TokenPaymentSession) TokenPayment() (*types.Transaction, error) {
	return _TokenPayment.Contract.TokenPayment(&_TokenPayment.TransactOpts)
}

// TokenPayment is a paid mutator transaction binding the contract method 0xf218dd97.
//
// Solidity: function TokenPayment() returns(address)
func (_TokenPayment *TokenPaymentTransactorSession) TokenPayment() (*types.Transaction, error) {
	return _TokenPayment.Contract.TokenPayment(&_TokenPayment.TransactOpts)
}

// TokenPricePath is a paid mutator transaction binding the contract method 0xc9969bcf.
//
// Solidity: function TokenPricePath() returns(bytes32)
func (_TokenPayment *TokenPaymentTransactor) TokenPricePath(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPayment.contract.Transact(opts, "TokenPricePath")
}

// TokenPricePath is a paid mutator transaction binding the contract method 0xc9969bcf.
//
// Solidity: function TokenPricePath() returns(bytes32)
func (_TokenPayment *TokenPaymentSession) TokenPricePath() (*types.Transaction, error) {
	return _TokenPayment.Contract.TokenPricePath(&_TokenPayment.TransactOpts)
}

// TokenPricePath is a paid mutator transaction binding the contract method 0xc9969bcf.
//
// Solidity: function TokenPricePath() returns(bytes32)
func (_TokenPayment *TokenPaymentTransactorSession) TokenPricePath() (*types.Transaction, error) {
	return _TokenPayment.Contract.TokenPricePath(&_TokenPayment.TransactOpts)
}

// GetPrice is a paid mutator transaction binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) returns(uint256 price)
func (_TokenPayment *TokenPaymentTransactor) GetPrice(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _TokenPayment.contract.Transact(opts, "getPrice", token)
}

// GetPrice is a paid mutator transaction binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) returns(uint256 price)
func (_TokenPayment *TokenPaymentSession) GetPrice(token common.Address) (*types.Transaction, error) {
	return _TokenPayment.Contract.GetPrice(&_TokenPayment.TransactOpts, token)
}

// GetPrice is a paid mutator transaction binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) returns(uint256 price)
func (_TokenPayment *TokenPaymentTransactorSession) GetPrice(token common.Address) (*types.Transaction, error) {
	return _TokenPayment.Contract.GetPrice(&_TokenPayment.TransactOpts, token)
}

// GetTokenFor is a paid mutator transaction binding the contract method 0xc0e04395.
//
// Solidity: function getTokenFor(address to) returns(address token)
func (_TokenPayment *TokenPaymentTransactor) GetTokenFor(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _TokenPayment.contract.Transact(opts, "getTokenFor", to)
}

// GetTokenFor is a paid mutator transaction binding the contract method 0xc0e04395.
//
// Solidity: function getTokenFor(address to) returns(address token)
func (_TokenPayment *TokenPaymentSession) GetTokenFor(to common.Address) (*types.Transaction, error) {
	return _TokenPayment.Contract.GetTokenFor(&_TokenPayment.TransactOpts, to)
}

// GetTokenFor is a paid mutator transaction binding the contract method 0xc0e04395.
//
// Solidity: function getTokenFor(address to) returns(address token)
func (_TokenPayment *TokenPaymentTransactorSession) GetTokenFor(to common.Address) (*types.Transaction, error) {
	return _TokenPayment.Contract.GetTokenFor(&_TokenPayment.TransactOpts, to)
}

// Pay is a paid mutator transaction binding the contract method 0xdd9a76ff.
//
// Solidity: function pay(address coinbase, address txTo, uint256 txGasPrice, uint256 gasToPay) returns()
func (_TokenPayment *TokenPaymentTransactor) Pay(opts *bind.TransactOpts, coinbase common.Address, txTo common.Address, txGasPrice *big.Int, gasToPay *big.Int) (*types.Transaction, error) {
	return _TokenPayment.contract.Transact(opts, "pay", coinbase, txTo, txGasPrice, gasToPay)
}

// Pay is a paid mutator transaction binding the contract method 0xdd9a76ff.
//
// Solidity: function pay(address coinbase, address txTo, uint256 txGasPrice, uint256 gasToPay) returns()
func (_TokenPayment *TokenPaymentSession) Pay(coinbase common.Address, txTo common.Address, txGasPrice *big.Int, gasToPay *big.Int) (*types.Transaction, error) {
	return _TokenPayment.Contract.Pay(&_TokenPayment.TransactOpts, coinbase, txTo, txGasPrice, gasToPay)
}

// Pay is a paid mutator transaction binding the contract method 0xdd9a76ff.
//
// Solidity: function pay(address coinbase, address txTo, uint256 txGasPrice, uint256 gasToPay) returns()
func (_TokenPayment *TokenPaymentTransactorSession) Pay(coinbase common.Address, txTo common.Address, txGasPrice *big.Int, gasToPay *big.Int) (*types.Transaction, error) {
	return _TokenPayment.Contract.Pay(&_TokenPayment.TransactOpts, coinbase, txTo, txGasPrice, gasToPay)
}

// SetPrice is a paid mutator transaction binding the contract method 0x00e4768b.
//
// Solidity: function setPrice(address token, uint256 price) returns()
func (_TokenPayment *TokenPaymentTransactor) SetPrice(opts *bind.TransactOpts, token common.Address, price *big.Int) (*types.Transaction, error) {
	return _TokenPayment.contract.Transact(opts, "setPrice", token, price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x00e4768b.
//
// Solidity: function setPrice(address token, uint256 price) returns()
func (_TokenPayment *TokenPaymentSession) SetPrice(token common.Address, price *big.Int) (*types.Transaction, error) {
	return _TokenPayment.Contract.SetPrice(&_TokenPayment.TransactOpts, token, price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x00e4768b.
//
// Solidity: function setPrice(address token, uint256 price) returns()
func (_TokenPayment *TokenPaymentTransactorSession) SetPrice(token common.Address, price *big.Int) (*types.Transaction, error) {
	return _TokenPayment.Contract.SetPrice(&_TokenPayment.TransactOpts, token, price)
}

// SetTokenFor is a paid mutator transaction binding the contract method 0x90b63f21.
//
// Solidity: function setTokenFor(address to, address token) returns()
func (_TokenPayment *TokenPaymentTransactor) SetTokenFor(opts *bind.TransactOpts, to common.Address, token common.Address) (*types.Transaction, error) {
	return _TokenPayment.contract.Transact(opts, "setTokenFor", to, token)
}

// SetTokenFor is a paid mutator transaction binding the contract method 0x90b63f21.
//
// Solidity: function setTokenFor(address to, address token) returns()
func (_TokenPayment *TokenPaymentSession) SetTokenFor(to common.Address, token common.Address) (*types.Transaction, error) {
	return _TokenPayment.Contract.SetTokenFor(&_TokenPayment.TransactOpts, to, token)
}

// SetTokenFor is a paid mutator transaction binding the contract method 0x90b63f21.
//
// Solidity: function setTokenFor(address to, address token) returns()
func (_TokenPayment *TokenPaymentTransactorSession) SetTokenFor(to common.Address, token common.Address) (*types.Transaction, error) {
	return _TokenPayment.Contract.SetTokenFor(&_TokenPayment.TransactOpts, to, token)
}

// DsABI is the input ABI used to generate the binding from.
const DsABI = "[]"

// DsBin is the compiled bytecode used for deploying new contracts.
var DsBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220b49dbed9f3f1f5ee4ca8a479262617c08c6f95b9e7963dc8135a4e44e48a59a864736f6c63430006080033"

// DeployDs deploys a new Ethereum contract, binding an instance of Ds to it.
func DeployDs(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ds, error) {
	parsed, err := abi.JSON(strings.NewReader(DsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ds{DsCaller: DsCaller{contract: contract}, DsTransactor: DsTransactor{contract: contract}, DsFilterer: DsFilterer{contract: contract}}, nil
}

// Ds is an auto generated Go binding around an Ethereum contract.
type Ds struct {
	DsCaller     // Read-only binding to the contract
	DsTransactor // Write-only binding to the contract
	DsFilterer   // Log filterer for contract events
}

// DsCaller is an auto generated read-only Go binding around an Ethereum contract.
type DsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DsSession struct {
	Contract     *Ds               // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DsCallerSession struct {
	Contract *DsCaller     // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// DsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DsTransactorSession struct {
	Contract     *DsTransactor     // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DsRaw is an auto generated low-level Go binding around an Ethereum contract.
type DsRaw struct {
	Contract *Ds // Generic contract binding to access the raw methods on
}

// DsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DsCallerRaw struct {
	Contract *DsCaller // Generic read-only contract binding to access the raw methods on
}

// DsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DsTransactorRaw struct {
	Contract *DsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDs creates a new instance of Ds, bound to a specific deployed contract.
func NewDs(address common.Address, backend bind.ContractBackend) (*Ds, error) {
	contract, err := bindDs(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ds{DsCaller: DsCaller{contract: contract}, DsTransactor: DsTransactor{contract: contract}, DsFilterer: DsFilterer{contract: contract}}, nil
}

// NewDsCaller creates a new read-only instance of Ds, bound to a specific deployed contract.
func NewDsCaller(address common.Address, caller bind.ContractCaller) (*DsCaller, error) {
	contract, err := bindDs(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DsCaller{contract: contract}, nil
}

// NewDsTransactor creates a new write-only instance of Ds, bound to a specific deployed contract.
func NewDsTransactor(address common.Address, transactor bind.ContractTransactor) (*DsTransactor, error) {
	contract, err := bindDs(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DsTransactor{contract: contract}, nil
}

// NewDsFilterer creates a new log filterer instance of Ds, bound to a specific deployed contract.
func NewDsFilterer(address common.Address, filterer bind.ContractFilterer) (*DsFilterer, error) {
	contract, err := bindDs(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DsFilterer{contract: contract}, nil
}

// bindDs binds a generic wrapper to an already deployed contract.
func bindDs(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ds *DsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ds.Contract.DsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ds *DsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ds.Contract.DsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ds *DsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ds.Contract.DsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ds *DsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ds.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ds *DsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ds.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ds *DsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ds.Contract.contract.Transact(opts, method, params...)
}
