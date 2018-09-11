package main

import (
	"encoding/json"
	"fmt"
)

/*
{
	"company": "itcast",
	"subjects": [
		"GO",
		"C++",
		"Python",
		"Test"
	],
	"isok": true,
	"price": 666.666
}
*/

//结构成员变量名首字母必须大写
//type IT struct {
//	Company string
//	Subjects []string
//	Isok bool
//	Price float64
//}

type IT struct {
	Company string	`json:"-"` //-代表此字段不会输出到屏幕
	Subjects []string	`json:"subjects"` //二次编码,后面显示小写的subjects
	Isok bool		`json:",string"` //把bool值转化为字符串
	Price float64	`json:"price,string"` //二次编码后转换为字符串
}

func main()  {
	//定义一个结构体变量,同时初始化
	s := IT{"itcast", []string{"GO", "C++", "Python", "Test"}, true, 666.666}

	//编码, 根据内容生成json文本
	//buf, err := json.Marshal(s)
	buf, err := json.MarshalIndent(s, "", " ")// prefix 前缀 indent 缩进


	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	//fmt.Println("buf = ", buf)	// 返回的buf是切片类型
	fmt.Println("buf = ", string(buf)) // 转化为字符串类型
}
