package main 

import (
    "fmt"
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

func  sayHelloWorld(w http.ResponseWriter, r *http.Request){
    params := mux.Vars(r)
    w.WriteHeader(http.StatusOK)  // 设置响应状态码为 200
    fmt.Fprintf(w, "Hello, %s!", params["name"])  // 发送响应到客户端
}

func main (){
	r:=mux.NewRouter()
	r.HandleFunc("/hello/{name:[a-z]+}", sayHelloWorld)
	log.Fatal(http.ListenAndServe(":8080",r))
}