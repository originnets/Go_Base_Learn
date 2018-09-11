package main

import "fmt"

//此通道只能写, 不能读
func producer(out chan <- int) {
	for i:=0 ;i<10 ;i++  {
		out <- i * i	//写入channel
	}
	close(out)	//关闭channel
}

//此通道只能读, 不能写
func consumer(in <- chan int){
	for num := range in{	//读取channel数据
		fmt.Println("num = ", num)
	}
}

func main()  {
	// 创建一个双向通道
	ch := make(chan int)


	// 新建一个协程
	// 生产者, 生产数字, 写入channel
	go producer(ch)	//channel传参,引用传递

	//主协程
	// 消费者, 从channel读取内容,打印
	consumer(ch)
}
