package main

import "fmt"

type Person struct {
	name string
	sex	byte
	age	int
}


//修改成员变量的值

//参数为普通变量,非指针,值语义,一份拷贝(值传递)
func (p Person) SetinfoValue(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
	fmt.Println("p = ", p)
	fmt.Printf("SetinfoValue &p = %p\n", &p)
}

//接收者为指针变量,引用传递
func (p *Person) SetinfoPointer(n string, s byte, a int) {
	p.name = n
	p.sex = s
	p.age = a
	fmt.Printf("SetinfoPointer *p = %v\n", *p)
}


func main() {
	var p1 Person
	var p2 Person
	fmt.Printf("&p1 = %p\n", &p1) //打印地址
	//值语义
	p1.SetinfoValue("mike",'m', 18)
	fmt.Println("p1 = ", p1)

	//引用语义
	p2.SetinfoPointer("mike",'m', 18)
	fmt.Println("p2 = ", p2)
}