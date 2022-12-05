package nameServices

import (
	"QA-System-Server/app/models"
	"QA-System-Server/config/database"
)

// GetName
//
//	@Description: 获取试卷或投票名称的Service
//	@param id 试卷或投票ID
//	@param time 获取时间
//	@return *string 试卷或投票名称
//	@return error 可能的异常
func GetName(id int) ([]models.NameMap, error) {
	var nameMaps []models.NameMap

	result := database.DB.Where(
		models.NameMap{
			ID: id,
		}).Find(&nameMaps)
	if result.Error != nil {
		return nil, result.Error
	}
	return nameMaps, nil
}
