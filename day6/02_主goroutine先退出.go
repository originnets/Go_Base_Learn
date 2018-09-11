package main

import (
	"fmt"
	"runtime"
	"time"
)

//主协程先退出了, 其他子线程也要跟着退出

func NewTask() {
	for {
		fmt.Println("this is a newtask")
		time.Sleep(time.Second) // 延时1秒
	}
}

func main() {

	//子协程1
	go NewTask()
	//子协程2
	go func() {
		i := 0
		for {
			i++
			fmt.Println("子协程 i = ", i)
			time.Sleep(time.Second)

		}
	}()

	// 主协程
	i := 0
	for {
		i++
		fmt.Println("主协程 i = ", i)
		time.Sleep(time.Second)

		if i == 2 {
			break
		}
	}
}
