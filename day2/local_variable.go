package main

import (
	"fmt"
)

func test() {
	a := 10
	fmt.Println("a = ", a)
}

func main() {
	// 定义在 {}里面的变量就是局部变量， 只能在{}里面有效
	// 执行到定义变量时， 才开始分配空间 ， 离开作用域是自动释放
	// 作用域，变量起左右的范围

	// a = 111 在这里直接使用是没有办法编译通过的，因为test中的a 是局部变量

	{
		i := 10
		fmt.Println("i = ", i)
	} // 在{} 里面创建的一个局部变量 在{}外面是没有办法直接在外部使用的
	// i = 111 // 这样是错误的


	if flag := 3; flag == 3 {
		fmt.Println("flag = ", flag)
	}

	//flag = 4 这里也是错误的

}

// go语言是用{}作为分割作用域
