package controllers

import "github.com/astaxie/beego"

type MerchantLogController struct {
	beego.Controller
}

// @router /merchant/log [get]
func (this *MerchantLogController) Get() {
	this.TplName = "merchantLog.html"
}

// @router /merchant/log [post]
func (this *MerchantLogController) Post() {
	this.TplName = "merchantLog.html"
}
