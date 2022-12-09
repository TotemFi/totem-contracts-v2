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

// RewardManagerMetaData contains all meta data concerning the RewardManager contract.
var RewardManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractTotemToken\",\"name\":\"_totemToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OperatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OperatorRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RewarderAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"RewarderRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"SetOperator\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"rewarder\",\"type\":\"address\"}],\"name\":\"SetRewarder\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolAddress\",\"type\":\"address\"}],\"name\":\"addPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addRewarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isRewarder\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceRewarder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"rewardUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOperator\",\"type\":\"address\"}],\"name\":\"setOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// RewardManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use RewardManagerMetaData.ABI instead.
var RewardManagerABI = RewardManagerMetaData.ABI

// RewardManager is an auto generated Go binding around an Ethereum contract.
type RewardManager struct {
	RewardManagerCaller     // Read-only binding to the contract
	RewardManagerTransactor // Write-only binding to the contract
	RewardManagerFilterer   // Log filterer for contract events
}

// RewardManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type RewardManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RewardManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RewardManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RewardManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RewardManagerSession struct {
	Contract     *RewardManager    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RewardManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RewardManagerCallerSession struct {
	Contract *RewardManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// RewardManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RewardManagerTransactorSession struct {
	Contract     *RewardManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// RewardManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type RewardManagerRaw struct {
	Contract *RewardManager // Generic contract binding to access the raw methods on
}

// RewardManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RewardManagerCallerRaw struct {
	Contract *RewardManagerCaller // Generic read-only contract binding to access the raw methods on
}

// RewardManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RewardManagerTransactorRaw struct {
	Contract *RewardManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRewardManager creates a new instance of RewardManager, bound to a specific deployed contract.
func NewRewardManager(address common.Address, backend bind.ContractBackend) (*RewardManager, error) {
	contract, err := bindRewardManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RewardManager{RewardManagerCaller: RewardManagerCaller{contract: contract}, RewardManagerTransactor: RewardManagerTransactor{contract: contract}, RewardManagerFilterer: RewardManagerFilterer{contract: contract}}, nil
}

// NewRewardManagerCaller creates a new read-only instance of RewardManager, bound to a specific deployed contract.
func NewRewardManagerCaller(address common.Address, caller bind.ContractCaller) (*RewardManagerCaller, error) {
	contract, err := bindRewardManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RewardManagerCaller{contract: contract}, nil
}

// NewRewardManagerTransactor creates a new write-only instance of RewardManager, bound to a specific deployed contract.
func NewRewardManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*RewardManagerTransactor, error) {
	contract, err := bindRewardManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RewardManagerTransactor{contract: contract}, nil
}

// NewRewardManagerFilterer creates a new log filterer instance of RewardManager, bound to a specific deployed contract.
func NewRewardManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*RewardManagerFilterer, error) {
	contract, err := bindRewardManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RewardManagerFilterer{contract: contract}, nil
}

// bindRewardManager binds a generic wrapper to an already deployed contract.
func bindRewardManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RewardManagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardManager *RewardManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewardManager.Contract.RewardManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardManager *RewardManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardManager.Contract.RewardManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardManager *RewardManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardManager.Contract.RewardManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RewardManager *RewardManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RewardManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RewardManager *RewardManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RewardManager *RewardManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RewardManager.Contract.contract.Transact(opts, method, params...)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address account) view returns(bool)
func (_RewardManager *RewardManagerCaller) IsOperator(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "isOperator", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address account) view returns(bool)
func (_RewardManager *RewardManagerSession) IsOperator(account common.Address) (bool, error) {
	return _RewardManager.Contract.IsOperator(&_RewardManager.CallOpts, account)
}

// IsOperator is a free data retrieval call binding the contract method 0x6d70f7ae.
//
// Solidity: function isOperator(address account) view returns(bool)
func (_RewardManager *RewardManagerCallerSession) IsOperator(account common.Address) (bool, error) {
	return _RewardManager.Contract.IsOperator(&_RewardManager.CallOpts, account)
}

