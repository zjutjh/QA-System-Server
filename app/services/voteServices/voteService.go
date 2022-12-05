package voteServices

import (
	"QA-System-Server/app/models"
	"QA-System-Server/config/database"
)

// GetVoteList
//
//	@Description: 获取投票的Service
//	@param id 投票ID
//	@return []models.Vote 投票列表
//	@return error 可能的异常
func GetVoteList(id int) ([]models.Vote, error) {
	var votes []models.Vote

	result := database.DB.Where(
		models.Vote{
			VoteId: id,
		}).Find(&votes)
	if result.Error != nil {
		return nil, result.Error
	}

	return votes, nil
}
