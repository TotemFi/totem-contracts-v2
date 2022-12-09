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

// WrappedTokenDistributorUpgradeableMetaData contains all meta data concerning the WrappedTokenDistributorUpgradeable contract.
var WrappedTokenDistributorUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DistributedBTC\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"swapRouterAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"BUSDContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WrappedTokenContractAddress\",\"type\":\"address\"}],\"name\":\"__WrappedTokenDistributor_initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"getEstimatedWrappedTokenForUSD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPathForUSDToWrappedToken\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSwapRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// WrappedTokenDistributorUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use WrappedTokenDistributorUpgradeableMetaData.ABI instead.
var WrappedTokenDistributorUpgradeableABI = WrappedTokenDistributorUpgradeableMetaData.ABI

// WrappedTokenDistributorUpgradeable is an auto generated Go binding around an Ethereum contract.
type WrappedTokenDistributorUpgradeable struct {
	WrappedTokenDistributorUpgradeableCaller     // Read-only binding to the contract
	WrappedTokenDistributorUpgradeableTransactor // Write-only binding to the contract
	WrappedTokenDistributorUpgradeableFilterer   // Log filterer for contract events
}

// WrappedTokenDistributorUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type WrappedTokenDistributorUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappedTokenDistributorUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WrappedTokenDistributorUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappedTokenDistributorUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WrappedTokenDistributorUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappedTokenDistributorUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WrappedTokenDistributorUpgradeableSession struct {
	Contract     *WrappedTokenDistributorUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                       // Call options to use throughout this session
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// WrappedTokenDistributorUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WrappedTokenDistributorUpgradeableCallerSession struct {
	Contract *WrappedTokenDistributorUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                             // Call options to use throughout this session
}

// WrappedTokenDistributorUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WrappedTokenDistributorUpgradeableTransactorSession struct {
	Contract     *WrappedTokenDistributorUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                             // Transaction auth options to use throughout this session
}

// WrappedTokenDistributorUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type WrappedTokenDistributorUpgradeableRaw struct {
	Contract *WrappedTokenDistributorUpgradeable // Generic contract binding to access the raw methods on
}

// WrappedTokenDistributorUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WrappedTokenDistributorUpgradeableCallerRaw struct {
	Contract *WrappedTokenDistributorUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// WrappedTokenDistributorUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WrappedTokenDistributorUpgradeableTransactorRaw struct {
	Contract *WrappedTokenDistributorUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWrappedTokenDistributorUpgradeable creates a new instance of WrappedTokenDistributorUpgradeable, bound to a specific deployed contract.
func NewWrappedTokenDistributorUpgradeable(address common.Address, backend bind.ContractBackend) (*WrappedTokenDistributorUpgradeable, error) {
	contract, err := bindWrappedTokenDistributorUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WrappedTokenDistributorUpgradeable{WrappedTokenDistributorUpgradeableCaller: WrappedTokenDistributorUpgradeableCaller{contract: contract}, WrappedTokenDistributorUpgradeableTransactor: WrappedTokenDistributorUpgradeableTransactor{contract: contract}, WrappedTokenDistributorUpgradeableFilterer: WrappedTokenDistributorUpgradeableFilterer{contract: contract}}, nil
}

// NewWrappedTokenDistributorUpgradeableCaller creates a new read-only instance of WrappedTokenDistributorUpgradeable, bound to a specific deployed contract.
func NewWrappedTokenDistributorUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*WrappedTokenDistributorUpgradeableCaller, error) {
	contract, err := bindWrappedTokenDistributorUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WrappedTokenDistributorUpgradeableCaller{contract: contract}, nil
}

// NewWrappedTokenDistributorUpgradeableTransactor creates a new write-only instance of WrappedTokenDistributorUpgradeable, bound to a specific deployed contract.
func NewWrappedTokenDistributorUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*WrappedTokenDistributorUpgradeableTransactor, error) {
	contract, err := bindWrappedTokenDistributorUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WrappedTokenDistributorUpgradeableTransactor{contract: contract}, nil
}

