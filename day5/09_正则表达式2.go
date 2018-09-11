package main

import (
	"fmt"
	"regexp"
)

func main() {
	buf := "43.14 567 agsdg 1.23 7. 8.99 lsdljgq 6.66 7.8  "
	//定义一个正则表达式 + 表示前一个字符一次或多次
	reg := regexp.MustCompile(`\d+\.\d+`)

	if reg == nil {
		fmt.Println("error")
		return
	}

	//提取关键信息

	//result := reg.FindAllString(buf, -1)
	result := reg.FindAllStringSubmatch(buf, -1)	//数组的方式打印
	fmt.Println("result = ", result)
}
