package admin

import (
	"github.com/astaxie/beego"
)

// Cron的CronController操作
type CronController struct {
	beego.Controller
}

// 注册内部控制器路由器
func (c *CronController) URLMapping() {
	// 将方法映射到函数
	c.Mapping("GetAll", c.GetAll)
}

// 文章抓取
func (c *CronController) GetAll() {
	c.TplName = "admin/article-get.html"
}
