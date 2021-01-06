package utils

import (
	"fmt"
	"io/ioutil"
)

func GetViewPaths() (name []string) {
	/*
		func ReadDir
		func ReadDir(dirname string) ([]os.FileInfo, error)
		返回dirname指定的目录的目录信息的有序列表。
	*/
	dir, err := ioutil.ReadDir("views/home")
	//fmt.Println("=====>", dir)
	if err != nil {
		fmt.Println(err)
		return name
	}

	for _, fi := range dir {
		if fi.IsDir() { // 目录，递归遍历
			name = append(name, fi.Name()) // 切片中添加元素
		}
	}
	return
}
