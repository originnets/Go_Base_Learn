package routers

import (
	"github.com/astaxie/beego"
	"blog/controllers"
)

func init() {
	beego.Include(&controllers.IndexController{})
}
