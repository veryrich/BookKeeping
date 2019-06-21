package main

import (
	_ "BookKeeping/routers"
	"github.com/astaxie/beego"
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

// 分页功能
//获取下一页页码
func ShowNextPage(pageIndex, pageCount int) int {

	if pageIndex == pageCount {
		return pageIndex
	}
	return pageIndex + 1
}

// todo: 以后优化掉分页功能
//获取上一页页码
func ShowPrePage(pageIndex int) int {
	if pageIndex == 1 {
		return pageIndex
	}
	return pageIndex - 1
}

func main() {

	// 开启session
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider = "memory"
	beego.BConfig.WebConfig.Session.SessionName = "BookKeeping"


	beego.AddFuncMap("next", ShowNextPage)
	beego.AddFuncMap("pre", ShowPrePage)

	beego.Run()

}
