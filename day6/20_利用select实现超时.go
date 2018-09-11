package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	quit := make(chan bool)
	//新建一个协程
	go func() {
		for {
			select {
			case num := <- ch :
				fmt.Println("num = ", num)
			case <-time.After(3 * time.Second):	 //设置定时3秒没有操作,自动执行这个, 因为外部3秒没有写入就会自动执行这个
				fmt.Println("超时")
				quit <- true	//向quit channel写一个数据
			}
		}
	}()

	// 主协程

	for i:=0 ;i<5 ;i++  {
		ch <- i	//向ch channel写一个数据
		time.Sleep(time.Second)
	}

	<-quit	//向quit channel读数据
	fmt.Println("程序结束了")
}
