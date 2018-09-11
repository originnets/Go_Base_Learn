package main

import "fmt"

type Person struct {
	name string
	sex	byte
	age	int
}

func (p Person) SetInfoValue (n string, s byte, a int)  {
	p.name = n
	p.sex = s
	p.age = a
	fmt.Printf("SetInfoValue &p = %p\n", &p)
}

func (p * Person) SetInfoPointer (n string, s byte, a int)  {
	p.name = n
	p.sex = s
	p.age = a

	fmt.Printf("SetInfoPointer p = %p\n", p)
}


func main() {
	// 假如,结构体是一个指针变量他能够调用的方法,这些方法就是一个集合,简称方法集
	p := &Person{"mike", 'm', 18}

	// 把(*p)转换成p后再调用
	(*p).SetInfoPointer("mike", 'm', 18)
	//p.SetInfoPointer("mike", 'm', 18)

	//内部做的转换, 先把指针p,转换成*p后再调用
	//(*p)SetInfoValue("mike", 'm', 18)
	p.SetInfoValue("mike", 'm', 18)
}