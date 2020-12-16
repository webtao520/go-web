### struct

Go语言中，也和C或者其他语言一样，我们可以声明新的类型，作为其它类型的属性或字段的容器。例如，我们可以创建一个自定义类型person代表一个人的实体。
这个实体拥有属性：姓名和年龄。这样的类型我们称之struct。如下代码所示:

type person struct {
    name string
    age int
}

看到了吗？声明一个struct如此简单，上面的类型包含有两个字段

一个string类型的字段name，用来保存用户名称这个属性
一个int类型的字段age,用来保存用户年龄这个属性
如何使用struct呢？请看下面的代码

    type person struct {
        name string
        age int 
    }

    var P person // p 现在就是person类型的变量了

    P.name = "Astaxie"  // 赋值"Astaxie"给P的name属性.
    P.age = 25  // 赋值"25"给变量P的age属性
    fmt.Printf("The person's name is %s", P.name)  // 访问P的name属性.

除了上面这种P的声明使用之外，还有另外几种声明使用方式：

1. 按照顺序提供初始化值
  P := person{“Tom”, 25}
2. 通过field:value 的方式初始化，这样可以任意顺序
P:=person{age:24,name:"TOM"} 
3. 当然也可以通过new函数分配一个指针，此处P的类型为*person
 P:=new(person)

 下面我们看一个完整的使用struct的例子

+ 案例
  * 1.go


### struct的匿名字段

我们上面介绍了如何定义一个struct，定义的时候是字段名与其类型一一对应，实际上Go支持只提供类型，而不写字段名的方式，
也就是匿名字段，也称为嵌入字段。

当匿名字段是一个struct的时候，那么这个struct所拥有的全部字段都被隐式地引入了当前定义的这个struct。

让我们来看一个例子，让上面说的这些更具体化

+ 案例
  * 2.go

我们看到Student访问属性age和name的时候，就像访问自己所有用的字段一样，对，匿名字段就是这样，能够实现字段的继承。是不是很酷啊？还有比这个更酷的呢，
那就是student还能访问Human这个字段作为字段名。请看下面的代码，是不是更酷了。

mark.Human = Human{"Marcus", 55, 220}
mark.Human.age -= 1

通过匿名访问和修改字段相当的有用，但是不仅仅是struct字段哦，所有的内置类型和自定义类型都是可以作为匿名字段的。请看下面的例子

+ 案例
  * 3.go

从上面例子我们看出来struct不仅仅能够将struct作为匿名字段，自定义类型、内置类型都可以作为匿名字段，
而且可以在相应的字段上面进行函数操作（如例子中的append）。
这里有一个问题：如果human里面有一个字段叫做phone，而student也有一个字段叫做phone，那么该怎么办呢？
Go里面很简单的解决了这个问题，最外层的优先访问，也就是当你通过student.phone访问的时候，是访问student里面的字段，而不是human里面的字段。
这样就允许我们去重载通过匿名字段继承的一些字段，当然如果我们想访问重载后对应匿名类型里面的字段，可以通过匿名字段名来访问。请看下面的例子

+ 案例
  * 4.go
