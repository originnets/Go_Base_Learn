package controllers

type IndexController struct {
	BaseController
}

// @router / [get]
func (this *IndexController) Index() {
	this.TplName="index.html"
}

// @router /about [get]
func (this *IndexController) About() {
	this.TplName="about.html"
}

// @router /comment [get]
func (this *IndexController) Comment() {
	this.TplName="comment.html"
}

// @router /details [get]
func (this *IndexController) Details() {
	this.TplName="details.html"
}

// @router /message [get]
func (this *IndexController) Message() {
	this.TplName="message.html"
}

// @router /error [get]
func (this *IndexController) Error() {
	this.TplName="404.html"
}