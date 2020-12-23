package routes

import (
	"net/http"
	"chitchat/handlers"
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
var  webRoutes = WebRoutes{
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
	
	
}
