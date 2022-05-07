package router

import (
	"QA-System-Server/app/controllers/scoreController"
	"github.com/gin-gonic/gin"
)

func scoreRouterInit(r *gin.RouterGroup) {
	r.Any("/submit", scoreController.GetScore)
}
