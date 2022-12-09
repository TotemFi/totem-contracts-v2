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

// StakingPoolImplementationForTestMetaData contains all meta data concerning the StakingPoolImplementationForTest contract.
var StakingPoolImplementationForTestMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ApproveTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"stakeAmounts\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"predictions\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"indexOfFirstStake\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"indexOfLastStake\",\"type\":\"uint256\"}],\"name\":\"BatchStake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DistributedBTC\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"PoolActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"PoolDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"PoolDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"PoolLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"PoolMatured\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"PoolSorted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"PoolUnLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ReceivedTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pricePrediction\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"indexOfStake\",\"type\":\"uint256\"}],\"name\":\"Stake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferTokens\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unstake\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakingReturn\",\"type\":\"uint256\"}],\"name\":\"WithdrawStakingReturn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totemPrize\",\"type\":\"uint256\"}],\"name\":\"WithdrawTotemPrize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"wrappedTokenPrize\",\"type\":\"uint256\"}],\"name\":\"WithdrawWrappedTokenPrize\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"swapRouterAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"BUSDContractAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"WrappedTokenContractAddress\",\"type\":\"address\"}],\"name\":\"__WrappedTokenDistributor_initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"averagePricePrediction\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_stakingAmounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_predictions\",\"type\":\"uint256[]\"}],\"name\":\"batchStake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collaborativeRange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collaborativeReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"declareEmergency\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deletePool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"emergentWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"prediction\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_range\",\"type\":\"uint256\"}],\"name\":\"getDifference\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"getEstimatedWrappedTokenForUSD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stakeIndex\",\"type\":\"uint256\"}],\"name\":\"getIndexedPrize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stakeIndex\",\"type\":\"uint256\"}],\"name\":\"getIndexedStakingReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLPS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPathForUSDToWrappedToken\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getPrize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracleContract\",\"type\":\"address\"}],\"name\":\"getPrizeTokenDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getStakingReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenTaxRate\",\"type\":\"uint256\"}],\"name\":\"getStakingTax\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSwapRouter\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUSDBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUSDToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWrappedTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"stakeIndex\",\"type\":\"uint256\"}],\"name\":\"hasUnStaked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"stakeIndex\",\"type\":\"uint256\"}],\"name\":\"indexedClaimReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isAnEmergency\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isDeleted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isEnhancedEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isLocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isMatured\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"launchDate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"launchDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturityTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maturingPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"usdPrizeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prizeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakeApr\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collaborativeReward\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"oracleDecimals\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isEnhancedEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isMatured\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturingPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maturityTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumStakeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracleDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolCreator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolType\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"potentialCollabReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"predictions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stakedBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stakedTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountWithdrawn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastWithdrawalTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pricePrediction\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"difference\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rank\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"prizeRewardWithdrawn\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"didUnstake\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"prizeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"prizeRewardRates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rank\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"percentage\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"usdAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"purchaseWrappedToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardManager\",\"outputs\":[{\"internalType\":\"contractIRewardManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_activationStatus\",\"type\":\"bool\"}],\"name\":\"setActivationStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"setOracleToZero\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[25]\",\"name\":\"addrArray\",\"type\":\"address[25]\"},{\"internalType\":\"uint256[25]\",\"name\":\"indexArray\",\"type\":\"uint256[25]\"}],\"name\":\"setSortedStakers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sizeAllocation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sizeLimitRangeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sortedStakers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pricePrediction\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeApr\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakeTaxRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"stakerAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingPoolImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totemToken\",\"outputs\":[{\"internalType\":\"contractITotemToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_predictionPrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_samePredictionPrizeToken\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_prizePrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_oracleContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_oracleDecimals\",\"type\":\"uint256\"}],\"name\":\"updateMaturingPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"upgradeEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdPrizeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stuckToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"withdrawStuckTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrappedToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrappedTokenSymbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// StakingPoolImplementationForTestABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingPoolImplementationForTestMetaData.ABI instead.
var StakingPoolImplementationForTestABI = StakingPoolImplementationForTestMetaData.ABI

// StakingPoolImplementationForTest is an auto generated Go binding around an Ethereum contract.
type StakingPoolImplementationForTest struct {
	StakingPoolImplementationForTestCaller     // Read-only binding to the contract
	StakingPoolImplementationForTestTransactor // Write-only binding to the contract
	StakingPoolImplementationForTestFilterer   // Log filterer for contract events
}

// StakingPoolImplementationForTestCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingPoolImplementationForTestCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingPoolImplementationForTestTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingPoolImplementationForTestTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingPoolImplementationForTestFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingPoolImplementationForTestFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingPoolImplementationForTestSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingPoolImplementationForTestSession struct {
	Contract     *StakingPoolImplementationForTest // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                     // Call options to use throughout this session
	TransactOpts bind.TransactOpts                 // Transaction auth options to use throughout this session
}

// StakingPoolImplementationForTestCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingPoolImplementationForTestCallerSession struct {
	Contract *StakingPoolImplementationForTestCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                           // Call options to use throughout this session
}

// StakingPoolImplementationForTestTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingPoolImplementationForTestTransactorSession struct {
	Contract     *StakingPoolImplementationForTestTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                           // Transaction auth options to use throughout this session
}

// StakingPoolImplementationForTestRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingPoolImplementationForTestRaw struct {
	Contract *StakingPoolImplementationForTest // Generic contract binding to access the raw methods on
}

// StakingPoolImplementationForTestCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingPoolImplementationForTestCallerRaw struct {
	Contract *StakingPoolImplementationForTestCaller // Generic read-only contract binding to access the raw methods on
}

// StakingPoolImplementationForTestTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingPoolImplementationForTestTransactorRaw struct {
	Contract *StakingPoolImplementationForTestTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingPoolImplementationForTest creates a new instance of StakingPoolImplementationForTest, bound to a specific deployed contract.
func NewStakingPoolImplementationForTest(address common.Address, backend bind.ContractBackend) (*StakingPoolImplementationForTest, error) {
	contract, err := bindStakingPoolImplementationForTest(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationForTest{StakingPoolImplementationForTestCaller: StakingPoolImplementationForTestCaller{contract: contract}, StakingPoolImplementationForTestTransactor: StakingPoolImplementationForTestTransactor{contract: contract}, StakingPoolImplementationForTestFilterer: StakingPoolImplementationForTestFilterer{contract: contract}}, nil
}

// NewStakingPoolImplementationForTestCaller creates a new read-only instance of StakingPoolImplementationForTest, bound to a specific deployed contract.
func NewStakingPoolImplementationForTestCaller(address common.Address, caller bind.ContractCaller) (*StakingPoolImplementationForTestCaller, error) {
	contract, err := bindStakingPoolImplementationForTest(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationForTestCaller{contract: contract}, nil
}

// NewStakingPoolImplementationForTestTransactor creates a new write-only instance of StakingPoolImplementationForTest, bound to a specific deployed contract.
func NewStakingPoolImplementationForTestTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingPoolImplementationForTestTransactor, error) {
	contract, err := bindStakingPoolImplementationForTest(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationForTestTransactor{contract: contract}, nil
}

// NewStakingPoolImplementationForTestFilterer creates a new log filterer instance of StakingPoolImplementationForTest, bound to a specific deployed contract.
func NewStakingPoolImplementationForTestFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingPoolImplementationForTestFilterer, error) {
	contract, err := bindStakingPoolImplementationForTest(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationForTestFilterer{contract: contract}, nil
}

// bindStakingPoolImplementationForTest binds a generic wrapper to an already deployed contract.
func bindStakingPoolImplementationForTest(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingPoolImplementationForTestABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingPoolImplementationForTest.Contract.StakingPoolImplementationForTestCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.StakingPoolImplementationForTestTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.StakingPoolImplementationForTestTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingPoolImplementationForTest.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.contract.Transact(opts, method, params...)
}

// AveragePricePrediction is a free data retrieval call binding the contract method 0x785df1f8.
//
// Solidity: function averagePricePrediction() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) AveragePricePrediction(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "averagePricePrediction")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AveragePricePrediction is a free data retrieval call binding the contract method 0x785df1f8.
//
// Solidity: function averagePricePrediction() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) AveragePricePrediction() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.AveragePricePrediction(&_StakingPoolImplementationForTest.CallOpts)
}

// AveragePricePrediction is a free data retrieval call binding the contract method 0x785df1f8.
//
// Solidity: function averagePricePrediction() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) AveragePricePrediction() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.AveragePricePrediction(&_StakingPoolImplementationForTest.CallOpts)
}

// CollaborativeRange is a free data retrieval call binding the contract method 0xeb850ae1.
//
// Solidity: function collaborativeRange() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) CollaborativeRange(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "collaborativeRange")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollaborativeRange is a free data retrieval call binding the contract method 0xeb850ae1.
//
// Solidity: function collaborativeRange() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) CollaborativeRange() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.CollaborativeRange(&_StakingPoolImplementationForTest.CallOpts)
}

// CollaborativeRange is a free data retrieval call binding the contract method 0xeb850ae1.
//
// Solidity: function collaborativeRange() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) CollaborativeRange() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.CollaborativeRange(&_StakingPoolImplementationForTest.CallOpts)
}

// CollaborativeReward is a free data retrieval call binding the contract method 0xf3d9dc0d.
//
// Solidity: function collaborativeReward() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) CollaborativeReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "collaborativeReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CollaborativeReward is a free data retrieval call binding the contract method 0xf3d9dc0d.
//
// Solidity: function collaborativeReward() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) CollaborativeReward() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.CollaborativeReward(&_StakingPoolImplementationForTest.CallOpts)
}

// CollaborativeReward is a free data retrieval call binding the contract method 0xf3d9dc0d.
//
// Solidity: function collaborativeReward() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) CollaborativeReward() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.CollaborativeReward(&_StakingPoolImplementationForTest.CallOpts)
}

// GetDifference is a free data retrieval call binding the contract method 0x4c7c77fc.
//
// Solidity: function getDifference(uint256 prediction, uint256 _range) view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) GetDifference(opts *bind.CallOpts, prediction *big.Int, _range *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "getDifference", prediction, _range)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDifference is a free data retrieval call binding the contract method 0x4c7c77fc.
//
// Solidity: function getDifference(uint256 prediction, uint256 _range) view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) GetDifference(prediction *big.Int, _range *big.Int) (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.GetDifference(&_StakingPoolImplementationForTest.CallOpts, prediction, _range)
}

// GetDifference is a free data retrieval call binding the contract method 0x4c7c77fc.
//
// Solidity: function getDifference(uint256 prediction, uint256 _range) view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) GetDifference(prediction *big.Int, _range *big.Int) (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.GetDifference(&_StakingPoolImplementationForTest.CallOpts, prediction, _range)
}

// GetEstimatedWrappedTokenForUSD is a free data retrieval call binding the contract method 0x3fcdd82c.
//
// Solidity: function getEstimatedWrappedTokenForUSD(uint256 _amount) view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) GetEstimatedWrappedTokenForUSD(opts *bind.CallOpts, _amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "getEstimatedWrappedTokenForUSD", _amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetEstimatedWrappedTokenForUSD is a free data retrieval call binding the contract method 0x3fcdd82c.
//
// Solidity: function getEstimatedWrappedTokenForUSD(uint256 _amount) view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) GetEstimatedWrappedTokenForUSD(_amount *big.Int) (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.GetEstimatedWrappedTokenForUSD(&_StakingPoolImplementationForTest.CallOpts, _amount)
}

// GetEstimatedWrappedTokenForUSD is a free data retrieval call binding the contract method 0x3fcdd82c.
//
// Solidity: function getEstimatedWrappedTokenForUSD(uint256 _amount) view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) GetEstimatedWrappedTokenForUSD(_amount *big.Int) (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.GetEstimatedWrappedTokenForUSD(&_StakingPoolImplementationForTest.CallOpts, _amount)
}

// GetIndexedPrize is a free data retrieval call binding the contract method 0xb7b30170.
//
// Solidity: function getIndexedPrize(address _staker, uint256 _stakeIndex) view returns(uint256, uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) GetIndexedPrize(opts *bind.CallOpts, _staker common.Address, _stakeIndex *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "getIndexedPrize", _staker, _stakeIndex)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetIndexedPrize is a free data retrieval call binding the contract method 0xb7b30170.
//
// Solidity: function getIndexedPrize(address _staker, uint256 _stakeIndex) view returns(uint256, uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) GetIndexedPrize(_staker common.Address, _stakeIndex *big.Int) (*big.Int, *big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.GetIndexedPrize(&_StakingPoolImplementationForTest.CallOpts, _staker, _stakeIndex)
}

// GetIndexedPrize is a free data retrieval call binding the contract method 0xb7b30170.
//
// Solidity: function getIndexedPrize(address _staker, uint256 _stakeIndex) view returns(uint256, uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) GetIndexedPrize(_staker common.Address, _stakeIndex *big.Int) (*big.Int, *big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.GetIndexedPrize(&_StakingPoolImplementationForTest.CallOpts, _staker, _stakeIndex)
}

// GetIndexedStakingRewardA is a free data retrieval call binding the contract method 0x2342e480.
//
// Solidity: function getIndexedStakingReward(address _staker, uint256 _stakeIndex) view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) GetIndexedStakingRewardA(opts *bind.CallOpts, _staker common.Address, _stakeIndex *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "getIndexedStakingReward", _staker, _stakeIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIndexedStakingRewardA is a free data retrieval call binding the contract method 0x2342e480.
//
// Solidity: function getIndexedStakingReward(address _staker, uint256 _stakeIndex) view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) GetIndexedStakingRewardA(_staker common.Address, _stakeIndex *big.Int) (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.GetIndexedStakingRewardA(&_StakingPoolImplementationForTest.CallOpts, _staker, _stakeIndex)
}

// GetIndexedStakingRewardA is a free data retrieval call binding the contract method 0x2342e480.
//
// Solidity: function getIndexedStakingReward(address _staker, uint256 _stakeIndex) view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) GetIndexedStakingRewardA(_staker common.Address, _stakeIndex *big.Int) (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.GetIndexedStakingRewardA(&_StakingPoolImplementationForTest.CallOpts, _staker, _stakeIndex)
}

