// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

interface IRewardManager {

    function setOperator(address _newOperator) external;

    function addPool(address _poolAddress) external;

    function rewardUser(address _user, uint256 _amount) external;

    event SetOperator(address operator);
    event SetRewarder(address rewarder);

}