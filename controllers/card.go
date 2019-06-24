package controllers

import (
	"BookKeeping/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"html/template"
	"math"
)

type CardController struct {
	beego.Controller
}

// @router /card [get]
func (this *CardController) Get() {

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
	cards, count := models.QueryCards(pageSize, start)

	//获取总页数
	pageCount := math.Ceil(float64(count) / float64(pageSize))

	this.Data["cards"] = cards
	this.Data["count"] = count
	this.Data["pageCount"] = int(pageCount)
	this.Data["pageIndex"] = pageIndex

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

	// 校验数据，并添加卡片
	if operator != nil {

		if name != "" && cardNumber != "" {
			this.Data["status"] = "姓名或卡号不能为空"
		}

		res := models.CreateCard(name, cardNumber, operator.(string))
		if !res {
			this.Data["status"] = "添加失败，卡号重复"
		} else {
			this.Data["status"] = "添加卡片成功"

			// 如果添加卡片成功，记录操作日志
			logRes := models.CardLogging(name, cardNumber, operator.(string), "增加")
			if !logRes {
				logs.Error(logRes)
			}
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

// @router /card/edit/ [get]
func (this *CardController) EditCard() {
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	oldCardNumber := this.Input().Get("cardNumber")
	this.Data["oldCardNumber"] = oldCardNumber
	this.TplName = "cardEdit.html"
}

// @router /card/edit/ [post]
func (this *CardController) UpdateCard() {
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())

	// 准备数据
	newName := this.GetString("name")
	newCardNumber := this.GetString("newCardNumber")
	oldCardNumber := this.GetString("oldCardNumber")
	operator := this.GetSession(SESSION_USER_KEY)

	// 校验数据，并修改卡片
	if operator != nil {
		res := models.UpdateCard(oldCardNumber, newName, newCardNumber, operator.(string))
		if !res {
			this.Data["status"] = "修改失败，卡号不存在，或不安全的操作"
		} else {
			this.Data["status"] = "修改卡片成功"

			// 如果修改卡片成功，记录操作日志
			logRes := models.CardLogging(newName, newCardNumber, operator.(string), "修改")
			if !logRes {
				logs.Error(logRes)
			}
		}
	} else {
		this.Data["status"] = "非法操作，用户未登录"

	}

	this.TplName = "cardEdit.html"
}

// @router /card/delete/ [get]
func (this *CardController) DeleteCard() {

	operator := this.GetSession(SESSION_USER_KEY)

	if operator != nil {
		cardNumber := this.Input().Get("cardNumber")
		name := models.ReadOne(cardNumber)
		models.DeleteCard(cardNumber)

		// 记录日志
		models.CardLogging(name, cardNumber, operator.(string), "删除")

		this.Redirect("/card", 302)
	} else {
		this.Data["status"] = "非法操作，用户未登录"
		this.Redirect("/login", 302)

	}

}
