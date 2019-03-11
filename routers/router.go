// @APIVersion 1.0.0
// @Title TrendAI Backend APIs
// @Description APIs document for TrendAI Backend app
// @Contact zenthangplus@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/trend-ai/TrendAI_mobile_backend/controllers"
	"github.com/trend-ai/TrendAI_mobile_backend/services/authentications"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSBefore(authentications.JwtAuthenticationFilter()),
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/trends",
			beego.NSBefore(authentications.JwtAuthenticationFilter()),
			beego.NSInclude(
				&controllers.TrendsController{},
			),
		),
		beego.NSNamespace("/auth",
			beego.NSInclude(
				&controllers.AuthController{},
			),
		),
		beego.NSNamespace("/admin",
			beego.NSInclude(
				&controllers.AdminController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
