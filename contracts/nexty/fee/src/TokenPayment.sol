// SPDX-License-Identifier: MIT
pragma solidity ^0.6.0;

import "./lib/ds.sol";
import "./interfaces/IConfig.sol";
import "./PayerCode.sol";

contract TokenPayment is PayerCode, IConfig {
    bytes32 public constant TRUE = 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF;
    bytes32 public constant TeamMemberPath  = 'TeamMember-';    // bytes11

    constructor(
        address[]   memory admins,  // trusted team to set the token prices
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

    function getAppTokenAndPrice(address app) external override view returns (address token, uint price) {
        token = _getAppToken(app);
        if (token != address(0x0)) {
            return (token, _getPrice(token));
        }
        return (app, _getPrice(app));
    }

    function setAppToken(address app, address token) external {
        require(_isAdmin(msg.sender), "for team member only");
        _setAppToken(app, token);
    }

    /**
     * this is for new dapp contract to call on construction/initilization
     */
    function setAppToken(address token) external {
        _setAppToken(msg.sender, token);
    }

    function getPrice(address token) external override view returns (uint price) {
        return _getPrice(token);
    }

    function setPrice(address token, uint price) external {
        require(_isAdmin(msg.sender), "for team member only");
        _setPrice(token, price);
    }
}
