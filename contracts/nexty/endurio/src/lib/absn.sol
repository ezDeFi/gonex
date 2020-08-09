// SPDX-License-Identifier: UNLICENSED
pragma solidity >= 0.6.2;

import "./util.sol";
import "./map.sol";

library absn {
    using map for map.AddressBool;

    address constant ZERO_ADDRESS = address(0x0);

    using absn for Absorption;
    struct Absorption {
        uint deadline; // the last block the absorption is valid
        uint supply;   // StablizeToken supply at the activation block
        uint target;   // StablizeToken target supply for the active
        bool isHalted;
        bool isPreemptive;
    }

    function exists(Absorption storage self) internal view returns(bool) {
        return 0 < self.deadline;
    }

    function isExpired(Absorption storage self) internal view returns(bool) {
        return self.exists() && self.deadline < block.number;
    }

    function isAbsorbing(Absorption storage self, uint supply) internal view returns(bool) {
        return self.exists() &&
            !self.isHalted &&                               // not halted
            supply != self.target &&                        // target not reached &&
            util.inOrder(self.supply, supply, self.target); // not over-absorbed
    }

    using absn for Proposal;
    struct Proposal {
        // address of the proposer
        address maker;

        // amount of StablizeToken to absorb, positive for inflation, negative for deflation
        int amount;

        // NTY amount staked for the preemptive proposal
        uint stake;

        // slashingDivisor = stake * slashingRate / ZOOM / |amount|
        // toSlash = |deviation| * slashingDivisor
        uint slashingRate;

        // lockdown duration (in blocks from the activation)
        uint lockdownExpiration;

        // block number the proposal is proposed
        uint number;

        // voters map
        map.AddressBool votes;
    }

    function vote(Proposal storage self, bool up) internal {
        self.votes.set(msg.sender, up);
    }

    function exists(Proposal storage self) internal view returns (bool) {
        return (self.maker != ZERO_ADDRESS);
    }

    function clear(Proposal storage self) internal {
        self.votes.clear();
    }

    struct Preemptive {
        address maker;
        int amount;
        uint stake;
        uint slashingFactor;

        // block number the lockdown will end
        uint unlockNumber;
    }
    using absn for Preemptive;

    function exists(Preemptive storage self) internal view returns (bool) {
        return self.maker != ZERO_ADDRESS;
    }

    function isLocked(Preemptive storage self) internal view returns (bool) {
        return self.exists() && block.number < self.unlockNumber;
    }

    function unlockable(Preemptive storage self) internal view returns (bool) {
        return self.exists() && self.unlockNumber <= block.number;
    }
}
