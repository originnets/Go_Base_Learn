package main

import "fmt"

func testa()  {
	fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaa")
}
func testb(x int)  {
	var a [10]int //定义一个数组
	a[x] = 111 //当x 大于10是,会导致数组越界,产生一个panic,导致程序崩溃 //内部实现的panic

}

func testc()  {
	fmt.Println("cccccccccccccccccccccccccc")
}


func main()  {
	testa()
	testb(11)
	testc()
}