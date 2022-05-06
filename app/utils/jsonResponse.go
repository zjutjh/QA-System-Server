package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JsonSuccessResponse(c *gin.Context, data interface{}, id string) {
	c.JSON(http.StatusOK, gin.H{
		"msg":  "SUCCESS",
		"name": id,
		"data": data,
	})
}
