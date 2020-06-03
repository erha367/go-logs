package router

import (
	"gin-example/application/controller/DemoController"
	"gin-example/router/middleware"
	"github.com/gin-gonic/gin"
)

func ApiRouter() (router *gin.Engine) {
	router = gin.New()
	//中间件
	router.Use(middleware.Logger())
	router.Use(gin.Recovery())

	group := router.Group("/demo")
	{
		group.GET("/list", DemoController.List)
		group.GET("/create", DemoController.Create)
	}
	return
}
