pragma solidity ^0.5.2;

import "./lib/util.sol";
import "./lib/map.sol";
import "./lib/absn.sol";
import "./Absorbable.sol";

/**
 * Pre-emptive absorption propsosal and voting logic.
 */
contract Preemptivable is Absorbable {
    using map for map.ProposalMap;
    using map for map.AddressBool;
    using absn for absn.Proposal;
    using absn for absn.Preemptive;

    event Revoke(address indexed maker);
    event Propose(
        address indexed maker,
        int256 amount,
        uint256 stake,
        uint256 lockdownExpiration,
        uint256 slashingRate
    );
    event Preemptive(
        address indexed maker,
        uint256 stake,
        uint256 lockdownExpiration,
        uint256 unlockNumber
    );
    event Slash(address indexed maker, uint256 amount);
    event Unlock(address indexed maker);

    address constant ZERO_ADDRESS = address(0x0);
    uint constant SLASHING_RATE_ZOOM = 1000;

    // adapting global default parameters, only used if proposal maker doesn't specify them
    uint internal globalLockdownExpiration = 2 weeks / 2 seconds;
    uint internal globalSlashingRate = SLASHING_RATE_ZOOM; // neuture rate

    // adapting global requirement
    uint internal globalSuccessRank = 0;
    uint internal globalSuccessStake = 0;

    // proposal params must not lower than 1/3 of global params
    uint constant PARAM_TOLERANCE = 3;

    // proposal must have atleast globalLockdownExpiration/4 block to be voted
    // note: using globalLockdownExpiration instead of proposal's value for safety
    uint constant MIN_VOTING_DURATION = 4;

    // map (maker => Proposal)
    map.ProposalMap internal proposals;

    // revoked proposals' votes to clear by consensus
    // since the clearing job is too expensive for transaction to perform.
    map.AddressBool[] votesToClear;

    constructor (
        uint absorptionPace,
        uint absorptionExpiration,
        uint initialSlashingPace,
        uint initialLockdownExpiration
    )
        Absorbable(
            absorptionPace,
            absorptionExpiration
        )
        public
    {
        if (initialLockdownExpiration > 0) {
            globalLockdownExpiration = initialLockdownExpiration;
        }
        if (initialSlashingPace > 0) {
            globalSlashingRate = initialSlashingPace;
        }
    }

    // Token transfer's fallback
    // bytes _data = uint[2] = (wantAmount, assistingID)
    // RULE : delegateCall never used
    //
    // buy/sell order is created by sending token to this address,
    // with extra data = (wantAmount, assistingID)
    function tokenFallback(
        address maker,  // actual tx sender
        uint value,     // amount of ERC223(msg.sender) received
        bytes calldata data)
        external
    {
        // if MNTY is received and data contains 4 params
        if (data.length == 32*4 && msg.sender == address(VolatileToken)) {
            // pre-emptive absorption proposal
            require(!proposals.has(maker), "already has a proposal");

            (   int amount,
                uint slashingRate,
                uint lockdownExpiration,
                bytes32 reserve // reserve params to distinguish proposal and trading request
            ) = abi.decode(data, (int, uint, uint, bytes32));

            // unused
            reserve = bytes32(0);

            propose(maker, amount, value, slashingRate, lockdownExpiration);
            return;
        }

        // not a pre-emptive proposal, fallback to Orderbook trader order
        bytes32 index;
        uint wantAmount;
        bytes32 assistingID;
        if (data.length == 32*3) {
            (index, wantAmount, assistingID) = abi.decode(data, (bytes32, uint, bytes32));
        } else {
            (index, wantAmount) = abi.decode(data, (bytes32, uint));
        }

        super.trade(maker, index, value, wantAmount, assistingID);
    }

    function onBlockInitialized(uint target) public consensus {
        // cleaning up
        for (uint i = 0; i < votesToClear.length; i++) {
            votesToClear[i].clear();
        }
        delete votesToClear;

        if (lockdown.unlockable()) {
            unlock();
        }
        checkAndTriggerPreemptive();

        if (target > 0) { // price block
            if (lockdown.isLocked() && last.isPreemptive) {
                uint supply = StablizeToken.totalSupply();
                int deviation = util.sub(target, supply);
                // halt the PeA if lockdown is violated
                last.isHalted = deviation != 0 && checkAndSlash(deviation);
            }
        }
        super.onBlockInitialized(target);
    }

    function unlock() internal {
        if (!lockdown.exists()) {
            return;
        }
        if (lockdown.stake > 0) {
            VolatileToken.transfer(lockdown.maker, lockdown.stake);
        }
        emit Unlock(lockdown.maker);
        delete lockdown;
    }

    /**
     * @dev slash the initiator whenever the price is moving in
     * opposition direction with the initiator's direction,
     * the initiator's deposited balance will be minus by slashed
     *
     * slashed = MIN(PeA.Stake, MAX(1, -Deviation/PeA.Amount / PeA.SlashingPace))
     *
     * @return true if the lockdown is violated and get slashed
     */
    function checkAndSlash(int deviation) internal returns (bool) {
        if (!util.inOrder(lockdown.amount, 0, deviation)) {
            // same direction, no slashing
            return false;
        }

        // lockdown violated
        uint toSlash = util.mulCap(util.abs(deviation), lockdown.slashingFactor);
        if (lockdown.stake < toSlash) {
            toSlash = lockdown.stake;
        }

        // this slashed NTY will be burnt by the consensus by calling setBalance
        lockdown.stake -= toSlash;
        VolatileToken.dexBurn(toSlash);

        emit Slash(lockdown.maker, toSlash);

        if (lockdown.stake == 0) {
            // there's nothing at stake anymore, clear the lockdown and its absorption
            stopAbsorption();
            unlock();
        }
        return true;
    }

    /**
     * propose allows Preemptive initiator to lock their MNTY in and
     * introduces new proposal.
     */
    function propose(
        address maker,
        int amount,
        uint stake,
        uint slashingRate,
        uint lockdownExpiration
    )
        internal
    {
        require(stake >= globalSuccessStake - globalSuccessStake / PARAM_TOLERANCE, "stake too low");
        require(amount != 0, "zero absorption");

        absn.Proposal memory proposal;

        if (slashingRate > 0) {
            require(
                slashingRate >=
                globalSlashingRate - globalSlashingRate / PARAM_TOLERANCE,
                "slashing rate param too low");
            proposal.slashingRate = slashingRate;
        } else {
            proposal.slashingRate = globalSlashingRate;
        }
        uint slashingFactor = util.mulCap(stake, proposal.slashingRate) / SLASHING_RATE_ZOOM / util.abs(amount);
        require(slashingFactor > 0, "slashing factor calculated to zero");

        if (lockdownExpiration > 0) {
            require(
                lockdownExpiration >=
                globalLockdownExpiration - globalLockdownExpiration / PARAM_TOLERANCE,
                "lockdown duration param too short");
            proposal.lockdownExpiration = lockdownExpiration;
        } else {
            proposal.lockdownExpiration = globalLockdownExpiration;
        }

        proposal.maker = maker;
        proposal.amount = amount;
        proposal.stake = stake;
        proposal.number = block.number;
        proposals.push(proposal);

        emit Propose(
            maker,
            amount,
            stake,
            proposal.slashingRate,
            proposal.lockdownExpiration
        );
    }

    function revoke(address maker) external {
        absn.Proposal storage p = proposals.get(maker);
        require(maker == p.maker, "only maker can revoke proposal");
        votesToClear.push(p.votes); // leave the job for consensus
        VolatileToken.transfer(p.maker, p.stake);
        proposals.remove(maker);
        emit Revoke(maker);
    }

    function vote(address maker, bool up) external {
        require(proposals.has(maker), "no such proposal");
        absn.Proposal storage proposal = proposals.get(maker);
        proposal.vote(up);
        // emit Vote(maker, up);
    }

    // check and trigger a new Preemptive when one is eligible
    // return the true if a new preemptive is activated
    function checkAndTriggerPreemptive() internal returns (bool) {
        if (lockdown.isLocked()) {
            // there's current active or lockdown absorption
            return false;
        }
        (address maker, uint rank) = winningProposal();
        if (maker == ZERO_ADDRESS) {
            // no eligible proposals
            return false;
        }
        absn.Proposal storage proposal = proposals.get(maker);
        adaptGlobalParams(proposal, rank);
        triggerPreemptive(proposal);
        return true;
    }

    // adapt the global params to the last winning preemptive
    function adaptGlobalParams(absn.Proposal storage proposal, uint rank) internal {
        globalSuccessStake = adaptParam(globalSuccessStake, proposal.stake);
        globalSlashingRate = adaptParam(globalSlashingRate, proposal.slashingRate);
        globalLockdownExpiration = adaptParam(globalLockdownExpiration, proposal.lockdownExpiration);
        globalSuccessRank = adaptParam(globalSuccessRank, rank);
    }

    function adaptParam(uint oldValue, uint newValue)
        internal
        pure
        returns(uint)
    {
        if (newValue == oldValue) {
            return oldValue;
        }
        uint value = util.avgCap(oldValue, newValue);
        if (newValue < oldValue || oldValue == 0) {
            return value;
        }
        uint change = value - oldValue; // safe
        if (change > oldValue / PARAM_TOLERANCE) {
            return oldValue + oldValue / PARAM_TOLERANCE; // safe
        }
        return value;
    }

    function getGlobalParams()
        external
        view
        returns(
            uint stake,
            uint slashingRate,
            uint lockdownExpiration,
            uint rank
        )
    {
        return (globalSuccessStake, globalSlashingRate, globalLockdownExpiration, globalSuccessRank);
    }

    // trigger an absorption from a maker's proposal
    function triggerPreemptive(absn.Proposal storage proposal) internal {
        proposal.votes.clear(); // clear the votes (consensus only)
        uint slashingFactor = util.mulCap(proposal.stake, proposal.slashingRate) / SLASHING_RATE_ZOOM / util.abs(proposal.amount);
        lockdown = absn.Preemptive(
            proposal.maker,
            proposal.amount,
            proposal.stake,
            slashingFactor,
            block.number + proposal.lockdownExpiration
        );
        proposals.remove(proposal.maker);
        uint supply = StablizeToken.totalSupply();
        uint target = util.add(supply, lockdown.amount);
        triggerAbsorption(target, supply, true, true);
        emit Preemptive(
            proposal.maker,
            proposal.stake,
            proposal.slashingRate,
            lockdown.unlockNumber
        );
    }

    // expensive calculation, only consensus can affort this
    function calcRank(absn.Proposal storage proposal) internal view returns (int) {
        int voteCount = countVote(proposal);
        if (voteCount <= 0) {
            return 0;
        }
        return util.mulCap(voteCount, int(proposal.stake));
    }

    // expensive calculation, only consensus can affort this
    function countVote(absn.Proposal storage proposal) internal view returns(int) {
        int voteCount = 0;
        for (uint i = 0; i < proposal.votes.count(); ++i) {
            (address voter, bool up) = proposal.votes.get(i);
            int weight = int(voter.balance + VolatileToken.balanceOf(voter));
            if (up) {
                voteCount += weight;
            } else {
                voteCount -= weight;
            }
        }
        return voteCount;
    }

    // expensive calculation, only consensus can affort this
    function winningProposal() internal view returns(address, uint) {
        int rankRequirement = int(globalSuccessRank - globalSuccessRank / PARAM_TOLERANCE);
        uint durationRequirement = globalLockdownExpiration / MIN_VOTING_DURATION;
        int bestRank = 0;
        address bestMaker = ZERO_ADDRESS;
        bool enoughVotingTime;
        for (uint i = 0; i < proposals.count(); ++i) {
            absn.Proposal storage proposal = proposals.get(i);
            int rank = calcRank(proposal);
            if (rank < rankRequirement) {
                // not good enough
                continue;
            }
            if (rank > bestRank) {
                bestRank = rank;
                bestMaker = proposal.maker;
            }
            // there's atleast 1 activable proposal
            if (!enoughVotingTime && durationRequirement <= block.number - proposal.number) {
                enoughVotingTime = true;
            }
        }
        if (!enoughVotingTime) {
            return (ZERO_ADDRESS, 0);
        }
        return (bestMaker, uint(bestRank));
    }

    function totalVote(address maker) public view returns(int) {
        absn.Proposal storage p = proposals.get(maker);
        return countVote(p);
    }

    function getProposalCount() public view returns(uint) {
        return proposals.count();
    }

    function getProposal(uint idx) public view
        returns (
            address maker,
            uint stake,
            int amount,
            uint slashingRate,
            uint lockdownExpiration,
            uint number
        )
    {
        absn.Proposal storage p = proposals.get(idx);
        return (p.maker, p.stake, p.amount, p.slashingRate, p.lockdownExpiration, p.number);
    }

    function getLockdown() public view
        returns(
            address maker,
            uint stake,
            int amount,
            uint slashingFactor,
            uint unlockNumber
        )
    {
        return (lockdown.maker, lockdown.stake, lockdown.amount, lockdown.slashingFactor, lockdown.unlockNumber);
    }
}