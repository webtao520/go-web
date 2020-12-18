package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	//"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
    r.ParseForm()       //解析url传递的参数，对于POST则解析响应包的主体（request body）
    //注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息 map[age:[666] password:[] username:[]]
	//fmt.Printf("%T\n",r.Form["username"]) // []string 切片
	//fmt.Println(r.Form["username"][0]) 
	
	/*
	_,err:=strconv.Atoi(r.Form.Get("age"))
	if err!=nil{
		fmt.Println("数字转化出错了，那么可能就不是数字")
	}
	*/



	/*
	if len(r.Form["username"][0]) == 0{
		fmt.Println("用户名不能为空...")
	}
	*/
	/*
	if r.Form.Get("username") == "" {
		fmt.Println("用户名不能为空...")
	}
	*/
	/*
	r.Form对不同类型的表单元素的留空有不同的处理， 对于空文本框、空文本区域以及文件上传，元素的值为空值,而如果是未选中的复选框和单选按钮，
	则根本不会在r.Form中产生相应条目，如果我们用上面例子中的方式去获取数据时程序就会报错。所以我们需要通过r.Form.Get()来获取值，因为如果字段不存在，
	通过该方式获取的是空值。但是通过r.Form.Get()只能获取单个的值，如果是map的值，必须通过上面的方式来获取。
	*/



    //fmt.Println("path-", r.URL.Path) // / 
    //fmt.Println("scheme-", r.URL.Scheme)
	//fmt.Println(r.Form["url_long"])
	//fmt.Println("username-:", r.Form["username"]) // [admin]
	/*
    for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Printf("%T\n",v)   // []srting
		fmt.Println(v)
        fmt.Println("val:", strings.Join(v, ""))
    }
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
	*/
}

func login(w http.ResponseWriter, r *http.Request) {
    fmt.Println("method:", r.Method) //获取请求的方法
    if r.Method == "GET" {
        t, _ := template.ParseFiles("login.gtpl")
        log.Println(t.Execute(w, nil))
    } else {
        // 请求的是登录数据，那么执行登录的逻辑判断
        fmt.Println("username:", r.FormValue("username"))
        fmt.Println("password:", r.Form["password"])
    }
}

func main() {
    http.HandleFunc("/", sayhelloName)       // 设置访问的路由  后端处理路由
    http.HandleFunc("/login", login)         // 设置访问的路由 前台路由
    err := http.ListenAndServe(":9090", nil) // 设置监听的端口
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