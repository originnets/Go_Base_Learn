package main

import (
	"fmt"
	"regexp"
)

// ``反引号表示原生字符串

func main()  {
	buf := "abc azc a7c aac 888 a9c tac"
	//1. 解析规则, 它会解析正则表达式, 如果成功返回正则相关对象, 否则返回错误
	//reg1 := regexp.MustCompile(`a.c`) // MustCompile(``) 设置正则规则
	reg1 := regexp.MustCompile(`a[0-9]c`)
	if reg1 == nil {	//规则设置失败, 返回nil
		fmt.Println("err = ", reg1)
		return
	}

	//2. 根据规则提取关键信息
	result1 := reg1.FindAllString(buf, -1)	//小于0表示全部匹配
	fmt.Println("result1= ", result1)

}
