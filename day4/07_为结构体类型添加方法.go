package main

import "fmt"

type Person struct {
	name string
	sex	byte
	age	int
}


//带有接收者的函数叫方法
func (tmp Person) PrintInfo()  {
	fmt.Println("tmp = " ,tmp)
}



// 通过一个函数,给成员赋值

func (p *Person) SetInfo(name string, sex byte, age int) {
	p.name = name
	p.sex = sex
	p.age = age
}

// 只要接收者类型不一样,这个方法就算同名,也是不同的方法,不会出现重复定义的错误
type long int
func (tep long) test() {

}
type char byte
func (tep char) test() {

}

//type porinter *int
//pointer 为接收者类型, 它本身不能是指针类型
//func (tmp porinter ) test02() {
//
//}

func main() {
	//定义同时初始化
	p := Person{"mike", 'm', 18}
	p.PrintInfo()

	//定义一个结构体变量
	var p2 Person //或者	p2 := *new(Person)
	(&p2).SetInfo("youyou", 'f', 22)
	p2.PrintInfo()
}