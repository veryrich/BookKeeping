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

}
