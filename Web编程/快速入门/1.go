package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloWorld(w http.ResponseWriter, r *http.Request) {
	// ParseForm解析URL中的查询字符串，并将解析结果更新到r.Form字段
	r.ParseForm()       //  解析参数
	fmt.Println(r.Form) // 在服务端打印请求参数
	//fmt.Println("URL:", r.URL.Path) // 请求 URL
	//fmt.Println("Scheme", r.URL.Scheme)

	for k, v := range r.Form {
		fmt.Println(k, ":", strings.Join(v, ""))
	}

	/**
	func (r *Request) ParseForm() error
	ParseForm解析URL中的查询字符串，并将解析结果更新到r.Form字段。


	Form url.Values
	// PostForm是解析好的POST或PUT的表单数据。
	// 本字段只有在调用ParseForm后才有效。在客户端，会忽略请求中的本字段而使用Body替代。
	*/

}

func main() {
	// func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))
	http.HandleFunc("/", sayHelloWorld)      // HandleFunc注册一个处理器函数handler和对应的模式pattern。
	err := http.ListenAndServe(":9091", nil) //  调用的处理器，如为nil会调用http.DefaultServeMux
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
