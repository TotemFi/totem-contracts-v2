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

// TotemTokenMetaData contains all meta data concerning the TotemToken contract.
var TotemTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"locker\",\"type\":\"address\"}],\"name\":\"SetLocker\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADVISORS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"AdvisorsAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COMMUNITY_DEVELOPMENT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CommunityDevelopmentAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DECIMALS\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIAL_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LAUNCH_POOL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIQUIDITY_POOL\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LiquidityPoolAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PRIVATE_SALE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PUBLIC_SALE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PrivateSaleAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PublicSaleAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SEED_INVESTMENT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_REWARDS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STRATEGIC_ROUND\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SYMBOL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SeedInvestmentAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"StakingRewardsAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"StrategicRoundAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TEAM_ALLOCATION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TeamAllocationAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributeTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"locker\",\"outputs\":[{\"internalType\":\"contractILocker\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_CommunityDevelopmentAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_StakingRewardsAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_LiquidityPoolAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_PublicSaleAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_AdvisorsAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_SeedInvestmentAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_PrivateSaleAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_TeamAllocationAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_StrategicRoundAddr\",\"type\":\"address\"}],\"name\":\"setDistributionTeamsAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_locker\",\"type\":\"address\"}],\"name\":\"setLocker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"setTaxExemptStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTaxRate\",\"type\":\"uint256\"}],\"name\":\"setTaxRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTaxationWallet\",\"type\":\"address\"}],\"name\":\"setTaxationWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"taxExempt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taxRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taxationWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TotemTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use TotemTokenMetaData.ABI instead.
var TotemTokenABI = TotemTokenMetaData.ABI

// TotemToken is an auto generated Go binding around an Ethereum contract.
type TotemToken struct {
	TotemTokenCaller     // Read-only binding to the contract
	TotemTokenTransactor // Write-only binding to the contract
	TotemTokenFilterer   // Log filterer for contract events
}

// TotemTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TotemTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TotemTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TotemTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TotemTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TotemTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TotemTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TotemTokenSession struct {
	Contract     *TotemToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TotemTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TotemTokenCallerSession struct {
	Contract *TotemTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// TotemTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TotemTokenTransactorSession struct {
	Contract     *TotemTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// TotemTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type TotemTokenRaw struct {
	Contract *TotemToken // Generic contract binding to access the raw methods on
}

// TotemTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TotemTokenCallerRaw struct {
	Contract *TotemTokenCaller // Generic read-only contract binding to access the raw methods on
}

// TotemTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TotemTokenTransactorRaw struct {
	Contract *TotemTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTotemToken creates a new instance of TotemToken, bound to a specific deployed contract.
func NewTotemToken(address common.Address, backend bind.ContractBackend) (*TotemToken, error) {
	contract, err := bindTotemToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TotemToken{TotemTokenCaller: TotemTokenCaller{contract: contract}, TotemTokenTransactor: TotemTokenTransactor{contract: contract}, TotemTokenFilterer: TotemTokenFilterer{contract: contract}}, nil
}

// NewTotemTokenCaller creates a new read-only instance of TotemToken, bound to a specific deployed contract.
func NewTotemTokenCaller(address common.Address, caller bind.ContractCaller) (*TotemTokenCaller, error) {
	contract, err := bindTotemToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TotemTokenCaller{contract: contract}, nil
}

// NewTotemTokenTransactor creates a new write-only instance of TotemToken, bound to a specific deployed contract.
func NewTotemTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*TotemTokenTransactor, error) {
	contract, err := bindTotemToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TotemTokenTransactor{contract: contract}, nil
}

// NewTotemTokenFilterer creates a new log filterer instance of TotemToken, bound to a specific deployed contract.
func NewTotemTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*TotemTokenFilterer, error) {
	contract, err := bindTotemToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TotemTokenFilterer{contract: contract}, nil
}

// bindTotemToken binds a generic wrapper to an already deployed contract.
func bindTotemToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TotemTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TotemToken *TotemTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TotemToken.Contract.TotemTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TotemToken *TotemTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TotemToken.Contract.TotemTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TotemToken *TotemTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TotemToken.Contract.TotemTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TotemToken *TotemTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TotemToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TotemToken *TotemTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TotemToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TotemToken *TotemTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TotemToken.Contract.contract.Transact(opts, method, params...)
}

