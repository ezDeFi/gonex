pragma solidity ^0.4.25;

contract Nexty {
    mapping(address => address) signers;
    address[] sealers;
    
    constructor() public {
        CODE_HERE
    }
    
    function join() public {
        sealers.push(msg.sender);
    }
    
    function leave() public {
        for (uint i = 0; i < sealers.length; i++) {
            if (msg.sender == sealers[i]) {
                sealers[i] = sealers[sealers.length - 1];
                delete sealers[sealers.length - 1];
                sealers.length--;
                return;
            }
        }
    }
}