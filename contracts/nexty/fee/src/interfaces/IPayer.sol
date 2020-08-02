// SPDX-License-Identifier: MIT
pragma solidity ^0.6.0;

import "./IERC20.sol";

/**
 * (delegated)
 */
interface IPayer {
    function payment(
        address to,
        uint txGas,
        bytes4 txMethodSig,
        bytes32[] calldata txMethodParams
    )
        external
        view
        returns (
            address token,
            uint price,         // Token/NTY price (in wei) << 128
            uint payGasLimit    // gasLimit to call and pre-paid for method pay
        );

    function pay(
        address coinbase,
        address token,
        uint value
    )
        external;
}
