package routers

import (
	"day1/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{}, "get:Get")
    beego.Router("/user/?:id", &controllers.UserController{},"get:GetInfo")
    beego.Router("/download/*.*", &controllers.UserController{}, "get:GetFileName")
	beego.Router("/download/*", &controllers.UserController{}, "get:GetFileName;post:PostData") //同时指定get和post方法
}
