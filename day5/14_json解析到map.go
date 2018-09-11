package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	buf := `
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

	m := make(map[string]interface{}, 4) //定义一个万能字典类型map

	err := json.Unmarshal([]byte(buf), &m)

	if err != nil {
		fmt.Println("err =", err)
		return
	}
	fmt.Println("m =", m)

	//类型断言

	for key, value := range m {
		//fmt.Printf("%v ----> %v\n", key, value)
		switch data := value.(type) {
		case string:
			fmt.Printf("%v ---> %v\n", key, data)
		case []interface{}:   //subjects go用了interface类型
			fmt.Printf("%v ---> %v\n", key, data)
		case bool:
			fmt.Printf("%v ---> %v\n", key, data)
		case float64:
			fmt.Printf("%v ---> %v\n", key, data)
		}
	}
}