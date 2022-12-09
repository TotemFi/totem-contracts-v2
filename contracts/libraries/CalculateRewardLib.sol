// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "../libraries/BasisPoints.sol";
import "../Staking/StakingPoolStorageStructure.sol";

library CalculateRewardLib {

    using BasisPoints for uint256;
    using SafeMath for uint256;

    uint256 public constant foo = 0;

    function getTotalStakedBalance(StakingPoolStorageStructure.StakeWithPrediction[] storage _staker)
        public
        view
        returns (uint256)
    {
        if (_staker.length == 0) return 0;

        uint256 totalStakedBalance = 0;
        for (uint256 i = 0; i < _staker.length; i++) {
            if (!_staker[i].didUnstake) {
                totalStakedBalance = totalStakedBalance.add(
                    _staker[i].stakedBalance
                );
            }
        }

        return totalStakedBalance;
    }

    /**
     * @notice the reward formula is:
          ((1 + stakeAPR +enhancedReward)^((MaturingDate - StakingDate)/365) - 1) * StakingBalance
    */
    function _getStakingRewardPerStake(
        StakingPoolStorageStructure.StakeWithPrediction[] storage _staker, 
        uint256 _stakeIndex,
        StakingPoolStorageStructure.LibParams storage _lps
    )
        internal
        view
        returns (uint256)
    {
        uint256 maturityDate = 
            _lps.launchDate + 
            _lps.lockTime + 
            _lps.maturityTime;

        uint256 timeTo =
            block.timestamp > maturityDate ? maturityDate : block.timestamp;

        uint256 enhancedApr;
        if ( _lps.isEnhancedEnabled ) {
            enhancedApr = _getEnhancedRewardRate(
                _staker[_stakeIndex].stakedTime,
                _lps
            );
        }

        uint256 rewardPerStake = _calcStakingReturn(
            _lps.stakeApr.add(enhancedApr),
            timeTo.sub(_staker[_stakeIndex].stakedTime),
            _staker[_stakeIndex].stakedBalance
        );

        rewardPerStake = rewardPerStake.sub(_staker[_stakeIndex].amountWithdrawn);

        return rewardPerStake;
    }

    function _getEnhancedRewardRate(
        uint256 stakedTime,
        StakingPoolStorageStructure.LibParams storage _lps
    )
        internal
        view
        returns (uint256)
    {

        if (!_lps.isEnhancedEnabled) {
            return 0;
        }

        uint256 lockDate = _lps.launchDate.add(_lps.lockTime);
        uint256 difference = lockDate.sub(stakedTime);

        if (difference < 48 hours) {
            return 0;
        } else if (difference < 72 hours) {
            return 100;
        } else if (difference < 96 hours) {
            return 200;
        } else if (difference < 120 hours) {
            return 300;
        } else if (difference < 144 hours) {
            return 400;
        } else {
            return 500;
        }
    }

    function _calcStakingReturn(uint256 totalRewardRate, uint256 timeDuration, uint256 totalStakedBalance) 
        internal 
        pure
        returns (uint256) 
    {
        uint256 yearInSeconds = 365 days;

        uint256 first = (yearInSeconds**2)
            .mul(10**8);

        uint256 second = timeDuration
            .mul(totalRewardRate) 
            .mul(yearInSeconds)
            .mul(5000);
        
        uint256 third = totalRewardRate
            .mul(yearInSeconds**2)
            .mul(5000);

        uint256 forth = (timeDuration**2)
            .mul(totalRewardRate**2)
            .div(6);

        uint256 fifth = timeDuration
            .mul(totalRewardRate**2)
            .mul(yearInSeconds)
            .div(2);

        uint256 sixth = (totalRewardRate**2)
            .mul(yearInSeconds**2)
            .div(3);
 
        uint256 rewardPerStake = first.add(second).add(forth).add(sixth);

        rewardPerStake = rewardPerStake.sub(third).sub(fifth);

        rewardPerStake = rewardPerStake
            .mul(totalRewardRate)
            .mul(timeDuration);

        rewardPerStake = rewardPerStake
            .mul(totalStakedBalance)
            .div(yearInSeconds**3)
            .div(10**12);

        return rewardPerStake; 
    }

    function _getPercentageReward(
        uint256 _rank, 
        StakingPoolStorageStructure.PrizeRewardRate[] storage _prizeRewardRates
    )
        internal
        view
        returns (uint256)
    {
        for (uint256 i = 0; i < _prizeRewardRates.length; i++) {
            if (_rank <= _prizeRewardRates[i].rank) {
                return _prizeRewardRates[i].percentage;
            }
        }

        return 0;
    }        



}