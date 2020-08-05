// SPDX-License-Identifier: MIT
pragma solidity ^0.6.0;

import "./IERC20.sol";

/**
 * (delegated)
 */
interface IPayer {
    function pay(
        address coinbase,
        uint paymentGas,
        address txTo,
        uint txGas,
        uint txGasPrice
    ) external;
}
