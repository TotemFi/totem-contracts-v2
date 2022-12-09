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

// ITotemTokenMetaData contains all meta data concerning the ITotemToken contract.
var ITotemTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributeTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_CommunityDevelopmentAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_StakingRewardsAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_LiquidityPoolAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_PublicSaleAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_AdvisorsAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_SeedInvestmentAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_PrivateSaleAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_TeamAllocationAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_StrategicRoundAddr\",\"type\":\"address\"}],\"name\":\"setDistributionTeamsAddresses\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_locker\",\"type\":\"address\"}],\"name\":\"setLocker\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"setTaxExemptStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newTaxRate\",\"type\":\"uint256\"}],\"name\":\"setTaxRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newTaxationWallet\",\"type\":\"address\"}],\"name\":\"setTaxationWallet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_msgSender\",\"type\":\"address\"}],\"name\":\"taxExempt\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taxRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"taxationWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ITotemTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use ITotemTokenMetaData.ABI instead.
var ITotemTokenABI = ITotemTokenMetaData.ABI

// ITotemToken is an auto generated Go binding around an Ethereum contract.
type ITotemToken struct {
	ITotemTokenCaller     // Read-only binding to the contract
	ITotemTokenTransactor // Write-only binding to the contract
	ITotemTokenFilterer   // Log filterer for contract events
}

// ITotemTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type ITotemTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITotemTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ITotemTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITotemTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ITotemTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ITotemTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ITotemTokenSession struct {
	Contract     *ITotemToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ITotemTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ITotemTokenCallerSession struct {
	Contract *ITotemTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ITotemTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ITotemTokenTransactorSession struct {
	Contract     *ITotemTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ITotemTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type ITotemTokenRaw struct {
	Contract *ITotemToken // Generic contract binding to access the raw methods on
}

// ITotemTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ITotemTokenCallerRaw struct {
	Contract *ITotemTokenCaller // Generic read-only contract binding to access the raw methods on
}

// ITotemTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ITotemTokenTransactorRaw struct {
	Contract *ITotemTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewITotemToken creates a new instance of ITotemToken, bound to a specific deployed contract.
func NewITotemToken(address common.Address, backend bind.ContractBackend) (*ITotemToken, error) {
	contract, err := bindITotemToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ITotemToken{ITotemTokenCaller: ITotemTokenCaller{contract: contract}, ITotemTokenTransactor: ITotemTokenTransactor{contract: contract}, ITotemTokenFilterer: ITotemTokenFilterer{contract: contract}}, nil
}

// NewITotemTokenCaller creates a new read-only instance of ITotemToken, bound to a specific deployed contract.
func NewITotemTokenCaller(address common.Address, caller bind.ContractCaller) (*ITotemTokenCaller, error) {
	contract, err := bindITotemToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ITotemTokenCaller{contract: contract}, nil
}

// NewITotemTokenTransactor creates a new write-only instance of ITotemToken, bound to a specific deployed contract.
func NewITotemTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*ITotemTokenTransactor, error) {
	contract, err := bindITotemToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ITotemTokenTransactor{contract: contract}, nil
}

// NewITotemTokenFilterer creates a new log filterer instance of ITotemToken, bound to a specific deployed contract.
func NewITotemTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*ITotemTokenFilterer, error) {
	contract, err := bindITotemToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ITotemTokenFilterer{contract: contract}, nil
}

// bindITotemToken binds a generic wrapper to an already deployed contract.
func bindITotemToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ITotemTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITotemToken *ITotemTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITotemToken.Contract.ITotemTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITotemToken *ITotemTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITotemToken.Contract.ITotemTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITotemToken *ITotemTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITotemToken.Contract.ITotemTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ITotemToken *ITotemTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ITotemToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ITotemToken *ITotemTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITotemToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ITotemToken *ITotemTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ITotemToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ITotemToken *ITotemTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ITotemToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ITotemToken *ITotemTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ITotemToken.Contract.BalanceOf(&_ITotemToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ITotemToken *ITotemTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ITotemToken.Contract.BalanceOf(&_ITotemToken.CallOpts, account)
}

// DistributeTokens is a paid mutator transaction binding the contract method 0x9ab1b484.
//
// Solidity: function distributeTokens() returns()
func (_ITotemToken *ITotemTokenTransactor) DistributeTokens(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITotemToken.contract.Transact(opts, "distributeTokens")
}

