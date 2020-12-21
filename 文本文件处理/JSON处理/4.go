package main 

import (
	//"fmt"
	"encoding/json"
	"os"
)

type Server struct {
	 // ID 不会导出到json中
	 ID int `json:"_"`
	 // ServerName2 的值会进行二次JSON编码
	ServerName  string `json:"serverName"`
	ServerName2 string `json:"serverName2,string"`
	// 如果ServerIP 为空，则不输出到json串中
	ServerIP string `json:"serverIP,omitempty"`	
}


func main(){
	s:=Server{
		ID:3,
		ServerName:`Go "1.0" `,
		ServerName2:`Go "1.0" `,
		ServerIP:``,
   }
   
   b,_:=json.Marshal(s)
   //fmt.Println(string(b)) // {"_":3,"serverName":"Go \"1.0\" ","serverName2":"\"Go \"1.0\" \""}
   os.Stdout.Write(b)  // {"_":3,"serverName":"Go \"1.0\" ","serverName2":"\"Go \"1.0\" \""}
}



