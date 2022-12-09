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

// IRewardManagerMetaData contains all meta data concerning the IRewardManager contract.
var IRewardManagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"SetOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewarder\",\"type\":\"address\"}],\"name\":\"SetRewarder\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolAddress\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"rewardUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOperator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IRewardManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use IRewardManagerMetaData.ABI instead.
var IRewardManagerABI = IRewardManagerMetaData.ABI

// IRewardManager is an auto generated Go binding around an Ethereum contract.
type IRewardManager struct {
	IRewardManagerCaller     // Read-only binding to the contract
	IRewardManagerTransactor // Write-only binding to the contract
	IRewardManagerFilterer   // Log filterer for contract events
}

// IRewardManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type IRewardManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRewardManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IRewardManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRewardManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IRewardManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IRewardManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IRewardManagerSession struct {
	Contract     *IRewardManager   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IRewardManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IRewardManagerCallerSession struct {
	Contract *IRewardManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IRewardManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IRewardManagerTransactorSession struct {
	Contract     *IRewardManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IRewardManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type IRewardManagerRaw struct {
	Contract *IRewardManager // Generic contract binding to access the raw methods on
}

// IRewardManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IRewardManagerCallerRaw struct {
	Contract *IRewardManagerCaller // Generic read-only contract binding to access the raw methods on
}

// IRewardManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IRewardManagerTransactorRaw struct {
	Contract *IRewardManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIRewardManager creates a new instance of IRewardManager, bound to a specific deployed contract.
func NewIRewardManager(address common.Address, backend bind.ContractBackend) (*IRewardManager, error) {
	contract, err := bindIRewardManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IRewardManager{IRewardManagerCaller: IRewardManagerCaller{contract: contract}, IRewardManagerTransactor: IRewardManagerTransactor{contract: contract}, IRewardManagerFilterer: IRewardManagerFilterer{contract: contract}}, nil
}

// NewIRewardManagerCaller creates a new read-only instance of IRewardManager, bound to a specific deployed contract.
func NewIRewardManagerCaller(address common.Address, caller bind.ContractCaller) (*IRewardManagerCaller, error) {
	contract, err := bindIRewardManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IRewardManagerCaller{contract: contract}, nil
}

// NewIRewardManagerTransactor creates a new write-only instance of IRewardManager, bound to a specific deployed contract.
func NewIRewardManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*IRewardManagerTransactor, error) {
	contract, err := bindIRewardManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IRewardManagerTransactor{contract: contract}, nil
}

// NewIRewardManagerFilterer creates a new log filterer instance of IRewardManager, bound to a specific deployed contract.
func NewIRewardManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*IRewardManagerFilterer, error) {
	contract, err := bindIRewardManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IRewardManagerFilterer{contract: contract}, nil
}

// bindIRewardManager binds a generic wrapper to an already deployed contract.
func bindIRewardManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IRewardManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRewardManager *IRewardManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRewardManager.Contract.IRewardManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRewardManager *IRewardManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRewardManager.Contract.IRewardManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRewardManager *IRewardManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRewardManager.Contract.IRewardManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IRewardManager *IRewardManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IRewardManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IRewardManager *IRewardManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IRewardManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IRewardManager *IRewardManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IRewardManager.Contract.contract.Transact(opts, method, params...)
}

// AddPool is a paid mutator transaction binding the contract method 0xd914cd4b.
//
// Solidity: function addPool(address _poolAddress) returns()
func (_IRewardManager *IRewardManagerTransactor) AddPool(opts *bind.TransactOpts, _poolAddress common.Address) (*types.Transaction, error) {
	return _IRewardManager.contract.Transact(opts, "addPool", _poolAddress)
}

// AddPool is a paid mutator transaction binding the contract method 0xd914cd4b.
//
// Solidity: function addPool(address _poolAddress) returns()
func (_IRewardManager *IRewardManagerSession) AddPool(_poolAddress common.Address) (*types.Transaction, error) {
	return _IRewardManager.Contract.AddPool(&_IRewardManager.TransactOpts, _poolAddress)
}

// AddPool is a paid mutator transaction binding the contract method 0xd914cd4b.
//
// Solidity: function addPool(address _poolAddress) returns()
func (_IRewardManager *IRewardManagerTransactorSession) AddPool(_poolAddress common.Address) (*types.Transaction, error) {
	return _IRewardManager.Contract.AddPool(&_IRewardManager.TransactOpts, _poolAddress)
}

// RewardUser is a paid mutator transaction binding the contract method 0xe4e103dc.
//
// Solidity: function rewardUser(address _user, uint256 _amount) returns()
func (_IRewardManager *IRewardManagerTransactor) RewardUser(opts *bind.TransactOpts, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IRewardManager.contract.Transact(opts, "rewardUser", _user, _amount)
}

// RewardUser is a paid mutator transaction binding the contract method 0xe4e103dc.
//
// Solidity: function rewardUser(address _user, uint256 _amount) returns()
func (_IRewardManager *IRewardManagerSession) RewardUser(_user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IRewardManager.Contract.RewardUser(&_IRewardManager.TransactOpts, _user, _amount)
}

