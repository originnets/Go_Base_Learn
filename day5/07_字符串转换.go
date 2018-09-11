package main

import (
	"fmt"
	"strconv"
)


//转换为字符串后追加到字节数组
func Append1() {
	slice := make([]byte, 0 , 1024)	//创建一个byte类型的切片,长度为0, cap为1024
	slice = strconv.AppendBool(slice, true)	//先将布尔值转换为byte再追加到切片中
	slice = strconv.AppendInt(slice, 1234, 10) //先将1234以10进制的方式转换为byte类型再追加到切片中, 1234 为要追加的数, 10表示指定为10进制的方式追加
	slice = strconv.AppendQuote(slice, "abcgohello")  //将"abcgohello"转化为byte类型后 追加到切片中

	fmt.Println("slice = ", string(slice)) // 转换成string后再打印(方便查看)
}


//其他类型转换为字符串
func Format1() {
	var str string
	str = strconv.FormatBool(false)	//把布尔值转换为字符串
	fmt.Printf("str = %s[%T]\n" , str, str)

	str = strconv.FormatFloat(3.14, 'f', -1, 64) //'f'指打印格式, 以小数方式, -1只小鼠点位数(紧缩模式), 64 以float64处理 //浮点型转换为字符
	fmt.Printf("str = %s[%T]\n" , str, str)

}

//整型转字符型,常用
func Itoa1() {
	str := strconv.Itoa(666)
	fmt.Printf("str = %s[%T]\n" , str, str)
}


//字符串转其他类型
func Parse1() {
	var flag bool
	var err error
	flag, err = strconv.ParseBool("false")  //将字符串转换为bool值 ,当字符串不能转换时返回err, bool一般为true和false
	if err == nil {
		fmt.Println("flage = ", flag)
	} else {
		fmt.Println("err = ", err)
	}
}

//把字符串转换为整型
func Atoi1() {
	a, err := strconv.Atoi("567")		//将符串转转换为整型, 当不能转换时返回err
	if err == nil {
		fmt.Printf("a = %d[%T]", a, a)
	} else {
		fmt.Println("err = ", err)
	}

	//if a, err := strconv.Atoi("567"); err == nil {
	//		fmt.Printf("a = %d[%T]", a, a)
	//	} else {
	//		fmt.Println("err = ", err)
	//	}
	//a, _ := strconv.Atoi("567")		//丢弃err值
	//fmt.Printf("a = %d[%T]", a, a)
}

func main() {
	Append1()
	Format1()
	Itoa1()
	Parse1()
	Atoi1()
}
