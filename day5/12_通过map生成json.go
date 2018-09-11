package main

import (
	"encoding/json"
	"fmt"
)

func main()  {

	//创建一个map
	m := make(map[string]interface{}, 4)  // 定义一个万能接收interface万能类型

	m["company"] = "itcast"
	m["subjects"] = []string{"GO", "C++", "Python", "Test"}
	m["isok"] = true
	m["price"] = 6666.666

	//编码, 根据内容生成json文本
	//result, err := json.Marshal(m)
	result, err := json.MarshalIndent(m, "", "	")// prefix 前缀 indent 缩进
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	//fmt.Println("result = ", result)	// 返回的buf是切片类型
	fmt.Println("result = ", string(result)) // 转化为字符串类型

}