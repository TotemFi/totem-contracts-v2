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

// AdvisorsVestingMetaData contains all meta data concerning the AdvisorsVesting contract.
var AdvisorsVestingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractTotemToken\",\"name\":\"_totemToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"}],\"name\":\"StartTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"registeredAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"}],\"name\":\"VestingScheduleRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"registeredAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountWithdrawn\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"LOCK_PERIODS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELEASE_PERIODS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOTAL_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAW_INTERVAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalAmount\",\"type\":\"uint256\"}],\"name\":\"addRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isStartTimeSet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockPeriods\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"recipients\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountWithdrawn\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"releasePeriods\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newStartTime\",\"type\":\"uint256\"}],\"name\":\"setStartTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totemToken\",\"outputs\":[{\"internalType\":\"contractTotemToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unallocatedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"vested\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountVested\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"withdrawable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AdvisorsVestingABI is the input ABI used to generate the binding from.
// Deprecated: Use AdvisorsVestingMetaData.ABI instead.
var AdvisorsVestingABI = AdvisorsVestingMetaData.ABI

// AdvisorsVesting is an auto generated Go binding around an Ethereum contract.
type AdvisorsVesting struct {
	AdvisorsVestingCaller     // Read-only binding to the contract
	AdvisorsVestingTransactor // Write-only binding to the contract
	AdvisorsVestingFilterer   // Log filterer for contract events
}

// AdvisorsVestingCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdvisorsVestingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdvisorsVestingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdvisorsVestingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdvisorsVestingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdvisorsVestingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdvisorsVestingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdvisorsVestingSession struct {
	Contract     *AdvisorsVesting  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdvisorsVestingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdvisorsVestingCallerSession struct {
	Contract *AdvisorsVestingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AdvisorsVestingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdvisorsVestingTransactorSession struct {
	Contract     *AdvisorsVestingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AdvisorsVestingRaw is an auto generated low-level Go binding around an Ethereum contract.
type AdvisorsVestingRaw struct {
	Contract *AdvisorsVesting // Generic contract binding to access the raw methods on
}

// AdvisorsVestingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdvisorsVestingCallerRaw struct {
	Contract *AdvisorsVestingCaller // Generic read-only contract binding to access the raw methods on
}

// AdvisorsVestingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdvisorsVestingTransactorRaw struct {
	Contract *AdvisorsVestingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdvisorsVesting creates a new instance of AdvisorsVesting, bound to a specific deployed contract.
func NewAdvisorsVesting(address common.Address, backend bind.ContractBackend) (*AdvisorsVesting, error) {
	contract, err := bindAdvisorsVesting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AdvisorsVesting{AdvisorsVestingCaller: AdvisorsVestingCaller{contract: contract}, AdvisorsVestingTransactor: AdvisorsVestingTransactor{contract: contract}, AdvisorsVestingFilterer: AdvisorsVestingFilterer{contract: contract}}, nil
}

// NewAdvisorsVestingCaller creates a new read-only instance of AdvisorsVesting, bound to a specific deployed contract.
func NewAdvisorsVestingCaller(address common.Address, caller bind.ContractCaller) (*AdvisorsVestingCaller, error) {
	contract, err := bindAdvisorsVesting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdvisorsVestingCaller{contract: contract}, nil
}

// NewAdvisorsVestingTransactor creates a new write-only instance of AdvisorsVesting, bound to a specific deployed contract.
func NewAdvisorsVestingTransactor(address common.Address, transactor bind.ContractTransactor) (*AdvisorsVestingTransactor, error) {
	contract, err := bindAdvisorsVesting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdvisorsVestingTransactor{contract: contract}, nil
}

// NewAdvisorsVestingFilterer creates a new log filterer instance of AdvisorsVesting, bound to a specific deployed contract.
func NewAdvisorsVestingFilterer(address common.Address, filterer bind.ContractFilterer) (*AdvisorsVestingFilterer, error) {
	contract, err := bindAdvisorsVesting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdvisorsVestingFilterer{contract: contract}, nil
}

// bindAdvisorsVesting binds a generic wrapper to an already deployed contract.
func bindAdvisorsVesting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AdvisorsVestingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdvisorsVesting *AdvisorsVestingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AdvisorsVesting.Contract.AdvisorsVestingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdvisorsVesting *AdvisorsVestingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdvisorsVesting.Contract.AdvisorsVestingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdvisorsVesting *AdvisorsVestingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdvisorsVesting.Contract.AdvisorsVestingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdvisorsVesting *AdvisorsVestingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AdvisorsVesting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdvisorsVesting *AdvisorsVestingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdvisorsVesting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdvisorsVesting *AdvisorsVestingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdvisorsVesting.Contract.contract.Transact(opts, method, params...)
}

