package main

import (
	"math/rand"
	"time"
	"fmt"
)

// 生成随机一个长度为10的切片
func InitData(s []int)  {

	//设置种子
	rand.Seed(time.Now().UnixNano())

	for i:=0; i<len(s); i++ {
		s[i] = rand.Intn(100)  //限制100以内的随机数
	}
}

// 冒泡排序的实现
func BubbleSort(n int, s []int) {

	for i:=0; i < n-1; i++ {
		for j:=0; j<n-1-i; j++{
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

func main()  {
	n := 10
	//创建一个切片, len为n
	s := make([]int, n)

	InitData(s) //初始化函数

	fmt.Println("排序前:", s)

	BubbleSort(n, s) //冒泡排序, 类似以引用传递(实际上是值传递, 但他把处理后的数据返回了)

	fmt.Println("排序后:", s)

}
