package main

import "fmt"

//回调函数， 函数有一个参数是函数类型， 这个函数就是毁掉函数
//计算器， 可以进行四则运算
//多态， 多种形态， 调用同一个接口， 不同的表现，可以实现不同的表现， 加减乘除


type FuncType func(int, int) int

//实现加法

func Add(a, b int) int {
	return a + b
}

//实现减法
func Minus(a, b int) int {
	return  a - b
}


func Calc(a, b int, FuncTest FuncType) (result int) {
	fmt.Println("Calc")
	result = FuncTest(a, b) // 调用函数，并传参数
	// result = Add(a, b) //Add()必须先定义后，才能调用 , 这样就不能实现多态
	return
}

func main() {
	a := Calc(1, 1 , Add)
	fmt.Println("a = ", a)
	b := Calc(1, 1 , Minus)
	fmt.Println("b = ", b)
}
