package admin

import (
	_ "fmt"
	"time"

	"cms/models/admin"
	"cms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type UserController struct {
	BaseController
}

// 站点用户列表页面
func (c *UserController) List() {
	limit, _ := beego.AppConfig.Int64("limit") // 一页的数量
	page, _ := c.GetInt64("page", 1)           // 页数
	offset := (page - 1) * limit               // 偏移量
	//用户筛选
	name := c.GetString("name")
	status, _ := c.GetInt("status", 0)

	c.Data["Status"] = status
	c.Data["Name"] = name

	//  数据库查询
	o := orm.NewOrm()

	var users []*admin.User
	// return a QuerySeter for table operations.
	qs := o.QueryTable(new(admin.User)) //QuerySeter  QueryTable("user"), QueryTable(&user{}) or QueryTable((*User)(nil)),

	// 状态 qs.Filter("UserName", "slene")
	if status != 0 {
		qs = qs.Filter("status", status)
	}

	// 名称
	if name != "" {
		qs = qs.Filter("name__icontains", name)
	}

	// 获取数据
	_, err := qs.OrderBy("-id").Limit(limit).Offset(offset).All(&users)
	if err != nil {
		c.Abort("404")
	}
	// 统计
	count, err := qs.Count()
	if err != nil {
		c.Abort("404")
	}
	c.Data["Data"] = &users
	c.Data["Paginator"] = utils.GenPaginator(page, limit, count)
	c.Data["StatusText"] = admin.Status
	c.TplName = "admin/user-list.html"
}

//  添加用户页面
func (c *UserController) Add() {
	c.TplName = "admin/user-add.html"
}

//保存用户数据
func (c *UserController) Save() {
	name := c.GetString("name")
	password := c.GetString("password")
	repassword := c.GetString("repassword")
	email := c.GetString("email")

	o := orm.NewOrm()
	//组合用户结构体数据
	user := admin.User{
		Name:     name,
		Password: utils.PasswordMD5(password, name),
		Email:    email,
		Created:  time.Now(),
		Status:   1,
	}

	// 定义响应map
	response := make(map[string]interface{})

	// 表单验证是用于数据验证和错误收集的模块
	valid := validation.Validation{}
	valid.Required(user.Name, "Name")
	valid.Required(user.Password, "Password")
	valid.Required(user.Email, "Email")
	valid.Email(user.Email, "Email")

	if password != repassword {
		response["msg"] = "新增失败"
		response["code"] = 500
		response["err"] = "密码不一致"
		c.Data["json"] = response
		c.ServeJSON()
		c.StopRun() // 提前终止运行

	}

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			//log.Println(err.Key, err.Message) //  Email  Must be a valid email address
			response["msg"] = "新增失败"
			response["code"] = 500
			response["err"] = err.Key + " " + err.Message
			c.Data["json"] = response
			c.ServeJSON()
			c.StopRun() // 提前终止运行
		}
	}

	if id, err := o.Insert(&user); err == nil {
		response["msg"] = "新增成功"
		response["code"] = 200
		response["id"] = id
	} else {
		// 失败
		response["msg"] = "新增失败!"
		response["code"] = 500
		response["err"] = err.Error()
	}

	c.Data["json"] = response
	// JSON 数据直接输出
	c.ServeJSON() // 调用 ServeJSON 之后，会设置 content-type 为 application/json，然后同时把数据进行 JSON 序列化输出。
	c.StopRun()
}

//  查看编辑页面
func (c *UserController) Put() {
	id, err := c.GetInt("id", 0)
	if id == 0 {
		c.Abort("404")
	}
	// 根据id读取基础数据
	o := orm.NewOrm()
	var users []*admin.User
	qs := o.QueryTable(new(admin.User))
	err = qs.Filter("id", id).One(&users)
	if err != nil {
		c.Abort("404")
	}
	// 0xc000053200 0xc000053228 11
	//fmt.Println(&users[0].Id, &users[0].Email, users[0].Name) // {3 11 834c5dcd6f8ed2605debbdcdc3b83dfb 11@qq.com 2021-01-05 14:19:32 +0800 CST 1}
	c.Data["Data"] = &users[0]
	c.TplName = "admin/user-edit.html"
}

//  更新用户数据
func (c *UserController) Update() {

	id, _ := c.GetInt("id", 0)
	name := c.GetString("name")
	password := c.GetString("password")
	repassword := c.GetString("repassword")
	email := c.GetString("email")
	response := make(map[string]interface{})

	o := orm.NewOrm()

	user := admin.User{Id: id}
	if o.Read(&user) == nil {
		user.Name = name
		user.Email = email
		if password != "" {
			user.Password = utils.PasswordMD5(password, name)
		}
		// 数据校验
		valid := validation.Validation{}
		valid.Required(user.Name, "Name")
		valid.Required(user.Email, "Email")
		valid.Email(user.Email, "Email")

		if password != repassword {
			response["msg"] = "新增失败! "
			response["code"] = 500
			response["err"] = "密码不一致"
			c.Data["json"] = response
			c.ServeJSON()
			c.StopRun()
		}

		// 查看验证是否有错误
		if valid.HasErrors() {
			// 如果有错误信息，证明验证没通过
			// 打印错误信息
			for _, err := range valid.Errors {
				response["msg"] = "新增失败"
				response["code"] = 500
				response["err"] = err.Key + " " + err.Message
				c.Data["json"] = response
				c.ServeJSON()
				c.StopRun()
			}
		}

		// 更新数据
		if _, err := o.Update(&user); err == nil {
			response["msg"] = "修改成功! "
			response["code"] = 200
			response["id"] = id
		} else {
			response["msg"] = "修改失败! "
			response["code"] = 500
			response["err"] = err.Error()
		}
	} else {
		response["msg"] = "修改失败! "
		response["code"] = 500
		response["err"] = "ID 不能为空! "
	}
	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()

	/*
		c.Data["json"] = c.Input()
		c.ServeJSON()
		c.StopRun()
	*/

}

// 用户状态更新
func (c *UserController) Delete() {
	id, _ := c.GetInt("id", 0)
	status, _ := c.GetInt("status", 1)

	//  设置响应map
	response := make(map[string]interface{})

	o := orm.NewOrm()
	user := admin.User{Id: id}
	if o.Read(&user) == nil {
		if status == 1 {
			status = 2
		} else {
			status = 1
		}
		user.Status = status

		// 更新状态数据
		if _, err := o.Update(&user); err == nil {
			response["msg"] = "禁用成功! "
			response["code"] = 200
			response["id"] = id
		} else {
			response["msg"] = "禁用失败! "
			response["code"] = 500
			response["err"] = err.Error()
		}
	} else {
		response["msg"] = "禁用失败! "
		response["code"] = 500
		response["err"] = "ID 不能为空! "
	}
	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()
}
