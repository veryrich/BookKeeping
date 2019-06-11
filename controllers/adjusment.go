package controllers

import "github.com/astaxie/beego"

type AdjustmentController struct {
	beego.Controller
}

// @router /adjustment [get]
func (this *AdjustmentController) Get() {
	this.TplName = "adjustment.html"
}

// @router /adjustment [post]
func (this *AdjustmentController) Post() {
	this.TplName = "adjustment.html"
}
