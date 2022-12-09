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

// USDRetrieverMetaData contains all meta data concerning the USDRetriever contract.
var USDRetrieverMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ApproveTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ReceivedTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferTokens\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getUSDBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUSDToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// USDRetrieverABI is the input ABI used to generate the binding from.
// Deprecated: Use USDRetrieverMetaData.ABI instead.
var USDRetrieverABI = USDRetrieverMetaData.ABI

// USDRetriever is an auto generated Go binding around an Ethereum contract.
type USDRetriever struct {
	USDRetrieverCaller     // Read-only binding to the contract
	USDRetrieverTransactor // Write-only binding to the contract
	USDRetrieverFilterer   // Log filterer for contract events
}

// USDRetrieverCaller is an auto generated read-only Go binding around an Ethereum contract.
type USDRetrieverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDRetrieverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type USDRetrieverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDRetrieverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type USDRetrieverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDRetrieverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type USDRetrieverSession struct {
	Contract     *USDRetriever     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// USDRetrieverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type USDRetrieverCallerSession struct {
	Contract *USDRetrieverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// USDRetrieverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type USDRetrieverTransactorSession struct {
	Contract     *USDRetrieverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// USDRetrieverRaw is an auto generated low-level Go binding around an Ethereum contract.
type USDRetrieverRaw struct {
	Contract *USDRetriever // Generic contract binding to access the raw methods on
}

// USDRetrieverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type USDRetrieverCallerRaw struct {
	Contract *USDRetrieverCaller // Generic read-only contract binding to access the raw methods on
}

// USDRetrieverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type USDRetrieverTransactorRaw struct {
	Contract *USDRetrieverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUSDRetriever creates a new instance of USDRetriever, bound to a specific deployed contract.
