package main

import (
	"chitchat/config"
	"chitchat/routes"
	"log"
	"net/http"
)

func main (){
	startWebServer()
}

// 通过指定端口启动web服务器
func startWebServer(){
    r:=routes.NewRouter() // 通过 router.go 中定义的路由器来分发请求
	// 处理静态资源文件
	assets := http.FileServer(http.Dir(config.ViperConfig.App.Static)) //  &{public}
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))

	http.Handle("/", r)
	log.Println("Starting HTTP service at " + config.ViperConfig.App.Address)
	err := http.ListenAndServe(config.ViperConfig.App.Address, nil)
	if err !=nil {
		log.Println("An error occured starting HTTP listener at "+  config.ViperConfig.App.Address)
		log.Println("Error: "+err.Error())
	}
	
}