// DistributeTokens is a paid mutator transaction binding the contract method 0x9ab1b484.
//
// Solidity: function distributeTokens() returns()
func (_ITotemToken *ITotemTokenSession) DistributeTokens() (*types.Transaction, error) {
	return _ITotemToken.Contract.DistributeTokens(&_ITotemToken.TransactOpts)
}

// DistributeTokens is a paid mutator transaction binding the contract method 0x9ab1b484.
//
// Solidity: function distributeTokens() returns()
func (_ITotemToken *ITotemTokenTransactorSession) DistributeTokens() (*types.Transaction, error) {
	return _ITotemToken.Contract.DistributeTokens(&_ITotemToken.TransactOpts)
}

// SetDistributionTeamsAddresses is a paid mutator transaction binding the contract method 0xa8d61e9c.
//
// Solidity: function setDistributionTeamsAddresses(address _CommunityDevelopmentAddr, address _StakingRewardsAddr, address _LiquidityPoolAddr, address _PublicSaleAddr, address _AdvisorsAddr, address _SeedInvestmentAddr, address _PrivateSaleAddr, address _TeamAllocationAddr, address _StrategicRoundAddr) returns()
func (_ITotemToken *ITotemTokenTransactor) SetDistributionTeamsAddresses(opts *bind.TransactOpts, _CommunityDevelopmentAddr common.Address, _StakingRewardsAddr common.Address, _LiquidityPoolAddr common.Address, _PublicSaleAddr common.Address, _AdvisorsAddr common.Address, _SeedInvestmentAddr common.Address, _PrivateSaleAddr common.Address, _TeamAllocationAddr common.Address, _StrategicRoundAddr common.Address) (*types.Transaction, error) {
	return _ITotemToken.contract.Transact(opts, "setDistributionTeamsAddresses", _CommunityDevelopmentAddr, _StakingRewardsAddr, _LiquidityPoolAddr, _PublicSaleAddr, _AdvisorsAddr, _SeedInvestmentAddr, _PrivateSaleAddr, _TeamAllocationAddr, _StrategicRoundAddr)
}

// SetDistributionTeamsAddresses is a paid mutator transaction binding the contract method 0xa8d61e9c.
//
// Solidity: function setDistributionTeamsAddresses(address _CommunityDevelopmentAddr, address _StakingRewardsAddr, address _LiquidityPoolAddr, address _PublicSaleAddr, address _AdvisorsAddr, address _SeedInvestmentAddr, address _PrivateSaleAddr, address _TeamAllocationAddr, address _StrategicRoundAddr) returns()
func (_ITotemToken *ITotemTokenSession) SetDistributionTeamsAddresses(_CommunityDevelopmentAddr common.Address, _StakingRewardsAddr common.Address, _LiquidityPoolAddr common.Address, _PublicSaleAddr common.Address, _AdvisorsAddr common.Address, _SeedInvestmentAddr common.Address, _PrivateSaleAddr common.Address, _TeamAllocationAddr common.Address, _StrategicRoundAddr common.Address) (*types.Transaction, error) {
	return _ITotemToken.Contract.SetDistributionTeamsAddresses(&_ITotemToken.TransactOpts, _CommunityDevelopmentAddr, _StakingRewardsAddr, _LiquidityPoolAddr, _PublicSaleAddr, _AdvisorsAddr, _SeedInvestmentAddr, _PrivateSaleAddr, _TeamAllocationAddr, _StrategicRoundAddr)
}

// SetDistributionTeamsAddresses is a paid mutator transaction binding the contract method 0xa8d61e9c.
//
// Solidity: function setDistributionTeamsAddresses(address _CommunityDevelopmentAddr, address _StakingRewardsAddr, address _LiquidityPoolAddr, address _PublicSaleAddr, address _AdvisorsAddr, address _SeedInvestmentAddr, address _PrivateSaleAddr, address _TeamAllocationAddr, address _StrategicRoundAddr) returns()
func (_ITotemToken *ITotemTokenTransactorSession) SetDistributionTeamsAddresses(_CommunityDevelopmentAddr common.Address, _StakingRewardsAddr common.Address, _LiquidityPoolAddr common.Address, _PublicSaleAddr common.Address, _AdvisorsAddr common.Address, _SeedInvestmentAddr common.Address, _PrivateSaleAddr common.Address, _TeamAllocationAddr common.Address, _StrategicRoundAddr common.Address) (*types.Transaction, error) {
	return _ITotemToken.Contract.SetDistributionTeamsAddresses(&_ITotemToken.TransactOpts, _CommunityDevelopmentAddr, _StakingRewardsAddr, _LiquidityPoolAddr, _PublicSaleAddr, _AdvisorsAddr, _SeedInvestmentAddr, _PrivateSaleAddr, _TeamAllocationAddr, _StrategicRoundAddr)
}

