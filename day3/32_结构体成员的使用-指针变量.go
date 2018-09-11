package main

import "fmt"

//定义一个结构体类型
type Student struct {
	id int
	name string
	sex byte // 字符类型
	age int
	addr string
}


func main() {
	//指针有合法指向,才能操作成员
	//先定义一个结构体普通变量
	var s Student
	//再定义一个指针变量,保存s的地址
	p1 := &s
	//通过指针操作成员, pi.id 和(*p1).id 完全等价, 只能使用.运算符

	p1.id = 1
	(*p1).name = "mike" //字符串
	p1.sex = 'm'	//字符
	p1.age = 19
	p1.addr = "bj"
	fmt.Println("s = ", s)

	//通过new函数申请一个结构体
	p2 := new(Student)
	p2.id = 1
	(*p2).name = "mike" //字符串
	p2.sex = 'm'	//字符
	p2.age = 19
	p2.addr = "bj"
	fmt.Println("p2 = ", p2)
	fmt.Println("*p2 = ", *p2)
}