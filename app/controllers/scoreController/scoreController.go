package scoreController

import (
	"QA-System-Server/app/apiExpection"
	"QA-System-Server/app/models"
	"QA-System-Server/app/services/questionServices"
	"QA-System-Server/app/utils"
	"github.com/gin-gonic/gin"
	"math"
	"sort"
	"strconv"
	"strings"
)

func GetScore(c *gin.Context) {
	var scoreForm models.ScoreForm

	err := c.ShouldBindJSON(&scoreForm)
	if err != nil {
		_ = c.AbortWithError(200, apiExpection.ParamError)
		return
	}

	var questions []models.Question
	sum := 0

	questions, err = questionServices.GetQuestions(scoreForm.ID)
	if err != nil {
		_ = c.AbortWithError(200, apiExpection.ServerError)
	}

	if len(questions) != len(scoreForm.Ans) {
		_ = c.AbortWithError(200, apiExpection.ParamError)
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

		sort.SliceStable(scoreForm.Ans[i].Key, func(i, j int) bool {
			return scoreForm.Ans[i].Key[i] < scoreForm.Ans[i].Key[j]
		})

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

	score := math.Trunc((100.0/float64(len(questions)*sum))*1e2+0.5) * 1e-2

	if err != nil {
		_ = c.AbortWithError(200, apiExpection.ServerError)
	} else {
		utils.JsonSuccessResponse(c, score, scoreForm.ID)
	}
}
