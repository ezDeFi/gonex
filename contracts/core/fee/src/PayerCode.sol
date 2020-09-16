// SPDX-License-Identifier: MIT
pragma solidity ^0.6.0;

import "./lib/ds.sol";
import "./interfaces/IERC20.sol";
import "./interfaces/IPayer.sol";
import "./interfaces/IConfig.sol";

/**
 * Will be deployed by the consensus to 0x56789
 * (callcode)
 */
abstract contract PayerCode is IPayer {
    IConfig public constant TokenPayment            = IConfig(0x56789);
    bytes32 public constant FeeTokenKey             = 'FeeToken';           // if this key is set, the value is always use for token fee
    bytes32 public constant FeeTokenFallbackKey     = 'FeeTokenFallback';   // if all else fall, use this token to pay
    bytes32 public constant AppTokenPath            = 'AppToken---';        // bytes11: token for a specific address
    bytes32 public constant TokenPricePath          = 'TokenPrice-';        // bytes11: user price for each token

    /**
     * @dev perfrom no overflow check at all, miner need to check to protect their own income
     *
     * (callcode)
     */
    function pay(uint fee, address txTo) external override {
        (address token, uint price) = _payment(txTo);
        require(price > 0, "payment token price not set");
        // require(fee * (10**18) / (10**18) == fee, "token overflow");
        uint tokenToPay = fee * (10**18) / price; // unsafe
        // require(tokenToPay > 0, "zero token payment");
        IERC20(token).transfer(block.coinbase, tokenToPay);
    }

    function _payment(address txTo) internal view returns (
        address token,
        uint price  // Token/ZD price (in wei) (decimals = 18)
    ) {
        // first check the fee token setting
        token = ds.loadAddress(FeeTokenKey);
        if (token != address(0x0)) {
            return (token, _getAnyPrice(token));
        }
        // check the personal app token mapping
        token = _getAppToken(txTo);
        if (token != address(0x0)) {
            return (token, _getAnyPrice(token));
        }
        // get the global app token and price
        (token, price) = TokenPayment.getAppTokenAndPrice(txTo);
        // get the personal price to override the global price
        uint p = _getPrice(token);
        if (p > 0) {
            return (token, p); // global app token and personal price
        }
        if (price > 0) {
            return (token, price); // global app token and global price
        }
        // if all else fail, try the fallback fee token setting
        token = ds.loadAddress(FeeTokenFallbackKey);
        require(token != address(0x0), "no token for payment");
        return (token, _getAnyPrice(token));
    }

    // get personal price and then global price
    function _getAnyPrice(address token) internal view returns (uint price) {
        // allows user to override the prices using account personal storage
        uint price = _getPrice(token);
        if (price > 0) {
            return price;
        }
        // then load the price from the global contract
        return TokenPayment.getPrice(token);
    }

    function _getAppToken(address app) internal view returns (address) {
        return ds.loadAddress(ds.key11a(AppTokenPath, app));
    }

    function _getPrice(address token) internal view returns (uint price) {
        return uint(ds.load(ds.key11a(TokenPricePath, token)));
    }

    function _setAppToken(address app, address token) internal {
        ds.storeAddress(ds.key11a(AppTokenPath, app), token);
    }

    function _setPrice(address token, uint price) internal {
        ds.store(ds.key11a(TokenPricePath, token), bytes32(price));
    }
}
