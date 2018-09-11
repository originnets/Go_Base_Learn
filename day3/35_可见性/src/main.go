package main

import (
	"test"
	"fmt"
) //引用test下的包

func main() {

	//包名.函数名
	test.MyFunc() // 调用test中的MyFunc函数
	//test.myFunc() // 该函数无法调用 报错undefined: test.myFunc

	//包名.结构体类型名
	s1 := test.Stu{18}
	fmt.Println(s1)
}
