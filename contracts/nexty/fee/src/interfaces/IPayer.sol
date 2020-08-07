// SPDX-License-Identifier: MIT
pragma solidity ^0.6.0;

import "./IERC20.sol";

/**
 * (callcode)
 */
interface IPayer {
    function pay(
        address coinbase,
        address txTo,
        uint txGasPrice,
        uint gasToPay
    ) external;
}
