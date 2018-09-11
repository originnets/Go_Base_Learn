package main

import (
	//"nginxconf/controllers"
	_ "nginxconf/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