// GetLPS is a free data retrieval call binding the contract method 0x9e3dd80f.
//
// Solidity: function getLPS() view returns(uint256, uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) GetLPS(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "getLPS")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetLPS is a free data retrieval call binding the contract method 0x9e3dd80f.
//
// Solidity: function getLPS() view returns(uint256, uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) GetLPS() (*big.Int, *big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.GetLPS(&_StakingPoolImplementationForTest.CallOpts)
}

// GetLPS is a free data retrieval call binding the contract method 0x9e3dd80f.
//
// Solidity: function getLPS() view returns(uint256, uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) GetLPS() (*big.Int, *big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.GetLPS(&_StakingPoolImplementationForTest.CallOpts)
}

// GetPathForUSDToWrappedToken is a free data retrieval call binding the contract method 0xa35ef3a7.
//
// Solidity: function getPathForUSDToWrappedToken() view returns(address[])
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) GetPathForUSDToWrappedToken(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "getPathForUSDToWrappedToken")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPathForUSDToWrappedToken is a free data retrieval call binding the contract method 0xa35ef3a7.
//
// Solidity: function getPathForUSDToWrappedToken() view returns(address[])
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) GetPathForUSDToWrappedToken() ([]common.Address, error) {
	return _StakingPoolImplementationForTest.Contract.GetPathForUSDToWrappedToken(&_StakingPoolImplementationForTest.CallOpts)
}

// GetPathForUSDToWrappedToken is a free data retrieval call binding the contract method 0xa35ef3a7.
//
// Solidity: function getPathForUSDToWrappedToken() view returns(address[])
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) GetPathForUSDToWrappedToken() ([]common.Address, error) {
	return _StakingPoolImplementationForTest.Contract.GetPathForUSDToWrappedToken(&_StakingPoolImplementationForTest.CallOpts)
}

// GetPrize is a free data retrieval call binding the contract method 0x2d5cd1d6.
//
// Solidity: function getPrize(address _staker) view returns(uint256, uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) GetPrize(opts *bind.CallOpts, _staker common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "getPrize", _staker)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPrize is a free data retrieval call binding the contract method 0x2d5cd1d6.
//
// Solidity: function getPrize(address _staker) view returns(uint256, uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) GetPrize(_staker common.Address) (*big.Int, *big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.GetPrize(&_StakingPoolImplementationForTest.CallOpts, _staker)
}

// GetPrize is a free data retrieval call binding the contract method 0x2d5cd1d6.
//
// Solidity: function getPrize(address _staker) view returns(uint256, uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) GetPrize(_staker common.Address) (*big.Int, *big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.GetPrize(&_StakingPoolImplementationForTest.CallOpts, _staker)
}

// GetPrizeTokenDecimals is a free data retrieval call binding the contract method 0xcee0c889.
//
// Solidity: function getPrizeTokenDecimals(address _oracleContract) view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) GetPrizeTokenDecimals(opts *bind.CallOpts, _oracleContract common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "getPrizeTokenDecimals", _oracleContract)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPrizeTokenDecimals is a free data retrieval call binding the contract method 0xcee0c889.
//
// Solidity: function getPrizeTokenDecimals(address _oracleContract) view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) GetPrizeTokenDecimals(_oracleContract common.Address) (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.GetPrizeTokenDecimals(&_StakingPoolImplementationForTest.CallOpts, _oracleContract)
}

// GetPrizeTokenDecimals is a free data retrieval call binding the contract method 0xcee0c889.
//
// Solidity: function getPrizeTokenDecimals(address _oracleContract) view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) GetPrizeTokenDecimals(_oracleContract common.Address) (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.GetPrizeTokenDecimals(&_StakingPoolImplementationForTest.CallOpts, _oracleContract)
}

// GetStakers is a free data retrieval call binding the contract method 0x43352d61.
//
// Solidity: function getStakers() view returns(address[], uint256[])
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) GetStakers(opts *bind.CallOpts) ([]common.Address, []*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "getStakers")

	if err != nil {
		return *new([]common.Address), *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return out0, out1, err

}

// GetStakers is a free data retrieval call binding the contract method 0x43352d61.
//
// Solidity: function getStakers() view returns(address[], uint256[])
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) GetStakers() ([]common.Address, []*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.GetStakers(&_StakingPoolImplementationForTest.CallOpts)
}

// GetStakers is a free data retrieval call binding the contract method 0x43352d61.
//
// Solidity: function getStakers() view returns(address[], uint256[])
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) GetStakers() ([]common.Address, []*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.GetStakers(&_StakingPoolImplementationForTest.CallOpts)
}

// GetStakingRewardA is a free data retrieval call binding the contract method 0x78f4d413.
//
// Solidity: function getStakingReward(address _staker) view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) GetStakingRewardA(opts *bind.CallOpts, _staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "getStakingReward", _staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakingRewardA is a free data retrieval call binding the contract method 0x78f4d413.
//
// Solidity: function getStakingReward(address _staker) view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) GetStakingRewardA(_staker common.Address) (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.GetStakingRewardA(&_StakingPoolImplementationForTest.CallOpts, _staker)
}

// GetStakingRewardA is a free data retrieval call binding the contract method 0x78f4d413.
//
// Solidity: function getStakingReward(address _staker) view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) GetStakingRewardA(_staker common.Address) (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.GetStakingRewardA(&_StakingPoolImplementationForTest.CallOpts, _staker)
}

// GetStakingTax is a free data retrieval call binding the contract method 0x395b3dac.
//
// Solidity: function getStakingTax(uint256 amount, uint256 tokenTaxRate) view returns(uint256, uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) GetStakingTax(opts *bind.CallOpts, amount *big.Int, tokenTaxRate *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "getStakingTax", amount, tokenTaxRate)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetStakingTax is a free data retrieval call binding the contract method 0x395b3dac.
//
// Solidity: function getStakingTax(uint256 amount, uint256 tokenTaxRate) view returns(uint256, uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) GetStakingTax(amount *big.Int, tokenTaxRate *big.Int) (*big.Int, *big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.GetStakingTax(&_StakingPoolImplementationForTest.CallOpts, amount, tokenTaxRate)
}

// GetStakingTax is a free data retrieval call binding the contract method 0x395b3dac.
//
// Solidity: function getStakingTax(uint256 amount, uint256 tokenTaxRate) view returns(uint256, uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) GetStakingTax(amount *big.Int, tokenTaxRate *big.Int) (*big.Int, *big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.GetStakingTax(&_StakingPoolImplementationForTest.CallOpts, amount, tokenTaxRate)
}

// GetSwapRouter is a free data retrieval call binding the contract method 0x725c9c49.
//
// Solidity: function getSwapRouter() view returns(address)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) GetSwapRouter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "getSwapRouter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetSwapRouter is a free data retrieval call binding the contract method 0x725c9c49.
//
// Solidity: function getSwapRouter() view returns(address)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) GetSwapRouter() (common.Address, error) {
	return _StakingPoolImplementationForTest.Contract.GetSwapRouter(&_StakingPoolImplementationForTest.CallOpts)
}

// GetSwapRouter is a free data retrieval call binding the contract method 0x725c9c49.
//
// Solidity: function getSwapRouter() view returns(address)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) GetSwapRouter() (common.Address, error) {
	return _StakingPoolImplementationForTest.Contract.GetSwapRouter(&_StakingPoolImplementationForTest.CallOpts)
}

// GetUSDBalance is a free data retrieval call binding the contract method 0xc1201054.
//
// Solidity: function getUSDBalance() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) GetUSDBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "getUSDBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUSDBalance is a free data retrieval call binding the contract method 0xc1201054.
//
// Solidity: function getUSDBalance() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) GetUSDBalance() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.GetUSDBalance(&_StakingPoolImplementationForTest.CallOpts)
}

// GetUSDBalance is a free data retrieval call binding the contract method 0xc1201054.
//
// Solidity: function getUSDBalance() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) GetUSDBalance() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.GetUSDBalance(&_StakingPoolImplementationForTest.CallOpts)
}

// GetUSDToken is a free data retrieval call binding the contract method 0xf5e96fc9.
//
// Solidity: function getUSDToken() view returns(address)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) GetUSDToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "getUSDToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUSDToken is a free data retrieval call binding the contract method 0xf5e96fc9.
//
// Solidity: function getUSDToken() view returns(address)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) GetUSDToken() (common.Address, error) {
	return _StakingPoolImplementationForTest.Contract.GetUSDToken(&_StakingPoolImplementationForTest.CallOpts)
}

// GetUSDToken is a free data retrieval call binding the contract method 0xf5e96fc9.
//
// Solidity: function getUSDToken() view returns(address)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) GetUSDToken() (common.Address, error) {
	return _StakingPoolImplementationForTest.Contract.GetUSDToken(&_StakingPoolImplementationForTest.CallOpts)
}

// GetWrappedTokenBalance is a free data retrieval call binding the contract method 0xf8e829c3.
//
// Solidity: function getWrappedTokenBalance() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) GetWrappedTokenBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "getWrappedTokenBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWrappedTokenBalance is a free data retrieval call binding the contract method 0xf8e829c3.
//
// Solidity: function getWrappedTokenBalance() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) GetWrappedTokenBalance() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.GetWrappedTokenBalance(&_StakingPoolImplementationForTest.CallOpts)
}

// GetWrappedTokenBalance is a free data retrieval call binding the contract method 0xf8e829c3.
//
// Solidity: function getWrappedTokenBalance() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) GetWrappedTokenBalance() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.GetWrappedTokenBalance(&_StakingPoolImplementationForTest.CallOpts)
}

// HasUnStaked is a free data retrieval call binding the contract method 0x45e91319.
//
// Solidity: function hasUnStaked(address staker, uint256 stakeIndex) view returns(bool)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) HasUnStaked(opts *bind.CallOpts, staker common.Address, stakeIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "hasUnStaked", staker, stakeIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasUnStaked is a free data retrieval call binding the contract method 0x45e91319.
//
// Solidity: function hasUnStaked(address staker, uint256 stakeIndex) view returns(bool)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) HasUnStaked(staker common.Address, stakeIndex *big.Int) (bool, error) {
	return _StakingPoolImplementationForTest.Contract.HasUnStaked(&_StakingPoolImplementationForTest.CallOpts, staker, stakeIndex)
}

// HasUnStaked is a free data retrieval call binding the contract method 0x45e91319.
//
// Solidity: function hasUnStaked(address staker, uint256 stakeIndex) view returns(bool)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) HasUnStaked(staker common.Address, stakeIndex *big.Int) (bool, error) {
	return _StakingPoolImplementationForTest.Contract.HasUnStaked(&_StakingPoolImplementationForTest.CallOpts, staker, stakeIndex)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) IsActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "isActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) IsActive() (bool, error) {
	return _StakingPoolImplementationForTest.Contract.IsActive(&_StakingPoolImplementationForTest.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) IsActive() (bool, error) {
	return _StakingPoolImplementationForTest.Contract.IsActive(&_StakingPoolImplementationForTest.CallOpts)
}

// IsAnEmergency is a free data retrieval call binding the contract method 0xa08b2c79.
//
// Solidity: function isAnEmergency() view returns(bool)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) IsAnEmergency(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "isAnEmergency")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAnEmergency is a free data retrieval call binding the contract method 0xa08b2c79.
//
// Solidity: function isAnEmergency() view returns(bool)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) IsAnEmergency() (bool, error) {
	return _StakingPoolImplementationForTest.Contract.IsAnEmergency(&_StakingPoolImplementationForTest.CallOpts)
}

// IsAnEmergency is a free data retrieval call binding the contract method 0xa08b2c79.
//
// Solidity: function isAnEmergency() view returns(bool)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) IsAnEmergency() (bool, error) {
	return _StakingPoolImplementationForTest.Contract.IsAnEmergency(&_StakingPoolImplementationForTest.CallOpts)
}

// IsDeleted is a free data retrieval call binding the contract method 0xd7efb6b7.
//
// Solidity: function isDeleted() view returns(bool)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) IsDeleted(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "isDeleted")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDeleted is a free data retrieval call binding the contract method 0xd7efb6b7.
//
// Solidity: function isDeleted() view returns(bool)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) IsDeleted() (bool, error) {
	return _StakingPoolImplementationForTest.Contract.IsDeleted(&_StakingPoolImplementationForTest.CallOpts)
}

// IsDeleted is a free data retrieval call binding the contract method 0xd7efb6b7.
//
// Solidity: function isDeleted() view returns(bool)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) IsDeleted() (bool, error) {
	return _StakingPoolImplementationForTest.Contract.IsDeleted(&_StakingPoolImplementationForTest.CallOpts)
}

// IsEnhancedEnabled is a free data retrieval call binding the contract method 0x4bc4eeb2.
//
// Solidity: function isEnhancedEnabled() view returns(bool)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) IsEnhancedEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "isEnhancedEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsEnhancedEnabled is a free data retrieval call binding the contract method 0x4bc4eeb2.
//
// Solidity: function isEnhancedEnabled() view returns(bool)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) IsEnhancedEnabled() (bool, error) {
	return _StakingPoolImplementationForTest.Contract.IsEnhancedEnabled(&_StakingPoolImplementationForTest.CallOpts)
}

// IsEnhancedEnabled is a free data retrieval call binding the contract method 0x4bc4eeb2.
//
// Solidity: function isEnhancedEnabled() view returns(bool)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) IsEnhancedEnabled() (bool, error) {
	return _StakingPoolImplementationForTest.Contract.IsEnhancedEnabled(&_StakingPoolImplementationForTest.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) IsLocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "isLocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) IsLocked() (bool, error) {
	return _StakingPoolImplementationForTest.Contract.IsLocked(&_StakingPoolImplementationForTest.CallOpts)
}

// IsLocked is a free data retrieval call binding the contract method 0xa4e2d634.
//
// Solidity: function isLocked() view returns(bool)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) IsLocked() (bool, error) {
	return _StakingPoolImplementationForTest.Contract.IsLocked(&_StakingPoolImplementationForTest.CallOpts)
}

