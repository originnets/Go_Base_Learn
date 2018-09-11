package main

import "fmt"


func xxx(arg ... interface{}) { //这个可以接收任意类型

}

func main()  {
	//空接口,万能类型,可以保存任意类型的值
	var v1 interface{} = 1	//将int类型赋值给interface{}
	fmt.Println("v1 = ", v1)

	var v2 interface{} = "abc"	//将string类型赋值给interface{}
	fmt.Println("v2 = ", v2)

	var v3 interface{} = &v2	//将*interface{}类型赋值给interface{}
	fmt.Println("v3 = ", v3)

	var v4 interface{} = struct {
		X int
	}{1} //将结构体(struct)类型赋值给interface{}

	fmt.Println("v4 = ", v4)

	var v5 interface{} = &struct {
		X int
	}{1}	//将结构体(struct)指针类型赋值给interface{}
	fmt.Println("v5 = ", v5)
}
