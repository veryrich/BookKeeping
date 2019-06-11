package controllers

import "github.com/astaxie/beego"

type OutcomeController struct {
	beego.Controller
}

// @router /outcome [get]
func (this *OutcomeController) Get() {
	this.TplName = "outcome.html"
}

// @router /outcome [post]
func (this *OutcomeController) Post() {
	this.TplName = "outcome.html"
}