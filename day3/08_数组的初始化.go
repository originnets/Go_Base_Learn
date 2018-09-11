package main

import "fmt"

func main() {
	//声明定义同时赋值, 叫初始化
	//全部初始化
	var a [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("a = ", a)

	//推导类型初始化
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b = ", b)

	//部分初始化,没有初始化的元素,自动赋值为0
	c := [5]int{1, 2, 3}
	fmt.Println("c = ", c)

	//指定某个元素初始化 [5]int{2:10, 4:20} ,下标为2的赋值为10, 下标为4的赋值为20
	d := [5]int{2:10, 4:20}
	fmt.Println("d = ", d)
}
