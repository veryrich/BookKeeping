package models

import (
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

type Cards struct {
	Id         int
	Name       string
	CardNumber string `orm:"unique" describe:"卡号"`
	Operator   string
	Created    time.Time `orm:"auto_now_add;type(datetime)" describe:"创建时间"`
	Updated    time.Time `orm:"auto_now;type(datetime)" describe:"修改时间"`
}

// 注册数据库
func init() {
	orm.RegisterModel(new(Cards))
}

func CreateCard(name, cardNumber, operator string) bool {
	o := orm.NewOrm()
	card := Cards{Name: name, CardNumber: cardNumber, Operator: operator}
	_, err := o.Insert(&card)

	if err != nil {
		return false
	}

	return true

}

func ListAllCards() []*Cards {
	o := orm.NewOrm()

	//var lists []orm.ParamsList
	//
	//num, err := o.Raw("select * from cards").ValuesList(&lists)
	//if err == nil && num > 0 {
	//	return lists
	//}
	var cards []*Cards
	query := o.QueryTable("cards")
	_, e := query.All(&cards)
	if e != nil {
		logs.Error(e)
	}

	return cards

}
