package main

import "fmt"

// defer的执行顺序， 先进后出， 后进先出


func test(x int) {
	result := 100 / x
	fmt.Println("result = ", result)
}
func main()  {

	//fmt.Println("aaaaaaaaaaaaaaaaaaaa")
	//
	//fmt.Println("bbbbbbbbbbbbbbbbbbbb")
	//// 调用一个函数， 导致内存出问题 , 后面的函数就不能执行
	//test(0)
	//fmt.Println("cccccccccccccccccccc")


	defer fmt.Println("aaaaaaaaaaaaaaaaaaaa")

	defer  fmt.Println("bbbbbbbbbbbbbbbbbbbb")
	// 利用defer调用一个函数， 导致内存出问题 , 后面的函数能执行
	defer  test(0)
	defer  fmt.Println("cccccccccccccccccccc")

}
