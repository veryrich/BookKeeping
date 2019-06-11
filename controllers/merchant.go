package controllers

import "github.com/astaxie/beego"

type MerchantController struct {
	beego.Controller
}

// @router /merchant [get]
func (this *MerchantController) Get() {
	this.TplName = "merchant.html"
}

// @router /merchant [post]
func (this *MerchantController) Post() {
	this.TplName = "merchant.html"
}
