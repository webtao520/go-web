package admin

import (
	"github.com/astaxie/beego/orm"
)

type Setting struct {
	Name  string `orm:"size(255);pk"`  // string 
	Value string `orm:"type(text)"`
}

//   需要在init中注册定义的model
func init (){
	orm.RegisterModel(new(Setting))
}