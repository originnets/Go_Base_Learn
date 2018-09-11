package main

import (
	"fmt"
)

func main()  {
	//定义一个数组, [10]int 和[5]int是不同类型
	//[数字], 这个数字作为数组元素个数
	var a [10]int
	var b [5]int
	//查看元素的个数
	fmt.Printf("len(a) = %d , len(b) = %d\n", len(a), len(b))
	// 注意: 定义数组是,指定的数组元素个数必须是常量
	//n := 10
	//var c [n]int //报错non-constant array bound n
	//fmt.Printf("%d", c)
	//操作数组元素, 从0开始,到len()-1,不对称元素,这个数组,叫下标
	//这时的下边,可以时变量或常量
	//赋值
	a[0] = 1
	i := 1
	a[i] = 2 // a[i] = 2

	//给每个元素赋值
	for i := 0; i < len(a); i++{
		a[i] = i+1
	}
	//打印
	//第一个返回下标,第二个返回元素
	for i, data := range a{
		fmt.Printf("a[%d] = %d\n", i, data)
	}
}
