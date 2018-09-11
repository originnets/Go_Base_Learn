package main

import "fmt"

func swap(a, b *int) { // swap(a, b *int)接受地址传递, 这里的a和b是地址需要加上*才表示值
	*a, *b = *b, *a
	fmt.Printf("a= %d, b= %d\n",*a ,*b)
}

func main() {
	a, b := 10, 20

	// 通过一个函数交换a和b的内容
	swap(&a, &b) //地址传递
	fmt.Printf("a= %d, b= %d",a ,b) // 这边打印值还是已经交换
}
