package controllers

import "github.com/astaxie/beego"

type CardLogController struct {
	beego.Controller
}

// @router /card/log [get]
func (this *CardLogController) Get() {
	this.TplName = "cardLog.html"
}

// @router /card/log [post]
func (this *CardLogController) Post() {
	this.TplName = "cardLog.html"
}