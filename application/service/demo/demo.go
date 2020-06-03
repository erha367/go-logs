package demo

import (
	"gin-example/application/database"
	"gin-example/application/entity/form"
	"gin-example/application/library"
	"gin-example/application/models"
	"go.uber.org/zap"
)

//创建demo
func CreateDemo(demo form.DemoSearch) uint {
	tx := database.Begin("demos")
	demoModel := models.Demo{
		Name: demo.Name,
		Age:  demo.Age,
	}

	tx.Create(&demoModel)
	library.Logger.Info("find the demo ", zap.Uint("insertId", demoModel.ID))
	tx.Commit()
	return demoModel.ID
}

//检索demo
func Search(ds form.DemoSearch) (users []models.Demo) {
	dbDemos := database.GetSlaveDB("demos")
	if ds.Name != "" {
		dbDemos = dbDemos.Where("name = ?", ds.Name)
	}

	if ds.Age >= 0 {
		//dbDemos = dbDemos.Where("age > ?", ds.Age)
	}
	dbDemos.Find(&users)
	return
}