// IsMatured is a free data retrieval call binding the contract method 0x7f2b6a0d.
//
// Solidity: function isMatured() view returns(bool)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) IsMatured(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "isMatured")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMatured is a free data retrieval call binding the contract method 0x7f2b6a0d.
//
// Solidity: function isMatured() view returns(bool)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) IsMatured() (bool, error) {
	return _StakingPoolImplementationForTest.Contract.IsMatured(&_StakingPoolImplementationForTest.CallOpts)
}

// IsMatured is a free data retrieval call binding the contract method 0x7f2b6a0d.
//
// Solidity: function isMatured() view returns(bool)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) IsMatured() (bool, error) {
	return _StakingPoolImplementationForTest.Contract.IsMatured(&_StakingPoolImplementationForTest.CallOpts)
}

// LaunchDate is a free data retrieval call binding the contract method 0xf8eeed62.
//
// Solidity: function launchDate() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) LaunchDate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "launchDate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LaunchDate is a free data retrieval call binding the contract method 0xf8eeed62.
//
// Solidity: function launchDate() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) LaunchDate() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.LaunchDate(&_StakingPoolImplementationForTest.CallOpts)
}

// LaunchDate is a free data retrieval call binding the contract method 0xf8eeed62.
//
// Solidity: function launchDate() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) LaunchDate() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.LaunchDate(&_StakingPoolImplementationForTest.CallOpts)
}

// LockTime is a free data retrieval call binding the contract method 0x0d668087.
//
// Solidity: function lockTime() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) LockTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "lockTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockTime is a free data retrieval call binding the contract method 0x0d668087.
//
// Solidity: function lockTime() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) LockTime() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.LockTime(&_StakingPoolImplementationForTest.CallOpts)
}

// LockTime is a free data retrieval call binding the contract method 0x0d668087.
//
// Solidity: function lockTime() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) LockTime() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.LockTime(&_StakingPoolImplementationForTest.CallOpts)
}

// Lps is a free data retrieval call binding the contract method 0xe0818669.
//
// Solidity: function lps() view returns(uint256 launchDate, uint256 lockTime, uint256 maturityTime, uint256 maturingPrice, uint256 usdPrizeAmount, uint256 prizeAmount, uint256 stakeApr, uint256 collaborativeReward, uint256 oracleDecimals, bool isEnhancedEnabled, bool isMatured)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) Lps(opts *bind.CallOpts) (struct {
	LaunchDate          *big.Int
	LockTime            *big.Int
	MaturityTime        *big.Int
	MaturingPrice       *big.Int
	UsdPrizeAmount      *big.Int
	PrizeAmount         *big.Int
	StakeApr            *big.Int
	CollaborativeReward *big.Int
	OracleDecimals      *big.Int
	IsEnhancedEnabled   bool
	IsMatured           bool
}, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "lps")

	outstruct := new(struct {
		LaunchDate          *big.Int
		LockTime            *big.Int
		MaturityTime        *big.Int
		MaturingPrice       *big.Int
		UsdPrizeAmount      *big.Int
		PrizeAmount         *big.Int
		StakeApr            *big.Int
		CollaborativeReward *big.Int
		OracleDecimals      *big.Int
		IsEnhancedEnabled   bool
		IsMatured           bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LaunchDate = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.LockTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MaturityTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.MaturingPrice = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.UsdPrizeAmount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.PrizeAmount = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.StakeApr = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.CollaborativeReward = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.OracleDecimals = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.IsEnhancedEnabled = *abi.ConvertType(out[9], new(bool)).(*bool)
	outstruct.IsMatured = *abi.ConvertType(out[10], new(bool)).(*bool)

	return *outstruct, err

}

// Lps is a free data retrieval call binding the contract method 0xe0818669.
//
// Solidity: function lps() view returns(uint256 launchDate, uint256 lockTime, uint256 maturityTime, uint256 maturingPrice, uint256 usdPrizeAmount, uint256 prizeAmount, uint256 stakeApr, uint256 collaborativeReward, uint256 oracleDecimals, bool isEnhancedEnabled, bool isMatured)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) Lps() (struct {
	LaunchDate          *big.Int
	LockTime            *big.Int
	MaturityTime        *big.Int
	MaturingPrice       *big.Int
	UsdPrizeAmount      *big.Int
	PrizeAmount         *big.Int
	StakeApr            *big.Int
	CollaborativeReward *big.Int
	OracleDecimals      *big.Int
	IsEnhancedEnabled   bool
	IsMatured           bool
}, error) {
	return _StakingPoolImplementationForTest.Contract.Lps(&_StakingPoolImplementationForTest.CallOpts)
}

// Lps is a free data retrieval call binding the contract method 0xe0818669.
//
// Solidity: function lps() view returns(uint256 launchDate, uint256 lockTime, uint256 maturityTime, uint256 maturingPrice, uint256 usdPrizeAmount, uint256 prizeAmount, uint256 stakeApr, uint256 collaborativeReward, uint256 oracleDecimals, bool isEnhancedEnabled, bool isMatured)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) Lps() (struct {
	LaunchDate          *big.Int
	LockTime            *big.Int
	MaturityTime        *big.Int
	MaturingPrice       *big.Int
	UsdPrizeAmount      *big.Int
	PrizeAmount         *big.Int
	StakeApr            *big.Int
	CollaborativeReward *big.Int
	OracleDecimals      *big.Int
	IsEnhancedEnabled   bool
	IsMatured           bool
}, error) {
	return _StakingPoolImplementationForTest.Contract.Lps(&_StakingPoolImplementationForTest.CallOpts)
}

// MaturingPrice is a free data retrieval call binding the contract method 0xd025188b.
//
// Solidity: function maturingPrice() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) MaturingPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "maturingPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaturingPrice is a free data retrieval call binding the contract method 0xd025188b.
//
// Solidity: function maturingPrice() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) MaturingPrice() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.MaturingPrice(&_StakingPoolImplementationForTest.CallOpts)
}

// MaturingPrice is a free data retrieval call binding the contract method 0xd025188b.
//
// Solidity: function maturingPrice() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) MaturingPrice() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.MaturingPrice(&_StakingPoolImplementationForTest.CallOpts)
}

// MaturityTime is a free data retrieval call binding the contract method 0x4e8bfdaa.
//
// Solidity: function maturityTime() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) MaturityTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "maturityTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaturityTime is a free data retrieval call binding the contract method 0x4e8bfdaa.
//
// Solidity: function maturityTime() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) MaturityTime() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.MaturityTime(&_StakingPoolImplementationForTest.CallOpts)
}

// MaturityTime is a free data retrieval call binding the contract method 0x4e8bfdaa.
//
// Solidity: function maturityTime() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) MaturityTime() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.MaturityTime(&_StakingPoolImplementationForTest.CallOpts)
}

// MinimumStakeAmount is a free data retrieval call binding the contract method 0x6b036f45.
//
// Solidity: function minimumStakeAmount() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) MinimumStakeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "minimumStakeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumStakeAmount is a free data retrieval call binding the contract method 0x6b036f45.
//
// Solidity: function minimumStakeAmount() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) MinimumStakeAmount() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.MinimumStakeAmount(&_StakingPoolImplementationForTest.CallOpts)
}

// MinimumStakeAmount is a free data retrieval call binding the contract method 0x6b036f45.
//
// Solidity: function minimumStakeAmount() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) MinimumStakeAmount() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.MinimumStakeAmount(&_StakingPoolImplementationForTest.CallOpts)
}

// OracleContract is a free data retrieval call binding the contract method 0xbece7532.
//
// Solidity: function oracleContract() view returns(address)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) OracleContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "oracleContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OracleContract is a free data retrieval call binding the contract method 0xbece7532.
//
// Solidity: function oracleContract() view returns(address)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) OracleContract() (common.Address, error) {
	return _StakingPoolImplementationForTest.Contract.OracleContract(&_StakingPoolImplementationForTest.CallOpts)
}

// OracleContract is a free data retrieval call binding the contract method 0xbece7532.
//
// Solidity: function oracleContract() view returns(address)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) OracleContract() (common.Address, error) {
	return _StakingPoolImplementationForTest.Contract.OracleContract(&_StakingPoolImplementationForTest.CallOpts)
}

// OracleDecimals is a free data retrieval call binding the contract method 0xe68b52e7.
//
// Solidity: function oracleDecimals() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) OracleDecimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "oracleDecimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OracleDecimals is a free data retrieval call binding the contract method 0xe68b52e7.
//
// Solidity: function oracleDecimals() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) OracleDecimals() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.OracleDecimals(&_StakingPoolImplementationForTest.CallOpts)
}

// OracleDecimals is a free data retrieval call binding the contract method 0xe68b52e7.
//
// Solidity: function oracleDecimals() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) OracleDecimals() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.OracleDecimals(&_StakingPoolImplementationForTest.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) Owner() (common.Address, error) {
	return _StakingPoolImplementationForTest.Contract.Owner(&_StakingPoolImplementationForTest.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) Owner() (common.Address, error) {
	return _StakingPoolImplementationForTest.Contract.Owner(&_StakingPoolImplementationForTest.CallOpts)
}

// PoolCreator is a free data retrieval call binding the contract method 0xc6c1decd.
//
// Solidity: function poolCreator() view returns(address)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) PoolCreator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "poolCreator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolCreator is a free data retrieval call binding the contract method 0xc6c1decd.
//
// Solidity: function poolCreator() view returns(address)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) PoolCreator() (common.Address, error) {
	return _StakingPoolImplementationForTest.Contract.PoolCreator(&_StakingPoolImplementationForTest.CallOpts)
}

// PoolCreator is a free data retrieval call binding the contract method 0xc6c1decd.
//
// Solidity: function poolCreator() view returns(address)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) PoolCreator() (common.Address, error) {
	return _StakingPoolImplementationForTest.Contract.PoolCreator(&_StakingPoolImplementationForTest.CallOpts)
}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() view returns(string)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) PoolType(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "poolType")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() view returns(string)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) PoolType() (string, error) {
	return _StakingPoolImplementationForTest.Contract.PoolType(&_StakingPoolImplementationForTest.CallOpts)
}

// PoolType is a free data retrieval call binding the contract method 0xb1dd61b6.
//
// Solidity: function poolType() view returns(string)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) PoolType() (string, error) {
	return _StakingPoolImplementationForTest.Contract.PoolType(&_StakingPoolImplementationForTest.CallOpts)
}

// PotentialCollabReward is a free data retrieval call binding the contract method 0xe425df8c.
//
// Solidity: function potentialCollabReward() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) PotentialCollabReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "potentialCollabReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PotentialCollabReward is a free data retrieval call binding the contract method 0xe425df8c.
//
// Solidity: function potentialCollabReward() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) PotentialCollabReward() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.PotentialCollabReward(&_StakingPoolImplementationForTest.CallOpts)
}

// PotentialCollabReward is a free data retrieval call binding the contract method 0xe425df8c.
//
// Solidity: function potentialCollabReward() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) PotentialCollabReward() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.PotentialCollabReward(&_StakingPoolImplementationForTest.CallOpts)
}

// Predictions is a free data retrieval call binding the contract method 0x850fa7cb.
//
// Solidity: function predictions(address , uint256 ) view returns(uint256 stakedBalance, uint256 stakedTime, uint256 amountWithdrawn, uint256 lastWithdrawalTime, uint256 pricePrediction, uint256 difference, uint256 rank, bool prizeRewardWithdrawn, bool didUnstake)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) Predictions(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	StakedBalance        *big.Int
	StakedTime           *big.Int
	AmountWithdrawn      *big.Int
	LastWithdrawalTime   *big.Int
	PricePrediction      *big.Int
	Difference           *big.Int
	Rank                 *big.Int
	PrizeRewardWithdrawn bool
	DidUnstake           bool
}, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "predictions", arg0, arg1)

	outstruct := new(struct {
		StakedBalance        *big.Int
		StakedTime           *big.Int
		AmountWithdrawn      *big.Int
		LastWithdrawalTime   *big.Int
		PricePrediction      *big.Int
		Difference           *big.Int
		Rank                 *big.Int
		PrizeRewardWithdrawn bool
		DidUnstake           bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StakedBalance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StakedTime = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AmountWithdrawn = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LastWithdrawalTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.PricePrediction = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Difference = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Rank = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.PrizeRewardWithdrawn = *abi.ConvertType(out[7], new(bool)).(*bool)
	outstruct.DidUnstake = *abi.ConvertType(out[8], new(bool)).(*bool)

	return *outstruct, err

}

// Predictions is a free data retrieval call binding the contract method 0x850fa7cb.
//
// Solidity: function predictions(address , uint256 ) view returns(uint256 stakedBalance, uint256 stakedTime, uint256 amountWithdrawn, uint256 lastWithdrawalTime, uint256 pricePrediction, uint256 difference, uint256 rank, bool prizeRewardWithdrawn, bool didUnstake)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) Predictions(arg0 common.Address, arg1 *big.Int) (struct {
	StakedBalance        *big.Int
	StakedTime           *big.Int
	AmountWithdrawn      *big.Int
	LastWithdrawalTime   *big.Int
	PricePrediction      *big.Int
	Difference           *big.Int
	Rank                 *big.Int
	PrizeRewardWithdrawn bool
	DidUnstake           bool
}, error) {
	return _StakingPoolImplementationForTest.Contract.Predictions(&_StakingPoolImplementationForTest.CallOpts, arg0, arg1)
}

// Predictions is a free data retrieval call binding the contract method 0x850fa7cb.
//
// Solidity: function predictions(address , uint256 ) view returns(uint256 stakedBalance, uint256 stakedTime, uint256 amountWithdrawn, uint256 lastWithdrawalTime, uint256 pricePrediction, uint256 difference, uint256 rank, bool prizeRewardWithdrawn, bool didUnstake)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) Predictions(arg0 common.Address, arg1 *big.Int) (struct {
	StakedBalance        *big.Int
	StakedTime           *big.Int
	AmountWithdrawn      *big.Int
	LastWithdrawalTime   *big.Int
	PricePrediction      *big.Int
	Difference           *big.Int
	Rank                 *big.Int
	PrizeRewardWithdrawn bool
	DidUnstake           bool
}, error) {
	return _StakingPoolImplementationForTest.Contract.Predictions(&_StakingPoolImplementationForTest.CallOpts, arg0, arg1)
}

