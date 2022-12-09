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

// TeamVestingMetaData contains all meta data concerning the TeamVesting contract.
var TeamVestingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractTotemToken\",\"name\":\"_totemToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"}],\"name\":\"StartTimeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"registeredAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"}],\"name\":\"VestingScheduleRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"registeredAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountWithdrawn\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"LOCK_PERIODS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELEASE_PERIODS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOTAL_AMOUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAW_INTERVAL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_totalAmount\",\"type\":\"uint256\"}],\"name\":\"addRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isStartTimeSet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockPeriods\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"recipients\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountWithdrawn\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"releasePeriods\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newStartTime\",\"type\":\"uint256\"}],\"name\":\"setStartTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totemToken\",\"outputs\":[{\"internalType\":\"contractTotemToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unallocatedAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"vested\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"_amountVested\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawInterval\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"}],\"name\":\"withdrawable\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// TeamVestingABI is the input ABI used to generate the binding from.
// Deprecated: Use TeamVestingMetaData.ABI instead.
var TeamVestingABI = TeamVestingMetaData.ABI

// TeamVesting is an auto generated Go binding around an Ethereum contract.
type TeamVesting struct {
	TeamVestingCaller     // Read-only binding to the contract
	TeamVestingTransactor // Write-only binding to the contract
	TeamVestingFilterer   // Log filterer for contract events
}

// TeamVestingCaller is an auto generated read-only Go binding around an Ethereum contract.
type TeamVestingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeamVestingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TeamVestingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeamVestingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TeamVestingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TeamVestingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TeamVestingSession struct {
	Contract     *TeamVesting      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TeamVestingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TeamVestingCallerSession struct {
	Contract *TeamVestingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// TeamVestingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TeamVestingTransactorSession struct {
	Contract     *TeamVestingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// TeamVestingRaw is an auto generated low-level Go binding around an Ethereum contract.
type TeamVestingRaw struct {
	Contract *TeamVesting // Generic contract binding to access the raw methods on
}

// TeamVestingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TeamVestingCallerRaw struct {
	Contract *TeamVestingCaller // Generic read-only contract binding to access the raw methods on
}

// TeamVestingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TeamVestingTransactorRaw struct {
	Contract *TeamVestingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTeamVesting creates a new instance of TeamVesting, bound to a specific deployed contract.
func NewTeamVesting(address common.Address, backend bind.ContractBackend) (*TeamVesting, error) {
	contract, err := bindTeamVesting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TeamVesting{TeamVestingCaller: TeamVestingCaller{contract: contract}, TeamVestingTransactor: TeamVestingTransactor{contract: contract}, TeamVestingFilterer: TeamVestingFilterer{contract: contract}}, nil
}

// NewTeamVestingCaller creates a new read-only instance of TeamVesting, bound to a specific deployed contract.
func NewTeamVestingCaller(address common.Address, caller bind.ContractCaller) (*TeamVestingCaller, error) {
	contract, err := bindTeamVesting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TeamVestingCaller{contract: contract}, nil
}

// NewTeamVestingTransactor creates a new write-only instance of TeamVesting, bound to a specific deployed contract.
func NewTeamVestingTransactor(address common.Address, transactor bind.ContractTransactor) (*TeamVestingTransactor, error) {
	contract, err := bindTeamVesting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TeamVestingTransactor{contract: contract}, nil
}

// NewTeamVestingFilterer creates a new log filterer instance of TeamVesting, bound to a specific deployed contract.
func NewTeamVestingFilterer(address common.Address, filterer bind.ContractFilterer) (*TeamVestingFilterer, error) {
	contract, err := bindTeamVesting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TeamVestingFilterer{contract: contract}, nil
}

// bindTeamVesting binds a generic wrapper to an already deployed contract.
func bindTeamVesting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TeamVestingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeamVesting *TeamVestingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeamVesting.Contract.TeamVestingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeamVesting *TeamVestingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeamVesting.Contract.TeamVestingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeamVesting *TeamVestingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeamVesting.Contract.TeamVestingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TeamVesting *TeamVestingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TeamVesting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TeamVesting *TeamVestingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeamVesting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TeamVesting *TeamVestingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TeamVesting.Contract.contract.Transact(opts, method, params...)
}

