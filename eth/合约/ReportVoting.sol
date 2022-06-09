// SPDX-License-Identifier: MIT
pragma solidity ^0.8.1;

contract ReportVoting {

	uint public RVNUM;

    struct Voter {
        uint voteTimeStamp; //投票时的区块时间
        bool initialized;   //判断是否投过票的标志
    }

    //提案内容
    struct Vote {
		address ip;  //被举报人的地址
        string vName;        //举报理由
		address chairperson;    //举报人
        uint voteCount;      //附议人数
        uint limitTime;      //附议限制时间
        mapping(address => Voter) voters;
    }

    //所有提案列表
    mapping(uint => Vote) public vote;

	//附议事件
	event VoteEvt(string indexed eventType, address _voter, uint timestamp);

	//提案事件
	event VOTE(address ip, uint _proposalId, uint _limitTime);

	//创建新提案
	function createVote(address ip, string memory _pName) public{
	  //新提案
	  	RVNUM++;
        Vote storage v = vote[RVNUM];
		v.ip = ip;
        v.vName = _pName;
    	v.chairperson = msg.sender;
    	v.limitTime = 604800 + block.timestamp;
        v.voteCount = 0;

      emit VOTE(msg.sender, RVNUM, 604800);
	}

	//进行附议
	function Voting(uint pId) public {
	  //提案是否存在
	  if (vote[pId].limitTime == 0)
	    revert("proposal not exist");

	  uint currentTime = block.timestamp;

	  //是否已超过提案时限
	  if (vote[pId].limitTime < currentTime)
	    revert("exceed voting time");

	  //是否已经投过票
	  if (vote[pId].voters[msg.sender].initialized == true)
	   revert("already vote");

	  //新投票信息
	  Voter memory voter = Voter({
	     voteTimeStamp: block.timestamp,
	     initialized: true
	  });

	  //记录投票信息
	  vote[pId].voters[msg.sender] = voter;
	  vote[pId].voteCount++;

	  emit VoteEvt("vote", msg.sender, block.timestamp);
	}

	//查询是否附议
	function VotingT_F(uint pId, address voterAddr) public view returns (uint){
	  //提案是否存在
	  if (vote[pId].limitTime == 0)
	    revert("proposal not exist");

	  //返回投票时间
	  return vote[pId].voters[voterAddr].voteTimeStamp;
	}
}
