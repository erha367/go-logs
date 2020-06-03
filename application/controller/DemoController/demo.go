package DemoController

import (
	"gin-example/application/entity/form"
	"gin-example/application/library"
	demoService "gin-example/application/service/demo"
	"github.com/gin-gonic/gin"
)

func List(context *gin.Context) {
	var demoSearch form.DemoSearch
	err := context.ShouldBind(&demoSearch)
	if err != nil {
		library.FailJson(context, err)
	} else {
		list := demoService.Search(demoSearch)
		library.OkJson(context, list)
	}
}

func Create(context *gin.Context) {
	var demo form.DemoSearch
	err := context.ShouldBind(&demo)
	if err != nil {
		library.FailJson(context, err)
	} else {
		id := demoService.CreateDemo(demo)
		library.OkJson(context, id)
	}
}
