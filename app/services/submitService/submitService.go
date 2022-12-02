package submitService

import (
	"QA-System-Server/app/models"
	"QA-System-Server/config/database"
)

// CreateSubmit
//
//	@Description: 创建提交信息的Service
//	@param data Submit数据
//	@return error 可能的异常
func CreateSubmit(data models.Submit) error {
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
//	@return *models.Submit Submit数据
//	@return error 可能的异常
func FetchSubmit(ID, UID string) (*models.Submit, error) {
	var submit models.Submit
	result := database.DB.Where(
		models.Submit{
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
func UpdateSubmit(data models.Submit) error {
	result := database.DB.Model(models.Submit{}).Where(
		&models.Submit{
			PaperID: data.PaperID,
			UID:     data.UID,
		}).Updates(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
