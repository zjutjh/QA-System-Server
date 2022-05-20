package questionServices

import (
	"QA-System-Server/app/models"
	"QA-System-Server/config/database"
)

func GetQuestions(id int) ([]models.Question, error) {
	var questions []models.Question

	result := database.DB.Where(
		models.Question{
			QuestionnaireID: id,
		}).Find(&questions)
	if result.Error != nil {
		return nil, result.Error
	}

	return questions, nil
}
