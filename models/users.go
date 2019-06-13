package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Users struct {
	Id       uint64
	Name     string    `orm:"unique" description:"用户名"`
	Password string
	Role     uint8     `orm:"default(1)" description:"用户权限，0为管理员，1为普通用户"`
	Created  time.Time `orm:"auto_now_add;type(datetime)" description:"创建时间"`
	Updated  time.Time `orm:"auto_now;type(datetime)" description:"上次登录时间"`
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
