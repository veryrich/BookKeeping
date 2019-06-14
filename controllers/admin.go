package controllers

import (
	"BookKeeping/models"
	"github.com/astaxie/beego"
	"html/template"
)

type AdminController struct {
	beego.Controller
}

// @router /admin/user [get]
func (this *AdminController) Get() {
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())

	this.TplName = "admin/user.html"
}

// @router /admin/user [post]
func (this *AdminController) Post() {
	this.Data["xsrfdata"] = template.HTML(this.XSRFFormHTML())

	// 获取用户数据
	username := this.GetString("username")
	password := this.GetString("password")
	passwordConfirm := this.GetString("confirmPassword")
	isAdmin := this.GetString("isAdmin")

	//开始创建
	if len(username) >= 5 && len(password) >= 5 {
		if password == passwordConfirm {
			res := models.CreateUser(username, password, isAdmin)
			if !res {
				this.Data["status"] = "创建失败，用户名不能重复"

			} else {
				this.Data["status"] = "创建成功"

			}
		} else {
			this.Data["status"] = "非法输入,绕过前端校验"

		}
	} else {
		this.Data["status"] = "非法输入,绕过前端校验"
	}

	this.TplName = "admin/user.html"
}
