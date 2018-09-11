package main

import "fmt"

type mystr string //自定义类型,给一个类型改名

type Person struct { //定义一个结构体
	name string
	sex	byte
	age	int
}


type Student struct {
	Person  //结构体匿名字段
	int		//基础类型的匿名字段
	mystr
}

func main()  {
	s := Student{Person{"mirk", 'm', 18}, 666, "hehehe"}
	s.Person = Person{"go", 'm', 22}
	fmt.Printf("s = %+v\n", s)
	fmt.Println(s.Person, s.int, s.mystr)
	fmt.Println(s.name,s.sex,s.age, s.int, s.mystr)
}