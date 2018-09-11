package main

import (
	"fmt"
	"strings"
)

//Contains 是否包含这个字符串
func Contains1() {
	//查看"hellogo"是否包含"hello", 包含返回true, 不包含返回false
	fmt.Println(strings.Contains("hellogo", "hello"))
	fmt.Println(strings.Contains("hellogo", "abc"))
}

//Join 拼接
func Joins1() {
	s := []string{"abc", "hello", "make", "go"}
	buf := strings.Join(s, "x")		//将切片s用x进行拼接在一起组合成新的字符串
	buf2 := strings.Join(s, " ")	//将切片空格用x进行拼接在一起组合成新的字符串
	fmt.Printf("buf = %s , buf2 = %s \n", buf, buf2)

}

//Index	查找子串所在的位置,不存在返回-1
func Index1()  {
	fmt.Println(strings.Index("abcdhello", "hello")) //查找子串所在的位置
	fmt.Println(strings.Index("abcdhello", "go"))	//当不存在时返回-1
}


// Repeat重复字串多少次, 最后返回重复的字符串
func Repeat1() {
	buf := strings.Repeat("go", 3) //重复"go" 3次, 并返回一个字符串
	fmt.Println("buf = ", buf)
}

//Replace 在字符攒中,把旧的字符串替换为新的字符串, n 表示替换的次数, 小于0表示全部替换,返回一个字符串
func Replace1() {
	fmt.Println(strings.Replace("oink oink oink", "k", "key", 2)) // 把"k"字符串替换为"key"字符串,并替换2次
	fmt.Println(strings.Replace("oink oink oink", "nk", "moo", -1)) // 把"nk"字符串替换为"moo"字符串,并全部替换
}

//Split 以指定的分隔符分隔, 返回一个切片(slice)
func Split1() {
	buf := "hello@abc@go@mike"
	s2 := strings.Split(buf,"@") //以@进行分割,返回一个切片(slice)
	fmt.Println("s2 = ", s2)
}

//Trim 在字符串的头部和尾部去除指定的字符串
func Trim1() {
	buf := strings.Trim("   are you ok?   ", " ") //去除头部和尾部的空格,返回一个字符串
	fmt.Printf("buf = %s\n", buf)
}

// Fields 去除字符串头部和尾部的空格,并放到一个切片中(slice)
func Fields1()  {
	buf := strings.Fields("   are you ok?   ") //去除头部和尾部的空格并放到一个切片中(slice)
	fmt.Printf("buf = %s\n", buf)
	for i, data := range buf {
		fmt.Println(i, ",", data)
	}
}


func main() {
	Contains1()
	Joins1()
	Index1()
	Repeat1()
	Replace1()
	Split1()
	Trim1()
	Fields1()
}