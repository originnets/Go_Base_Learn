package main

import "fmt"

func main()  {
	a := []int{1, 2, 3, 4, 5}
	s := a[0:3:5]

	fmt.Println("s = ", s)
	fmt.Println("len(s) = ", len(s)) // 长度 3-0
	fmt.Println("cap(s) = ", cap(s)) // 容量 5-0

	s1 := a[1:4:5]
	fmt.Println("s1 = ", s1)	// 1表示:从下标1开始取,3表示:取到下标为4 但是是左闭右开区间,所以取不到下标为4的值
	fmt.Println("len(s1) = ", len(s1)) // 长度 4-1
	fmt.Println("cap(s1) = ", cap(s1)) // 容量 5-1
}