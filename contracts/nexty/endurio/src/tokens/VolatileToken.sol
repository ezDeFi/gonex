// SPDX-License-Identifier: UNLICENSED
pragma solidity >= 0.6.2;

import "@openzeppelin/contracts/math/SafeMath.sol";
import "./ERC223.sol";

/*
    . Exchanged with NTY with rate 1 MegaNTY = 1000000 NTY
    . Mint. / burn. able(free) by owner = orderbook contract
*/

contract VolatileToken is ERC223 {
    // track the amount of token has been inflated, negative for deflation
    // only use for report, so ignore any overflow
    int inflated;

    constructor (
        address orderbook,      // mandatory
        address prefundAddress, // optional
        uint prefundAmount      // optional
    )
        public
        ERC20("Wrapped NTY", "WNTY")
    {
        if (prefundAmount > 0 ) {
            _mint(prefundAddress, prefundAmount * 10**18); // TODO: pass wei instead of coin
        }
        initialize(orderbook);
    }

    // override ERC223.dexMint
    function dexMint(uint value) public override {
        super.dexMint(value);
        inflated += int(value);
    }

    // override ERC223.dexBurn
    function dexBurn(uint value) public override {
        super.dexBurn(value);
        inflated -= int(value);
    }

    function totalInflated()
        external
        view
        returns(int)
    {
        return inflated;
    }

    // deposit (MNTY <- NTY)
    function deposit()
        external
        payable
        returns(bool)
    {
        depositTo(msg.sender);
    }

    // withdraw (MNTY -> NTY)
    function withdraw(uint _amount)
        external
        returns(bool)
    {
        withdrawTo(_amount, msg.sender);
    }

    // withdrawTo (MNTY -> NTY -> address)
    function withdrawTo(uint _amount, address payable _to)
        public
        returns(bool)
    {
        address _sender = msg.sender;
        _burn(_sender, _amount);

        /************************************************************************/
        /* concensus garantures, this contract always got enough NTY to withdraw */
        /************************************************************************/

        _to.transfer(_amount);
    }

    // depositTo (addresss <- MNTY <- NTY)
    function depositTo(
        address _to
    )
        public
        payable
        returns(bool)
    {
        uint _amount = msg.value;
        _mint(_to, _amount);
        return true;
    }

    // deposit and order (NTY -> MNTY -> USD)
    function depositAndTrade(
        bytes32 index,
        uint haveAmount,
        uint wantAmount,
        bytes32 assistingID
    )
        public
        payable
    {
        depositTo(msg.sender);
        trade(index, haveAmount, wantAmount, assistingID);
    }

    // create selling order (MNTY -> USD)
    // with verbose data = (wantAmount, assistingID)
    function trade(
        bytes32 index,
        uint haveAmount,
        uint wantAmount,
        bytes32 assistingID
    )
        public
    {
        bytes memory data = abi.encode(index, wantAmount, assistingID);
        transfer(dex, haveAmount, data);
    }

    // deposit and propose()
    function depositAndPropose(
        int amount,             // absorption amount of StablizeToken
        uint stake,             // staked amount of VolatileToken
        uint slashingRate,
        uint lockdownExpiration
    )
        public
        payable
    {
        depositTo(msg.sender);
        propose(amount, stake, slashingRate, lockdownExpiration);
    }

    // propose a new pre-emptive absorption
    // with verbose data = (amount, slashingRate, lockdownExpiration);
    function propose(
        int amount,             // absorption amount of StablizeToken
        uint stake,             // staked amount of VolatileToken
        uint slashingRate,
        uint lockdownExpiration
    )
        public
    {
        bytes memory data = abi.encode(amount, slashingRate, lockdownExpiration, bytes32(0));
        transfer(dex, stake, data);
    }
}