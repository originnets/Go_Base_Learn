package main

import "fmt"

func main()  {
	//a := 10 //定义整形变量a

	var p *int
	//指向一个合法的内存
	//p = &a

	// p 是*int, 指向int类型
	// new() 生成一个随机内存地址
	p = new(int)
	*p = 666
	fmt.Printf("p = %v, *p = %v\n", p, *p)

	q := new(int) // 自推导类型
	*q = 777
	fmt.Println("q = ,*q = ",q ,*q)

}
