package models

// VoteSubmitForm
//
//	@Description: 投票候选人的需要在前端展示的模型
type VoteSubmitForm struct {
	// 投票人学号
	Uid string `json:"uid"`
	// 投票结果
	Result []int `json:"result"`
}
