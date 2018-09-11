package main

import "fmt"

func main() {
	ch := make(chan int, 5) //创建一个有缓冲channel

	//新建一个协程
	go func() {
		for i:=0 ;i<20 ;i++  {
			ch <- i // 往channel写数据
		}
		//不需要再写数据时,主动关闭channel
		close(ch)
		//ch <- 666 //关闭channel后无法发数据, 但可以收数据
	}()

	// 取channel中的数据
	for num := range ch {
		fmt.Println("num =", num)
	}
}

func main1() {
	ch := make(chan int) //创建一个无缓冲channel

	//新建一个协程
	go func() {
		for i:=0 ;i<5 ;i++  {
			ch <- i // 往channel写数据
		}
		//不需要再写数据时,主动关闭channel
		close(ch)
		//ch <- 666 //关闭channel后无法发数据, 但可以收数据
	}()


	for {
		//如果ok为true时, 说明channel没有关闭
		if num, ok := <- ch; ok == true {
			fmt.Println("num =", num)
		} else { //channel关闭
			break
		}
	}

}
