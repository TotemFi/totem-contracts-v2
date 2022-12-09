// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// OperatorMetaData contains all meta data concerning the Operator contract.
var OperatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OperatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OperatorRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// OperatorABI is the input ABI used to generate the binding from.
// Deprecated: Use OperatorMetaData.ABI instead.
var OperatorABI = OperatorMetaData.ABI

// Operator is an auto generated Go binding around an Ethereum contract.
type Operator struct {
	OperatorCaller     // Read-only binding to the contract
	OperatorTransactor // Write-only binding to the contract
	OperatorFilterer   // Log filterer for contract events
}

// OperatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type OperatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OperatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OperatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OperatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OperatorSession struct {
	Contract     *Operator         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OperatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OperatorCallerSession struct {
	Contract *OperatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// OperatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OperatorTransactorSession struct {
	Contract     *OperatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// OperatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type OperatorRaw struct {
	Contract *Operator // Generic contract binding to access the raw methods on
}

// OperatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OperatorCallerRaw struct {
	Contract *OperatorCaller // Generic read-only contract binding to access the raw methods on
}

// OperatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OperatorTransactorRaw struct {
	Contract *OperatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOperator creates a new instance of Operator, bound to a specific deployed contract.
func NewOperator(address common.Address, backend bind.ContractBackend) (*Operator, error) {
	contract, err := bindOperator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Operator{OperatorCaller: OperatorCaller{contract: contract}, OperatorTransactor: OperatorTransactor{contract: contract}, OperatorFilterer: OperatorFilterer{contract: contract}}, nil
}

// NewOperatorCaller creates a new read-only instance of Operator, bound to a specific deployed contract.
func NewOperatorCaller(address common.Address, caller bind.ContractCaller) (*OperatorCaller, error) {
	contract, err := bindOperator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OperatorCaller{contract: contract}, nil
}

// NewOperatorTransactor creates a new write-only instance of Operator, bound to a specific deployed contract.
func NewOperatorTransactor(address common.Address, transactor bind.ContractTransactor) (*OperatorTransactor, error) {
	contract, err := bindOperator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OperatorTransactor{contract: contract}, nil
}

// NewOperatorFilterer creates a new log filterer instance of Operator, bound to a specific deployed contract.
func NewOperatorFilterer(address common.Address, filterer bind.ContractFilterer) (*OperatorFilterer, error) {
	contract, err := bindOperator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OperatorFilterer{contract: contract}, nil
}

// bindOperator binds a generic wrapper to an already deployed contract.
func bindOperator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OperatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Operator *OperatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Operator.Contract.OperatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Operator *OperatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Operator.Contract.OperatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Operator *OperatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Operator.Contract.OperatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Operator *OperatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Operator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Operator *OperatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Operator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Operator *OperatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Operator.Contract.contract.Transact(opts, method, params...)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address account) view returns(bool)
func (_Operator *OperatorCaller) IsOperator(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Operator.contract.Call(opts, &out, "isOperator", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address account) view returns(bool)
func (_Operator *OperatorSession) IsOperator(account common.Address) (bool, error) {
	return _Operator.Contract.IsOperator(&_Operator.CallOpts, account)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address account) view returns(bool)
func (_Operator *OperatorCallerSession) IsOperator(account common.Address) (bool, error) {
	return _Operator.Contract.IsOperator(&_Operator.CallOpts, account)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address account) returns()
func (_Operator *OperatorTransactor) AddOperator(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "addOperator", account)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address account) returns()
func (_Operator *OperatorSession) AddOperator(account common.Address) (*types.Transaction, error) {
	return _Operator.Contract.AddOperator(&_Operator.TransactOpts, account)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address account) returns()
func (_Operator *OperatorTransactorSession) AddOperator(account common.Address) (*types.Transaction, error) {
	return _Operator.Contract.AddOperator(&_Operator.TransactOpts, account)
}

// RenounceOperator is a paid mutator transaction binding the contract method 0x2ab6f8db.
//
// Solidity: function renounceOperator() returns()
func (_Operator *OperatorTransactor) RenounceOperator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Operator.contract.Transact(opts, "renounceOperator")
}

// RenounceOperator is a paid mutator transaction binding the contract method 0x2ab6f8db.
//
// Solidity: function renounceOperator() returns()
func (_Operator *OperatorSession) RenounceOperator() (*types.Transaction, error) {
	return _Operator.Contract.RenounceOperator(&_Operator.TransactOpts)
}

