### interface

Go语言里面设计最精妙的应该算interface，它让面向对象，内容组织实现非常的方便，当你看完这一章，你就会被interface的巧妙设计所折服。

什么是interface
简单的说，interface是一组method签名的组合，我们通过interface来定义对象的一组行为。

我们前面一章最后一个例子中Student和Employee都能SayHi，虽然他们的内部实现不一样，但是那不重要，重要的是他们都能say hi

让我们来继续做更多的扩展，Student和Employee实现另一个方法Sing，然后Student实现方法BorrowMoney而Employee实现SpendSalary。

这样Student实现了三个方法：SayHi、Sing、BorrowMoney；而Employee实现了SayHi、Sing、SpendSalary。

上面这些方法的组合称为interface(被对象Student和Employee实现)。例如Student和Employee都实现了interface：SayHi和Sing，
也就是这两个对象是该interface类型。而Employee没有实现这个interface：SayHi、Sing和BorrowMoney，因为Employee没有实现BorrowMoney这个方法。

### interface类型

interface类型定义了一组方法，如果某个对象实现了某个接口的所有方法，则此对象就实现了此接口。详细的语法参考下面这个例子

    type Human struct {
        name string
        age int
        phone string
    }

    type Student struct {
        Human //匿名字段Human
        school string
        loan float32
    }

    type Employee struct {
        Human //匿名字段Human
        company string
        money float32
    }

    //Human对象实现Sayhi方法
    func (h *Human) SayHi() {
        fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
    }

    // Human对象实现Sing方法
    func (h *Human) Sing(lyrics string) {
        fmt.Println("La la, la la la, la la la la la...", lyrics)
    }

    //Human对象实现Guzzle方法
    func (h *Human) Guzzle(beerStein string) {
        fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
    }

    // Employee重载Human的Sayhi方法
    func (e *Employee) SayHi() {
        fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
            e.company, e.phone) //此句可以分成多行
    }

    //Student实现BorrowMoney方法
    func (s *Student) BorrowMoney(amount float32) {
        s.loan += amount // (again and again and...)
    }

    //Employee实现SpendSalary方法
    func (e *Employee) SpendSalary(amount float32) {
        e.money -= amount // More vodka please!!! Get me through the day!
    }

    // 定义interface
    type Men interface {
        SayHi()
        Sing(lyrics string)
        Guzzle(beerStein string)
    }

    type YoungChap interface {
        SayHi()
        Sing(song string)
        BorrowMoney(amount float32)
    }

    type ElderlyGent interface {
        SayHi()
        Sing(song string)
        SpendSalary(amount float32)
    }

通过上面的代码我们可以知道，interface可以被任意的对象实现。我们看到上面的Men interface被Human、Student和Employee实现。
同理，一个对象可以实现任意多个interface，例如上面的Student实现了Men和YoungChap两个interface。

最后，任意的类型都实现了空interface(我们这样定义：interface{})，也就是包含0个method的interface

### interface值

那么interface里面到底能存什么值呢？如果我们定义了一个interface的变量，那么这个变量里面可以存实现这个interface的任意类型的对象。
例如上面例子中，我们定义了一个Men interface类型的变量m，那么m里面可以存Human、Student或者Employee值。

因为m能够持有这三种类型的对象，所以我们可以定义一个包含Men类型元素的slice，这个slice可以被赋予实现了Men接口的任意结构的对象，这个和我们传统意义上面的slice有所不同。

让我们来看一下下面这个例子:

+ 案例
  * 1.go

通过上面的代码，你会发现interface就是一组抽象方法的集合，它必须由其他非interface类型实现，而不能自我实现，
Go通过interface实现了duck-typing:即”当看到一只鸟走起来像鸭子、游泳起来像鸭子、叫起来也像鸭子，那么这只鸟就可以被称为鸭子”  

### 空interface

空interface(interface{})不包含任何的method，正因为如此，所有的类型都实现了空interface。空interface对于描述起不到任何的作用(因为它不包含任何的method），
但是空interface在我们需要存储任意类型的数值的时候相当有用，因为它可以存储任意类型的数值。它有点类似于C语言的void*类型。

    // 定义a为空接口
    var a interface{}
    var i int =5 
    s:="hello world"
    // a 可以存储任意类型的数值
    a=i 
    a=s

一个函数把interface{}作为参数，那么他可以接受任意类型的值作为参数，如果一个函数返回interface{},那么也就可以返回任意类型的值。是不是很有用啊！

### interface函数参数

interface的变量可以持有任意实现该interface类型的对象，这给我们编写函数(包括method)提供了一些额外的思考，
我们是不是可以通过定义interface参数，让函数接受各种类型的参数。

举个例子：fmt.Println是我们常用的一个函数，但是你是否注意到它可以接受任意类型的数据。打开fmt的源码文件，你会看到这样一个定义:


type Stringer interface {
     String() string
}

也就是说，任何实现了String方法的类型都能作为参数被fmt.Println调用,让我们来试一试

+ 案例
  * 2.go

### interface变量存储的类型

我们知道interface的变量里面可以存储任意类型的数值(该类型实现了interface)。那么我们怎么反向知道这个变量里面实际保存了的是哪个类型的对象呢？目前常用的有两种方法：

Comma-ok断言

Go语言里面有一个语法，可以直接判断是否是该类型的变量： value, ok = element.(T)，这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型。

如果element里面确实存储了T类型的数值，那么ok返回true，否则返回false。

让我们通过一个例子来更加深入的理解。

+ 案例
  * 3.go

是不是很简单啊，同时你是否注意到了多个if里面，还记得我前面介绍流程时讲过，if里面允许初始化变量。

也许你注意到了，我们断言的类型越多，那么if else也就越多，所以才引出了下面要介绍的switch。


switch测试

最好的讲解就是代码例子，现在让我们重写上面的这个实现

+ 案例
  * 4.go

这里有一点需要强调的是：element.(type)语法不能在switch外的任何逻辑里面使用，如果你要在switch外面判断一个类型就使用comma-ok。

### 嵌入interface

Go里面真正吸引人的是它内置的逻辑语法，就像我们在学习Struct时学习的匿名字段，多么的优雅啊，那么相同的逻辑引入到interface里面，那不是更加完美了。
如果一个interface1作为interface2的一个嵌入字段，那么interface2隐式的包含了interface1里面的method。

我们可以看到源码包container/heap里面有这样的一个定义


    type Interface interface {
        sort.Interface //嵌入字段sort.Interface
        Push(x interface{}) //a Push method to push elements into the heap
        Pop() interface{} //a Pop elements that pops elements from the heap
    }
    我们看到sort.Interface其实就是嵌入字段，把sort.Interface的所有method给隐式的包含进来了。也就是下面三个方法：


    type Interface interface {
        // Len is the number of elements in the collection.
        Len() int
        // Less returns whether the element with index i should sort
        // before the element with index j.
        Less(i, j int) bool
        // Swap swaps the elements with indexes i and j.
        Swap(i, j int)
    }
    另一个例子就是io包下面的 io.ReadWriter ，它包含了io包下面的Reader和Writer两个interface：


    // io.ReadWriter
    type ReadWriter interface {
        Reader
        Writer
    }


