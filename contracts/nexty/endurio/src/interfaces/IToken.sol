pragma solidity ^0.5.2;

interface IToken {
    function dex() external returns (address);
    function dexMint(uint _amount) external;
    function dexBurn(uint _amount) external;
    function transferToDex(address _from, uint _amount) external;

    // IERC20
    function transfer(address to, uint256 value) external returns (bool);
    function approve(address spender, uint256 value) external returns (bool);
    function transferFrom(address from, address to, uint256 value) external returns (bool);
    function totalSupply() external view returns (uint256);
    function balanceOf(address who) external view returns (uint256);
    function allowance(address owner, address spender) external view returns (uint256);
    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(address indexed owner, address indexed spender, uint256 value);
}