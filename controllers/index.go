package controllers

import "github.com/astaxie/beego"

type IndexController struct {
	beego.Controller
}

// @router /index [get]
func (c *IndexController) Get() {
	c.Data["hello"] = "hello world"
	c.TplName = "index.html"
}
