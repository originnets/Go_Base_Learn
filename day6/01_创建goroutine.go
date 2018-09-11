package main

import (
	"fmt"
	"time"
)

// go 创建go协程

func NewTask() {
	for {
		fmt.Println("this is a newtask")
		time.Sleep(time.Second) // 延时1秒
	}
}

func main()  {
	go NewTask() // 新建一个协程,新建一个任务, 需要放到非协程任务前面

	for {
		fmt.Println("this is a main")
		time.Sleep(time.Second) // 延时1秒
	}

	// 这样的话就是一直执行上面的任务
	// NewTask()
}
