package main

import (
	"os"
	"fmt"
)

//利用os模块获取命令行参数

func main() {
	//接受用户传递的参数，都是以字符串方式传递的
	list := os.Args
	n := len(list)
	fmt.Println("n = ", n)

	//for i := 0; i < n; i ++ {
	//	fmt.Printf("list[%d] = %s \n", i, list[i])
	//}

	for i , data := range list {
		fmt.Printf("list[%d] = %s\n", i, data)
	}

}
