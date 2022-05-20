package questionController

import (
	"QA-System-Server/app/apiExpection"
	"QA-System-Server/app/models"
	"QA-System-Server/app/services/nameServices"
	"QA-System-Server/app/services/questionServices"
	"QA-System-Server/app/utils"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"strings"
)

func GetQuestions(c *gin.Context) {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
	id := c.Query("id")
	id_, _ := strconv.Atoi(id)

	questions, err := questionServices.GetQuestions(id_)
	if err != nil {
		log.Println("table question error")
		_ = c.AbortWithError(200, apiExpection.ServerError)
		return
	}
	name, err_ := nameServices.GetName(id)
	if err_ == apiExpection.ParamError {
		_ = c.AbortWithError(200, err_)
		return
	} else if err_ != nil {
		log.Println("table name_map error")
		_ = c.AbortWithError(200, apiExpection.ServerError)
		return
	}

	questionsSplit := make([]models.QuestionSplit, len(questions))
	for i, value := range questions {
		questionsSplit[i].Topic = value.Stem
		questionsSplit[i].TypeNum = value.TypeNum
		questionsSplit[i].Options = strings.Split(value.Options, ";")
	}

	utils.JsonSuccessResponse(c, questionsSplit, *name)
}
