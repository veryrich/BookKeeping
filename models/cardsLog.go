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

func ListCardLogs(pageSize, start int) (cardRes []*CardLog, count int64) {
	o := orm.NewOrm()

	var cards []*CardLog
	query := o.QueryTable("card_log")
	count, err := query.Count()
	if err != nil {
		logs.Error("获取count失败")
	}

	//查询数据库部分数据
	_, err = query.Limit(pageSize, start).OrderBy("-created").All(&cards)
	if err != nil {
		logs.Error(err)
	}

	return cards, count

}
