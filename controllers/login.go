package controllers

import (
	"BookKeeping/models"
	"github.com/astaxie/beego"
	"html/template"
	"time"
)

type LoginController struct {
	beego.Controller
}

const SESSION_USER_KEY = "BookKeeping"

// @Failure 403
// @router / [get]
// @router /login [get]
func (this *LoginController) Get() {
	// 默认访问登录页面
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	this.TplName = "login.html"
}

// @router /login [post]
func (this *LoginController) Post() {

	// login页面 用户登录
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())
	user := this.GetString("username")
	pwd := this.GetString("password")

	userIfLogin := models.Login(user, pwd)

	if userIfLogin {
		// 登录成功设置cookie
		this.Ctx.SetCookie("username", user, time.Second*3600)

		// 登录成功设置session，登录不成功则返回登录失败
		this.SetSession(SESSION_USER_KEY, user) // 保存session
		this.Ctx.Redirect(302, "/index")
	} else {
		this.Data["status"] = "登录失败! 用户名或密码不正确"
	}

	this.TplName = "login.html"
}

// @router /logout [get]
func (this *LoginController) Logout() {
	// 注销登录
	user := this.GetString("username")
	this.Ctx.SetCookie("username", user, -1)
	this.DelSession(SESSION_USER_KEY)
	this.DestroySession()
	this.Redirect("/login", 302)
}
