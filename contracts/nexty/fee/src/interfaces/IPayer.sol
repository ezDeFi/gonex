// SPDX-License-Identifier: MIT
pragma solidity ^0.6.0;

import "./IERC20.sol";

/**
 * (callcode)
 */
interface IPayer {
    function pay(
        address coinbase,
        address to,
        uint fee
    ) external;
}
