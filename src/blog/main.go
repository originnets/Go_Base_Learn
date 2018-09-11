package main

import (
	_ "blog/routers"
	"github.com/astaxie/beego"
	"strings"
)

func initTemplate() {
	beego.AddFuncMap("equrl", func(x, y string) bool {
		x1 := strings.Trim(x, "/")
		y1 := strings.Trim(y, "/")
		return strings.Compare(x1,y1) == 0
	})
}

func main() {
	initTemplate()
	beego.Run()
}


