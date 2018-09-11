package main

import (
	"math/rand"
	"fmt"
	"time"
)

func main()  {
	//设置种子, 只需一次
	//如果种子参数一样 ,那么每次运行程序产生的随机数都一样,想要不一样需要改变种子参数

	//rand.Seed(666)
	rand.Seed(time.Now().UnixNano()) //以当前时间作为种子参数

	for i:=0; i < 100; i++{
		//产生随机数
		//fmt.Printf("rand = %d\n", rand.Int()) //随机很大的数
		fmt.Printf("rand = %d\n", rand.Intn(100)) //限制在100以内的数
	}
}