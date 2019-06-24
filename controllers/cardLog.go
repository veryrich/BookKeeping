package controllers

import (
	"BookKeeping/models"
	"github.com/astaxie/beego"
	"math"
)

type CardLogController struct {
	beego.Controller
}

// @router /card/log [get]
func (this *CardLogController) Get() {

	//确定每页显示数
	pageSize := 10

	//获取页码
	pageIndex, err := this.GetInt("pageIndex")
	if err != nil {
		pageIndex = int(1)
	}

	//确定数据的起始位置
	start := (pageIndex - 1) * pageSize

	//从db中获取数据
	cards, count := models.ListCardLogs(pageSize, start)

	//获取总页数
	pageCount := math.Ceil(float64(count) / float64(pageSize))

	this.Data["cards"] = cards
	this.Data["count"] = count
	this.Data["pageCount"] = int(pageCount)
	this.Data["pageIndex"] = pageIndex

	this.TplName = "cardLog.html"
}
