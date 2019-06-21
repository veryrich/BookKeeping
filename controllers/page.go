package controllers

import (
	"BookKeeping/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"math"
)

type PageController struct {
	beego.Controller
}

type Page struct {
	PageNo     int
	PageSize   int
	TotalPage  int
	TotalCount int
	FirstPage  bool
	LastPage   bool
	List       interface{}
}

func PageUtil(count int, pageNo int, pageSize int, list interface{}) Page {
	tp := count / pageSize
	if count%pageSize > 0 {
		tp = count/pageSize + 1
	}
	return Page{PageNo: pageNo, PageSize: pageSize, TotalPage: tp, TotalCount: count, FirstPage: pageNo == 1, LastPage: pageNo == tp, List: list}
}

// @router /page [get]
func (this *PageController) Get() {
	o := orm.NewOrm()
	query := o.QueryTable("cards")
	count, _ := query.Count()

	//确定每页显示数
	pageSize := 5
	//获取总页数
	pageCount := math.Ceil(float64(count) / float64(pageSize))

	//获取页码
	pageIndex, err := this.GetInt("pageIndex")
	if err != nil {
		pageIndex = int(1)
	}

	//确定数据的起始位置
	start := (pageIndex - 1) * pageSize
	//查询数据库部分数据
	var cards []*models.Cards
	query.Limit(pageSize, start).All(&cards)

	//把数据传递给视图，并在视图中显示。
	this.Data["data"] = cards

	this.Data["count"] = count
	this.Data["pageCount"] = int(pageCount)
	this.Data["pageIndex"] = pageIndex

	this.TplName = "page.html"
}

//  todo: 合并分页功能到卡片管理中
