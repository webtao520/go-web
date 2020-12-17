package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
   "strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()       //解析url传递的参数，对于POST则解析响应包的主体（request body）
    //注意:如果没有调用ParseForm方法，下面无法获取表单的数据
    //fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
    //fmt.Println("path-", r.URL.Path) // / 
    //fmt.Println("scheme-", r.URL.Scheme)
	//fmt.Println(r.Form["url_long"])
	//fmt.Println("username-:", r.Form["username"]) // [admin]
    for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Printf("%T\n",v)  //[]srting
		fmt.Println(v)
        fmt.Println("val:", strings.Join(v, ""))
    }
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
	
}

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //获取请求的方法
    if r.Method == "GET" {
        t, _ := template.ParseFiles("login.gtpl")
        log.Println(t.Execute(w, nil))
    } else {
        //请求的是登录数据，那么执行登录的逻辑判断
        fmt.Println("username:", r.FormValue("username"))
        fmt.Println("password:", r.FormValue["password"])
    }
}

func main() {
    http.HandleFunc("/", sayhelloName)       //设置访问的路由  后端处理路由
    http.HandleFunc("/login", login)         //设置访问的路由 前台路由
    err := http.ListenAndServe(":9090", nil) //设置监听的端口
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}

/**
func Join(a []string, sep string) string
将一系列字符串连接为一个字符串，之间用sep来分隔。


Tips:
Request本身也提供了FormValue()函数来获取用户提交的参数。如r.Form[“username”]也可写成r.FormValue(“username”)。
调用r.FormValue时会自动调用r.ParseForm，
所以不必提前调用。r.FormValue只会返回同名参数中的第一个，若参数不存在则返回空字符串。
*/