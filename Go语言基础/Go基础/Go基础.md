### Go基础

这小节我们将要介绍如何定义变量、常量、Go内置类型以及Go程序设计中的一些技巧。

### 定义变量

Go语言里面定义变量有多种方式。

使用var关键字是Go最基本的定义变量方式，与C语言不同的是Go把变量类型放在变量名后面：


//定义一个名称为“variableName”，类型为"type"的变量

var variableName type

定义多个变量


//定义三个类型都是“type”的变量

var vname1, vname2, vname3 type

定义变量并初始化值


//初始化“variableName”的变量为“value”值，类型是“type”

var variableName type = value

同时初始化多个变量

/*
    定义三个类型都是"type"的变量,并且分别初始化为相应的值
    vname1为v1，vname2为v2，vname3为v3
*/

var vname1, vname2, vname3 type= v1, v2, v3

你是不是觉得上面这样的定义有点繁琐？没关系，因为Go语言的设计者也发现了，有一种写法可以让它变得简单一点。
我们可以直接忽略类型声明，那么上面的代码变成这样了：


/*
    定义三个变量，它们分别初始化为相应的值
    vname1为v1，vname2为v2，vname3为v3
    然后Go会根据其相应值的类型来帮你初始化它们
*/
    var vname1, vname2, vname3 = v1, v2, v3

你觉得上面的还是有些繁琐？好吧，我也觉得。让我们继续简化：


/*
    定义三个变量，它们分别初始化为相应的值
    vname1为v1，vname2为v2，vname3为v3
    编译器会根据初始化的值自动推导出相应的类型
*/
    vname1, vname2, vname3 := v1, v2, v3

现在是不是看上去非常简洁了？:=这个符号直接取代了var和type,这种形式叫做简短声明。不过它有一个限制，那就是它只能用在函数内部；
在函数外部使用则会无法编译通过，所以一般用var方式来定义全局变量。

_（下划线）是个特殊的变量名，任何赋予它的值都会被丢弃。在这个例子中，我们将值35赋予b，并同时丢弃34：

    _, b := 34, 35
Go对于已声明但未使用的变量会在编译阶段报错，比如下面的代码就会产生一个错误：声明了i但未使用。


    package main

    func main() {
        var i int
    }


### 常量

所谓常量，也就是在程序编译阶段就确定下来的值，而程序在运行时无法改变该值。在Go程序中，
常量可定义为数值、布尔值或字符串等类型。

它的语法如下：

    const constantName = value
    //如果需要，也可以明确指定常量的类型：
    const Pi float32 = 3.1415926


下面是一些常量声明的例子：

    const Pi = 3.1415926
    const i = 10000
    const MaxThread = 10
    const prefix = "astaxie_"

Go 常量和一般程序语言不同的是，可以指定相当多的小数位数(例如200位)，
若指定給float32自动缩短为32bit，指定给float64自动缩短为64bit，详情参考链接    

### 内置基础类型 

Boolean

在Go中，布尔值的类型为bool，值是true或false，默认为false。

// 实例代码
var isActive  bool  // 全局变量声明
var enabled, disabled = true,false 
func test(){
    var available  bool  // 一般声明
    valid:=false  // 简短声明
    available =true // 赋值操作
}


### 数值类型

整数类型有无符号和带符号两种。Go同时支持int和uint，这两种类型的长度相同，但具体长度取决于不同编译器的实现。
Go里面也有直接定义好位数的类型：rune, int8, int16, int32, int64和byte, uint8, uint16, uint32, uint64。
其中rune是int32的别称，byte是uint8的别称。

需要注意的一点是，这些类型的变量之间不允许互相赋值或操作，不然会在编译时引起编译器报错。

如下的代码会产生错误：invalid operation: a + b (mismatched types int8 and int32)

var a int8


