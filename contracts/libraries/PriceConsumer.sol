// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

import "@chainlink/contracts/src/v0.8/interfaces/AggregatorV3Interface.sol";

library PriceConsumer {
    /**
     * @param oracle The chainlink node oracle address to send requests
     * @notice Returns decimals for oracle contract
     */
    function getDecimals(address oracle) internal view returns (uint8) {
        AggregatorV3Interface priceFeed = AggregatorV3Interface(oracle);
        uint8 decimals = priceFeed.decimals();
        return decimals;
    }

    /**
     * @notice Returns the latest price from oracle contract
     */
    function getLatestPrice(address oracle) internal view returns (uint256) {
        AggregatorV3Interface priceFeed = AggregatorV3Interface(oracle);
        (, int256 price, , , ) = priceFeed.latestRoundData();

        return price >= 0 ? uint256(price) : 0;
    }
}
