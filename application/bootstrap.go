package application

import (
	"flag"
	"gin-example/application/library"
	"gin-example/config"
)

func Bootstrap() {
	flag.Parse()             //解析命令行
	config.LoadConfig()      //加载配置
	library.InitLogger()     //初始化日志
	//初始化其他
}
