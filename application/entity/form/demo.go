package form

type DemoSearch struct {
	Name string `form:"name" binding:"required"`
	Age  uint   `form:"age"`
}
