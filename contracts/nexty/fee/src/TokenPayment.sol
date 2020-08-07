// SPDX-License-Identifier: MIT
pragma solidity ^0.6.0;

import "./lib/ds.sol";
import "./interfaces/IConfig.sol";
import "./PayerCode.sol";

contract TokenPayment is PayerCode, IConfig {
    bytes32 public constant TRUE = 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF;
    bytes32 public constant TeamMemberPath  = 'TeamMember-'; // bytes11
    bytes32 public constant TokenForPath    = 'TokenFor---'; // bytes11: token for a specific address

    constructor(
        address[]   memory admins,    // trusted team to set the token prices
        address[]   memory tokens,  // initial mapping(token address => price) keys
        uint[]      memory prices   // initial mapping(token address => price) values
    ) public {
        for (uint i = 0; i < admins.length; i++) {
            _addAdmin(admins[i]);
        }
        for (uint i = 0; i < tokens.length; i++) {
            _setPrice(tokens[i], prices[i]);
        }
    }

    function _addAdmin(address member) internal {
        ds.store(ds.key11a(TeamMemberPath, member), TRUE);
    }

    function _isAdmin(address member) internal view returns (bool) {
        return ds.load(ds.key11a(TeamMemberPath, member)) != 0;
    }

    function getTokenFor(address to) external override view returns (address token) {
        token = _getTokenFor(to);
        if (token != address(0x0)) {
            return token;
        }
        return to;
    }

    function setTokenFor(address to, address token) external {
        require(_isAdmin(msg.sender), "for team member only");
        _setTokenFor(to, token);
    }

    function _getTokenFor(address to) internal view returns (address) {
        return ds.loadAddress(ds.key11a(TokenForPath, to));
    }

    function _setTokenFor(address to, address token) internal {
        ds.storeAddress(ds.key11a(TokenForPath, to), token);
    }

    function getPrice(address token) external override view returns (uint price) {
        return uint(ds.load(ds.key11a(TokenPricePath, token)));
    }

    function setPrice(address token, uint price) external {
        require(_isAdmin(msg.sender), "for team member only");
        _setPrice(token, price);
    }

    function _setPrice(address token, uint price) internal {
        ds.store(ds.key11a(TokenPricePath, token), bytes32(price));
    }
}
