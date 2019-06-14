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
	res := models.ListAllCards()

	this.Data["cards"] = res

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
	operator := this.GetSession(SESSION_USER_KEY)

	if operator != nil {
		res := models.CreateCard(name, cardNumber, operator.(string))
		if !res {
			this.Data["status"] = "添加失败，卡号重复"
		} else {
			this.Data["status"] = "添加卡片成功"
		}
	} else {
		this.Data["status"] = "非法操作，用户未登录"

	}

	this.TplName = "cardAdd.html"
}

// @router /card [post]
func (this *CardController) Post() {
	this.TplName = "card.html"
}
