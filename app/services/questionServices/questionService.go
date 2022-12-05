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

// CreateSubmit
//
//	@Description: 创建提交信息的Service
//	@param data Submit数据
//	@return error 可能的异常
func CreateSubmit(data models.ExamSubmit) error {
	result := database.DB.Create(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// FetchSubmit
//
//	@Description: 获取提交信息的Service
//	@param ID 试卷ID
//	@param UID 答题人学号
//	@return *models.ExamSubmit Submit数据
//	@return error 可能的异常
func FetchSubmit(ID, UID string) (*models.ExamSubmit, error) {
	var submit models.ExamSubmit
	result := database.DB.Where(
		models.ExamSubmit{
			UID:     UID,
			PaperID: ID,
		}).Find(&submit)
	if result.Error != nil {
		return nil, result.Error
	}
	return &submit, nil
}

// UpdateSubmit
//
//	@Description: 更新提交信息的Service
//	@param data 更新后的Submit数据
//	@return error 可能的异常
func UpdateSubmit(data models.ExamSubmit) error {
	result := database.DB.Model(models.ExamSubmit{}).Where(
		&models.ExamSubmit{
			PaperID: data.PaperID,
			UID:     data.UID,
		}).Updates(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
