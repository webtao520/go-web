package handlers

import (
	"chitchat/models"
	_ "fmt"
	"net/http"
)

// GET/login 登陆页面
func Login(writer http.ResponseWriter, request *http.Request){
	generateHTML(writer,nil,"auth.layout","navbar","login")
}


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
	//fmt.Println("================>",request.PostForm) //  map[email:[1183681473@qq.com] name:[呃呃呃] password:[123456ee ]]
	user:=models.User{
		Name: request.PostFormValue("name"),
		Email: request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}
	
	if err := user.Create(); err != nil {
		danger(err,"Cannot create user")
	}
	http.Redirect(writer,request,"/login",302)  // 重定向
}

// POST / Authenticate 登陆验证
// 通过邮箱和密码字段对用户进行认证
// ResponseWriter接口被HTTP处理器用于构造HTTP回复。
// Request类型代表一个服务端接受到的或者客户端发送出去的HTTP请求。
func Authenticate(writer http.ResponseWriter, request *http.Request){
	  err:=request.ParseForm()
	  user,err:=models.UserByEmail(request.PostFormValue("email"))	
	  if err !=nil{
		   danger(err,"Cannot find user")
	  }  
	  if user.Password == models.Encrypt(request.PostFormValue("password")) {
		  session,err:=user.CreateSession()
		  if err !=nil {
			  danger(err,"Cannot create session")
		  }
		  cookie:=http.Cookie{
				Name: "_cookie",
				Value: session.Uuid,
				HttpOnly: true, // 通过js 脚本无法获取cookie 信息， 防止XSS攻击
		  }
		  http.SetCookie(writer,&cookie)
		  http.Redirect(writer,request,"/",302)
	  }else{
		  http.Redirect(writer,request,"/login",302)
	  }
	  //fmt.Printf("%+v", user)   // {Id:0 Uuid: Name: Email: Password: CreatedAt:0001-01-01 00:00:00 +0000 UTC}
}

// GET / Logout 
// 用户退出 删除数据中的 session
func Logout(writer http.ResponseWriter, request *http.Request) {
	cookie,err:=request.Cookie("_cookie")
	// Cookie返回请求中名为name的cookie，如果未找到该cookie会返回nil, ErrNoCookie。
	if err != http.ErrNoCookie{
		warning(err, "Failed to get cookie")
		session:=models.Session{Uuid: cookie.Value}
		session.DeleteByUUID()
	}
	http.Redirect(writer,request,"/",302)
}

