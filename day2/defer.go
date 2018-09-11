package main

import "fmt"

// defer的使用
func main()  {

	//defer 延迟调用，main函数结束前调用
	defer  fmt.Println("bbbbbbbbbbbbbbbbbbbb")

	fmt.Println("aaaaaaaaaaaaaaaaaaaa")

}
