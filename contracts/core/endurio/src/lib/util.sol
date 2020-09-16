// SPDX-License-Identifier: UNLICENSED
pragma solidity >= 0.6.2;

library util {
    uint256 constant MaxUint256 = 0xFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF;
    int256 constant MaxInt256   = 0x7FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF;
    int256 constant MinInt256   = MaxInt256 + 1;

    function abs(int a) internal pure returns (uint) {
        return uint(a > 0 ? a : -a);
    }

    // subtract 2 uints and convert result to int
    function sub(uint a, uint b) internal pure returns(int) {
        // require(|a-b| < 2**128)
        return a > b ? int(a - b) : -int(b - a);
    }

    // TODO: apply SafeMath
    function add(uint a, int b) internal pure returns(uint) {
        if (b < 0) {
            return a - uint(-b);
        }
        return a + uint(b);
    }

    function inOrder(uint a, uint b, uint c) internal pure returns (bool) {
        return (a <= b && b <= c) || (a >= b && b >= c);
    }

    function inStrictOrder(uint a, uint b, uint c) internal pure returns (bool) {
        return (a < b && b < c) || (a > b && b > c);
    }

    function inOrder(int a, int b, int c) internal pure returns (bool) {
        return (a <= b && b <= c) || (a >= b && b >= c);
    }

    function inStrictOrder(int a, int b, int c) internal pure returns (bool) {
        return (a < b && b < c) || (a > b && b > c);
    }

    // capped multiply
    // if the calculation is overflown, return the max or min value of the type
    function mulCap(int a, int b) internal pure returns (int) {
        // Gas optimization: this is cheaper than requiring 'a' not being zero, but the
        // benefit is lost if 'b' is also tested.
        // See: https://github.com/OpenZeppelin/openzeppelin-solidity/pull/522
        if (a == 0) {
            return 0;
        }

        int c = a * b;
        if (c / a == b) {
            return c;
        }

        if (inStrictOrder(a, 0, b)) {
            // negative overflown
            return MinInt256;
        }
        // positive overflown
        return MaxInt256;
    }

    // unsigned capped multiply
    // if the calculation is overflown, return the max value of uint256
    function mulCap(uint a, uint b) internal pure returns (uint) {
        // Gas optimization: this is cheaper than requiring 'a' not being zero, but the
        // benefit is lost if 'b' is also tested.
        // See: https://github.com/OpenZeppelin/openzeppelin-solidity/pull/522
        if (a == 0) {
            return 0;
        }

        uint c = a * b;
        if (c / a == b) {
            return c;
        }

        // overflown
        return MaxUint256;
    }

    // capped average
    // if the calculation is overflown, return the max value of uint
    function avgCap(uint a, uint b) internal pure returns (uint) {
        uint c = a + b;
        if (c >= a) {
            return c >> 1;
        }
        // (a+b) is overflown
        uint d = (a >> 1) + (b >> 1);
        if (d >= a) {
            return d;
        }
        // (a/2+b/2) is also overflown
        return MaxUint256;
    }
}