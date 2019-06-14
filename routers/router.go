package routers

import (
	"BookKeeping/controllers"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"strings"
)

func init() {
	beego.Include(&controllers.IndexController{})
	beego.Include(&controllers.LoginController{})
	beego.Include(&controllers.IncomeController{})
	beego.Include(&controllers.OutcomeController{})
	beego.Include(&controllers.AdjustmentController{})
	beego.Include(&controllers.MerchantController{})
	beego.Include(&controllers.CardController{})
	beego.Include(&controllers.StatisticsMerchantController{})
	beego.Include(&controllers.StatisticsServiceChargeController{})
	beego.Include(&controllers.StatisticsBalanceController{})
	beego.Include(&controllers.StatisticsBalanceDtailedController{})
	beego.Include(&controllers.MerchantLogController{})
	beego.Include(&controllers.CardLogController{})
	beego.Include(&controllers.AdminController{})

	// 注册过滤器 todo:为了方便调试，暂时取消登录校验，上线后需要打开
	//beego.InsertFilter("/*", beego.BeforeRouter, FilterUser)

}

// 过滤器定义，未登录用户重定向至登录页面
var FilterUser = func(ctx *context.Context) {
	_, ok := ctx.Input.Session(controllers.SESSION_USER_KEY).(string)
	ok2 := strings.Contains(ctx.Request.RequestURI, "/login")
	if !ok && !ok2 {
		ctx.Redirect(302, "/login")
	}

	fmt.Println(ok, ok2)

}