// RenounceOperator is a paid mutator transaction binding the contract method 0x2ab6f8db.
//
// Solidity: function renounceOperator() returns()
func (_Operator *OperatorTransactorSession) RenounceOperator() (*types.Transaction, error) {
	return _Operator.Contract.RenounceOperator(&_Operator.TransactOpts)
}

// OperatorOperatorAddedIterator is returned from FilterOperatorAdded and is used to iterate over the raw logs and unpacked data for OperatorAdded events raised by the Operator contract.
type OperatorOperatorAddedIterator struct {
	Event *OperatorOperatorAdded // Event containing the contract specifics and raw log

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
func (it *OperatorOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorOperatorAdded)
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
		it.Event = new(OperatorOperatorAdded)
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
func (it *OperatorOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorOperatorAdded represents a OperatorAdded event raised by the Operator contract.
type OperatorOperatorAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOperatorAdded is a free log retrieval operation binding the contract event 0xac6fa858e9350a46cec16539926e0fde25b7629f84b5a72bffaae4df888ae86d.
//
// Solidity: event OperatorAdded(address indexed account)
func (_Operator *OperatorFilterer) FilterOperatorAdded(opts *bind.FilterOpts, account []common.Address) (*OperatorOperatorAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Operator.contract.FilterLogs(opts, "OperatorAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &OperatorOperatorAddedIterator{contract: _Operator.contract, event: "OperatorAdded", logs: logs, sub: sub}, nil
}

// WatchOperatorAdded is a free log subscription operation binding the contract event 0xac6fa858e9350a46cec16539926e0fde25b7629f84b5a72bffaae4df888ae86d.
//
// Solidity: event OperatorAdded(address indexed account)
func (_Operator *OperatorFilterer) WatchOperatorAdded(opts *bind.WatchOpts, sink chan<- *OperatorOperatorAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Operator.contract.WatchLogs(opts, "OperatorAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorOperatorAdded)
				if err := _Operator.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
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

// ParseOperatorAdded is a log parse operation binding the contract event 0xac6fa858e9350a46cec16539926e0fde25b7629f84b5a72bffaae4df888ae86d.
//
// Solidity: event OperatorAdded(address indexed account)
func (_Operator *OperatorFilterer) ParseOperatorAdded(log types.Log) (*OperatorOperatorAdded, error) {
	event := new(OperatorOperatorAdded)
	if err := _Operator.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OperatorOperatorRemovedIterator is returned from FilterOperatorRemoved and is used to iterate over the raw logs and unpacked data for OperatorRemoved events raised by the Operator contract.
type OperatorOperatorRemovedIterator struct {
	Event *OperatorOperatorRemoved // Event containing the contract specifics and raw log

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
func (it *OperatorOperatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OperatorOperatorRemoved)
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
		it.Event = new(OperatorOperatorRemoved)
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
func (it *OperatorOperatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OperatorOperatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OperatorOperatorRemoved represents a OperatorRemoved event raised by the Operator contract.
type OperatorOperatorRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemoved is a free log retrieval operation binding the contract event 0x80c0b871b97b595b16a7741c1b06fed0c6f6f558639f18ccbce50724325dc40d.
//
// Solidity: event OperatorRemoved(address indexed account)
func (_Operator *OperatorFilterer) FilterOperatorRemoved(opts *bind.FilterOpts, account []common.Address) (*OperatorOperatorRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Operator.contract.FilterLogs(opts, "OperatorRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &OperatorOperatorRemovedIterator{contract: _Operator.contract, event: "OperatorRemoved", logs: logs, sub: sub}, nil
}

// WatchOperatorRemoved is a free log subscription operation binding the contract event 0x80c0b871b97b595b16a7741c1b06fed0c6f6f558639f18ccbce50724325dc40d.
//
// Solidity: event OperatorRemoved(address indexed account)
func (_Operator *OperatorFilterer) WatchOperatorRemoved(opts *bind.WatchOpts, sink chan<- *OperatorOperatorRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Operator.contract.WatchLogs(opts, "OperatorRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OperatorOperatorRemoved)
				if err := _Operator.contract.UnpackLog(event, "OperatorRemoved", log); err != nil {
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

// ParseOperatorRemoved is a log parse operation binding the contract event 0x80c0b871b97b595b16a7741c1b06fed0c6f6f558639f18ccbce50724325dc40d.
//
// Solidity: event OperatorRemoved(address indexed account)
func (_Operator *OperatorFilterer) ParseOperatorRemoved(log types.Log) (*OperatorOperatorRemoved, error) {
	event := new(OperatorOperatorRemoved)
	if err := _Operator.contract.UnpackLog(event, "OperatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
