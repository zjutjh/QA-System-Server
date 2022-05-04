package main

import (
	"QA-System-Server/config/database"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()

	r := gin.Default()
	r.Use(cors.Default())
}
