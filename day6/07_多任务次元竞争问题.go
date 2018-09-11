package main

import (
	"fmt"
	"time"
)

//定义一个打印机, 参数为字符串,按每个字符打印
//打印机属于公共资源

func Printer(str string)  {
	for _, data := range str{
		fmt.Printf("%c", data)
		time.Sleep(time.Second)	//延时一秒
	}
	fmt.Printf("\n")
}

func Person1() {
	Printer("Hello")
}

func Person2() {
	Printer("World")
}
func main() {
	//Printer("Hello")
	//Printer("World")
	//新建2个协程, 代表2个人,2个人同时使用打印机
	go Person1()
	go Person2()

	//特地不让主协程结束, 死循环
	for {}
}