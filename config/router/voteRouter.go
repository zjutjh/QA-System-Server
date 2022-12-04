package router

import (
	"QA-System-Server/app/controllers/examController"
	"QA-System-Server/app/controllers/scoreController"
	"github.com/gin-gonic/gin"
)

func voteRouterInit(r *gin.RouterGroup) {
	r.GET("/getVoteList", examController.GetExam)
	r.POST("/submitVote", scoreController.GetScore)
	//TODO impl tow controllers
}
