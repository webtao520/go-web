package main

type Element interface{}
type List []Element // slice 切片


type  Person struct{
	name string
	age int 
}

// 打印 实现了fmt string 方法
func (p Person) String() string{
	return "(name: " + p.name + " - age: "+strconv.Itoa(p.age)+ " years)"
}


for index, element := range list{
	switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case Person:
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		default:
			fmt.Println("list[%d] is of a different type", index)
	}
  }
}