// PrizeAmount is a free data retrieval call binding the contract method 0x785fa627.
//
// Solidity: function prizeAmount() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) PrizeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "prizeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrizeAmount is a free data retrieval call binding the contract method 0x785fa627.
//
// Solidity: function prizeAmount() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) PrizeAmount() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.PrizeAmount(&_StakingPoolImplementationForTest.CallOpts)
}

// PrizeAmount is a free data retrieval call binding the contract method 0x785fa627.
//
// Solidity: function prizeAmount() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) PrizeAmount() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.PrizeAmount(&_StakingPoolImplementationForTest.CallOpts)
}

// PrizeRewardRates is a free data retrieval call binding the contract method 0xf1b17a8b.
//
// Solidity: function prizeRewardRates(uint256 ) view returns(uint256 rank, uint256 percentage)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) PrizeRewardRates(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Rank       *big.Int
	Percentage *big.Int
}, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "prizeRewardRates", arg0)

	outstruct := new(struct {
		Rank       *big.Int
		Percentage *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Rank = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Percentage = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PrizeRewardRates is a free data retrieval call binding the contract method 0xf1b17a8b.
//
// Solidity: function prizeRewardRates(uint256 ) view returns(uint256 rank, uint256 percentage)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) PrizeRewardRates(arg0 *big.Int) (struct {
	Rank       *big.Int
	Percentage *big.Int
}, error) {
	return _StakingPoolImplementationForTest.Contract.PrizeRewardRates(&_StakingPoolImplementationForTest.CallOpts, arg0)
}

// PrizeRewardRates is a free data retrieval call binding the contract method 0xf1b17a8b.
//
// Solidity: function prizeRewardRates(uint256 ) view returns(uint256 rank, uint256 percentage)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) PrizeRewardRates(arg0 *big.Int) (struct {
	Rank       *big.Int
	Percentage *big.Int
}, error) {
	return _StakingPoolImplementationForTest.Contract.PrizeRewardRates(&_StakingPoolImplementationForTest.CallOpts, arg0)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) RewardManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "rewardManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) RewardManager() (common.Address, error) {
	return _StakingPoolImplementationForTest.Contract.RewardManager(&_StakingPoolImplementationForTest.CallOpts)
}

// RewardManager is a free data retrieval call binding the contract method 0x0f4ef8a6.
//
// Solidity: function rewardManager() view returns(address)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) RewardManager() (common.Address, error) {
	return _StakingPoolImplementationForTest.Contract.RewardManager(&_StakingPoolImplementationForTest.CallOpts)
}

// SizeAllocation is a free data retrieval call binding the contract method 0x51eacbf0.
//
// Solidity: function sizeAllocation() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) SizeAllocation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "sizeAllocation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SizeAllocation is a free data retrieval call binding the contract method 0x51eacbf0.
//
// Solidity: function sizeAllocation() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) SizeAllocation() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.SizeAllocation(&_StakingPoolImplementationForTest.CallOpts)
}

// SizeAllocation is a free data retrieval call binding the contract method 0x51eacbf0.
//
// Solidity: function sizeAllocation() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) SizeAllocation() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.SizeAllocation(&_StakingPoolImplementationForTest.CallOpts)
}

// SizeLimitRangeRate is a free data retrieval call binding the contract method 0x95c2ba47.
//
// Solidity: function sizeLimitRangeRate() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) SizeLimitRangeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "sizeLimitRangeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SizeLimitRangeRate is a free data retrieval call binding the contract method 0x95c2ba47.
//
// Solidity: function sizeLimitRangeRate() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) SizeLimitRangeRate() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.SizeLimitRangeRate(&_StakingPoolImplementationForTest.CallOpts)
}

// SizeLimitRangeRate is a free data retrieval call binding the contract method 0x95c2ba47.
//
// Solidity: function sizeLimitRangeRate() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) SizeLimitRangeRate() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.SizeLimitRangeRate(&_StakingPoolImplementationForTest.CallOpts)
}

// SortedStakers is a free data retrieval call binding the contract method 0xde09ee20.
//
// Solidity: function sortedStakers(uint256 ) view returns(address stakerAddress, uint256 index)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) SortedStakers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StakerAddress common.Address
	Index         *big.Int
}, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "sortedStakers", arg0)

	outstruct := new(struct {
		StakerAddress common.Address
		Index         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StakerAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Index = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// SortedStakers is a free data retrieval call binding the contract method 0xde09ee20.
//
// Solidity: function sortedStakers(uint256 ) view returns(address stakerAddress, uint256 index)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) SortedStakers(arg0 *big.Int) (struct {
	StakerAddress common.Address
	Index         *big.Int
}, error) {
	return _StakingPoolImplementationForTest.Contract.SortedStakers(&_StakingPoolImplementationForTest.CallOpts, arg0)
}

// SortedStakers is a free data retrieval call binding the contract method 0xde09ee20.
//
// Solidity: function sortedStakers(uint256 ) view returns(address stakerAddress, uint256 index)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) SortedStakers(arg0 *big.Int) (struct {
	StakerAddress common.Address
	Index         *big.Int
}, error) {
	return _StakingPoolImplementationForTest.Contract.SortedStakers(&_StakingPoolImplementationForTest.CallOpts, arg0)
}

// StakeApr is a free data retrieval call binding the contract method 0xaa7bcb57.
//
// Solidity: function stakeApr() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) StakeApr(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "stakeApr")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeApr is a free data retrieval call binding the contract method 0xaa7bcb57.
//
// Solidity: function stakeApr() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) StakeApr() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.StakeApr(&_StakingPoolImplementationForTest.CallOpts)
}

// StakeApr is a free data retrieval call binding the contract method 0xaa7bcb57.
//
// Solidity: function stakeApr() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) StakeApr() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.StakeApr(&_StakingPoolImplementationForTest.CallOpts)
}

// StakeTaxRate is a free data retrieval call binding the contract method 0x6847d0c5.
//
// Solidity: function stakeTaxRate() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) StakeTaxRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "stakeTaxRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeTaxRate is a free data retrieval call binding the contract method 0x6847d0c5.
//
// Solidity: function stakeTaxRate() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) StakeTaxRate() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.StakeTaxRate(&_StakingPoolImplementationForTest.CallOpts)
}

// StakeTaxRate is a free data retrieval call binding the contract method 0x6847d0c5.
//
// Solidity: function stakeTaxRate() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) StakeTaxRate() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.StakeTaxRate(&_StakingPoolImplementationForTest.CallOpts)
}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 ) view returns(address stakerAddress, uint256 index)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) Stakers(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StakerAddress common.Address
	Index         *big.Int
}, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "stakers", arg0)

	outstruct := new(struct {
		StakerAddress common.Address
		Index         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StakerAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Index = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 ) view returns(address stakerAddress, uint256 index)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) Stakers(arg0 *big.Int) (struct {
	StakerAddress common.Address
	Index         *big.Int
}, error) {
	return _StakingPoolImplementationForTest.Contract.Stakers(&_StakingPoolImplementationForTest.CallOpts, arg0)
}

// Stakers is a free data retrieval call binding the contract method 0xfd5e6dd1.
//
// Solidity: function stakers(uint256 ) view returns(address stakerAddress, uint256 index)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) Stakers(arg0 *big.Int) (struct {
	StakerAddress common.Address
	Index         *big.Int
}, error) {
	return _StakingPoolImplementationForTest.Contract.Stakers(&_StakingPoolImplementationForTest.CallOpts, arg0)
}

// StakingPoolImplementation is a free data retrieval call binding the contract method 0x2aa8195e.
//
// Solidity: function stakingPoolImplementation() view returns(address)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) StakingPoolImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "stakingPoolImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingPoolImplementation is a free data retrieval call binding the contract method 0x2aa8195e.
//
// Solidity: function stakingPoolImplementation() view returns(address)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) StakingPoolImplementation() (common.Address, error) {
	return _StakingPoolImplementationForTest.Contract.StakingPoolImplementation(&_StakingPoolImplementationForTest.CallOpts)
}

// StakingPoolImplementation is a free data retrieval call binding the contract method 0x2aa8195e.
//
// Solidity: function stakingPoolImplementation() view returns(address)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) StakingPoolImplementation() (common.Address, error) {
	return _StakingPoolImplementationForTest.Contract.StakingPoolImplementation(&_StakingPoolImplementationForTest.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "totalStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) TotalStaked() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.TotalStaked(&_StakingPoolImplementationForTest.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) TotalStaked() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.TotalStaked(&_StakingPoolImplementationForTest.CallOpts)
}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) TotemToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "totemToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) TotemToken() (common.Address, error) {
	return _StakingPoolImplementationForTest.Contract.TotemToken(&_StakingPoolImplementationForTest.CallOpts)
}

// TotemToken is a free data retrieval call binding the contract method 0xe8153c93.
//
// Solidity: function totemToken() view returns(address)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) TotemToken() (common.Address, error) {
	return _StakingPoolImplementationForTest.Contract.TotemToken(&_StakingPoolImplementationForTest.CallOpts)
}

// UpgradeEnabled is a free data retrieval call binding the contract method 0x8cf0e21e.
//
// Solidity: function upgradeEnabled() view returns(bool)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) UpgradeEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "upgradeEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UpgradeEnabled is a free data retrieval call binding the contract method 0x8cf0e21e.
//
// Solidity: function upgradeEnabled() view returns(bool)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) UpgradeEnabled() (bool, error) {
	return _StakingPoolImplementationForTest.Contract.UpgradeEnabled(&_StakingPoolImplementationForTest.CallOpts)
}

// UpgradeEnabled is a free data retrieval call binding the contract method 0x8cf0e21e.
//
// Solidity: function upgradeEnabled() view returns(bool)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) UpgradeEnabled() (bool, error) {
	return _StakingPoolImplementationForTest.Contract.UpgradeEnabled(&_StakingPoolImplementationForTest.CallOpts)
}

// UsdPrizeAmount is a free data retrieval call binding the contract method 0x73f31703.
//
// Solidity: function usdPrizeAmount() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) UsdPrizeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "usdPrizeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UsdPrizeAmount is a free data retrieval call binding the contract method 0x73f31703.
//
// Solidity: function usdPrizeAmount() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) UsdPrizeAmount() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.UsdPrizeAmount(&_StakingPoolImplementationForTest.CallOpts)
}

// UsdPrizeAmount is a free data retrieval call binding the contract method 0x73f31703.
//
// Solidity: function usdPrizeAmount() view returns(uint256)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) UsdPrizeAmount() (*big.Int, error) {
	return _StakingPoolImplementationForTest.Contract.UsdPrizeAmount(&_StakingPoolImplementationForTest.CallOpts)
}

// WrappedToken is a free data retrieval call binding the contract method 0x996c6cc3.
//
// Solidity: function wrappedToken() view returns(address)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) WrappedToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "wrappedToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WrappedToken is a free data retrieval call binding the contract method 0x996c6cc3.
//
// Solidity: function wrappedToken() view returns(address)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) WrappedToken() (common.Address, error) {
	return _StakingPoolImplementationForTest.Contract.WrappedToken(&_StakingPoolImplementationForTest.CallOpts)
}

// WrappedToken is a free data retrieval call binding the contract method 0x996c6cc3.
//
// Solidity: function wrappedToken() view returns(address)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) WrappedToken() (common.Address, error) {
	return _StakingPoolImplementationForTest.Contract.WrappedToken(&_StakingPoolImplementationForTest.CallOpts)
}

// WrappedTokenSymbol is a free data retrieval call binding the contract method 0x238a6c74.
//
// Solidity: function wrappedTokenSymbol() view returns(string)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCaller) WrappedTokenSymbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StakingPoolImplementationForTest.contract.Call(opts, &out, "wrappedTokenSymbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// WrappedTokenSymbol is a free data retrieval call binding the contract method 0x238a6c74.
//
// Solidity: function wrappedTokenSymbol() view returns(string)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) WrappedTokenSymbol() (string, error) {
	return _StakingPoolImplementationForTest.Contract.WrappedTokenSymbol(&_StakingPoolImplementationForTest.CallOpts)
}

// WrappedTokenSymbol is a free data retrieval call binding the contract method 0x238a6c74.
//
// Solidity: function wrappedTokenSymbol() view returns(string)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestCallerSession) WrappedTokenSymbol() (string, error) {
	return _StakingPoolImplementationForTest.Contract.WrappedTokenSymbol(&_StakingPoolImplementationForTest.CallOpts)
}

// WrappedTokenDistributorInitialize is a paid mutator transaction binding the contract method 0xa8f7b346.
//
// Solidity: function __WrappedTokenDistributor_initialize(address swapRouterAddress, address BUSDContractAddress, address WrappedTokenContractAddress) returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactor) WrappedTokenDistributorInitialize(opts *bind.TransactOpts, swapRouterAddress common.Address, BUSDContractAddress common.Address, WrappedTokenContractAddress common.Address) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.contract.Transact(opts, "__WrappedTokenDistributor_initialize", swapRouterAddress, BUSDContractAddress, WrappedTokenContractAddress)
}

// WrappedTokenDistributorInitialize is a paid mutator transaction binding the contract method 0xa8f7b346.
//
// Solidity: function __WrappedTokenDistributor_initialize(address swapRouterAddress, address BUSDContractAddress, address WrappedTokenContractAddress) returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) WrappedTokenDistributorInitialize(swapRouterAddress common.Address, BUSDContractAddress common.Address, WrappedTokenContractAddress common.Address) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.WrappedTokenDistributorInitialize(&_StakingPoolImplementationForTest.TransactOpts, swapRouterAddress, BUSDContractAddress, WrappedTokenContractAddress)
}

