package main

import (
	"gin/database"
	"gin/routes"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
	_ "gin/docs"
)

func main() {
	db := database.Init()
	router := routes.Setup(db)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run()
}