// SetLocker is a paid mutator transaction binding the contract method 0x171060ec.
//
// Solidity: function setLocker(address _locker) returns()
func (_ITotemToken *ITotemTokenTransactor) SetLocker(opts *bind.TransactOpts, _locker common.Address) (*types.Transaction, error) {
	return _ITotemToken.contract.Transact(opts, "setLocker", _locker)
}

// SetLocker is a paid mutator transaction binding the contract method 0x171060ec.
//
// Solidity: function setLocker(address _locker) returns()
func (_ITotemToken *ITotemTokenSession) SetLocker(_locker common.Address) (*types.Transaction, error) {
	return _ITotemToken.Contract.SetLocker(&_ITotemToken.TransactOpts, _locker)
}

// SetLocker is a paid mutator transaction binding the contract method 0x171060ec.
//
// Solidity: function setLocker(address _locker) returns()
func (_ITotemToken *ITotemTokenTransactorSession) SetLocker(_locker common.Address) (*types.Transaction, error) {
	return _ITotemToken.Contract.SetLocker(&_ITotemToken.TransactOpts, _locker)
}

// SetTaxExemptStatus is a paid mutator transaction binding the contract method 0xc125c861.
//
// Solidity: function setTaxExemptStatus(address account, bool status) returns()
func (_ITotemToken *ITotemTokenTransactor) SetTaxExemptStatus(opts *bind.TransactOpts, account common.Address, status bool) (*types.Transaction, error) {
	return _ITotemToken.contract.Transact(opts, "setTaxExemptStatus", account, status)
}

// SetTaxExemptStatus is a paid mutator transaction binding the contract method 0xc125c861.
//
// Solidity: function setTaxExemptStatus(address account, bool status) returns()
func (_ITotemToken *ITotemTokenSession) SetTaxExemptStatus(account common.Address, status bool) (*types.Transaction, error) {
	return _ITotemToken.Contract.SetTaxExemptStatus(&_ITotemToken.TransactOpts, account, status)
}

// SetTaxExemptStatus is a paid mutator transaction binding the contract method 0xc125c861.
//
// Solidity: function setTaxExemptStatus(address account, bool status) returns()
func (_ITotemToken *ITotemTokenTransactorSession) SetTaxExemptStatus(account common.Address, status bool) (*types.Transaction, error) {
	return _ITotemToken.Contract.SetTaxExemptStatus(&_ITotemToken.TransactOpts, account, status)
}

// SetTaxRate is a paid mutator transaction binding the contract method 0xc6d69a30.
//
// Solidity: function setTaxRate(uint256 newTaxRate) returns()
func (_ITotemToken *ITotemTokenTransactor) SetTaxRate(opts *bind.TransactOpts, newTaxRate *big.Int) (*types.Transaction, error) {
	return _ITotemToken.contract.Transact(opts, "setTaxRate", newTaxRate)
}

// SetTaxRate is a paid mutator transaction binding the contract method 0xc6d69a30.
//
// Solidity: function setTaxRate(uint256 newTaxRate) returns()
func (_ITotemToken *ITotemTokenSession) SetTaxRate(newTaxRate *big.Int) (*types.Transaction, error) {
	return _ITotemToken.Contract.SetTaxRate(&_ITotemToken.TransactOpts, newTaxRate)
}

// SetTaxRate is a paid mutator transaction binding the contract method 0xc6d69a30.
//
// Solidity: function setTaxRate(uint256 newTaxRate) returns()
func (_ITotemToken *ITotemTokenTransactorSession) SetTaxRate(newTaxRate *big.Int) (*types.Transaction, error) {
	return _ITotemToken.Contract.SetTaxRate(&_ITotemToken.TransactOpts, newTaxRate)
}

// SetTaxationWallet is a paid mutator transaction binding the contract method 0x28e2f997.
//
// Solidity: function setTaxationWallet(address newTaxationWallet) returns()
func (_ITotemToken *ITotemTokenTransactor) SetTaxationWallet(opts *bind.TransactOpts, newTaxationWallet common.Address) (*types.Transaction, error) {
	return _ITotemToken.contract.Transact(opts, "setTaxationWallet", newTaxationWallet)
}

