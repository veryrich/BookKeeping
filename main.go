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

func main() {

	// 开启session
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider = "memory"
	beego.BConfig.WebConfig.Session.SessionName = "BookKeeping"

	beego.Run()

}
