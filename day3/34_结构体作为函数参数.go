package main

import "fmt"

type Student struct {
	id int
	name string
	sex byte
	age int
	addr string
}



func test01(s Student){
	s.id = 666
	fmt.Println("test01:", s )
}


func test02(s *Student) {
	s.id = 666
	fmt.Println("*test01:", *s )
}

func main()  {
	s := Student{1,"mirk",'m', 18, "bj"}

	test01(s) //这个是值传递, 形参无法改变实参
	fmt.Println("main:", s )
	test02(&s) // 地址传递(引用传递),形参可以改变实参
	fmt.Println("main:", s )
}
