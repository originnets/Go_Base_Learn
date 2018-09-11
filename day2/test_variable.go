package main

import "fmt"

// 不同作用域同名变量
// 不同作用域， 语序定义同名变量
// 使用变量的原则， 就近原则

var a byte // 全局变量


func test() {
	fmt.Printf("3: %T\n ", a) // 这里没有定义a 调用的是全局变量中的a
}

func main() {
	var a int // 局部变量

	fmt.Printf("1: %T\n ", a)
	{
		var a float32
		fmt.Printf("2: %T\n ", a)
	}
	test() // 调用 test()函数
}
