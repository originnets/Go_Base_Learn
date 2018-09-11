package main

import "fmt"


//全局变量只能使用var定义
// var a int  // 定义变量
// a = 10 		//给变量赋值

//var a int = 10  // 定义变量并赋值
var a = 10		// 使用推导类型进行定义和赋值


// 定义在函数外部的变量就是全局变量
// 全局变量在任何地方都能使用

func main() {
	fmt.Println("a = ", a)
}
