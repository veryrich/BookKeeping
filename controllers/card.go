package controllers

import (
	"BookKeeping/models"
	"github.com/astaxie/beego"
	"html/template"
)

type CardController struct {
	beego.Controller
}

// @router /card [get]
func (this *CardController) Get() {
	this.TplName = "card.html"
}

// @router /card/add [get]
func (this *CardController) ShowCard() {
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	this.TplName = "cardAdd.html"
}

// @router /card/add [post]
func (this *CardController) AddCard() {
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	name := this.GetString("name")
	cardNumber := this.GetString("cardNumber")

	models.CreateCard(name,cardNumber,"test")

	this.TplName = "cardAdd.html"
}

// @router /card [post]
func (this *CardController) Post() {
	this.TplName = "card.html"
}
