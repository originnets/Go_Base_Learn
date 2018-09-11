package main

import (
	"fmt"
	"time"
)

//全局变量,创建一个channel
var ch = make(chan int)


//打印机属于公共资源
func Printer(str string)  {
	for _, data := range str{
		fmt.Printf("%c", data)
		time.Sleep(time.Second)	//延时一秒
	}
	fmt.Printf("\n")
}

// person1执行完后, 才能到person2执行

func Person1() {
	Printer("Hello")
	ch <- 666	//给管道写数据, 发送数据
}

func Person2() {
	s,ok := <- ch	//从管带取数据,接收, 如果通道没有数据就会阻塞
	Printer("World")
	fmt.Println(s, ok)
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
