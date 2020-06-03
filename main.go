package main

import (
	"gin-example/application"
	"gin-example/application/database"
	"gin-example/application/library"
	"gin-example/config"
	"gin-example/router"
	"github.com/gin-gonic/gin"
)

func main() {
	defer database.CloseDatabases()
	application.Bootstrap()
	database.InitDatabases() //初始化DB
	gin.SetMode(config.Mode())
	apiRouter := router.ApiRouter()
	err := apiRouter.Run()
	if err != nil {
		library.Logger.Fatal("项目启动失败")
	}
}
