package main

import "fmt"

type Test struct {
	id int
	name string
	sex byte
}

func main()  {
	s := make([]interface{}, 3)
	s[0] = 1
	s[1] = "测试"
	s[2] = Test{2,"aaa", 'c'}

	for index, data := range s {
		switch value := data.(type) {
		case int:
			fmt.Printf("s[%d] = %d\n", index, value)
		case string:
			fmt.Printf("s[%d] = %s\n", index, value)
		case Test:
			fmt.Printf("s[%d], id = %d, name = %s , sex = %c \n", index, value.id , value.name, value.sex)
		}
	}
	for index, data := range s {
		if value, ok := data.(int) ; ok ==true {
			fmt.Printf("s[%d] = %d\n", index, value)
		}
		if value, ok := data.(string) ; ok ==true {
			fmt.Printf("s[%d] = %s\n", index, value)
		}
		if value, ok := data.(Test) ; ok ==true {
			fmt.Printf("s[%d], id = %d, name = %s , sex = %c \n", index, value.id , value.name, value.sex)
		}
	}
}
