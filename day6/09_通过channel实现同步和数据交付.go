package main

import (
	"fmt"
	"time"
)

func main() {
	//创建channel
	ch := make(chan string)

	go func() {
		defer fmt.Println("子协程调用完毕")

		for i:=0;i<2 ; i++ {
			fmt.Println("i =", i)
			time.Sleep(time.Second)
		}
		ch <- "我是子协程,工作完毕"	//发送数据
	}()
	//str := <- ch // 没有数据前 阻塞
	str, ok := <- ch	//接收数据,没有数据前 阻塞
	fmt.Println( str, ok)
}
