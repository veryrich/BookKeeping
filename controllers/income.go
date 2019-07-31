package controllers

import (
	"BookKeeping/models"
	"fmt"
	"github.com/astaxie/beego"
	"html/template"
)

type IncomeController struct {
	beego.Controller
}

// @router /income [get]
func (this *IncomeController) Get() {
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())

	this.TplName = "income.html"
}

// @router /income [post]
func (this *IncomeController) Post() {
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())

	name := this.GetString("name")
	cardNumber := this.GetString("cardNumber")
	Amount := this.GetString("amount")
	ServiceCharge := this.GetString("serviceCharge")

	fmt.Println(name, cardNumber, Amount, ServiceCharge)

	this.TplName = "income.html"
}

// @router /income/query [get]
func (this *IncomeController) GetCardNumber() {
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())

	name := this.Input().Get("name")
	res := models.CardFilter(name, "")

	this.Data["json"] = res
	this.ServeJSON()
	return
}
