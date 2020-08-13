// SPDX-License-Identifier: MIT
pragma solidity ^0.6.0;

import "./IERC20.sol";

interface IConfig {
    function getPrice(address token) external view returns (uint price);
    function getAppTokenAndPrice(address app) external view returns (address token, uint price);
}
