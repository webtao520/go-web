package main

import (
	"fmt"
)

type Men interface {
	SayHi()
	Sing(lyrics string)
}

type Human struct {
	name string 
	age int 
	phone string
}

type Student struct {
	Human //匿名字段
	school string
	loan float32
}

type Employee struct {
    Human //匿名字段
    company string
    money float32
}

// Human 实现SayHi方法
func(h Human) SayHi(){
  fmt.Printf("Hi, I am %s you can call me on %s\n",h.name,h.phone)
}

//Human实现Sing方法
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

// Employee 重载Human的SayHi方法
func(e Employee) SayHi(){
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
	e.company, e.phone)
}

// Interface Men 被Human，Student，Employee 实现


func main(){
		mike:=Student{Human: Human{"Mike",25,"222-222-XXX"},school:"MIT",loan:0.00}
		paul:=Student{Human:Human{"Paul",26,"111-222-XXX"},school:"Harvard",loan:100}
		sam:=Employee{Human:Human{"Sam", 36, "444-222-XXX"},company:"Golang Inc.",money:1000}
		tom := Employee{Human:Human{"Tom", 37, "222-444-XXX"}, company:"Things Ltd.", money:5000}

		// 定义Men类型的变量i 接口类型
		var i Men 

		// i 能存储Student
		i=mike 
		fmt.Println("This is Mike,a Student:")
		i.SayHi()
		i.Sing("November rain")

	    //i也能存储Employee
		i = tom
		fmt.Println("This is tom, an Employee:")
		i.SayHi()
		i.Sing("Born to be wild")

		   //定义了slice Men
		fmt.Println("Let's use a slice of Men and see what happens")

		x:=make([]Men,3)
		// 这三个都是不同类型的元素，但是他们实现了interface同一个接口
		x[0],x[1],x[2]=paul,sam,mike
		
		for _,value:=range x{
			//fmt.Printf("%T\n",value) // main.Student
			value.SayHi()
		}


}



