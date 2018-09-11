package main

import "fmt"

//有多少个[]就是多少维
//有多少个[]就要用多少个循环

func main() {

	var a [3][4]int
	k := 0
	for i := 0; i < 3; i++{
		for j := 0; j < 4; j++ {
			k ++
			a[i][j] = k
			fmt.Printf("a[%d][%d] = %d\t", i, j, k )
		}
		fmt.Printf("\n")
	}
	fmt.Println("a = ", a)

	//有3个元素, 每个元素又是一个一维数组[4]int
	//全部初始化
	b := [3][4]int{ {1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println("b = ", b)

	//部分初始化
	c := [3][4]int{ {1, 2, 3, 4}, {5, 6, 7, 8},}
	fmt.Println("c = ", c)

	//部分初始化
	d := [3][4]int{ {1, 2, 3, 4}, {9, 10, 11, 12}}
	fmt.Println("d = ", d)

	//部分初始化 [3][4]int{1: {1, 2, 3, 4}} 中的1:{1, 2, 3, 4}表示在二维数组中的第2个元素插入{1, 2, 3, 4}这个数组
	e := [3][4]int{1: {1, 2, 3, 4}}
	fmt.Println("e = ", e)
}