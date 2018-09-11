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
	var p Person
	//p := Person{"mirk", 'm', 18}
	p.SetInfoPointer("mirk", 'm', 18)
	//内部自动先把p, 转换为&p再调用, (&p).SetInfoPointer("mirk", 'm', 18)

	p.SetInfoValue("mirk", 'm', 18)
}