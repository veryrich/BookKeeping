package controllers

import (
	"BookKeeping/models"
	"github.com/astaxie/beego"
)

type CardLogController struct {
	beego.Controller
}

// @router /card/log [get]
func (this *CardLogController) Get() {
	res := models.ListCardLogs()
	this.Data["cardLogs"] = res

	this.TplName = "cardLog.html"
}