// ADVISORS is a free data retrieval call binding the contract method 0x4ed87713.
//
// Solidity: function ADVISORS() view returns(uint256)
func (_TotemToken *TotemTokenCaller) ADVISORS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TotemToken.contract.Call(opts, &out, "ADVISORS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ADVISORS is a free data retrieval call binding the contract method 0x4ed87713.
//
// Solidity: function ADVISORS() view returns(uint256)
func (_TotemToken *TotemTokenSession) ADVISORS() (*big.Int, error) {
	return _TotemToken.Contract.ADVISORS(&_TotemToken.CallOpts)
}

// ADVISORS is a free data retrieval call binding the contract method 0x4ed87713.
//
// Solidity: function ADVISORS() view returns(uint256)
func (_TotemToken *TotemTokenCallerSession) ADVISORS() (*big.Int, error) {
	return _TotemToken.Contract.ADVISORS(&_TotemToken.CallOpts)
}

// AdvisorsAddr is a free data retrieval call binding the contract method 0xb3b37fe1.
//
// Solidity: function AdvisorsAddr() view returns(address)
func (_TotemToken *TotemTokenCaller) AdvisorsAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TotemToken.contract.Call(opts, &out, "AdvisorsAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdvisorsAddr is a free data retrieval call binding the contract method 0xb3b37fe1.
//
// Solidity: function AdvisorsAddr() view returns(address)
func (_TotemToken *TotemTokenSession) AdvisorsAddr() (common.Address, error) {
	return _TotemToken.Contract.AdvisorsAddr(&_TotemToken.CallOpts)
}

// AdvisorsAddr is a free data retrieval call binding the contract method 0xb3b37fe1.
//
// Solidity: function AdvisorsAddr() view returns(address)
func (_TotemToken *TotemTokenCallerSession) AdvisorsAddr() (common.Address, error) {
	return _TotemToken.Contract.AdvisorsAddr(&_TotemToken.CallOpts)
}

// COMMUNITYDEVELOPMENT is a free data retrieval call binding the contract method 0x02f73b2b.
//
// Solidity: function COMMUNITY_DEVELOPMENT() view returns(uint256)
func (_TotemToken *TotemTokenCaller) COMMUNITYDEVELOPMENT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TotemToken.contract.Call(opts, &out, "COMMUNITY_DEVELOPMENT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COMMUNITYDEVELOPMENT is a free data retrieval call binding the contract method 0x02f73b2b.
//
// Solidity: function COMMUNITY_DEVELOPMENT() view returns(uint256)
func (_TotemToken *TotemTokenSession) COMMUNITYDEVELOPMENT() (*big.Int, error) {
	return _TotemToken.Contract.COMMUNITYDEVELOPMENT(&_TotemToken.CallOpts)
}

// COMMUNITYDEVELOPMENT is a free data retrieval call binding the contract method 0x02f73b2b.
//
// Solidity: function COMMUNITY_DEVELOPMENT() view returns(uint256)
func (_TotemToken *TotemTokenCallerSession) COMMUNITYDEVELOPMENT() (*big.Int, error) {
	return _TotemToken.Contract.COMMUNITYDEVELOPMENT(&_TotemToken.CallOpts)
}

// CommunityDevelopmentAddr is a free data retrieval call binding the contract method 0x337e4ece.
//
// Solidity: function CommunityDevelopmentAddr() view returns(address)
func (_TotemToken *TotemTokenCaller) CommunityDevelopmentAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TotemToken.contract.Call(opts, &out, "CommunityDevelopmentAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CommunityDevelopmentAddr is a free data retrieval call binding the contract method 0x337e4ece.
//
// Solidity: function CommunityDevelopmentAddr() view returns(address)
func (_TotemToken *TotemTokenSession) CommunityDevelopmentAddr() (common.Address, error) {
	return _TotemToken.Contract.CommunityDevelopmentAddr(&_TotemToken.CallOpts)
}

// CommunityDevelopmentAddr is a free data retrieval call binding the contract method 0x337e4ece.
//
// Solidity: function CommunityDevelopmentAddr() view returns(address)
func (_TotemToken *TotemTokenCallerSession) CommunityDevelopmentAddr() (common.Address, error) {
	return _TotemToken.Contract.CommunityDevelopmentAddr(&_TotemToken.CallOpts)
}

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() view returns(uint8)
func (_TotemToken *TotemTokenCaller) DECIMALS(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TotemToken.contract.Call(opts, &out, "DECIMALS")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() view returns(uint8)
func (_TotemToken *TotemTokenSession) DECIMALS() (uint8, error) {
	return _TotemToken.Contract.DECIMALS(&_TotemToken.CallOpts)
}

// DECIMALS is a free data retrieval call binding the contract method 0x2e0f2625.
//
// Solidity: function DECIMALS() view returns(uint8)
func (_TotemToken *TotemTokenCallerSession) DECIMALS() (uint8, error) {
	return _TotemToken.Contract.DECIMALS(&_TotemToken.CallOpts)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() view returns(uint256)
func (_TotemToken *TotemTokenCaller) INITIALSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TotemToken.contract.Call(opts, &out, "INITIAL_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() view returns(uint256)
func (_TotemToken *TotemTokenSession) INITIALSUPPLY() (*big.Int, error) {
	return _TotemToken.Contract.INITIALSUPPLY(&_TotemToken.CallOpts)
}

// INITIALSUPPLY is a free data retrieval call binding the contract method 0x2ff2e9dc.
//
// Solidity: function INITIAL_SUPPLY() view returns(uint256)
func (_TotemToken *TotemTokenCallerSession) INITIALSUPPLY() (*big.Int, error) {
	return _TotemToken.Contract.INITIALSUPPLY(&_TotemToken.CallOpts)
}

// LAUNCHPOOL is a free data retrieval call binding the contract method 0xe04d2a71.
//
// Solidity: function LAUNCH_POOL() view returns(uint256)
func (_TotemToken *TotemTokenCaller) LAUNCHPOOL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TotemToken.contract.Call(opts, &out, "LAUNCH_POOL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LAUNCHPOOL is a free data retrieval call binding the contract method 0xe04d2a71.
//
// Solidity: function LAUNCH_POOL() view returns(uint256)
func (_TotemToken *TotemTokenSession) LAUNCHPOOL() (*big.Int, error) {
	return _TotemToken.Contract.LAUNCHPOOL(&_TotemToken.CallOpts)
}

// LAUNCHPOOL is a free data retrieval call binding the contract method 0xe04d2a71.
//
// Solidity: function LAUNCH_POOL() view returns(uint256)
func (_TotemToken *TotemTokenCallerSession) LAUNCHPOOL() (*big.Int, error) {
	return _TotemToken.Contract.LAUNCHPOOL(&_TotemToken.CallOpts)
}

// LIQUIDITYPOOL is a free data retrieval call binding the contract method 0x2cca9dfd.
//
// Solidity: function LIQUIDITY_POOL() view returns(uint256)
func (_TotemToken *TotemTokenCaller) LIQUIDITYPOOL(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TotemToken.contract.Call(opts, &out, "LIQUIDITY_POOL")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LIQUIDITYPOOL is a free data retrieval call binding the contract method 0x2cca9dfd.
//
// Solidity: function LIQUIDITY_POOL() view returns(uint256)
func (_TotemToken *TotemTokenSession) LIQUIDITYPOOL() (*big.Int, error) {
	return _TotemToken.Contract.LIQUIDITYPOOL(&_TotemToken.CallOpts)
}

// LIQUIDITYPOOL is a free data retrieval call binding the contract method 0x2cca9dfd.
//
// Solidity: function LIQUIDITY_POOL() view returns(uint256)
func (_TotemToken *TotemTokenCallerSession) LIQUIDITYPOOL() (*big.Int, error) {
	return _TotemToken.Contract.LIQUIDITYPOOL(&_TotemToken.CallOpts)
}

// LiquidityPoolAddr is a free data retrieval call binding the contract method 0x124a0b1c.
//
// Solidity: function LiquidityPoolAddr() view returns(address)
func (_TotemToken *TotemTokenCaller) LiquidityPoolAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TotemToken.contract.Call(opts, &out, "LiquidityPoolAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LiquidityPoolAddr is a free data retrieval call binding the contract method 0x124a0b1c.
//
// Solidity: function LiquidityPoolAddr() view returns(address)
func (_TotemToken *TotemTokenSession) LiquidityPoolAddr() (common.Address, error) {
	return _TotemToken.Contract.LiquidityPoolAddr(&_TotemToken.CallOpts)
}

// LiquidityPoolAddr is a free data retrieval call binding the contract method 0x124a0b1c.
//
// Solidity: function LiquidityPoolAddr() view returns(address)
func (_TotemToken *TotemTokenCallerSession) LiquidityPoolAddr() (common.Address, error) {
	return _TotemToken.Contract.LiquidityPoolAddr(&_TotemToken.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_TotemToken *TotemTokenCaller) NAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TotemToken.contract.Call(opts, &out, "NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_TotemToken *TotemTokenSession) NAME() (string, error) {
	return _TotemToken.Contract.NAME(&_TotemToken.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_TotemToken *TotemTokenCallerSession) NAME() (string, error) {
	return _TotemToken.Contract.NAME(&_TotemToken.CallOpts)
}

// PRIVATESALE is a free data retrieval call binding the contract method 0x98c83a16.
//
// Solidity: function PRIVATE_SALE() view returns(uint256)
func (_TotemToken *TotemTokenCaller) PRIVATESALE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TotemToken.contract.Call(opts, &out, "PRIVATE_SALE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PRIVATESALE is a free data retrieval call binding the contract method 0x98c83a16.
//
// Solidity: function PRIVATE_SALE() view returns(uint256)
func (_TotemToken *TotemTokenSession) PRIVATESALE() (*big.Int, error) {
	return _TotemToken.Contract.PRIVATESALE(&_TotemToken.CallOpts)
}

// PRIVATESALE is a free data retrieval call binding the contract method 0x98c83a16.
//
// Solidity: function PRIVATE_SALE() view returns(uint256)
func (_TotemToken *TotemTokenCallerSession) PRIVATESALE() (*big.Int, error) {
	return _TotemToken.Contract.PRIVATESALE(&_TotemToken.CallOpts)
}

// PUBLICSALE is a free data retrieval call binding the contract method 0x5f89584e.
//
// Solidity: function PUBLIC_SALE() view returns(uint256)
func (_TotemToken *TotemTokenCaller) PUBLICSALE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TotemToken.contract.Call(opts, &out, "PUBLIC_SALE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PUBLICSALE is a free data retrieval call binding the contract method 0x5f89584e.
//
// Solidity: function PUBLIC_SALE() view returns(uint256)
func (_TotemToken *TotemTokenSession) PUBLICSALE() (*big.Int, error) {
	return _TotemToken.Contract.PUBLICSALE(&_TotemToken.CallOpts)
}

// PUBLICSALE is a free data retrieval call binding the contract method 0x5f89584e.
//
// Solidity: function PUBLIC_SALE() view returns(uint256)
func (_TotemToken *TotemTokenCallerSession) PUBLICSALE() (*big.Int, error) {
	return _TotemToken.Contract.PUBLICSALE(&_TotemToken.CallOpts)
}

// PrivateSaleAddr is a free data retrieval call binding the contract method 0x98f70100.
//
// Solidity: function PrivateSaleAddr() view returns(address)
func (_TotemToken *TotemTokenCaller) PrivateSaleAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TotemToken.contract.Call(opts, &out, "PrivateSaleAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrivateSaleAddr is a free data retrieval call binding the contract method 0x98f70100.
//
// Solidity: function PrivateSaleAddr() view returns(address)
func (_TotemToken *TotemTokenSession) PrivateSaleAddr() (common.Address, error) {
	return _TotemToken.Contract.PrivateSaleAddr(&_TotemToken.CallOpts)
}

// PrivateSaleAddr is a free data retrieval call binding the contract method 0x98f70100.
//
// Solidity: function PrivateSaleAddr() view returns(address)
func (_TotemToken *TotemTokenCallerSession) PrivateSaleAddr() (common.Address, error) {
	return _TotemToken.Contract.PrivateSaleAddr(&_TotemToken.CallOpts)
}

// PublicSaleAddr is a free data retrieval call binding the contract method 0x93b23dce.
//
// Solidity: function PublicSaleAddr() view returns(address)
func (_TotemToken *TotemTokenCaller) PublicSaleAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TotemToken.contract.Call(opts, &out, "PublicSaleAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PublicSaleAddr is a free data retrieval call binding the contract method 0x93b23dce.
//
// Solidity: function PublicSaleAddr() view returns(address)
func (_TotemToken *TotemTokenSession) PublicSaleAddr() (common.Address, error) {
	return _TotemToken.Contract.PublicSaleAddr(&_TotemToken.CallOpts)
}

// PublicSaleAddr is a free data retrieval call binding the contract method 0x93b23dce.
//
// Solidity: function PublicSaleAddr() view returns(address)
func (_TotemToken *TotemTokenCallerSession) PublicSaleAddr() (common.Address, error) {
	return _TotemToken.Contract.PublicSaleAddr(&_TotemToken.CallOpts)
}

// SEEDINVESTMENT is a free data retrieval call binding the contract method 0x45f38dbb.
//
// Solidity: function SEED_INVESTMENT() view returns(uint256)
func (_TotemToken *TotemTokenCaller) SEEDINVESTMENT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TotemToken.contract.Call(opts, &out, "SEED_INVESTMENT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SEEDINVESTMENT is a free data retrieval call binding the contract method 0x45f38dbb.
//
// Solidity: function SEED_INVESTMENT() view returns(uint256)
func (_TotemToken *TotemTokenSession) SEEDINVESTMENT() (*big.Int, error) {
	return _TotemToken.Contract.SEEDINVESTMENT(&_TotemToken.CallOpts)
}

// SEEDINVESTMENT is a free data retrieval call binding the contract method 0x45f38dbb.
//
// Solidity: function SEED_INVESTMENT() view returns(uint256)
func (_TotemToken *TotemTokenCallerSession) SEEDINVESTMENT() (*big.Int, error) {
	return _TotemToken.Contract.SEEDINVESTMENT(&_TotemToken.CallOpts)
}

// STAKINGREWARDS is a free data retrieval call binding the contract method 0x07bce0cd.
//
// Solidity: function STAKING_REWARDS() view returns(uint256)
func (_TotemToken *TotemTokenCaller) STAKINGREWARDS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TotemToken.contract.Call(opts, &out, "STAKING_REWARDS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STAKINGREWARDS is a free data retrieval call binding the contract method 0x07bce0cd.
//
// Solidity: function STAKING_REWARDS() view returns(uint256)
func (_TotemToken *TotemTokenSession) STAKINGREWARDS() (*big.Int, error) {
	return _TotemToken.Contract.STAKINGREWARDS(&_TotemToken.CallOpts)
}

// STAKINGREWARDS is a free data retrieval call binding the contract method 0x07bce0cd.
//
// Solidity: function STAKING_REWARDS() view returns(uint256)
func (_TotemToken *TotemTokenCallerSession) STAKINGREWARDS() (*big.Int, error) {
	return _TotemToken.Contract.STAKINGREWARDS(&_TotemToken.CallOpts)
}

// STRATEGICROUND is a free data retrieval call binding the contract method 0x842460a5.
//
// Solidity: function STRATEGIC_ROUND() view returns(uint256)
func (_TotemToken *TotemTokenCaller) STRATEGICROUND(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TotemToken.contract.Call(opts, &out, "STRATEGIC_ROUND")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// STRATEGICROUND is a free data retrieval call binding the contract method 0x842460a5.
//
// Solidity: function STRATEGIC_ROUND() view returns(uint256)
func (_TotemToken *TotemTokenSession) STRATEGICROUND() (*big.Int, error) {
	return _TotemToken.Contract.STRATEGICROUND(&_TotemToken.CallOpts)
}

// STRATEGICROUND is a free data retrieval call binding the contract method 0x842460a5.
//
// Solidity: function STRATEGIC_ROUND() view returns(uint256)
func (_TotemToken *TotemTokenCallerSession) STRATEGICROUND() (*big.Int, error) {
	return _TotemToken.Contract.STRATEGICROUND(&_TotemToken.CallOpts)
}

// SYMBOL is a free data retrieval call binding the contract method 0xf76f8d78.
//
// Solidity: function SYMBOL() view returns(string)
func (_TotemToken *TotemTokenCaller) SYMBOL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TotemToken.contract.Call(opts, &out, "SYMBOL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SYMBOL is a free data retrieval call binding the contract method 0xf76f8d78.
//
// Solidity: function SYMBOL() view returns(string)
func (_TotemToken *TotemTokenSession) SYMBOL() (string, error) {
	return _TotemToken.Contract.SYMBOL(&_TotemToken.CallOpts)
}

// SYMBOL is a free data retrieval call binding the contract method 0xf76f8d78.
//
// Solidity: function SYMBOL() view returns(string)
func (_TotemToken *TotemTokenCallerSession) SYMBOL() (string, error) {
	return _TotemToken.Contract.SYMBOL(&_TotemToken.CallOpts)
}

// SeedInvestmentAddr is a free data retrieval call binding the contract method 0x9d51606d.
//
// Solidity: function SeedInvestmentAddr() view returns(address)
func (_TotemToken *TotemTokenCaller) SeedInvestmentAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TotemToken.contract.Call(opts, &out, "SeedInvestmentAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SeedInvestmentAddr is a free data retrieval call binding the contract method 0x9d51606d.
//
// Solidity: function SeedInvestmentAddr() view returns(address)
func (_TotemToken *TotemTokenSession) SeedInvestmentAddr() (common.Address, error) {
	return _TotemToken.Contract.SeedInvestmentAddr(&_TotemToken.CallOpts)
}

// SeedInvestmentAddr is a free data retrieval call binding the contract method 0x9d51606d.
//
// Solidity: function SeedInvestmentAddr() view returns(address)
func (_TotemToken *TotemTokenCallerSession) SeedInvestmentAddr() (common.Address, error) {
	return _TotemToken.Contract.SeedInvestmentAddr(&_TotemToken.CallOpts)
}

// StakingRewardsAddr is a free data retrieval call binding the contract method 0x06f2d223.
//
// Solidity: function StakingRewardsAddr() view returns(address)
func (_TotemToken *TotemTokenCaller) StakingRewardsAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TotemToken.contract.Call(opts, &out, "StakingRewardsAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingRewardsAddr is a free data retrieval call binding the contract method 0x06f2d223.
//
// Solidity: function StakingRewardsAddr() view returns(address)
func (_TotemToken *TotemTokenSession) StakingRewardsAddr() (common.Address, error) {
	return _TotemToken.Contract.StakingRewardsAddr(&_TotemToken.CallOpts)
}

// StakingRewardsAddr is a free data retrieval call binding the contract method 0x06f2d223.
//
// Solidity: function StakingRewardsAddr() view returns(address)
func (_TotemToken *TotemTokenCallerSession) StakingRewardsAddr() (common.Address, error) {
	return _TotemToken.Contract.StakingRewardsAddr(&_TotemToken.CallOpts)
}

// StrategicRoundAddr is a free data retrieval call binding the contract method 0x1f48847b.
//
// Solidity: function StrategicRoundAddr() view returns(address)
func (_TotemToken *TotemTokenCaller) StrategicRoundAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TotemToken.contract.Call(opts, &out, "StrategicRoundAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StrategicRoundAddr is a free data retrieval call binding the contract method 0x1f48847b.
//
// Solidity: function StrategicRoundAddr() view returns(address)
func (_TotemToken *TotemTokenSession) StrategicRoundAddr() (common.Address, error) {
	return _TotemToken.Contract.StrategicRoundAddr(&_TotemToken.CallOpts)
}

// StrategicRoundAddr is a free data retrieval call binding the contract method 0x1f48847b.
//
// Solidity: function StrategicRoundAddr() view returns(address)
func (_TotemToken *TotemTokenCallerSession) StrategicRoundAddr() (common.Address, error) {
	return _TotemToken.Contract.StrategicRoundAddr(&_TotemToken.CallOpts)
}

// TEAMALLOCATION is a free data retrieval call binding the contract method 0xfd99cbed.
//
// Solidity: function TEAM_ALLOCATION() view returns(uint256)
func (_TotemToken *TotemTokenCaller) TEAMALLOCATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TotemToken.contract.Call(opts, &out, "TEAM_ALLOCATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TEAMALLOCATION is a free data retrieval call binding the contract method 0xfd99cbed.
//
// Solidity: function TEAM_ALLOCATION() view returns(uint256)
func (_TotemToken *TotemTokenSession) TEAMALLOCATION() (*big.Int, error) {
	return _TotemToken.Contract.TEAMALLOCATION(&_TotemToken.CallOpts)
}

// TEAMALLOCATION is a free data retrieval call binding the contract method 0xfd99cbed.
//
// Solidity: function TEAM_ALLOCATION() view returns(uint256)
func (_TotemToken *TotemTokenCallerSession) TEAMALLOCATION() (*big.Int, error) {
	return _TotemToken.Contract.TEAMALLOCATION(&_TotemToken.CallOpts)
}

// TeamAllocationAddr is a free data retrieval call binding the contract method 0x29470766.
//
// Solidity: function TeamAllocationAddr() view returns(address)
func (_TotemToken *TotemTokenCaller) TeamAllocationAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TotemToken.contract.Call(opts, &out, "TeamAllocationAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TeamAllocationAddr is a free data retrieval call binding the contract method 0x29470766.
//
// Solidity: function TeamAllocationAddr() view returns(address)
func (_TotemToken *TotemTokenSession) TeamAllocationAddr() (common.Address, error) {
	return _TotemToken.Contract.TeamAllocationAddr(&_TotemToken.CallOpts)
}

// TeamAllocationAddr is a free data retrieval call binding the contract method 0x29470766.
//
// Solidity: function TeamAllocationAddr() view returns(address)
func (_TotemToken *TotemTokenCallerSession) TeamAllocationAddr() (common.Address, error) {
	return _TotemToken.Contract.TeamAllocationAddr(&_TotemToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TotemToken *TotemTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TotemToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TotemToken *TotemTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TotemToken.Contract.Allowance(&_TotemToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TotemToken *TotemTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TotemToken.Contract.Allowance(&_TotemToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TotemToken *TotemTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TotemToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TotemToken *TotemTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _TotemToken.Contract.BalanceOf(&_TotemToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TotemToken *TotemTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _TotemToken.Contract.BalanceOf(&_TotemToken.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TotemToken *TotemTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TotemToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TotemToken *TotemTokenSession) Decimals() (uint8, error) {
	return _TotemToken.Contract.Decimals(&_TotemToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TotemToken *TotemTokenCallerSession) Decimals() (uint8, error) {
	return _TotemToken.Contract.Decimals(&_TotemToken.CallOpts)
}

// Locker is a free data retrieval call binding the contract method 0xd7b96d4e.
//
// Solidity: function locker() view returns(address)
func (_TotemToken *TotemTokenCaller) Locker(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TotemToken.contract.Call(opts, &out, "locker")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Locker is a free data retrieval call binding the contract method 0xd7b96d4e.
//
// Solidity: function locker() view returns(address)
func (_TotemToken *TotemTokenSession) Locker() (common.Address, error) {
	return _TotemToken.Contract.Locker(&_TotemToken.CallOpts)
}

// Locker is a free data retrieval call binding the contract method 0xd7b96d4e.
//
// Solidity: function locker() view returns(address)
func (_TotemToken *TotemTokenCallerSession) Locker() (common.Address, error) {
	return _TotemToken.Contract.Locker(&_TotemToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TotemToken *TotemTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TotemToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TotemToken *TotemTokenSession) Name() (string, error) {
	return _TotemToken.Contract.Name(&_TotemToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TotemToken *TotemTokenCallerSession) Name() (string, error) {
	return _TotemToken.Contract.Name(&_TotemToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TotemToken *TotemTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TotemToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TotemToken *TotemTokenSession) Owner() (common.Address, error) {
	return _TotemToken.Contract.Owner(&_TotemToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TotemToken *TotemTokenCallerSession) Owner() (common.Address, error) {
	return _TotemToken.Contract.Owner(&_TotemToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TotemToken *TotemTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TotemToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TotemToken *TotemTokenSession) Symbol() (string, error) {
	return _TotemToken.Contract.Symbol(&_TotemToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TotemToken *TotemTokenCallerSession) Symbol() (string, error) {
	return _TotemToken.Contract.Symbol(&_TotemToken.CallOpts)
}

// TaxExempt is a free data retrieval call binding the contract method 0xd1ecfc68.
//
// Solidity: function taxExempt(address ) view returns(bool)
func (_TotemToken *TotemTokenCaller) TaxExempt(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _TotemToken.contract.Call(opts, &out, "taxExempt", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TaxExempt is a free data retrieval call binding the contract method 0xd1ecfc68.
//
// Solidity: function taxExempt(address ) view returns(bool)
func (_TotemToken *TotemTokenSession) TaxExempt(arg0 common.Address) (bool, error) {
	return _TotemToken.Contract.TaxExempt(&_TotemToken.CallOpts, arg0)
}

// TaxExempt is a free data retrieval call binding the contract method 0xd1ecfc68.
//
// Solidity: function taxExempt(address ) view returns(bool)
func (_TotemToken *TotemTokenCallerSession) TaxExempt(arg0 common.Address) (bool, error) {
	return _TotemToken.Contract.TaxExempt(&_TotemToken.CallOpts, arg0)
}

// TaxRate is a free data retrieval call binding the contract method 0x771a3a1d.
//
// Solidity: function taxRate() view returns(uint256)
func (_TotemToken *TotemTokenCaller) TaxRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TotemToken.contract.Call(opts, &out, "taxRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TaxRate is a free data retrieval call binding the contract method 0x771a3a1d.
//
// Solidity: function taxRate() view returns(uint256)
func (_TotemToken *TotemTokenSession) TaxRate() (*big.Int, error) {
	return _TotemToken.Contract.TaxRate(&_TotemToken.CallOpts)
}

// TaxRate is a free data retrieval call binding the contract method 0x771a3a1d.
//
// Solidity: function taxRate() view returns(uint256)
func (_TotemToken *TotemTokenCallerSession) TaxRate() (*big.Int, error) {
	return _TotemToken.Contract.TaxRate(&_TotemToken.CallOpts)
}

// TaxationWallet is a free data retrieval call binding the contract method 0x4a5db0a9.
//
// Solidity: function taxationWallet() view returns(address)
func (_TotemToken *TotemTokenCaller) TaxationWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TotemToken.contract.Call(opts, &out, "taxationWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TaxationWallet is a free data retrieval call binding the contract method 0x4a5db0a9.
//
// Solidity: function taxationWallet() view returns(address)
func (_TotemToken *TotemTokenSession) TaxationWallet() (common.Address, error) {
	return _TotemToken.Contract.TaxationWallet(&_TotemToken.CallOpts)
}

// TaxationWallet is a free data retrieval call binding the contract method 0x4a5db0a9.
//
// Solidity: function taxationWallet() view returns(address)
func (_TotemToken *TotemTokenCallerSession) TaxationWallet() (common.Address, error) {
	return _TotemToken.Contract.TaxationWallet(&_TotemToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TotemToken *TotemTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TotemToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TotemToken *TotemTokenSession) TotalSupply() (*big.Int, error) {
	return _TotemToken.Contract.TotalSupply(&_TotemToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TotemToken *TotemTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _TotemToken.Contract.TotalSupply(&_TotemToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TotemToken *TotemTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TotemToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TotemToken *TotemTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TotemToken.Contract.Approve(&_TotemToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TotemToken *TotemTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TotemToken.Contract.Approve(&_TotemToken.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TotemToken *TotemTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TotemToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TotemToken *TotemTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TotemToken.Contract.DecreaseAllowance(&_TotemToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TotemToken *TotemTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TotemToken.Contract.DecreaseAllowance(&_TotemToken.TransactOpts, spender, subtractedValue)
}

// DistributeTokens is a paid mutator transaction binding the contract method 0x9ab1b484.
//
// Solidity: function distributeTokens() returns()
func (_TotemToken *TotemTokenTransactor) DistributeTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TotemToken.contract.Transact(opts, "distributeTokens")
}

// DistributeTokens is a paid mutator transaction binding the contract method 0x9ab1b484.
//
// Solidity: function distributeTokens() returns()
func (_TotemToken *TotemTokenSession) DistributeTokens() (*types.Transaction, error) {
	return _TotemToken.Contract.DistributeTokens(&_TotemToken.TransactOpts)
}

// DistributeTokens is a paid mutator transaction binding the contract method 0x9ab1b484.
//
// Solidity: function distributeTokens() returns()
func (_TotemToken *TotemTokenTransactorSession) DistributeTokens() (*types.Transaction, error) {
	return _TotemToken.Contract.DistributeTokens(&_TotemToken.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TotemToken *TotemTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TotemToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TotemToken *TotemTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TotemToken.Contract.IncreaseAllowance(&_TotemToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TotemToken *TotemTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TotemToken.Contract.IncreaseAllowance(&_TotemToken.TransactOpts, spender, addedValue)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TotemToken *TotemTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TotemToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TotemToken *TotemTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _TotemToken.Contract.RenounceOwnership(&_TotemToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TotemToken *TotemTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TotemToken.Contract.RenounceOwnership(&_TotemToken.TransactOpts)
}

// SetDistributionTeamsAddresses is a paid mutator transaction binding the contract method 0xa8d61e9c.
//
// Solidity: function setDistributionTeamsAddresses(address _CommunityDevelopmentAddr, address _StakingRewardsAddr, address _LiquidityPoolAddr, address _PublicSaleAddr, address _AdvisorsAddr, address _SeedInvestmentAddr, address _PrivateSaleAddr, address _TeamAllocationAddr, address _StrategicRoundAddr) returns()
func (_TotemToken *TotemTokenTransactor) SetDistributionTeamsAddresses(opts *bind.TransactOpts, _CommunityDevelopmentAddr common.Address, _StakingRewardsAddr common.Address, _LiquidityPoolAddr common.Address, _PublicSaleAddr common.Address, _AdvisorsAddr common.Address, _SeedInvestmentAddr common.Address, _PrivateSaleAddr common.Address, _TeamAllocationAddr common.Address, _StrategicRoundAddr common.Address) (*types.Transaction, error) {
	return _TotemToken.contract.Transact(opts, "setDistributionTeamsAddresses", _CommunityDevelopmentAddr, _StakingRewardsAddr, _LiquidityPoolAddr, _PublicSaleAddr, _AdvisorsAddr, _SeedInvestmentAddr, _PrivateSaleAddr, _TeamAllocationAddr, _StrategicRoundAddr)
}

// SetDistributionTeamsAddresses is a paid mutator transaction binding the contract method 0xa8d61e9c.
//
// Solidity: function setDistributionTeamsAddresses(address _CommunityDevelopmentAddr, address _StakingRewardsAddr, address _LiquidityPoolAddr, address _PublicSaleAddr, address _AdvisorsAddr, address _SeedInvestmentAddr, address _PrivateSaleAddr, address _TeamAllocationAddr, address _StrategicRoundAddr) returns()
func (_TotemToken *TotemTokenSession) SetDistributionTeamsAddresses(_CommunityDevelopmentAddr common.Address, _StakingRewardsAddr common.Address, _LiquidityPoolAddr common.Address, _PublicSaleAddr common.Address, _AdvisorsAddr common.Address, _SeedInvestmentAddr common.Address, _PrivateSaleAddr common.Address, _TeamAllocationAddr common.Address, _StrategicRoundAddr common.Address) (*types.Transaction, error) {
	return _TotemToken.Contract.SetDistributionTeamsAddresses(&_TotemToken.TransactOpts, _CommunityDevelopmentAddr, _StakingRewardsAddr, _LiquidityPoolAddr, _PublicSaleAddr, _AdvisorsAddr, _SeedInvestmentAddr, _PrivateSaleAddr, _TeamAllocationAddr, _StrategicRoundAddr)
}

// SetDistributionTeamsAddresses is a paid mutator transaction binding the contract method 0xa8d61e9c.
//
// Solidity: function setDistributionTeamsAddresses(address _CommunityDevelopmentAddr, address _StakingRewardsAddr, address _LiquidityPoolAddr, address _PublicSaleAddr, address _AdvisorsAddr, address _SeedInvestmentAddr, address _PrivateSaleAddr, address _TeamAllocationAddr, address _StrategicRoundAddr) returns()
func (_TotemToken *TotemTokenTransactorSession) SetDistributionTeamsAddresses(_CommunityDevelopmentAddr common.Address, _StakingRewardsAddr common.Address, _LiquidityPoolAddr common.Address, _PublicSaleAddr common.Address, _AdvisorsAddr common.Address, _SeedInvestmentAddr common.Address, _PrivateSaleAddr common.Address, _TeamAllocationAddr common.Address, _StrategicRoundAddr common.Address) (*types.Transaction, error) {
	return _TotemToken.Contract.SetDistributionTeamsAddresses(&_TotemToken.TransactOpts, _CommunityDevelopmentAddr, _StakingRewardsAddr, _LiquidityPoolAddr, _PublicSaleAddr, _AdvisorsAddr, _SeedInvestmentAddr, _PrivateSaleAddr, _TeamAllocationAddr, _StrategicRoundAddr)
}

// SetLocker is a paid mutator transaction binding the contract method 0x171060ec.
//
// Solidity: function setLocker(address _locker) returns()
func (_TotemToken *TotemTokenTransactor) SetLocker(opts *bind.TransactOpts, _locker common.Address) (*types.Transaction, error) {
	return _TotemToken.contract.Transact(opts, "setLocker", _locker)
}

// SetLocker is a paid mutator transaction binding the contract method 0x171060ec.
//
// Solidity: function setLocker(address _locker) returns()
func (_TotemToken *TotemTokenSession) SetLocker(_locker common.Address) (*types.Transaction, error) {
	return _TotemToken.Contract.SetLocker(&_TotemToken.TransactOpts, _locker)
}

// SetLocker is a paid mutator transaction binding the contract method 0x171060ec.
//
// Solidity: function setLocker(address _locker) returns()
func (_TotemToken *TotemTokenTransactorSession) SetLocker(_locker common.Address) (*types.Transaction, error) {
	return _TotemToken.Contract.SetLocker(&_TotemToken.TransactOpts, _locker)
}

// SetTaxExemptStatus is a paid mutator transaction binding the contract method 0xc125c861.
//
// Solidity: function setTaxExemptStatus(address account, bool status) returns()
func (_TotemToken *TotemTokenTransactor) SetTaxExemptStatus(opts *bind.TransactOpts, account common.Address, status bool) (*types.Transaction, error) {
	return _TotemToken.contract.Transact(opts, "setTaxExemptStatus", account, status)
}

// SetTaxExemptStatus is a paid mutator transaction binding the contract method 0xc125c861.
//
// Solidity: function setTaxExemptStatus(address account, bool status) returns()
func (_TotemToken *TotemTokenSession) SetTaxExemptStatus(account common.Address, status bool) (*types.Transaction, error) {
	return _TotemToken.Contract.SetTaxExemptStatus(&_TotemToken.TransactOpts, account, status)
}

// SetTaxExemptStatus is a paid mutator transaction binding the contract method 0xc125c861.
//
// Solidity: function setTaxExemptStatus(address account, bool status) returns()
func (_TotemToken *TotemTokenTransactorSession) SetTaxExemptStatus(account common.Address, status bool) (*types.Transaction, error) {
	return _TotemToken.Contract.SetTaxExemptStatus(&_TotemToken.TransactOpts, account, status)
}

// SetTaxRate is a paid mutator transaction binding the contract method 0xc6d69a30.
//
// Solidity: function setTaxRate(uint256 newTaxRate) returns()
func (_TotemToken *TotemTokenTransactor) SetTaxRate(opts *bind.TransactOpts, newTaxRate *big.Int) (*types.Transaction, error) {
	return _TotemToken.contract.Transact(opts, "setTaxRate", newTaxRate)
}

// SetTaxRate is a paid mutator transaction binding the contract method 0xc6d69a30.
//
// Solidity: function setTaxRate(uint256 newTaxRate) returns()
func (_TotemToken *TotemTokenSession) SetTaxRate(newTaxRate *big.Int) (*types.Transaction, error) {
	return _TotemToken.Contract.SetTaxRate(&_TotemToken.TransactOpts, newTaxRate)
}

// SetTaxRate is a paid mutator transaction binding the contract method 0xc6d69a30.
//
// Solidity: function setTaxRate(uint256 newTaxRate) returns()
func (_TotemToken *TotemTokenTransactorSession) SetTaxRate(newTaxRate *big.Int) (*types.Transaction, error) {
	return _TotemToken.Contract.SetTaxRate(&_TotemToken.TransactOpts, newTaxRate)
}

// SetTaxationWallet is a paid mutator transaction binding the contract method 0x28e2f997.
//
// Solidity: function setTaxationWallet(address newTaxationWallet) returns()
func (_TotemToken *TotemTokenTransactor) SetTaxationWallet(opts *bind.TransactOpts, newTaxationWallet common.Address) (*types.Transaction, error) {
	return _TotemToken.contract.Transact(opts, "setTaxationWallet", newTaxationWallet)
}

// SetTaxationWallet is a paid mutator transaction binding the contract method 0x28e2f997.
//
// Solidity: function setTaxationWallet(address newTaxationWallet) returns()
func (_TotemToken *TotemTokenSession) SetTaxationWallet(newTaxationWallet common.Address) (*types.Transaction, error) {
	return _TotemToken.Contract.SetTaxationWallet(&_TotemToken.TransactOpts, newTaxationWallet)
}

// SetTaxationWallet is a paid mutator transaction binding the contract method 0x28e2f997.
//
// Solidity: function setTaxationWallet(address newTaxationWallet) returns()
func (_TotemToken *TotemTokenTransactorSession) SetTaxationWallet(newTaxationWallet common.Address) (*types.Transaction, error) {
	return _TotemToken.Contract.SetTaxationWallet(&_TotemToken.TransactOpts, newTaxationWallet)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_TotemToken *TotemTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TotemToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_TotemToken *TotemTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TotemToken.Contract.Transfer(&_TotemToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_TotemToken *TotemTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TotemToken.Contract.Transfer(&_TotemToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_TotemToken *TotemTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TotemToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_TotemToken *TotemTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TotemToken.Contract.TransferFrom(&_TotemToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_TotemToken *TotemTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TotemToken.Contract.TransferFrom(&_TotemToken.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TotemToken *TotemTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TotemToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TotemToken *TotemTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TotemToken.Contract.TransferOwnership(&_TotemToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TotemToken *TotemTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TotemToken.Contract.TransferOwnership(&_TotemToken.TransactOpts, newOwner)
}

// TotemTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TotemToken contract.
type TotemTokenApprovalIterator struct {
	Event *TotemTokenApproval // Event containing the contract specifics and raw log

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
func (it *TotemTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TotemTokenApproval)
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
		it.Event = new(TotemTokenApproval)
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
func (it *TotemTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TotemTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TotemTokenApproval represents a Approval event raised by the TotemToken contract.
type TotemTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TotemToken *TotemTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TotemTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TotemToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TotemTokenApprovalIterator{contract: _TotemToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TotemToken *TotemTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TotemTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TotemToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TotemTokenApproval)
				if err := _TotemToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TotemToken *TotemTokenFilterer) ParseApproval(log types.Log) (*TotemTokenApproval, error) {
	event := new(TotemTokenApproval)
	if err := _TotemToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TotemTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TotemToken contract.
type TotemTokenOwnershipTransferredIterator struct {
	Event *TotemTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TotemTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TotemTokenOwnershipTransferred)
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
		it.Event = new(TotemTokenOwnershipTransferred)
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
func (it *TotemTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TotemTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TotemTokenOwnershipTransferred represents a OwnershipTransferred event raised by the TotemToken contract.
type TotemTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TotemToken *TotemTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TotemTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TotemToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TotemTokenOwnershipTransferredIterator{contract: _TotemToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TotemToken *TotemTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TotemTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TotemToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TotemTokenOwnershipTransferred)
				if err := _TotemToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_TotemToken *TotemTokenFilterer) ParseOwnershipTransferred(log types.Log) (*TotemTokenOwnershipTransferred, error) {
	event := new(TotemTokenOwnershipTransferred)
	if err := _TotemToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TotemTokenSetLockerIterator is returned from FilterSetLocker and is used to iterate over the raw logs and unpacked data for SetLocker events raised by the TotemToken contract.
type TotemTokenSetLockerIterator struct {
	Event *TotemTokenSetLocker // Event containing the contract specifics and raw log

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
func (it *TotemTokenSetLockerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TotemTokenSetLocker)
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
		it.Event = new(TotemTokenSetLocker)
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
func (it *TotemTokenSetLockerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TotemTokenSetLockerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TotemTokenSetLocker represents a SetLocker event raised by the TotemToken contract.
type TotemTokenSetLocker struct {
	Locker common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetLocker is a free log retrieval operation binding the contract event 0x2ad3bcdfc9b535ec0e163460ef3e0b8015fca3aa8e9cc778ca27876a07bd2827.
//
// Solidity: event SetLocker(address indexed locker)
func (_TotemToken *TotemTokenFilterer) FilterSetLocker(opts *bind.FilterOpts, locker []common.Address) (*TotemTokenSetLockerIterator, error) {

	var lockerRule []interface{}
	for _, lockerItem := range locker {
		lockerRule = append(lockerRule, lockerItem)
	}

	logs, sub, err := _TotemToken.contract.FilterLogs(opts, "SetLocker", lockerRule)
	if err != nil {
		return nil, err
	}
	return &TotemTokenSetLockerIterator{contract: _TotemToken.contract, event: "SetLocker", logs: logs, sub: sub}, nil
}

// WatchSetLocker is a free log subscription operation binding the contract event 0x2ad3bcdfc9b535ec0e163460ef3e0b8015fca3aa8e9cc778ca27876a07bd2827.
//
// Solidity: event SetLocker(address indexed locker)
func (_TotemToken *TotemTokenFilterer) WatchSetLocker(opts *bind.WatchOpts, sink chan<- *TotemTokenSetLocker, locker []common.Address) (event.Subscription, error) {

	var lockerRule []interface{}
	for _, lockerItem := range locker {
		lockerRule = append(lockerRule, lockerItem)
	}

	logs, sub, err := _TotemToken.contract.WatchLogs(opts, "SetLocker", lockerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TotemTokenSetLocker)
				if err := _TotemToken.contract.UnpackLog(event, "SetLocker", log); err != nil {
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
func (_TotemToken *TotemTokenFilterer) ParseSetLocker(log types.Log) (*TotemTokenSetLocker, error) {
	event := new(TotemTokenSetLocker)
	if err := _TotemToken.contract.UnpackLog(event, "SetLocker", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TotemTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TotemToken contract.
type TotemTokenTransferIterator struct {
	Event *TotemTokenTransfer // Event containing the contract specifics and raw log

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
func (it *TotemTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TotemTokenTransfer)
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
		it.Event = new(TotemTokenTransfer)
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
func (it *TotemTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TotemTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TotemTokenTransfer represents a Transfer event raised by the TotemToken contract.
type TotemTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TotemToken *TotemTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TotemTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TotemToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TotemTokenTransferIterator{contract: _TotemToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TotemToken *TotemTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TotemTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TotemToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TotemTokenTransfer)
				if err := _TotemToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TotemToken *TotemTokenFilterer) ParseTransfer(log types.Log) (*TotemTokenTransfer, error) {
	event := new(TotemTokenTransfer)
	if err := _TotemToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
