package utils

import (
	"fmt"
	"strconv"
)

// 在模板中注册函数
func IndexForOne(i int,p,limit int64) int64{
	// fmt.Printf("%T",i) int
	s:=strconv.Itoa(i)  // Itoa 是 FormatInt(i, 10) 的简写。
	index,_:=strconv.ParseInt(s,10,64)
	return (p-1)*limit + index + 1

}

func IndexDecrOne(i interface{}) int64 {
	//fmt.Printf("%T\n",i) // int64
	index, _ := ToInt64(i)
	return index - 1
}

func IndexAddOne(i interface{}) int64 {
	// fmt.Printf("%T\n",i) // int64
	index, _ := ToInt64(i)
	return index + 1
}