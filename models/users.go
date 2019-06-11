package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Users struct {
	Id       uint64
	Name     string `orm:"unique" description:"用户名"`
	Password string `orm:"unique" description:"密码"`
	Role     uint8  `orm:"default(1)" description:"用户权限，0为管理员，1为普通用户"`
}

// 注册数据库
func init() {
	orm.RegisterModel(new(Users))
}

func Login(username, pwd string) bool {
	// 对比账号密码

	o := orm.NewOrm()
	user := Users{Name: username, Password: pwd}
	err := o.Read(&user, "Name", "Password")

	if err != nil {
		return false
	}

	return true
}

func main() {
	// 开启session
	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionProvider = "memory"
	beego.BConfig.WebConfig.Session.SessionName = "BookKeeping"

	beego.Run()

}
