package main

import "fmt"

func main()  {
	//定义一个变量 , 类型位map[int]string map是无序的
	var m1 map[int]string

	fmt.Println("m1 = ", m1)
	//对于map只用len, 没有cap
	fmt.Println("len = ", len(m1))

	//可以通过make创建map
	m2 := make(map[int]string)
	fmt.Println("m2 = ", m2)
	fmt.Println("len = ", len(m2))

	//可以通过make创建map,可以指定长度, 只是指定了容量, 但是里面没有数据, 这个容量是随着数据的到少自动扩容
	m3 := make(map[int]string, 10)
	m3[1] = "python" //元素的操作
	m3[2] = "go"
	m3[3] = "c++"
	fmt.Println("m3 = ", m3)
	fmt.Println("len = ", len(m3))


	//初始化
	//键值是唯一的

	m4 := map[int]string{1:"python", 2:"go", 3:"c++"}
	fmt.Println("m4 = ", m4)

}
