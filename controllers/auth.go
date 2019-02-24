package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/trend-ai/TrendAI_mobile_backend/conf"
	"github.com/trend-ai/TrendAI_mobile_backend/models"
	"gopkg.in/mgo.v2/bson"
	"strconv"
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
	var res models.Response
	defer func() {
		o.Data["json"] = res
		o.ServeJSON()
	}()

	var packet struct {
		AccessToken       string `json:"access_token"`
		AccessTokenSecret string `json:"access_token_secret"`
	}

	err := json.Unmarshal(o.Ctx.Input.RequestBody, &packet)
	if err != nil {
		logs.Error("Parse request error", err)
		o.Ctx.Output.SetStatus(522)
		res = err.Error()
		return
	}

	// http.Client will automatically authorize Requests
	config := oauth1.NewConfig(conf.Get().TwitterConsumerKey, conf.Get().TwitterConsumerSecret)
	token := oauth1.NewToken(packet.AccessToken, packet.AccessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)

	// twitter client
	client := twitter.NewClient(httpClient)

	// Validate credentials
	IncludeEmail := true
	twitterUser, _, err := client.Accounts.VerifyCredentials(&twitter.AccountVerifyParams{
		IncludeEmail: &IncludeEmail,
	})
	if err != nil {
		logs.Error("Authentication failed:", err.Error())
		o.Ctx.Output.SetStatus(401)
		res = models.NewResponseWithError("unauthorized", "Could't validate your credentials")
		return
	}

	// Get internal user which matched with twitter email
	var user models.User
	err = models.GetUserCollection().Find(bson.M{"email": twitterUser.Email}).One(&user)
	if err != nil {
		user = models.User{
			Name:      twitterUser.Name,
			Email:     twitterUser.Email,
			CreatedAt: bson.Now(),
		}
	} else {
		// Re-sync twitter data
		user.Name = twitterUser.Name
	}

	// Current external user
	externalUser := models.ExternalUser{
		AppId:           conf.Get().TwitterAppId,
		UserId:          strconv.FormatInt(twitterUser.ID, 10),
		CreatedAt:       bson.Now(),
		LastConnectedAt: bson.Now(),
	}

	// Check current external user exists in this user
	foundKey := -1
	for i, v := range user.ExternalUsers {
		if v.AppId == externalUser.AppId && v.UserId == externalUser.UserId {
			foundKey = i
			break
		}
	}

	// If not exists, assign external user for this user
	if foundKey < 0 {
		user.ExternalUsers = append(user.ExternalUsers, externalUser)
	} else {
		//If exists, update last connected at
		user.ExternalUsers[foundKey].LastConnectedAt = bson.Now()
	}

	// Save data
	if user.Id.Valid() {
		err = models.GetUserCollection().UpdateId(user.Id, user)
	} else {
		err = models.GetUserCollection().Insert(user)
	}

	// If saving fail, then response error
	if err != nil {
		logs.Error("Couldn't update user", err, user.Id)
		o.Ctx.Output.SetStatus(401)
		res = models.NewResponseWithError("unauthorized", "Unauthorized")
		return
	}

	res = user
}
