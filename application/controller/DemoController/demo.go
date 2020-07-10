package DemoController

import (
	"eeo_webcast_service/application/entity/form"
	"eeo_webcast_service/application/library"
	demoService "eeo_webcast_service/application/service/demo"
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

}
