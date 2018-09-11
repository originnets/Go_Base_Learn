package main

import "fmt"

func main()  {
	//const (
	//	a = iota
	//	b = iota
	//	c = iota
	//	d = iota
	//)
	//或者
	const (
		a = iota
		b
		c
		d
	)

	const (
		e = iota
		f = iota
		g = iota
		h = iota
	)
	fmt.Println(a, b, c, d, e, f, g, h) //iota是枚举类型，自动生成当出现一个const时会自动从0开始计数
}
