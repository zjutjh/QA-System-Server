package router

import (
	"QA-System-Server/app/controllers/questionController"
	"github.com/gin-gonic/gin"
)

func questionRouterInit(r *gin.RouterGroup) {
	r.Any("/getQuestion", questionController.GetQuestions)
}
