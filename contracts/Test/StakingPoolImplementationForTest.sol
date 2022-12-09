// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;


import "../Staking/StakingPoolImplementation.sol";

contract StakingPoolImplementationForTest is StakingPoolImplementation {
    event PoolUnLocked();

    
    // TODO: we should decide which functions we may want to override so that we can add 
    // "virtual" to them
    /**
     * @dev If pool is locked unwantedly, we can reverse it like this
     */
    function lockPool() public onlyPoolCreator override {
        _unlockPool();
    }

    function _unlockPool() internal {
        isLocked = false;

        emit PoolUnLocked();
    }
}