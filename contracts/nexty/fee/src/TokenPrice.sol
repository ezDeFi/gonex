// SPDX-License-Identifier: MIT
pragma solidity ^0.6.0;

import "./lib/ds.sol";
import "./interfaces/IERC20.sol";
import "./interfaces/IPayer.sol";

/**
 * Will be deployed by the consensus to 0x6789A
 */
contract TokenPrice {
    bytes32 public constant TokenPricePath = 'TokenPrice-'; // bytes11
    bytes32 public constant TeamMemberPath = 'TeamMember-'; // bytes11
    bytes32 public constant TRUE = 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF;

    constructor(
        address[] memory team,      // trusted team to set the token prices
        address[] memory tokens,    // intitial token addresses
        uint[] memory prices        // intitial token prices
    )
        public
    {
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

    function getPrice(address token)
        external
        view
        returns (uint price)
    {
        return uint(ds.load(ds.key11a(TokenPricePath, token)));
    }

    function setPrice(
        address token,
        uint price
    )
        external
    {
        require(isTeamMember(msg.sender), "for team member only");
        _setPrice(token, price);
    }

    function _setPrice(address token, uint price) internal {
        // TODO: keep a list
        ds.store(ds.key11a(TokenPricePath, token), bytes32(price));
    }
}
