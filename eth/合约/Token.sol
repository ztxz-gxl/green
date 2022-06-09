// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;



// import "https://github.com/OpenZeppelin/openzeppelin-contracts/blob/master/contracts/token/ERC20/ERC20.sol";

import "./ERC20.sol";

contract Token is ERC20 {
    constructor() ERC20("Green","GTH") {

        _mint(msg.sender, 10000000000000000 * 10**uint(decimals()));
    }
}
