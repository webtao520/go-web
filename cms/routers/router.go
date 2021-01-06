package routers

import (
	"cms/controllers/admin"
	"cms/controllers/common"

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
		/*
			func NSInclude
			func NSInclude(cList ...ControllerInterface) LinkNamespace
			命名空间包含ControllerInterface
		*/
		beego.NSInclude(
			&admin.MenuController{}, // 导航管理
			&admin.CronController{}, // 文章管理 - 文章抓取
		),
		//************************* 文章管理 - 文章抓取 ***********************************
		//beego.NSRouter("/get-review", &admin.CronController{}, "get:GetReview"),

		//************************* 用户管理 ***********************************
		beego.NSRouter("/user", &admin.UserController{}, "get:List;post:Save"),
		beego.NSRouter("/user/add", &admin.UserController{}, "get:Add"),
		beego.NSRouter("/user/edit", &admin.UserController{}, "get:Put"),
		beego.NSRouter("/user/update", &admin.UserController{}, "Post:Update"),
		beego.NSRouter("/user/status", &admin.UserController{}, "Post:Delete"),
		// 后台登录
		/*
			     func NSRouter(rootpath string, c ControllerInterface, mappingMethods ...string) LinkNamespace
				调用命名空间路由器
		*/
		beego.NSRouter("/login", &admin.LoginController{}, "Get:Sign;Post:Login"),

		//************************* 站点设置 ***********************************
		beego.NSRouter("/setting", &admin.SettingController{}, "get:Add"),
		beego.NSRouter("/notice", &admin.SettingController{}, "get:Notice"),
		beego.NSRouter("/notice/save", &admin.SettingController{}, "post:NoticeSave"),
		beego.NSRouter("/setting/save", &admin.SettingController{}, "post:Save"),
	)
	/*
		func AddNamespace(nl ...*Namespace)
		将名称空间注册到beego。处理程序支持多命名空间
	*/
	beego.AddNamespace(adminNs)

	beego.Router("/admin", &admin.MainController{}, "get:Index")

	// 文件上传
	beego.Router("/uploads.html", &common.UploadsController{}, "Post:Uploads")

}
