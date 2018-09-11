package main

import "fmt"

func testa()  {
	fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaa")
}
func testb()  {
	//fmt.Println("bbbbbbbbbbbbbbbbbbbbbbbbb")
	//显式调用panic函数,导致程序中断
	panic("this is a panic test")
}

func testc()  {
	fmt.Println("cccccccccccccccccccccccccc")
}


func main()  {
	testa()
	testb()
	testc()
}