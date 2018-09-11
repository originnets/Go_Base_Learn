package main

import "fmt"

// select 的每个case语句里必须是一个IO操作, 如果两个条件都满足的话,就会随机先执行一个

func fiboncci(ch chan <- int, quit <- chan bool)  {
	x,y := 1, 1
	for {
		//监听channel数据的流动
		select {
		case ch <- x:	//把数据写到channel中
			x,y = y, x+y
		case flage := <- quit:
			fmt.Println("flage =", flage)
			return // 退出
		}
	}
}

func main()  {
	ch := make(chan int) // 通讯数字
	quit := make(chan bool)	//程序是否结束


	// 消费者, 从channel读取内容
	go func() {
		for i:=0; i<8; i++ {
			num := <- ch	//读取channel内容
			fmt.Println(num)
		}
		//可以停止了,往quit写一个值
		quit <- true
	}()


	//生产者, 产生数字, 写入channel

	fiboncci(ch, quit)
}
