package router

import (
	"github.com/gin-gonic/gin"
	"eeo_webcast_service/application/controller/DemoController"
	"eeo_webcast_service/router/middleware"
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
