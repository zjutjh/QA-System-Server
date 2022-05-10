package scoreController

import (
	"QA-System-Server/app/apiExpection"
	"QA-System-Server/app/models"
	"QA-System-Server/app/services/nameServices"
	"QA-System-Server/app/services/questionServices"
	"QA-System-Server/app/utils"
	"github.com/gin-gonic/gin"
	"math"
	"sort"
	"strings"
)

func GetScore(c *gin.Context) {
	var scoreForm models.ScoreForm
	sum := 0

	err := c.ShouldBindJSON(&scoreForm)
	if err != nil {
		_ = c.AbortWithError(200, apiExpection.ParamError)
		return
	}

	questions, err := questionServices.GetQuestions(scoreForm.PaperCode)
	if err != nil {
		_ = c.AbortWithError(200, apiExpection.ServerError)
	}

	name, err_ := nameServices.GetName(scoreForm.PaperCode)
	if err_ != nil {
		_ = c.AbortWithError(200, apiExpection.ServerError)
	}

	if len(questions) != len(scoreForm.AnsList) {
		_ = c.AbortWithError(200, apiExpection.ParamError)
	}
	sort.SliceStable(scoreForm.AnsList, func(i, j int) bool {
		return scoreForm.AnsList[i].ID < scoreForm.AnsList[j].ID
	})
	for i, value := range questions {
		flag := true
		answers := strings.Split(value.Answer, ";")

		if len(answers) != len(scoreForm.AnsList[i].Key) {
			continue
		}
		if len(scoreForm.AnsList[i].Key) > 1 {
			sort.SliceStable(scoreForm.AnsList[i].Key, func(j, k int) bool {
				return scoreForm.AnsList[i].Key[j] < scoreForm.AnsList[i].Key[k]
			})
		}
		for j := 0; j < len(answers) && flag; j++ {
			if answers[j] != scoreForm.AnsList[i].Key[j] {
				flag = false
				break
			}
		}
		if !flag {
			continue
		}
		sum++
	}
	score := math.Round(((100.0/float64(len(questions)))*float64(sum))*1e2+0.5) * 1e-2

	if err != nil {
		_ = c.AbortWithError(200, apiExpection.ServerError)
	} else {
		utils.JsonSuccessResponse(c, score, *name)
	}
}
