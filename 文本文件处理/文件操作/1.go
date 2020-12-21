package main

import (
	//"fmt"
	"fmt"
	"os"
)

func main (){
	os.Mkdir("astaxie", 0777)
	os.MkdirAll("astaxie/test1/test2",0777)
	err:=os.Remove("astaxie") //  当目录下有文件或者其他目录时会报错
	if err !=nil {
		fmt.Println(err)
	}
    os.RemoveAll("astaxie")
}