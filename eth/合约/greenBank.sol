// SPDX-License-Identifier: MIT
pragma solidity ^0.8.1;
import "./partTake.sol";
import "./IERC20.sol"; 
import "./safeTrans.sol";

contract GreenBank is PartTake {
    using safeTrans for uint256;

    uint BankBalance;

    uint public BankAdminOperationNum;

    struct BankVoucher {
        address participant; //存入者
        uint partakeAmount; //存入金额
        uint deadline; //截止时间
        uint interest; //利息
        bool takeOut; //是否取出
    }

    struct BankAdmin{
        address ip;//操作地址
        uint operationAmount;//操作数额
        string operation;//存入或者取出操作
        uint operationTime; //操作时间
    }

    mapping(bytes32 => BankVoucher) public BankVoucherHash; //hash -> BankVoucher
    mapping(uint => BankAdmin) public BankAdminList;//BankAdminOperationNum -> BankAdmin

    event ADDBANK(address ip,bytes32 hash,bytes32 time);

    function Exchange(uint _amount) OnlySignIn public {
        require(_amount >= newPersonal[msg.sender].personalAmount && _amount >= 100,"Sorry, your credit is running low");
        newPersonal[msg.sender].personalAmount = newPersonal[msg.sender].personalAmount.sub(_amount);
        approve(address(this),_amount);
        IERC20(address(this)).transferFrom(msg.sender,address(this),_amount);
        payable(address(this)).transfer(_amount);
    }

    function Bank(uint _amount,uint _year,uint _interest) OnlySignIn public{
        require(_amount >= 100,"Sorry, your credit is running low");
        uint time = block.timestamp;
        bytes32 h = keccak256(abi.encodePacked(time));
        bytes32 hashs = keccak256(abi.encodePacked(msg.sender,h));
        BankVoucher storage b = BankVoucherHash[hashs];
        b.participant = msg.sender;
        b.partakeAmount = _amount;
        b.deadline = time + _year * 31536000;
        b.interest = _amount * _interest * 1 / 10000;
        b.takeOut = false;
        newPersonal[msg.sender].personalAmount = newPersonal[msg.sender].personalAmount.sub(_amount);
        BankBalance = BankBalance.add(_amount);
        approve(address(this),_amount);
        IERC20(address(this)).transferFrom(msg.sender,address(this),_amount);
        emit ADDBANK(msg.sender,hashs,h);
    }

    function TakeOutBank(uint _amount,string memory _operation) public {
        require(_amount <= BankBalance ,"the amount cannot be greater than the balance");
        BankAdmin storage b = BankAdminList[BankAdminOperationNum];
        b.ip = msg.sender;
        b.operationAmount = _amount;
        b.operation = _operation;
        b.operationTime = block.timestamp;
        BankAdminOperationNum ++;
        BankBalance = BankBalance.sub(_amount);
        IERC20(address(this)).transfes(msg.sender,_amount);
    }

    function Deposit(uint _amount,string memory _operation) public {
        BankAdmin storage b = BankAdminList[BankAdminOperationNum];
        b.ip = msg.sender;
        b.operationAmount = _amount;
        b.operation = _operation;
        b.operationTime = block.timestamp;
        BankAdminOperationNum ++;
        BankBalance = BankBalance.add(_amount);
        approve(address(this),_amount);
        IERC20(address(this)).transferFrom(msg.sender,address(this),_amount);
    }

    function GetInterest(bytes32 _hash) OnlySignIn public {
        require(BankVoucherHash[_hash].participant == msg.sender,"hash is abnormal");
        require(!BankVoucherHash[_hash].takeOut,"Reward removed");
        BankVoucherHash[_hash].takeOut = true;
        BankBalance = BankBalance.sub(BankVoucherHash[_hash].partakeAmount);
        newPersonal[msg.sender].personalAmount = newPersonal[msg.sender].personalAmount.add(BankVoucherHash[_hash].partakeAmount + BankVoucherHash[_hash].interest);
        IERC20(address(this)).transfes(msg.sender,BankVoucherHash[_hash].partakeAmount + BankVoucherHash[_hash].interest);
    }
}
