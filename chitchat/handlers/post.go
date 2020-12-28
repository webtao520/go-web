package handlers

import (
	"fmt"
	"net/http"
	"chitchat/models"
)

//  POST /thread/post
//  再指定群组下创建新主题
func PostThread(writer http.ResponseWriter, request *http.Request) {
	sess,err:=session(writer,request)
	if err !=nil{
		// 未登陆 跳转道登陆页面
		http.Redirect(writer,request,"/login",302)
	}else{
		err =request.ParseForm()
		if err !=nil{
			danger(err, "Cannot parse form")
		}
		user,err:=sess.User()
		if err !=nil {
			danger(err, "Cannot get user from session")
		}
		//  接受 post 提交数据
		body:=request.PostFormValue("body")
		uuid:=request.PostFormValue("uuid")
		thread, _ := models.ThreadByUUID(uuid)
		// TODO 本地化
	    if _,err:=user.CreatePost(thread,body);err !=nil {
			danger(err, "Cannot create post")
		}
		// func Sprintf(format string, a ...interface{}) string
		// Sprintf根据format参数生成格式化的字符串并返回该字符串。
		url := fmt.Sprint("/thread/read?id=", uuid)
		http.Redirect(writer,request,url,302)
	}
}
