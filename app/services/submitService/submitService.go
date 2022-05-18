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
