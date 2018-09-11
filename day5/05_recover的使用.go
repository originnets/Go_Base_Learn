package main

import "fmt"

func testa()  {
	fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaa")
}

func testb(x int)  {
	//设置recover,这个可以恢复panic导致程序崩溃跳过错误,向下执行
	defer func() {
		//recover()
		//fmt.Println(recover()) //可以打印panic的错误信息
		if err := recover(); err != nil {
			fmt.Println(err)		//当他有错误时打印错误
		}
	}() //别忘了(), 这个作用是调用匿名函数

	var a [10]int //定义一个数组
	a[x] = 111 //当x 大于10是,会导致数组越界,产生一个panic,导致程序崩溃 //内部实现的panic

}

func testc()  {
	fmt.Println("cccccccccccccccccccccccccc")
}


func main()  {
	testa()
	testb(9)
	testc()
}
