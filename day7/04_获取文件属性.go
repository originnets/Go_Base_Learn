package main

import (
	"fmt"
	"os"
)

func main()  {
	list := os.Args

	if len(list) != 2 {
		fmt.Println("useage: xxx file")
		return
	}

	FileName := list[1]

	info, err := os.Stat(FileName)

	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	fmt.Println("name = ", info.Name())
	fmt.Println("size", info.Size())
}