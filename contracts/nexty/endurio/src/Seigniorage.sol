// SPDX-License-Identifier: UNLICENSED
pragma solidity >= 0.6.2;

import "./Preemptivable.sol";

/**
 * Seigniorage Shares Stablecoin
 */
contract Seigniorage is Preemptivable {
    constructor (
        uint absorptionDuration,
        uint absorptionExpiration,
        uint initialSlashingPace,
        uint initialLockdownExpiration
    )
        Preemptivable(
            absorptionDuration,
            absorptionExpiration,
            initialSlashingPace,
            initialLockdownExpiration
        )
        public
    {
    }
}