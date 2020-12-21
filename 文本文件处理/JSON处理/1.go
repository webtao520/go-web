package main 

import (
	//"encoding/json"
	"fmt"
)

type Server struct {
	 ServerName string
	 ServerIP string
}

type Serverslice struct {
	Servers []Server
}

func main (){
	var s Serverslice
	 // 对s进行字面量
	//  fmt.Printf("%T\n",s)
	str:= `{"servers":[{"serverName":"Shanghai_VPN","serverIP":"127.0.0.1"},{"serverName":"Beijing_VPN","serverIP":"127.0.0.2"}]}`
	json.Unmarshal([]byte(str),&s)
	fmt.Println(s)
}