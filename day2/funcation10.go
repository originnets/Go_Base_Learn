package main

import "fmt"

//闭包捕获外部变量的特点

func main() {
	a := 10
	str := "make"
	func() {
		//闭包以引用的方式捕获外部变量(里面更改外面也会更改)
		a = 666
		str = "go"
		fmt.Printf("内部: a = %d, str = %s\n", a, str)
	}() //()代表直接调用
	fmt.Printf("外部: a = %d, str = %s\n", a, str)
}