package main 


import (
	"fmt"
)

func main(){
	
for i := 0; i < 5; i++ {
    defer fmt.Printf("%d ", i) //  先进后出 栈
}
}