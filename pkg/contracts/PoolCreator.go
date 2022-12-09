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

// PoolCreatorMetaData contains all meta data concerning the PoolCreator contract.
var PoolCreatorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PoolCreatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"PoolCreatorRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addPoolCreator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isPoolCreator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renouncePoolCreator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PoolCreatorABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolCreatorMetaData.ABI instead.
var PoolCreatorABI = PoolCreatorMetaData.ABI

// PoolCreator is an auto generated Go binding around an Ethereum contract.
type PoolCreator struct {
	PoolCreatorCaller     // Read-only binding to the contract
	PoolCreatorTransactor // Write-only binding to the contract
	PoolCreatorFilterer   // Log filterer for contract events
}

// PoolCreatorCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolCreatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolCreatorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolCreatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolCreatorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolCreatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolCreatorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolCreatorSession struct {
	Contract     *PoolCreator      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolCreatorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolCreatorCallerSession struct {
	Contract *PoolCreatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PoolCreatorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolCreatorTransactorSession struct {
	Contract     *PoolCreatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PoolCreatorRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolCreatorRaw struct {
	Contract *PoolCreator // Generic contract binding to access the raw methods on
}

// PoolCreatorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolCreatorCallerRaw struct {
	Contract *PoolCreatorCaller // Generic read-only contract binding to access the raw methods on
}

// PoolCreatorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolCreatorTransactorRaw struct {
	Contract *PoolCreatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoolCreator creates a new instance of PoolCreator, bound to a specific deployed contract.
func NewPoolCreator(address common.Address, backend bind.ContractBackend) (*PoolCreator, error) {
	contract, err := bindPoolCreator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoolCreator{PoolCreatorCaller: PoolCreatorCaller{contract: contract}, PoolCreatorTransactor: PoolCreatorTransactor{contract: contract}, PoolCreatorFilterer: PoolCreatorFilterer{contract: contract}}, nil
}

// NewPoolCreatorCaller creates a new read-only instance of PoolCreator, bound to a specific deployed contract.
func NewPoolCreatorCaller(address common.Address, caller bind.ContractCaller) (*PoolCreatorCaller, error) {
	contract, err := bindPoolCreator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolCreatorCaller{contract: contract}, nil
}

// NewPoolCreatorTransactor creates a new write-only instance of PoolCreator, bound to a specific deployed contract.
func NewPoolCreatorTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolCreatorTransactor, error) {
	contract, err := bindPoolCreator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolCreatorTransactor{contract: contract}, nil
}

// NewPoolCreatorFilterer creates a new log filterer instance of PoolCreator, bound to a specific deployed contract.
func NewPoolCreatorFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolCreatorFilterer, error) {
	contract, err := bindPoolCreator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolCreatorFilterer{contract: contract}, nil
}

// bindPoolCreator binds a generic wrapper to an already deployed contract.
func bindPoolCreator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PoolCreatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolCreator *PoolCreatorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolCreator.Contract.PoolCreatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolCreator *PoolCreatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolCreator.Contract.PoolCreatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolCreator *PoolCreatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolCreator.Contract.PoolCreatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolCreator *PoolCreatorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolCreator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolCreator *PoolCreatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolCreator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolCreator *PoolCreatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolCreator.Contract.contract.Transact(opts, method, params...)
}

