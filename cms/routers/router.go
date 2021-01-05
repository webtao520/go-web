package routers

import (
	"cms/controllers/admin"

	"github.com/astaxie/beego"
)

func init() {
	/*
		 NewNamespace(prefix string, funcs …interface{})
		 得到新的名称空间
		初始化 namespace 对象,下面这些函数都是 namespace 对象的方法,但是强烈推荐使用 NS 开头的相应函数注册，
		因为这样更容易通过 gofmt 工具看的更清楚路由的级别关系
	*/
	adminNs := beego.NewNamespace("/admin",
		beego.NSRouter("/user", &admin.UserController{}, "get:List"),

		// 后台登录
		/*
			     func NSRouter(rootpath string, c ControllerInterface, mappingMethods ...string) LinkNamespace
				调用命名空间路由器
		*/
		beego.NSRouter("/login", &admin.LoginController{}, "Get:Sign;Post:Login"),
	)
	/*
		func AddNamespace(nl ...*Namespace)
		将名称空间注册到beego。处理程序支持多命名空间
	*/
	beego.AddNamespace(adminNs)

	beego.Router("/admin", &admin.MainController{}, "get:Index")
}
