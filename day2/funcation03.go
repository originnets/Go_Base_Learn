package main

import "fmt"


//多个返回值

func MyFunc1() (int, int, int) {
	return 1, 2, 3
}

//go 推荐写法
//func MyFunc2() (a int , b int , c int) {
//	a, b, c = 111, 222, 333
//	return
//}
//或
func MyFunc2() (a, b, c int) {
	a, b, c = 111, 222, 333
	return
}


func main() {
	//a, b, c := MyFunc1()
	//fmt.Printf("%d, %d, %d",a ,b, c)
	a, b, c := MyFunc2()
	fmt.Printf("%d, %d, %d",a ,b, c)

}
