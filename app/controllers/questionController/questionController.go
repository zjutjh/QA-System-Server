package questionController

import (
	"QA-System-Server/app/apiExpection"
	"QA-System-Server/app/services/questionServices"
	"QA-System-Server/app/utils"
	"github.com/gin-gonic/gin"
)

func GetQuestions(c *gin.Context) {
	id := c.Param("id")
	questions, err := questionServices.GetQuestions(id)
	if err != nil {
		_ = c.AbortWithError(200, apiExpection.ServerError)
	} else {
		utils.JsonSuccessResponse(c, questions, id)
	}
}
