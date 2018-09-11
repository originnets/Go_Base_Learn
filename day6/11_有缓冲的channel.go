package main

import (
	"fmt"
)

func main() {
	//创建一个有缓存的channel,channel就是异步的
	ch := make(chan int, 3) //代表可以存放三个值

	//len(ch)缓冲区剩余数据个数, cap(ch)缓冲区大小
	fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch))

	// 新建一个协程
	go func() {
		for i:=0 ;i<10 ;i++  {
			ch <- i	//往chan写内容,当ch存放满的时候阻塞, 这里就是3个
			fmt.Printf("子协程[%d]: len(ch) = %d, cap(ch) = %d\n", i, len(ch), cap(ch))
		}
	}()

	// 延时
	//time.Sleep(1 * time.Second)

	for i:=0 ;i<10 ; i++{
		num := <- ch //读取管道中的内容,没有内容前 阻塞
		fmt.Println("num =", num)
	}
}
