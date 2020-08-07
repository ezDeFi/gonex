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
    bytes32 public constant TokenPricePath          = 'TokenPrice-';        // bytes11: user price for each token
    bytes32 public constant FeeTokenPath            = 'FeeToken';           // if this key is set, the value is always use for token fee
    bytes32 public constant FeeTokenFallbackPath    = 'FeeTokenFallback';   // if all else fall, use this token to pay

    /**
     * @dev perfrom no overflow check at all, miner should do that to protect their income
     *
     * (callcode)
     */
    function pay(
        address coinbase,
        address to,
        uint fee
    ) external override {
        (address token, uint price) = _payment(to);
        require(price > 0, "payment token price not set");
        // require(fee * (10**18) / (10**18) == fee, "token overflow");
        uint tokenToPay = fee * (10**18) / price; // unsafe
        // require(tokenToPay > 0, "zero token payment");
        IERC20(token).transfer(coinbase, tokenToPay);
    }

    // (callcode)
    function _payment(
        address to
    ) internal view returns (
        address token,
        uint price  // Token/NTY price (in wei) (decimals = 18)
    ) {
        token = _getTokenToPay(to);
        require(token != address(0x0), "no token for payment");
        return (token, _getPrice(token));
    }

    // (callcode)
    function _getTokenToPay(address to) internal view returns (address token) {
        // first check the fee token setting
        token = ds.loadAddress(FeeTokenPath);
        if (token != address(0x0)) {
            return token;
        }
        // then try the interacting token
        token = TokenPayment.getTokenFor(to);
        if (token != address(0x0)) {
            return token;
        }
        // if all else fail, try the fallback fee token setting
        return ds.loadAddress(FeeTokenFallbackPath);
    }

    // (callcode))
    function _getPrice(address token) internal view returns (uint price) {
        // allows user to override the prices using account personal storage
        uint personalPrice = uint(ds.load(ds.key11a(TokenPricePath, token)));
        if (personalPrice > 0) {
            return personalPrice;
        }
        // then load the price from the global contract
        return TokenPayment.getPrice(token);
    }
}
