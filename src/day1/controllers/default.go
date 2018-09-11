package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "abyun.org"
	c.Data["Email"] = "ssghuo@163.com"
	c.TplName = "index.tpl"
}