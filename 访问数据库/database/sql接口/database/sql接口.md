### database/sql接口

Go与PHP不同的地方是Go官方没有提供数据库驱动，而是为开发数据库驱动定义了一些标准接口，开发者可以根据定义的接口来开发相应的数据库驱动，
这样做有一个好处，只要是按照标准接口开发的代码， 以后需要迁移数据库时，不需要任何修改。那么Go都定义了哪些标准接口呢？让我们来详细的分析一下

### sql.Register
这个存在于database/sql的函数是用来注册数据库驱动的，当第三方开发者开发数据库驱动时，都会实现init函数，在init里面会调用这个Register(name string, 
driver driver.Driver)完成本驱动的注册。

我们来看一下mymysql、sqlite3的驱动里面都是怎么调用的：




