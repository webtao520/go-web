package admin

import (
	"cms/models/admin"
	"cms/utils"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type ArticleController struct {
	BaseController
}

/**
List 列表数据执行sql

SELECT
	T0.`id`,
	T0.`title`,
	T0.`tag`,
	T0.`remark`,
	T0.`desc`,
	T0.`html`,
	T0.`created`,
	T0.`updated`,
	T0.`status`,
	T0.`pv`,
	T0.`review`,
	T0.`recommend`,
	T0.`like`,
	T0.`other`,
	T0.`url`,
	T0.`cover`,
	T0.`user_id`,
	T0.`category_id`,
	T1.`id`,
	T1.`name`,
	T1.`password`,
	T1.`email`,
	T1.`created`,
	T1.`status`,
	T2.`id`,
	T2.`name`,
	T2.`pid`,
	T2.`sort`,
	T2.`status`
FROM
	`article` T0
	INNER JOIN `user` T1 ON T1.`id` = T0.`user_id`
	INNER JOIN `category` T2 ON T2.`id` = T0.`category_id`
WHERE
	T1.`name` IS NOT NULL
	AND T2.`name` IS NOT NULL
	AND T0.`status` < 3
ORDER BY
	T0.`id` DESC,
	T0.`pv` DESC
	LIMIT 15


================================================================================================================================


SELECT
	COUNT( * )
FROM
	`article` T0
	INNER JOIN `user` T1 ON T1.`id` = T0.`user_id`
	INNER JOIN `category` T2 ON T2.`id` = T0.`category_id`
WHERE
	T1.`name` IS NOT NULL
	AND T2.`name` IS NOT NULL
	AND T0.`status` < 3
*/

// 控制器绑定 List 方法
func (c *ArticleController) List() {
	o := orm.NewOrm()

	limit := int64(15)
	page, _ := c.GetInt64("page", 1) // 页数
	offset := (page - 1) * limit     // 偏移量

	start := c.GetString("start")
	end := c.GetString("end")
	status, _ := c.GetInt("status", 0)
	title := c.GetString("title")

	// 前台渲染数据
	c.Data["start"] = start
	c.Data["End"] = end
	c.Data["Status"] = status
	c.Data["Title"] = title

	article := new(admin.Article) // 开辟内存地址
	var articles []*admin.Article
	qs := o.QueryTable(article)
	// 增加sql 过滤
	qs = qs.Filter("User__Name__isnull", false)
	//where User.Name IS NOT NULL
	qs = qs.Filter("Category__Name__isnull", false)
	//where Category.Name IS NOT  NULL

	// 状态
	if status != 0 {
		qs = qs.Filter("status", status)
		// where  status = status
	} else {
		qs = qs.Filter("status__lt", 3)
		// where status > <3
	}

	// 开始时间
	if start != "" {
		qs = qs.Filter("created__gte", start)
	}

	// 结束时间
	if end != "" {
		qs = qs.Filter("created__lte", end)
	}

	// 标题
	if title != "" {
		qs = qs.Filter("title__icontains", title)
		// where title like '%title%'
	}

	// 获取数据 将使用左连接加载所有相关字段 减号 - 表示 DESC「降序」 的排列
	qs.OrderBy("-id", "-pv").RelatedSel().Limit(limit).Offset(offset).All(&articles)

	// 统计
	count, _ := qs.Count()

	c.Data["Data"] = &articles
	c.Data["Paginator"] = utils.GenPaginator(page, limit, count)
	c.Data["StatusText"] = admin.Status
	c.Data["RecommendText"] = admin.Recommend

	c.TplName = "admin/article-list.html"
}

/*
func Add(c *ArticleController){

}
*/

//  后台添加文章 实例化 生产ArticleController的变量， 这样就能够调用 Add 方法
func (c *ArticleController) Add() {
	/**
	SELECT
		T0.`id`,
		T0.`name`,
		T0.`pid`,
		T0.`sort`,
		T0.`status`
	FROM
		`category` T0
	WHERE
		T0.`status` = 1
	*/

	//  获取分类信息
	o := orm.NewOrm()
	category := new(admin.Category) // 返回的是指针 开辟新的内存空间
	// All 结果集切片
	var categorys []*admin.Category
	qs := o.QueryTable(category) // 获取 QuerySeter 对象
	qs = qs.Filter("status", 1)  // where status =1
	qs.All(&categorys, "Id", "Name")
	c.Data["Category"] = categorys
	c.TplName = "admin/article-add.html"
}

// 保存文章
func (c *ArticleController) Save() {
	title := c.GetString("title")  //d标题
	tag := c.GetString("tag")      // tag标签
	cate, _ := c.GetInt("cate", 0) // 分类
	//fmt.Printf("%T====%T===%T",title,tag,cate) // string====string===int
	remark := c.GetString("remark")
	desc := c.GetString("desc_content")
	html := c.GetString("desc_html")
	url := c.GetString("url")
	cover := c.GetString("cover")

	// 创建orm
	o := orm.NewOrm()
	//  初始化结构体
	/*
		Id        int
		Title     string
		Tag       string
		Remark    string
		Desc      string    `orm:"type(text)"`
		Html      string    `orm:"type(text)"`
		Created   time.Time `orm:"auto_now_add;type(datetime)"` // 第一次保存时才设置时间
		Updated   time.Time `orm:"auto_now;type(datetime)"`     // 每次 model 保存时都会对时间自动更新
		Status    int       `orm:"default(1)"`
		Pv        int       `orm:"default(0)"`
		Review    int       `orm:"default(0)"`
		Recommend int       `orm:"default(0)"`
		Like      int       `orm:"default(0)"`
		Other     string    `orm:"type(text)"`
		Url       string
		Cover     string
		User      *User     `orm:"rel(fk)"`  //  设置一对多关系
		Category  *Category `orm:"rel(one)"` // 设置一对一关系
	*/
	article := admin.Article{
		Title:    title,
		Tag:      tag,
		Desc:     desc,
		Html:     html,
		Remark:   remark,
		Url:      url,
		Cover:    cover,
		Status:   1,
		User:     &admin.User{3, "", "", "", time.Now(), 0}, // orm:rel(fk) 一对多的关系
		Category: &admin.Category{cate, "", 0, 0, 0},
	}

	/*
				{
					"Id": 0,
					"Title": "444",
					"Tag": "444",
					"Remark": "444",
					"Desc": "4444",
					"Html": "\u003cp\u003e4444\u003c/p\u003e\r\n",
					"Created": "0001-01-01T00:00:00Z",
					"Updated": "0001-01-01T00:00:00Z",
					"Status": 1,
					"Pv": 0,
					"Review": 0,
					"Recommend": 0,
					"Like": 0,
					"Other": "",
					"Url": "444",
					"Cover": "/static/uploads/2021011510201081.jpg",
					"User": {
						"Id": 1,
						"Name": "",
						"Password": "",
						"Email": "",
						"Created": "2021-01-15T10:20:13.6418482+08:00",
						"Status": 0
					},
					"Category": {
						"Id": 1,
						"Name": "",
						"Pid": 0,
						"Sort": 0,
						"Status": 0
					}
				}


		 //  打印调试数据
		 c.Data["json"]=&article
		 c.ServeJSON()
		 c.StopRun()
	*/

	// 设置响应信息
	response := make(map[string]interface{})

	// 数据校验
	valid := validation.Validation{}
	valid.Required(article.Title, "Title") //Required 不为空，即各个类型要求不为其零值
	valid.Required(article.Html, "Html")
	valid.Required(article.Tag, "Tag")
	valid.Required(article.Desc, "Desc")

	if valid.HasErrors() {
		// 如果有错误信息，证明验证没通过
		// 打印错误信息
		for _, err := range valid.Errors {
			response["msg"] = "新增失败!"
			response["code"] = 500
			response["err"] = err.Key + " " + err.Message
			c.Data["json"] = response
			c.ServeJSON()
			c.StopRun()
		}
	}

	if id, err := o.Insert(&article); err == nil {
		response["msg"] = "新增成功"
		response["code"] = 200
		response["id"] = id
	} else {
		response["msg"] = "新增失败"
		response["code"] = 500
		response["err"] = err.Error()
	}

	c.Data["json"] = response

	c.ServeJSON()
	c.StopRun()
}

// 是否推荐
func (c *ArticleController) Top() {
	id, _ := c.GetInt("id", 0)

	response := make(map[string]interface{})

	o := orm.NewOrm()
	article := admin.Article{Id: id}
	if o.Read(&article) == nil {

		recommend := article.Recommend
		if recommend == 0 {
			article.Recommend = 1
		} else {
			article.Recommend = 0
		}

		if _, err := o.Update(&article,"Recommend"); err == nil { // Update 默认更新所有的字段，可以更新指定的字段
			response["msg"] = "操作成功！"
			response["code"] = 200
			response["id"] = id
		} else {
			response["msg"] = "操作失败！"
			response["code"] = 500
			response["err"] = err.Error()
		}
	} else {
		response["msg"] = "删除失败！"
		response["code"] = 500
		response["err"] = "ID 不能为空！"
	}

	c.Data["json"] = response
	c.ServeJSON()
	c.StopRun()
}

// 删除文章
func (c *ArticleController) Delete(){
	// 获取删除 id 号
	id,_:=c.GetInt("id",0)	

	//响应信息 开辟内存空间  只开辟内存空间才能使用
	response :=make(map[string]interface{})


}