package main

import "fmt"

func main() {

	s := "王思聪"
	if s == "王思聪"{
		fmt.Println("王思聪")
	} else {
		fmt.Println("上班去")
	}

	if a := 10; a == 10{  // if可以又初始化语句
		fmt.Println("a==10")
	}

	// switch case 的使用
	num := 10
	switch num {
	case 1:
		fmt.Printf("按下的是%d楼\n", num)
		//break //go语言保留了break关键字，跳出switch语言， 默认不写就会跳出
		//fallthrough 不跳出switch语句，后面无条件执行 ，关闭break
	case 2:
		fmt.Printf("按下的是%d楼\n", num)
	case 3:
		fmt.Printf("按下的是%d楼\n", num)
	case 4:
		fmt.Printf("按下的是%d楼\n", num)
	case 5:
		fmt.Printf("按下的是%d楼\n", num)
	default:
		fmt.Printf("按下的是%d楼\n", num)
	}
}
