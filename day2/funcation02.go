package main

import "fmt"


//无参有返回值， 有返回值的函数需要通过return中断函数， 通过return返回
// 后面的int表示返回一个int类型
func MyFunc1() int {
	return 666
}

// 给返回值起一个变量名 , go推荐写法
func MyFunc2() (result int) {
	return 666
}

//常用写法
func MyFunc3() (result int) {
	result =666
	return // 这里返回就是返回result
}


func main() {
	//无参返回值函数调用
	var a int
	a = MyFunc1()
	fmt.Println("a = " , a)

	b := MyFunc1()
	fmt.Println("b = " , b)

	c := MyFunc2()
	fmt.Println("c = " , c)

	d := MyFunc3()
	fmt.Println("d = " , d)
}
