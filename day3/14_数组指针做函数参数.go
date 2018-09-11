package main

import "fmt"

// a指向实现数组a, 它指向数组,他是数组指针
// *a代表指针所指向的内存, 就是实参a
// 利用指针就可以改变a的值

func modify1(a *[5]int)  {
	(*a)[0] = 666
	fmt.Println("modify: *a = ", *a)
}

func main()  {
	a := [5]int{1, 2, 3, 4, 5} //初始化数组
	modify1(&a) //地址传递
	fmt.Println("mian: a = ", a)
}

//输出结果
//modify: *a =  [666 2 3 4 5]
//mian: a =  [666 2 3 4 5]