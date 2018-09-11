package main

import "fmt"

//函数的递归调用

func test(a int) {
	if a == 1 {  // 函数调用自身的终止条件非常重要
		fmt.Println("a = ",a )
		return // 终止函数调用
	}
	//函数调用本身
	test(a - 1)
	fmt.Println("a = ",a)
}


func NewAdd() (sum int) {
	for i := 1; i <=100; i++{
		sum += i
	}
	return
}

func NewAdd1(i int) (sum int) {
	if i == 1 {
		return i
	}
	return i + NewAdd1(i-1)
}

func NewAdd2(i int) (sum int) {
	if i == 100 {
		return i
	}
	return i + NewAdd2(i + 1)
}

func main() {
	test(3)
	fmt.Println("main")
	sum := NewAdd()
	fmt.Println("sum = ", sum)
	sum1 := NewAdd1(100)
	fmt.Println("sum = ", sum1)
	sum2 := NewAdd2(1)
	fmt.Println("sum = ", sum2)
}
