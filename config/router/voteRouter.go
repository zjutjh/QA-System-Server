package router

import (
	"QA-System-Server/app/controllers/voteController"
	"github.com/gin-gonic/gin"
)

func voteRouterInit(r *gin.RouterGroup) {
	r.GET("/getVoteList", voteController.GetVoteList)
	r.POST("/submitVote", voteController.SubmitVote)
	//TODO impl two controllers
}
