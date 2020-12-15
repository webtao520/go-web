package main 

import (
	"fmt"
)

const (
	x= iota // x == 0
	y=iota  // y == 1
	z=iota  // z == 2
	w    // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w = iota，因此w == 3。其实上面y和z可同样不用"= iota"
)

const v = iota // 每遇到一个const 关键字， iota 就会重置，此时 v == 0 

const (
	h, i, j = iota, iota, iota //h=0,i=0,j=0 iota在同一行值相同
)

const (
	a       = iota //a=0
    b       = "B"
    c       = iota             //c=2
    d, e, f = iota, iota, iota //d=3,e=3,f=3
    g       = iota   
)


func main() {
	fmt.Println(a, b, c, d, e, f, g, h, i, j, x, y, z, w, v)
	
	fmt.Println("====g===>",g)
}

/**
PS D:\goLang\github\go-web\Go语言基础\Go基础> go run 4.go
0 B 2 3 3 3 4 0 0 0 0 1 2 3 0
*/