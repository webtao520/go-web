package main 


import (
	"fmt"
)

func main(){
	s:="hello"
	fmt.Println(s[0]) // 获取字节的值 &s[0] 是非法的
	fmt.Println(s[2:])
	s="c" + s[1:] // 字符串虽不能进行更改，但可进行切片操作
	fmt.Printf("%s\n",s) // cello
}