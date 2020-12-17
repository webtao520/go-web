package main 

import (
	"fmt"
)

type Human struct {
	name string
	age  int 
	phone string
}

type Student  struct {
	Human // 匿名字段
	school  string
}

type Employee   struct {
	Human // 匿名字段
	company string
}

// 在Human上面定义一个method
func (h *Human) SayHi(){
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func main (){
	mark:=Student{Human:Human{"Mark", 25, "222-222-YYYY"},school:"MIT"}
	sam:=Employee{Human:Human{"Sam", 45, "111-888-XXXX"},company:"Golang Inc"}
	mark.SayHi()
    sam.SayHi()
}


/**
	PS D:\goLang\github\go-web\Go语言基础\面向对象> go run 4.go
	Hi, I am Mark you can call me on 222-222-YYYY
	Hi, I am Sam you can call me on 111-888-XXXX
*/