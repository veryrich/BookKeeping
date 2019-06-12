package controllers

import "github.com/astaxie/beego"

type StatisticsMerchantController struct {
	beego.Controller
}

// @router /statistics/merchant [get]
func (this *StatisticsServiceChargeController) Get() {
	this.TplName = "statistics/merchant.html"
}

// @router /statistics/merchant [post]
func (this *StatisticsServiceChargeController) Post() {
	this.TplName = "statistics/merchant.html"
}
