package handlers

import (
	"net/http"
)

// 注册页面 GET
func Signup(writer http.ResponseWriter, request *http.Request){
	generateHTML(writer, nil, "auth.layout", "navbar", "signup")
}

// 注册新用户 POST
func SignupAccount(writer http.ResponseWriter, request *http.Request){
	err:=request.ParseForm()
	if err !=nil {
		danger(err,"Cannot parse form")
	}
	// 模型和表映射
}

