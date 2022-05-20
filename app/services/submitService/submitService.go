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

func FetchSubmit(ID, UID string) (*models.Submit, error) {
	var submit models.Submit
	result := database.DB.Where(
		models.Submit{
			UID: UID,
			ID:  ID,
		}).Find(&submit)
	if result.Error != nil {
		return nil, result.Error
	}
	return &submit, nil
}

func UpdateSubmit(data models.Submit) error {
	result := database.DB.Model(models.Submit{}).Where(
		&models.Submit{
			ID:   data.ID,
			Name: data.Name,
		}).Updates(&data)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
