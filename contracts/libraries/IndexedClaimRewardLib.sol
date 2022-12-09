// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "./CalculateRewardLib.sol";
import "../libraries/BasisPoints.sol";
import "../Staking/StakingPoolStorageStructure.sol";

library IndexedClaimRewardLib {

    using CalculateRewardLib for *;
    using BasisPoints for uint256; 
    using SafeMath for uint256;

    uint256 public constant foo = 0;

    function withdrawIndexedStakingReturn(
        StakingPoolStorageStructure.StakeWithPrediction[] storage _staker, 
        uint256 _stakeIndex,
        StakingPoolStorageStructure.LibParams storage _lps
    ) 
        public
    {
        if (_staker.length == 0) return;
        if (_stakeIndex >= _staker.length) return;

        uint256 rewardPerStake = CalculateRewardLib._getStakingRewardPerStake(
            _staker, 
            _stakeIndex,
            _lps
        );

        _staker[_stakeIndex].lastWithdrawalTime = block.timestamp;
        _staker[_stakeIndex].amountWithdrawn = _staker[_stakeIndex].amountWithdrawn.add(
            rewardPerStake
        );
    }

    function withdrawIndexedPrize(
        StakingPoolStorageStructure.StakeWithPrediction[] storage _staker, 
        uint256 _stakeIndex
    ) 
        public 
    {
        if (_staker.length == 0) return;
        if (_staker[_stakeIndex].prizeRewardWithdrawn) return;

        _staker[_stakeIndex].prizeRewardWithdrawn = true;
    }

    function withdrawIndexedStakedBalance(
        StakingPoolStorageStructure.StakeWithPrediction[] storage _staker, 
        uint256 _stakeIndex
    ) 
        public
    {
        if (_staker.length == 0) return;
        if (_stakeIndex >= _staker.length) return;

        _staker[_stakeIndex].didUnstake = true;
    }

    function getIndexedStakedBalance(
        StakingPoolStorageStructure.StakeWithPrediction[] storage _staker, 
        uint256 _stakeIndex
    )
        public
        view
        returns (uint256)
    {
        if (_staker.length == 0) return 0;
        if (_stakeIndex >= _staker.length) return 0; 

        uint256 totalStakedBalance = 0;

        if (!_staker[_stakeIndex].didUnstake) {
            totalStakedBalance = totalStakedBalance.add(
                _staker[_stakeIndex].stakedBalance
            );
        }

        return totalStakedBalance;
    }

    function getIndexedStakingReturn(
        StakingPoolStorageStructure.StakeWithPrediction[] storage _staker, 
        uint256 _stakeIndex,
        StakingPoolStorageStructure.LibParams storage _lps
    ) 
        public
        view 
        returns (uint256) 
    {
        if (_staker.length == 0) return 0;
        if (_stakeIndex >= _staker.length) return 0;

        uint256 reward = 0;
        
        uint256 rewardPerStake = CalculateRewardLib._getStakingRewardPerStake(
            _staker, 
            _stakeIndex,
            _lps
        );
        reward = reward.add(rewardPerStake);

        return reward;
    }

    function getIndexedPrize(
        StakingPoolStorageStructure.StakeWithPrediction[] storage _staker, 
        uint256 _stakeIndex,
        StakingPoolStorageStructure.LibParams storage _lps,
        StakingPoolStorageStructure.PrizeRewardRate[] storage _prizeRewardRates
    )
        public
        view
        returns (uint256, uint256)
    {
        if (!_lps.isMatured) return (0, 0);

        if (_staker.length == 0) return (0, 0);

        if (_stakeIndex >= _staker.length) return (0,0);

        if (_staker[_stakeIndex].prizeRewardWithdrawn) return (0, 0);

        uint256 maturingWrappedTokenPrizeAmount =
            (_lps.usdPrizeAmount.mul(10**_lps.oracleDecimals)).div(_lps.maturingPrice);

        uint256 reward = 0;
        uint256 wrappedTokenReward = 0;

        uint256 _percent = CalculateRewardLib._getPercentageReward(
            _staker[_stakeIndex].rank,
            _prizeRewardRates
        );

        reward = reward.add(
                        _lps.prizeAmount.mulBP(_percent)
                    );

        wrappedTokenReward = wrappedTokenReward.add(
                        maturingWrappedTokenPrizeAmount
                            .mulBP(_percent)
                    );            

        if (_lps.collaborativeReward > 0) {
            reward = reward.addBP(_lps.collaborativeReward);
            wrappedTokenReward = wrappedTokenReward.addBP(_lps.collaborativeReward);
        }

        return (reward, wrappedTokenReward);
    }
}

