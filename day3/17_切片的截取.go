package main

import "fmt"

func main()  {
	array := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//[low:high:max] 取下标从low开始的元素,到high-1的下标元素结束, 长度为:len=high-low, 容量:cap=max-low

	s1 := array[:] //[:]相当于[0:len(array):len(array)] 不指定容量默认就是长度和容量一样

	fmt.Println("s1 = ", s1)
	fmt.Printf("len = %d, cap = %d\n", len(s1), cap(s1))

	//操作摸个元素, 和数组的操作方式一样

	data := array[1]
	fmt.Println("data = ", data)

	s2 := array[3:6:7] //a[3], a[4],a[5] len = 6-3 cap = 7-3
	fmt.Println("s2 = ", s2)
	fmt.Printf("len = %d, cap = %d\n", len(s2), cap(s2))

	s3 := array[:6] //从0开始, 取6个元素,长度是6, 容量也是6, 常用的
	fmt.Println("s3 = ", s3)
	fmt.Printf("len = %d, cap = %d\n", len(s3), cap(s3))

	s4 := array[3:] //从下标为3开始, 到结尾,长度是7 容量也是7
	fmt.Println("s4 = ", s4)
	fmt.Printf("len = %d, cap = %d\n", len(s4), cap(s4))
}
