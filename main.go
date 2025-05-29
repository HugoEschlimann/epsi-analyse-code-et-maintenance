package main

import (
	"gin/database"
	"gin/routes"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
	_ "gin/docs"
)

// @title Library Management API
// @version 1.0
// @description This is a sample API for managing a library system.
// @host localhost:8080
// @BasePath /api
func main() {
	db := database.Init()
	router := routes.Setup(db)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run()
}
