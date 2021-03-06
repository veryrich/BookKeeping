package models

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"time"
)

type Cards struct {
	Id            int
	Name          string
	CardNumber    string  `orm:"unique" description:"卡号"`
	Amount        float64 `orm:"digits(12);decimals(2);default(0)" description:"余额"`
	ServiceCharge float64 `orm:"digits(12);decimals(2);default(0)" description:"手续费"`
	Operator      string
	Created       time.Time `orm:"auto_now_add;type(datetime)" description:"创建时间"`
	Updated       time.Time `orm:"auto_now;type(datetime)" description:"修改时间"`
}

// 注册数据库
func init() {
	orm.RegisterModel(new(Cards))
}

// todo: 优化 1.返回文字错误到controllers。 2.利用结构体解析用户给的数据

func ReadOne(cardNumber string) string {
	o := orm.NewOrm()
	card := Cards{CardNumber: cardNumber}
	err := o.Read(&card, "card_number")

	// 记录相关日志
	if err == orm.ErrNoRows {
		logs.Notice("查询不到卡片", err)
	} else if err == orm.ErrMissPK {
		logs.Notice("找不到主键", err)
	} else if err == nil {
		logs.Debug("执行ReadOne查询被删除卡片的属主")
	} else {
		logs.Notice("获取卡片属主失败", err)
	}

	return card.Name
}

func CreateCard(name, cardNumber, operator string) bool {
	// 创建卡片，根据传入值创建
	o := orm.NewOrm()
	card := Cards{Name: name, CardNumber: cardNumber, Operator: operator}
	_, err := o.Insert(&card)

	if err != nil {
		return false
	}

	return true

}

func UpdateCard(oldCardNumber, newName, newCardNumber, operator string) bool {

	o := orm.NewOrm()
	card := Cards{CardNumber: oldCardNumber}
	card.CardNumber = oldCardNumber

	// 先读取一下，确定数据存在后，再进行修改
	if o.Read(&card, "card_number") == nil {
		card.CardNumber = newCardNumber
		card.Name = newName
		card.Operator = operator
		if num, err := o.Update(&card, "card_number", "name", "operator"); err == nil {
			logs.Info(num)
			return true
		}
	} else {
		fmt.Println("未读取到该卡")
	}
	return false
}

// 取出所有卡片，并返回卡片结果集，和数量
func QueryCards(pageSize, start int) (cardRes []*Cards, count int64) {
	o := orm.NewOrm()

	var cards []*Cards
	query := o.QueryTable("cards")
	count, err := query.Count()
	if err != nil {
		logs.Error("获取count失败")
	}

	//查询数据库部分数据
	_, err = query.Limit(pageSize, start).All(&cards)
	if err != nil {
		logs.Error(err)
	}

	return cards, count
}

func DeleteCard(cardNumber string) {
	// 删除卡片
	o := orm.NewOrm()
	card := Cards{CardNumber: cardNumber}
	num, err := o.Delete(&card, "card_number")
	logs.Debug(num, err)

}

func CardFilter(name, cardNumber string) []*Cards {
	// 卡片筛选功能，根据传入值去数据库取数据
	o := orm.NewOrm()

	var cards []*Cards
	query := o.QueryTable("cards")

	switch {
	case name != "" && cardNumber == "":
		_, err := query.Filter("name__contains", name).All(&cards)
		if err != nil {
			logs.Info(err)
		}

	case cardNumber != "" && name == "":
		_, err := query.Filter("card_number", cardNumber).All(&cards)
		if err != nil {
			logs.Info(err)
		}

	case name != "" && cardNumber != "":
		_, err := query.Filter("card_number", cardNumber).Filter("name", name).All(&cards)
		if err != nil {
			logs.Info(err)
		}

	}

	return cards

}