// LOCKPERIODS is a free data retrieval call binding the contract method 0x494051a0.
//
// Solidity: function LOCK_PERIODS() view returns(uint256)
func (_AdvisorsVesting *AdvisorsVestingCaller) LOCKPERIODS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdvisorsVesting.contract.Call(opts, &out, "LOCK_PERIODS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LOCKPERIODS is a free data retrieval call binding the contract method 0x494051a0.
//
// Solidity: function LOCK_PERIODS() view returns(uint256)
func (_AdvisorsVesting *AdvisorsVestingSession) LOCKPERIODS() (*big.Int, error) {
	return _AdvisorsVesting.Contract.LOCKPERIODS(&_AdvisorsVesting.CallOpts)
}

// LOCKPERIODS is a free data retrieval call binding the contract method 0x494051a0.
//
// Solidity: function LOCK_PERIODS() view returns(uint256)
func (_AdvisorsVesting *AdvisorsVestingCallerSession) LOCKPERIODS() (*big.Int, error) {
	return _AdvisorsVesting.Contract.LOCKPERIODS(&_AdvisorsVesting.CallOpts)
}

// RELEASEPERIODS is a free data retrieval call binding the contract method 0x00ad3a23.
//
// Solidity: function RELEASE_PERIODS() view returns(uint256)
func (_AdvisorsVesting *AdvisorsVestingCaller) RELEASEPERIODS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdvisorsVesting.contract.Call(opts, &out, "RELEASE_PERIODS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RELEASEPERIODS is a free data retrieval call binding the contract method 0x00ad3a23.
//
// Solidity: function RELEASE_PERIODS() view returns(uint256)
func (_AdvisorsVesting *AdvisorsVestingSession) RELEASEPERIODS() (*big.Int, error) {
	return _AdvisorsVesting.Contract.RELEASEPERIODS(&_AdvisorsVesting.CallOpts)
}

// RELEASEPERIODS is a free data retrieval call binding the contract method 0x00ad3a23.
//
// Solidity: function RELEASE_PERIODS() view returns(uint256)
func (_AdvisorsVesting *AdvisorsVestingCallerSession) RELEASEPERIODS() (*big.Int, error) {
	return _AdvisorsVesting.Contract.RELEASEPERIODS(&_AdvisorsVesting.CallOpts)
}

// TOTALAMOUNT is a free data retrieval call binding the contract method 0xa2d7f5e3.
//
// Solidity: function TOTAL_AMOUNT() view returns(uint256)
func (_AdvisorsVesting *AdvisorsVestingCaller) TOTALAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdvisorsVesting.contract.Call(opts, &out, "TOTAL_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOTALAMOUNT is a free data retrieval call binding the contract method 0xa2d7f5e3.
//
// Solidity: function TOTAL_AMOUNT() view returns(uint256)
func (_AdvisorsVesting *AdvisorsVestingSession) TOTALAMOUNT() (*big.Int, error) {
	return _AdvisorsVesting.Contract.TOTALAMOUNT(&_AdvisorsVesting.CallOpts)
}

// TOTALAMOUNT is a free data retrieval call binding the contract method 0xa2d7f5e3.
//
// Solidity: function TOTAL_AMOUNT() view returns(uint256)
func (_AdvisorsVesting *AdvisorsVestingCallerSession) TOTALAMOUNT() (*big.Int, error) {
	return _AdvisorsVesting.Contract.TOTALAMOUNT(&_AdvisorsVesting.CallOpts)
}

// WITHDRAWINTERVAL is a free data retrieval call binding the contract method 0x30c6b5eb.
//
// Solidity: function WITHDRAW_INTERVAL() view returns(uint256)
func (_AdvisorsVesting *AdvisorsVestingCaller) WITHDRAWINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdvisorsVesting.contract.Call(opts, &out, "WITHDRAW_INTERVAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WITHDRAWINTERVAL is a free data retrieval call binding the contract method 0x30c6b5eb.
//
// Solidity: function WITHDRAW_INTERVAL() view returns(uint256)
func (_AdvisorsVesting *AdvisorsVestingSession) WITHDRAWINTERVAL() (*big.Int, error) {
	return _AdvisorsVesting.Contract.WITHDRAWINTERVAL(&_AdvisorsVesting.CallOpts)
}

// WITHDRAWINTERVAL is a free data retrieval call binding the contract method 0x30c6b5eb.
//
// Solidity: function WITHDRAW_INTERVAL() view returns(uint256)
func (_AdvisorsVesting *AdvisorsVestingCallerSession) WITHDRAWINTERVAL() (*big.Int, error) {
	return _AdvisorsVesting.Contract.WITHDRAWINTERVAL(&_AdvisorsVesting.CallOpts)
}

// IsStartTimeSet is a free data retrieval call binding the contract method 0xf361c59b.
//
// Solidity: function isStartTimeSet() view returns(bool)
func (_AdvisorsVesting *AdvisorsVestingCaller) IsStartTimeSet(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AdvisorsVesting.contract.Call(opts, &out, "isStartTimeSet")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStartTimeSet is a free data retrieval call binding the contract method 0xf361c59b.
//
// Solidity: function isStartTimeSet() view returns(bool)
func (_AdvisorsVesting *AdvisorsVestingSession) IsStartTimeSet() (bool, error) {
	return _AdvisorsVesting.Contract.IsStartTimeSet(&_AdvisorsVesting.CallOpts)
}

// IsStartTimeSet is a free data retrieval call binding the contract method 0xf361c59b.
//
// Solidity: function isStartTimeSet() view returns(bool)
func (_AdvisorsVesting *AdvisorsVestingCallerSession) IsStartTimeSet() (bool, error) {
	return _AdvisorsVesting.Contract.IsStartTimeSet(&_AdvisorsVesting.CallOpts)
}

// LockPeriods is a free data retrieval call binding the contract method 0x2509e086.
//
// Solidity: function lockPeriods() view returns(uint256)
func (_AdvisorsVesting *AdvisorsVestingCaller) LockPeriods(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdvisorsVesting.contract.Call(opts, &out, "lockPeriods")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockPeriods is a free data retrieval call binding the contract method 0x2509e086.
//
// Solidity: function lockPeriods() view returns(uint256)
func (_AdvisorsVesting *AdvisorsVestingSession) LockPeriods() (*big.Int, error) {
	return _AdvisorsVesting.Contract.LockPeriods(&_AdvisorsVesting.CallOpts)
}

// LockPeriods is a free data retrieval call binding the contract method 0x2509e086.
//
// Solidity: function lockPeriods() view returns(uint256)
func (_AdvisorsVesting *AdvisorsVestingCallerSession) LockPeriods() (*big.Int, error) {
	return _AdvisorsVesting.Contract.LockPeriods(&_AdvisorsVesting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AdvisorsVesting *AdvisorsVestingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AdvisorsVesting.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AdvisorsVesting *AdvisorsVestingSession) Owner() (common.Address, error) {
	return _AdvisorsVesting.Contract.Owner(&_AdvisorsVesting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AdvisorsVesting *AdvisorsVestingCallerSession) Owner() (common.Address, error) {
	return _AdvisorsVesting.Contract.Owner(&_AdvisorsVesting.CallOpts)
}

// Recipients is a free data retrieval call binding the contract method 0xeb820312.
//
// Solidity: function recipients(address ) view returns(uint256 totalAmount, uint256 amountWithdrawn)
func (_AdvisorsVesting *AdvisorsVestingCaller) Recipients(opts *bind.CallOpts, arg0 common.Address) (struct {
	TotalAmount     *big.Int
	AmountWithdrawn *big.Int
}, error) {
	var out []interface{}
	err := _AdvisorsVesting.contract.Call(opts, &out, "recipients", arg0)

	outstruct := new(struct {
		TotalAmount     *big.Int
		AmountWithdrawn *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalAmount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AmountWithdrawn = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Recipients is a free data retrieval call binding the contract method 0xeb820312.
//
// Solidity: function recipients(address ) view returns(uint256 totalAmount, uint256 amountWithdrawn)
func (_AdvisorsVesting *AdvisorsVestingSession) Recipients(arg0 common.Address) (struct {
	TotalAmount     *big.Int
	AmountWithdrawn *big.Int
}, error) {
	return _AdvisorsVesting.Contract.Recipients(&_AdvisorsVesting.CallOpts, arg0)
}

// Recipients is a free data retrieval call binding the contract method 0xeb820312.
//
// Solidity: function recipients(address ) view returns(uint256 totalAmount, uint256 amountWithdrawn)
func (_AdvisorsVesting *AdvisorsVestingCallerSession) Recipients(arg0 common.Address) (struct {
	TotalAmount     *big.Int
	AmountWithdrawn *big.Int
}, error) {
	return _AdvisorsVesting.Contract.Recipients(&_AdvisorsVesting.CallOpts, arg0)
}

// ReleasePeriods is a free data retrieval call binding the contract method 0xf57c8cdf.
//
// Solidity: function releasePeriods() view returns(uint256)
func (_AdvisorsVesting *AdvisorsVestingCaller) ReleasePeriods(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdvisorsVesting.contract.Call(opts, &out, "releasePeriods")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReleasePeriods is a free data retrieval call binding the contract method 0xf57c8cdf.
//
// Solidity: function releasePeriods() view returns(uint256)
func (_AdvisorsVesting *AdvisorsVestingSession) ReleasePeriods() (*big.Int, error) {
	return _AdvisorsVesting.Contract.ReleasePeriods(&_AdvisorsVesting.CallOpts)
}

// ReleasePeriods is a free data retrieval call binding the contract method 0xf57c8cdf.
//
// Solidity: function releasePeriods() view returns(uint256)
func (_AdvisorsVesting *AdvisorsVestingCallerSession) ReleasePeriods() (*big.Int, error) {
	return _AdvisorsVesting.Contract.ReleasePeriods(&_AdvisorsVesting.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_AdvisorsVesting *AdvisorsVestingCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdvisorsVesting.contract.Call(opts, &out, "startTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_AdvisorsVesting *AdvisorsVestingSession) StartTime() (*big.Int, error) {
	return _AdvisorsVesting.Contract.StartTime(&_AdvisorsVesting.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_AdvisorsVesting *AdvisorsVestingCallerSession) StartTime() (*big.Int, error) {
	return _AdvisorsVesting.Contract.StartTime(&_AdvisorsVesting.CallOpts)
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_AdvisorsVesting *AdvisorsVestingCaller) TotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdvisorsVesting.contract.Call(opts, &out, "totalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_AdvisorsVesting *AdvisorsVestingSession) TotalAmount() (*big.Int, error) {
	return _AdvisorsVesting.Contract.TotalAmount(&_AdvisorsVesting.CallOpts)
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_AdvisorsVesting *AdvisorsVestingCallerSession) TotalAmount() (*big.Int, error) {
	return _AdvisorsVesting.Contract.TotalAmount(&_AdvisorsVesting.CallOpts)
}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_AdvisorsVesting *AdvisorsVestingCaller) TotemToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AdvisorsVesting.contract.Call(opts, &out, "totemToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_AdvisorsVesting *AdvisorsVestingSession) TotemToken() (common.Address, error) {
	return _AdvisorsVesting.Contract.TotemToken(&_AdvisorsVesting.CallOpts)
}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_AdvisorsVesting *AdvisorsVestingCallerSession) TotemToken() (common.Address, error) {
	return _AdvisorsVesting.Contract.TotemToken(&_AdvisorsVesting.CallOpts)
}

// UnallocatedAmount is a free data retrieval call binding the contract method 0x4afe5bf3.
//
// Solidity: function unallocatedAmount() view returns(uint256)
func (_AdvisorsVesting *AdvisorsVestingCaller) UnallocatedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdvisorsVesting.contract.Call(opts, &out, "unallocatedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnallocatedAmount is a free data retrieval call binding the contract method 0x4afe5bf3.
//
// Solidity: function unallocatedAmount() view returns(uint256)
func (_AdvisorsVesting *AdvisorsVestingSession) UnallocatedAmount() (*big.Int, error) {
	return _AdvisorsVesting.Contract.UnallocatedAmount(&_AdvisorsVesting.CallOpts)
}

// UnallocatedAmount is a free data retrieval call binding the contract method 0x4afe5bf3.
//
// Solidity: function unallocatedAmount() view returns(uint256)
func (_AdvisorsVesting *AdvisorsVestingCallerSession) UnallocatedAmount() (*big.Int, error) {
	return _AdvisorsVesting.Contract.UnallocatedAmount(&_AdvisorsVesting.CallOpts)
}

// Vested is a free data retrieval call binding the contract method 0x7102b728.
//
// Solidity: function vested(address beneficiary) view returns(uint256 _amountVested)
func (_AdvisorsVesting *AdvisorsVestingCaller) Vested(opts *bind.CallOpts, beneficiary common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdvisorsVesting.contract.Call(opts, &out, "vested", beneficiary)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Vested is a free data retrieval call binding the contract method 0x7102b728.
//
// Solidity: function vested(address beneficiary) view returns(uint256 _amountVested)
func (_AdvisorsVesting *AdvisorsVestingSession) Vested(beneficiary common.Address) (*big.Int, error) {
	return _AdvisorsVesting.Contract.Vested(&_AdvisorsVesting.CallOpts, beneficiary)
}

// Vested is a free data retrieval call binding the contract method 0x7102b728.
//
// Solidity: function vested(address beneficiary) view returns(uint256 _amountVested)
func (_AdvisorsVesting *AdvisorsVestingCallerSession) Vested(beneficiary common.Address) (*big.Int, error) {
	return _AdvisorsVesting.Contract.Vested(&_AdvisorsVesting.CallOpts, beneficiary)
}

// WithdrawInterval is a free data retrieval call binding the contract method 0x162075d8.
//
// Solidity: function withdrawInterval() view returns(uint256)
func (_AdvisorsVesting *AdvisorsVestingCaller) WithdrawInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AdvisorsVesting.contract.Call(opts, &out, "withdrawInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawInterval is a free data retrieval call binding the contract method 0x162075d8.
//
// Solidity: function withdrawInterval() view returns(uint256)
func (_AdvisorsVesting *AdvisorsVestingSession) WithdrawInterval() (*big.Int, error) {
	return _AdvisorsVesting.Contract.WithdrawInterval(&_AdvisorsVesting.CallOpts)
}

// WithdrawInterval is a free data retrieval call binding the contract method 0x162075d8.
//
// Solidity: function withdrawInterval() view returns(uint256)
func (_AdvisorsVesting *AdvisorsVestingCallerSession) WithdrawInterval() (*big.Int, error) {
	return _AdvisorsVesting.Contract.WithdrawInterval(&_AdvisorsVesting.CallOpts)
}

// Withdrawable is a free data retrieval call binding the contract method 0xce513b6f.
//
// Solidity: function withdrawable(address beneficiary) view returns(uint256 amount)
func (_AdvisorsVesting *AdvisorsVestingCaller) Withdrawable(opts *bind.CallOpts, beneficiary common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AdvisorsVesting.contract.Call(opts, &out, "withdrawable", beneficiary)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Withdrawable is a free data retrieval call binding the contract method 0xce513b6f.
//
// Solidity: function withdrawable(address beneficiary) view returns(uint256 amount)
func (_AdvisorsVesting *AdvisorsVestingSession) Withdrawable(beneficiary common.Address) (*big.Int, error) {
	return _AdvisorsVesting.Contract.Withdrawable(&_AdvisorsVesting.CallOpts, beneficiary)
}

// Withdrawable is a free data retrieval call binding the contract method 0xce513b6f.
//
// Solidity: function withdrawable(address beneficiary) view returns(uint256 amount)
func (_AdvisorsVesting *AdvisorsVestingCallerSession) Withdrawable(beneficiary common.Address) (*big.Int, error) {
	return _AdvisorsVesting.Contract.Withdrawable(&_AdvisorsVesting.CallOpts, beneficiary)
}

// AddRecipient is a paid mutator transaction binding the contract method 0xf7982243.
//
// Solidity: function addRecipient(address _newRecipient, uint256 _totalAmount) returns()
func (_AdvisorsVesting *AdvisorsVestingTransactor) AddRecipient(opts *bind.TransactOpts, _newRecipient common.Address, _totalAmount *big.Int) (*types.Transaction, error) {
	return _AdvisorsVesting.contract.Transact(opts, "addRecipient", _newRecipient, _totalAmount)
}

// AddRecipient is a paid mutator transaction binding the contract method 0xf7982243.
//
// Solidity: function addRecipient(address _newRecipient, uint256 _totalAmount) returns()
func (_AdvisorsVesting *AdvisorsVestingSession) AddRecipient(_newRecipient common.Address, _totalAmount *big.Int) (*types.Transaction, error) {
	return _AdvisorsVesting.Contract.AddRecipient(&_AdvisorsVesting.TransactOpts, _newRecipient, _totalAmount)
}

// AddRecipient is a paid mutator transaction binding the contract method 0xf7982243.
//
// Solidity: function addRecipient(address _newRecipient, uint256 _totalAmount) returns()
func (_AdvisorsVesting *AdvisorsVestingTransactorSession) AddRecipient(_newRecipient common.Address, _totalAmount *big.Int) (*types.Transaction, error) {
	return _AdvisorsVesting.Contract.AddRecipient(&_AdvisorsVesting.TransactOpts, _newRecipient, _totalAmount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AdvisorsVesting *AdvisorsVestingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdvisorsVesting.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AdvisorsVesting *AdvisorsVestingSession) RenounceOwnership() (*types.Transaction, error) {
	return _AdvisorsVesting.Contract.RenounceOwnership(&_AdvisorsVesting.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AdvisorsVesting *AdvisorsVestingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AdvisorsVesting.Contract.RenounceOwnership(&_AdvisorsVesting.TransactOpts)
}

// SetStartTime is a paid mutator transaction binding the contract method 0x3e0a322d.
//
// Solidity: function setStartTime(uint256 _newStartTime) returns()
func (_AdvisorsVesting *AdvisorsVestingTransactor) SetStartTime(opts *bind.TransactOpts, _newStartTime *big.Int) (*types.Transaction, error) {
	return _AdvisorsVesting.contract.Transact(opts, "setStartTime", _newStartTime)
}

// SetStartTime is a paid mutator transaction binding the contract method 0x3e0a322d.
//
// Solidity: function setStartTime(uint256 _newStartTime) returns()
func (_AdvisorsVesting *AdvisorsVestingSession) SetStartTime(_newStartTime *big.Int) (*types.Transaction, error) {
	return _AdvisorsVesting.Contract.SetStartTime(&_AdvisorsVesting.TransactOpts, _newStartTime)
}

// SetStartTime is a paid mutator transaction binding the contract method 0x3e0a322d.
//
// Solidity: function setStartTime(uint256 _newStartTime) returns()
func (_AdvisorsVesting *AdvisorsVestingTransactorSession) SetStartTime(_newStartTime *big.Int) (*types.Transaction, error) {
	return _AdvisorsVesting.Contract.SetStartTime(&_AdvisorsVesting.TransactOpts, _newStartTime)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AdvisorsVesting *AdvisorsVestingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AdvisorsVesting.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AdvisorsVesting *AdvisorsVestingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AdvisorsVesting.Contract.TransferOwnership(&_AdvisorsVesting.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AdvisorsVesting *AdvisorsVestingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AdvisorsVesting.Contract.TransferOwnership(&_AdvisorsVesting.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_AdvisorsVesting *AdvisorsVestingTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdvisorsVesting.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_AdvisorsVesting *AdvisorsVestingSession) Withdraw() (*types.Transaction, error) {
	return _AdvisorsVesting.Contract.Withdraw(&_AdvisorsVesting.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_AdvisorsVesting *AdvisorsVestingTransactorSession) Withdraw() (*types.Transaction, error) {
	return _AdvisorsVesting.Contract.Withdraw(&_AdvisorsVesting.TransactOpts)
}

// AdvisorsVestingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AdvisorsVesting contract.
type AdvisorsVestingOwnershipTransferredIterator struct {
	Event *AdvisorsVestingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AdvisorsVestingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvisorsVestingOwnershipTransferred)
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
		it.Event = new(AdvisorsVestingOwnershipTransferred)
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
func (it *AdvisorsVestingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdvisorsVestingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdvisorsVestingOwnershipTransferred represents a OwnershipTransferred event raised by the AdvisorsVesting contract.
type AdvisorsVestingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AdvisorsVesting *AdvisorsVestingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AdvisorsVestingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AdvisorsVesting.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AdvisorsVestingOwnershipTransferredIterator{contract: _AdvisorsVesting.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AdvisorsVesting *AdvisorsVestingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AdvisorsVestingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AdvisorsVesting.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdvisorsVestingOwnershipTransferred)
				if err := _AdvisorsVesting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AdvisorsVesting *AdvisorsVestingFilterer) ParseOwnershipTransferred(log types.Log) (*AdvisorsVestingOwnershipTransferred, error) {
	event := new(AdvisorsVestingOwnershipTransferred)
	if err := _AdvisorsVesting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdvisorsVestingStartTimeSetIterator is returned from FilterStartTimeSet and is used to iterate over the raw logs and unpacked data for StartTimeSet events raised by the AdvisorsVesting contract.
type AdvisorsVestingStartTimeSetIterator struct {
	Event *AdvisorsVestingStartTimeSet // Event containing the contract specifics and raw log

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
func (it *AdvisorsVestingStartTimeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvisorsVestingStartTimeSet)
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
		it.Event = new(AdvisorsVestingStartTimeSet)
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
func (it *AdvisorsVestingStartTimeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdvisorsVestingStartTimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdvisorsVestingStartTimeSet represents a StartTimeSet event raised by the AdvisorsVesting contract.
type AdvisorsVestingStartTimeSet struct {
	StartTime *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStartTimeSet is a free log retrieval operation binding the contract event 0xaad53c4362ef2fe5a5390cc046e71fd8423a0a8dceebc156ac9bbcd15997eec2.
//
// Solidity: event StartTimeSet(uint256 startTime)
func (_AdvisorsVesting *AdvisorsVestingFilterer) FilterStartTimeSet(opts *bind.FilterOpts) (*AdvisorsVestingStartTimeSetIterator, error) {

	logs, sub, err := _AdvisorsVesting.contract.FilterLogs(opts, "StartTimeSet")
	if err != nil {
		return nil, err
	}
	return &AdvisorsVestingStartTimeSetIterator{contract: _AdvisorsVesting.contract, event: "StartTimeSet", logs: logs, sub: sub}, nil
}

// WatchStartTimeSet is a free log subscription operation binding the contract event 0xaad53c4362ef2fe5a5390cc046e71fd8423a0a8dceebc156ac9bbcd15997eec2.
//
// Solidity: event StartTimeSet(uint256 startTime)
func (_AdvisorsVesting *AdvisorsVestingFilterer) WatchStartTimeSet(opts *bind.WatchOpts, sink chan<- *AdvisorsVestingStartTimeSet) (event.Subscription, error) {

	logs, sub, err := _AdvisorsVesting.contract.WatchLogs(opts, "StartTimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdvisorsVestingStartTimeSet)
				if err := _AdvisorsVesting.contract.UnpackLog(event, "StartTimeSet", log); err != nil {
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

// ParseStartTimeSet is a log parse operation binding the contract event 0xaad53c4362ef2fe5a5390cc046e71fd8423a0a8dceebc156ac9bbcd15997eec2.
//
// Solidity: event StartTimeSet(uint256 startTime)
func (_AdvisorsVesting *AdvisorsVestingFilterer) ParseStartTimeSet(log types.Log) (*AdvisorsVestingStartTimeSet, error) {
	event := new(AdvisorsVestingStartTimeSet)
	if err := _AdvisorsVesting.contract.UnpackLog(event, "StartTimeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdvisorsVestingVestingScheduleRegisteredIterator is returned from FilterVestingScheduleRegistered and is used to iterate over the raw logs and unpacked data for VestingScheduleRegistered events raised by the AdvisorsVesting contract.
type AdvisorsVestingVestingScheduleRegisteredIterator struct {
	Event *AdvisorsVestingVestingScheduleRegistered // Event containing the contract specifics and raw log

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
func (it *AdvisorsVestingVestingScheduleRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvisorsVestingVestingScheduleRegistered)
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
		it.Event = new(AdvisorsVestingVestingScheduleRegistered)
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
func (it *AdvisorsVestingVestingScheduleRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdvisorsVestingVestingScheduleRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdvisorsVestingVestingScheduleRegistered represents a VestingScheduleRegistered event raised by the AdvisorsVesting contract.
type AdvisorsVestingVestingScheduleRegistered struct {
	RegisteredAddress common.Address
	TotalAmount       *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterVestingScheduleRegistered is a free log retrieval operation binding the contract event 0x16e42921daee38136dc17f8420c371163ec8299e5299fb480d7aeff85cfac3eb.
//
// Solidity: event VestingScheduleRegistered(address registeredAddress, uint256 totalAmount)
func (_AdvisorsVesting *AdvisorsVestingFilterer) FilterVestingScheduleRegistered(opts *bind.FilterOpts) (*AdvisorsVestingVestingScheduleRegisteredIterator, error) {

	logs, sub, err := _AdvisorsVesting.contract.FilterLogs(opts, "VestingScheduleRegistered")
	if err != nil {
		return nil, err
	}
	return &AdvisorsVestingVestingScheduleRegisteredIterator{contract: _AdvisorsVesting.contract, event: "VestingScheduleRegistered", logs: logs, sub: sub}, nil
}

// WatchVestingScheduleRegistered is a free log subscription operation binding the contract event 0x16e42921daee38136dc17f8420c371163ec8299e5299fb480d7aeff85cfac3eb.
//
// Solidity: event VestingScheduleRegistered(address registeredAddress, uint256 totalAmount)
func (_AdvisorsVesting *AdvisorsVestingFilterer) WatchVestingScheduleRegistered(opts *bind.WatchOpts, sink chan<- *AdvisorsVestingVestingScheduleRegistered) (event.Subscription, error) {

	logs, sub, err := _AdvisorsVesting.contract.WatchLogs(opts, "VestingScheduleRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdvisorsVestingVestingScheduleRegistered)
				if err := _AdvisorsVesting.contract.UnpackLog(event, "VestingScheduleRegistered", log); err != nil {
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

// ParseVestingScheduleRegistered is a log parse operation binding the contract event 0x16e42921daee38136dc17f8420c371163ec8299e5299fb480d7aeff85cfac3eb.
//
// Solidity: event VestingScheduleRegistered(address registeredAddress, uint256 totalAmount)
func (_AdvisorsVesting *AdvisorsVestingFilterer) ParseVestingScheduleRegistered(log types.Log) (*AdvisorsVestingVestingScheduleRegistered, error) {
	event := new(AdvisorsVestingVestingScheduleRegistered)
	if err := _AdvisorsVesting.contract.UnpackLog(event, "VestingScheduleRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AdvisorsVestingWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the AdvisorsVesting contract.
type AdvisorsVestingWithdrawIterator struct {
	Event *AdvisorsVestingWithdraw // Event containing the contract specifics and raw log

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
func (it *AdvisorsVestingWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvisorsVestingWithdraw)
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
		it.Event = new(AdvisorsVestingWithdraw)
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
func (it *AdvisorsVestingWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdvisorsVestingWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdvisorsVestingWithdraw represents a Withdraw event raised by the AdvisorsVesting contract.
type AdvisorsVestingWithdraw struct {
	RegisteredAddress common.Address
	AmountWithdrawn   *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address registeredAddress, uint256 amountWithdrawn)
func (_AdvisorsVesting *AdvisorsVestingFilterer) FilterWithdraw(opts *bind.FilterOpts) (*AdvisorsVestingWithdrawIterator, error) {

	logs, sub, err := _AdvisorsVesting.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &AdvisorsVestingWithdrawIterator{contract: _AdvisorsVesting.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address registeredAddress, uint256 amountWithdrawn)
func (_AdvisorsVesting *AdvisorsVestingFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *AdvisorsVestingWithdraw) (event.Subscription, error) {

	logs, sub, err := _AdvisorsVesting.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdvisorsVestingWithdraw)
				if err := _AdvisorsVesting.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address registeredAddress, uint256 amountWithdrawn)
func (_AdvisorsVesting *AdvisorsVestingFilterer) ParseWithdraw(log types.Log) (*AdvisorsVestingWithdraw, error) {
	event := new(AdvisorsVestingWithdraw)
	if err := _AdvisorsVesting.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
