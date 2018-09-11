package main

import "fmt"

// defer和匿名函数结合使用


func main()  {
	// defer和匿名函数无传参
	//a := 10
	//b := 20
	//defer func() {
	//	fmt.Printf("内部:a = %d, b = %d\n",a ,b)
	//}() //()代表此匿名函数
	//a = 111
	//b = 222
	//fmt.Printf("外部:a = %d, b = %d\n",a ,b)

//结果为:
//外部:a = 111, b = 222
//内部:a = 111, b = 222

	// defer和匿名函数有传参
	a := 10
	b := 20
	defer func(a, b int) {
		fmt.Printf("内部:a = %d, b = %d\n",a ,b)
	}(a, b) //()代表此匿名函数, 把参数传递过去， 参数已经先传递过去，只是没有调用,所以之还是为a = 10 ，b = 20
	a = 111
	b = 222
	fmt.Printf("外部:a = %d, b = %d\n",a ,b)

//结果为:
//外部:a = 111, b = 222
//内部:a = 10, b = 20
}
