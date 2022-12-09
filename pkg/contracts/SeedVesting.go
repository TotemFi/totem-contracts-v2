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

// SeedVestingMetaData contains all meta data concerning the SeedVesting contract.
var SeedVestingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractTotemToken\",\"name\":\"_totemToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"}],\"name\":\"StartTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"registeredAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"}],\"name\":\"VestingScheduleRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"registeredAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountWithdrawn\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"LOCK_PERIODS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELEASE_PERIODS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOTAL_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAW_INTERVAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalAmount\",\"type\":\"uint256\"}],\"name\":\"addRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isStartTimeSet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockPeriods\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"recipients\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountWithdrawn\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"releasePeriods\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newStartTime\",\"type\":\"uint256\"}],\"name\":\"setStartTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totemToken\",\"outputs\":[{\"internalType\":\"contractTotemToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unallocatedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"vested\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountVested\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"withdrawable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// SeedVestingABI is the input ABI used to generate the binding from.
// Deprecated: Use SeedVestingMetaData.ABI instead.
var SeedVestingABI = SeedVestingMetaData.ABI

// SeedVesting is an auto generated Go binding around an Ethereum contract.
type SeedVesting struct {
	SeedVestingCaller     // Read-only binding to the contract
	SeedVestingTransactor // Write-only binding to the contract
	SeedVestingFilterer   // Log filterer for contract events
}