// RewardUser is a paid mutator transaction binding the contract method 0xe4e103dc.
//
// Solidity: function rewardUser(address _user, uint256 _amount) returns()
func (_IRewardManager *IRewardManagerTransactorSession) RewardUser(_user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _IRewardManager.Contract.RewardUser(&_IRewardManager.TransactOpts, _user, _amount)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _newOperator) returns()
func (_IRewardManager *IRewardManagerTransactor) SetOperator(opts *bind.TransactOpts, _newOperator common.Address) (*types.Transaction, error) {
	return _IRewardManager.contract.Transact(opts, "setOperator", _newOperator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _newOperator) returns()
func (_IRewardManager *IRewardManagerSession) SetOperator(_newOperator common.Address) (*types.Transaction, error) {
	return _IRewardManager.Contract.SetOperator(&_IRewardManager.TransactOpts, _newOperator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _newOperator) returns()
func (_IRewardManager *IRewardManagerTransactorSession) SetOperator(_newOperator common.Address) (*types.Transaction, error) {
	return _IRewardManager.Contract.SetOperator(&_IRewardManager.TransactOpts, _newOperator)
}

// IRewardManagerSetOperatorIterator is returned from FilterSetOperator and is used to iterate over the raw logs and unpacked data for SetOperator events raised by the IRewardManager contract.
type IRewardManagerSetOperatorIterator struct {
	Event *IRewardManagerSetOperator // Event containing the contract specifics and raw log

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
func (it *IRewardManagerSetOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRewardManagerSetOperator)
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
		it.Event = new(IRewardManagerSetOperator)
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
func (it *IRewardManagerSetOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRewardManagerSetOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRewardManagerSetOperator represents a SetOperator event raised by the IRewardManager contract.
type IRewardManagerSetOperator struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetOperator is a free log retrieval operation binding the contract event 0xdbebfba65bd6398fb722063efc10c99f624f9cd8ba657201056af918a676d5ee.
//
// Solidity: event SetOperator(address operator)
func (_IRewardManager *IRewardManagerFilterer) FilterSetOperator(opts *bind.FilterOpts) (*IRewardManagerSetOperatorIterator, error) {

	logs, sub, err := _IRewardManager.contract.FilterLogs(opts, "SetOperator")
	if err != nil {
		return nil, err
	}
	return &IRewardManagerSetOperatorIterator{contract: _IRewardManager.contract, event: "SetOperator", logs: logs, sub: sub}, nil
}

// WatchSetOperator is a free log subscription operation binding the contract event 0xdbebfba65bd6398fb722063efc10c99f624f9cd8ba657201056af918a676d5ee.
//
// Solidity: event SetOperator(address operator)
func (_IRewardManager *IRewardManagerFilterer) WatchSetOperator(opts *bind.WatchOpts, sink chan<- *IRewardManagerSetOperator) (event.Subscription, error) {

	logs, sub, err := _IRewardManager.contract.WatchLogs(opts, "SetOperator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRewardManagerSetOperator)
				if err := _IRewardManager.contract.UnpackLog(event, "SetOperator", log); err != nil {
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

// ParseSetOperator is a log parse operation binding the contract event 0xdbebfba65bd6398fb722063efc10c99f624f9cd8ba657201056af918a676d5ee.
//
// Solidity: event SetOperator(address operator)
func (_IRewardManager *IRewardManagerFilterer) ParseSetOperator(log types.Log) (*IRewardManagerSetOperator, error) {
	event := new(IRewardManagerSetOperator)
	if err := _IRewardManager.contract.UnpackLog(event, "SetOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IRewardManagerSetRewarderIterator is returned from FilterSetRewarder and is used to iterate over the raw logs and unpacked data for SetRewarder events raised by the IRewardManager contract.
type IRewardManagerSetRewarderIterator struct {
	Event *IRewardManagerSetRewarder // Event containing the contract specifics and raw log

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
func (it *IRewardManagerSetRewarderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IRewardManagerSetRewarder)
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
		it.Event = new(IRewardManagerSetRewarder)
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
func (it *IRewardManagerSetRewarderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IRewardManagerSetRewarderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IRewardManagerSetRewarder represents a SetRewarder event raised by the IRewardManager contract.
type IRewardManagerSetRewarder struct {
	Rewarder common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetRewarder is a free log retrieval operation binding the contract event 0xcf0aff36caea97f7ad632b334936cb196014c193ac1a790b578f12a70d9836db.
//
// Solidity: event SetRewarder(address rewarder)
func (_IRewardManager *IRewardManagerFilterer) FilterSetRewarder(opts *bind.FilterOpts) (*IRewardManagerSetRewarderIterator, error) {

	logs, sub, err := _IRewardManager.contract.FilterLogs(opts, "SetRewarder")
	if err != nil {
		return nil, err
	}
	return &IRewardManagerSetRewarderIterator{contract: _IRewardManager.contract, event: "SetRewarder", logs: logs, sub: sub}, nil
}

// WatchSetRewarder is a free log subscription operation binding the contract event 0xcf0aff36caea97f7ad632b334936cb196014c193ac1a790b578f12a70d9836db.
//
// Solidity: event SetRewarder(address rewarder)
func (_IRewardManager *IRewardManagerFilterer) WatchSetRewarder(opts *bind.WatchOpts, sink chan<- *IRewardManagerSetRewarder) (event.Subscription, error) {

	logs, sub, err := _IRewardManager.contract.WatchLogs(opts, "SetRewarder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IRewardManagerSetRewarder)
				if err := _IRewardManager.contract.UnpackLog(event, "SetRewarder", log); err != nil {
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

// ParseSetRewarder is a log parse operation binding the contract event 0xcf0aff36caea97f7ad632b334936cb196014c193ac1a790b578f12a70d9836db.
//
// Solidity: event SetRewarder(address rewarder)
func (_IRewardManager *IRewardManagerFilterer) ParseSetRewarder(log types.Log) (*IRewardManagerSetRewarder, error) {
	event := new(IRewardManagerSetRewarder)
	if err := _IRewardManager.contract.UnpackLog(event, "SetRewarder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
