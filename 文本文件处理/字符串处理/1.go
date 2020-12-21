package main

import (
	"fmt"
	"strings"
)

func main(){
	//fmt.Println(strings.Contains("seafood","foo")) //  true
	//fmt.Println(strings.Contains("golang","go")) // true
	//s:=[]string{"foo","bar","baz"}
	//fmt.Println(strings.Join(s,", ")) // foo, bar, baz
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr"))
}