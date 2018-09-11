package main

import "fmt"

func main() {
	type long int64  // 把int64取一个别名，后面运用这个别名的就是int64类型
	var a long = 10000000000
	var b int64 = 1000000
	fmt.Printf("%T , %T", a, b)
}
