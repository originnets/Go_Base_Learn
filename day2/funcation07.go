package main

import "fmt"

//函数类型


func Add(a, b int) int {
	return a + b
}

func Minus(a, b int) int {
	return  a - b
}

// 函数也是一种数据类型，通过type各一个函数类型取一个别名
//FuncType 它是一个函数类型

type FuncType func(int, int) int //没有函数名字， 没有{}

func main() {
	result := Add(1, 1) 	//传统方式调用
	fmt.Println("result = ", result)

	//声明一个函数类型的变量， 取变量名为FuncTest
	var FuncTest FuncType
	FuncTest = Add 				//是变量就可以赋值
	result2 := FuncTest(10, 20) 	//等价于Add(10, 20)
	fmt.Println("result2 = ", result2)

	FuncTest = Minus 			//是变量就可以赋值
	result3 := FuncTest(10, 5) 	//等价于Minus(10, 5)
	fmt.Println("result3 = ", result3)
}
