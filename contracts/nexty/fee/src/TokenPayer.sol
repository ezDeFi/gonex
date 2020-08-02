// SPDX-License-Identifier: MIT
pragma solidity ^0.6.0;

import "./lib/ds.sol";
import "./interfaces/IERC20.sol";
import "./interfaces/IPayer.sol";
import "./TokenPrice.sol";

/**
 * Will be deployed by the consensus to 0x56789
 * (delegated)
 */
contract TokenPayer is IPayer {
    address public constant TokenPricesContract = address(0x6789A);
    bytes32 public constant TokenPricePath = 'TokenPrice-';             // bytes11: user price for each token
    bytes32 public constant FeeTokenPath = 'FeeToken';                  // if this key is set, the value is always use for token fee
    bytes32 public constant FeeTokenFallbackPath = 'FeeTokenFallback';  // if all else fall, use this token to pay
    uint public constant ERC20PayGasLimit = 34000;                      // gas limit for typical ERC20 transfer is 32k

    function getPrice(address to)
        internal
        view
        returns (uint price)
    {
        // allows user to override the prices using account personal storage
        uint personalPrice = uint(ds.load(ds.key11a(TokenPricePath, to)));
        if (personalPrice > 0) {
            return personalPrice;
        }
        // then load the price from the global contract
        return TokenPrice(TokenPricesContract).getPrice(to);
    }

    function payment(
        address to,
        uint txGas,                         // ignored
        bytes4 txMethodSig,                 // ignored
        bytes32[] calldata txMethodParams   // ignored
    )
        external
        view
        override
        returns (
            address token,
            uint price,         // Token/NTY price (in wei) << 128
            uint payGasLimit    // gasLimit to call and pre-paid for method pay
        )
    {
        // first check the explicit fee token setting
        token = ds.loadAddress(FeeTokenPath);
        if (token != address(0x0)) {
            price = getPrice(token);
            if (price > 0) {
                return (token, price, ERC20PayGasLimit);
            }
        }
        // then try the interacting token
        price = getPrice(to);
        if (price > 0) {
            return (to, price, ERC20PayGasLimit);
        }
        // if all else fail, try the fallback fee token setting
        token = ds.loadAddress(FeeTokenFallbackPath);
        if (token != address(0x0)) {
            price = getPrice(token);
            if (price > 0) {
                return (token, price, ERC20PayGasLimit);
            }
        }
        // no token for payment
        return (address(0x0), 0, 0);
    }

    function pay(
        address coinbase,
        address token,
        uint value
    )
        external
        override
    {
        IERC20(token).transfer(coinbase, value);
    }
}
