package main

import (
	"fmt"
	"runtime"
)

func main()  {
	// 定义一个子线程
	go func() {
		for i:=0;i<5;i++{
			fmt.Println("Go")
		}
	}()
	for i:=0;i<2;i++{
		runtime.Gosched() // 让出时间片, 先让其他协程执行完,在回来执行此协程
		fmt.Println("Hello")
	}

}
