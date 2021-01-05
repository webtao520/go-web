package admin

import (
	"cms/models/admin"
	"cms/utils"

	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (ctl *LoginController) Sign() {
	//fmt.Println("sign.....")
	ctl.TplName = "admin/login.html"
}

// 结构体绑定方法
func (ct1 *LoginController) Login() {
	username := ct1.GetString("username")
	password := ct1.GetString("password")
	//  密码处理
	password = utils.PasswordMD5(password, username)
	// 声明一个响应变量
	response := make(map[string]interface{})
	// 处理登录逻辑
	if user, ok := admin.Login(username, password); ok {
		// 设置登录session
		ct1.SetSession("User", *user)
		response["code"] = 200
		response["msg"] = "登录成功"
	} else {
		response["code"] = 500
		response["msg"] = "登录失败！"
	}
	ct1.Data["json"] = response
	// 发送带有编码字符集的json响应。
	ct1.ServeJSON()
}
