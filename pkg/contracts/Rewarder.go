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

// RewarderMetaData contains all meta data concerning the Rewarder contract.
var RewarderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RewarderAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RewarderRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addRewarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isRewarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceRewarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RewarderABI is the input ABI used to generate the binding from.
// Deprecated: Use RewarderMetaData.ABI instead.
var RewarderABI = RewarderMetaData.ABI

// Rewarder is an auto generated Go binding around an Ethereum contract.
type Rewarder struct {
	RewarderCaller     // Read-only binding to the contract
	RewarderTransactor // Write-only binding to the contract
	RewarderFilterer   // Log filterer for contract events
}

// RewarderCaller is an auto generated read-only Go binding around an Ethereum contract.
type RewarderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewarderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RewarderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewarderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RewarderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewarderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RewarderSession struct {
	Contract     *Rewarder         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RewarderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RewarderCallerSession struct {
	Contract *RewarderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// RewarderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RewarderTransactorSession struct {
	Contract     *RewarderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// RewarderRaw is an auto generated low-level Go binding around an Ethereum contract.
type RewarderRaw struct {
	Contract *Rewarder // Generic contract binding to access the raw methods on
}

// RewarderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RewarderCallerRaw struct {
	Contract *RewarderCaller // Generic read-only contract binding to access the raw methods on
}

// RewarderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RewarderTransactorRaw struct {
	Contract *RewarderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRewarder creates a new instance of Rewarder, bound to a specific deployed contract.
func NewRewarder(address common.Address, backend bind.ContractBackend) (*Rewarder, error) {
	contract, err := bindRewarder(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rewarder{RewarderCaller: RewarderCaller{contract: contract}, RewarderTransactor: RewarderTransactor{contract: contract}, RewarderFilterer: RewarderFilterer{contract: contract}}, nil
}

// NewRewarderCaller creates a new read-only instance of Rewarder, bound to a specific deployed contract.
func NewRewarderCaller(address common.Address, caller bind.ContractCaller) (*RewarderCaller, error) {
	contract, err := bindRewarder(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewarderCaller{contract: contract}, nil
}

// NewRewarderTransactor creates a new write-only instance of Rewarder, bound to a specific deployed contract.
func NewRewarderTransactor(address common.Address, transactor bind.ContractTransactor) (*RewarderTransactor, error) {
	contract, err := bindRewarder(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewarderTransactor{contract: contract}, nil
}

// NewRewarderFilterer creates a new log filterer instance of Rewarder, bound to a specific deployed contract.
func NewRewarderFilterer(address common.Address, filterer bind.ContractFilterer) (*RewarderFilterer, error) {
	contract, err := bindRewarder(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewarderFilterer{contract: contract}, nil
}

// bindRewarder binds a generic wrapper to an already deployed contract.
func bindRewarder(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RewarderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rewarder *RewarderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rewarder.Contract.RewarderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rewarder *RewarderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewarder.Contract.RewarderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rewarder *RewarderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rewarder.Contract.RewarderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rewarder *RewarderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Rewarder.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rewarder *RewarderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewarder.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rewarder *RewarderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rewarder.Contract.contract.Transact(opts, method, params...)
}

// IsRewarder is a free data retrieval call binding the contract method 0x8e478cab.
//
// Solidity: function isRewarder(address account) view returns(bool)
func (_Rewarder *RewarderCaller) IsRewarder(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _Rewarder.contract.Call(opts, &out, "isRewarder", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRewarder is a free data retrieval call binding the contract method 0x8e478cab.
//
// Solidity: function isRewarder(address account) view returns(bool)
func (_Rewarder *RewarderSession) IsRewarder(account common.Address) (bool, error) {
	return _Rewarder.Contract.IsRewarder(&_Rewarder.CallOpts, account)
}

// IsRewarder is a free data retrieval call binding the contract method 0x8e478cab.
//
// Solidity: function isRewarder(address account) view returns(bool)
func (_Rewarder *RewarderCallerSession) IsRewarder(account common.Address) (bool, error) {
	return _Rewarder.Contract.IsRewarder(&_Rewarder.CallOpts, account)
}

// AddRewarder is a paid mutator transaction binding the contract method 0x56d3590b.
//
// Solidity: function addRewarder(address account) returns()
func (_Rewarder *RewarderTransactor) AddRewarder(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Rewarder.contract.Transact(opts, "addRewarder", account)
}

// AddRewarder is a paid mutator transaction binding the contract method 0x56d3590b.
//
// Solidity: function addRewarder(address account) returns()
func (_Rewarder *RewarderSession) AddRewarder(account common.Address) (*types.Transaction, error) {
	return _Rewarder.Contract.AddRewarder(&_Rewarder.TransactOpts, account)
}

// AddRewarder is a paid mutator transaction binding the contract method 0x56d3590b.
//
// Solidity: function addRewarder(address account) returns()
func (_Rewarder *RewarderTransactorSession) AddRewarder(account common.Address) (*types.Transaction, error) {
	return _Rewarder.Contract.AddRewarder(&_Rewarder.TransactOpts, account)
}

// RenounceRewarder is a paid mutator transaction binding the contract method 0x40cc8518.
//
// Solidity: function renounceRewarder() returns()
func (_Rewarder *RewarderTransactor) RenounceRewarder(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rewarder.contract.Transact(opts, "renounceRewarder")
}

// RenounceRewarder is a paid mutator transaction binding the contract method 0x40cc8518.
//
// Solidity: function renounceRewarder() returns()
func (_Rewarder *RewarderSession) RenounceRewarder() (*types.Transaction, error) {
	return _Rewarder.Contract.RenounceRewarder(&_Rewarder.TransactOpts)
}

// RenounceRewarder is a paid mutator transaction binding the contract method 0x40cc8518.
//
// Solidity: function renounceRewarder() returns()
func (_Rewarder *RewarderTransactorSession) RenounceRewarder() (*types.Transaction, error) {
	return _Rewarder.Contract.RenounceRewarder(&_Rewarder.TransactOpts)
}

// RewarderRewarderAddedIterator is returned from FilterRewarderAdded and is used to iterate over the raw logs and unpacked data for RewarderAdded events raised by the Rewarder contract.
type RewarderRewarderAddedIterator struct {
	Event *RewarderRewarderAdded // Event containing the contract specifics and raw log

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
func (it *RewarderRewarderAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewarderRewarderAdded)
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
		it.Event = new(RewarderRewarderAdded)
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
func (it *RewarderRewarderAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewarderRewarderAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewarderRewarderAdded represents a RewarderAdded event raised by the Rewarder contract.
type RewarderRewarderAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRewarderAdded is a free log retrieval operation binding the contract event 0x9dfd431959d2d3358e3eb909555ad574123ea5881ff0e05a80f66d4984710c1b.
//
// Solidity: event RewarderAdded(address indexed account)
func (_Rewarder *RewarderFilterer) FilterRewarderAdded(opts *bind.FilterOpts, account []common.Address) (*RewarderRewarderAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Rewarder.contract.FilterLogs(opts, "RewarderAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &RewarderRewarderAddedIterator{contract: _Rewarder.contract, event: "RewarderAdded", logs: logs, sub: sub}, nil
}

// WatchRewarderAdded is a free log subscription operation binding the contract event 0x9dfd431959d2d3358e3eb909555ad574123ea5881ff0e05a80f66d4984710c1b.
//
// Solidity: event RewarderAdded(address indexed account)
func (_Rewarder *RewarderFilterer) WatchRewarderAdded(opts *bind.WatchOpts, sink chan<- *RewarderRewarderAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Rewarder.contract.WatchLogs(opts, "RewarderAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewarderRewarderAdded)
				if err := _Rewarder.contract.UnpackLog(event, "RewarderAdded", log); err != nil {
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

// ParseRewarderAdded is a log parse operation binding the contract event 0x9dfd431959d2d3358e3eb909555ad574123ea5881ff0e05a80f66d4984710c1b.
//
// Solidity: event RewarderAdded(address indexed account)
func (_Rewarder *RewarderFilterer) ParseRewarderAdded(log types.Log) (*RewarderRewarderAdded, error) {
	event := new(RewarderRewarderAdded)
	if err := _Rewarder.contract.UnpackLog(event, "RewarderAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewarderRewarderRemovedIterator is returned from FilterRewarderRemoved and is used to iterate over the raw logs and unpacked data for RewarderRemoved events raised by the Rewarder contract.
type RewarderRewarderRemovedIterator struct {
	Event *RewarderRewarderRemoved // Event containing the contract specifics and raw log

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
func (it *RewarderRewarderRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewarderRewarderRemoved)
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
		it.Event = new(RewarderRewarderRemoved)
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
func (it *RewarderRewarderRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewarderRewarderRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewarderRewarderRemoved represents a RewarderRemoved event raised by the Rewarder contract.
type RewarderRewarderRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRewarderRemoved is a free log retrieval operation binding the contract event 0xce699c579f0b70ea4ccd6a4b38be26726a2c248b89c7102ccbc5d0f3060ef6d0.
//
// Solidity: event RewarderRemoved(address indexed account)
func (_Rewarder *RewarderFilterer) FilterRewarderRemoved(opts *bind.FilterOpts, account []common.Address) (*RewarderRewarderRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Rewarder.contract.FilterLogs(opts, "RewarderRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &RewarderRewarderRemovedIterator{contract: _Rewarder.contract, event: "RewarderRemoved", logs: logs, sub: sub}, nil
}

// WatchRewarderRemoved is a free log subscription operation binding the contract event 0xce699c579f0b70ea4ccd6a4b38be26726a2c248b89c7102ccbc5d0f3060ef6d0.
//
// Solidity: event RewarderRemoved(address indexed account)
func (_Rewarder *RewarderFilterer) WatchRewarderRemoved(opts *bind.WatchOpts, sink chan<- *RewarderRewarderRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Rewarder.contract.WatchLogs(opts, "RewarderRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewarderRewarderRemoved)
				if err := _Rewarder.contract.UnpackLog(event, "RewarderRemoved", log); err != nil {
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

// ParseRewarderRemoved is a log parse operation binding the contract event 0xce699c579f0b70ea4ccd6a4b38be26726a2c248b89c7102ccbc5d0f3060ef6d0.
//
// Solidity: event RewarderRemoved(address indexed account)
func (_Rewarder *RewarderFilterer) ParseRewarderRemoved(log types.Log) (*RewarderRewarderRemoved, error) {
	event := new(RewarderRewarderRemoved)
	if err := _Rewarder.contract.UnpackLog(event, "RewarderRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
