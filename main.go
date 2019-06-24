package main

import (
	_ "BookKeeping/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// 初始化数据库
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "rich:www.123.com@tcp(192.168.120.78:3306)/BookKeeping?charset=utf8mb4&loc=Asia%2FShanghai")
	orm.RunSyncdb("default", false, true)
	orm.Debug = true

}

// 分页功能  todo: 优化分页功能为 1....9,10,11 or 1,2,3....11
//获取下一页页码
func ShowNextPage(pageIndex, pageCount int) int {

	if pageIndex == pageCount {
		return pageIndex
	}
	return pageIndex + 1
}

//获取上一页页码
func ShowPrePage(pageIndex int) int {
	if pageIndex == 1 {
		return pageIndex
	}
	return pageIndex - 1
}

func main() {
	// 设置日志
	logs.SetLogger(logs.AdapterFile, `{"filename":"BookKeeping.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	logs.Async()

	// 开启session
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider = "memory"
	beego.BConfig.WebConfig.Session.SessionName = "BookKeeping"

	beego.AddFuncMap("next", ShowNextPage)
	beego.AddFuncMap("pre", ShowPrePage)

	beego.Run()

}
