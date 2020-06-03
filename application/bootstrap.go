package application

import (
	"flag"
	"gin-example/application/database"
	"gin-example/application/library"
	"gin-example/config"
)

func Bootstrap() {
	flag.Parse()             //解析命令行
	config.LoadConfig()      //加载配置
	library.InitLogger()     //初始化日志
	database.InitDatabases() //初始化DB
	//初始化其他
}
