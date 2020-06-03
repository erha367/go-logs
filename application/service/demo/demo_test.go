package demo

import (
	"gin-example/application"
	"gin-example/application/entity/form"
	"testing"
)

var _ = func() bool {
	testing.Init()
	return true
}

func TestCreateDemo(t *testing.T) {
	application.Bootstrap()
	insertId := CreateDemo(form.DemoSearch{Name: "test", Age: 1})
	if insertId == 0 {
		t.Error("demo create fail")
	}
}

func TestSearch(t *testing.T) {
	application.Bootstrap()
	list := Search(form.DemoSearch{Name: "test", Age: 1})
	if len(list) < 0 {
		t.Error("demo search fail")
	}
}
