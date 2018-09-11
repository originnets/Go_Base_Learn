package main

import (
	"fmt"
	"time"
)

func main1()  {
	// 延时两秒后打印一句话
	timer := time.NewTimer(2* time.Second)
	s := <-timer.C
	fmt.Println("时间到", s)

}

func main2()  {
	// 延时两秒后打印一句话
	time.Sleep(2* time.Second)
	fmt.Println("时间到")

}

func main()  {
	s := <- time.After(2 * time.Second) //定时2s即阻塞2s, 2s后产生一个事件, 往channel写内容
	fmt.Println("时间到",s)
}
