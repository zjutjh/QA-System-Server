package router

import (
	"QA-System-Server/app/controllers/examController"
	"github.com/gin-gonic/gin"
)

func examRouterInit(r *gin.RouterGroup) {
	r.GET("/getExam", examController.GetExam)
	r.POST("/submit", examController.GetScore)
}
