package utils

import (
	"strconv"
	"strings"
)

// 在模板中注册函数
func IndexForOne(i int, p, limit int64) int64 {
	//fmt.Printf("%T", i)  //int
	/*
		func FormatInt
		func FormatInt(i int64, base int) string
		返回i的base进制的字符串表示。base 必须在2到36之间，结果中会使用小写字母'a'到'z'表示大于10的数字。
	*/
	s := strconv.Itoa(i) // Itoa 是 FormatInt(i, 10) 的简写。  func Itoa(i int) string
	index, _ := strconv.ParseInt(s, 10, 64)
	return (p-1)*limit + index + 1

}

func IndexDecrOne(i interface{}) int64 {
	//fmt.Printf("%T\n", i) // int64
	index, _ := ToInt64(i)
	return index - 1
}

func IndexAddOne(i interface{}) int64 {
	// fmt.Printf("%T\n",i) // int64
	index, _ := ToInt64(i)
	return index + 1
}

func StringReplace(str, old, new string) string {
	return strings.Replace(str, old, new, -1)
}
