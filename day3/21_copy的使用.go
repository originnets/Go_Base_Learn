package main

import "fmt"

func main()  {
	serSlice := []int{1, 2}
	dstSlice := []int{6, 6, 6, 6, 6}
	copy(dstSlice, serSlice) //将serSlice的元素copy到dstSlice中,但是copy后dstSlice的容量不会发生改变

	fmt.Printf("ser=%d , dst=%d", serSlice, dstSlice)
}
