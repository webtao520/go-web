package routes

import (
	"chitchat/handlers"
	"net/http"
)

// 定义一个 WebRoute 结构体用于存放单个路由
type WebRoute struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// 声明 WebRoutes 切片存放所有 Web 路由
type WebRoutes []WebRoute

// 定义所有 Web 路由
var webRoutes = WebRoutes{
	// 首页
	{
		"home",
		"GET",
		"/",
		handlers.Index,
	},
	// 注册
	{
		"signup",
		"GET",
		"/signup",
		handlers.Signup,
	},
	// 注册信息入表
	{
		"signupAccount",
		"POST",
		"/signup_account",
		handlers.SignupAccount,
	},
	//登陆
	{
		"login",
		"GET",
		"/login",
		handlers.Login,
	},
	// 登陆验证
	{
		"auth",
		"POST",
		"/authenticate",
		handlers.Authenticate,
	},
	// 退出
	{
		"logout",
		"GET",
		"/logout",
		handlers.Logout,
	},
	// 创建主题页面
	{
		"newThread",
		"GET",
		"/thread/new",
		handlers.NewThread,
	},
	//  创建主题后端逻辑
	{
		"createThread",
		"POST",
		"/thread/create",
		handlers.CreateThread,
	},
	// 回复页面
	{
		"readThread",
		"GET",
		"/thread/read",
		handlers.ReadThread,
	},
	// 回复逻辑
	{
		"postThread",
		"POST",
		"/thread/post",
		handlers.PostThread,
	},
}
