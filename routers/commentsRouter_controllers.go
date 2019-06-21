package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["BookKeeping/controllers:AdjustmentController"] = append(beego.GlobalControllerRouter["BookKeeping/controllers:AdjustmentController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/adjustment`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookKeeping/controllers:AdjustmentController"] = append(beego.GlobalControllerRouter["BookKeeping/controllers:AdjustmentController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/adjustment`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookKeeping/controllers:AdminController"] = append(beego.GlobalControllerRouter["BookKeeping/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/admin/user`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookKeeping/controllers:AdminController"] = append(beego.GlobalControllerRouter["BookKeeping/controllers:AdminController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/admin/user`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookKeeping/controllers:CardController"] = append(beego.GlobalControllerRouter["BookKeeping/controllers:CardController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/card`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookKeeping/controllers:CardController"] = append(beego.GlobalControllerRouter["BookKeeping/controllers:CardController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/card`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookKeeping/controllers:CardController"] = append(beego.GlobalControllerRouter["BookKeeping/controllers:CardController"],
        beego.ControllerComments{
            Method: "ShowCard",
            Router: `/card/add`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookKeeping/controllers:CardController"] = append(beego.GlobalControllerRouter["BookKeeping/controllers:CardController"],
        beego.ControllerComments{
            Method: "AddCard",
            Router: `/card/add`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookKeeping/controllers:CardController"] = append(beego.GlobalControllerRouter["BookKeeping/controllers:CardController"],
        beego.ControllerComments{
            Method: "DeleteCard",
            Router: `/card/delete/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookKeeping/controllers:CardController"] = append(beego.GlobalControllerRouter["BookKeeping/controllers:CardController"],
        beego.ControllerComments{
            Method: "EditCard",
            Router: `/card/edit/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookKeeping/controllers:CardController"] = append(beego.GlobalControllerRouter["BookKeeping/controllers:CardController"],
        beego.ControllerComments{
            Method: "UpdateCard",
            Router: `/card/edit/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookKeeping/controllers:CardLogController"] = append(beego.GlobalControllerRouter["BookKeeping/controllers:CardLogController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/card/log`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookKeeping/controllers:IncomeController"] = append(beego.GlobalControllerRouter["BookKeeping/controllers:IncomeController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/income`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookKeeping/controllers:IncomeController"] = append(beego.GlobalControllerRouter["BookKeeping/controllers:IncomeController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/income`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookKeeping/controllers:IndexController"] = append(beego.GlobalControllerRouter["BookKeeping/controllers:IndexController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/index`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookKeeping/controllers:LoginController"] = append(beego.GlobalControllerRouter["BookKeeping/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookKeeping/controllers:LoginController"] = append(beego.GlobalControllerRouter["BookKeeping/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookKeeping/controllers:LoginController"] = append(beego.GlobalControllerRouter["BookKeeping/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookKeeping/controllers:LoginController"] = append(beego.GlobalControllerRouter["BookKeeping/controllers:LoginController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookKeeping/controllers:MerchantController"] = append(beego.GlobalControllerRouter["BookKeeping/controllers:MerchantController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/merchant`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookKeeping/controllers:MerchantController"] = append(beego.GlobalControllerRouter["BookKeeping/controllers:MerchantController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/merchant`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookKeeping/controllers:MerchantLogController"] = append(beego.GlobalControllerRouter["BookKeeping/controllers:MerchantLogController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/merchant/log`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookKeeping/controllers:MerchantLogController"] = append(beego.GlobalControllerRouter["BookKeeping/controllers:MerchantLogController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/merchant/log`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookKeeping/controllers:OutcomeController"] = append(beego.GlobalControllerRouter["BookKeeping/controllers:OutcomeController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/outcome`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookKeeping/controllers:OutcomeController"] = append(beego.GlobalControllerRouter["BookKeeping/controllers:OutcomeController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/outcome`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookKeeping/controllers:PageController"] = append(beego.GlobalControllerRouter["BookKeeping/controllers:PageController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/page`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookKeeping/controllers:StatisticsBalanceController"] = append(beego.GlobalControllerRouter["BookKeeping/controllers:StatisticsBalanceController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/statistics/balance`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookKeeping/controllers:StatisticsBalanceController"] = append(beego.GlobalControllerRouter["BookKeeping/controllers:StatisticsBalanceController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/statistics/balance`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookKeeping/controllers:StatisticsBalanceDtailedController"] = append(beego.GlobalControllerRouter["BookKeeping/controllers:StatisticsBalanceDtailedController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/statistics/balance/detailed`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookKeeping/controllers:StatisticsBalanceDtailedController"] = append(beego.GlobalControllerRouter["BookKeeping/controllers:StatisticsBalanceDtailedController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/statistics/balance/detailed`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookKeeping/controllers:StatisticsMerchantController"] = append(beego.GlobalControllerRouter["BookKeeping/controllers:StatisticsMerchantController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/statistics/servicecharge`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookKeeping/controllers:StatisticsMerchantController"] = append(beego.GlobalControllerRouter["BookKeeping/controllers:StatisticsMerchantController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/statistics/servicecharge`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookKeeping/controllers:StatisticsServiceChargeController"] = append(beego.GlobalControllerRouter["BookKeeping/controllers:StatisticsServiceChargeController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/statistics/merchant`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["BookKeeping/controllers:StatisticsServiceChargeController"] = append(beego.GlobalControllerRouter["BookKeeping/controllers:StatisticsServiceChargeController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/statistics/merchant`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
