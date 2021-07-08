pragma solidity >=0.5.16;

// compiled with solc 0.5.17

contract Transfer {
    address public owner;

    event Sent(address receiver, uint256 amount);

    constructor() public {
        owner = msg.sender;
    }

    function send(address payable receiver, uint256 amount) public {
        require(msg.sender == owner, "sender is not owner");

        receiver.transfer(amount);
        emit Sent(receiver, amount);
    }

    function() external payable {}
}
