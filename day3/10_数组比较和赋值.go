package main

import "fmt"

func main()  {
	//支持比较, 但是只支持 == 或者 != 比较是不是每个元素都一样 2个数组比较,数组类型要一样
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{1, 2, 3, 4, 5}
	c := [5]int{1, 2, 3}

	fmt.Println(" a == b", a == b )
	fmt.Println(" a == c", a == c )

	//同类型的数组可以赋值,元素个数不一样,表示类型不一样比如 [6]int 和 [5]int表示的类型不一样

	var d[5]int
	d = a // 把a数值赋值给d
	fmt.Println("d = ", d)
}
