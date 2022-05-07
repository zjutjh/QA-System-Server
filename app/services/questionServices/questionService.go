package questionServices

import (
	"QA-System-Server/app/models"
	"QA-System-Server/config/database"
	"math"
	"sort"
	"strconv"
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

func GetScore(scoreForm models.ScoreForm) (float64, error) {
	var questions []models.Question
	var sum int64 = 0

	result := database.DB.Find(&questions).Where("id = ?", scoreForm.ID)
	if result.Error != nil {
		return -1, result.Error
	}

	sort.SliceStable(scoreForm.Ans, func(i, j int) bool {
		return scoreForm.Ans[i].ID < scoreForm.Ans[j].ID
	})

	for i, value := range questions {
		flag := true
		answers := strings.Split(value.Answer, ";")

		if len(answers) != len(scoreForm.Ans[i].Key) {
			continue
		}
		for j := 0; j < len(answers) && flag; j++ {
			key, _ := strconv.Atoi(answers[j])
			if key != scoreForm.Ans[i].Key[j] {
				flag = false
				break
			}
		}

		if !flag {
			continue
		}

		sum++
	}

	return math.Trunc((100.0/float64(result.RowsAffected*sum))*1e2+0.5) * 1e-2, nil
}
