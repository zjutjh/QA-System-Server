package examNameServices

import (
	"QA-System-Server/app/models"
	"QA-System-Server/config/database"
)

// GetExamName
//
//	@Description: 获取试卷名称的Service
//	@param id 试卷ID
//	@param time 获取时间
//	@return *string 试卷名称
//	@return error 可能的异常
func GetExamName(id string) ([]models.ExamNameMap, error) {
	var examNameMaps []models.ExamNameMap

	result := database.DB.Where(
		models.ExamNameMap{
			ExamID: id,
		}).Find(&examNameMaps)
	if result.Error != nil {
		return nil, result.Error
	}
	return examNameMaps, nil
}
