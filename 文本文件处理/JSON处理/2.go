package main

import (
	"encoding/json"
	"fmt"
)

func main(){
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err:=json.Unmarshal(b,&f) // Unmarshal函数解析json编码的数据并将结果存入v指向的值。
	if err !=nil{

	}else{
		fmt.Println(f)  //map[Age:6 Name:Wednesday Parents:[Gomez Morticia]]
	}



	//fmt.Printf("======%T\n",f)  //map[string]interface {}
	//fmt.Println("=================>",f)

	m := f.(map[string]interface{}) // 将 f 转化map[string]interface{}



	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string", vv)
		case int:
			fmt.Println(k, "is int", vv)
		case float64:
			fmt.Println(k,"is float64",vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}


}