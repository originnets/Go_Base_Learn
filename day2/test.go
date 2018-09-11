package main

import "fmt"

type FuncType func(int, int) int

func Add(a, b int) int {
	return a + b
}

func Test(a, b int, FuncTest FuncType) (result int) {
	fmt.Println("Test")
	result = FuncTest(a, b)
	return
}

func main() {
	back := Test(1, 2 ,Add)
	fmt.Println("back = ", back)
}

