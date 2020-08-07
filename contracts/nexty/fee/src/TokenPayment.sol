// SPDX-License-Identifier: MIT
pragma solidity ^0.6.0;

import "./lib/ds.sol";
import "./interfaces/IConfig.sol";
import "./PayerCode.sol";

contract TokenPayment is PayerCode, IConfig {
    bytes32 public constant TRUE = 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF;
    bytes32 public constant TeamMemberPath = 'TeamMember-'; // bytes11
    bytes32 public constant TokenForPath = 'TokenFor---'; // bytes11: token for a specific address

    constructor(
        address[]   memory team,    // trusted team to set the token prices
        address[]   memory tokens,  // intitial token addresses
        uint[]      memory prices   // intitial token prices
    ) public {
        for (uint i = 0; i < team.length; i++) {
            addTeamMember(team[i]);
        }
        for (uint i = 0; i < tokens.length; i++) {
            _setPrice(tokens[i], prices[i]);
        }
    }

    function addTeamMember(address member) internal {
        ds.store(ds.key11a(TeamMemberPath, member), TRUE);
    }

    function isTeamMember(address member) internal view returns (bool) {
        return ds.load(ds.key11a(TeamMemberPath, member)) != 0;
    }

    function setTokenFor(address to, address token) external {
        require(isTeamMember(msg.sender), "for team member only");
        ds.storeAddress(ds.key11a(TokenForPath, to), token);
    }

    function _getTokenFor(address to) internal view returns (address) {
        return ds.loadAddress(ds.key11a(TokenForPath, to));
    }

    function getTokenFor(address to) external override view returns (address token) {
        token = _getTokenFor(to);
        if (token != address(0x0)) {
            return token;
        }
        return to;
    }

    function getPrice(address token) external override view returns (uint price) {
        return uint(ds.load(ds.key11a(TokenPricePath, token)));
    }

    function setPrice(address token, uint price) external {
        require(isTeamMember(msg.sender), "for team member only");
        _setPrice(token, price);
    }

    function _setPrice(address token, uint price) internal {
        ds.store(ds.key11a(TokenPricePath, token), bytes32(price));
    }
}
