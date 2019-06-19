package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

type CardLog struct {
	Id         int
	Name       string
	CardNumber string
	Operator   string
	Action     string
	Created    time.Time `orm:"auto_now_add;type(datetime)" describe:"创建时间"`
	Updated    time.Time `orm:"auto_now;type(datetime)" describe:"修改时间"`
}

// 注册数据库
func init() {
	orm.RegisterModel(new(CardLog))
}

func CardLogging(name, cardNumber, operator, action string) bool {
	o := orm.NewOrm()
	cardLog := CardLog{Name: name, CardNumber: cardNumber, Operator: operator, Action: action}
	_, err := o.Insert(&cardLog)

	if err != nil {
		return false
	}

	return true
}

func ListCardLogs() []*CardLog {
	o := orm.NewOrm()

	var cards []*CardLog
	query := o.QueryTable("card_log")
	_, e := query.All(&cards)
	if e != nil {
		logs.Error(e)
	}

	return cards
}
