// SPDX-License-Identifier: MIT
pragma solidity ^0.6.0;

import "./ds.sol";

library array {
    bytes32 public constant ArrayLenPostfix = 'Length'; // bytes16
    uint public constant NOT_FOUND = uint(-1);

    struct ArrayAddress {
        bytes32 id; // bytes15
    }
    using array for ArrayAddress;

    function newArrayAddress(bytes16 id)
        internal
        pure
        returns (ArrayAddress memory arr)
    {
        arr.id = id;
    }

    function get(
        ArrayAddress memory arr,
        uint index
    )
        internal
        view
        returns (address)
    {
        bytes32 key = key(arr.id, index);
        return ds.loadAddress(key);
    }

    function set(
        ArrayAddress memory arr,
        uint index,
        address adr
    )
        internal
    {
        bytes32 key = key(arr.id, index);
        ds.store(key, bytes32(bytes20(adr)));
    }

    function push(
        ArrayAddress memory arr,
        address adr
    )
        internal
        returns (uint)
    {
        bytes32 key = key(arr.id, arr.lenInc());
        ds.store(key, bytes32(bytes20(adr)));
    }

    // unordered item remove by index
    function remove(
        ArrayAddress memory arr,
        uint index
    )
        internal
    {
        uint decLen = arr.decLen(); // len-1
        require(index <= decLen, "ArrayAddress.remove: index out of bound");
        // only swap if it's not the last item
        if (index < decLen) {
            // arr[index] = arr[len-1]
            arr.set(index, arr.get(decLen));
        }
        arr.set(decLen, address(0x0));
    }

    // unordered item remove by index
    function remove(
        ArrayAddress memory arr,
        address adr
    )
        internal
    {
        uint index = arr.find(adr);
        require(index != NOT_FOUND, "ArrayAddress.remove: address not found");
        arr.remove(index);
    }

    function find(
        ArrayAddress memory arr,
        address adr
    )
        internal
        view
        returns (uint index)
    {
        uint len = arr.len();
        for (; index < len; ++index) {
            if (arr.get(index) == adr) {
                return index;
            }
        }
        return NOT_FOUND;
    }

    function len(
        ArrayAddress memory arr
    )
        internal
        view
        returns (uint)
    {
        bytes32 key = key(arr.id, uint(ArrayLenPostfix));
        return uint(ds.load(key));
    }

    // returns len++ (not ++len)
    function lenInc(
        ArrayAddress memory arr
    )
        internal
        returns (uint oldLen)
    {
        bytes32 key = key(arr.id, uint(ArrayLenPostfix));
        oldLen = uint(ds.load(key));
        ds.store(key, bytes32(oldLen+1));
    }

    // returns --len
    function decLen(
        ArrayAddress memory arr
    )
        internal
        returns (uint newLen)
    {
        bytes32 key = key(arr.id, uint(ArrayLenPostfix));
        newLen = uint(ds.load(key));
        ds.store(key, bytes32(--newLen));
    }

    function key(
        bytes32 id,
        uint index
    )
        internal
        pure
        returns (bytes32 ret)
    {
        // solhint-disable-next-line no-inline-assembly
        assembly {
            mstore(ret, id)             // 0..14
            mstore(add(ret,15), index)  // 15..30
            mstore8(add(ret,31), 0x0)   // 31
        }
    }
}