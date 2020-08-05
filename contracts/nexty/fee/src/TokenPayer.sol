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

    /**
     * @dev perfrom no overflow check at all, tx pool will check it anyway
     */
    function pay(
        address coinbase,
        address txTo,
        uint txGas,
        uint txGasPrice
    ) external override {
        (address token, uint price) = _payment(txTo);
        // require(token != address(0x0) && price > 0, "payment not available");
        uint fee = txGas * txGasPrice;
        // require(fee > 0 && fee / txGas == txGasPrice, "fee overflow");
        // require(fee * (10**18) / (10**18) == fee, "token overflow");
        uint tokenToPay = fee * (10**18) / price;
        // require(tokenToPay > 0, "zero token fee");
        IERC20(token).transfer(coinbase, tokenToPay);
    }

    function _payment(
        address to
    ) internal view returns (
            address token,
            uint price  // Token/NTY price (in wei) (decimals = 18)
        )
    {
        // first check the explicit fee token setting
        token = ds.loadAddress(FeeTokenPath);
        if (token != address(0x0)) {
            price = getPrice(token);
            if (price > 0) {
                return (token, price);
            }
        }
        // then try the interacting token
        price = getPrice(to);
        if (price > 0) {
            return (to, price);
        }
        // if all else fail, try the fallback fee token setting
        token = ds.loadAddress(FeeTokenFallbackPath);
        if (token != address(0x0)) {
            price = getPrice(token);
            if (price > 0) {
                return (token, price);
            }
        }
        // no token for payment
        return (address(0x0), 0);
    }
}
