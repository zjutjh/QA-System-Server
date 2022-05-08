package questionServices

import (
	"QA-System-Server/app/models"
	"QA-System-Server/config/database"
)

func GetQuestions(id string) ([]models.Question, error) {
	var questions []models.Question

	result := database.DB.Where("questionnaire_id = ?", id).Find(&questions)
	if result.Error != nil {
		return nil, result.Error
	}

	return questions, nil
}
