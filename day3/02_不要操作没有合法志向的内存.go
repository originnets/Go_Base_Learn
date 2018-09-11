package main

import (
	"fmt"
)

func main() {
	var p *int
	p = nil
	fmt.Printf("p = %v \n", p)

	//*p = 666 // err, 因为p没有合法指向

	var a int
	p = &a //p指向a
	fmt.Printf("p = %v \n", p)
	*p = 666
	fmt.Println("a = ", a)
}