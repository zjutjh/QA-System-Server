package examController

import (
	"QA-System-Server/app/apiExpection"
	"QA-System-Server/app/models"
	"QA-System-Server/app/services/examNameServices"
	"QA-System-Server/app/services/questionServices"
	"QA-System-Server/app/utils"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"strings"
)

// GetExam
//
//	@Description: 获取试卷功能的Controller
//	@param c gin.Context
func GetExam(c *gin.Context) {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
	idRaw := c.Query("id")
	timeRaw := c.Query("time")

	id, err := strconv.Atoi(idRaw)
	if err != nil {
		log.Println("exam id error")
		_ = c.AbortWithError(200, apiExpection.ParamError)
		return
	}
	time, err := strconv.ParseInt(timeRaw, 10, 64)
	if err != nil {
		log.Println("timestamp error")
		_ = c.AbortWithError(200, apiExpection.ParamError)
		return
	}

	examNameMaps, err := examNameServices.GetExamName(idRaw)
	if err != nil {
		log.Println("name_map table error")
		_ = c.AbortWithError(200, apiExpection.ServerError)
		return
	}
	if len(examNameMaps) == 0 {
		log.Println("this exam is not existing")
		_ = c.AbortWithError(200, apiExpection.ParamError)
		return
	}
	DeadLine := examNameMaps[0].DeadLine.Time
	ExamName := examNameMaps[0].ExamName
	if time > DeadLine.Unix() {
		utils.JsonSuccessResponse(c, gin.H{
			"data": "超出试卷作答时间",
			"name": ExamName,
		}, "EXPIRED")
		return
	}

	questions, err := questionServices.GetQuestions(id)
	if err != nil {
		log.Println("exam table error")
		_ = c.AbortWithError(200, apiExpection.ServerError)
		return
	}

	questionsSplit := make([]models.QuestionSplit, len(questions))
	for i, value := range questions {
		questionsSplit[i].Topic = value.Topic
		questionsSplit[i].TypeNum = value.TypeNum
		questionsSplit[i].Options = strings.Split(value.Options, ";")
	}

	utils.JsonSuccessResponse(c, gin.H{
		"data": questionsSplit,
		"name": ExamName,
	}, "SUCCESS")
}
