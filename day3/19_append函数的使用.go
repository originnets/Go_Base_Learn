package main

import "fmt"

func main()  {
	s1 := []int{} //定义一个空切片
	fmt.Printf("len = %d, cap = %d\n", len(s1), cap(s1))
	fmt.Println("s1 = ", s1)


	//在原切片的末尾添加元素, 容量自动变化,一旦超过原来的容量,一般以2倍的容量扩容
	s1 = append(s1, 1)
	s1 = append(s1, 2, 2 , 2)
	s1 = append(s1, 3)
	fmt.Printf("len = %d, cap = %d\n", len(s1), cap(s1))
	fmt.Println("s1 = ", s1)


	s2 := []int{1,2,3}
	s2 = append(s2, 5, 5, 5)
	s2 = append(s2, 6)
	s2 = append(s2, 7)
	fmt.Printf("len = %d, cap = %d\n", len(s2), cap(s2))
	fmt.Println("s2 = ", s2)
}
