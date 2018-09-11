package main

import "fmt"

//函数的定义

//func FuncName(/*参数列表*/) (01 type1, 02 type2 /*返回类型*/){
//	//函数体
//	return v1, v2 //返回多个值
//}

//无参无返回值函数的定义

func MyFunc() {
	a := 6666
	fmt.Printf("a = %d", a)
}

//有参无返回值函数的定义
// 定义函数时， 在函数名后面()定义的参数叫形参
// 参数传递，只能由实参传给形参， 不能反向传递，为单向传递
func MyFunc1(a int) { //普通参数列表,即固定参数传递
	fmt.Printf("a = %d", a)
}


func MyFunc2(a, b int) {
	fmt.Printf("a = %d , b = %d", a, b)
}

func MyFunc3(a int, b string) {
	fmt.Printf("a = %d , b = %s", a, b)
}



func MyFunc4(args ... int) {
	//像 ... int这样的类型， ... type 为不定参数类型

	fmt.Println("len(args) = ", len(args))

	for i, data :=range args {
		fmt.Printf("args[%d]=%d\n", i, data)
	}
	// 或者
	//for i, _ :=range args {
	//	fmt.Printf("args[%d]=%d\n", i, args[i])
	//}
}


func MyFunc5(a string, args ... int) {
	// 注意: 不定参数， 一定(只能)放在形参的最后一个, 固定参数一定要传参
	fmt.Println(a)
	for i, data :=range args {
		fmt.Printf("args[%d]=%d\n", i, data)
	}
}





func main() {
	MyFunc() // 调用MyFunc函数
	// 调用函数传递的参数叫实参
	MyFunc1(666)
	MyFunc2(666, 222)
	MyFunc3(666, "aaaa")
	MyFunc4()
	MyFunc4(1)
	MyFunc4(1, 2, 3)
	MyFunc5("aa", 1,2, 3)

}
