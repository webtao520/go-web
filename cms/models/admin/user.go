package admin

import (
	_ "fmt"
	"time"

	"github.com/astaxie/beego/orm"
)

type User struct {
	Id int 
	Name  string
	Password string
	Email string
	Created  time.Time `orm:"auto_now_add;type(datetime)"`  // time.Time 字段的对应 db 类型使用 datetime  auto_now_add 第一次保存时才设置时间
	Status    int
}

func init (){
	// 需要在init 中注册定义的model
	orm.RegisterModel(new(User))
}

// 登录逻辑
func Login(username,password string)(*User,bool){
	o:=orm.NewOrm()
	var (
		user User
		err error
	)
	ok:=false
	o.Using("default")
	qs:=o.QueryTable(&user)
	cond:=orm.NewCondition() // 返回一个新的结构体
	cond=cond.And("status",1).And("Name",username).Or("Email",username) // 并将表达式添加到条件
	qs=qs.SetCond(cond)
	//查询一行数据并映射到容器
	if err=qs.One(&user);err==nil{
		if user.Password == password {
			ok = true
		}
	}
	return &user,ok
}