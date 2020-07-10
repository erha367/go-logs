package database

import (
	"eeo_webcast_service/application"
	"eeo_webcast_service/application/model"
)

func Migrate() {
	application.Bootstrap()
	db := GetMasterDB("demos")
	db.AutoMigrate(&model.Demo{})
}
