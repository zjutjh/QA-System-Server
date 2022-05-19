package submitService

import (
	"QA-System-Server/app/models"
	"QA-System-Server/config/database"
)

func CreateSubmit(data models.Submit) error {
	result := database.DB.Create(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func FetchSubmit(UID string) (*models.Submit, error) {
	var submit models.Submit
	result := database.DB.Where(models.Submit{
		UID: UID,
	}).Find(&submit)
	if result.Error != nil {
		return nil, result.Error
	}
	return &submit, nil
}
