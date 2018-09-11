package main

import (
	"fmt"
)

type Person struct {
	name string
	sex byte
	age int
}

//Person类型, 实现了一个方法

func (temp *Person) PrintInfo() {
	fmt.Printf("name = %s, sex = %c , age= %d \n", temp.name, temp.sex, temp.age)
}


//有一个学生,继承Persion字段,成员和方法读继承了
type Student struct {
	Person //匿名字段
	id	int
	addr string
}

// Student也实现了一个方法 , 这个方法和Person方法同名,这种方法叫重写

func (tmp *Student) PrintInfo () {
	fmt.Println("Student : tmp = ", tmp)
}


func main() {
	s := Student{Person{"mirk", 'm', 28},666, "bj"}

	// 就近原则 : 先找到本作用域的方法, 早不到再用继承的方法
	s.PrintInfo() //这里的PrintInfo方法是使用了Student中Person里的方法

	// 显示调用继承的方法
	s.Person.PrintInfo() //这里的PrintInfo方法是使用了指定的Person里的方法
}

