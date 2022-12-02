package questionServices

import (
	"QA-System-Server/app/models"
	"QA-System-Server/config/database"
)

// GetQuestions
//
//	@Description: 获取题目的Service
//	@param id 试卷 ID
//	@return []models.Question 题目列表
//	@return error 可能的异常
func GetQuestions(id int) ([]models.Question, error) {
	var questions []models.Question

	result := database.DB.Where(
		models.Question{
			ExamId: id,
		}).Find(&questions)
	if result.Error != nil {
		return nil, result.Error
	}

	return questions, nil
}
