package models

import (
	"github.com/jinzhu/gorm"
)

type Demo struct {
	Model
	Name string `sql:"type:varchar(100);not null;default:''"`
	Age  uint   `sql:"not null;default:0"`
}

//钩子函数
func (demo *Demo) BeforeSave(scope *gorm.Scope) (err error) {
	demo.Age = demo.Age + 1
	return nil
}
