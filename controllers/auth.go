package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type AuthController struct {
	beego.Controller
}

// @Title Login
// @Description Login API
// @Param	access_token	body	string	true	"Twitter access token"
// @Param	access_token_secret	body	string	true	"Twitter access token secret"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /login [post]
func (o *AuthController) Login() {
	var response interface{}
	o.Data["json"] = &response

	var packet struct {
		AccessToken       string `json:"access_token"`
		AccessTokenSecret string `json:"access_token_secret"`
	}

	err := json.Unmarshal(o.Ctx.Input.RequestBody, &packet)
	if err != nil {
		logs.Error("Parse request errors", err)
		response = err.Error()
		o.Ctx.Output.SetStatus(522)
		o.ServeJSON()
		return
	}

	response = struct {
		Error   bool
		Message string
	}{
		Error:   false,
		Message: "OK",
	}
}
