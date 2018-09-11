package main

import (
	"fmt"
	"time"
)

func main()  {
	//创建一个循环定时器,这里就是1秒钟执行一次
	ticker := time.NewTicker(1 * time.Second)

	i := 0
	for {
		<-ticker.C	//设置阻塞, 返回的为时间
		//s := <-ticker.C
		i ++
		fmt.Println("i = ", i)
		if i == 5 {
			ticker.Stop() // 停止这个定时器
			break
		}
	}
}
