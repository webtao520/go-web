package admin

import (
	_ "fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"cms/models/admin"
	"cms/utils"
)

type UserController struct {
	BaseController
}

// 站点用户列表页面
func (c *UserController) List(){
	limit,_:=beego.AppConfig.Int64("limit") // 一页的数量
	page,_:=c.GetInt64("page",1) // 页数
	offset:=(page - 1) *limit // 偏移量
	//用户筛选
	name:=c.GetString("name")
	status,_:=c.GetInt("status", 0)
	
	c.Data["Status"] = status
	c.Data["Name"] = name

	//  数据库查询
	o:=orm.NewOrm()
	
	var users []*admin.User
	// return a QuerySeter for table operations.
	qs:=o.QueryTable(new(admin.User))   //QuerySeter  QueryTable("user"), QueryTable(&user{}) or QueryTable((*User)(nil)),

	// 状态 qs.Filter("UserName", "slene")
	if status !=0 {
		qs = qs.Filter("status",status)
	}

	// 名称
	if name !="" {
		qs =qs.Filter("name__icontains",name)
	}

	// 获取数据
	_,err:=qs.OrderBy("-id").Limit(limit).Offset(offset).All(&users)
	if err !=nil {
		c.Abort("404")
	}
	// 统计
	count,err:=qs.Count()
	if err !=nil {
		c.Abort("404")
	}
	c.Data["Data"] = &users
	c.Data["Paginator"] = utils.GenPaginator(page, limit, count)
	c.Data["StatusText"] = admin.Status
	c.TplName = "admin/user-list.html"
}