func NewUSDRetriever(address common.Address, backend bind.ContractBackend) (*USDRetriever, error) {
	contract, err := bindUSDRetriever(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &USDRetriever{USDRetrieverCaller: USDRetrieverCaller{contract: contract}, USDRetrieverTransactor: USDRetrieverTransactor{contract: contract}, USDRetrieverFilterer: USDRetrieverFilterer{contract: contract}}, nil
}

// NewUSDRetrieverCaller creates a new read-only instance of USDRetriever, bound to a specific deployed contract.
func NewUSDRetrieverCaller(address common.Address, caller bind.ContractCaller) (*USDRetrieverCaller, error) {
	contract, err := bindUSDRetriever(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &USDRetrieverCaller{contract: contract}, nil
}

// NewUSDRetrieverTransactor creates a new write-only instance of USDRetriever, bound to a specific deployed contract.
func NewUSDRetrieverTransactor(address common.Address, transactor bind.ContractTransactor) (*USDRetrieverTransactor, error) {
	contract, err := bindUSDRetriever(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &USDRetrieverTransactor{contract: contract}, nil
}

// NewUSDRetrieverFilterer creates a new log filterer instance of USDRetriever, bound to a specific deployed contract.
func NewUSDRetrieverFilterer(address common.Address, filterer bind.ContractFilterer) (*USDRetrieverFilterer, error) {
	contract, err := bindUSDRetriever(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &USDRetrieverFilterer{contract: contract}, nil
}

// bindUSDRetriever binds a generic wrapper to an already deployed contract.
func bindUSDRetriever(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(USDRetrieverABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_USDRetriever *USDRetrieverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USDRetriever.Contract.USDRetrieverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_USDRetriever *USDRetrieverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDRetriever.Contract.USDRetrieverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_USDRetriever *USDRetrieverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USDRetriever.Contract.USDRetrieverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_USDRetriever *USDRetrieverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USDRetriever.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_USDRetriever *USDRetrieverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDRetriever.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_USDRetriever *USDRetrieverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USDRetriever.Contract.contract.Transact(opts, method, params...)
}

// GetUSDBalance is a free data retrieval call binding the contract method 0xc1201054.
//
// Solidity: function getUSDBalance() view returns(uint256)
func (_USDRetriever *USDRetrieverCaller) GetUSDBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _USDRetriever.contract.Call(opts, &out, "getUSDBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUSDBalance is a free data retrieval call binding the contract method 0xc1201054.
//
// Solidity: function getUSDBalance() view returns(uint256)
func (_USDRetriever *USDRetrieverSession) GetUSDBalance() (*big.Int, error) {
	return _USDRetriever.Contract.GetUSDBalance(&_USDRetriever.CallOpts)
}

// GetUSDBalance is a free data retrieval call binding the contract method 0xc1201054.
//
// Solidity: function getUSDBalance() view returns(uint256)
func (_USDRetriever *USDRetrieverCallerSession) GetUSDBalance() (*big.Int, error) {
	return _USDRetriever.Contract.GetUSDBalance(&_USDRetriever.CallOpts)
}

// GetUSDToken is a free data retrieval call binding the contract method 0xf5e96fc9.
//
// Solidity: function getUSDToken() view returns(address)
func (_USDRetriever *USDRetrieverCaller) GetUSDToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _USDRetriever.contract.Call(opts, &out, "getUSDToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUSDToken is a free data retrieval call binding the contract method 0xf5e96fc9.
//
// Solidity: function getUSDToken() view returns(address)
func (_USDRetriever *USDRetrieverSession) GetUSDToken() (common.Address, error) {
	return _USDRetriever.Contract.GetUSDToken(&_USDRetriever.CallOpts)
}

// GetUSDToken is a free data retrieval call binding the contract method 0xf5e96fc9.
//
// Solidity: function getUSDToken() view returns(address)
func (_USDRetriever *USDRetrieverCallerSession) GetUSDToken() (common.Address, error) {
	return _USDRetriever.Contract.GetUSDToken(&_USDRetriever.CallOpts)
}

// USDRetrieverApproveTokensIterator is returned from FilterApproveTokens and is used to iterate over the raw logs and unpacked data for ApproveTokens events raised by the USDRetriever contract.
type USDRetrieverApproveTokensIterator struct {
	Event *USDRetrieverApproveTokens // Event containing the contract specifics and raw log

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
func (it *USDRetrieverApproveTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDRetrieverApproveTokens)
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
		it.Event = new(USDRetrieverApproveTokens)
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
func (it *USDRetrieverApproveTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDRetrieverApproveTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDRetrieverApproveTokens represents a ApproveTokens event raised by the USDRetriever contract.
type USDRetrieverApproveTokens struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterApproveTokens is a free log retrieval operation binding the contract event 0xd77df63d9076d067d9fd7af16c3d67db48b84064c3314f5900ae919922d364a1.
//
// Solidity: event ApproveTokens(address indexed to, uint256 amount)
func (_USDRetriever *USDRetrieverFilterer) FilterApproveTokens(opts *bind.FilterOpts, to []common.Address) (*USDRetrieverApproveTokensIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USDRetriever.contract.FilterLogs(opts, "ApproveTokens", toRule)
	if err != nil {
		return nil, err
	}
	return &USDRetrieverApproveTokensIterator{contract: _USDRetriever.contract, event: "ApproveTokens", logs: logs, sub: sub}, nil
}

// WatchApproveTokens is a free log subscription operation binding the contract event 0xd77df63d9076d067d9fd7af16c3d67db48b84064c3314f5900ae919922d364a1.
//
// Solidity: event ApproveTokens(address indexed to, uint256 amount)
func (_USDRetriever *USDRetrieverFilterer) WatchApproveTokens(opts *bind.WatchOpts, sink chan<- *USDRetrieverApproveTokens, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USDRetriever.contract.WatchLogs(opts, "ApproveTokens", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDRetrieverApproveTokens)
				if err := _USDRetriever.contract.UnpackLog(event, "ApproveTokens", log); err != nil {
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

// ParseApproveTokens is a log parse operation binding the contract event 0xd77df63d9076d067d9fd7af16c3d67db48b84064c3314f5900ae919922d364a1.
//
// Solidity: event ApproveTokens(address indexed to, uint256 amount)
func (_USDRetriever *USDRetrieverFilterer) ParseApproveTokens(log types.Log) (*USDRetrieverApproveTokens, error) {
	event := new(USDRetrieverApproveTokens)
	if err := _USDRetriever.contract.UnpackLog(event, "ApproveTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// USDRetrieverReceivedTokensIterator is returned from FilterReceivedTokens and is used to iterate over the raw logs and unpacked data for ReceivedTokens events raised by the USDRetriever contract.
type USDRetrieverReceivedTokensIterator struct {
	Event *USDRetrieverReceivedTokens // Event containing the contract specifics and raw log

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
func (it *USDRetrieverReceivedTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDRetrieverReceivedTokens)
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
		it.Event = new(USDRetrieverReceivedTokens)
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
func (it *USDRetrieverReceivedTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDRetrieverReceivedTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDRetrieverReceivedTokens represents a ReceivedTokens event raised by the USDRetriever contract.
type USDRetrieverReceivedTokens struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedTokens is a free log retrieval operation binding the contract event 0x2946de6c4ec03d8d15126164a7c0da68d7c6835173e41827a7a715f8becb07a8.
//
// Solidity: event ReceivedTokens(address indexed from, uint256 amount)
func (_USDRetriever *USDRetrieverFilterer) FilterReceivedTokens(opts *bind.FilterOpts, from []common.Address) (*USDRetrieverReceivedTokensIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _USDRetriever.contract.FilterLogs(opts, "ReceivedTokens", fromRule)
	if err != nil {
		return nil, err
	}
	return &USDRetrieverReceivedTokensIterator{contract: _USDRetriever.contract, event: "ReceivedTokens", logs: logs, sub: sub}, nil
}

// WatchReceivedTokens is a free log subscription operation binding the contract event 0x2946de6c4ec03d8d15126164a7c0da68d7c6835173e41827a7a715f8becb07a8.
//
// Solidity: event ReceivedTokens(address indexed from, uint256 amount)
func (_USDRetriever *USDRetrieverFilterer) WatchReceivedTokens(opts *bind.WatchOpts, sink chan<- *USDRetrieverReceivedTokens, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _USDRetriever.contract.WatchLogs(opts, "ReceivedTokens", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDRetrieverReceivedTokens)
				if err := _USDRetriever.contract.UnpackLog(event, "ReceivedTokens", log); err != nil {
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

// ParseReceivedTokens is a log parse operation binding the contract event 0x2946de6c4ec03d8d15126164a7c0da68d7c6835173e41827a7a715f8becb07a8.
//
// Solidity: event ReceivedTokens(address indexed from, uint256 amount)
func (_USDRetriever *USDRetrieverFilterer) ParseReceivedTokens(log types.Log) (*USDRetrieverReceivedTokens, error) {
	event := new(USDRetrieverReceivedTokens)
	if err := _USDRetriever.contract.UnpackLog(event, "ReceivedTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// USDRetrieverTransferTokensIterator is returned from FilterTransferTokens and is used to iterate over the raw logs and unpacked data for TransferTokens events raised by the USDRetriever contract.
type USDRetrieverTransferTokensIterator struct {
	Event *USDRetrieverTransferTokens // Event containing the contract specifics and raw log

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
func (it *USDRetrieverTransferTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDRetrieverTransferTokens)
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
		it.Event = new(USDRetrieverTransferTokens)
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
func (it *USDRetrieverTransferTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDRetrieverTransferTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDRetrieverTransferTokens represents a TransferTokens event raised by the USDRetriever contract.
type USDRetrieverTransferTokens struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferTokens is a free log retrieval operation binding the contract event 0x7cbfb8a5c69722670db81365c97149301df8bf0b0afe48f73787d6a6a4385954.
//
// Solidity: event TransferTokens(address indexed to, uint256 amount)
func (_USDRetriever *USDRetrieverFilterer) FilterTransferTokens(opts *bind.FilterOpts, to []common.Address) (*USDRetrieverTransferTokensIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USDRetriever.contract.FilterLogs(opts, "TransferTokens", toRule)
	if err != nil {
		return nil, err
	}
	return &USDRetrieverTransferTokensIterator{contract: _USDRetriever.contract, event: "TransferTokens", logs: logs, sub: sub}, nil
}

// WatchTransferTokens is a free log subscription operation binding the contract event 0x7cbfb8a5c69722670db81365c97149301df8bf0b0afe48f73787d6a6a4385954.
//
// Solidity: event TransferTokens(address indexed to, uint256 amount)
func (_USDRetriever *USDRetrieverFilterer) WatchTransferTokens(opts *bind.WatchOpts, sink chan<- *USDRetrieverTransferTokens, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USDRetriever.contract.WatchLogs(opts, "TransferTokens", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDRetrieverTransferTokens)
				if err := _USDRetriever.contract.UnpackLog(event, "TransferTokens", log); err != nil {
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

// ParseTransferTokens is a log parse operation binding the contract event 0x7cbfb8a5c69722670db81365c97149301df8bf0b0afe48f73787d6a6a4385954.
//
// Solidity: event TransferTokens(address indexed to, uint256 amount)
func (_USDRetriever *USDRetrieverFilterer) ParseTransferTokens(log types.Log) (*USDRetrieverTransferTokens, error) {
	event := new(USDRetrieverTransferTokens)
	if err := _USDRetriever.contract.UnpackLog(event, "TransferTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
