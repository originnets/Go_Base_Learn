package main

import (
	"fmt"
	"time"
)

func main() {
	//创建一个无缓存的channel
	ch := make(chan int, 0) //后面0或者不写 表示无缓冲

	//len(ch)缓冲区剩余数据个数, cap(ch)缓冲区大小, 因为无缓冲区,所以len(ch)和cap(ch))一直为0
	fmt.Printf("len(ch) = %d, cap(ch) = %d\n", len(ch), cap(ch))

	// 新建一个协程
	go func() {
		for i:=0 ;i<3 ;i++  {
			fmt.Printf("子协程: i=%d\n", i)
			ch <- i	//往chan写内容
		}
	}()

	// 延时
	time.Sleep(2* time.Second)

	for i:=0 ;i<3 ; i++{
		num := <- ch //读取管道中的内容,没有内容前 阻塞
		fmt.Println("num =", num)
	}
}
