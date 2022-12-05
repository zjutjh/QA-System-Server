package voteController

import (
	"QA-System-Server/app/apiExpection"
	"QA-System-Server/app/models"
	"QA-System-Server/app/services/nameServices"
	"QA-System-Server/app/services/voteServices"
	"QA-System-Server/app/utils"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

func GetVoteList(c *gin.Context) {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
	idRaw := c.Query("voteid")
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

	nameMaps, err := nameServices.GetName(id)
	if err != nil {
		log.Println("name_map table error")
		_ = c.AbortWithError(200, apiExpection.ServerError)
		return
	}
	if len(nameMaps) == 0 {
		log.Println("this exam is not existing")
		_ = c.AbortWithError(200, apiExpection.ParamError)
		return
	}
	DeadLine := nameMaps[0].DeadLine.Time
	ExamName := nameMaps[0].Name
	if time > DeadLine.Unix() {
		utils.JsonSuccessResponse(c, gin.H{
			"data": "投票已截止",
			"name": ExamName,
		}, "EXPIRED")
		return
	}

	votes, err := voteServices.GetVoteList(id)
	if err != nil {
		log.Println("vote table error")
		_ = c.AbortWithError(200, apiExpection.ServerError)
		return
	}

	voteViews := make([]models.VoteView, len(votes))
	for i, vote := range votes {
		voteViews[i].CandidateId = vote.CandidateId
		voteViews[i].Describe = vote.Describe
	}

	utils.JsonSuccessResponse(c, gin.H{
		"data": voteViews,
		"name": ExamName,
	}, "SUCCESS")
}

func SubmitVote(c *gin.Context) {
	log.SetFlags(log.Lshortfile | log.Ldate | log.Lmicroseconds)
}
