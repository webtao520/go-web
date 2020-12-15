package main 

import (
	"fmt"
)

func main(){
	// 声明一个数组 字节
	var array = [10]byte{'a','b','c','d','e','f','g','h','i','j'} 
	// 声明两个slice
	var aSlice,bSlice []byte 
	// 演示一些简单操作
	aSlice =  array[:3] // 等价于 aSlice =array[0:3]  aSlice 包含元素： a,b,c,  不包含索引是3的元素
	aSlice=array[5:] // // 等价于aSlice = array[5:10] aSlice包含元素: f,g,h,i,j
	aSlice=array[:] // 等价于aSlice = array[0:10] 这样aSlice包含了全部的元素

	// 从slice中获取slice
aSlice = array[3:7]  // aSlice包含元素: d,e,f,g，len=4，cap=7
bSlice = aSlice[1:3] // bSlice 包含aSlice[1], aSlice[2] 也就是含有: e,f
bSlice = aSlice[:3]  // bSlice 包含 aSlice[0], aSlice[1], aSlice[2] 也就是含有: d,e,f
bSlice = aSlice[0:5] // 对slice的slice可以在cap范围内扩展，此时bSlice包含：d,e,f,g,h
bSlice = aSlice[:]   // bSlice包含所有aSlice的元素: d,e,f,g


	
}