// WrappedTokenDistributorInitialize is a paid mutator transaction binding the contract method 0xa8f7b346.
//
// Solidity: function __WrappedTokenDistributor_initialize(address swapRouterAddress, address BUSDContractAddress, address WrappedTokenContractAddress) returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactorSession) WrappedTokenDistributorInitialize(swapRouterAddress common.Address, BUSDContractAddress common.Address, WrappedTokenContractAddress common.Address) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.WrappedTokenDistributorInitialize(&_StakingPoolImplementationForTest.TransactOpts, swapRouterAddress, BUSDContractAddress, WrappedTokenContractAddress)
}

// BatchStake is a paid mutator transaction binding the contract method 0xd49fab8a.
//
// Solidity: function batchStake(uint256[] _stakingAmounts, uint256[] _predictions) returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactor) BatchStake(opts *bind.TransactOpts, _stakingAmounts []*big.Int, _predictions []*big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.contract.Transact(opts, "batchStake", _stakingAmounts, _predictions)
}

// BatchStake is a paid mutator transaction binding the contract method 0xd49fab8a.
//
// Solidity: function batchStake(uint256[] _stakingAmounts, uint256[] _predictions) returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) BatchStake(_stakingAmounts []*big.Int, _predictions []*big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.BatchStake(&_StakingPoolImplementationForTest.TransactOpts, _stakingAmounts, _predictions)
}

// BatchStake is a paid mutator transaction binding the contract method 0xd49fab8a.
//
// Solidity: function batchStake(uint256[] _stakingAmounts, uint256[] _predictions) returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactorSession) BatchStake(_stakingAmounts []*big.Int, _predictions []*big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.BatchStake(&_StakingPoolImplementationForTest.TransactOpts, _stakingAmounts, _predictions)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactor) ClaimReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.contract.Transact(opts, "claimReward")
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) ClaimReward() (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.ClaimReward(&_StakingPoolImplementationForTest.TransactOpts)
}

// ClaimReward is a paid mutator transaction binding the contract method 0xb88a802f.
//
// Solidity: function claimReward() returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactorSession) ClaimReward() (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.ClaimReward(&_StakingPoolImplementationForTest.TransactOpts)
}

// DeclareEmergency is a paid mutator transaction binding the contract method 0x31e244e5.
//
// Solidity: function declareEmergency() returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactor) DeclareEmergency(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.contract.Transact(opts, "declareEmergency")
}

// DeclareEmergency is a paid mutator transaction binding the contract method 0x31e244e5.
//
// Solidity: function declareEmergency() returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) DeclareEmergency() (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.DeclareEmergency(&_StakingPoolImplementationForTest.TransactOpts)
}

// DeclareEmergency is a paid mutator transaction binding the contract method 0x31e244e5.
//
// Solidity: function declareEmergency() returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactorSession) DeclareEmergency() (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.DeclareEmergency(&_StakingPoolImplementationForTest.TransactOpts)
}

// DeletePool is a paid mutator transaction binding the contract method 0x8bfbbcbe.
//
// Solidity: function deletePool() returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactor) DeletePool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.contract.Transact(opts, "deletePool")
}

// DeletePool is a paid mutator transaction binding the contract method 0x8bfbbcbe.
//
// Solidity: function deletePool() returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) DeletePool() (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.DeletePool(&_StakingPoolImplementationForTest.TransactOpts)
}

// DeletePool is a paid mutator transaction binding the contract method 0x8bfbbcbe.
//
// Solidity: function deletePool() returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactorSession) DeletePool() (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.DeletePool(&_StakingPoolImplementationForTest.TransactOpts)
}

// EmergentWithdraw is a paid mutator transaction binding the contract method 0x8a9920b7.
//
// Solidity: function emergentWithdraw() returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactor) EmergentWithdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.contract.Transact(opts, "emergentWithdraw")
}

// EmergentWithdraw is a paid mutator transaction binding the contract method 0x8a9920b7.
//
// Solidity: function emergentWithdraw() returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) EmergentWithdraw() (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.EmergentWithdraw(&_StakingPoolImplementationForTest.TransactOpts)
}

// EmergentWithdraw is a paid mutator transaction binding the contract method 0x8a9920b7.
//
// Solidity: function emergentWithdraw() returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactorSession) EmergentWithdraw() (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.EmergentWithdraw(&_StakingPoolImplementationForTest.TransactOpts)
}

// EndPool is a paid mutator transaction binding the contract method 0x2d42cf7d.
//
// Solidity: function endPool() returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactor) EndPool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.contract.Transact(opts, "endPool")
}

// EndPool is a paid mutator transaction binding the contract method 0x2d42cf7d.
//
// Solidity: function endPool() returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) EndPool() (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.EndPool(&_StakingPoolImplementationForTest.TransactOpts)
}

// EndPool is a paid mutator transaction binding the contract method 0x2d42cf7d.
//
// Solidity: function endPool() returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactorSession) EndPool() (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.EndPool(&_StakingPoolImplementationForTest.TransactOpts)
}

// IndexedClaimReward is a paid mutator transaction binding the contract method 0xb1b1e1c0.
//
// Solidity: function indexedClaimReward(uint256 stakeIndex) returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactor) IndexedClaimReward(opts *bind.TransactOpts, stakeIndex *big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.contract.Transact(opts, "indexedClaimReward", stakeIndex)
}

// IndexedClaimReward is a paid mutator transaction binding the contract method 0xb1b1e1c0.
//
// Solidity: function indexedClaimReward(uint256 stakeIndex) returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) IndexedClaimReward(stakeIndex *big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.IndexedClaimReward(&_StakingPoolImplementationForTest.TransactOpts, stakeIndex)
}

// IndexedClaimReward is a paid mutator transaction binding the contract method 0xb1b1e1c0.
//
// Solidity: function indexedClaimReward(uint256 stakeIndex) returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactorSession) IndexedClaimReward(stakeIndex *big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.IndexedClaimReward(&_StakingPoolImplementationForTest.TransactOpts, stakeIndex)
}

// LockPool is a paid mutator transaction binding the contract method 0x4026478e.
//
// Solidity: function lockPool() returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactor) LockPool(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.contract.Transact(opts, "lockPool")
}

// LockPool is a paid mutator transaction binding the contract method 0x4026478e.
//
// Solidity: function lockPool() returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) LockPool() (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.LockPool(&_StakingPoolImplementationForTest.TransactOpts)
}

// LockPool is a paid mutator transaction binding the contract method 0x4026478e.
//
// Solidity: function lockPool() returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactorSession) LockPool() (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.LockPool(&_StakingPoolImplementationForTest.TransactOpts)
}

// PurchaseWrappedToken is a paid mutator transaction binding the contract method 0x4ddc1bdc.
//
// Solidity: function purchaseWrappedToken(uint256 usdAmount, uint256 deadline) returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactor) PurchaseWrappedToken(opts *bind.TransactOpts, usdAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.contract.Transact(opts, "purchaseWrappedToken", usdAmount, deadline)
}

// PurchaseWrappedToken is a paid mutator transaction binding the contract method 0x4ddc1bdc.
//
// Solidity: function purchaseWrappedToken(uint256 usdAmount, uint256 deadline) returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) PurchaseWrappedToken(usdAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.PurchaseWrappedToken(&_StakingPoolImplementationForTest.TransactOpts, usdAmount, deadline)
}

// PurchaseWrappedToken is a paid mutator transaction binding the contract method 0x4ddc1bdc.
//
// Solidity: function purchaseWrappedToken(uint256 usdAmount, uint256 deadline) returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactorSession) PurchaseWrappedToken(usdAmount *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.PurchaseWrappedToken(&_StakingPoolImplementationForTest.TransactOpts, usdAmount, deadline)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.RenounceOwnership(&_StakingPoolImplementationForTest.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.RenounceOwnership(&_StakingPoolImplementationForTest.TransactOpts)
}

// SetActivationStatus is a paid mutator transaction binding the contract method 0xa118c9a7.
//
// Solidity: function setActivationStatus(bool _activationStatus) returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactor) SetActivationStatus(opts *bind.TransactOpts, _activationStatus bool) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.contract.Transact(opts, "setActivationStatus", _activationStatus)
}

// SetActivationStatus is a paid mutator transaction binding the contract method 0xa118c9a7.
//
// Solidity: function setActivationStatus(bool _activationStatus) returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) SetActivationStatus(_activationStatus bool) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.SetActivationStatus(&_StakingPoolImplementationForTest.TransactOpts, _activationStatus)
}

// SetActivationStatus is a paid mutator transaction binding the contract method 0xa118c9a7.
//
// Solidity: function setActivationStatus(bool _activationStatus) returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactorSession) SetActivationStatus(_activationStatus bool) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.SetActivationStatus(&_StakingPoolImplementationForTest.TransactOpts, _activationStatus)
}

// SetOracleToZero is a paid mutator transaction binding the contract method 0xdf7f92e0.
//
// Solidity: function setOracleToZero() returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactor) SetOracleToZero(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.contract.Transact(opts, "setOracleToZero")
}

// SetOracleToZero is a paid mutator transaction binding the contract method 0xdf7f92e0.
//
// Solidity: function setOracleToZero() returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) SetOracleToZero() (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.SetOracleToZero(&_StakingPoolImplementationForTest.TransactOpts)
}

// SetOracleToZero is a paid mutator transaction binding the contract method 0xdf7f92e0.
//
// Solidity: function setOracleToZero() returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactorSession) SetOracleToZero() (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.SetOracleToZero(&_StakingPoolImplementationForTest.TransactOpts)
}

// SetSortedStakers is a paid mutator transaction binding the contract method 0x444a4902.
//
// Solidity: function setSortedStakers(address[25] addrArray, uint256[25] indexArray) returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactor) SetSortedStakers(opts *bind.TransactOpts, addrArray [25]common.Address, indexArray [25]*big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.contract.Transact(opts, "setSortedStakers", addrArray, indexArray)
}

// SetSortedStakers is a paid mutator transaction binding the contract method 0x444a4902.
//
// Solidity: function setSortedStakers(address[25] addrArray, uint256[25] indexArray) returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) SetSortedStakers(addrArray [25]common.Address, indexArray [25]*big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.SetSortedStakers(&_StakingPoolImplementationForTest.TransactOpts, addrArray, indexArray)
}

// SetSortedStakers is a paid mutator transaction binding the contract method 0x444a4902.
//
// Solidity: function setSortedStakers(address[25] addrArray, uint256[25] indexArray) returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactorSession) SetSortedStakers(addrArray [25]common.Address, indexArray [25]*big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.SetSortedStakers(&_StakingPoolImplementationForTest.TransactOpts, addrArray, indexArray)
}

// Stake is a paid mutator transaction binding the contract method 0x7b0472f0.
//
// Solidity: function stake(uint256 _amount, uint256 _pricePrediction) returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactor) Stake(opts *bind.TransactOpts, _amount *big.Int, _pricePrediction *big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.contract.Transact(opts, "stake", _amount, _pricePrediction)
}

// Stake is a paid mutator transaction binding the contract method 0x7b0472f0.
//
// Solidity: function stake(uint256 _amount, uint256 _pricePrediction) returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) Stake(_amount *big.Int, _pricePrediction *big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.Stake(&_StakingPoolImplementationForTest.TransactOpts, _amount, _pricePrediction)
}

// Stake is a paid mutator transaction binding the contract method 0x7b0472f0.
//
// Solidity: function stake(uint256 _amount, uint256 _pricePrediction) returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactorSession) Stake(_amount *big.Int, _pricePrediction *big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.Stake(&_StakingPoolImplementationForTest.TransactOpts, _amount, _pricePrediction)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.TransferOwnership(&_StakingPoolImplementationForTest.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.TransferOwnership(&_StakingPoolImplementationForTest.TransactOpts, newOwner)
}

// UpdateMaturingPrice is a paid mutator transaction binding the contract method 0xe113af51.
//
// Solidity: function updateMaturingPrice(uint256 _predictionPrice, bool _samePredictionPrizeToken, uint256 _prizePrice, address _oracleContract, uint256 _oracleDecimals) returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactor) UpdateMaturingPrice(opts *bind.TransactOpts, _predictionPrice *big.Int, _samePredictionPrizeToken bool, _prizePrice *big.Int, _oracleContract common.Address, _oracleDecimals *big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.contract.Transact(opts, "updateMaturingPrice", _predictionPrice, _samePredictionPrizeToken, _prizePrice, _oracleContract, _oracleDecimals)
}

// UpdateMaturingPrice is a paid mutator transaction binding the contract method 0xe113af51.
//
// Solidity: function updateMaturingPrice(uint256 _predictionPrice, bool _samePredictionPrizeToken, uint256 _prizePrice, address _oracleContract, uint256 _oracleDecimals) returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) UpdateMaturingPrice(_predictionPrice *big.Int, _samePredictionPrizeToken bool, _prizePrice *big.Int, _oracleContract common.Address, _oracleDecimals *big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.UpdateMaturingPrice(&_StakingPoolImplementationForTest.TransactOpts, _predictionPrice, _samePredictionPrizeToken, _prizePrice, _oracleContract, _oracleDecimals)
}

// UpdateMaturingPrice is a paid mutator transaction binding the contract method 0xe113af51.
//
// Solidity: function updateMaturingPrice(uint256 _predictionPrice, bool _samePredictionPrizeToken, uint256 _prizePrice, address _oracleContract, uint256 _oracleDecimals) returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactorSession) UpdateMaturingPrice(_predictionPrice *big.Int, _samePredictionPrizeToken bool, _prizePrice *big.Int, _oracleContract common.Address, _oracleDecimals *big.Int) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.UpdateMaturingPrice(&_StakingPoolImplementationForTest.TransactOpts, _predictionPrice, _samePredictionPrizeToken, _prizePrice, _oracleContract, _oracleDecimals)
}

// WithdrawStuckTokens is a paid mutator transaction binding the contract method 0x5d2631e2.
//
// Solidity: function withdrawStuckTokens(address _stuckToken, uint256 amount, address receiver) returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactor) WithdrawStuckTokens(opts *bind.TransactOpts, _stuckToken common.Address, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.contract.Transact(opts, "withdrawStuckTokens", _stuckToken, amount, receiver)
}

