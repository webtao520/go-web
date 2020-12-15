package main 

import (
	"fmt"
)

func main(){
	s:="hello"
	//将字符串转换为 []byte 类型
	c:=[]byte(s)  
	c[0]='c'
	fmt.Println(c) // [104 101 108 108 111]
	// 再转换回string 类型
	s2:=string(c)
	fmt.Printf("%s\n",s2) // cello
}