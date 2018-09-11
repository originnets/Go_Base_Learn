package main

import "fmt"

func main() {
	m1 := map[int]string{1 : "make", 2 : "youyou" , 3 : "test"}
	fmt.Println("m1 = ", m1)
	// 赋值, 如果已经存在的key值, 再赋值就是修改该内容,如果没有的话就是追加
	m1[3] = "test1"
	fmt.Println("m1 = ", m1)
	// 赋值, 如果key不存在值就是追加一个元素, map底层自动扩容, 和append有点类似
	m1[4] = "go"

}
