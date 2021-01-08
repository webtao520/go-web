package admin

import (
	"cms/models/admin"
	"cms/utils/article"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"strings"

	"github.com/astaxie/beego/orm"
)

type  ArticleResourcesController struct {
	BaseController
}


// 手动抓取
func (c *ArticleResourcesController) GetArticle(){
   t1:=time.Now()
   //fmt.Println("======",t1) // 2021-01-06 14:23:03.6181887 +0800 CST m=+22.721840101
   response:=make(map[string]interface{})
	//  获取提交的url
	urlPath:=c.GetString("url")
	if urlPath == "" {
		response["msg"]="抓取失败! "
		response["code"]=500
		response["err"]="url must no be null"
		c.Data["json"]=response
		c.ServeJSON()
		c.StopRun()
	}
	/*
		func Parse(rawurl string) (url *URL, err error)
		Parse函数解析rawurl为一个URL结构体，rawurl可以是绝对地址，也可以是相对地址。
	*/
	u,err:=url.Parse(urlPath)
	if err !=nil {
	   response["msg"]="新增失败！"
	   response["code"]=500
	   response["err"]=err.Error()
	}

  
	//fmt.Println(u,"===",t1)  // http://gocn.vip === 2021-01-06 14:33:13.5311077 +0800 CST m=+21.500889501
	host:=u.Host
	ho:=strings.Split(host, ":")
	var data map[string]interface{} 
	//var md,title,cover,html string
	var title string
	//var html  string
	switch ho[0] {
	case "gocn.vip":
		gocn:=article.Gocn{}
		data=gocn.Get(urlPath)
		title = data["title"].(string) //  interface 类型转换其它类型
		fmt.Println(t1,title)
	case "book.douban.com":	
		 douban:=article.Douban{}
		 data=douban.Get(urlPath)
		 review := data["list"].([]map[string]string)

		 	// 数据库操作
	o:=orm.NewOrm()
	art:=admin.Article{
		Title: "",
		Review: len(review),
		Category: &admin.Category{Id:1},
		User:     &admin.User{Id: 1},
		Status:   1,
	}

	id,err:=o.Insert(&art)
	if err !=nil {
		response["msg"] = "新增失败！"
		response["code"] = 500
		response["err"] = err.Error()
		c.Data["json"] = response
		c.ServeJSON()
		c.StopRun()
	}

	aid:=int(id)
	if aid == 0 {
		response["msg"] = "新增失败！"
		response["code"] = 500
		response["err"] = "文章ID不能为空!"
		c.Data["json"] = response
		c.ServeJSON()
		c.StopRun()
	}

	 count := 0 
     for _,v := range  review{
		customerId,_:=admin.AddCustomer(
			&admin.Customer{
				Username: v["username"],
				Nickname: v["username"],
			})	
		rev:=admin.Review{}
		rev.ArticleId	=aid 
		rev.Review =  v["content"]
		rev.Name=v["username"]
		rev.Status = 1
		rev.Customer=&admin.Customer{Id: int(customerId)}
		loc, _ := time.LoadLocation("Local")
	

		//fmt.Println("时间：", v["commenttime"])
		the_time, _ := time.ParseInLocation("2006-01-02 15:04:05", v["commenttime"], loc)
		//fmt.Println("=========>>>>",the_time) //  0001-01-01 00:00:00 +0000 UTC
		rev.Created=the_time
		if _,err:=o.Insert(&rev);err !=nil{
			fmt.Println(err.Error())
			continue
		}
		count++
	 }
      // []map[string]string
	  response["msg"] = "总共评论 " + strconv.Itoa(len(review)) + "条，爬取成功 " + strconv.Itoa(count) + "条！"
	  response["code"] = 200
	  response["id"] = aid
	  response["title"] = art.Title
	  response["elapsed"] = fmt.Sprintf("%s", time.Since(t1)) // Since返回从t到现在经过的时间
	  c.Data["json"] = response
	  c.ServeJSON()
	  c.StopRun()
	default:
		response["msg"] = "新增失败！"
		response["code"] = 500
		response["err"] = "未知源!"
		c.Data["json"] = response
		c.ServeJSON()
		c.StopRun()
	 //fmt.Printf("%T\n",count) // int 
	}
	var cover = "90"
	o := orm.NewOrm()
	article := admin.Article{
		Title:    title,
		Tag:      "",
		Remark:   "",
		Status:   1,
		User:     &admin.User{1, "", "", "", time.Now(), 0},
		Category: &admin.Category{1, "", 0, 0, 0},
		Cover:    cover,
	}

	if id, err := o.Insert(&article); err == nil {
		response["msg"] = "新增成功！"
		response["code"] = 200
		response["id"] = id
		response["title"] = title
		response["elapsed"] = fmt.Sprintf("%s", time.Since(t1))
	} else {
		response["msg"] = "新增失败！"
		response["code"] = 500
		response["err"] = err.Error()
	}

	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()
}
