package main

import "fmt"

func main() {
	//for 初始化条件 ; 判断条件 ； 条件变化{
	// } 1+2+3+....100 累加

	var sum int
	sum1 := 1
	//sum2 := 1

	for i := 1; i <=100; i++ {
		sum = sum + i
	}
	fmt.Println(sum)

	for sum1 <= 100 {
		sum1 ++
	}
	fmt.Println(sum1)

	// 无限循环(死循环)
	//for{
	//	sum2++
	//	fmt.Println(sum2)
	//}

	//通过for循环打印每个字符
	str := "abcdefghijk"
	for i := 0; i<len(str);i++ {
		fmt.Printf("str[%d]=%c\n", i,str[i])
	}
	//通过迭代器打印每个元素， 默认返回2个值
	for i, data := range str {
		fmt.Printf("str[%d]=%c\n", i, data)
	}

	for i := range str { //通过迭代器打印每个元素，这样默认忽略第二个
		fmt.Printf("str[%d]=%c\n", i, str[i])
	}

	for i, _ := range str { //通过迭代器打印每个元素，这样默认忽略第二个
		fmt.Printf("str[%d]=%c\n", i, str[i])
	}
}
