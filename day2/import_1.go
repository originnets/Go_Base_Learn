package main

// 方法1:
//import "fmt"
//import "os"

// 方法2:
//import (
//	"fmt"
//	"os"
//) // 导入包 ， 必须使用， 否则编译不通过
// import . "fmt" //调用函数，无需通过包名

//import io "fmt" // 给包取别名

import _ "fmt"  // 忽略此包 主要是为了引用init函数

func main()  {
//	fmt.Println("this is a test")
//	fmt.Println("os.Agrs = " os.Args)
//	Println("this is a test") //这里不需要写 fmt.Println
//	io.Println("this is a test")
}

