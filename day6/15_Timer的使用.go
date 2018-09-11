package main

import (
	"fmt"
	"time"
)
//验证time.NewTimer(), 时间到了, 只响应一次
func Verification() {

	// 创建一个定时器, 设置时间为1秒, 1秒后, 就回往time 通道写内容(当前时间)

	timer := time.NewTimer(1 * time.Second)

	for{	// 因为只响应一次, timer.C的内容也只写一次,导致没有数据可读导致死锁
		s := <-timer.C
		fmt.Println("时间到:", s)
	}
}


func main() {

	// 创建一个定时器, 设置时间为2秒, 2秒后, 就回往time 通道写内容(当前时间)

	timer := time.NewTimer(2 * time.Second)

	fmt.Println("当前时间:", time.Now())

	//2s后, 就会往time.C写数去,有数据后,就可以读取
	t := <- timer.C	//channel没有数据前阻塞
	fmt.Println("t = ", t)
	Verification()
}

