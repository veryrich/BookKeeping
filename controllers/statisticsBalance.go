package controllers

import "github.com/astaxie/beego"

type StatisticsBalanceController struct {
	beego.Controller
}

// @router /statistics/balance [get]
func (this *StatisticsBalanceController) Get() {
	this.TplName = "statistics/balance.html"
}

// @router /statistics/balance [post]
func (this *StatisticsBalanceController) Post() {
	this.TplName = "statistics/balance.html"
}
