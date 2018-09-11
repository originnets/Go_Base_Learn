package main

import "fmt"

type Person struct {
	name string
	sex	byte
	age	int
}


type Student struct {
	Person  //只有类型, 没有名字 , 匿名字段,继承Person的所有的成员
	id int
	addr string
}

func main() {
	s1 := Student{Person{"mirk", 'm', 18}, 1, "bj"}
	fmt.Printf("s1 = %+v\n", s1)
	//赋值
	s1.name = "yoyo"
	s1.sex = 'f'
	s1.age = 22
	s1.id = 666
	s1.addr = "sz"

	s1.Person = Person{"go",'m',26 }
	fmt.Println(s1.name, s1.sex, s1.age, s1.addr, s1.id)
}
