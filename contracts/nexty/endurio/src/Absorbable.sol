pragma solidity ^0.5.2;

import "./lib/util.sol";
import "./lib/dex.sol";
import "./lib/absn.sol";
import "./Orderbook.sol";

/**
 * Mechanisms of absorption logic: active, passive and pre-emptive.
 */
contract Absorbable is Orderbook {
    using dex for dex.Book;
    using absn for absn.Absorption;
    using absn for absn.Preemptive;

    event Absorption(int256 amount, uint256 supply, bool emptive);
    event Stop();

    IToken VolatileToken;
    IToken StablizeToken; // spelling intented

    // constants
    uint EXPIRATION = 1 weeks / 2 seconds;
    int DURATION = int(EXPIRATION / 2);

    // last absorption
    absn.Absorption internal last;
    // lockdown tracks the active and current being lockdown pre-emptive absorption
    absn.Preemptive internal lockdown;

    constructor (
        uint absorptionDuration,
        uint absorptionExpiration
    )
        public
    {
        if (absorptionExpiration > 0) EXPIRATION = absorptionExpiration;
        DURATION = int(absorptionDuration > 0 ? absorptionDuration : absorptionExpiration / 2);
    }

    function registerTokens(
        address volatileToken,
        address stablizeToken
    )
        public
    {
        require(address(VolatileToken) == address(0), "VolatileToken already set");
        require(address(StablizeToken) == address(0), "StablizeToken already set");
        VolatileToken = IToken(volatileToken);
        StablizeToken = IToken(stablizeToken);
        super.registerTokens(volatileToken, stablizeToken);
        // trigger the first blank absorption
        uint supply = StablizeToken.totalSupply();
        triggerAbsorption(supply, supply, false, false);
    }

    function toString(address x) internal pure returns (string memory) {
        bytes memory b = new bytes(20);
        for (uint i = 0; i < 20; i++)
            b[i] = byte(uint8(uint(x) / (2**(8*(19 - i)))));
        return string(b);
    }

    modifier consensus() {
        require(msg.sender == address(0x0), "consensus only");
        _;
    }

    // for ethstat
    function getRemainToAbsorb() public view returns (bool, int) {
        if (last.target == 0) {
            return (false, 0);
        }
        return (true, util.sub(last.target, StablizeToken.totalSupply()));
    }

    // called by the consensus on each block
    // median price = target / StablizeToken.totalSupply()
    // zero target is fed for no median price available
    function onBlockInitialized(uint target) public consensus {
        if (last.isExpired()) {
            // absorption takes no longer than one duration
            stopAbsorption();
        }
        uint supply = StablizeToken.totalSupply();
        if (target > 0) { // absorption block
            if (shouldTriggerPassive()) {
                triggerAbsorption(target, supply, false, false);
            } else if (shouldTriggerActive(supply, target)) {
                triggerAbsorption(target, supply, true, false);
            }
        }
        if (last.isAbsorbing(supply)) {
            absorb();
        }
    }

    function calcNextAbsorption() internal view returns(int) {
        int total = util.sub(last.target, last.supply);
        int remain = util.sub(last.target, StablizeToken.totalSupply());
        if (!util.inOrder(0, remain, total)) {
            // target reached or passed
            return 0;
        }
        int amount = total / DURATION;
        if (last.isPreemptive) {
            amount /= 2;
        }
        if (!util.inOrder(0, amount, remain)) {
            // don't over absorb
            return remain;
        }
        return amount;
    }

    // shouldTriggerPassive returns whether a new passive absorption can be activated
    // passive condition: 1 duration without any active absorption or absorption never occurs
    function shouldTriggerPassive() internal view returns (bool) {
        return last.isExpired();
    }

    // shouldTriggerActive returns whether the new target is sufficient to trigger a new active absorption
    // make things simple by compare only the (target-supply) instead (target-supply)/supply
    function shouldTriggerActive(uint supply, uint target) internal view returns (bool) {
        if (target == supply) {
            return false;
        }
        if (last.target == last.supply) {
            return true;
        }
        // int a = util.sub(target, supply);
        // int b = util.sub(last.target, last.supply);
        // return a/b >= 2 || util.inOrder(-2, b/a, 0);
        if (target > supply) {
            uint a = target - supply;
            if (last.target > last.supply) {
                uint b = last.target - last.supply;
                return a / b >= 2;
            } else {
                uint b = last.supply - last.target;
                return b / a <= 2;
            }
        } else {
            uint a = supply - target;
            if (last.target < last.supply) {
                uint b = last.supply - last.target;
                return a / b >= 2;
            } else {
                uint b = last.target - last.supply;
                return b / a <= 2;
            }
        }
    }

    function triggerAbsorption(uint target, uint supply, bool emptive, bool isPreemptive) internal {
        last = absn.Absorption(block.number + EXPIRATION,
            supply,
            target,
            false,
            isPreemptive);
        int amount = util.sub(target, supply);
        emit Absorption(amount, supply, emptive);
    }

    function stopAbsorption() internal {
        delete last;
        emit Stop();
    }

    function absorb() internal {
        int amount = calcNextAbsorption();
        dex.Book storage book = books[amount > 0 ? Ask : Bid];
        bool useHaveAmount = book.haveToken == StablizeToken;

        (uint totalBMT, uint totalAMT) = book.absorb(useHaveAmount, util.abs(amount));

        if (!last.isPreemptive) {
            return;
        }

        // preemptive
        if (totalAMT == 0 || totalBMT == 0) {
            // no main absorb, no side absorb
            return;
        }

        // check the remain absorption
        uint supply = StablizeToken.totalSupply();
        if (!last.isAbsorbing(supply)) {
            return; // target reached, skip side-absorption
        }

        // book.absorbPreemptive(useHaveAmount, util.abs(amount), lockdown.maker);
        (uint haveAMT, uint wantAMT) = useHaveAmount ? (totalAMT, totalBMT) : (totalBMT, totalAMT);

        address initiator = lockdown.maker;
        if (haveAMT > book.haveToken.allowance(initiator, address(this)) ||
            haveAMT > book.haveToken.balanceOf(initiator)) {
            // not enough allowance for side absorption
            return;
        }

        book.haveToken.transferFrom(initiator, book.haveToken.dex(), haveAMT);
        book.haveToken.dexBurn(haveAMT);
        book.wantToken.dexMint(wantAMT);
        book.wantToken.transfer(initiator, wantAMT);
        // accumulate the side-absorb
        // totalAMT += useHaveAmount ? haveAMT : wantAMT;
        // totalBMT += useHaveAmount ? wantAMT : haveAMT;
        // emit SideFill(initiator, haveAMT, wantAMT);
    }

    /// MUST BE REMOVED ///
    function testAbsorb(
        int absorption,
        address initiator
    )
        external
    {
        bool isPreemptive = initiator != address(0x0);
        int amount = absorption;
        if (isPreemptive) {
            amount /= 2;
        }

        dex.Book storage book = books[amount > 0 ? Ask : Bid];
        bool useHaveAmount = book.haveToken == StablizeToken;

        (uint totalBMT, uint totalAMT) = book.absorb(useHaveAmount, util.abs(amount));

        if (!isPreemptive) {
            return;
        }

        // preemptive
        if (totalAMT == 0 || totalBMT == 0) {
            // no main absorb, no side absorb
            return;
        }

        // book.absorbPreemptive(useHaveAmount, util.abs(amount), lockdown.maker);
        (uint haveAMT, uint wantAMT) = useHaveAmount ? (totalAMT, totalBMT) : (totalBMT, totalAMT);

        if (haveAMT > book.haveToken.allowance(initiator, address(this)) ||
            haveAMT > book.haveToken.balanceOf(initiator)) {
            // not enough allowance for side absorption
            return;
        }

        book.haveToken.transferFrom(initiator, book.haveToken.dex(), haveAMT);
        book.haveToken.dexBurn(haveAMT);
        book.wantToken.dexMint(wantAMT);
        book.wantToken.transfer(initiator, wantAMT);
    }
}