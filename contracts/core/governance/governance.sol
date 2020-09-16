pragma solidity >=0.6.2;

import "node_modules/@openzeppelin/contracts/math/SafeMath.sol";

/**
 * @title sealers management smart contract
 */
contract Governance {
    using SafeMath for uint256;

    // zero address
    address constant ZERO_ADDRESS = address(0x0);

    enum Status {
        PENDING_ACTIVE,     // Sealer deposited enough NTFs into registration contract successfully.

        ACTIVE,             // Sealer send request to become a sealer
                            // and added into activation sealer set successfully

        PENDING_WITHDRAW,   // Sealer send request to exit from activation sealer set successfully.
                            // Sealer casted out of activation sealer set

        WITHDRAWN,          // Sealer already withdrawn their deposit NTFs successfully.
                            // They can only make withdrawal after withdrawal period.

        PENALIZED           //Sealer marked as penalized node (update by consensus or voting result via dapp)
                            //and cannot become active sealer and cannot withdraw balance neither.
    }

    struct Account {
        Status status;
        // ntf amount deposited
        uint256 balance;
        // delegated address to seal blocks
        address signer;
        // withdrawable block number after leaving
        uint256 unlockHeight;
    }

    // Consensus variables

    // index = 0
    // signers array
    address[] public signers;

    // index = 1
    // signer => coinbase (beneficiary address) map
    mapping(address => address) public signerCoinbase;

    // End of consensus variables

    // coinbase => NTF Account map
    mapping(address => Account) public account;

    // minimum of deposited NTF to join
    uint256 public stakeRequire;
    // minimum number of blocks signer has to wait from leaving block to withdraw the fund
    uint256 public stakeLockHeight;

    event Deposited(address _sealer, uint _amount);
    event Joined(address _sealer, address _signer);
    event Left(address _sealer, address _signer);
    event Withdrawn(address _sealer, uint256 _amount);
    event Banned(address _sealer);
    event Unbanned(address _sealer);

    /**
    * Check if address is a valid destination to transfer tokens to
    * - must not be zero address
    * - must not be the token address
    * - must not be the sender's address
    */
    modifier validSigner(address _signer) {
        require(signerCoinbase[_signer] == ZERO_ADDRESS, "coinbase already used");
        require(_signer != ZERO_ADDRESS, "signer zero");
        require(_signer != address(this), "same contract's address");
        require(_signer != msg.sender, "same sender's address");
        _;
    }

    modifier notBanned() {
        require(account[msg.sender].status != Status.PENALIZED, "banned ");
        _;
    }

    modifier joinable() {
        require(account[msg.sender].status != Status.ACTIVE, "already joined ");
        require(account[msg.sender].balance >= stakeRequire, "not enough ntf");
        _;
    }

    modifier leaveable() {
        require(account[msg.sender].status == Status.ACTIVE, "not joined ");
        _;
    }

    modifier withdrawable() {
        require(isWithdrawable(msg.sender), "unable to withdraw at the moment");
        _;
    }

    /**
    * contract initialize
    */
    constructor(uint256 _stakeRequire, uint256 _stakeLockHeight, address[] memory _signers) public {
        stakeRequire = _stakeRequire;
        stakeLockHeight = _stakeLockHeight;
        for (uint i = 0; i < _signers.length; i++) {
            signers.push(_signers[i]);
            signerCoinbase[_signers[i]] = _signers[i];
            account[_signers[i]].signer = _signers[i];
            account[_signers[i]].status = Status.ACTIVE;
        }
    }

    // Get ban status of a sealer's address
    function isBanned(address _address) public view returns(bool) {
        return (account[_address].status == Status.PENALIZED);
    }

    ////////////////////////////////

    function addSigner(address _signer) internal {
        signers.push(_signer);
    }

    function removeSigner(address _signer) internal {
        for (uint i = 0; i < signers.length; i++) {
            if (_signer == signers[i]) {
                signers[i] = signers[signers.length - 1];
                signers.pop();
                return;
            }
        }
    }

    function deposit() external payable {
        _deposit(msg.value);
    }

    receive() external payable {
        _deposit(msg.value);
    }

    function _deposit(uint256 _amount) internal {
        account[msg.sender].balance = (account[msg.sender].balance).add(_amount);
        emit Deposited(msg.sender, _amount);
    }

    /**
    * To allow deposited NTF participate joining in as sealer.
    * Participate already must deposit enough NTF via Deposit function.
    * It takes signer as parameter.
    * @param _signer Destination address
    */
    function join(address _signer) public notBanned joinable validSigner(_signer) returns (bool) {
        account[msg.sender].signer = _signer;
        account[msg.sender].status = Status.ACTIVE;
        signerCoinbase[_signer] = msg.sender;
        addSigner(_signer);
        emit Joined(msg.sender, _signer);
        return true;
    }

    /**
    * Request to exit out of activation sealer set
    */
    function leave() public notBanned leaveable returns (bool) {
        address _signer = account[msg.sender].signer;

        account[msg.sender].signer = ZERO_ADDRESS;
        account[msg.sender].status = Status.PENDING_WITHDRAW;
        account[msg.sender].unlockHeight = stakeLockHeight.add(block.number);
        delete signerCoinbase[_signer];
        removeSigner(_signer);
        emit Left(msg.sender, _signer);
        return true;
    }

    /**
    * To withdraw sealer’s NTF balance when they already exited and after withdrawal period.
    */
    function withdraw() public notBanned withdrawable returns (bool) {
        uint256 amount = account[msg.sender].balance;
        account[msg.sender].balance = 0;
        account[msg.sender].status = Status.WITHDRAWN;
        msg.sender.transfer(amount);
        emit Withdrawn(msg.sender, amount);
        return true;
    }

    function getStatusCode(Status _status) private pure returns(uint256){
        if (_status == Status.PENDING_ACTIVE) return 0;
        if (_status == Status.ACTIVE) return 1;
        if (_status == Status.PENDING_WITHDRAW) return 2;
        if (_status == Status.WITHDRAWN) return 3;
        return 127;
    }

    function getStatus(address _address) public view returns(uint256) {
        return getStatusCode(account[_address].status);
    }

    function getBalance(address _address) public view returns(uint256) {
        return account[_address].balance;
    }

    function getCoinbase(address _address) public view returns(address) {
        return account[_address].signer;
    }

    function getUnlockHeight(address _address) public view returns(uint256) {
        return account[_address].unlockHeight;
    }

    function isWithdrawable(address _address) public view returns(bool) {
        return
        (account[_address].status != Status.ACTIVE) &&
        (account[_address].status != Status.PENALIZED) &&
        (account[_address].unlockHeight < block.number);
    }
}