// NewWrappedTokenDistributorUpgradeableFilterer creates a new log filterer instance of WrappedTokenDistributorUpgradeable, bound to a specific deployed contract.
func NewWrappedTokenDistributorUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*WrappedTokenDistributorUpgradeableFilterer, error) {
	contract, err := bindWrappedTokenDistributorUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WrappedTokenDistributorUpgradeableFilterer{contract: contract}, nil
}

// bindWrappedTokenDistributorUpgradeable binds a generic wrapper to an already deployed contract.
func bindWrappedTokenDistributorUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WrappedTokenDistributorUpgradeableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WrappedTokenDistributorUpgradeable *WrappedTokenDistributorUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WrappedTokenDistributorUpgradeable.Contract.WrappedTokenDistributorUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WrappedTokenDistributorUpgradeable *WrappedTokenDistributorUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedTokenDistributorUpgradeable.Contract.WrappedTokenDistributorUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WrappedTokenDistributorUpgradeable *WrappedTokenDistributorUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrappedTokenDistributorUpgradeable.Contract.WrappedTokenDistributorUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WrappedTokenDistributorUpgradeable *WrappedTokenDistributorUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WrappedTokenDistributorUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WrappedTokenDistributorUpgradeable *WrappedTokenDistributorUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedTokenDistributorUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WrappedTokenDistributorUpgradeable *WrappedTokenDistributorUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrappedTokenDistributorUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// GetEstimatedWrappedTokenForUSD is a free data retrieval call binding the contract method 0x3fcdd82c.
//
// Solidity: function getEstimatedWrappedTokenForUSD(uint256 _amount) view returns(uint256)
func (_WrappedTokenDistributorUpgradeable *WrappedTokenDistributorUpgradeableCaller) GetEstimatedWrappedTokenForUSD(opts *bind.CallOpts, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _WrappedTokenDistributorUpgradeable.contract.Call(opts, &out, "getEstimatedWrappedTokenForUSD", _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEstimatedWrappedTokenForUSD is a free data retrieval call binding the contract method 0x3fcdd82c.
//
// Solidity: function getEstimatedWrappedTokenForUSD(uint256 _amount) view returns(uint256)
func (_WrappedTokenDistributorUpgradeable *WrappedTokenDistributorUpgradeableSession) GetEstimatedWrappedTokenForUSD(_amount *big.Int) (*big.Int, error) {
	return _WrappedTokenDistributorUpgradeable.Contract.GetEstimatedWrappedTokenForUSD(&_WrappedTokenDistributorUpgradeable.CallOpts, _amount)
}

// GetEstimatedWrappedTokenForUSD is a free data retrieval call binding the contract method 0x3fcdd82c.
//
// Solidity: function getEstimatedWrappedTokenForUSD(uint256 _amount) view returns(uint256)
func (_WrappedTokenDistributorUpgradeable *WrappedTokenDistributorUpgradeableCallerSession) GetEstimatedWrappedTokenForUSD(_amount *big.Int) (*big.Int, error) {
	return _WrappedTokenDistributorUpgradeable.Contract.GetEstimatedWrappedTokenForUSD(&_WrappedTokenDistributorUpgradeable.CallOpts, _amount)
}

// GetPathForUSDToWrappedToken is a free data retrieval call binding the contract method 0xa35ef3a7.
//
// Solidity: function getPathForUSDToWrappedToken() view returns(address[])
func (_WrappedTokenDistributorUpgradeable *WrappedTokenDistributorUpgradeableCaller) GetPathForUSDToWrappedToken(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _WrappedTokenDistributorUpgradeable.contract.Call(opts, &out, "getPathForUSDToWrappedToken")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPathForUSDToWrappedToken is a free data retrieval call binding the contract method 0xa35ef3a7.
//
// Solidity: function getPathForUSDToWrappedToken() view returns(address[])
func (_WrappedTokenDistributorUpgradeable *WrappedTokenDistributorUpgradeableSession) GetPathForUSDToWrappedToken() ([]common.Address, error) {
	return _WrappedTokenDistributorUpgradeable.Contract.GetPathForUSDToWrappedToken(&_WrappedTokenDistributorUpgradeable.CallOpts)
}

// GetPathForUSDToWrappedToken is a free data retrieval call binding the contract method 0xa35ef3a7.
//
// Solidity: function getPathForUSDToWrappedToken() view returns(address[])
func (_WrappedTokenDistributorUpgradeable *WrappedTokenDistributorUpgradeableCallerSession) GetPathForUSDToWrappedToken() ([]common.Address, error) {
	return _WrappedTokenDistributorUpgradeable.Contract.GetPathForUSDToWrappedToken(&_WrappedTokenDistributorUpgradeable.CallOpts)
}

// GetSwapRouter is a free data retrieval call binding the contract method 0x725c9c49.
//
// Solidity: function getSwapRouter() view returns(address)
func (_WrappedTokenDistributorUpgradeable *WrappedTokenDistributorUpgradeableCaller) GetSwapRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WrappedTokenDistributorUpgradeable.contract.Call(opts, &out, "getSwapRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSwapRouter is a free data retrieval call binding the contract method 0x725c9c49.
//
// Solidity: function getSwapRouter() view returns(address)
func (_WrappedTokenDistributorUpgradeable *WrappedTokenDistributorUpgradeableSession) GetSwapRouter() (common.Address, error) {
	return _WrappedTokenDistributorUpgradeable.Contract.GetSwapRouter(&_WrappedTokenDistributorUpgradeable.CallOpts)
}

// GetSwapRouter is a free data retrieval call binding the contract method 0x725c9c49.
//
// Solidity: function getSwapRouter() view returns(address)
func (_WrappedTokenDistributorUpgradeable *WrappedTokenDistributorUpgradeableCallerSession) GetSwapRouter() (common.Address, error) {
	return _WrappedTokenDistributorUpgradeable.Contract.GetSwapRouter(&_WrappedTokenDistributorUpgradeable.CallOpts)
}

// WrappedTokenDistributorInitialize is a paid mutator transaction binding the contract method 0xa8f7b346.
//
// Solidity: function __WrappedTokenDistributor_initialize(address swapRouterAddress, address BUSDContractAddress, address WrappedTokenContractAddress) returns()
func (_WrappedTokenDistributorUpgradeable *WrappedTokenDistributorUpgradeableTransactor) WrappedTokenDistributorInitialize(opts *bind.TransactOpts, swapRouterAddress common.Address, BUSDContractAddress common.Address, WrappedTokenContractAddress common.Address) (*types.Transaction, error) {
	return _WrappedTokenDistributorUpgradeable.contract.Transact(opts, "__WrappedTokenDistributor_initialize", swapRouterAddress, BUSDContractAddress, WrappedTokenContractAddress)
}

// WrappedTokenDistributorInitialize is a paid mutator transaction binding the contract method 0xa8f7b346.
//
// Solidity: function __WrappedTokenDistributor_initialize(address swapRouterAddress, address BUSDContractAddress, address WrappedTokenContractAddress) returns()
func (_WrappedTokenDistributorUpgradeable *WrappedTokenDistributorUpgradeableSession) WrappedTokenDistributorInitialize(swapRouterAddress common.Address, BUSDContractAddress common.Address, WrappedTokenContractAddress common.Address) (*types.Transaction, error) {
	return _WrappedTokenDistributorUpgradeable.Contract.WrappedTokenDistributorInitialize(&_WrappedTokenDistributorUpgradeable.TransactOpts, swapRouterAddress, BUSDContractAddress, WrappedTokenContractAddress)
}

// WrappedTokenDistributorInitialize is a paid mutator transaction binding the contract method 0xa8f7b346.
//
// Solidity: function __WrappedTokenDistributor_initialize(address swapRouterAddress, address BUSDContractAddress, address WrappedTokenContractAddress) returns()
func (_WrappedTokenDistributorUpgradeable *WrappedTokenDistributorUpgradeableTransactorSession) WrappedTokenDistributorInitialize(swapRouterAddress common.Address, BUSDContractAddress common.Address, WrappedTokenContractAddress common.Address) (*types.Transaction, error) {
	return _WrappedTokenDistributorUpgradeable.Contract.WrappedTokenDistributorInitialize(&_WrappedTokenDistributorUpgradeable.TransactOpts, swapRouterAddress, BUSDContractAddress, WrappedTokenContractAddress)
}

// WrappedTokenDistributorUpgradeableDistributedBTCIterator is returned from FilterDistributedBTC and is used to iterate over the raw logs and unpacked data for DistributedBTC events raised by the WrappedTokenDistributorUpgradeable contract.
type WrappedTokenDistributorUpgradeableDistributedBTCIterator struct {
	Event *WrappedTokenDistributorUpgradeableDistributedBTC // Event containing the contract specifics and raw log

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
func (it *WrappedTokenDistributorUpgradeableDistributedBTCIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedTokenDistributorUpgradeableDistributedBTC)
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
		it.Event = new(WrappedTokenDistributorUpgradeableDistributedBTC)
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
func (it *WrappedTokenDistributorUpgradeableDistributedBTCIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappedTokenDistributorUpgradeableDistributedBTCIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappedTokenDistributorUpgradeableDistributedBTC represents a DistributedBTC event raised by the WrappedTokenDistributorUpgradeable contract.
type WrappedTokenDistributorUpgradeableDistributedBTC struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDistributedBTC is a free log retrieval operation binding the contract event 0xe690d17e165a33f4a6cf4c2374c904d5f425eb2563320d08ed595794a1ba33aa.
//
// Solidity: event DistributedBTC(address indexed to, uint256 amount)
func (_WrappedTokenDistributorUpgradeable *WrappedTokenDistributorUpgradeableFilterer) FilterDistributedBTC(opts *bind.FilterOpts, to []common.Address) (*WrappedTokenDistributorUpgradeableDistributedBTCIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrappedTokenDistributorUpgradeable.contract.FilterLogs(opts, "DistributedBTC", toRule)
	if err != nil {
		return nil, err
	}
	return &WrappedTokenDistributorUpgradeableDistributedBTCIterator{contract: _WrappedTokenDistributorUpgradeable.contract, event: "DistributedBTC", logs: logs, sub: sub}, nil
}

// WatchDistributedBTC is a free log subscription operation binding the contract event 0xe690d17e165a33f4a6cf4c2374c904d5f425eb2563320d08ed595794a1ba33aa.
//
// Solidity: event DistributedBTC(address indexed to, uint256 amount)
func (_WrappedTokenDistributorUpgradeable *WrappedTokenDistributorUpgradeableFilterer) WatchDistributedBTC(opts *bind.WatchOpts, sink chan<- *WrappedTokenDistributorUpgradeableDistributedBTC, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrappedTokenDistributorUpgradeable.contract.WatchLogs(opts, "DistributedBTC", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappedTokenDistributorUpgradeableDistributedBTC)
				if err := _WrappedTokenDistributorUpgradeable.contract.UnpackLog(event, "DistributedBTC", log); err != nil {
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

// ParseDistributedBTC is a log parse operation binding the contract event 0xe690d17e165a33f4a6cf4c2374c904d5f425eb2563320d08ed595794a1ba33aa.
//
// Solidity: event DistributedBTC(address indexed to, uint256 amount)
func (_WrappedTokenDistributorUpgradeable *WrappedTokenDistributorUpgradeableFilterer) ParseDistributedBTC(log types.Log) (*WrappedTokenDistributorUpgradeableDistributedBTC, error) {
	event := new(WrappedTokenDistributorUpgradeableDistributedBTC)
	if err := _WrappedTokenDistributorUpgradeable.contract.UnpackLog(event, "DistributedBTC", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
