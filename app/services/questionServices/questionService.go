package questionServices

import (
	"QA-System-Server/app/models"
	"QA-System-Server/config/database"
	"strings"
)

func GetQuestions(id string) ([]models.QuestionSplit, error) {
	var questions []models.Question

	result := database.DB.Find(&questions).Where("id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	questionsSplit := make([]models.QuestionSplit, result.RowsAffected)

	for i, value := range questions {
		questionsSplit[i].Stem = value.Stem
		questionsSplit[i].TypeNum = value.TypeNum
		questionsSplit[i].Options = strings.Split(value.Options, ";")
	}
	return questionsSplit, nil
}
