package admin 


import (
	"github.com/astaxie/beego/orm"
)

type Category struct {
	Id int 
	Name string
	Pid int 
	Sort int 
	Status  int
}

// 需要在 init 中注册定义的model
func init (){
	orm.RegisterModel(new (Category))
}