package controllers

import "github.com/astaxie/beego"

type CardController struct {
	beego.Controller
}

// @router /card [get]
func (this *CardController) Get() {
	this.TplName = "card.html"
}

// @router /card [post]
func (this *CardController) Post() {
	this.TplName = "card.html"
}
