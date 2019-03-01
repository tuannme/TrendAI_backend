package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["github.com/trend-ai/TrendAI_mobile_backend/controllers:AdminController"] = append(beego.GlobalControllerRouter["github.com/trend-ai/TrendAI_mobile_backend/controllers:AdminController"],
		beego.ControllerComments{
			Method:           "PushInterestCategories",
			Router:           `/interest_categories`,
			AllowHTTPMethods: []string{"put"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/trend-ai/TrendAI_mobile_backend/controllers:AuthController"] = append(beego.GlobalControllerRouter["github.com/trend-ai/TrendAI_mobile_backend/controllers:AuthController"],
		beego.ControllerComments{
			Method:           "Login",
			Router:           `/login`,
			AllowHTTPMethods: []string{"post"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/trend-ai/TrendAI_mobile_backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/trend-ai/TrendAI_mobile_backend/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Get",
			Router:           `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/trend-ai/TrendAI_mobile_backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/trend-ai/TrendAI_mobile_backend/controllers:UserController"],
		beego.ControllerComments{
			Method:           "Patch",
			Router:           `/`,
			AllowHTTPMethods: []string{"patch"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

	beego.GlobalControllerRouter["github.com/trend-ai/TrendAI_mobile_backend/controllers:UserController"] = append(beego.GlobalControllerRouter["github.com/trend-ai/TrendAI_mobile_backend/controllers:UserController"],
		beego.ControllerComments{
			Method:           "GetCategories",
			Router:           `/categories`,
			AllowHTTPMethods: []string{"get"},
			MethodParams:     param.Make(),
			Filters:          nil,
			Params:           nil})

}
