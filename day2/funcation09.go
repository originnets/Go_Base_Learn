package main

import "fmt"

//匿名函数


func main() {
	a := 10
	str := "make"
	// 匿名函数，没有函数名字, 函数定义， 还没有调用

	f1 := func() {
		fmt.Println("a = ", a)
		fmt.Println("str = ",str)
	}

	f1() // 调用匿名函数

	// 给一个函数类型取别名

	type FuncType func() // 函数没有参数， 没有返回值
	//声明变量
	var f2 FuncType
	f2 = f1
	f2()

	//定义匿名函数，同时调用
	func(){
		fmt.Printf("a = %d, str = %s\n",a ,str)
	}() //后面的()代表调用自匿名函数，传参数


	//定义匿名函数，同时调用， 带参数无返回值
	func(i, j int){
		fmt.Printf("i = %d, j = %d\n",i ,j)
	}(10, 2)

	//定义匿名函数，同时调用， 带参数有返回值

	x, y := func(i, j int) (max, min int) {
		if i > j {
			max = i
			min = j
		}else {
			max = j
			min = i
		}
		return //带返回值一定要return
	}(10, 20)

	fmt.Printf("x = %d, y = %d\n",x ,y)
}