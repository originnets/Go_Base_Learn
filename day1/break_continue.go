package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for {
		i++
		time.Sleep(time.Second) // 延时1s

		if i== 5 {
			//break // 跳出循环，如果嵌套多个循环，跳出最近的循环
			//continue // 跳过本次循环，下一次循环继续执行
		}
		fmt.Println("i = ", i)
	}
}
