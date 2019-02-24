package controllers

import (
	"github.com/astaxie/beego"
	"github.com/trend-ai/TrendAI_mobile_backend/models"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title Get
// @Description Get current user
// @Success 200 {object} models.UserResponse
// @Failure 401 {object} models.ResponseWithError
// @router / [get]
func (o *UserController) Get() {
	var res models.Response
	defer func() {
		o.Data["json"] = res
		o.ServeJSON()
	}()
	user := o.Ctx.Input.GetData("user").(models.User)
	res = user.ToResponse()
}
