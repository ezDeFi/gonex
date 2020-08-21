pragma solidity ^0.5.2;

import "openzeppelin-solidity/contracts/math/SafeMath.sol";
import "./ERC223.sol";

/*
    ...
*/

contract StableToken is ERC223{
    string public constant symbol = "NEWSD";
    string public constant name = "New Stable Dollar";
    uint public constant decimals = 18;

    constructor (
        address orderbook,      // mandatory
        address prefundAddress, // optional
        uint prefundAmount      // optional
    )
        public
    {
        if (prefundAmount > 0 ) {
            _mint(prefundAddress, prefundAmount * 10**decimals);
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
        transfer(owner(), haveAmount, data);
    }
}