package questionController

import (
	"QA-System-Server/app/apiExpection"
	"QA-System-Server/app/models"
	"QA-System-Server/app/services/questionServices"
	"QA-System-Server/app/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

func GetQuestions(c *gin.Context) {
	id := c.Query("id")
	println(id)
	questions, err := questionServices.GetQuestions(id)

	questionsSplit := make([]models.QuestionSplit, len(questions))
	for i, value := range questions {
		questionsSplit[i].Stem = value.Stem
		questionsSplit[i].TypeNum = value.TypeNum
		questionsSplit[i].Options = strings.Split(value.Options, ";")
	}

	if err != nil {
		_ = c.AbortWithError(200, apiExpection.ServerError)
	} else {
		utils.JsonSuccessResponse(c, questionsSplit, id)
	}
}