// WithdrawStuckTokens is a paid mutator transaction binding the contract method 0x5d2631e2.
//
// Solidity: function withdrawStuckTokens(address _stuckToken, uint256 amount, address receiver) returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestSession) WithdrawStuckTokens(_stuckToken common.Address, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.WithdrawStuckTokens(&_StakingPoolImplementationForTest.TransactOpts, _stuckToken, amount, receiver)
}

// WithdrawStuckTokens is a paid mutator transaction binding the contract method 0x5d2631e2.
//
// Solidity: function withdrawStuckTokens(address _stuckToken, uint256 amount, address receiver) returns()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestTransactorSession) WithdrawStuckTokens(_stuckToken common.Address, amount *big.Int, receiver common.Address) (*types.Transaction, error) {
	return _StakingPoolImplementationForTest.Contract.WithdrawStuckTokens(&_StakingPoolImplementationForTest.TransactOpts, _stuckToken, amount, receiver)
}

// StakingPoolImplementationForTestApproveTokensIterator is returned from FilterApproveTokens and is used to iterate over the raw logs and unpacked data for ApproveTokens events raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestApproveTokensIterator struct {
	Event *StakingPoolImplementationForTestApproveTokens // Event containing the contract specifics and raw log

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
func (it *StakingPoolImplementationForTestApproveTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolImplementationForTestApproveTokens)
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
		it.Event = new(StakingPoolImplementationForTestApproveTokens)
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
func (it *StakingPoolImplementationForTestApproveTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolImplementationForTestApproveTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolImplementationForTestApproveTokens represents a ApproveTokens event raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestApproveTokens struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterApproveTokens is a free log retrieval operation binding the contract event 0xd77df63d9076d067d9fd7af16c3d67db48b84064c3314f5900ae919922d364a1.
//
// Solidity: event ApproveTokens(address indexed to, uint256 amount)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) FilterApproveTokens(opts *bind.FilterOpts, to []common.Address) (*StakingPoolImplementationForTestApproveTokensIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakingPoolImplementationForTest.contract.FilterLogs(opts, "ApproveTokens", toRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationForTestApproveTokensIterator{contract: _StakingPoolImplementationForTest.contract, event: "ApproveTokens", logs: logs, sub: sub}, nil
}

