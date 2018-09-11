package main

import (
	"fmt"
	"runtime"
)

func main()  {
	n := runtime.GOMAXPROCS(2)	//指定多少核运行, 这里是1, 表示单核, 返回是当前计算机的核数
	fmt.Println("n = ", n)
	for {
		go fmt.Println(1)

		fmt.Println(0)
	}
}