package main

import (
	"QA-System-Server/app/midwares"
	"QA-System-Server/config/database"
	"QA-System-Server/config/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	database.Init()

	r := gin.Default()
	r.Use(cors.Default())
	r.Use(midwares.ErrHandler())
	r.NoMethod(midwares.HandleNotFound)
	r.NoRoute(midwares.HandleNotFound)

	router.Init(r)

	err := r.Run()
	if err != nil {
		log.Fatal("ServerStartFailed", err)
	}
}
