package main

import "fmt"


//有参数，有返回值

func MaxAndMin(a, b int) (max, min int) {
	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}
	return  //有返回值的函数， 必须通过return返回
}

func main() {
	//max, min := MaxAndMin(20, 10)
	//fmt.Printf("Max:%d, min:%d\n",max, min)
	// 通过匿名变量丢弃某个变量
	max, _ := MaxAndMin(20, 10)
	fmt.Printf("Max:%d",max)

}
