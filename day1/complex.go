package main

import "fmt"

func main () {
	var t complex128
	t = 2.1 + 3.14i

	fmt.Println("real(t)= ", real(t), "imag(t)", imag(t))
}