package main

import (
	"eeo_webcast_service/application"
	"eeo_webcast_service/application/database"
	"eeo_webcast_service/application/library"
	"eeo_webcast_service/config"
	"eeo_webcast_service/router"
	"github.com/gin-gonic/gin"
)

func main() {
	defer database.CloseDatabases()
	application.Bootstrap()
	database.InitDatabases() //初始化DB
	gin.SetMode(config.Mode())
	apiRouter := router.ApiRouter()
	err := apiRouter.Run(`:88`)
	if err != nil {
		library.Logger.Fatal("项目启动失败")
	}
}
