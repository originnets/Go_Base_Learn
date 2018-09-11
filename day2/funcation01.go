package main

import "fmt"

func MyFunc1(tmp ... int) {
	for _, data := range tmp {
		fmt.Println("data = ", data)
	}
}

func MyFunc2(tmp ... int) {
	for _, data := range tmp {
		fmt.Println("data = ", data)
	}
}




//通过Run进行传参
func Run(args ... int) {
	// 把全部元素传递给MyFunc1
	//MyFunc1(args...)

	//只想传递某些参数给另外的一个函数使用
	MyFunc2(args[:2] ...) // args[0]~args[2](不包括args[2]),传递过去
	//MyFunc2(args[2:] ...) // 从args[2]开始(包括本身),把后面所有的元素传递过去
}

//main --> Run -> other
func main() {
	Run(1, 2, 3, 4)
}
