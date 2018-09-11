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
	name string //与Person同名
}

func main() {
	//声明(定义)一个变量
	var s Student
	//默认规则(就近原则), 如果能再本作用域,就操作此成员, 如果没有找到,就找继承的字段
	s.name = "mike" //默认操作的是Student的name
	s.sex = 'm'
	s.age = 18
	s.addr = "bj"

	//显示调用
	s.Person.name = "youyou"	//Person的name
	fmt.Printf("s = %+v", s)
}