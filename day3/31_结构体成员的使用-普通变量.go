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
	//定义一个结构体普通变量

	var s Student //或者s := Student{}
	//操作成员, 需要使用(.)运算符
	s.id = 1
	s.name = "mike" //字符串
	s.sex = 'm'	//字符
	s.age = 19
	s.addr = "bj"
	fmt.Println("s = ", s)

}