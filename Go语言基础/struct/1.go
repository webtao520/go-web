package main 

import (
	"fmt"
)

// 声明一个新的类型
type person struct {
	name string
	age int 
}

// 比较两个人的年纪，返回年纪大的那个人，并且返回年纪差
// struct 也是传值的
func Older(p1,p2 person)(person,int){
	if p1.age > p1.age { // 比较p1 和 p2 这两个人的年纪
		return p1,p1.age-p2.age
	}
	return p2,p2.age-p1.age
}

func main(){
	var tom person
	// 赋值初始化
	tom.name,tom.age="tom",18
	//两个字段都写情书的初始化
	bob:=person{age:25,name:"BOb"}
	//按照struct定义顺序初始化值
	paul:=person{"Paul",43}

	tb_Older, tb_diff := Older(tom, bob)
    tp_Older, tp_diff := Older(tom, paul)
    bp_Older, bp_diff := Older(bob, paul)

    fmt.Printf("Of %s and %s, %s is older by %d years\n",
        tom.name, bob.name, tb_Older.name, tb_diff)

    fmt.Printf("Of %s and %s, %s is older by %d years\n",
        tom.name, paul.name, tp_Older.name, tp_diff)

    fmt.Printf("Of %s and %s, %s is older by %d years\n",
        bob.name, paul.name, bp_Older.name, bp_diff)

}