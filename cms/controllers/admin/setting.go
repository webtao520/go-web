package admin

import (
	_"fmt"
	"cms/models/admin"
	"github.com/astaxie/beego/orm"
	"cms/utils"
)

type SettingController struct {
	BaseController
}


func (c *SettingController) Add(){
	// 基础数据
	o:=orm.NewOrm()
	var setting []*admin.Setting
	o.QueryTable(new(admin.Setting)).All(&setting)
	for _,v:=range setting{
		c.Data[v.Name]=v.Value
	}
	path := utils.GetViewPaths()
	c.Data["View"] = path
	c.TplName = "admin/setting.html"
}

// 站点公告
func (c *SettingController) Notice(){
	// 加载基础数据
	o:=orm.NewOrm()
	var setting  admin.Setting
	o.QueryTable(new(admin.Setting)).Filter("name","notice") .One(&setting)
	// fmt.Println(setting)  //{notice 欢迎来到使用 Go Blog 。} [0xc000004940]
	c.Data["Notice"]=setting.Value
	c.TplName = "admin/notice.html"
}


// 保存站点公共
func (c *SettingController) NoticeSave(){
	response:=make(map[string]interface{})
	notice:=c.GetString("notice")
	o:=orm.NewOrm()
	// 开启事物
	err:=o.Begin()
	_,err=o.Delete(&admin.Setting{Name:"notice"})

	num,err:=o.Insert(&admin.Setting{Name:"notice",Value: notice})
	if err !=nil {
		err=o.Rollback()
	}else{
		err = o.Commit()
	}
	if err != nil {
		response["msg"] = "操作失败！"
		response["code"] = 500
		response["err"] = err.Error()
	} else {
		response["msg"] = "操作成功！"
		response["code"] = 200
		response["id"] = num
	}

	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()
}

// 保存站点设置
func  (c *SettingController)  Save(){
  // 设置响应map
  response:=make(map[string]interface{})

  title:=c.GetString("title")
  keyword:=c.GetString("keyword")
  description := c.GetString("description")
  tag:=c.GetString("tag")
  image := c.GetString("image")

  template := c.GetString("template")
  limit := c.GetString("limit")


  // 保存数据
  o:=orm.NewOrm()
  err:=o.Begin()
  // 处理历史数据
	//_, err = o.Delete(&admin.Setting{Name: "name"})
	_, err = o.Delete(&admin.Setting{Name: "title"})
	_, err = o.Delete(&admin.Setting{Name: "tag"})
	_, err = o.Delete(&admin.Setting{Name: "template"})
	//_, err = o.Delete(&admin.Setting{Name: "remark"})
	_, err = o.Delete(&admin.Setting{Name: "image"})
	_, err = o.Delete(&admin.Setting{Name: "keyword"})
	_, err = o.Delete(&admin.Setting{Name: "description"})
	//_, err = o.Delete(&admin.Setting{Name: "remark_markdown_doc"})
	//_, err = o.Delete(&admin.Setting{Name: "remark_html_code"})
	_, err = o.Delete(&admin.Setting{Name: "limit"})

	
	settings := []admin.Setting{
		{Name: "title", Value: title},
		{Name: "template", Value: template},
		//{Name: "name", Value: name},
		{Name: "limit", Value: limit},
		{Name: "tag", Value: tag},
		//{Name: "remark_markdown_doc", Value: remark_markdown_doc},
		//{Name: "remark_html_code", Value: remark_html_code},
		{Name: "image", Value: image},
		{Name: "keyword", Value: keyword},
		{Name: "description", Value: description},
	}
	// 插入一些模型到数据库  
	num,err:=o.InsertMulti(7,settings)

	// 事务处理
	if err !=nil {
		err =o.Rollback()
	}else{
		err=o.Commit()
	}

	if err !=nil {
		response["msg"] = "操作失败！"
		response["code"] = 500
		response["err"] = err.Error()
	}else {
		response["msg"] = "操作成功！"
		response["code"] = 200
		response["id"] = num
	}
	c.Data["json"]= response
	c.ServeJSON()
	c.StopRun()
}


