package main

import "fmt"

//数组做函数参数,它是值传递
//实参数组的每个元素给形参数组拷贝一份
//形参的实在是实参数组的复制品, 是值传递,形参无法改变实参

func modify(a [5]int)  {
	a[0] = 666
	fmt.Println("modify: a = ", a)
}

func main()  {

	a := [5]int{1, 2, 3, 4, 5} //初始化数组
	modify(a) //数组传递过去
	fmt.Println("mian: a = ", a)
}

//输出结果
//modify: a =  [666 2 3 4 5]
//mian: a =  [1 2 3 4 5]