// IsPoolCreator is a free data retrieval call binding the contract method 0x9a94775d.
//
// Solidity: function isPoolCreator(address account) view returns(bool)
func (_PoolCreator *PoolCreatorCaller) IsPoolCreator(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _PoolCreator.contract.Call(opts, &out, "isPoolCreator", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPoolCreator is a free data retrieval call binding the contract method 0x9a94775d.
//
// Solidity: function isPoolCreator(address account) view returns(bool)
func (_PoolCreator *PoolCreatorSession) IsPoolCreator(account common.Address) (bool, error) {
	return _PoolCreator.Contract.IsPoolCreator(&_PoolCreator.CallOpts, account)
}

// IsPoolCreator is a free data retrieval call binding the contract method 0x9a94775d.
//
// Solidity: function isPoolCreator(address account) view returns(bool)
func (_PoolCreator *PoolCreatorCallerSession) IsPoolCreator(account common.Address) (bool, error) {
	return _PoolCreator.Contract.IsPoolCreator(&_PoolCreator.CallOpts, account)
}

// AddPoolCreator is a paid mutator transaction binding the contract method 0x8c2d741c.
//
// Solidity: function addPoolCreator(address account) returns()
func (_PoolCreator *PoolCreatorTransactor) AddPoolCreator(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _PoolCreator.contract.Transact(opts, "addPoolCreator", account)
}

// AddPoolCreator is a paid mutator transaction binding the contract method 0x8c2d741c.
//
// Solidity: function addPoolCreator(address account) returns()
func (_PoolCreator *PoolCreatorSession) AddPoolCreator(account common.Address) (*types.Transaction, error) {
	return _PoolCreator.Contract.AddPoolCreator(&_PoolCreator.TransactOpts, account)
}

// AddPoolCreator is a paid mutator transaction binding the contract method 0x8c2d741c.
//
// Solidity: function addPoolCreator(address account) returns()
func (_PoolCreator *PoolCreatorTransactorSession) AddPoolCreator(account common.Address) (*types.Transaction, error) {
	return _PoolCreator.Contract.AddPoolCreator(&_PoolCreator.TransactOpts, account)
}

// RenouncePoolCreator is a paid mutator transaction binding the contract method 0xe281cc7e.
//
// Solidity: function renouncePoolCreator() returns()
func (_PoolCreator *PoolCreatorTransactor) RenouncePoolCreator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolCreator.contract.Transact(opts, "renouncePoolCreator")
}

// RenouncePoolCreator is a paid mutator transaction binding the contract method 0xe281cc7e.
//
// Solidity: function renouncePoolCreator() returns()
func (_PoolCreator *PoolCreatorSession) RenouncePoolCreator() (*types.Transaction, error) {
	return _PoolCreator.Contract.RenouncePoolCreator(&_PoolCreator.TransactOpts)
}

// RenouncePoolCreator is a paid mutator transaction binding the contract method 0xe281cc7e.
//
// Solidity: function renouncePoolCreator() returns()
func (_PoolCreator *PoolCreatorTransactorSession) RenouncePoolCreator() (*types.Transaction, error) {
	return _PoolCreator.Contract.RenouncePoolCreator(&_PoolCreator.TransactOpts)
}

// PoolCreatorPoolCreatorAddedIterator is returned from FilterPoolCreatorAdded and is used to iterate over the raw logs and unpacked data for PoolCreatorAdded events raised by the PoolCreator contract.
type PoolCreatorPoolCreatorAddedIterator struct {
	Event *PoolCreatorPoolCreatorAdded // Event containing the contract specifics and raw log

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
func (it *PoolCreatorPoolCreatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolCreatorPoolCreatorAdded)
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
		it.Event = new(PoolCreatorPoolCreatorAdded)
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
func (it *PoolCreatorPoolCreatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolCreatorPoolCreatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolCreatorPoolCreatorAdded represents a PoolCreatorAdded event raised by the PoolCreator contract.
type PoolCreatorPoolCreatorAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPoolCreatorAdded is a free log retrieval operation binding the contract event 0xac89bb7b3d0c5a763a97f31bc75f8faee00426e7b235f02ca76897d55caf7b61.
//
// Solidity: event PoolCreatorAdded(address indexed account)
func (_PoolCreator *PoolCreatorFilterer) FilterPoolCreatorAdded(opts *bind.FilterOpts, account []common.Address) (*PoolCreatorPoolCreatorAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PoolCreator.contract.FilterLogs(opts, "PoolCreatorAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &PoolCreatorPoolCreatorAddedIterator{contract: _PoolCreator.contract, event: "PoolCreatorAdded", logs: logs, sub: sub}, nil
}

// WatchPoolCreatorAdded is a free log subscription operation binding the contract event 0xac89bb7b3d0c5a763a97f31bc75f8faee00426e7b235f02ca76897d55caf7b61.
//
// Solidity: event PoolCreatorAdded(address indexed account)
func (_PoolCreator *PoolCreatorFilterer) WatchPoolCreatorAdded(opts *bind.WatchOpts, sink chan<- *PoolCreatorPoolCreatorAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PoolCreator.contract.WatchLogs(opts, "PoolCreatorAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolCreatorPoolCreatorAdded)
				if err := _PoolCreator.contract.UnpackLog(event, "PoolCreatorAdded", log); err != nil {
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

// ParsePoolCreatorAdded is a log parse operation binding the contract event 0xac89bb7b3d0c5a763a97f31bc75f8faee00426e7b235f02ca76897d55caf7b61.
//
// Solidity: event PoolCreatorAdded(address indexed account)
func (_PoolCreator *PoolCreatorFilterer) ParsePoolCreatorAdded(log types.Log) (*PoolCreatorPoolCreatorAdded, error) {
	event := new(PoolCreatorPoolCreatorAdded)
	if err := _PoolCreator.contract.UnpackLog(event, "PoolCreatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolCreatorPoolCreatorRemovedIterator is returned from FilterPoolCreatorRemoved and is used to iterate over the raw logs and unpacked data for PoolCreatorRemoved events raised by the PoolCreator contract.
type PoolCreatorPoolCreatorRemovedIterator struct {
	Event *PoolCreatorPoolCreatorRemoved // Event containing the contract specifics and raw log

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
func (it *PoolCreatorPoolCreatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolCreatorPoolCreatorRemoved)
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
		it.Event = new(PoolCreatorPoolCreatorRemoved)
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
func (it *PoolCreatorPoolCreatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolCreatorPoolCreatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolCreatorPoolCreatorRemoved represents a PoolCreatorRemoved event raised by the PoolCreator contract.
type PoolCreatorPoolCreatorRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPoolCreatorRemoved is a free log retrieval operation binding the contract event 0x2bc71cc60b543df5fdd80ef5a45e1cdb344843769b649e9d901de3be24aeb24e.
//
// Solidity: event PoolCreatorRemoved(address indexed account)
func (_PoolCreator *PoolCreatorFilterer) FilterPoolCreatorRemoved(opts *bind.FilterOpts, account []common.Address) (*PoolCreatorPoolCreatorRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PoolCreator.contract.FilterLogs(opts, "PoolCreatorRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &PoolCreatorPoolCreatorRemovedIterator{contract: _PoolCreator.contract, event: "PoolCreatorRemoved", logs: logs, sub: sub}, nil
}

// WatchPoolCreatorRemoved is a free log subscription operation binding the contract event 0x2bc71cc60b543df5fdd80ef5a45e1cdb344843769b649e9d901de3be24aeb24e.
//
// Solidity: event PoolCreatorRemoved(address indexed account)
func (_PoolCreator *PoolCreatorFilterer) WatchPoolCreatorRemoved(opts *bind.WatchOpts, sink chan<- *PoolCreatorPoolCreatorRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _PoolCreator.contract.WatchLogs(opts, "PoolCreatorRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolCreatorPoolCreatorRemoved)
				if err := _PoolCreator.contract.UnpackLog(event, "PoolCreatorRemoved", log); err != nil {
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

// ParsePoolCreatorRemoved is a log parse operation binding the contract event 0x2bc71cc60b543df5fdd80ef5a45e1cdb344843769b649e9d901de3be24aeb24e.
//
// Solidity: event PoolCreatorRemoved(address indexed account)
func (_PoolCreator *PoolCreatorFilterer) ParsePoolCreatorRemoved(log types.Log) (*PoolCreatorPoolCreatorRemoved, error) {
	event := new(PoolCreatorPoolCreatorRemoved)
	if err := _PoolCreator.contract.UnpackLog(event, "PoolCreatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
