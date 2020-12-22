package main

import (
	"net/http"
)

type HelloHander struct {

}

type WorldHander struct {

}


func (handler *HelloHander) ServeHTTP(w http.ResponseWriter,r *http.Request){
	sayHelloGolang(w, r)
}

func (handler *WorldHander) ServeHTTP(w http.ResponseWriter,r *http.Request){
	fmt.Fprintf(w, "Hello World!")
}

func sayHelloGolang(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello Golang!")
}

func main(){
	 hello:=HelloHander{}
	 world:=WorldHander{}
	 server:=http.Server{
		Addr: ":9091",
	 }

	 http.Handle("/hello", &hello)
	 http.Handle("/world", &world)
	 server.ListenAndServe()

}