// WatchApproveTokens is a free log subscription operation binding the contract event 0xd77df63d9076d067d9fd7af16c3d67db48b84064c3314f5900ae919922d364a1.
//
// Solidity: event ApproveTokens(address indexed to, uint256 amount)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) WatchApproveTokens(opts *bind.WatchOpts, sink chan<- *StakingPoolImplementationForTestApproveTokens, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakingPoolImplementationForTest.contract.WatchLogs(opts, "ApproveTokens", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolImplementationForTestApproveTokens)
				if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "ApproveTokens", log); err != nil {
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
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) ParseApproveTokens(log types.Log) (*StakingPoolImplementationForTestApproveTokens, error) {
	event := new(StakingPoolImplementationForTestApproveTokens)
	if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "ApproveTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolImplementationForTestBatchStakeIterator is returned from FilterBatchStake and is used to iterate over the raw logs and unpacked data for BatchStake events raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestBatchStakeIterator struct {
	Event *StakingPoolImplementationForTestBatchStake // Event containing the contract specifics and raw log

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
func (it *StakingPoolImplementationForTestBatchStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolImplementationForTestBatchStake)
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
		it.Event = new(StakingPoolImplementationForTestBatchStake)
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
func (it *StakingPoolImplementationForTestBatchStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolImplementationForTestBatchStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolImplementationForTestBatchStake represents a BatchStake event raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestBatchStake struct {
	User              common.Address
	StakeAmounts      []*big.Int
	Predictions       []*big.Int
	IndexOfFirstStake *big.Int
	IndexOfLastStake  *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterBatchStake is a free log retrieval operation binding the contract event 0x852fe7a2ad7b2c8c0a0e9b927734feb1acc3e8bc608b8921d4c4768c6193c33b.
//
// Solidity: event BatchStake(address indexed user, uint256[] stakeAmounts, uint256[] predictions, uint256 indexOfFirstStake, uint256 indexOfLastStake)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) FilterBatchStake(opts *bind.FilterOpts, user []common.Address) (*StakingPoolImplementationForTestBatchStakeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingPoolImplementationForTest.contract.FilterLogs(opts, "BatchStake", userRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationForTestBatchStakeIterator{contract: _StakingPoolImplementationForTest.contract, event: "BatchStake", logs: logs, sub: sub}, nil
}

// WatchBatchStake is a free log subscription operation binding the contract event 0x852fe7a2ad7b2c8c0a0e9b927734feb1acc3e8bc608b8921d4c4768c6193c33b.
//
// Solidity: event BatchStake(address indexed user, uint256[] stakeAmounts, uint256[] predictions, uint256 indexOfFirstStake, uint256 indexOfLastStake)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) WatchBatchStake(opts *bind.WatchOpts, sink chan<- *StakingPoolImplementationForTestBatchStake, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingPoolImplementationForTest.contract.WatchLogs(opts, "BatchStake", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolImplementationForTestBatchStake)
				if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "BatchStake", log); err != nil {
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

// ParseBatchStake is a log parse operation binding the contract event 0x852fe7a2ad7b2c8c0a0e9b927734feb1acc3e8bc608b8921d4c4768c6193c33b.
//
// Solidity: event BatchStake(address indexed user, uint256[] stakeAmounts, uint256[] predictions, uint256 indexOfFirstStake, uint256 indexOfLastStake)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) ParseBatchStake(log types.Log) (*StakingPoolImplementationForTestBatchStake, error) {
	event := new(StakingPoolImplementationForTestBatchStake)
	if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "BatchStake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolImplementationForTestDistributedBTCIterator is returned from FilterDistributedBTC and is used to iterate over the raw logs and unpacked data for DistributedBTC events raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestDistributedBTCIterator struct {
	Event *StakingPoolImplementationForTestDistributedBTC // Event containing the contract specifics and raw log

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
func (it *StakingPoolImplementationForTestDistributedBTCIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolImplementationForTestDistributedBTC)
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
		it.Event = new(StakingPoolImplementationForTestDistributedBTC)
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
func (it *StakingPoolImplementationForTestDistributedBTCIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolImplementationForTestDistributedBTCIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolImplementationForTestDistributedBTC represents a DistributedBTC event raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestDistributedBTC struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDistributedBTC is a free log retrieval operation binding the contract event 0xe690d17e165a33f4a6cf4c2374c904d5f425eb2563320d08ed595794a1ba33aa.
//
// Solidity: event DistributedBTC(address indexed to, uint256 amount)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) FilterDistributedBTC(opts *bind.FilterOpts, to []common.Address) (*StakingPoolImplementationForTestDistributedBTCIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakingPoolImplementationForTest.contract.FilterLogs(opts, "DistributedBTC", toRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationForTestDistributedBTCIterator{contract: _StakingPoolImplementationForTest.contract, event: "DistributedBTC", logs: logs, sub: sub}, nil
}

// WatchDistributedBTC is a free log subscription operation binding the contract event 0xe690d17e165a33f4a6cf4c2374c904d5f425eb2563320d08ed595794a1ba33aa.
//
// Solidity: event DistributedBTC(address indexed to, uint256 amount)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) WatchDistributedBTC(opts *bind.WatchOpts, sink chan<- *StakingPoolImplementationForTestDistributedBTC, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakingPoolImplementationForTest.contract.WatchLogs(opts, "DistributedBTC", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolImplementationForTestDistributedBTC)
				if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "DistributedBTC", log); err != nil {
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
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) ParseDistributedBTC(log types.Log) (*StakingPoolImplementationForTestDistributedBTC, error) {
	event := new(StakingPoolImplementationForTestDistributedBTC)
	if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "DistributedBTC", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolImplementationForTestOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestOwnershipTransferredIterator struct {
	Event *StakingPoolImplementationForTestOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakingPoolImplementationForTestOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolImplementationForTestOwnershipTransferred)
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
		it.Event = new(StakingPoolImplementationForTestOwnershipTransferred)
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
func (it *StakingPoolImplementationForTestOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolImplementationForTestOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolImplementationForTestOwnershipTransferred represents a OwnershipTransferred event raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingPoolImplementationForTestOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingPoolImplementationForTest.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationForTestOwnershipTransferredIterator{contract: _StakingPoolImplementationForTest.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingPoolImplementationForTestOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _StakingPoolImplementationForTest.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolImplementationForTestOwnershipTransferred)
				if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) ParseOwnershipTransferred(log types.Log) (*StakingPoolImplementationForTestOwnershipTransferred, error) {
	event := new(StakingPoolImplementationForTestOwnershipTransferred)
	if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolImplementationForTestPoolActivatedIterator is returned from FilterPoolActivated and is used to iterate over the raw logs and unpacked data for PoolActivated events raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestPoolActivatedIterator struct {
	Event *StakingPoolImplementationForTestPoolActivated // Event containing the contract specifics and raw log

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
func (it *StakingPoolImplementationForTestPoolActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolImplementationForTestPoolActivated)
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
		it.Event = new(StakingPoolImplementationForTestPoolActivated)
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
func (it *StakingPoolImplementationForTestPoolActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolImplementationForTestPoolActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolImplementationForTestPoolActivated represents a PoolActivated event raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestPoolActivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPoolActivated is a free log retrieval operation binding the contract event 0x3d31b97eef590df350bb2cdbc390034c613c9f23ea6ae1906682628c85248b53.
//
// Solidity: event PoolActivated()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) FilterPoolActivated(opts *bind.FilterOpts) (*StakingPoolImplementationForTestPoolActivatedIterator, error) {

	logs, sub, err := _StakingPoolImplementationForTest.contract.FilterLogs(opts, "PoolActivated")
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationForTestPoolActivatedIterator{contract: _StakingPoolImplementationForTest.contract, event: "PoolActivated", logs: logs, sub: sub}, nil
}

// WatchPoolActivated is a free log subscription operation binding the contract event 0x3d31b97eef590df350bb2cdbc390034c613c9f23ea6ae1906682628c85248b53.
//
// Solidity: event PoolActivated()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) WatchPoolActivated(opts *bind.WatchOpts, sink chan<- *StakingPoolImplementationForTestPoolActivated) (event.Subscription, error) {

	logs, sub, err := _StakingPoolImplementationForTest.contract.WatchLogs(opts, "PoolActivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolImplementationForTestPoolActivated)
				if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "PoolActivated", log); err != nil {
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

// ParsePoolActivated is a log parse operation binding the contract event 0x3d31b97eef590df350bb2cdbc390034c613c9f23ea6ae1906682628c85248b53.
//
// Solidity: event PoolActivated()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) ParsePoolActivated(log types.Log) (*StakingPoolImplementationForTestPoolActivated, error) {
	event := new(StakingPoolImplementationForTestPoolActivated)
	if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "PoolActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolImplementationForTestPoolDeactivatedIterator is returned from FilterPoolDeactivated and is used to iterate over the raw logs and unpacked data for PoolDeactivated events raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestPoolDeactivatedIterator struct {
	Event *StakingPoolImplementationForTestPoolDeactivated // Event containing the contract specifics and raw log

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
func (it *StakingPoolImplementationForTestPoolDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolImplementationForTestPoolDeactivated)
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
		it.Event = new(StakingPoolImplementationForTestPoolDeactivated)
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
func (it *StakingPoolImplementationForTestPoolDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolImplementationForTestPoolDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolImplementationForTestPoolDeactivated represents a PoolDeactivated event raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestPoolDeactivated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPoolDeactivated is a free log retrieval operation binding the contract event 0x763d36bfedf6707954f8bb39c2011a58a47a54844f2b3cdcd326c71cad0aa753.
//
// Solidity: event PoolDeactivated()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) FilterPoolDeactivated(opts *bind.FilterOpts) (*StakingPoolImplementationForTestPoolDeactivatedIterator, error) {

	logs, sub, err := _StakingPoolImplementationForTest.contract.FilterLogs(opts, "PoolDeactivated")
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationForTestPoolDeactivatedIterator{contract: _StakingPoolImplementationForTest.contract, event: "PoolDeactivated", logs: logs, sub: sub}, nil
}

// WatchPoolDeactivated is a free log subscription operation binding the contract event 0x763d36bfedf6707954f8bb39c2011a58a47a54844f2b3cdcd326c71cad0aa753.
//
// Solidity: event PoolDeactivated()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) WatchPoolDeactivated(opts *bind.WatchOpts, sink chan<- *StakingPoolImplementationForTestPoolDeactivated) (event.Subscription, error) {

	logs, sub, err := _StakingPoolImplementationForTest.contract.WatchLogs(opts, "PoolDeactivated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolImplementationForTestPoolDeactivated)
				if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "PoolDeactivated", log); err != nil {
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

// ParsePoolDeactivated is a log parse operation binding the contract event 0x763d36bfedf6707954f8bb39c2011a58a47a54844f2b3cdcd326c71cad0aa753.
//
// Solidity: event PoolDeactivated()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) ParsePoolDeactivated(log types.Log) (*StakingPoolImplementationForTestPoolDeactivated, error) {
	event := new(StakingPoolImplementationForTestPoolDeactivated)
	if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "PoolDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolImplementationForTestPoolDeletedIterator is returned from FilterPoolDeleted and is used to iterate over the raw logs and unpacked data for PoolDeleted events raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestPoolDeletedIterator struct {
	Event *StakingPoolImplementationForTestPoolDeleted // Event containing the contract specifics and raw log

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
func (it *StakingPoolImplementationForTestPoolDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolImplementationForTestPoolDeleted)
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
		it.Event = new(StakingPoolImplementationForTestPoolDeleted)
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
func (it *StakingPoolImplementationForTestPoolDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolImplementationForTestPoolDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolImplementationForTestPoolDeleted represents a PoolDeleted event raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestPoolDeleted struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPoolDeleted is a free log retrieval operation binding the contract event 0x9331b2a38c658c57a39ea19b2709d928c28a12cc8fbb838d12b85148550f937d.
//
// Solidity: event PoolDeleted()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) FilterPoolDeleted(opts *bind.FilterOpts) (*StakingPoolImplementationForTestPoolDeletedIterator, error) {

	logs, sub, err := _StakingPoolImplementationForTest.contract.FilterLogs(opts, "PoolDeleted")
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationForTestPoolDeletedIterator{contract: _StakingPoolImplementationForTest.contract, event: "PoolDeleted", logs: logs, sub: sub}, nil
}

// WatchPoolDeleted is a free log subscription operation binding the contract event 0x9331b2a38c658c57a39ea19b2709d928c28a12cc8fbb838d12b85148550f937d.
//
// Solidity: event PoolDeleted()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) WatchPoolDeleted(opts *bind.WatchOpts, sink chan<- *StakingPoolImplementationForTestPoolDeleted) (event.Subscription, error) {

	logs, sub, err := _StakingPoolImplementationForTest.contract.WatchLogs(opts, "PoolDeleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolImplementationForTestPoolDeleted)
				if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "PoolDeleted", log); err != nil {
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

// ParsePoolDeleted is a log parse operation binding the contract event 0x9331b2a38c658c57a39ea19b2709d928c28a12cc8fbb838d12b85148550f937d.
//
// Solidity: event PoolDeleted()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) ParsePoolDeleted(log types.Log) (*StakingPoolImplementationForTestPoolDeleted, error) {
	event := new(StakingPoolImplementationForTestPoolDeleted)
	if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "PoolDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolImplementationForTestPoolLockedIterator is returned from FilterPoolLocked and is used to iterate over the raw logs and unpacked data for PoolLocked events raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestPoolLockedIterator struct {
	Event *StakingPoolImplementationForTestPoolLocked // Event containing the contract specifics and raw log

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
func (it *StakingPoolImplementationForTestPoolLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolImplementationForTestPoolLocked)
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
		it.Event = new(StakingPoolImplementationForTestPoolLocked)
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
func (it *StakingPoolImplementationForTestPoolLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolImplementationForTestPoolLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolImplementationForTestPoolLocked represents a PoolLocked event raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestPoolLocked struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPoolLocked is a free log retrieval operation binding the contract event 0x2e136745550967f28b779ce8c395ca060a2f844edfb5b06631fe677f66f9e305.
//
// Solidity: event PoolLocked()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) FilterPoolLocked(opts *bind.FilterOpts) (*StakingPoolImplementationForTestPoolLockedIterator, error) {

	logs, sub, err := _StakingPoolImplementationForTest.contract.FilterLogs(opts, "PoolLocked")
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationForTestPoolLockedIterator{contract: _StakingPoolImplementationForTest.contract, event: "PoolLocked", logs: logs, sub: sub}, nil
}

// WatchPoolLocked is a free log subscription operation binding the contract event 0x2e136745550967f28b779ce8c395ca060a2f844edfb5b06631fe677f66f9e305.
//
// Solidity: event PoolLocked()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) WatchPoolLocked(opts *bind.WatchOpts, sink chan<- *StakingPoolImplementationForTestPoolLocked) (event.Subscription, error) {

	logs, sub, err := _StakingPoolImplementationForTest.contract.WatchLogs(opts, "PoolLocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolImplementationForTestPoolLocked)
				if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "PoolLocked", log); err != nil {
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

// ParsePoolLocked is a log parse operation binding the contract event 0x2e136745550967f28b779ce8c395ca060a2f844edfb5b06631fe677f66f9e305.
//
// Solidity: event PoolLocked()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) ParsePoolLocked(log types.Log) (*StakingPoolImplementationForTestPoolLocked, error) {
	event := new(StakingPoolImplementationForTestPoolLocked)
	if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "PoolLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolImplementationForTestPoolMaturedIterator is returned from FilterPoolMatured and is used to iterate over the raw logs and unpacked data for PoolMatured events raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestPoolMaturedIterator struct {
	Event *StakingPoolImplementationForTestPoolMatured // Event containing the contract specifics and raw log

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
func (it *StakingPoolImplementationForTestPoolMaturedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolImplementationForTestPoolMatured)
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
		it.Event = new(StakingPoolImplementationForTestPoolMatured)
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
func (it *StakingPoolImplementationForTestPoolMaturedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolImplementationForTestPoolMaturedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolImplementationForTestPoolMatured represents a PoolMatured event raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestPoolMatured struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPoolMatured is a free log retrieval operation binding the contract event 0xe8ef7a17c92d025d9c1ead1c6b3ca3b44d3283d0ee954f615c5c476911f629a5.
//
// Solidity: event PoolMatured()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) FilterPoolMatured(opts *bind.FilterOpts) (*StakingPoolImplementationForTestPoolMaturedIterator, error) {

	logs, sub, err := _StakingPoolImplementationForTest.contract.FilterLogs(opts, "PoolMatured")
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationForTestPoolMaturedIterator{contract: _StakingPoolImplementationForTest.contract, event: "PoolMatured", logs: logs, sub: sub}, nil
}

// WatchPoolMatured is a free log subscription operation binding the contract event 0xe8ef7a17c92d025d9c1ead1c6b3ca3b44d3283d0ee954f615c5c476911f629a5.
//
// Solidity: event PoolMatured()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) WatchPoolMatured(opts *bind.WatchOpts, sink chan<- *StakingPoolImplementationForTestPoolMatured) (event.Subscription, error) {

	logs, sub, err := _StakingPoolImplementationForTest.contract.WatchLogs(opts, "PoolMatured")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolImplementationForTestPoolMatured)
				if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "PoolMatured", log); err != nil {
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

// ParsePoolMatured is a log parse operation binding the contract event 0xe8ef7a17c92d025d9c1ead1c6b3ca3b44d3283d0ee954f615c5c476911f629a5.
//
// Solidity: event PoolMatured()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) ParsePoolMatured(log types.Log) (*StakingPoolImplementationForTestPoolMatured, error) {
	event := new(StakingPoolImplementationForTestPoolMatured)
	if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "PoolMatured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolImplementationForTestPoolSortedIterator is returned from FilterPoolSorted and is used to iterate over the raw logs and unpacked data for PoolSorted events raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestPoolSortedIterator struct {
	Event *StakingPoolImplementationForTestPoolSorted // Event containing the contract specifics and raw log

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
func (it *StakingPoolImplementationForTestPoolSortedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolImplementationForTestPoolSorted)
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
		it.Event = new(StakingPoolImplementationForTestPoolSorted)
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
func (it *StakingPoolImplementationForTestPoolSortedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolImplementationForTestPoolSortedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolImplementationForTestPoolSorted represents a PoolSorted event raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestPoolSorted struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPoolSorted is a free log retrieval operation binding the contract event 0x5c243cc6e4f210bf2dd86b8eca0ca891d3f8ad2eb09fbf91735eb348bc199c7e.
//
// Solidity: event PoolSorted()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) FilterPoolSorted(opts *bind.FilterOpts) (*StakingPoolImplementationForTestPoolSortedIterator, error) {

	logs, sub, err := _StakingPoolImplementationForTest.contract.FilterLogs(opts, "PoolSorted")
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationForTestPoolSortedIterator{contract: _StakingPoolImplementationForTest.contract, event: "PoolSorted", logs: logs, sub: sub}, nil
}

// WatchPoolSorted is a free log subscription operation binding the contract event 0x5c243cc6e4f210bf2dd86b8eca0ca891d3f8ad2eb09fbf91735eb348bc199c7e.
//
// Solidity: event PoolSorted()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) WatchPoolSorted(opts *bind.WatchOpts, sink chan<- *StakingPoolImplementationForTestPoolSorted) (event.Subscription, error) {

	logs, sub, err := _StakingPoolImplementationForTest.contract.WatchLogs(opts, "PoolSorted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolImplementationForTestPoolSorted)
				if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "PoolSorted", log); err != nil {
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

// ParsePoolSorted is a log parse operation binding the contract event 0x5c243cc6e4f210bf2dd86b8eca0ca891d3f8ad2eb09fbf91735eb348bc199c7e.
//
// Solidity: event PoolSorted()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) ParsePoolSorted(log types.Log) (*StakingPoolImplementationForTestPoolSorted, error) {
	event := new(StakingPoolImplementationForTestPoolSorted)
	if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "PoolSorted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolImplementationForTestPoolUnLockedIterator is returned from FilterPoolUnLocked and is used to iterate over the raw logs and unpacked data for PoolUnLocked events raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestPoolUnLockedIterator struct {
	Event *StakingPoolImplementationForTestPoolUnLocked // Event containing the contract specifics and raw log

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
func (it *StakingPoolImplementationForTestPoolUnLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolImplementationForTestPoolUnLocked)
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
		it.Event = new(StakingPoolImplementationForTestPoolUnLocked)
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
func (it *StakingPoolImplementationForTestPoolUnLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolImplementationForTestPoolUnLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolImplementationForTestPoolUnLocked represents a PoolUnLocked event raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestPoolUnLocked struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPoolUnLocked is a free log retrieval operation binding the contract event 0x5519a5cb9044a928ca9f5e80693e0c2ceaadbbcc2ce720498c600bb2562a5503.
//
// Solidity: event PoolUnLocked()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) FilterPoolUnLocked(opts *bind.FilterOpts) (*StakingPoolImplementationForTestPoolUnLockedIterator, error) {

	logs, sub, err := _StakingPoolImplementationForTest.contract.FilterLogs(opts, "PoolUnLocked")
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationForTestPoolUnLockedIterator{contract: _StakingPoolImplementationForTest.contract, event: "PoolUnLocked", logs: logs, sub: sub}, nil
}

// WatchPoolUnLocked is a free log subscription operation binding the contract event 0x5519a5cb9044a928ca9f5e80693e0c2ceaadbbcc2ce720498c600bb2562a5503.
//
// Solidity: event PoolUnLocked()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) WatchPoolUnLocked(opts *bind.WatchOpts, sink chan<- *StakingPoolImplementationForTestPoolUnLocked) (event.Subscription, error) {

	logs, sub, err := _StakingPoolImplementationForTest.contract.WatchLogs(opts, "PoolUnLocked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolImplementationForTestPoolUnLocked)
				if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "PoolUnLocked", log); err != nil {
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

// ParsePoolUnLocked is a log parse operation binding the contract event 0x5519a5cb9044a928ca9f5e80693e0c2ceaadbbcc2ce720498c600bb2562a5503.
//
// Solidity: event PoolUnLocked()
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) ParsePoolUnLocked(log types.Log) (*StakingPoolImplementationForTestPoolUnLocked, error) {
	event := new(StakingPoolImplementationForTestPoolUnLocked)
	if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "PoolUnLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolImplementationForTestReceivedTokensIterator is returned from FilterReceivedTokens and is used to iterate over the raw logs and unpacked data for ReceivedTokens events raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestReceivedTokensIterator struct {
	Event *StakingPoolImplementationForTestReceivedTokens // Event containing the contract specifics and raw log

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
func (it *StakingPoolImplementationForTestReceivedTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolImplementationForTestReceivedTokens)
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
		it.Event = new(StakingPoolImplementationForTestReceivedTokens)
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
func (it *StakingPoolImplementationForTestReceivedTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolImplementationForTestReceivedTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolImplementationForTestReceivedTokens represents a ReceivedTokens event raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestReceivedTokens struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReceivedTokens is a free log retrieval operation binding the contract event 0x2946de6c4ec03d8d15126164a7c0da68d7c6835173e41827a7a715f8becb07a8.
//
// Solidity: event ReceivedTokens(address indexed from, uint256 amount)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) FilterReceivedTokens(opts *bind.FilterOpts, from []common.Address) (*StakingPoolImplementationForTestReceivedTokensIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _StakingPoolImplementationForTest.contract.FilterLogs(opts, "ReceivedTokens", fromRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationForTestReceivedTokensIterator{contract: _StakingPoolImplementationForTest.contract, event: "ReceivedTokens", logs: logs, sub: sub}, nil
}

// WatchReceivedTokens is a free log subscription operation binding the contract event 0x2946de6c4ec03d8d15126164a7c0da68d7c6835173e41827a7a715f8becb07a8.
//
// Solidity: event ReceivedTokens(address indexed from, uint256 amount)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) WatchReceivedTokens(opts *bind.WatchOpts, sink chan<- *StakingPoolImplementationForTestReceivedTokens, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _StakingPoolImplementationForTest.contract.WatchLogs(opts, "ReceivedTokens", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolImplementationForTestReceivedTokens)
				if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "ReceivedTokens", log); err != nil {
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
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) ParseReceivedTokens(log types.Log) (*StakingPoolImplementationForTestReceivedTokens, error) {
	event := new(StakingPoolImplementationForTestReceivedTokens)
	if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "ReceivedTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolImplementationForTestStakeIterator is returned from FilterStake and is used to iterate over the raw logs and unpacked data for Stake events raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestStakeIterator struct {
	Event *StakingPoolImplementationForTestStake // Event containing the contract specifics and raw log

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
func (it *StakingPoolImplementationForTestStakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolImplementationForTestStake)
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
		it.Event = new(StakingPoolImplementationForTestStake)
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
func (it *StakingPoolImplementationForTestStakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolImplementationForTestStakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolImplementationForTestStake represents a Stake event raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestStake struct {
	User            common.Address
	Amount          *big.Int
	PricePrediction *big.Int
	IndexOfStake    *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStake is a free log retrieval operation binding the contract event 0xf556991011e831bcfac4f406d547e5e32cdd98267efab83935230d5f8d02c446.
//
// Solidity: event Stake(address indexed user, uint256 amount, uint256 pricePrediction, uint256 indexOfStake)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) FilterStake(opts *bind.FilterOpts, user []common.Address) (*StakingPoolImplementationForTestStakeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingPoolImplementationForTest.contract.FilterLogs(opts, "Stake", userRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationForTestStakeIterator{contract: _StakingPoolImplementationForTest.contract, event: "Stake", logs: logs, sub: sub}, nil
}

// WatchStake is a free log subscription operation binding the contract event 0xf556991011e831bcfac4f406d547e5e32cdd98267efab83935230d5f8d02c446.
//
// Solidity: event Stake(address indexed user, uint256 amount, uint256 pricePrediction, uint256 indexOfStake)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) WatchStake(opts *bind.WatchOpts, sink chan<- *StakingPoolImplementationForTestStake, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingPoolImplementationForTest.contract.WatchLogs(opts, "Stake", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolImplementationForTestStake)
				if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "Stake", log); err != nil {
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

// ParseStake is a log parse operation binding the contract event 0xf556991011e831bcfac4f406d547e5e32cdd98267efab83935230d5f8d02c446.
//
// Solidity: event Stake(address indexed user, uint256 amount, uint256 pricePrediction, uint256 indexOfStake)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) ParseStake(log types.Log) (*StakingPoolImplementationForTestStake, error) {
	event := new(StakingPoolImplementationForTestStake)
	if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "Stake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolImplementationForTestTransferTokensIterator is returned from FilterTransferTokens and is used to iterate over the raw logs and unpacked data for TransferTokens events raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestTransferTokensIterator struct {
	Event *StakingPoolImplementationForTestTransferTokens // Event containing the contract specifics and raw log

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
func (it *StakingPoolImplementationForTestTransferTokensIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolImplementationForTestTransferTokens)
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
		it.Event = new(StakingPoolImplementationForTestTransferTokens)
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
func (it *StakingPoolImplementationForTestTransferTokensIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolImplementationForTestTransferTokensIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolImplementationForTestTransferTokens represents a TransferTokens event raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestTransferTokens struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferTokens is a free log retrieval operation binding the contract event 0x7cbfb8a5c69722670db81365c97149301df8bf0b0afe48f73787d6a6a4385954.
//
// Solidity: event TransferTokens(address indexed to, uint256 amount)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) FilterTransferTokens(opts *bind.FilterOpts, to []common.Address) (*StakingPoolImplementationForTestTransferTokensIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakingPoolImplementationForTest.contract.FilterLogs(opts, "TransferTokens", toRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationForTestTransferTokensIterator{contract: _StakingPoolImplementationForTest.contract, event: "TransferTokens", logs: logs, sub: sub}, nil
}

// WatchTransferTokens is a free log subscription operation binding the contract event 0x7cbfb8a5c69722670db81365c97149301df8bf0b0afe48f73787d6a6a4385954.
//
// Solidity: event TransferTokens(address indexed to, uint256 amount)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) WatchTransferTokens(opts *bind.WatchOpts, sink chan<- *StakingPoolImplementationForTestTransferTokens, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StakingPoolImplementationForTest.contract.WatchLogs(opts, "TransferTokens", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolImplementationForTestTransferTokens)
				if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "TransferTokens", log); err != nil {
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
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) ParseTransferTokens(log types.Log) (*StakingPoolImplementationForTestTransferTokens, error) {
	event := new(StakingPoolImplementationForTestTransferTokens)
	if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "TransferTokens", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolImplementationForTestUnstakeIterator is returned from FilterUnstake and is used to iterate over the raw logs and unpacked data for Unstake events raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestUnstakeIterator struct {
	Event *StakingPoolImplementationForTestUnstake // Event containing the contract specifics and raw log

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
func (it *StakingPoolImplementationForTestUnstakeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolImplementationForTestUnstake)
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
		it.Event = new(StakingPoolImplementationForTestUnstake)
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
func (it *StakingPoolImplementationForTestUnstakeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolImplementationForTestUnstakeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolImplementationForTestUnstake represents a Unstake event raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestUnstake struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterUnstake is a free log retrieval operation binding the contract event 0x85082129d87b2fe11527cb1b3b7a520aeb5aa6913f88a3d8757fe40d1db02fdd.
//
// Solidity: event Unstake(address indexed user, uint256 amount)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) FilterUnstake(opts *bind.FilterOpts, user []common.Address) (*StakingPoolImplementationForTestUnstakeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingPoolImplementationForTest.contract.FilterLogs(opts, "Unstake", userRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationForTestUnstakeIterator{contract: _StakingPoolImplementationForTest.contract, event: "Unstake", logs: logs, sub: sub}, nil
}

// WatchUnstake is a free log subscription operation binding the contract event 0x85082129d87b2fe11527cb1b3b7a520aeb5aa6913f88a3d8757fe40d1db02fdd.
//
// Solidity: event Unstake(address indexed user, uint256 amount)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) WatchUnstake(opts *bind.WatchOpts, sink chan<- *StakingPoolImplementationForTestUnstake, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingPoolImplementationForTest.contract.WatchLogs(opts, "Unstake", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolImplementationForTestUnstake)
				if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "Unstake", log); err != nil {
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

// ParseUnstake is a log parse operation binding the contract event 0x85082129d87b2fe11527cb1b3b7a520aeb5aa6913f88a3d8757fe40d1db02fdd.
//
// Solidity: event Unstake(address indexed user, uint256 amount)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) ParseUnstake(log types.Log) (*StakingPoolImplementationForTestUnstake, error) {
	event := new(StakingPoolImplementationForTestUnstake)
	if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "Unstake", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolImplementationForTestWithdrawStakingReturnIterator is returned from FilterWithdrawStakingReturn and is used to iterate over the raw logs and unpacked data for WithdrawStakingReturn events raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestWithdrawStakingReturnIterator struct {
	Event *StakingPoolImplementationForTestWithdrawStakingReturn // Event containing the contract specifics and raw log

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
func (it *StakingPoolImplementationForTestWithdrawStakingReturnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolImplementationForTestWithdrawStakingReturn)
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
		it.Event = new(StakingPoolImplementationForTestWithdrawStakingReturn)
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
func (it *StakingPoolImplementationForTestWithdrawStakingReturnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolImplementationForTestWithdrawStakingReturnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolImplementationForTestWithdrawStakingReturn represents a WithdrawStakingReturn event raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestWithdrawStakingReturn struct {
	User          common.Address
	StakingReturn *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterWithdrawStakingReturn is a free log retrieval operation binding the contract event 0x115e2fe760f98d4fd0828d7741ae2070e7db700f892efbfe98448b53ac583886.
//
// Solidity: event WithdrawStakingReturn(address indexed user, uint256 stakingReturn)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) FilterWithdrawStakingReturn(opts *bind.FilterOpts, user []common.Address) (*StakingPoolImplementationForTestWithdrawStakingReturnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingPoolImplementationForTest.contract.FilterLogs(opts, "WithdrawStakingReturn", userRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationForTestWithdrawStakingReturnIterator{contract: _StakingPoolImplementationForTest.contract, event: "WithdrawStakingReturn", logs: logs, sub: sub}, nil
}

// WatchWithdrawStakingReturn is a free log subscription operation binding the contract event 0x115e2fe760f98d4fd0828d7741ae2070e7db700f892efbfe98448b53ac583886.
//
// Solidity: event WithdrawStakingReturn(address indexed user, uint256 stakingReturn)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) WatchWithdrawStakingReturn(opts *bind.WatchOpts, sink chan<- *StakingPoolImplementationForTestWithdrawStakingReturn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingPoolImplementationForTest.contract.WatchLogs(opts, "WithdrawStakingReturn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolImplementationForTestWithdrawStakingReturn)
				if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "WithdrawStakingReturn", log); err != nil {
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

// ParseWithdrawStakingReturn is a log parse operation binding the contract event 0x115e2fe760f98d4fd0828d7741ae2070e7db700f892efbfe98448b53ac583886.
//
// Solidity: event WithdrawStakingReturn(address indexed user, uint256 stakingReturn)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) ParseWithdrawStakingReturn(log types.Log) (*StakingPoolImplementationForTestWithdrawStakingReturn, error) {
	event := new(StakingPoolImplementationForTestWithdrawStakingReturn)
	if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "WithdrawStakingReturn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolImplementationForTestWithdrawTotemPrizeIterator is returned from FilterWithdrawTotemPrize and is used to iterate over the raw logs and unpacked data for WithdrawTotemPrize events raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestWithdrawTotemPrizeIterator struct {
	Event *StakingPoolImplementationForTestWithdrawTotemPrize // Event containing the contract specifics and raw log

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
func (it *StakingPoolImplementationForTestWithdrawTotemPrizeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolImplementationForTestWithdrawTotemPrize)
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
		it.Event = new(StakingPoolImplementationForTestWithdrawTotemPrize)
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
func (it *StakingPoolImplementationForTestWithdrawTotemPrizeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolImplementationForTestWithdrawTotemPrizeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolImplementationForTestWithdrawTotemPrize represents a WithdrawTotemPrize event raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestWithdrawTotemPrize struct {
	User       common.Address
	TotemPrize *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawTotemPrize is a free log retrieval operation binding the contract event 0x723b06f6c3a94dc82019852e6aa01c0241fba5f84410b9003ddc3440ad499e59.
//
// Solidity: event WithdrawTotemPrize(address indexed user, uint256 totemPrize)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) FilterWithdrawTotemPrize(opts *bind.FilterOpts, user []common.Address) (*StakingPoolImplementationForTestWithdrawTotemPrizeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingPoolImplementationForTest.contract.FilterLogs(opts, "WithdrawTotemPrize", userRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationForTestWithdrawTotemPrizeIterator{contract: _StakingPoolImplementationForTest.contract, event: "WithdrawTotemPrize", logs: logs, sub: sub}, nil
}

// WatchWithdrawTotemPrize is a free log subscription operation binding the contract event 0x723b06f6c3a94dc82019852e6aa01c0241fba5f84410b9003ddc3440ad499e59.
//
// Solidity: event WithdrawTotemPrize(address indexed user, uint256 totemPrize)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) WatchWithdrawTotemPrize(opts *bind.WatchOpts, sink chan<- *StakingPoolImplementationForTestWithdrawTotemPrize, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingPoolImplementationForTest.contract.WatchLogs(opts, "WithdrawTotemPrize", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolImplementationForTestWithdrawTotemPrize)
				if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "WithdrawTotemPrize", log); err != nil {
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

// ParseWithdrawTotemPrize is a log parse operation binding the contract event 0x723b06f6c3a94dc82019852e6aa01c0241fba5f84410b9003ddc3440ad499e59.
//
// Solidity: event WithdrawTotemPrize(address indexed user, uint256 totemPrize)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) ParseWithdrawTotemPrize(log types.Log) (*StakingPoolImplementationForTestWithdrawTotemPrize, error) {
	event := new(StakingPoolImplementationForTestWithdrawTotemPrize)
	if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "WithdrawTotemPrize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPoolImplementationForTestWithdrawWrappedTokenPrizeIterator is returned from FilterWithdrawWrappedTokenPrize and is used to iterate over the raw logs and unpacked data for WithdrawWrappedTokenPrize events raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestWithdrawWrappedTokenPrizeIterator struct {
	Event *StakingPoolImplementationForTestWithdrawWrappedTokenPrize // Event containing the contract specifics and raw log

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
func (it *StakingPoolImplementationForTestWithdrawWrappedTokenPrizeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPoolImplementationForTestWithdrawWrappedTokenPrize)
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
		it.Event = new(StakingPoolImplementationForTestWithdrawWrappedTokenPrize)
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
func (it *StakingPoolImplementationForTestWithdrawWrappedTokenPrizeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPoolImplementationForTestWithdrawWrappedTokenPrizeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPoolImplementationForTestWithdrawWrappedTokenPrize represents a WithdrawWrappedTokenPrize event raised by the StakingPoolImplementationForTest contract.
type StakingPoolImplementationForTestWithdrawWrappedTokenPrize struct {
	User              common.Address
	WrappedTokenPrize *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterWithdrawWrappedTokenPrize is a free log retrieval operation binding the contract event 0xd93fc6d1e8779271f55cdb0d5f7b6b05be64a2226a1ad050d3036f5609c1cdcf.
//
// Solidity: event WithdrawWrappedTokenPrize(address indexed user, uint256 wrappedTokenPrize)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) FilterWithdrawWrappedTokenPrize(opts *bind.FilterOpts, user []common.Address) (*StakingPoolImplementationForTestWithdrawWrappedTokenPrizeIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingPoolImplementationForTest.contract.FilterLogs(opts, "WithdrawWrappedTokenPrize", userRule)
	if err != nil {
		return nil, err
	}
	return &StakingPoolImplementationForTestWithdrawWrappedTokenPrizeIterator{contract: _StakingPoolImplementationForTest.contract, event: "WithdrawWrappedTokenPrize", logs: logs, sub: sub}, nil
}

// WatchWithdrawWrappedTokenPrize is a free log subscription operation binding the contract event 0xd93fc6d1e8779271f55cdb0d5f7b6b05be64a2226a1ad050d3036f5609c1cdcf.
//
// Solidity: event WithdrawWrappedTokenPrize(address indexed user, uint256 wrappedTokenPrize)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) WatchWithdrawWrappedTokenPrize(opts *bind.WatchOpts, sink chan<- *StakingPoolImplementationForTestWithdrawWrappedTokenPrize, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _StakingPoolImplementationForTest.contract.WatchLogs(opts, "WithdrawWrappedTokenPrize", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPoolImplementationForTestWithdrawWrappedTokenPrize)
				if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "WithdrawWrappedTokenPrize", log); err != nil {
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

// ParseWithdrawWrappedTokenPrize is a log parse operation binding the contract event 0xd93fc6d1e8779271f55cdb0d5f7b6b05be64a2226a1ad050d3036f5609c1cdcf.
//
// Solidity: event WithdrawWrappedTokenPrize(address indexed user, uint256 wrappedTokenPrize)
func (_StakingPoolImplementationForTest *StakingPoolImplementationForTestFilterer) ParseWithdrawWrappedTokenPrize(log types.Log) (*StakingPoolImplementationForTestWithdrawWrappedTokenPrize, error) {
	event := new(StakingPoolImplementationForTestWithdrawWrappedTokenPrize)
	if err := _StakingPoolImplementationForTest.contract.UnpackLog(event, "WithdrawWrappedTokenPrize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
