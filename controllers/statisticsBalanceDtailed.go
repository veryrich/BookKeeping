package controllers

import "github.com/astaxie/beego"

type StatisticsBalanceDtailedController struct {
	beego.Controller
}


// @router /statistics/balance/detailed [get]
func (this *StatisticsBalanceDtailedController) Get() {
	this.TplName = "statistics/detailed.html"
}

// @router /statistics/balance/detailed [post]
func (this *StatisticsBalanceDtailedController) Post() {
	this.TplName = "statistics/detailed.html"
}