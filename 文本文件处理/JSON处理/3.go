package main

import (
	"fmt"
)

func main() {
	mapA := make([string]interface{})

	mapB := make([string]interface{})

	mapA["name"] = "小文"

	mapA["age"] = 25

	mapB["mapA"] = mapA

	for k, v := range mapB {

		if k == "mapA" {
			for _, v1 := range v.(map[string]interface{}) { //这里把v再转成mapA的类型即可
				fmt.Println(v1)
			}
		}
	}
}