// IsRewarder is a free data retrieval call binding the contract method 0x8e478cab.
//
// Solidity: function isRewarder(address account) view returns(bool)
func (_RewardManager *RewardManagerCaller) IsRewarder(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "isRewarder", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRewarder is a free data retrieval call binding the contract method 0x8e478cab.
//
// Solidity: function isRewarder(address account) view returns(bool)
func (_RewardManager *RewardManagerSession) IsRewarder(account common.Address) (bool, error) {
	return _RewardManager.Contract.IsRewarder(&_RewardManager.CallOpts, account)
}

// IsRewarder is a free data retrieval call binding the contract method 0x8e478cab.
//
// Solidity: function isRewarder(address account) view returns(bool)
func (_RewardManager *RewardManagerCallerSession) IsRewarder(account common.Address) (bool, error) {
	return _RewardManager.Contract.IsRewarder(&_RewardManager.CallOpts, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RewardManager *RewardManagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RewardManager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RewardManager *RewardManagerSession) Owner() (common.Address, error) {
	return _RewardManager.Contract.Owner(&_RewardManager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_RewardManager *RewardManagerCallerSession) Owner() (common.Address, error) {
	return _RewardManager.Contract.Owner(&_RewardManager.CallOpts)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address account) returns()
func (_RewardManager *RewardManagerTransactor) AddOperator(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _RewardManager.contract.Transact(opts, "addOperator", account)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address account) returns()
func (_RewardManager *RewardManagerSession) AddOperator(account common.Address) (*types.Transaction, error) {
	return _RewardManager.Contract.AddOperator(&_RewardManager.TransactOpts, account)
}

// AddOperator is a paid mutator transaction binding the contract method 0x9870d7fe.
//
// Solidity: function addOperator(address account) returns()
func (_RewardManager *RewardManagerTransactorSession) AddOperator(account common.Address) (*types.Transaction, error) {
	return _RewardManager.Contract.AddOperator(&_RewardManager.TransactOpts, account)
}

// AddPool is a paid mutator transaction binding the contract method 0xd914cd4b.
//
// Solidity: function addPool(address _poolAddress) returns()
func (_RewardManager *RewardManagerTransactor) AddPool(opts *bind.TransactOpts, _poolAddress common.Address) (*types.Transaction, error) {
	return _RewardManager.contract.Transact(opts, "addPool", _poolAddress)
}

// AddPool is a paid mutator transaction binding the contract method 0xd914cd4b.
//
// Solidity: function addPool(address _poolAddress) returns()
func (_RewardManager *RewardManagerSession) AddPool(_poolAddress common.Address) (*types.Transaction, error) {
	return _RewardManager.Contract.AddPool(&_RewardManager.TransactOpts, _poolAddress)
}

// AddPool is a paid mutator transaction binding the contract method 0xd914cd4b.
//
// Solidity: function addPool(address _poolAddress) returns()
func (_RewardManager *RewardManagerTransactorSession) AddPool(_poolAddress common.Address) (*types.Transaction, error) {
	return _RewardManager.Contract.AddPool(&_RewardManager.TransactOpts, _poolAddress)
}

// AddRewarder is a paid mutator transaction binding the contract method 0x56d3590b.
//
// Solidity: function addRewarder(address account) returns()
func (_RewardManager *RewardManagerTransactor) AddRewarder(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _RewardManager.contract.Transact(opts, "addRewarder", account)
}

// AddRewarder is a paid mutator transaction binding the contract method 0x56d3590b.
//
// Solidity: function addRewarder(address account) returns()
func (_RewardManager *RewardManagerSession) AddRewarder(account common.Address) (*types.Transaction, error) {
	return _RewardManager.Contract.AddRewarder(&_RewardManager.TransactOpts, account)
}

// AddRewarder is a paid mutator transaction binding the contract method 0x56d3590b.
//
// Solidity: function addRewarder(address account) returns()
func (_RewardManager *RewardManagerTransactorSession) AddRewarder(account common.Address) (*types.Transaction, error) {
	return _RewardManager.Contract.AddRewarder(&_RewardManager.TransactOpts, account)
}

// RenounceOperator is a paid mutator transaction binding the contract method 0x2ab6f8db.
//
// Solidity: function renounceOperator() returns()
func (_RewardManager *RewardManagerTransactor) RenounceOperator(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardManager.contract.Transact(opts, "renounceOperator")
}

// RenounceOperator is a paid mutator transaction binding the contract method 0x2ab6f8db.
//
// Solidity: function renounceOperator() returns()
func (_RewardManager *RewardManagerSession) RenounceOperator() (*types.Transaction, error) {
	return _RewardManager.Contract.RenounceOperator(&_RewardManager.TransactOpts)
}

// RenounceOperator is a paid mutator transaction binding the contract method 0x2ab6f8db.
//
// Solidity: function renounceOperator() returns()
func (_RewardManager *RewardManagerTransactorSession) RenounceOperator() (*types.Transaction, error) {
	return _RewardManager.Contract.RenounceOperator(&_RewardManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RewardManager *RewardManagerTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardManager.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RewardManager *RewardManagerSession) RenounceOwnership() (*types.Transaction, error) {
	return _RewardManager.Contract.RenounceOwnership(&_RewardManager.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_RewardManager *RewardManagerTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _RewardManager.Contract.RenounceOwnership(&_RewardManager.TransactOpts)
}

// RenounceRewarder is a paid mutator transaction binding the contract method 0x40cc8518.
//
// Solidity: function renounceRewarder() returns()
func (_RewardManager *RewardManagerTransactor) RenounceRewarder(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RewardManager.contract.Transact(opts, "renounceRewarder")
}

// RenounceRewarder is a paid mutator transaction binding the contract method 0x40cc8518.
//
// Solidity: function renounceRewarder() returns()
func (_RewardManager *RewardManagerSession) RenounceRewarder() (*types.Transaction, error) {
	return _RewardManager.Contract.RenounceRewarder(&_RewardManager.TransactOpts)
}

// RenounceRewarder is a paid mutator transaction binding the contract method 0x40cc8518.
//
// Solidity: function renounceRewarder() returns()
func (_RewardManager *RewardManagerTransactorSession) RenounceRewarder() (*types.Transaction, error) {
	return _RewardManager.Contract.RenounceRewarder(&_RewardManager.TransactOpts)
}

// RewardUser is a paid mutator transaction binding the contract method 0xe4e103dc.
//
// Solidity: function rewardUser(address _user, uint256 _amount) returns()
func (_RewardManager *RewardManagerTransactor) RewardUser(opts *bind.TransactOpts, _user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardManager.contract.Transact(opts, "rewardUser", _user, _amount)
}

// RewardUser is a paid mutator transaction binding the contract method 0xe4e103dc.
//
// Solidity: function rewardUser(address _user, uint256 _amount) returns()
func (_RewardManager *RewardManagerSession) RewardUser(_user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardManager.Contract.RewardUser(&_RewardManager.TransactOpts, _user, _amount)
}

// RewardUser is a paid mutator transaction binding the contract method 0xe4e103dc.
//
// Solidity: function rewardUser(address _user, uint256 _amount) returns()
func (_RewardManager *RewardManagerTransactorSession) RewardUser(_user common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _RewardManager.Contract.RewardUser(&_RewardManager.TransactOpts, _user, _amount)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _newOperator) returns()
func (_RewardManager *RewardManagerTransactor) SetOperator(opts *bind.TransactOpts, _newOperator common.Address) (*types.Transaction, error) {
	return _RewardManager.contract.Transact(opts, "setOperator", _newOperator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _newOperator) returns()
func (_RewardManager *RewardManagerSession) SetOperator(_newOperator common.Address) (*types.Transaction, error) {
	return _RewardManager.Contract.SetOperator(&_RewardManager.TransactOpts, _newOperator)
}

// SetOperator is a paid mutator transaction binding the contract method 0xb3ab15fb.
//
// Solidity: function setOperator(address _newOperator) returns()
func (_RewardManager *RewardManagerTransactorSession) SetOperator(_newOperator common.Address) (*types.Transaction, error) {
	return _RewardManager.Contract.SetOperator(&_RewardManager.TransactOpts, _newOperator)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RewardManager *RewardManagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _RewardManager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RewardManager *RewardManagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RewardManager.Contract.TransferOwnership(&_RewardManager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_RewardManager *RewardManagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _RewardManager.Contract.TransferOwnership(&_RewardManager.TransactOpts, newOwner)
}

// RewardManagerOperatorAddedIterator is returned from FilterOperatorAdded and is used to iterate over the raw logs and unpacked data for OperatorAdded events raised by the RewardManager contract.
type RewardManagerOperatorAddedIterator struct {
	Event *RewardManagerOperatorAdded // Event containing the contract specifics and raw log

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
func (it *RewardManagerOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardManagerOperatorAdded)
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
		it.Event = new(RewardManagerOperatorAdded)
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
func (it *RewardManagerOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardManagerOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardManagerOperatorAdded represents a OperatorAdded event raised by the RewardManager contract.
type RewardManagerOperatorAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOperatorAdded is a free log retrieval operation binding the contract event 0xac6fa858e9350a46cec16539926e0fde25b7629f84b5a72bffaae4df888ae86d.
//
// Solidity: event OperatorAdded(address indexed account)
func (_RewardManager *RewardManagerFilterer) FilterOperatorAdded(opts *bind.FilterOpts, account []common.Address) (*RewardManagerOperatorAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _RewardManager.contract.FilterLogs(opts, "OperatorAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &RewardManagerOperatorAddedIterator{contract: _RewardManager.contract, event: "OperatorAdded", logs: logs, sub: sub}, nil
}

// WatchOperatorAdded is a free log subscription operation binding the contract event 0xac6fa858e9350a46cec16539926e0fde25b7629f84b5a72bffaae4df888ae86d.
//
// Solidity: event OperatorAdded(address indexed account)
func (_RewardManager *RewardManagerFilterer) WatchOperatorAdded(opts *bind.WatchOpts, sink chan<- *RewardManagerOperatorAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _RewardManager.contract.WatchLogs(opts, "OperatorAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardManagerOperatorAdded)
				if err := _RewardManager.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
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
func (_RewardManager *RewardManagerFilterer) ParseOperatorAdded(log types.Log) (*RewardManagerOperatorAdded, error) {
	event := new(RewardManagerOperatorAdded)
	if err := _RewardManager.contract.UnpackLog(event, "OperatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardManagerOperatorRemovedIterator is returned from FilterOperatorRemoved and is used to iterate over the raw logs and unpacked data for OperatorRemoved events raised by the RewardManager contract.
type RewardManagerOperatorRemovedIterator struct {
	Event *RewardManagerOperatorRemoved // Event containing the contract specifics and raw log

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
func (it *RewardManagerOperatorRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardManagerOperatorRemoved)
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
		it.Event = new(RewardManagerOperatorRemoved)
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
func (it *RewardManagerOperatorRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardManagerOperatorRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardManagerOperatorRemoved represents a OperatorRemoved event raised by the RewardManager contract.
type RewardManagerOperatorRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOperatorRemoved is a free log retrieval operation binding the contract event 0x80c0b871b97b595b16a7741c1b06fed0c6f6f558639f18ccbce50724325dc40d.
//
// Solidity: event OperatorRemoved(address indexed account)
func (_RewardManager *RewardManagerFilterer) FilterOperatorRemoved(opts *bind.FilterOpts, account []common.Address) (*RewardManagerOperatorRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _RewardManager.contract.FilterLogs(opts, "OperatorRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &RewardManagerOperatorRemovedIterator{contract: _RewardManager.contract, event: "OperatorRemoved", logs: logs, sub: sub}, nil
}

// WatchOperatorRemoved is a free log subscription operation binding the contract event 0x80c0b871b97b595b16a7741c1b06fed0c6f6f558639f18ccbce50724325dc40d.
//
// Solidity: event OperatorRemoved(address indexed account)
func (_RewardManager *RewardManagerFilterer) WatchOperatorRemoved(opts *bind.WatchOpts, sink chan<- *RewardManagerOperatorRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _RewardManager.contract.WatchLogs(opts, "OperatorRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardManagerOperatorRemoved)
				if err := _RewardManager.contract.UnpackLog(event, "OperatorRemoved", log); err != nil {
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
func (_RewardManager *RewardManagerFilterer) ParseOperatorRemoved(log types.Log) (*RewardManagerOperatorRemoved, error) {
	event := new(RewardManagerOperatorRemoved)
	if err := _RewardManager.contract.UnpackLog(event, "OperatorRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardManagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the RewardManager contract.
type RewardManagerOwnershipTransferredIterator struct {
	Event *RewardManagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RewardManagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardManagerOwnershipTransferred)
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
		it.Event = new(RewardManagerOwnershipTransferred)
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
func (it *RewardManagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardManagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardManagerOwnershipTransferred represents a OwnershipTransferred event raised by the RewardManager contract.
type RewardManagerOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RewardManager *RewardManagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RewardManagerOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RewardManager.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RewardManagerOwnershipTransferredIterator{contract: _RewardManager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RewardManager *RewardManagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RewardManagerOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _RewardManager.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardManagerOwnershipTransferred)
				if err := _RewardManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_RewardManager *RewardManagerFilterer) ParseOwnershipTransferred(log types.Log) (*RewardManagerOwnershipTransferred, error) {
	event := new(RewardManagerOwnershipTransferred)
	if err := _RewardManager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardManagerRewarderAddedIterator is returned from FilterRewarderAdded and is used to iterate over the raw logs and unpacked data for RewarderAdded events raised by the RewardManager contract.
type RewardManagerRewarderAddedIterator struct {
	Event *RewardManagerRewarderAdded // Event containing the contract specifics and raw log

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
func (it *RewardManagerRewarderAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardManagerRewarderAdded)
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
		it.Event = new(RewardManagerRewarderAdded)
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
func (it *RewardManagerRewarderAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardManagerRewarderAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardManagerRewarderAdded represents a RewarderAdded event raised by the RewardManager contract.
type RewardManagerRewarderAdded struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRewarderAdded is a free log retrieval operation binding the contract event 0x9dfd431959d2d3358e3eb909555ad574123ea5881ff0e05a80f66d4984710c1b.
//
// Solidity: event RewarderAdded(address indexed account)
func (_RewardManager *RewardManagerFilterer) FilterRewarderAdded(opts *bind.FilterOpts, account []common.Address) (*RewardManagerRewarderAddedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _RewardManager.contract.FilterLogs(opts, "RewarderAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return &RewardManagerRewarderAddedIterator{contract: _RewardManager.contract, event: "RewarderAdded", logs: logs, sub: sub}, nil
}

// WatchRewarderAdded is a free log subscription operation binding the contract event 0x9dfd431959d2d3358e3eb909555ad574123ea5881ff0e05a80f66d4984710c1b.
//
// Solidity: event RewarderAdded(address indexed account)
func (_RewardManager *RewardManagerFilterer) WatchRewarderAdded(opts *bind.WatchOpts, sink chan<- *RewardManagerRewarderAdded, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _RewardManager.contract.WatchLogs(opts, "RewarderAdded", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardManagerRewarderAdded)
				if err := _RewardManager.contract.UnpackLog(event, "RewarderAdded", log); err != nil {
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
func (_RewardManager *RewardManagerFilterer) ParseRewarderAdded(log types.Log) (*RewardManagerRewarderAdded, error) {
	event := new(RewardManagerRewarderAdded)
	if err := _RewardManager.contract.UnpackLog(event, "RewarderAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardManagerRewarderRemovedIterator is returned from FilterRewarderRemoved and is used to iterate over the raw logs and unpacked data for RewarderRemoved events raised by the RewardManager contract.
type RewardManagerRewarderRemovedIterator struct {
	Event *RewardManagerRewarderRemoved // Event containing the contract specifics and raw log

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
func (it *RewardManagerRewarderRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardManagerRewarderRemoved)
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
		it.Event = new(RewardManagerRewarderRemoved)
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
func (it *RewardManagerRewarderRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardManagerRewarderRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardManagerRewarderRemoved represents a RewarderRemoved event raised by the RewardManager contract.
type RewardManagerRewarderRemoved struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRewarderRemoved is a free log retrieval operation binding the contract event 0xce699c579f0b70ea4ccd6a4b38be26726a2c248b89c7102ccbc5d0f3060ef6d0.
//
// Solidity: event RewarderRemoved(address indexed account)
func (_RewardManager *RewardManagerFilterer) FilterRewarderRemoved(opts *bind.FilterOpts, account []common.Address) (*RewardManagerRewarderRemovedIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _RewardManager.contract.FilterLogs(opts, "RewarderRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return &RewardManagerRewarderRemovedIterator{contract: _RewardManager.contract, event: "RewarderRemoved", logs: logs, sub: sub}, nil
}

// WatchRewarderRemoved is a free log subscription operation binding the contract event 0xce699c579f0b70ea4ccd6a4b38be26726a2c248b89c7102ccbc5d0f3060ef6d0.
//
// Solidity: event RewarderRemoved(address indexed account)
func (_RewardManager *RewardManagerFilterer) WatchRewarderRemoved(opts *bind.WatchOpts, sink chan<- *RewardManagerRewarderRemoved, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _RewardManager.contract.WatchLogs(opts, "RewarderRemoved", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardManagerRewarderRemoved)
				if err := _RewardManager.contract.UnpackLog(event, "RewarderRemoved", log); err != nil {
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
func (_RewardManager *RewardManagerFilterer) ParseRewarderRemoved(log types.Log) (*RewardManagerRewarderRemoved, error) {
	event := new(RewardManagerRewarderRemoved)
	if err := _RewardManager.contract.UnpackLog(event, "RewarderRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardManagerSetOperatorIterator is returned from FilterSetOperator and is used to iterate over the raw logs and unpacked data for SetOperator events raised by the RewardManager contract.
type RewardManagerSetOperatorIterator struct {
	Event *RewardManagerSetOperator // Event containing the contract specifics and raw log

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
func (it *RewardManagerSetOperatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardManagerSetOperator)
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
		it.Event = new(RewardManagerSetOperator)
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
func (it *RewardManagerSetOperatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardManagerSetOperatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardManagerSetOperator represents a SetOperator event raised by the RewardManager contract.
type RewardManagerSetOperator struct {
	Operator common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetOperator is a free log retrieval operation binding the contract event 0xdbebfba65bd6398fb722063efc10c99f624f9cd8ba657201056af918a676d5ee.
//
// Solidity: event SetOperator(address operator)
func (_RewardManager *RewardManagerFilterer) FilterSetOperator(opts *bind.FilterOpts) (*RewardManagerSetOperatorIterator, error) {

	logs, sub, err := _RewardManager.contract.FilterLogs(opts, "SetOperator")
	if err != nil {
		return nil, err
	}
	return &RewardManagerSetOperatorIterator{contract: _RewardManager.contract, event: "SetOperator", logs: logs, sub: sub}, nil
}

// WatchSetOperator is a free log subscription operation binding the contract event 0xdbebfba65bd6398fb722063efc10c99f624f9cd8ba657201056af918a676d5ee.
//
// Solidity: event SetOperator(address operator)
func (_RewardManager *RewardManagerFilterer) WatchSetOperator(opts *bind.WatchOpts, sink chan<- *RewardManagerSetOperator) (event.Subscription, error) {

	logs, sub, err := _RewardManager.contract.WatchLogs(opts, "SetOperator")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardManagerSetOperator)
				if err := _RewardManager.contract.UnpackLog(event, "SetOperator", log); err != nil {
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
func (_RewardManager *RewardManagerFilterer) ParseSetOperator(log types.Log) (*RewardManagerSetOperator, error) {
	event := new(RewardManagerSetOperator)
	if err := _RewardManager.contract.UnpackLog(event, "SetOperator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RewardManagerSetRewarderIterator is returned from FilterSetRewarder and is used to iterate over the raw logs and unpacked data for SetRewarder events raised by the RewardManager contract.
type RewardManagerSetRewarderIterator struct {
	Event *RewardManagerSetRewarder // Event containing the contract specifics and raw log

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
func (it *RewardManagerSetRewarderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RewardManagerSetRewarder)
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
		it.Event = new(RewardManagerSetRewarder)
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
func (it *RewardManagerSetRewarderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RewardManagerSetRewarderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RewardManagerSetRewarder represents a SetRewarder event raised by the RewardManager contract.
type RewardManagerSetRewarder struct {
	Rewarder common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSetRewarder is a free log retrieval operation binding the contract event 0xcf0aff36caea97f7ad632b334936cb196014c193ac1a790b578f12a70d9836db.
//
// Solidity: event SetRewarder(address rewarder)
func (_RewardManager *RewardManagerFilterer) FilterSetRewarder(opts *bind.FilterOpts) (*RewardManagerSetRewarderIterator, error) {

	logs, sub, err := _RewardManager.contract.FilterLogs(opts, "SetRewarder")
	if err != nil {
		return nil, err
	}
	return &RewardManagerSetRewarderIterator{contract: _RewardManager.contract, event: "SetRewarder", logs: logs, sub: sub}, nil
}

// WatchSetRewarder is a free log subscription operation binding the contract event 0xcf0aff36caea97f7ad632b334936cb196014c193ac1a790b578f12a70d9836db.
//
// Solidity: event SetRewarder(address rewarder)
func (_RewardManager *RewardManagerFilterer) WatchSetRewarder(opts *bind.WatchOpts, sink chan<- *RewardManagerSetRewarder) (event.Subscription, error) {

	logs, sub, err := _RewardManager.contract.WatchLogs(opts, "SetRewarder")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RewardManagerSetRewarder)
				if err := _RewardManager.contract.UnpackLog(event, "SetRewarder", log); err != nil {
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
func (_RewardManager *RewardManagerFilterer) ParseSetRewarder(log types.Log) (*RewardManagerSetRewarder, error) {
	event := new(RewardManagerSetRewarder)
	if err := _RewardManager.contract.UnpackLog(event, "SetRewarder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
