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

// ILockerUserMetaData contains all meta data concerning the ILockerUser contract.
var ILockerUserMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"locker\",\"type\":\"address\"}],\"name\":\"SetLocker\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"locker\",\"outputs\":[{\"internalType\":\"contractILocker\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ILockerUserABI is the input ABI used to generate the binding from.
// Deprecated: Use ILockerUserMetaData.ABI instead.
var ILockerUserABI = ILockerUserMetaData.ABI

// ILockerUser is an auto generated Go binding around an Ethereum contract.
type ILockerUser struct {
	ILockerUserCaller     // Read-only binding to the contract
	ILockerUserTransactor // Write-only binding to the contract
	ILockerUserFilterer   // Log filterer for contract events
}

// ILockerUserCaller is an auto generated read-only Go binding around an Ethereum contract.
type ILockerUserCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILockerUserTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ILockerUserTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILockerUserFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ILockerUserFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILockerUserSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ILockerUserSession struct {
	Contract     *ILockerUser      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ILockerUserCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ILockerUserCallerSession struct {
	Contract *ILockerUserCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ILockerUserTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ILockerUserTransactorSession struct {
	Contract     *ILockerUserTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ILockerUserRaw is an auto generated low-level Go binding around an Ethereum contract.
type ILockerUserRaw struct {
	Contract *ILockerUser // Generic contract binding to access the raw methods on
}

// ILockerUserCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ILockerUserCallerRaw struct {
	Contract *ILockerUserCaller // Generic read-only contract binding to access the raw methods on
}

// ILockerUserTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ILockerUserTransactorRaw struct {
	Contract *ILockerUserTransactor // Generic write-only contract binding to access the raw methods on
}

// NewILockerUser creates a new instance of ILockerUser, bound to a specific deployed contract.
func NewILockerUser(address common.Address, backend bind.ContractBackend) (*ILockerUser, error) {
	contract, err := bindILockerUser(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ILockerUser{ILockerUserCaller: ILockerUserCaller{contract: contract}, ILockerUserTransactor: ILockerUserTransactor{contract: contract}, ILockerUserFilterer: ILockerUserFilterer{contract: contract}}, nil
}

// NewILockerUserCaller creates a new read-only instance of ILockerUser, bound to a specific deployed contract.
func NewILockerUserCaller(address common.Address, caller bind.ContractCaller) (*ILockerUserCaller, error) {
	contract, err := bindILockerUser(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ILockerUserCaller{contract: contract}, nil
}

// NewILockerUserTransactor creates a new write-only instance of ILockerUser, bound to a specific deployed contract.
func NewILockerUserTransactor(address common.Address, transactor bind.ContractTransactor) (*ILockerUserTransactor, error) {
	contract, err := bindILockerUser(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ILockerUserTransactor{contract: contract}, nil
}

// NewILockerUserFilterer creates a new log filterer instance of ILockerUser, bound to a specific deployed contract.
func NewILockerUserFilterer(address common.Address, filterer bind.ContractFilterer) (*ILockerUserFilterer, error) {
	contract, err := bindILockerUser(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ILockerUserFilterer{contract: contract}, nil
}

// bindILockerUser binds a generic wrapper to an already deployed contract.
func bindILockerUser(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ILockerUserABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILockerUser *ILockerUserRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILockerUser.Contract.ILockerUserCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILockerUser *ILockerUserRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILockerUser.Contract.ILockerUserTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILockerUser *ILockerUserRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILockerUser.Contract.ILockerUserTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILockerUser *ILockerUserCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILockerUser.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILockerUser *ILockerUserTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILockerUser.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILockerUser *ILockerUserTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILockerUser.Contract.contract.Transact(opts, method, params...)
}

// Locker is a free data retrieval call binding the contract method 0xd7b96d4e.
//
// Solidity: function locker() view returns(address)
func (_ILockerUser *ILockerUserCaller) Locker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILockerUser.contract.Call(opts, &out, "locker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Locker is a free data retrieval call binding the contract method 0xd7b96d4e.
//
// Solidity: function locker() view returns(address)
func (_ILockerUser *ILockerUserSession) Locker() (common.Address, error) {
	return _ILockerUser.Contract.Locker(&_ILockerUser.CallOpts)
}

// Locker is a free data retrieval call binding the contract method 0xd7b96d4e.
//
// Solidity: function locker() view returns(address)
func (_ILockerUser *ILockerUserCallerSession) Locker() (common.Address, error) {
	return _ILockerUser.Contract.Locker(&_ILockerUser.CallOpts)
}

// ILockerUserSetLockerIterator is returned from FilterSetLocker and is used to iterate over the raw logs and unpacked data for SetLocker events raised by the ILockerUser contract.
type ILockerUserSetLockerIterator struct {
	Event *ILockerUserSetLocker // Event containing the contract specifics and raw log

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
func (it *ILockerUserSetLockerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILockerUserSetLocker)
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
		it.Event = new(ILockerUserSetLocker)
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
func (it *ILockerUserSetLockerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILockerUserSetLockerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILockerUserSetLocker represents a SetLocker event raised by the ILockerUser contract.
type ILockerUserSetLocker struct {
	Locker common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetLocker is a free log retrieval operation binding the contract event 0x2ad3bcdfc9b535ec0e163460ef3e0b8015fca3aa8e9cc778ca27876a07bd2827.
//
// Solidity: event SetLocker(address indexed locker)
func (_ILockerUser *ILockerUserFilterer) FilterSetLocker(opts *bind.FilterOpts, locker []common.Address) (*ILockerUserSetLockerIterator, error) {

	var lockerRule []interface{}
	for _, lockerItem := range locker {
		lockerRule = append(lockerRule, lockerItem)
	}

	logs, sub, err := _ILockerUser.contract.FilterLogs(opts, "SetLocker", lockerRule)
	if err != nil {
		return nil, err
	}
	return &ILockerUserSetLockerIterator{contract: _ILockerUser.contract, event: "SetLocker", logs: logs, sub: sub}, nil
}

// WatchSetLocker is a free log subscription operation binding the contract event 0x2ad3bcdfc9b535ec0e163460ef3e0b8015fca3aa8e9cc778ca27876a07bd2827.
//
// Solidity: event SetLocker(address indexed locker)
func (_ILockerUser *ILockerUserFilterer) WatchSetLocker(opts *bind.WatchOpts, sink chan<- *ILockerUserSetLocker, locker []common.Address) (event.Subscription, error) {

	var lockerRule []interface{}
	for _, lockerItem := range locker {
		lockerRule = append(lockerRule, lockerItem)
	}

	logs, sub, err := _ILockerUser.contract.WatchLogs(opts, "SetLocker", lockerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILockerUserSetLocker)
				if err := _ILockerUser.contract.UnpackLog(event, "SetLocker", log); err != nil {
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

// ParseSetLocker is a log parse operation binding the contract event 0x2ad3bcdfc9b535ec0e163460ef3e0b8015fca3aa8e9cc778ca27876a07bd2827.
//
// Solidity: event SetLocker(address indexed locker)
func (_ILockerUser *ILockerUserFilterer) ParseSetLocker(log types.Log) (*ILockerUserSetLocker, error) {
	event := new(ILockerUserSetLocker)
	if err := _ILockerUser.contract.UnpackLog(event, "SetLocker", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
