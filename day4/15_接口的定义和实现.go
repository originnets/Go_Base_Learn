package main

import "fmt"

//定义接口类型
type Humaner interface {
	//方法,只有声明,没有实现,由别的类型(自定义类型)实现
	SayHi()
}

type Test struct {
	name string
	id int
	sex byte
}

type Student struct {
	name string
	id int
}

type Teacher struct {
	addr string
	group string
}

type MyStr string


//Student 实现此方法
func (tmp *Student) SayHi() {
	fmt.Printf("Student[%s, %d] SayHi\n", tmp.name, tmp.id)
}

//Teacher 实现此方法
func (tmp *Teacher) SayHi() {
	fmt.Printf("Teacher[%s, %s] SayHi\n", tmp.addr, tmp.group)
}

// MyStr 实现此方法

func (tmp *MyStr) SayHi() {
	fmt.Printf("MyStr[%s] SayHi\n", *tmp)
}


//定义一个函数,函数的阐述为接口类型
//只有一个函数, 可以有不同的表现,多态
func WhoSayHi(i Humaner) {
	i.SayHi()
}

func main01()  {
	//定义接口类型变量
	var i Humaner

	// 只要实现了此接口方法的类型,那么这个类型的变量(接收者类型)就可以给i赋值

	s := &Student{"mirk",666}
	i = s
	i.SayHi()

	t := &Teacher{"bj","go"}
	i = t
	i.SayHi()

	var str MyStr = "hello mike"
	i = &str
	i.SayHi()
}

func main()  {
	s := &Student{"maik", 666}
	t := &Teacher{"bj", "go"}
	var str MyStr = "hello mike"
	WhoSayHi(s)
	WhoSayHi(t)
	WhoSayHi(&str)

	//创建一个切片
	x := make([]Humaner, 3)
	x[0] = s
	x[1] = t
	x[2] = &str

	// 第一个返回下标, 第二个返回下标所对应的值
	for _,i := range x {
		i.SayHi()
	}
}