package main

import (
	_ "cms/routers"
	db "cms/service/databsae"
	"cms/utils"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/config"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" //  register db `default`, sql: unknown driver "mysql" (forgotten import?)
	"github.com/sirupsen/logrus"
)

func init() {
	// 读取配置信息
	conf, err := config.NewConfig("ini", "conf/app.conf")
	orm.Debug = true

	if err != nil {
		// Fatalf在标准日志记录器上以Fatal级别记录消息，然后进程将退出并将状态设置为1。
		logrus.Fatalf(err.Error())
	}
	// 加载数据库驱动 实例化数据库驱动实例
	database, _ := db.NewDataBase(conf.String("db::dbType"))
	// RegisterDriver注册一个数据库驱动程序时使用指定驱动程序名，这可以定义驱动程序是哪个数据库类型。
	orm.RegisterDriver(database.GetDriverName(), database.GetDriver())
	// RegisterDataBase设置数据库连接参数。使用数据库驱动程序self dataSource args。
	orm.RegisterDataBase(database.GetAliasName(), database.GetDriverName(), database.GetStr())
	// 在模板中注册一个函数
	beego.AddFuncMap("IndexForOne", utils.IndexForOne)
	beego.AddFuncMap("IndexDecrOne", utils.IndexDecrOne)
	beego.AddFuncMap("IndexAddOne", utils.IndexAddOne)
	beego.AddFuncMap("StringReplace", utils.StringReplace)
}

func main() {

	beego.Run()
}
