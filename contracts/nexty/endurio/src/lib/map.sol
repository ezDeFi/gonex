// SPDX-License-Identifier: UNLICENSED
pragma solidity >= 0.6.2;

import "./absn.sol";

library map {
    using absn for absn.Proposal;

    // Iterable map of (address => Proposal)
    // index = ordinals[a]-1
    // keys[index].maker = a
    using map for ProposalMap;
    struct ProposalMap {
        address[] keys;
        mapping (address => uint) ordinals;      // map the proposal's maker to (its index + 1)
        mapping (address => absn.Proposal) vals; // maker => proposal
    }

    function count(ProposalMap storage self) internal view returns (uint) {
        return self.keys.length;
    }

    function getKey(ProposalMap storage self, uint index) internal view returns (address) {
        return self.keys[index];
    }

    function get(ProposalMap storage self, uint index) internal view returns (absn.Proposal storage) {
        address key = self.getKey(index);
        return self.vals[key];
    }

    function get(ProposalMap storage self, address maker) internal view returns (absn.Proposal storage) {
        return self.vals[maker];
    }

    function has(ProposalMap storage self, address maker) internal view returns (bool) {
        return self.ordinals[maker] > 0;
    }

    function push(ProposalMap storage self, absn.Proposal memory proposal) internal returns (absn.Proposal storage) {
        address key = proposal.maker;
        uint ordinal = self.ordinals[key];
        require (ordinal == 0, "maker already has a proposal");
        self.vals[key] = proposal;
        self.keys.push(key);
        self.ordinals[key] = self.keys.length;
        return self.vals[key];
    }

    function remove(ProposalMap storage self, address maker) internal returns (bool success) {
        uint ordinal = self.ordinals[maker];
        require(ordinal > 0, "key not exist");
        self.remove(ordinal-1, maker);
        return true;
    }

    // index is the correct array index, which is (set.ordinals[item]-1)
    function remove(ProposalMap storage self, uint index) internal {
        address key = self.keys[index];
        require(key != address(0x0), "index not exist");
        self.remove(index, key);
    }

    // @require keys[index] == maker
    function remove(ProposalMap storage self, uint index, address maker) internal {
        delete self.ordinals[maker];
        delete self.vals[maker];

        if (self.keys.length-1 != index) {
            // swap with the last item in the keys and delete it
            self.keys[index] = self.keys[self.keys.length-1];
            self.ordinals[self.keys[index]] = index + 1;
        }
        // delete the last item from the array
        self.keys.pop();
    }

    function clear(ProposalMap storage self) internal {
        for (uint i = 0; i < self.keys.length; i++) {
            address key = self.keys[i];
            delete self.ordinals[key];
            delete self.vals[key];
        }
        delete self.keys;
    }

    ///////////////////////////////////////////////////////////////////////

    // Iterable map of (address => bool)
    // index = ordinals[a]-1
    // keys[index] = a
    using map for AddressBool;
    struct AddressBool {
        address[] keys;
        mapping (address => uint) ordinals; // map the voter's adress to (its index + 1)
        mapping (address => bool) vals;
    }

    function count(AddressBool storage self) internal view returns (uint) {
        return self.keys.length;
    }

    function getKey(AddressBool storage self, uint index) internal view returns (address) {
        return self.keys[index];
    }

    function get(AddressBool storage self, uint index) internal view returns (address, bool) {
        address key = self.getKey(index);
        return (key, self.vals[key]);
    }

    function get(AddressBool storage self, address key) internal view returns (bool) {
        return self.vals[key];
    }

    function has(AddressBool storage self, address key) internal view returns (bool) {
        return self.ordinals[key] > 0;
    }

    /**
     * @return true if new (key,val) is added, false if old key is map to a new value
     */
    function set(AddressBool storage self, address key, bool val) internal returns (bool) {
        self.vals[key] = val;
        uint ordinal = self.ordinals[key];
        if (ordinal == 0) {
            self.keys.push(key);
            self.ordinals[key] = self.keys.length;
            return true;
        }
        if (ordinal > self.keys.length || self.keys[ordinal-1] != key) {
            // storage inconsistency due to deleting proposal without clearing proposal.votes
            self.keys.push(key);
            self.ordinals[key] = self.keys.length;
            return true;
        }
        // key already has a proposal
        return false;
    }

    function remove(AddressBool storage self, address key) internal {
        uint ordinal = self.ordinals[key];
        require(ordinal > 0, "key not exist");
        self.remove(ordinal-1, key);
    }

    function remove(AddressBool storage self, uint index) internal {
        address key = self.keys[index];
        require(key != address(0x0), "index not exist");
        self.remove(index, key);
    }

    // @require keys[index] == key
    function remove(AddressBool storage self, uint index, address key) internal {
        delete self.ordinals[key];
        delete self.vals[key];

        if (self.keys.length-1 != index) {
            // swap with the last item in the keys and delete it
            self.keys[index] = self.keys[self.keys.length-1];
            self.ordinals[self.keys[index]] = index + 1;
        }
        // delete the last item from the array
        self.keys.pop();
    }

    function clear(AddressBool storage self) internal {
        for (uint i = 0; i < self.keys.length; i++) {
            address key = self.keys[i];
            delete self.ordinals[key];
            delete self.vals[key];
        }
        delete self.keys;
    }
}