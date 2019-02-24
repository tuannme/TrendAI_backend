package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/trend-ai/TrendAI_mobile_backend/models"
	"net/http"
	"time"
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

// @Title Patch
// @Description Update specific fields for current user
// @Success 200 {object} models.UserResponse
// @Failure 400 {object} models.ResponseWithError
// @router / [patch]
func (o *UserController) Patch() {
	var res models.Response
	defer func() {
		o.Data["json"] = res
		o.ServeJSON()
	}()

	var packet struct {
		Gender    int       `json:"gender"`
		Dob       time.Time `json:"dob"`
		Education string    `json:"education"`
	}

	err := json.Unmarshal(o.Ctx.Input.RequestBody, &packet)
	if err != nil {
		logs.Error("Parse request error", err)
		o.Ctx.Output.SetStatus(http.StatusBadRequest)
		res = models.NewResponseWithError("request_invalid", "Couldn't parse your request")
		return
	}
}
