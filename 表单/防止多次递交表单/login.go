package main

import (
	"fmt"
	"net/http"
	"log"
)

func login(w http.ResponseWriter, r *http.Request){
	 fmt.Println("method:",r.Method) // 获取请求的方法
	 if r.Method == "GET" {
		  
	 }
	 
}

func main(){
	http.HandleFunc("/login",login) // 注册一个处理器函数handler和对应的模式pattern
    err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err !=nil{
		log.Fatal("ListenAndServe: ", err)
	}
}