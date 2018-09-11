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
	fmt.Printf("SetInfoValue: %p, %v\n", &p, p)
}

func (p * Person) SetInfoPointer (n string, s byte, a int)  {
	p.name = n
	p.sex = s
	p.age = a
	fmt.Printf("SetInfoPointer: %p, %v\n", &p, p)
}


func main() {
	p := Person{"mrik", 'm', 18}
	fmt.Printf("main: %p, %v\n", &p, p)
	p.SetInfoPointer("miaa", 'f', 28) // 传统调用方式
	fmt.Printf("main: %p, %v\n", &p, p)
	// 保存方式入口地址

	pFunc := p.SetInfoPointer //这个就是方法值,调用函数时,无需再传递接收者,隐藏了接收者

	pFunc("miar", 'm', 30) //等价于 p.SetInfoPointer()
	fmt.Printf("main: %p, %v\n", &p, p)

	vFunc := p.SetInfoValue
	vFunc("miac", 'f', 29) //等价于 p.SetInfoValue() // 值引用 , 再main函数中不会发生改变
	fmt.Printf("main: %p, %v\n", &p, p)

}