// SetTaxationWallet is a paid mutator transaction binding the contract method 0x28e2f997.
//
// Solidity: function setTaxationWallet(address newTaxationWallet) returns()
func (_ITotemToken *ITotemTokenSession) SetTaxationWallet(newTaxationWallet common.Address) (*types.Transaction, error) {
	return _ITotemToken.Contract.SetTaxationWallet(&_ITotemToken.TransactOpts, newTaxationWallet)
}

// SetTaxationWallet is a paid mutator transaction binding the contract method 0x28e2f997.
//
// Solidity: function setTaxationWallet(address newTaxationWallet) returns()
func (_ITotemToken *ITotemTokenTransactorSession) SetTaxationWallet(newTaxationWallet common.Address) (*types.Transaction, error) {
	return _ITotemToken.Contract.SetTaxationWallet(&_ITotemToken.TransactOpts, newTaxationWallet)
}

// TaxExempt is a paid mutator transaction binding the contract method 0xd1ecfc68.
//
// Solidity: function taxExempt(address _msgSender) returns(bool)
func (_ITotemToken *ITotemTokenTransactor) TaxExempt(opts *bind.TransactOpts, _msgSender common.Address) (*types.Transaction, error) {
	return _ITotemToken.contract.Transact(opts, "taxExempt", _msgSender)
}

// TaxExempt is a paid mutator transaction binding the contract method 0xd1ecfc68.
//
// Solidity: function taxExempt(address _msgSender) returns(bool)
func (_ITotemToken *ITotemTokenSession) TaxExempt(_msgSender common.Address) (*types.Transaction, error) {
	return _ITotemToken.Contract.TaxExempt(&_ITotemToken.TransactOpts, _msgSender)
}

// TaxExempt is a paid mutator transaction binding the contract method 0xd1ecfc68.
//
// Solidity: function taxExempt(address _msgSender) returns(bool)
func (_ITotemToken *ITotemTokenTransactorSession) TaxExempt(_msgSender common.Address) (*types.Transaction, error) {
	return _ITotemToken.Contract.TaxExempt(&_ITotemToken.TransactOpts, _msgSender)
}

// TaxRate is a paid mutator transaction binding the contract method 0x771a3a1d.
//
// Solidity: function taxRate() returns(uint256)
func (_ITotemToken *ITotemTokenTransactor) TaxRate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITotemToken.contract.Transact(opts, "taxRate")
}

// TaxRate is a paid mutator transaction binding the contract method 0x771a3a1d.
//
// Solidity: function taxRate() returns(uint256)
func (_ITotemToken *ITotemTokenSession) TaxRate() (*types.Transaction, error) {
	return _ITotemToken.Contract.TaxRate(&_ITotemToken.TransactOpts)
}

// TaxRate is a paid mutator transaction binding the contract method 0x771a3a1d.
//
// Solidity: function taxRate() returns(uint256)
func (_ITotemToken *ITotemTokenTransactorSession) TaxRate() (*types.Transaction, error) {
	return _ITotemToken.Contract.TaxRate(&_ITotemToken.TransactOpts)
}

// TaxationWallet is a paid mutator transaction binding the contract method 0x4a5db0a9.
//
// Solidity: function taxationWallet() returns(address)
func (_ITotemToken *ITotemTokenTransactor) TaxationWallet(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ITotemToken.contract.Transact(opts, "taxationWallet")
}

// TaxationWallet is a paid mutator transaction binding the contract method 0x4a5db0a9.
//
// Solidity: function taxationWallet() returns(address)
func (_ITotemToken *ITotemTokenSession) TaxationWallet() (*types.Transaction, error) {
	return _ITotemToken.Contract.TaxationWallet(&_ITotemToken.TransactOpts)
}

// TaxationWallet is a paid mutator transaction binding the contract method 0x4a5db0a9.
//
// Solidity: function taxationWallet() returns(address)
func (_ITotemToken *ITotemTokenTransactorSession) TaxationWallet() (*types.Transaction, error) {
	return _ITotemToken.Contract.TaxationWallet(&_ITotemToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ITotemToken *ITotemTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ITotemToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ITotemToken *ITotemTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ITotemToken.Contract.Transfer(&_ITotemToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ITotemToken *ITotemTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ITotemToken.Contract.Transfer(&_ITotemToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ITotemToken *ITotemTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ITotemToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ITotemToken *ITotemTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ITotemToken.Contract.TransferFrom(&_ITotemToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ITotemToken *ITotemTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ITotemToken.Contract.TransferFrom(&_ITotemToken.TransactOpts, sender, recipient, amount)
}
