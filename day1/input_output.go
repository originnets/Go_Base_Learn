package main

import "fmt"

func main()  {
	var num int
	fmt.Printf("请输入:")
	fmt.Scanf("%d", &num) //或者 fmt.Scan(&num)
	fmt.Printf("你输入的是:%d\n", num)

}
