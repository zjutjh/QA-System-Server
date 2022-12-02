package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func JsonSuccessResponse(c *gin.Context, data interface{}, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"msg":  msg,
		"data": data,
	})
}
