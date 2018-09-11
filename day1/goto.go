package main

import "fmt"

func main()  {
	//break //break is not in a loop(循环), switch, or select
	//continue  //continue is not in a loop(循环)
	// goto 可以在用在任何地方， 但是不能跨函数使用
	fmt.Println("111111111111111111")

	goto END //goto 是关键字， END是用户起的名字， 即标签

	fmt.Println("222222222222222222")

	END: // goto跳转到这个END标签上
		fmt.Println("333333333333333333")
}