package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func MyDive(a, b int) (result int, err error)  {
	err = nil
	if b == 0 {
		err = errors.New("除数为0") //返回错误
	} else {
		result = a / b
	}
	return
}


func main()  {
	result, err := MyDive(10, 0)
	if err != nil {
		fmt.Println("err = ", err)
	} else {
		fmt.Println("result = ", result)
	}

	//和下面的代码一致

	//if result, err := MyDive(10, 0); err != nil {
	//	fmt.Println("err = ", err)
	//} else {
	//	fmt.Println("result = ", result)
	//}

}
