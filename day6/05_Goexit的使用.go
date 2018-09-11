package main

import (
	"fmt"
	"runtime"
)

func test()  {
	defer fmt.Println("ccccccccccccc")

	//return // 终止此函数
	runtime.Goexit() //终止所在的协程
	fmt.Println("ddddddddddddd")
}

func main()  {

	//创建一个新的协程
	go func() {
		fmt.Println("aaaaaaaaaaaaaa")

		//调用别的函数
		test()
		fmt.Println("bbbbbbbbbbbbbb")
	}()

	//特地写一个死循环, 目的不让主协程结束
	for {}
}