// SeedVestingCaller is an auto generated read-only Go binding around an Ethereum contract.
type SeedVestingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeedVestingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SeedVestingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeedVestingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SeedVestingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SeedVestingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SeedVestingSession struct {
	Contract     *SeedVesting      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SeedVestingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SeedVestingCallerSession struct {
	Contract *SeedVestingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SeedVestingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SeedVestingTransactorSession struct {
	Contract     *SeedVestingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SeedVestingRaw is an auto generated low-level Go binding around an Ethereum contract.
type SeedVestingRaw struct {
	Contract *SeedVesting // Generic contract binding to access the raw methods on
}

// SeedVestingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SeedVestingCallerRaw struct {
	Contract *SeedVestingCaller // Generic read-only contract binding to access the raw methods on
}

// SeedVestingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SeedVestingTransactorRaw struct {
	Contract *SeedVestingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSeedVesting creates a new instance of SeedVesting, bound to a specific deployed contract.
func NewSeedVesting(address common.Address, backend bind.ContractBackend) (*SeedVesting, error) {
	contract, err := bindSeedVesting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SeedVesting{SeedVestingCaller: SeedVestingCaller{contract: contract}, SeedVestingTransactor: SeedVestingTransactor{contract: contract}, SeedVestingFilterer: SeedVestingFilterer{contract: contract}}, nil
}

// NewSeedVestingCaller creates a new read-only instance of SeedVesting, bound to a specific deployed contract.
func NewSeedVestingCaller(address common.Address, caller bind.ContractCaller) (*SeedVestingCaller, error) {
	contract, err := bindSeedVesting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SeedVestingCaller{contract: contract}, nil
}

// NewSeedVestingTransactor creates a new write-only instance of SeedVesting, bound to a specific deployed contract.
func NewSeedVestingTransactor(address common.Address, transactor bind.ContractTransactor) (*SeedVestingTransactor, error) {
	contract, err := bindSeedVesting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SeedVestingTransactor{contract: contract}, nil
}

// NewSeedVestingFilterer creates a new log filterer instance of SeedVesting, bound to a specific deployed contract.
func NewSeedVestingFilterer(address common.Address, filterer bind.ContractFilterer) (*SeedVestingFilterer, error) {
	contract, err := bindSeedVesting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SeedVestingFilterer{contract: contract}, nil
}

// bindSeedVesting binds a generic wrapper to an already deployed contract.
func bindSeedVesting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SeedVestingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SeedVesting *SeedVestingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SeedVesting.Contract.SeedVestingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SeedVesting *SeedVestingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SeedVesting.Contract.SeedVestingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SeedVesting *SeedVestingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SeedVesting.Contract.SeedVestingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SeedVesting *SeedVestingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SeedVesting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SeedVesting *SeedVestingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SeedVesting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SeedVesting *SeedVestingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SeedVesting.Contract.contract.Transact(opts, method, params...)
}

// LOCKPERIODS is a free data retrieval call binding the contract method 0x494051a0.
//
// Solidity: function LOCK_PERIODS() view returns(uint256)
func (_SeedVesting *SeedVestingCaller) LOCKPERIODS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SeedVesting.contract.Call(opts, &out, "LOCK_PERIODS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LOCKPERIODS is a free data retrieval call binding the contract method 0x494051a0.
//
// Solidity: function LOCK_PERIODS() view returns(uint256)
func (_SeedVesting *SeedVestingSession) LOCKPERIODS() (*big.Int, error) {
	return _SeedVesting.Contract.LOCKPERIODS(&_SeedVesting.CallOpts)
}

// LOCKPERIODS is a free data retrieval call binding the contract method 0x494051a0.
//
// Solidity: function LOCK_PERIODS() view returns(uint256)
func (_SeedVesting *SeedVestingCallerSession) LOCKPERIODS() (*big.Int, error) {
	return _SeedVesting.Contract.LOCKPERIODS(&_SeedVesting.CallOpts)
}

// RELEASEPERIODS is a free data retrieval call binding the contract method 0x00ad3a23.
//
// Solidity: function RELEASE_PERIODS() view returns(uint256)
func (_SeedVesting *SeedVestingCaller) RELEASEPERIODS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SeedVesting.contract.Call(opts, &out, "RELEASE_PERIODS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RELEASEPERIODS is a free data retrieval call binding the contract method 0x00ad3a23.
//
// Solidity: function RELEASE_PERIODS() view returns(uint256)
func (_SeedVesting *SeedVestingSession) RELEASEPERIODS() (*big.Int, error) {
	return _SeedVesting.Contract.RELEASEPERIODS(&_SeedVesting.CallOpts)
}

// RELEASEPERIODS is a free data retrieval call binding the contract method 0x00ad3a23.
//
// Solidity: function RELEASE_PERIODS() view returns(uint256)
func (_SeedVesting *SeedVestingCallerSession) RELEASEPERIODS() (*big.Int, error) {
	return _SeedVesting.Contract.RELEASEPERIODS(&_SeedVesting.CallOpts)
}

// TOTALAMOUNT is a free data retrieval call binding the contract method 0xa2d7f5e3.
//
// Solidity: function TOTAL_AMOUNT() view returns(uint256)
func (_SeedVesting *SeedVestingCaller) TOTALAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SeedVesting.contract.Call(opts, &out, "TOTAL_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOTALAMOUNT is a free data retrieval call binding the contract method 0xa2d7f5e3.
//
// Solidity: function TOTAL_AMOUNT() view returns(uint256)
func (_SeedVesting *SeedVestingSession) TOTALAMOUNT() (*big.Int, error) {
	return _SeedVesting.Contract.TOTALAMOUNT(&_SeedVesting.CallOpts)
}

// TOTALAMOUNT is a free data retrieval call binding the contract method 0xa2d7f5e3.
//
// Solidity: function TOTAL_AMOUNT() view returns(uint256)
func (_SeedVesting *SeedVestingCallerSession) TOTALAMOUNT() (*big.Int, error) {
	return _SeedVesting.Contract.TOTALAMOUNT(&_SeedVesting.CallOpts)
}

// WITHDRAWINTERVAL is a free data retrieval call binding the contract method 0x30c6b5eb.
//
// Solidity: function WITHDRAW_INTERVAL() view returns(uint256)
func (_SeedVesting *SeedVestingCaller) WITHDRAWINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SeedVesting.contract.Call(opts, &out, "WITHDRAW_INTERVAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WITHDRAWINTERVAL is a free data retrieval call binding the contract method 0x30c6b5eb.
//
// Solidity: function WITHDRAW_INTERVAL() view returns(uint256)
func (_SeedVesting *SeedVestingSession) WITHDRAWINTERVAL() (*big.Int, error) {
	return _SeedVesting.Contract.WITHDRAWINTERVAL(&_SeedVesting.CallOpts)
}

// WITHDRAWINTERVAL is a free data retrieval call binding the contract method 0x30c6b5eb.
//
// Solidity: function WITHDRAW_INTERVAL() view returns(uint256)
func (_SeedVesting *SeedVestingCallerSession) WITHDRAWINTERVAL() (*big.Int, error) {
	return _SeedVesting.Contract.WITHDRAWINTERVAL(&_SeedVesting.CallOpts)
}

// IsStartTimeSet is a free data retrieval call binding the contract method 0xf361c59b.
//
// Solidity: function isStartTimeSet() view returns(bool)
func (_SeedVesting *SeedVestingCaller) IsStartTimeSet(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SeedVesting.contract.Call(opts, &out, "isStartTimeSet")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStartTimeSet is a free data retrieval call binding the contract method 0xf361c59b.
//
// Solidity: function isStartTimeSet() view returns(bool)
func (_SeedVesting *SeedVestingSession) IsStartTimeSet() (bool, error) {
	return _SeedVesting.Contract.IsStartTimeSet(&_SeedVesting.CallOpts)
}

// IsStartTimeSet is a free data retrieval call binding the contract method 0xf361c59b.
//
// Solidity: function isStartTimeSet() view returns(bool)
func (_SeedVesting *SeedVestingCallerSession) IsStartTimeSet() (bool, error) {
	return _SeedVesting.Contract.IsStartTimeSet(&_SeedVesting.CallOpts)
}

// LockPeriods is a free data retrieval call binding the contract method 0x2509e086.
//
// Solidity: function lockPeriods() view returns(uint256)
func (_SeedVesting *SeedVestingCaller) LockPeriods(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SeedVesting.contract.Call(opts, &out, "lockPeriods")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockPeriods is a free data retrieval call binding the contract method 0x2509e086.
//
// Solidity: function lockPeriods() view returns(uint256)
func (_SeedVesting *SeedVestingSession) LockPeriods() (*big.Int, error) {
	return _SeedVesting.Contract.LockPeriods(&_SeedVesting.CallOpts)
}

// LockPeriods is a free data retrieval call binding the contract method 0x2509e086.
//
// Solidity: function lockPeriods() view returns(uint256)
func (_SeedVesting *SeedVestingCallerSession) LockPeriods() (*big.Int, error) {
	return _SeedVesting.Contract.LockPeriods(&_SeedVesting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SeedVesting *SeedVestingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SeedVesting.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SeedVesting *SeedVestingSession) Owner() (common.Address, error) {
	return _SeedVesting.Contract.Owner(&_SeedVesting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SeedVesting *SeedVestingCallerSession) Owner() (common.Address, error) {
	return _SeedVesting.Contract.Owner(&_SeedVesting.CallOpts)
}

// Recipients is a free data retrieval call binding the contract method 0xeb820312.
//
// Solidity: function recipients(address ) view returns(uint256 totalAmount, uint256 amountWithdrawn)
func (_SeedVesting *SeedVestingCaller) Recipients(opts *bind.CallOpts, arg0 common.Address) (struct {
	TotalAmount     *big.Int
	AmountWithdrawn *big.Int
}, error) {
	var out []interface{}
	err := _SeedVesting.contract.Call(opts, &out, "recipients", arg0)

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
func (_SeedVesting *SeedVestingSession) Recipients(arg0 common.Address) (struct {
	TotalAmount     *big.Int
	AmountWithdrawn *big.Int
}, error) {
	return _SeedVesting.Contract.Recipients(&_SeedVesting.CallOpts, arg0)
}

// Recipients is a free data retrieval call binding the contract method 0xeb820312.
//
// Solidity: function recipients(address ) view returns(uint256 totalAmount, uint256 amountWithdrawn)
func (_SeedVesting *SeedVestingCallerSession) Recipients(arg0 common.Address) (struct {
	TotalAmount     *big.Int
	AmountWithdrawn *big.Int
}, error) {
	return _SeedVesting.Contract.Recipients(&_SeedVesting.CallOpts, arg0)
}

// ReleasePeriods is a free data retrieval call binding the contract method 0xf57c8cdf.
//
// Solidity: function releasePeriods() view returns(uint256)
func (_SeedVesting *SeedVestingCaller) ReleasePeriods(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SeedVesting.contract.Call(opts, &out, "releasePeriods")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReleasePeriods is a free data retrieval call binding the contract method 0xf57c8cdf.
//
// Solidity: function releasePeriods() view returns(uint256)
func (_SeedVesting *SeedVestingSession) ReleasePeriods() (*big.Int, error) {
	return _SeedVesting.Contract.ReleasePeriods(&_SeedVesting.CallOpts)
}

// ReleasePeriods is a free data retrieval call binding the contract method 0xf57c8cdf.
//
// Solidity: function releasePeriods() view returns(uint256)
func (_SeedVesting *SeedVestingCallerSession) ReleasePeriods() (*big.Int, error) {
	return _SeedVesting.Contract.ReleasePeriods(&_SeedVesting.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_SeedVesting *SeedVestingCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SeedVesting.contract.Call(opts, &out, "startTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_SeedVesting *SeedVestingSession) StartTime() (*big.Int, error) {
	return _SeedVesting.Contract.StartTime(&_SeedVesting.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_SeedVesting *SeedVestingCallerSession) StartTime() (*big.Int, error) {
	return _SeedVesting.Contract.StartTime(&_SeedVesting.CallOpts)
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_SeedVesting *SeedVestingCaller) TotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SeedVesting.contract.Call(opts, &out, "totalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_SeedVesting *SeedVestingSession) TotalAmount() (*big.Int, error) {
	return _SeedVesting.Contract.TotalAmount(&_SeedVesting.CallOpts)
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_SeedVesting *SeedVestingCallerSession) TotalAmount() (*big.Int, error) {
	return _SeedVesting.Contract.TotalAmount(&_SeedVesting.CallOpts)
}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_SeedVesting *SeedVestingCaller) TotemToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SeedVesting.contract.Call(opts, &out, "totemToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_SeedVesting *SeedVestingSession) TotemToken() (common.Address, error) {
	return _SeedVesting.Contract.TotemToken(&_SeedVesting.CallOpts)
}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_SeedVesting *SeedVestingCallerSession) TotemToken() (common.Address, error) {
	return _SeedVesting.Contract.TotemToken(&_SeedVesting.CallOpts)
}

// UnallocatedAmount is a free data retrieval call binding the contract method 0x4afe5bf3.
//
// Solidity: function unallocatedAmount() view returns(uint256)
func (_SeedVesting *SeedVestingCaller) UnallocatedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SeedVesting.contract.Call(opts, &out, "unallocatedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnallocatedAmount is a free data retrieval call binding the contract method 0x4afe5bf3.
//
// Solidity: function unallocatedAmount() view returns(uint256)
func (_SeedVesting *SeedVestingSession) UnallocatedAmount() (*big.Int, error) {
	return _SeedVesting.Contract.UnallocatedAmount(&_SeedVesting.CallOpts)
}

// UnallocatedAmount is a free data retrieval call binding the contract method 0x4afe5bf3.
//
// Solidity: function unallocatedAmount() view returns(uint256)
func (_SeedVesting *SeedVestingCallerSession) UnallocatedAmount() (*big.Int, error) {
	return _SeedVesting.Contract.UnallocatedAmount(&_SeedVesting.CallOpts)
}

// Vested is a free data retrieval call binding the contract method 0x7102b728.
//
// Solidity: function vested(address beneficiary) view returns(uint256 _amountVested)
func (_SeedVesting *SeedVestingCaller) Vested(opts *bind.CallOpts, beneficiary common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SeedVesting.contract.Call(opts, &out, "vested", beneficiary)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Vested is a free data retrieval call binding the contract method 0x7102b728.
//
// Solidity: function vested(address beneficiary) view returns(uint256 _amountVested)
func (_SeedVesting *SeedVestingSession) Vested(beneficiary common.Address) (*big.Int, error) {
	return _SeedVesting.Contract.Vested(&_SeedVesting.CallOpts, beneficiary)
}

// Vested is a free data retrieval call binding the contract method 0x7102b728.
//
// Solidity: function vested(address beneficiary) view returns(uint256 _amountVested)
func (_SeedVesting *SeedVestingCallerSession) Vested(beneficiary common.Address) (*big.Int, error) {
	return _SeedVesting.Contract.Vested(&_SeedVesting.CallOpts, beneficiary)
}

// WithdrawInterval is a free data retrieval call binding the contract method 0x162075d8.
//
// Solidity: function withdrawInterval() view returns(uint256)
func (_SeedVesting *SeedVestingCaller) WithdrawInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SeedVesting.contract.Call(opts, &out, "withdrawInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawInterval is a free data retrieval call binding the contract method 0x162075d8.
//
// Solidity: function withdrawInterval() view returns(uint256)
func (_SeedVesting *SeedVestingSession) WithdrawInterval() (*big.Int, error) {
	return _SeedVesting.Contract.WithdrawInterval(&_SeedVesting.CallOpts)
}

// WithdrawInterval is a free data retrieval call binding the contract method 0x162075d8.
//
// Solidity: function withdrawInterval() view returns(uint256)
func (_SeedVesting *SeedVestingCallerSession) WithdrawInterval() (*big.Int, error) {
	return _SeedVesting.Contract.WithdrawInterval(&_SeedVesting.CallOpts)
}

// Withdrawable is a free data retrieval call binding the contract method 0xce513b6f.
//
// Solidity: function withdrawable(address beneficiary) view returns(uint256 amount)
func (_SeedVesting *SeedVestingCaller) Withdrawable(opts *bind.CallOpts, beneficiary common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SeedVesting.contract.Call(opts, &out, "withdrawable", beneficiary)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Withdrawable is a free data retrieval call binding the contract method 0xce513b6f.
//
// Solidity: function withdrawable(address beneficiary) view returns(uint256 amount)
func (_SeedVesting *SeedVestingSession) Withdrawable(beneficiary common.Address) (*big.Int, error) {
	return _SeedVesting.Contract.Withdrawable(&_SeedVesting.CallOpts, beneficiary)
}

// Withdrawable is a free data retrieval call binding the contract method 0xce513b6f.
//
// Solidity: function withdrawable(address beneficiary) view returns(uint256 amount)
func (_SeedVesting *SeedVestingCallerSession) Withdrawable(beneficiary common.Address) (*big.Int, error) {
	return _SeedVesting.Contract.Withdrawable(&_SeedVesting.CallOpts, beneficiary)
}

// AddRecipient is a paid mutator transaction binding the contract method 0xf7982243.
//
// Solidity: function addRecipient(address _newRecipient, uint256 _totalAmount) returns()
func (_SeedVesting *SeedVestingTransactor) AddRecipient(opts *bind.TransactOpts, _newRecipient common.Address, _totalAmount *big.Int) (*types.Transaction, error) {
	return _SeedVesting.contract.Transact(opts, "addRecipient", _newRecipient, _totalAmount)
}

// AddRecipient is a paid mutator transaction binding the contract method 0xf7982243.
//
// Solidity: function addRecipient(address _newRecipient, uint256 _totalAmount) returns()
func (_SeedVesting *SeedVestingSession) AddRecipient(_newRecipient common.Address, _totalAmount *big.Int) (*types.Transaction, error) {
	return _SeedVesting.Contract.AddRecipient(&_SeedVesting.TransactOpts, _newRecipient, _totalAmount)
}

// AddRecipient is a paid mutator transaction binding the contract method 0xf7982243.
//
// Solidity: function addRecipient(address _newRecipient, uint256 _totalAmount) returns()
func (_SeedVesting *SeedVestingTransactorSession) AddRecipient(_newRecipient common.Address, _totalAmount *big.Int) (*types.Transaction, error) {
	return _SeedVesting.Contract.AddRecipient(&_SeedVesting.TransactOpts, _newRecipient, _totalAmount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SeedVesting *SeedVestingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SeedVesting.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SeedVesting *SeedVestingSession) RenounceOwnership() (*types.Transaction, error) {
	return _SeedVesting.Contract.RenounceOwnership(&_SeedVesting.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SeedVesting *SeedVestingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SeedVesting.Contract.RenounceOwnership(&_SeedVesting.TransactOpts)
}

// SetStartTime is a paid mutator transaction binding the contract method 0x3e0a322d.
//
// Solidity: function setStartTime(uint256 _newStartTime) returns()
func (_SeedVesting *SeedVestingTransactor) SetStartTime(opts *bind.TransactOpts, _newStartTime *big.Int) (*types.Transaction, error) {
	return _SeedVesting.contract.Transact(opts, "setStartTime", _newStartTime)
}

// SetStartTime is a paid mutator transaction binding the contract method 0x3e0a322d.
//
// Solidity: function setStartTime(uint256 _newStartTime) returns()
func (_SeedVesting *SeedVestingSession) SetStartTime(_newStartTime *big.Int) (*types.Transaction, error) {
	return _SeedVesting.Contract.SetStartTime(&_SeedVesting.TransactOpts, _newStartTime)
}

// SetStartTime is a paid mutator transaction binding the contract method 0x3e0a322d.
//
// Solidity: function setStartTime(uint256 _newStartTime) returns()
func (_SeedVesting *SeedVestingTransactorSession) SetStartTime(_newStartTime *big.Int) (*types.Transaction, error) {
	return _SeedVesting.Contract.SetStartTime(&_SeedVesting.TransactOpts, _newStartTime)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SeedVesting *SeedVestingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SeedVesting.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SeedVesting *SeedVestingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SeedVesting.Contract.TransferOwnership(&_SeedVesting.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SeedVesting *SeedVestingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SeedVesting.Contract.TransferOwnership(&_SeedVesting.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_SeedVesting *SeedVestingTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SeedVesting.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_SeedVesting *SeedVestingSession) Withdraw() (*types.Transaction, error) {
	return _SeedVesting.Contract.Withdraw(&_SeedVesting.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_SeedVesting *SeedVestingTransactorSession) Withdraw() (*types.Transaction, error) {
	return _SeedVesting.Contract.Withdraw(&_SeedVesting.TransactOpts)
}

// SeedVestingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SeedVesting contract.
type SeedVestingOwnershipTransferredIterator struct {
	Event *SeedVestingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SeedVestingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeedVestingOwnershipTransferred)
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
		it.Event = new(SeedVestingOwnershipTransferred)
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
func (it *SeedVestingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeedVestingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeedVestingOwnershipTransferred represents a OwnershipTransferred event raised by the SeedVesting contract.
type SeedVestingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SeedVesting *SeedVestingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SeedVestingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SeedVesting.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SeedVestingOwnershipTransferredIterator{contract: _SeedVesting.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SeedVesting *SeedVestingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SeedVestingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SeedVesting.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeedVestingOwnershipTransferred)
				if err := _SeedVesting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SeedVesting *SeedVestingFilterer) ParseOwnershipTransferred(log types.Log) (*SeedVestingOwnershipTransferred, error) {
	event := new(SeedVestingOwnershipTransferred)
	if err := _SeedVesting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SeedVestingStartTimeSetIterator is returned from FilterStartTimeSet and is used to iterate over the raw logs and unpacked data for StartTimeSet events raised by the SeedVesting contract.
type SeedVestingStartTimeSetIterator struct {
	Event *SeedVestingStartTimeSet // Event containing the contract specifics and raw log

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
func (it *SeedVestingStartTimeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeedVestingStartTimeSet)
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
		it.Event = new(SeedVestingStartTimeSet)
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
func (it *SeedVestingStartTimeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeedVestingStartTimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeedVestingStartTimeSet represents a StartTimeSet event raised by the SeedVesting contract.
type SeedVestingStartTimeSet struct {
	StartTime *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStartTimeSet is a free log retrieval operation binding the contract event 0xaad53c4362ef2fe5a5390cc046e71fd8423a0a8dceebc156ac9bbcd15997eec2.
//
// Solidity: event StartTimeSet(uint256 startTime)
func (_SeedVesting *SeedVestingFilterer) FilterStartTimeSet(opts *bind.FilterOpts) (*SeedVestingStartTimeSetIterator, error) {

	logs, sub, err := _SeedVesting.contract.FilterLogs(opts, "StartTimeSet")
	if err != nil {
		return nil, err
	}
	return &SeedVestingStartTimeSetIterator{contract: _SeedVesting.contract, event: "StartTimeSet", logs: logs, sub: sub}, nil
}

// WatchStartTimeSet is a free log subscription operation binding the contract event 0xaad53c4362ef2fe5a5390cc046e71fd8423a0a8dceebc156ac9bbcd15997eec2.
//
// Solidity: event StartTimeSet(uint256 startTime)
func (_SeedVesting *SeedVestingFilterer) WatchStartTimeSet(opts *bind.WatchOpts, sink chan<- *SeedVestingStartTimeSet) (event.Subscription, error) {

	logs, sub, err := _SeedVesting.contract.WatchLogs(opts, "StartTimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeedVestingStartTimeSet)
				if err := _SeedVesting.contract.UnpackLog(event, "StartTimeSet", log); err != nil {
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
func (_SeedVesting *SeedVestingFilterer) ParseStartTimeSet(log types.Log) (*SeedVestingStartTimeSet, error) {
	event := new(SeedVestingStartTimeSet)
	if err := _SeedVesting.contract.UnpackLog(event, "StartTimeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SeedVestingVestingScheduleRegisteredIterator is returned from FilterVestingScheduleRegistered and is used to iterate over the raw logs and unpacked data for VestingScheduleRegistered events raised by the SeedVesting contract.
type SeedVestingVestingScheduleRegisteredIterator struct {
	Event *SeedVestingVestingScheduleRegistered // Event containing the contract specifics and raw log

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
func (it *SeedVestingVestingScheduleRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeedVestingVestingScheduleRegistered)
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
		it.Event = new(SeedVestingVestingScheduleRegistered)
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
func (it *SeedVestingVestingScheduleRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeedVestingVestingScheduleRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeedVestingVestingScheduleRegistered represents a VestingScheduleRegistered event raised by the SeedVesting contract.
type SeedVestingVestingScheduleRegistered struct {
	RegisteredAddress common.Address
	TotalAmount       *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterVestingScheduleRegistered is a free log retrieval operation binding the contract event 0x16e42921daee38136dc17f8420c371163ec8299e5299fb480d7aeff85cfac3eb.
//
// Solidity: event VestingScheduleRegistered(address registeredAddress, uint256 totalAmount)
func (_SeedVesting *SeedVestingFilterer) FilterVestingScheduleRegistered(opts *bind.FilterOpts) (*SeedVestingVestingScheduleRegisteredIterator, error) {

	logs, sub, err := _SeedVesting.contract.FilterLogs(opts, "VestingScheduleRegistered")
	if err != nil {
		return nil, err
	}
	return &SeedVestingVestingScheduleRegisteredIterator{contract: _SeedVesting.contract, event: "VestingScheduleRegistered", logs: logs, sub: sub}, nil
}

// WatchVestingScheduleRegistered is a free log subscription operation binding the contract event 0x16e42921daee38136dc17f8420c371163ec8299e5299fb480d7aeff85cfac3eb.
//
// Solidity: event VestingScheduleRegistered(address registeredAddress, uint256 totalAmount)
func (_SeedVesting *SeedVestingFilterer) WatchVestingScheduleRegistered(opts *bind.WatchOpts, sink chan<- *SeedVestingVestingScheduleRegistered) (event.Subscription, error) {

	logs, sub, err := _SeedVesting.contract.WatchLogs(opts, "VestingScheduleRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeedVestingVestingScheduleRegistered)
				if err := _SeedVesting.contract.UnpackLog(event, "VestingScheduleRegistered", log); err != nil {
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
func (_SeedVesting *SeedVestingFilterer) ParseVestingScheduleRegistered(log types.Log) (*SeedVestingVestingScheduleRegistered, error) {
	event := new(SeedVestingVestingScheduleRegistered)
	if err := _SeedVesting.contract.UnpackLog(event, "VestingScheduleRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SeedVestingWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the SeedVesting contract.
type SeedVestingWithdrawIterator struct {
	Event *SeedVestingWithdraw // Event containing the contract specifics and raw log

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
func (it *SeedVestingWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SeedVestingWithdraw)
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
		it.Event = new(SeedVestingWithdraw)
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
func (it *SeedVestingWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SeedVestingWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SeedVestingWithdraw represents a Withdraw event raised by the SeedVesting contract.
type SeedVestingWithdraw struct {
	RegisteredAddress common.Address
	AmountWithdrawn   *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address registeredAddress, uint256 amountWithdrawn)
func (_SeedVesting *SeedVestingFilterer) FilterWithdraw(opts *bind.FilterOpts) (*SeedVestingWithdrawIterator, error) {

	logs, sub, err := _SeedVesting.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &SeedVestingWithdrawIterator{contract: _SeedVesting.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address registeredAddress, uint256 amountWithdrawn)
func (_SeedVesting *SeedVestingFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *SeedVestingWithdraw) (event.Subscription, error) {

	logs, sub, err := _SeedVesting.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SeedVestingWithdraw)
				if err := _SeedVesting.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_SeedVesting *SeedVestingFilterer) ParseWithdraw(log types.Log) (*SeedVestingWithdraw, error) {
	event := new(SeedVestingWithdraw)
	if err := _SeedVesting.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
