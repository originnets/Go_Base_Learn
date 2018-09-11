package main

import (
	"encoding/json"
	"fmt"
)

type IT struct {
	Company string	`json:"company"`
	Subjects []string	`json:"subjects"`
	Isok bool		`json:"isok"`
	Price float64	`json:"price"`
}


type IT2 struct {
	Subjects []string	`json:"subjects"`
}

func main() {
	jsonbuf := `
	{
		"company": "itcast",
		"isok": true,
		"price": 6666.666,
		"subjects": [
		"GO",
			"C++",
			"Python",
			"Test"
		]
	}`
	var tmp IT // 定义一个结构体

	err := json.Unmarshal([]byte(jsonbuf), &tmp) //第一个是切片类型, 西药先转换为切片类型,第二个参数要地址传递

	if err != nil {
		fmt.Println("err =", err)
		return
	}
	fmt.Println("tmp =", tmp)

	var tmp2 IT2 // 定义一个结构体

	err = json.Unmarshal([]byte(jsonbuf), &tmp2) //第一个是切片类型, 西药先转换为切片类型,第二个参数要地址传递

	if err != nil {
		fmt.Println("err =", err)
		return
	}
	fmt.Println("tmp2 =", tmp2)
}