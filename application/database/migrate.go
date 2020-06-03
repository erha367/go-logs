package database

import (
	"gin-example/application"
	"gin-example/application/models"
)

func Migrate() {
	application.Bootstrap()
	db := GetMasterDB("demos")
	db.AutoMigrate(&models.Demo{})
}
