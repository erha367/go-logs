package demo

import (
	"eeo_webcast_service/application/database"
	"eeo_webcast_service/application/entity/form"
	"eeo_webcast_service/application/model"
)

//创建demo
func CreateDemo(demo form.DemoSearch) {

}

//检索demo
func Search(ds form.DemoSearch) model.EeoWebcastCastInfo{
	var info model.EeoWebcastCastInfo
	db := database.GetSlaveDB("eo_oslive")
	db.Where("account = ?",ds.Name).First(&info)
	return info
}
