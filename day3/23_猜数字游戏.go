package main

import (
	"math/rand"
	"time"
	"fmt"
)

func CreatNUM1()(RandNum int) {
	p := new(int) //定义一个指针类型,并生成一个随机内存地址
	// 设置种子
	rand.Seed(time.Now().UnixNano())
	var num int
	for {
		num = rand.Intn(10000)  //一定是4位数
		if num >= 1000 {
			break
		}
	}

	fmt.Println("num = ", num)
	*p = num
	return num
}


func CreatNUM(p *int) {
	// 设置种子
	rand.Seed(time.Now().UnixNano())
	var num int
	for {
		num = rand.Intn(10000)  //一定是4位数
		if num >= 1000 {
			break
		}
	}
	//fmt.Println("num = ", num)
	*p = num
}


func GetNum(s []int, num int) {
	s[0] = num / 1000 //取千位
	s[1] = num % 1000 / 100 //取百位
	s[2] = num % 100 / 10 //取十位
	s[3] = num % 10 //取个位

}

func OneGame(RandSlice []int) {
	var num int
	for{	//没有全部猜对,就一直循环
		for {	//输入的不是四位数时就一直循环
			fmt.Printf("请输入你猜的四位数:")
			fmt.Scan(&num)
			if 999 < num && num < 10000 { // 输入的是四位数时就跳出循环
				break
			} else {
				fmt.Println("你输入的不是四位数,请重新输入:")
			}
		}
		//fmt.Println("num = ", num)
		KeyNum := make([]int, 4)	//定义一个切片接收键盘输入转换位切片的值
		GetNum(KeyNum, num) //把键盘输入的值转换成切片
		//fmt.Println("KeyNum = ", KeyNum)
		n := 0 //定义结束值
		for i:=0; i<4; i++{		//核心,判断随机生成的4位数和输入的值的大小
			if RandSlice[i] > KeyNum[i]{
				fmt.Printf("第%d位猜小了\n", i+1)
			}else if  RandSlice[i] < KeyNum[i]{
				fmt.Printf("第%d位猜大了\n", i+1)
			}else {
				//fmt.Printf("恭喜你猜对了第%d位\n", i+1)
				n ++
			}
		}
		if n == 4{	//当全部正确时就跳出循环
			fmt.Println("恭喜你全部猜对了")
			break // 跳出循环
		}
	}
}

func main() {
	//随机产生一个4位数
	var RandNum int
	//随机产生一个4位数
	CreatNUM(&RandNum)  // 地址传值
	//RandNum := CreatNUM1()
	//fmt.Println("RandNum = ", RandNum)
	RandSlice := make([]int, 4) // 生成一个切片,保存这个4位数的每一位
	GetNum(RandSlice, RandNum)	//将4位数转换成切片
	fmt.Println("RandSlice = ", RandSlice)
	OneGame(RandSlice)
}
