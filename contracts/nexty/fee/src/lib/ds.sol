// SPDX-License-Identifier: MIT
pragma solidity ^0.6.0;

/**
 * Key hashing or not, storage should be used like 'storage', not like in-memory variables.
 */
library ds {
    function load(bytes32 key)
        internal
        view
        returns (bytes32 value)
    {
        assembly {
            value := sload(key)
        }
    }

    function store(bytes32 key, bytes32 value)
        internal
    {
        assembly {
            sstore(key, value)
        }
    }

    function loadAddress(bytes32 key)
        internal
        view
        returns (address value)
    {
        assembly {
            value := sload(key)
        }
    }

    function storeAddress(bytes32 key, address value) internal {
        assembly {
            sstore(key, value)
        }
    }

    /**
     * @dev unsafe(path > bytes11)
     */
    function key11a(bytes32 path, address adr)
        internal
        pure
        returns (bytes32 key)
    {
        return path | bytes32((uint(adr) << 8));
        // assembly {
        //     mstore(key, path)           // 0..10
        //     mstore(add(key,11), adr)    // 11..30
        //     mstore8(add(key,31), 0x0)   // 31
        // }
    }

    /**
     * @dev unsafe(path > bytes11)
     */
    function keya11(address adr, bytes32 path)
        internal
        pure
        returns (bytes32 key)
    {
        return bytes32(bytes20(adr)) | path >> 160;
        // assembly {
        //     mstore(key, adr)            // 0..19
        //     mstore(add(key,20), path)   // 20..30
        //     mstore8(add(key,31), 0x0)   // 31
        // }
    }
}