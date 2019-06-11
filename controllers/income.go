package controllers

import "github.com/astaxie/beego"

type IncomeController struct {
	beego.Controller
}

// @router /income [get]
func (this *IncomeController) Get() {
	this.TplName = "income.html"
}

// @router /income [post]
func (this *IncomeController) Post() {
	this.TplName = "income.html"
}
