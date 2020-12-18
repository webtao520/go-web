package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // 获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix() // 获取当前Unix时间戳
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("login.gtpl")
		// func (t *Template) Execute(wr io.Writer, data interface{}) error
		// Execute方法将解析好的模板应用到data上，并将输出写入wr。如果执行时出现错误，会停止执行，但有可能已经写入wr部分数据。模板可以安全的并发执行。
		t.Execute(w, token)
	} else {
		// 请求的是登录数据，那莪执行登录的逻辑判断
		r.ParseForm() // ParseForm解析URL中的查询字符串，并将解析结果更新到r.Form字段
		token := r.Form.Get("token")
		if token != "" {
			//  验证token的合法性
		} else {
			// 不存在token报错
		}
		fmt.Println("username length:", len(r.Form["username"][0]))                 // 字符串长度
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端 // 返回s的HTML转义等价表示字符串。
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		// func HTMLEscape(w io.Writer, b []byte)
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端  // 函数返回其所有参数文本表示的HTML转义等价表示字符串。
	}

}

func main() {
	http.HandleFunc("/login", login)         // 注册一个处理器函数handler和对应的模式pattern
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
