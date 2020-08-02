// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package payer

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
const IPayerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"pay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"txGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"txMethodSig\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32[]\",\"name\":\"txMethodParams\",\"type\":\"bytes32[]\"}],\"name\":\"payment\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payGasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IPayerFuncSigs maps the 4-byte function signature to its string representation.
var IPayerFuncSigs = map[string]string{
	"b3d76188": "pay(address,address,uint256)",
	"dd50295e": "payment(address,uint256,bytes4,bytes32[])",
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

// Pay is a paid mutator transaction binding the contract method 0xb3d76188.
//
// Solidity: function pay(address coinbase, address token, uint256 value) returns()
func (_IPayer *IPayerTransactor) Pay(opts *bind.TransactOpts, coinbase common.Address, token common.Address, value *big.Int) (*types.Transaction, error) {
	return _IPayer.contract.Transact(opts, "pay", coinbase, token, value)
}

// Pay is a paid mutator transaction binding the contract method 0xb3d76188.
//
// Solidity: function pay(address coinbase, address token, uint256 value) returns()
func (_IPayer *IPayerSession) Pay(coinbase common.Address, token common.Address, value *big.Int) (*types.Transaction, error) {
	return _IPayer.Contract.Pay(&_IPayer.TransactOpts, coinbase, token, value)
}

// Pay is a paid mutator transaction binding the contract method 0xb3d76188.
//
// Solidity: function pay(address coinbase, address token, uint256 value) returns()
func (_IPayer *IPayerTransactorSession) Pay(coinbase common.Address, token common.Address, value *big.Int) (*types.Transaction, error) {
	return _IPayer.Contract.Pay(&_IPayer.TransactOpts, coinbase, token, value)
}

// Payment is a paid mutator transaction binding the contract method 0xdd50295e.
//
// Solidity: function payment(address to, uint256 txGas, bytes4 txMethodSig, bytes32[] txMethodParams) returns(address token, uint256 price, uint256 payGasLimit)
func (_IPayer *IPayerTransactor) Payment(opts *bind.TransactOpts, to common.Address, txGas *big.Int, txMethodSig [4]byte, txMethodParams [][32]byte) (*types.Transaction, error) {
	return _IPayer.contract.Transact(opts, "payment", to, txGas, txMethodSig, txMethodParams)
}

// Payment is a paid mutator transaction binding the contract method 0xdd50295e.
//
// Solidity: function payment(address to, uint256 txGas, bytes4 txMethodSig, bytes32[] txMethodParams) returns(address token, uint256 price, uint256 payGasLimit)
func (_IPayer *IPayerSession) Payment(to common.Address, txGas *big.Int, txMethodSig [4]byte, txMethodParams [][32]byte) (*types.Transaction, error) {
	return _IPayer.Contract.Payment(&_IPayer.TransactOpts, to, txGas, txMethodSig, txMethodParams)
}

// Payment is a paid mutator transaction binding the contract method 0xdd50295e.
//
// Solidity: function payment(address to, uint256 txGas, bytes4 txMethodSig, bytes32[] txMethodParams) returns(address token, uint256 price, uint256 payGasLimit)
func (_IPayer *IPayerTransactorSession) Payment(to common.Address, txGas *big.Int, txMethodSig [4]byte, txMethodParams [][32]byte) (*types.Transaction, error) {
	return _IPayer.Contract.Payment(&_IPayer.TransactOpts, to, txGas, txMethodSig, txMethodParams)
}

// TokenPayerABI is the input ABI used to generate the binding from.
const TokenPayerABI = "[{\"inputs\":[],\"name\":\"ERC20PayGasLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FeeTokenFallbackPath\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FeeTokenPath\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TokenPricePath\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TokenPricesContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"pay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"txGas\",\"type\":\"uint256\"},{\"internalType\":\"bytes4\",\"name\":\"txMethodSig\",\"type\":\"bytes4\"},{\"internalType\":\"bytes32[]\",\"name\":\"txMethodParams\",\"type\":\"bytes32[]\"}],\"name\":\"payment\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"payGasLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TokenPayerFuncSigs maps the 4-byte function signature to its string representation.
var TokenPayerFuncSigs = map[string]string{
	"388935d6": "ERC20PayGasLimit()",
	"5e8f4d74": "FeeTokenFallbackPath()",
	"c8c1956b": "FeeTokenPath()",
	"c9969bcf": "TokenPricePath()",
	"5e8c2e09": "TokenPricesContract()",
	"b3d76188": "pay(address,address,uint256)",
	"dd50295e": "payment(address,uint256,bytes4,bytes32[])",
}

// TokenPayerBin is the compiled bytecode used for deploying new contracts.
var TokenPayerBin = "0x608060405234801561001057600080fd5b5061045f806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c8063b3d761881161005b578063b3d76188146100c8578063c8c1956b14610100578063c9969bcf14610108578063dd50295e146101105761007d565b8063388935d6146100825780635e8c2e091461009c5780635e8f4d74146100c0575b600080fd5b61008a6101ce565b60408051918252519081900360200190f35b6100a46101d4565b604080516001600160a01b039092168252519081900360200190f35b61008a6101db565b6100fe600480360360608110156100de57600080fd5b506001600160a01b038135811691602081013590911690604001356101f2565b005b61008a610283565b61008a610292565b6101a66004803603608081101561012657600080fd5b6001600160a01b03823516916020810135916001600160e01b0319604083013516919081019060808101606082013564010000000081111561016757600080fd5b82018360208201111561017957600080fd5b8035906020019184602083028401116401000000008311171561019b57600080fd5b5090925090506102a4565b604080516001600160a01b039094168452602084019290925282820152519081900360600190f35b6184d081565b6206789a81565b6f466565546f6b656e46616c6c6261636b60801b81565b816001600160a01b031663a9059cbb84836040518363ffffffff1660e01b815260040180836001600160a01b03166001600160a01b0316815260200182815260200192505050602060405180830381600087803b15801561025257600080fd5b505af1158015610266573d6000803e3d6000fd5b505050506040513d602081101561027c57600080fd5b5050505050565b672332b2aa37b5b2b760c11b81565b6a546f6b656e50726963652d60a81b81565b60008060006102bd672332b2aa37b5b2b760c11b610361565b92506001600160a01b038316156102e8576102d783610365565b915081156102e857506184d0610356565b6102f188610365565b9150811561030557508691506184d0610356565b6103216f466565546f6b656e46616c6c6261636b60801b610361565b92506001600160a01b0383161561034c5761033b83610365565b9150811561034c57506184d0610356565b5060009150819050805b955095509592505050565b5490565b6000806103886103836a546f6b656e50726963652d60a81b85610415565b610361565b90508015610397579050610410565b604080516341976e0960e01b81526001600160a01b038516600482015290516206789a916341976e09916024808301926020929190829003018186803b1580156103e057600080fd5b505afa1580156103f4573d6000803e3d6000fd5b505050506040513d602081101561040a57600080fd5b50519150505b919050565b6000828152600b82905280601f539291505056fea26469706673582212202cc734ec6bb14805c69c6add5ac26e22b188ba11974b5fab2fb169a180f1b4e464736f6c63430006040033"

// DeployTokenPayer deploys a new Ethereum contract, binding an instance of TokenPayer to it.
func DeployTokenPayer(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TokenPayer, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenPayerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenPayerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenPayer{TokenPayerCaller: TokenPayerCaller{contract: contract}, TokenPayerTransactor: TokenPayerTransactor{contract: contract}, TokenPayerFilterer: TokenPayerFilterer{contract: contract}}, nil
}

// TokenPayer is an auto generated Go binding around an Ethereum contract.
type TokenPayer struct {
	TokenPayerCaller     // Read-only binding to the contract
	TokenPayerTransactor // Write-only binding to the contract
	TokenPayerFilterer   // Log filterer for contract events
}

// TokenPayerCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenPayerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenPayerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenPayerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenPayerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenPayerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenPayerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenPayerSession struct {
	Contract     *TokenPayer       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenPayerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenPayerCallerSession struct {
	Contract *TokenPayerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TokenPayerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenPayerTransactorSession struct {
	Contract     *TokenPayerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TokenPayerRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenPayerRaw struct {
	Contract *TokenPayer // Generic contract binding to access the raw methods on
}

// TokenPayerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenPayerCallerRaw struct {
	Contract *TokenPayerCaller // Generic read-only contract binding to access the raw methods on
}

// TokenPayerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenPayerTransactorRaw struct {
	Contract *TokenPayerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenPayer creates a new instance of TokenPayer, bound to a specific deployed contract.
func NewTokenPayer(address common.Address, backend bind.ContractBackend) (*TokenPayer, error) {
	contract, err := bindTokenPayer(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenPayer{TokenPayerCaller: TokenPayerCaller{contract: contract}, TokenPayerTransactor: TokenPayerTransactor{contract: contract}, TokenPayerFilterer: TokenPayerFilterer{contract: contract}}, nil
}

// NewTokenPayerCaller creates a new read-only instance of TokenPayer, bound to a specific deployed contract.
func NewTokenPayerCaller(address common.Address, caller bind.ContractCaller) (*TokenPayerCaller, error) {
	contract, err := bindTokenPayer(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenPayerCaller{contract: contract}, nil
}

// NewTokenPayerTransactor creates a new write-only instance of TokenPayer, bound to a specific deployed contract.
func NewTokenPayerTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenPayerTransactor, error) {
	contract, err := bindTokenPayer(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenPayerTransactor{contract: contract}, nil
}

// NewTokenPayerFilterer creates a new log filterer instance of TokenPayer, bound to a specific deployed contract.
func NewTokenPayerFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenPayerFilterer, error) {
	contract, err := bindTokenPayer(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenPayerFilterer{contract: contract}, nil
}

// bindTokenPayer binds a generic wrapper to an already deployed contract.
func bindTokenPayer(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenPayerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenPayer *TokenPayerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenPayer.Contract.TokenPayerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenPayer *TokenPayerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPayer.Contract.TokenPayerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenPayer *TokenPayerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenPayer.Contract.TokenPayerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenPayer *TokenPayerCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenPayer.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenPayer *TokenPayerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPayer.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenPayer *TokenPayerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenPayer.Contract.contract.Transact(opts, method, params...)
}

// ERC20PayGasLimit is a paid mutator transaction binding the contract method 0x388935d6.
//
// Solidity: function ERC20PayGasLimit() returns(uint256)
func (_TokenPayer *TokenPayerTransactor) ERC20PayGasLimit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPayer.contract.Transact(opts, "ERC20PayGasLimit")
}

// ERC20PayGasLimit is a paid mutator transaction binding the contract method 0x388935d6.
//
// Solidity: function ERC20PayGasLimit() returns(uint256)
func (_TokenPayer *TokenPayerSession) ERC20PayGasLimit() (*types.Transaction, error) {
	return _TokenPayer.Contract.ERC20PayGasLimit(&_TokenPayer.TransactOpts)
}

// ERC20PayGasLimit is a paid mutator transaction binding the contract method 0x388935d6.
//
// Solidity: function ERC20PayGasLimit() returns(uint256)
func (_TokenPayer *TokenPayerTransactorSession) ERC20PayGasLimit() (*types.Transaction, error) {
	return _TokenPayer.Contract.ERC20PayGasLimit(&_TokenPayer.TransactOpts)
}

// FeeTokenFallbackPath is a paid mutator transaction binding the contract method 0x5e8f4d74.
//
// Solidity: function FeeTokenFallbackPath() returns(bytes32)
func (_TokenPayer *TokenPayerTransactor) FeeTokenFallbackPath(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPayer.contract.Transact(opts, "FeeTokenFallbackPath")
}

// FeeTokenFallbackPath is a paid mutator transaction binding the contract method 0x5e8f4d74.
//
// Solidity: function FeeTokenFallbackPath() returns(bytes32)
func (_TokenPayer *TokenPayerSession) FeeTokenFallbackPath() (*types.Transaction, error) {
	return _TokenPayer.Contract.FeeTokenFallbackPath(&_TokenPayer.TransactOpts)
}

// FeeTokenFallbackPath is a paid mutator transaction binding the contract method 0x5e8f4d74.
//
// Solidity: function FeeTokenFallbackPath() returns(bytes32)
func (_TokenPayer *TokenPayerTransactorSession) FeeTokenFallbackPath() (*types.Transaction, error) {
	return _TokenPayer.Contract.FeeTokenFallbackPath(&_TokenPayer.TransactOpts)
}

// FeeTokenPath is a paid mutator transaction binding the contract method 0xc8c1956b.
//
// Solidity: function FeeTokenPath() returns(bytes32)
func (_TokenPayer *TokenPayerTransactor) FeeTokenPath(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPayer.contract.Transact(opts, "FeeTokenPath")
}

// FeeTokenPath is a paid mutator transaction binding the contract method 0xc8c1956b.
//
// Solidity: function FeeTokenPath() returns(bytes32)
func (_TokenPayer *TokenPayerSession) FeeTokenPath() (*types.Transaction, error) {
	return _TokenPayer.Contract.FeeTokenPath(&_TokenPayer.TransactOpts)
}

// FeeTokenPath is a paid mutator transaction binding the contract method 0xc8c1956b.
//
// Solidity: function FeeTokenPath() returns(bytes32)
func (_TokenPayer *TokenPayerTransactorSession) FeeTokenPath() (*types.Transaction, error) {
	return _TokenPayer.Contract.FeeTokenPath(&_TokenPayer.TransactOpts)
}

// TokenPricePath is a paid mutator transaction binding the contract method 0xc9969bcf.
//
// Solidity: function TokenPricePath() returns(bytes32)
func (_TokenPayer *TokenPayerTransactor) TokenPricePath(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPayer.contract.Transact(opts, "TokenPricePath")
}

// TokenPricePath is a paid mutator transaction binding the contract method 0xc9969bcf.
//
// Solidity: function TokenPricePath() returns(bytes32)
func (_TokenPayer *TokenPayerSession) TokenPricePath() (*types.Transaction, error) {
	return _TokenPayer.Contract.TokenPricePath(&_TokenPayer.TransactOpts)
}

// TokenPricePath is a paid mutator transaction binding the contract method 0xc9969bcf.
//
// Solidity: function TokenPricePath() returns(bytes32)
func (_TokenPayer *TokenPayerTransactorSession) TokenPricePath() (*types.Transaction, error) {
	return _TokenPayer.Contract.TokenPricePath(&_TokenPayer.TransactOpts)
}

// TokenPricesContract is a paid mutator transaction binding the contract method 0x5e8c2e09.
//
// Solidity: function TokenPricesContract() returns(address)
func (_TokenPayer *TokenPayerTransactor) TokenPricesContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPayer.contract.Transact(opts, "TokenPricesContract")
}

// TokenPricesContract is a paid mutator transaction binding the contract method 0x5e8c2e09.
//
// Solidity: function TokenPricesContract() returns(address)
func (_TokenPayer *TokenPayerSession) TokenPricesContract() (*types.Transaction, error) {
	return _TokenPayer.Contract.TokenPricesContract(&_TokenPayer.TransactOpts)
}

// TokenPricesContract is a paid mutator transaction binding the contract method 0x5e8c2e09.
//
// Solidity: function TokenPricesContract() returns(address)
func (_TokenPayer *TokenPayerTransactorSession) TokenPricesContract() (*types.Transaction, error) {
	return _TokenPayer.Contract.TokenPricesContract(&_TokenPayer.TransactOpts)
}

// Pay is a paid mutator transaction binding the contract method 0xb3d76188.
//
// Solidity: function pay(address coinbase, address token, uint256 value) returns()
func (_TokenPayer *TokenPayerTransactor) Pay(opts *bind.TransactOpts, coinbase common.Address, token common.Address, value *big.Int) (*types.Transaction, error) {
	return _TokenPayer.contract.Transact(opts, "pay", coinbase, token, value)
}

// Pay is a paid mutator transaction binding the contract method 0xb3d76188.
//
// Solidity: function pay(address coinbase, address token, uint256 value) returns()
func (_TokenPayer *TokenPayerSession) Pay(coinbase common.Address, token common.Address, value *big.Int) (*types.Transaction, error) {
	return _TokenPayer.Contract.Pay(&_TokenPayer.TransactOpts, coinbase, token, value)
}

// Pay is a paid mutator transaction binding the contract method 0xb3d76188.
//
// Solidity: function pay(address coinbase, address token, uint256 value) returns()
func (_TokenPayer *TokenPayerTransactorSession) Pay(coinbase common.Address, token common.Address, value *big.Int) (*types.Transaction, error) {
	return _TokenPayer.Contract.Pay(&_TokenPayer.TransactOpts, coinbase, token, value)
}

// Payment is a paid mutator transaction binding the contract method 0xdd50295e.
//
// Solidity: function payment(address to, uint256 txGas, bytes4 txMethodSig, bytes32[] txMethodParams) returns(address token, uint256 price, uint256 payGasLimit)
func (_TokenPayer *TokenPayerTransactor) Payment(opts *bind.TransactOpts, to common.Address, txGas *big.Int, txMethodSig [4]byte, txMethodParams [][32]byte) (*types.Transaction, error) {
	return _TokenPayer.contract.Transact(opts, "payment", to, txGas, txMethodSig, txMethodParams)
}

// Payment is a paid mutator transaction binding the contract method 0xdd50295e.
//
// Solidity: function payment(address to, uint256 txGas, bytes4 txMethodSig, bytes32[] txMethodParams) returns(address token, uint256 price, uint256 payGasLimit)
func (_TokenPayer *TokenPayerSession) Payment(to common.Address, txGas *big.Int, txMethodSig [4]byte, txMethodParams [][32]byte) (*types.Transaction, error) {
	return _TokenPayer.Contract.Payment(&_TokenPayer.TransactOpts, to, txGas, txMethodSig, txMethodParams)
}

// Payment is a paid mutator transaction binding the contract method 0xdd50295e.
//
// Solidity: function payment(address to, uint256 txGas, bytes4 txMethodSig, bytes32[] txMethodParams) returns(address token, uint256 price, uint256 payGasLimit)
func (_TokenPayer *TokenPayerTransactorSession) Payment(to common.Address, txGas *big.Int, txMethodSig [4]byte, txMethodParams [][32]byte) (*types.Transaction, error) {
	return _TokenPayer.Contract.Payment(&_TokenPayer.TransactOpts, to, txGas, txMethodSig, txMethodParams)
}

// TokenPriceABI is the input ABI used to generate the binding from.
const TokenPriceABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"team\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"prices\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"TRUE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TeamMemberPath\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TokenPricePath\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"setPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// TokenPriceFuncSigs maps the 4-byte function signature to its string representation.
var TokenPriceFuncSigs = map[string]string{
	"5994d984": "TRUE()",
	"654b8c2a": "TeamMemberPath()",
	"c9969bcf": "TokenPricePath()",
	"41976e09": "getPrice(address)",
	"00e4768b": "setPrice(address,uint256)",
}

// TokenPriceBin is the compiled bytecode used for deploying new contracts.
var TokenPriceBin = "0x608060405234801561001057600080fd5b5060405161051e38038061051e8339818101604052606081101561003357600080fd5b810190808051604051939291908464010000000082111561005357600080fd5b90830190602082018581111561006857600080fd5b825186602082028301116401000000008211171561008557600080fd5b82525081516020918201928201910280838360005b838110156100b257818101518382015260200161009a565b50505050905001604052602001805160405193929190846401000000008211156100db57600080fd5b9083019060208201858111156100f057600080fd5b825186602082028301116401000000008211171561010d57600080fd5b82525081516020918201928201910280838360005b8381101561013a578181015183820152602001610122565b505050509050016040526020018051604051939291908464010000000082111561016357600080fd5b90830190602082018581111561017857600080fd5b825186602082028301116401000000008211171561019557600080fd5b82525081516020918201928201910280838360005b838110156101c25781810151838201526020016101aa565b5050505090500160405250505060008090505b8351811015610208576102008482815181106101ed57fe5b602002602001015161025c60201b60201c565b6001016101d5565b5060005b82518110156102535761024b83828151811061022457fe5b602002602001015183838151811061023857fe5b602002602001015161029960201b60201c565b60010161020c565b505050506102ed565b6102966102816a5465616d4d656d6265722d60a81b836102d560201b6101d01760201c565b60001960001b6102e960201b6101e81760201c565b50565b6102d16102be6a546f6b656e50726963652d60a81b846102d560201b6101d01760201c565b8260001b6102e960201b6101e81760201c565b5050565b6000828152600b82905280601f5392915050565b9055565b610222806102fc6000396000f3fe608060405234801561001057600080fd5b50600436106100565760003560e01c8062e4768b1461005b57806341976e09146100895780635994d984146100c1578063654b8c2a146100c9578063c9969bcf146100d1575b600080fd5b6100876004803603604081101561007157600080fd5b506001600160a01b0381351690602001356100d9565b005b6100af6004803603602081101561009f57600080fd5b50356001600160a01b0316610138565b60408051918252519081900360200190f35b6100af610160565b6100af610166565b6100af610178565b6100e23361018a565b61012a576040805162461bcd60e51b8152602060048201526014602482015273666f72207465616d206d656d626572206f6e6c7960601b604482015290519081900360640190fd5b61013482826101af565b5050565b600061015a6101556a546f6b656e50726963652d60a81b846101d0565b6101e4565b92915050565b60001981565b6a5465616d4d656d6265722d60a81b81565b6a546f6b656e50726963652d60a81b81565b60006101a76101556a5465616d4d656d6265722d60a81b846101d0565b151592915050565b6101346101ca6a546f6b656e50726963652d60a81b846101d0565b826101e8565b6000828152600b82905280601f5392915050565b5490565b905556fea2646970667358221220bbe4e3af55d84fab59ac2b3200ca3d8250c3b08d6d04f212bf0e58bde27e619864736f6c63430006040033"

// DeployTokenPrice deploys a new Ethereum contract, binding an instance of TokenPrice to it.
func DeployTokenPrice(auth *bind.TransactOpts, backend bind.ContractBackend, team []common.Address, tokens []common.Address, prices []*big.Int) (common.Address, *types.Transaction, *TokenPrice, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenPriceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenPriceBin), backend, team, tokens, prices)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenPrice{TokenPriceCaller: TokenPriceCaller{contract: contract}, TokenPriceTransactor: TokenPriceTransactor{contract: contract}, TokenPriceFilterer: TokenPriceFilterer{contract: contract}}, nil
}

