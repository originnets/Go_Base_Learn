package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main()  {
	//var err1 error = fmt.Errorf("%s", "this is normol err")
	err1 := fmt.Errorf("%s", "this is normol err")
	fmt.Println("err1 = ", err1)

	err2 := errors.New("this is normol err2")
	fmt.Println("err2 = ", err2)
}
