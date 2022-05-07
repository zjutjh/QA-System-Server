package scoreController

import (
	"QA-System-Server/app/apiExpection"
	"QA-System-Server/app/models"
	"QA-System-Server/app/services/questionServices"
	"QA-System-Server/app/utils"
	"github.com/gin-gonic/gin"
)

func GetScore(c *gin.Context) {
	var scoreForm models.ScoreForm

	err := c.ShouldBindJSON(&scoreForm)
	if err != nil {
		_ = c.AbortWithError(200, apiExpection.ParamError)
		return
	}

	score, err := questionServices.GetScore(scoreForm)
	if err != nil {
		_ = c.AbortWithError(200, apiExpection.ServerError)
	} else {
		utils.JsonSuccessResponse(c, score, scoreForm.ID)
	}
}