// LOCKPERIODS is a free data retrieval call binding the contract method 0x494051a0.
//
// Solidity: function LOCK_PERIODS() view returns(uint256)
func (_TeamVesting *TeamVestingCaller) LOCKPERIODS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeamVesting.contract.Call(opts, &out, "LOCK_PERIODS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LOCKPERIODS is a free data retrieval call binding the contract method 0x494051a0.
//
// Solidity: function LOCK_PERIODS() view returns(uint256)
func (_TeamVesting *TeamVestingSession) LOCKPERIODS() (*big.Int, error) {
	return _TeamVesting.Contract.LOCKPERIODS(&_TeamVesting.CallOpts)
}

// LOCKPERIODS is a free data retrieval call binding the contract method 0x494051a0.
//
// Solidity: function LOCK_PERIODS() view returns(uint256)
func (_TeamVesting *TeamVestingCallerSession) LOCKPERIODS() (*big.Int, error) {
	return _TeamVesting.Contract.LOCKPERIODS(&_TeamVesting.CallOpts)
}

// RELEASEPERIODS is a free data retrieval call binding the contract method 0x00ad3a23.
//
// Solidity: function RELEASE_PERIODS() view returns(uint256)
func (_TeamVesting *TeamVestingCaller) RELEASEPERIODS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeamVesting.contract.Call(opts, &out, "RELEASE_PERIODS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RELEASEPERIODS is a free data retrieval call binding the contract method 0x00ad3a23.
//
// Solidity: function RELEASE_PERIODS() view returns(uint256)
func (_TeamVesting *TeamVestingSession) RELEASEPERIODS() (*big.Int, error) {
	return _TeamVesting.Contract.RELEASEPERIODS(&_TeamVesting.CallOpts)
}

// RELEASEPERIODS is a free data retrieval call binding the contract method 0x00ad3a23.
//
// Solidity: function RELEASE_PERIODS() view returns(uint256)
func (_TeamVesting *TeamVestingCallerSession) RELEASEPERIODS() (*big.Int, error) {
	return _TeamVesting.Contract.RELEASEPERIODS(&_TeamVesting.CallOpts)
}

// TOTALAMOUNT is a free data retrieval call binding the contract method 0xa2d7f5e3.
//
// Solidity: function TOTAL_AMOUNT() view returns(uint256)
func (_TeamVesting *TeamVestingCaller) TOTALAMOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeamVesting.contract.Call(opts, &out, "TOTAL_AMOUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOTALAMOUNT is a free data retrieval call binding the contract method 0xa2d7f5e3.
//
// Solidity: function TOTAL_AMOUNT() view returns(uint256)
func (_TeamVesting *TeamVestingSession) TOTALAMOUNT() (*big.Int, error) {
	return _TeamVesting.Contract.TOTALAMOUNT(&_TeamVesting.CallOpts)
}

// TOTALAMOUNT is a free data retrieval call binding the contract method 0xa2d7f5e3.
//
// Solidity: function TOTAL_AMOUNT() view returns(uint256)
func (_TeamVesting *TeamVestingCallerSession) TOTALAMOUNT() (*big.Int, error) {
	return _TeamVesting.Contract.TOTALAMOUNT(&_TeamVesting.CallOpts)
}

// WITHDRAWINTERVAL is a free data retrieval call binding the contract method 0x30c6b5eb.
//
// Solidity: function WITHDRAW_INTERVAL() view returns(uint256)
func (_TeamVesting *TeamVestingCaller) WITHDRAWINTERVAL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeamVesting.contract.Call(opts, &out, "WITHDRAW_INTERVAL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WITHDRAWINTERVAL is a free data retrieval call binding the contract method 0x30c6b5eb.
//
// Solidity: function WITHDRAW_INTERVAL() view returns(uint256)
func (_TeamVesting *TeamVestingSession) WITHDRAWINTERVAL() (*big.Int, error) {
	return _TeamVesting.Contract.WITHDRAWINTERVAL(&_TeamVesting.CallOpts)
}

// WITHDRAWINTERVAL is a free data retrieval call binding the contract method 0x30c6b5eb.
//
// Solidity: function WITHDRAW_INTERVAL() view returns(uint256)
func (_TeamVesting *TeamVestingCallerSession) WITHDRAWINTERVAL() (*big.Int, error) {
	return _TeamVesting.Contract.WITHDRAWINTERVAL(&_TeamVesting.CallOpts)
}

// IsStartTimeSet is a free data retrieval call binding the contract method 0xf361c59b.
//
// Solidity: function isStartTimeSet() view returns(bool)
func (_TeamVesting *TeamVestingCaller) IsStartTimeSet(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TeamVesting.contract.Call(opts, &out, "isStartTimeSet")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStartTimeSet is a free data retrieval call binding the contract method 0xf361c59b.
//
// Solidity: function isStartTimeSet() view returns(bool)
func (_TeamVesting *TeamVestingSession) IsStartTimeSet() (bool, error) {
	return _TeamVesting.Contract.IsStartTimeSet(&_TeamVesting.CallOpts)
}

// IsStartTimeSet is a free data retrieval call binding the contract method 0xf361c59b.
//
// Solidity: function isStartTimeSet() view returns(bool)
func (_TeamVesting *TeamVestingCallerSession) IsStartTimeSet() (bool, error) {
	return _TeamVesting.Contract.IsStartTimeSet(&_TeamVesting.CallOpts)
}

// LockPeriods is a free data retrieval call binding the contract method 0x2509e086.
//
// Solidity: function lockPeriods() view returns(uint256)
func (_TeamVesting *TeamVestingCaller) LockPeriods(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeamVesting.contract.Call(opts, &out, "lockPeriods")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockPeriods is a free data retrieval call binding the contract method 0x2509e086.
//
// Solidity: function lockPeriods() view returns(uint256)
func (_TeamVesting *TeamVestingSession) LockPeriods() (*big.Int, error) {
	return _TeamVesting.Contract.LockPeriods(&_TeamVesting.CallOpts)
}

// LockPeriods is a free data retrieval call binding the contract method 0x2509e086.
//
// Solidity: function lockPeriods() view returns(uint256)
func (_TeamVesting *TeamVestingCallerSession) LockPeriods() (*big.Int, error) {
	return _TeamVesting.Contract.LockPeriods(&_TeamVesting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TeamVesting *TeamVestingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeamVesting.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TeamVesting *TeamVestingSession) Owner() (common.Address, error) {
	return _TeamVesting.Contract.Owner(&_TeamVesting.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TeamVesting *TeamVestingCallerSession) Owner() (common.Address, error) {
	return _TeamVesting.Contract.Owner(&_TeamVesting.CallOpts)
}

// Recipients is a free data retrieval call binding the contract method 0xeb820312.
//
// Solidity: function recipients(address ) view returns(uint256 totalAmount, uint256 amountWithdrawn)
func (_TeamVesting *TeamVestingCaller) Recipients(opts *bind.CallOpts, arg0 common.Address) (struct {
	TotalAmount     *big.Int
	AmountWithdrawn *big.Int
}, error) {
	var out []interface{}
	err := _TeamVesting.contract.Call(opts, &out, "recipients", arg0)

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
func (_TeamVesting *TeamVestingSession) Recipients(arg0 common.Address) (struct {
	TotalAmount     *big.Int
	AmountWithdrawn *big.Int
}, error) {
	return _TeamVesting.Contract.Recipients(&_TeamVesting.CallOpts, arg0)
}

// Recipients is a free data retrieval call binding the contract method 0xeb820312.
//
// Solidity: function recipients(address ) view returns(uint256 totalAmount, uint256 amountWithdrawn)
func (_TeamVesting *TeamVestingCallerSession) Recipients(arg0 common.Address) (struct {
	TotalAmount     *big.Int
	AmountWithdrawn *big.Int
}, error) {
	return _TeamVesting.Contract.Recipients(&_TeamVesting.CallOpts, arg0)
}

// ReleasePeriods is a free data retrieval call binding the contract method 0xf57c8cdf.
//
// Solidity: function releasePeriods() view returns(uint256)
func (_TeamVesting *TeamVestingCaller) ReleasePeriods(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeamVesting.contract.Call(opts, &out, "releasePeriods")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReleasePeriods is a free data retrieval call binding the contract method 0xf57c8cdf.
//
// Solidity: function releasePeriods() view returns(uint256)
func (_TeamVesting *TeamVestingSession) ReleasePeriods() (*big.Int, error) {
	return _TeamVesting.Contract.ReleasePeriods(&_TeamVesting.CallOpts)
}

// ReleasePeriods is a free data retrieval call binding the contract method 0xf57c8cdf.
//
// Solidity: function releasePeriods() view returns(uint256)
func (_TeamVesting *TeamVestingCallerSession) ReleasePeriods() (*big.Int, error) {
	return _TeamVesting.Contract.ReleasePeriods(&_TeamVesting.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_TeamVesting *TeamVestingCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeamVesting.contract.Call(opts, &out, "startTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_TeamVesting *TeamVestingSession) StartTime() (*big.Int, error) {
	return _TeamVesting.Contract.StartTime(&_TeamVesting.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_TeamVesting *TeamVestingCallerSession) StartTime() (*big.Int, error) {
	return _TeamVesting.Contract.StartTime(&_TeamVesting.CallOpts)
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_TeamVesting *TeamVestingCaller) TotalAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeamVesting.contract.Call(opts, &out, "totalAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_TeamVesting *TeamVestingSession) TotalAmount() (*big.Int, error) {
	return _TeamVesting.Contract.TotalAmount(&_TeamVesting.CallOpts)
}

// TotalAmount is a free data retrieval call binding the contract method 0x1a39d8ef.
//
// Solidity: function totalAmount() view returns(uint256)
func (_TeamVesting *TeamVestingCallerSession) TotalAmount() (*big.Int, error) {
	return _TeamVesting.Contract.TotalAmount(&_TeamVesting.CallOpts)
}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_TeamVesting *TeamVestingCaller) TotemToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TeamVesting.contract.Call(opts, &out, "totemToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_TeamVesting *TeamVestingSession) TotemToken() (common.Address, error) {
	return _TeamVesting.Contract.TotemToken(&_TeamVesting.CallOpts)
}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_TeamVesting *TeamVestingCallerSession) TotemToken() (common.Address, error) {
	return _TeamVesting.Contract.TotemToken(&_TeamVesting.CallOpts)
}

// UnallocatedAmount is a free data retrieval call binding the contract method 0x4afe5bf3.
//
// Solidity: function unallocatedAmount() view returns(uint256)
func (_TeamVesting *TeamVestingCaller) UnallocatedAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeamVesting.contract.Call(opts, &out, "unallocatedAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UnallocatedAmount is a free data retrieval call binding the contract method 0x4afe5bf3.
//
// Solidity: function unallocatedAmount() view returns(uint256)
func (_TeamVesting *TeamVestingSession) UnallocatedAmount() (*big.Int, error) {
	return _TeamVesting.Contract.UnallocatedAmount(&_TeamVesting.CallOpts)
}

// UnallocatedAmount is a free data retrieval call binding the contract method 0x4afe5bf3.
//
// Solidity: function unallocatedAmount() view returns(uint256)
func (_TeamVesting *TeamVestingCallerSession) UnallocatedAmount() (*big.Int, error) {
	return _TeamVesting.Contract.UnallocatedAmount(&_TeamVesting.CallOpts)
}

// Vested is a free data retrieval call binding the contract method 0x7102b728.
//
// Solidity: function vested(address beneficiary) view returns(uint256 _amountVested)
func (_TeamVesting *TeamVestingCaller) Vested(opts *bind.CallOpts, beneficiary common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TeamVesting.contract.Call(opts, &out, "vested", beneficiary)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Vested is a free data retrieval call binding the contract method 0x7102b728.
//
// Solidity: function vested(address beneficiary) view returns(uint256 _amountVested)
func (_TeamVesting *TeamVestingSession) Vested(beneficiary common.Address) (*big.Int, error) {
	return _TeamVesting.Contract.Vested(&_TeamVesting.CallOpts, beneficiary)
}

// Vested is a free data retrieval call binding the contract method 0x7102b728.
//
// Solidity: function vested(address beneficiary) view returns(uint256 _amountVested)
func (_TeamVesting *TeamVestingCallerSession) Vested(beneficiary common.Address) (*big.Int, error) {
	return _TeamVesting.Contract.Vested(&_TeamVesting.CallOpts, beneficiary)
}

// WithdrawInterval is a free data retrieval call binding the contract method 0x162075d8.
//
// Solidity: function withdrawInterval() view returns(uint256)
func (_TeamVesting *TeamVestingCaller) WithdrawInterval(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TeamVesting.contract.Call(opts, &out, "withdrawInterval")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawInterval is a free data retrieval call binding the contract method 0x162075d8.
//
// Solidity: function withdrawInterval() view returns(uint256)
func (_TeamVesting *TeamVestingSession) WithdrawInterval() (*big.Int, error) {
	return _TeamVesting.Contract.WithdrawInterval(&_TeamVesting.CallOpts)
}

// WithdrawInterval is a free data retrieval call binding the contract method 0x162075d8.
//
// Solidity: function withdrawInterval() view returns(uint256)
func (_TeamVesting *TeamVestingCallerSession) WithdrawInterval() (*big.Int, error) {
	return _TeamVesting.Contract.WithdrawInterval(&_TeamVesting.CallOpts)
}

// Withdrawable is a free data retrieval call binding the contract method 0xce513b6f.
//
// Solidity: function withdrawable(address beneficiary) view returns(uint256 amount)
func (_TeamVesting *TeamVestingCaller) Withdrawable(opts *bind.CallOpts, beneficiary common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TeamVesting.contract.Call(opts, &out, "withdrawable", beneficiary)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Withdrawable is a free data retrieval call binding the contract method 0xce513b6f.
//
// Solidity: function withdrawable(address beneficiary) view returns(uint256 amount)
func (_TeamVesting *TeamVestingSession) Withdrawable(beneficiary common.Address) (*big.Int, error) {
	return _TeamVesting.Contract.Withdrawable(&_TeamVesting.CallOpts, beneficiary)
}

// Withdrawable is a free data retrieval call binding the contract method 0xce513b6f.
//
// Solidity: function withdrawable(address beneficiary) view returns(uint256 amount)
func (_TeamVesting *TeamVestingCallerSession) Withdrawable(beneficiary common.Address) (*big.Int, error) {
	return _TeamVesting.Contract.Withdrawable(&_TeamVesting.CallOpts, beneficiary)
}

// AddRecipient is a paid mutator transaction binding the contract method 0xf7982243.
//
// Solidity: function addRecipient(address _newRecipient, uint256 _totalAmount) returns()
func (_TeamVesting *TeamVestingTransactor) AddRecipient(opts *bind.TransactOpts, _newRecipient common.Address, _totalAmount *big.Int) (*types.Transaction, error) {
	return _TeamVesting.contract.Transact(opts, "addRecipient", _newRecipient, _totalAmount)
}

// AddRecipient is a paid mutator transaction binding the contract method 0xf7982243.
//
// Solidity: function addRecipient(address _newRecipient, uint256 _totalAmount) returns()
func (_TeamVesting *TeamVestingSession) AddRecipient(_newRecipient common.Address, _totalAmount *big.Int) (*types.Transaction, error) {
	return _TeamVesting.Contract.AddRecipient(&_TeamVesting.TransactOpts, _newRecipient, _totalAmount)
}

// AddRecipient is a paid mutator transaction binding the contract method 0xf7982243.
//
// Solidity: function addRecipient(address _newRecipient, uint256 _totalAmount) returns()
func (_TeamVesting *TeamVestingTransactorSession) AddRecipient(_newRecipient common.Address, _totalAmount *big.Int) (*types.Transaction, error) {
	return _TeamVesting.Contract.AddRecipient(&_TeamVesting.TransactOpts, _newRecipient, _totalAmount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TeamVesting *TeamVestingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeamVesting.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TeamVesting *TeamVestingSession) RenounceOwnership() (*types.Transaction, error) {
	return _TeamVesting.Contract.RenounceOwnership(&_TeamVesting.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TeamVesting *TeamVestingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TeamVesting.Contract.RenounceOwnership(&_TeamVesting.TransactOpts)
}

// SetStartTime is a paid mutator transaction binding the contract method 0x3e0a322d.
//
// Solidity: function setStartTime(uint256 _newStartTime) returns()
func (_TeamVesting *TeamVestingTransactor) SetStartTime(opts *bind.TransactOpts, _newStartTime *big.Int) (*types.Transaction, error) {
	return _TeamVesting.contract.Transact(opts, "setStartTime", _newStartTime)
}

// SetStartTime is a paid mutator transaction binding the contract method 0x3e0a322d.
//
// Solidity: function setStartTime(uint256 _newStartTime) returns()
func (_TeamVesting *TeamVestingSession) SetStartTime(_newStartTime *big.Int) (*types.Transaction, error) {
	return _TeamVesting.Contract.SetStartTime(&_TeamVesting.TransactOpts, _newStartTime)
}

// SetStartTime is a paid mutator transaction binding the contract method 0x3e0a322d.
//
// Solidity: function setStartTime(uint256 _newStartTime) returns()
func (_TeamVesting *TeamVestingTransactorSession) SetStartTime(_newStartTime *big.Int) (*types.Transaction, error) {
	return _TeamVesting.Contract.SetStartTime(&_TeamVesting.TransactOpts, _newStartTime)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TeamVesting *TeamVestingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TeamVesting.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TeamVesting *TeamVestingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TeamVesting.Contract.TransferOwnership(&_TeamVesting.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TeamVesting *TeamVestingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TeamVesting.Contract.TransferOwnership(&_TeamVesting.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_TeamVesting *TeamVestingTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TeamVesting.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_TeamVesting *TeamVestingSession) Withdraw() (*types.Transaction, error) {
	return _TeamVesting.Contract.Withdraw(&_TeamVesting.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_TeamVesting *TeamVestingTransactorSession) Withdraw() (*types.Transaction, error) {
	return _TeamVesting.Contract.Withdraw(&_TeamVesting.TransactOpts)
}

// TeamVestingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TeamVesting contract.
type TeamVestingOwnershipTransferredIterator struct {
	Event *TeamVestingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TeamVestingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeamVestingOwnershipTransferred)
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
		it.Event = new(TeamVestingOwnershipTransferred)
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
func (it *TeamVestingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeamVestingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeamVestingOwnershipTransferred represents a OwnershipTransferred event raised by the TeamVesting contract.
type TeamVestingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TeamVesting *TeamVestingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TeamVestingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TeamVesting.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TeamVestingOwnershipTransferredIterator{contract: _TeamVesting.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TeamVesting *TeamVestingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TeamVestingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TeamVesting.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeamVestingOwnershipTransferred)
				if err := _TeamVesting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TeamVesting *TeamVestingFilterer) ParseOwnershipTransferred(log types.Log) (*TeamVestingOwnershipTransferred, error) {
	event := new(TeamVestingOwnershipTransferred)
	if err := _TeamVesting.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeamVestingStartTimeSetIterator is returned from FilterStartTimeSet and is used to iterate over the raw logs and unpacked data for StartTimeSet events raised by the TeamVesting contract.
type TeamVestingStartTimeSetIterator struct {
	Event *TeamVestingStartTimeSet // Event containing the contract specifics and raw log

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
func (it *TeamVestingStartTimeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeamVestingStartTimeSet)
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
		it.Event = new(TeamVestingStartTimeSet)
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
func (it *TeamVestingStartTimeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeamVestingStartTimeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeamVestingStartTimeSet represents a StartTimeSet event raised by the TeamVesting contract.
type TeamVestingStartTimeSet struct {
	StartTime *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStartTimeSet is a free log retrieval operation binding the contract event 0xaad53c4362ef2fe5a5390cc046e71fd8423a0a8dceebc156ac9bbcd15997eec2.
//
// Solidity: event StartTimeSet(uint256 startTime)
func (_TeamVesting *TeamVestingFilterer) FilterStartTimeSet(opts *bind.FilterOpts) (*TeamVestingStartTimeSetIterator, error) {

	logs, sub, err := _TeamVesting.contract.FilterLogs(opts, "StartTimeSet")
	if err != nil {
		return nil, err
	}
	return &TeamVestingStartTimeSetIterator{contract: _TeamVesting.contract, event: "StartTimeSet", logs: logs, sub: sub}, nil
}

// WatchStartTimeSet is a free log subscription operation binding the contract event 0xaad53c4362ef2fe5a5390cc046e71fd8423a0a8dceebc156ac9bbcd15997eec2.
//
// Solidity: event StartTimeSet(uint256 startTime)
func (_TeamVesting *TeamVestingFilterer) WatchStartTimeSet(opts *bind.WatchOpts, sink chan<- *TeamVestingStartTimeSet) (event.Subscription, error) {

	logs, sub, err := _TeamVesting.contract.WatchLogs(opts, "StartTimeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeamVestingStartTimeSet)
				if err := _TeamVesting.contract.UnpackLog(event, "StartTimeSet", log); err != nil {
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
func (_TeamVesting *TeamVestingFilterer) ParseStartTimeSet(log types.Log) (*TeamVestingStartTimeSet, error) {
	event := new(TeamVestingStartTimeSet)
	if err := _TeamVesting.contract.UnpackLog(event, "StartTimeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeamVestingVestingScheduleRegisteredIterator is returned from FilterVestingScheduleRegistered and is used to iterate over the raw logs and unpacked data for VestingScheduleRegistered events raised by the TeamVesting contract.
type TeamVestingVestingScheduleRegisteredIterator struct {
	Event *TeamVestingVestingScheduleRegistered // Event containing the contract specifics and raw log

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
func (it *TeamVestingVestingScheduleRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeamVestingVestingScheduleRegistered)
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
		it.Event = new(TeamVestingVestingScheduleRegistered)
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
func (it *TeamVestingVestingScheduleRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeamVestingVestingScheduleRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeamVestingVestingScheduleRegistered represents a VestingScheduleRegistered event raised by the TeamVesting contract.
type TeamVestingVestingScheduleRegistered struct {
	RegisteredAddress common.Address
	TotalAmount       *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterVestingScheduleRegistered is a free log retrieval operation binding the contract event 0x16e42921daee38136dc17f8420c371163ec8299e5299fb480d7aeff85cfac3eb.
//
// Solidity: event VestingScheduleRegistered(address registeredAddress, uint256 totalAmount)
func (_TeamVesting *TeamVestingFilterer) FilterVestingScheduleRegistered(opts *bind.FilterOpts) (*TeamVestingVestingScheduleRegisteredIterator, error) {

	logs, sub, err := _TeamVesting.contract.FilterLogs(opts, "VestingScheduleRegistered")
	if err != nil {
		return nil, err
	}
	return &TeamVestingVestingScheduleRegisteredIterator{contract: _TeamVesting.contract, event: "VestingScheduleRegistered", logs: logs, sub: sub}, nil
}

// WatchVestingScheduleRegistered is a free log subscription operation binding the contract event 0x16e42921daee38136dc17f8420c371163ec8299e5299fb480d7aeff85cfac3eb.
//
// Solidity: event VestingScheduleRegistered(address registeredAddress, uint256 totalAmount)
func (_TeamVesting *TeamVestingFilterer) WatchVestingScheduleRegistered(opts *bind.WatchOpts, sink chan<- *TeamVestingVestingScheduleRegistered) (event.Subscription, error) {

	logs, sub, err := _TeamVesting.contract.WatchLogs(opts, "VestingScheduleRegistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeamVestingVestingScheduleRegistered)
				if err := _TeamVesting.contract.UnpackLog(event, "VestingScheduleRegistered", log); err != nil {
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
func (_TeamVesting *TeamVestingFilterer) ParseVestingScheduleRegistered(log types.Log) (*TeamVestingVestingScheduleRegistered, error) {
	event := new(TeamVestingVestingScheduleRegistered)
	if err := _TeamVesting.contract.UnpackLog(event, "VestingScheduleRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TeamVestingWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the TeamVesting contract.
type TeamVestingWithdrawIterator struct {
	Event *TeamVestingWithdraw // Event containing the contract specifics and raw log

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
func (it *TeamVestingWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TeamVestingWithdraw)
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
		it.Event = new(TeamVestingWithdraw)
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
func (it *TeamVestingWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TeamVestingWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TeamVestingWithdraw represents a Withdraw event raised by the TeamVesting contract.
type TeamVestingWithdraw struct {
	RegisteredAddress common.Address
	AmountWithdrawn   *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address registeredAddress, uint256 amountWithdrawn)
func (_TeamVesting *TeamVestingFilterer) FilterWithdraw(opts *bind.FilterOpts) (*TeamVestingWithdrawIterator, error) {

	logs, sub, err := _TeamVesting.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &TeamVestingWithdrawIterator{contract: _TeamVesting.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address registeredAddress, uint256 amountWithdrawn)
func (_TeamVesting *TeamVestingFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *TeamVestingWithdraw) (event.Subscription, error) {

	logs, sub, err := _TeamVesting.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TeamVestingWithdraw)
				if err := _TeamVesting.contract.UnpackLog(event, "Withdraw", log); err != nil {
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
func (_TeamVesting *TeamVestingFilterer) ParseWithdraw(log types.Log) (*TeamVestingWithdraw, error) {
	event := new(TeamVestingWithdraw)
	if err := _TeamVesting.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
