// SPDX-License-Identifier: MIT
pragma solidity ^0.8.1;
import "./Token.sol";
import "./ReportVoting.sol";
import "./safeTrans.sol";
import "./IERC20.sol";

contract Nuxt{
    uint public UserNum;

    struct Personal {
        string personalName; //用户名
        uint personalAmount; // 用户代币余额
        uint NeutralizedAmount; //中和量
        uint registrationTime;  //注册时间
        uint NumberOfReported;  //被举报次数
    }

    mapping(address => Personal) public newPersonal;
    mapping(address => bool) initialization;

    modifier OnlySignIn {
        require(initialization[msg.sender]);
        _;
    }

    event initializationEVT(address ip,uint time);

    function Initialization(string memory _name) public {
        require(!initialization[msg.sender]);
        require(keccak256(abi.encodePacked(_name))  != keccak256(abi.encodePacked("")),"name is null");
        Personal storage p = newPersonal[msg.sender];
        p.personalName = _name;
        p.personalAmount = 0;
        p.NeutralizedAmount = 0;
        p.registrationTime = block.timestamp;
        p.NumberOfReported = 0;
        initialization[msg.sender] = true;
        UserNum++;
        emit initializationEVT(msg.sender,block.timestamp);
    }

    function initializationSTU() public view returns(bool){
        return initialization[msg.sender];
    }
}

contract PartTake is Token,Nuxt,ReportVoting{
    using safeTrans for uint256;

    uint public NUM = 0;

    uint public currentYear;

    uint neutralizeTotals;

    uint neutralizePersonals;

    struct Neutralize {
        address publisher;  //发布者
        uint neutralizeTotal;  //需要中和的总碳排
        uint neutralizePersonals; //每个人的中和量
        uint peopleNUM;  //录入时总人数
        uint recordTime;  //录入时间
        bool UpToStandard;  //是否达标
        uint CurrentQuota; //目前额度
    }

    struct Message {
        address ip;  //用户地址
        string fun; //绿色方式
        uint time;  //上传时间
        bool stu; //状态
    }

    mapping(uint => Neutralize) public NeutralizeYear; //year -> Neutralize
    mapping(uint => Message) public messageID;

    event takePart(address ip,uint id,uint time);

    function ReleaseNeutralize(uint _year,uint _neutralizeTotal) public{
        require(_year > currentYear,"The year has been set!");
        Neutralize storage n = NeutralizeYear[_year];
        n.publisher = msg.sender;
        n.neutralizeTotal = _neutralizeTotal;
        n.neutralizePersonals = _neutralizeTotal / UserNum;
        n.peopleNUM = UserNum;
        n.recordTime =  block.timestamp;
        n.UpToStandard = false;
        n.CurrentQuota = 0;
        currentYear = _year;
        neutralizeTotals = _neutralizeTotal;
        neutralizePersonals = _neutralizeTotal / UserNum;
    }

    function TakeAPart(string memory _fun,uint _NeutralizedAmount,uint _personalAmount) OnlySignIn public{
        require(keccak256(abi.encodePacked(_fun)) != keccak256(abi.encodePacked("")),"_fun is null");
        Message storage m = messageID[NUM];
        m.ip = msg.sender;
        m.fun = _fun;
        m.time = block.timestamp;
        m.stu = true;
        NeutralizeYear[currentYear].CurrentQuota += _NeutralizedAmount;
        NUM++;
        newPersonal[msg.sender].NeutralizedAmount = newPersonal[msg.sender].NeutralizedAmount.add(_NeutralizedAmount);
        newPersonal[msg.sender].personalAmount = newPersonal[msg.sender].personalAmount.add(_personalAmount);
        IERC20(address(this)).transfes(msg.sender,_personalAmount);
        emit takePart(msg.sender,NUM - 1,block.timestamp);
    }

    function Report(address _ip,string memory _reason) OnlySignIn public{
        require(_ip != address(0),"address is null");
        createVote(_ip,_reason);
    }

    function finish(uint pId,uint _reward,uint MASSID) OnlySignIn public {
        if (vote[pId].limitTime == 0)
            revert("proposal not exist");

        uint currentTime = block.timestamp;

        if (vote[pId].limitTime > currentTime)
            revert("The vote is not over yet");

        if(vote[pId].voteCount >= UserNum / 2){
            if(newPersonal[vote[pId].ip].personalAmount >= (_reward * 3 / 2)){
                approve(address(this),(_reward * 3 / 2));
                IERC20(address(this)).transferFrom(msg.sender,address(this),(_reward * 3 / 2));
                newPersonal[vote[pId].ip].personalAmount = newPersonal[vote[pId].ip].personalAmount.sub(_reward * 3 / 2);
                messageID[MASSID].stu = false;
            }
        }

    }

    function SubERC(address ip,uint amount)public {
        newPersonal[ip].personalAmount = newPersonal[ip].personalAmount.sub(amount);
    }

    function AddERC(address ip,uint amount)public {
        newPersonal[ip].personalAmount = newPersonal[ip].personalAmount.add(amount);
    }

}
