package controllers

import "github.com/astaxie/beego"

type StatisticsServiceChargeController struct {
	beego.Controller
}

// @router /statistics/servicecharge [get]
func (this *StatisticsMerchantController) Get() {
	this.TplName = "statistics/serviceCharge.html"
}

// @router /statistics/servicecharge [post]
func (this *StatisticsMerchantController) Post() {
	this.TplName = "statistics/serviceCharge.html"
}
