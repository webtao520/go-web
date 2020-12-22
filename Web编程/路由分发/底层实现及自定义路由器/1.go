package main

import (
    "fmt"
    "net/http"
)

type MyHander struct {

}

func (handler *MyHander) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
    if r.URL.Path == "/" {
        sayHelloGolang(w, r)
        return
    }
    http.NotFound(w, r)
    return
}

func sayHelloGolang(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello Golang!")
}

func main()  {
    handler := MyHander{}
    http.ListenAndServe(":9091", &handler)
}