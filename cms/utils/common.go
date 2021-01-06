/*
   公共函数库
*/
package utils

import (
	"crypto/md5"
	"encoding/hex" // hex包实现了16进制字符表示的编解码
	"errors"
	_ "fmt"
	"strconv"
)

func PasswordMD5(passwd, salt string) string {
	h := md5.New()
	// 后面增加一个无意义字符串
	h.Write([]byte(passwd + salt + "@.YnO-"))
	cipherStr := h.Sum(nil)
	result := hex.EncodeToString(cipherStr)
	return result
}

// ToInt64 类型转换，获得int64
func ToInt64(v interface{}) (re int64, err error) {
	switch v.(type) {
	case string:
		/**
		返回字符串表示的整数值，接受正负号。
		base指定进制（2到36），如果base为0，则会从字符串前置判断，"0x"是16进制，"0"是8进制，否则是10进制；
		bitSize指定结果必须能无溢出赋值的整数类型，0、8、16、32、64 分别代表 int、int8、int16、int32、int64；返回的err是*NumErr类型的，如果语法有误，
		err.Error = ErrSyntax；如果结果超出类型范围 err.Error = ErrRange。
		*/
		re, err = strconv.ParseInt(v.(string), 10, 64) // 返回字符串表示的整数值，接受正负号 func ParseInt(s string, base int, bitSize int) (i int64, err error)
	case float64:
		re = int64(v.(float64))
	case float32:
		re = int64(v.(float32))
	case int64:
		re = v.(int64)
	case int32:
		re = int64(v.(int32)) // TODO   老版本code   (  re = v.(int64)  )
	default:
		err = errors.New("不能转换")
	}
	return
}
