package admin

import (
	"github.com/astaxie/beego"
)

// Cron的CronController操作
type CronController struct {
	beego.Controller
}

// 文章抓取
func (c *CronController) GetAll() {
	c.TplName = "admin/article-get.html"
}
