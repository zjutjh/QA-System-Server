package questionServices

import (
	"QA-System-Server/app/models"
	"QA-System-Server/config/database"
	"strings"
)

func GetQuestions(id string) ([]models.QuestionSplit, error) {
	var questions []models.Question
	var questionsSplit []models.QuestionSplit

	result := database.DB.Find(&questions).Where("id = ?", id)
	if result.Error != nil {
		return nil, result.Error
	}

	for i, value := range questions {
		questionsSplit[i].Stem = value.Stem
		questionsSplit[i].TypeNum = value.TypeNum
		questionsSplit[i].Options = strings.Split(value.Options, ";")
	}
	return questionsSplit, nil
}
