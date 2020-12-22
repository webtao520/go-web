package main 

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type HelloWorldHandler struct {

}

func sayHelloWorld(w http.ResponseWriter, r *http.Request)  {
    params := mux.Vars(r)
    w.WriteHeader(http.StatusOK)  // 设置响应状态码为 200
    fmt.Fprintf(w, "Hello, %s!", params["name"])  // 发送响应到客户端
}

//  接收者，绑定方法
func (handler *HelloWorldHandler)  ServeHTTP(w http.ResponseWriter, r *http.Request){
	   params:=mux.Vars(r) // 解析路由参数
	   w.WriteHeader(http.StatusOK) // 设置响应状态码为200
	   fmt.Fprintf(w,"你好, %s!",params["name"]) // 发送响应到客户端
}


func main (){
    r := mux.NewRouter()
    r.HandleFunc("/hello/{name:[a-z]+}", sayHelloWorld)
    r.Handle("/zh/hello/{name}", &HelloWorldHandler{})
    log.Fatal(http.ListenAndServe(":8080", r))
}