// TokenPrice is an auto generated Go binding around an Ethereum contract.
type TokenPrice struct {
	TokenPriceCaller     // Read-only binding to the contract
	TokenPriceTransactor // Write-only binding to the contract
	TokenPriceFilterer   // Log filterer for contract events
}

// TokenPriceCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenPriceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenPriceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenPriceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenPriceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenPriceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenPriceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenPriceSession struct {
	Contract     *TokenPrice       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenPriceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenPriceCallerSession struct {
	Contract *TokenPriceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TokenPriceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenPriceTransactorSession struct {
	Contract     *TokenPriceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TokenPriceRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenPriceRaw struct {
	Contract *TokenPrice // Generic contract binding to access the raw methods on
}

// TokenPriceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenPriceCallerRaw struct {
	Contract *TokenPriceCaller // Generic read-only contract binding to access the raw methods on
}

// TokenPriceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenPriceTransactorRaw struct {
	Contract *TokenPriceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenPrice creates a new instance of TokenPrice, bound to a specific deployed contract.
func NewTokenPrice(address common.Address, backend bind.ContractBackend) (*TokenPrice, error) {
	contract, err := bindTokenPrice(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenPrice{TokenPriceCaller: TokenPriceCaller{contract: contract}, TokenPriceTransactor: TokenPriceTransactor{contract: contract}, TokenPriceFilterer: TokenPriceFilterer{contract: contract}}, nil
}

// NewTokenPriceCaller creates a new read-only instance of TokenPrice, bound to a specific deployed contract.
func NewTokenPriceCaller(address common.Address, caller bind.ContractCaller) (*TokenPriceCaller, error) {
	contract, err := bindTokenPrice(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenPriceCaller{contract: contract}, nil
}

// NewTokenPriceTransactor creates a new write-only instance of TokenPrice, bound to a specific deployed contract.
func NewTokenPriceTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenPriceTransactor, error) {
	contract, err := bindTokenPrice(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenPriceTransactor{contract: contract}, nil
}

// NewTokenPriceFilterer creates a new log filterer instance of TokenPrice, bound to a specific deployed contract.
func NewTokenPriceFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenPriceFilterer, error) {
	contract, err := bindTokenPrice(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenPriceFilterer{contract: contract}, nil
}

// bindTokenPrice binds a generic wrapper to an already deployed contract.
func bindTokenPrice(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenPriceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenPrice *TokenPriceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenPrice.Contract.TokenPriceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenPrice *TokenPriceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPrice.Contract.TokenPriceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenPrice *TokenPriceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenPrice.Contract.TokenPriceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenPrice *TokenPriceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _TokenPrice.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenPrice *TokenPriceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPrice.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenPrice *TokenPriceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenPrice.Contract.contract.Transact(opts, method, params...)
}

// TRUE is a paid mutator transaction binding the contract method 0x5994d984.
//
// Solidity: function TRUE() returns(bytes32)
func (_TokenPrice *TokenPriceTransactor) TRUE(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPrice.contract.Transact(opts, "TRUE")
}

// TRUE is a paid mutator transaction binding the contract method 0x5994d984.
//
// Solidity: function TRUE() returns(bytes32)
func (_TokenPrice *TokenPriceSession) TRUE() (*types.Transaction, error) {
	return _TokenPrice.Contract.TRUE(&_TokenPrice.TransactOpts)
}

// TRUE is a paid mutator transaction binding the contract method 0x5994d984.
//
// Solidity: function TRUE() returns(bytes32)
func (_TokenPrice *TokenPriceTransactorSession) TRUE() (*types.Transaction, error) {
	return _TokenPrice.Contract.TRUE(&_TokenPrice.TransactOpts)
}

// TeamMemberPath is a paid mutator transaction binding the contract method 0x654b8c2a.
//
// Solidity: function TeamMemberPath() returns(bytes32)
func (_TokenPrice *TokenPriceTransactor) TeamMemberPath(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPrice.contract.Transact(opts, "TeamMemberPath")
}

// TeamMemberPath is a paid mutator transaction binding the contract method 0x654b8c2a.
//
// Solidity: function TeamMemberPath() returns(bytes32)
func (_TokenPrice *TokenPriceSession) TeamMemberPath() (*types.Transaction, error) {
	return _TokenPrice.Contract.TeamMemberPath(&_TokenPrice.TransactOpts)
}

// TeamMemberPath is a paid mutator transaction binding the contract method 0x654b8c2a.
//
// Solidity: function TeamMemberPath() returns(bytes32)
func (_TokenPrice *TokenPriceTransactorSession) TeamMemberPath() (*types.Transaction, error) {
	return _TokenPrice.Contract.TeamMemberPath(&_TokenPrice.TransactOpts)
}

// TokenPricePath is a paid mutator transaction binding the contract method 0xc9969bcf.
//
// Solidity: function TokenPricePath() returns(bytes32)
func (_TokenPrice *TokenPriceTransactor) TokenPricePath(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPrice.contract.Transact(opts, "TokenPricePath")
}

// TokenPricePath is a paid mutator transaction binding the contract method 0xc9969bcf.
//
// Solidity: function TokenPricePath() returns(bytes32)
func (_TokenPrice *TokenPriceSession) TokenPricePath() (*types.Transaction, error) {
	return _TokenPrice.Contract.TokenPricePath(&_TokenPrice.TransactOpts)
}

// TokenPricePath is a paid mutator transaction binding the contract method 0xc9969bcf.
//
// Solidity: function TokenPricePath() returns(bytes32)
func (_TokenPrice *TokenPriceTransactorSession) TokenPricePath() (*types.Transaction, error) {
	return _TokenPrice.Contract.TokenPricePath(&_TokenPrice.TransactOpts)
}

// GetPrice is a paid mutator transaction binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) returns(uint256 price)
func (_TokenPrice *TokenPriceTransactor) GetPrice(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _TokenPrice.contract.Transact(opts, "getPrice", token)
}

// GetPrice is a paid mutator transaction binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) returns(uint256 price)
func (_TokenPrice *TokenPriceSession) GetPrice(token common.Address) (*types.Transaction, error) {
	return _TokenPrice.Contract.GetPrice(&_TokenPrice.TransactOpts, token)
}

// GetPrice is a paid mutator transaction binding the contract method 0x41976e09.
//
// Solidity: function getPrice(address token) returns(uint256 price)
func (_TokenPrice *TokenPriceTransactorSession) GetPrice(token common.Address) (*types.Transaction, error) {
	return _TokenPrice.Contract.GetPrice(&_TokenPrice.TransactOpts, token)
}

// SetPrice is a paid mutator transaction binding the contract method 0x00e4768b.
//
// Solidity: function setPrice(address token, uint256 price) returns()
func (_TokenPrice *TokenPriceTransactor) SetPrice(opts *bind.TransactOpts, token common.Address, price *big.Int) (*types.Transaction, error) {
	return _TokenPrice.contract.Transact(opts, "setPrice", token, price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x00e4768b.
//
// Solidity: function setPrice(address token, uint256 price) returns()
func (_TokenPrice *TokenPriceSession) SetPrice(token common.Address, price *big.Int) (*types.Transaction, error) {
	return _TokenPrice.Contract.SetPrice(&_TokenPrice.TransactOpts, token, price)
}

// SetPrice is a paid mutator transaction binding the contract method 0x00e4768b.
//
// Solidity: function setPrice(address token, uint256 price) returns()
func (_TokenPrice *TokenPriceTransactorSession) SetPrice(token common.Address, price *big.Int) (*types.Transaction, error) {
	return _TokenPrice.Contract.SetPrice(&_TokenPrice.TransactOpts, token, price)
}

// DsABI is the input ABI used to generate the binding from.
const DsABI = "[]"

// DsBin is the compiled bytecode used for deploying new contracts.
var DsBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea264697066735822122029136e0c7497f95a20c6079b34f36f2cba9e2ad98f2b4210505fa1965368037664736f6c63430006040033"

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
