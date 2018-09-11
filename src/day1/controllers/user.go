package controllers

import "github.com/astaxie/beego"

type UserController struct {
	beego.Controller
}

func (this *UserController) Get(){
	this.Ctx.WriteString("hello")
}


func (this *UserController) GetInfo() {
	id := this.Ctx.Input.Param(":id")	//获取url的传值
	this.Ctx.WriteString("getInfo data, id = " + id)
}


func (this *UserController) GetFileName() {
	//path := this.Ctx.Input.Param(":path")
	//ext := this.Ctx.Input.Param(":ext")
	//this.Ctx.WriteString("FileName = " + path + "." + ext)
	splat := this.Ctx.Input.Param(":splat")
	this.Ctx.WriteString("FileName = " + splat)
}

func (this *UserController) PostData() {
	this.Ctx.WriteString("this is post func")
}