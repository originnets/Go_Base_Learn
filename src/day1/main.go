package main

import (
	_ "day1/routers" //调用routers的init函数
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

