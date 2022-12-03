package models

// Vote
//
//	@Description: 投票候选人的模型
type Vote struct {
	// 候选人ID
	CandidateId int
	// 投票ID
	VoteId int
	// 描述
	Describe string
	// 得票数
	Votes int
}
