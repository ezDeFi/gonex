// SPDX-License-Identifier: UNLICENSED
pragma solidity >= 0.6.2;

import "@openzeppelin/contracts/math/SafeMath.sol";
import "./ERC223.sol";

/*
    ...
*/

contract StableToken is ERC223 {
    constructor (
        address orderbook,      // mandatory
        address prefundAddress, // optional
        uint prefundAmount      // optional
    )
        public
        ERC20("New Stable Dollar", "NEWSD")
    {
        if (prefundAmount > 0 ) {
            _mint(prefundAddress, prefundAmount * 10**18);
        }
        initialize(orderbook);
    }

    // order USD -> MNTY
    function trade(
        bytes32 index,
        uint haveAmount,
        uint wantAmount,
        bytes32 assistingID
    )
        public
        payable
    {
        bytes memory data = abi.encode(index, wantAmount, assistingID);
        transfer(dex, haveAmount, data);
    }
}