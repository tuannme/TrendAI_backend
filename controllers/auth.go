package controllers

import (
	"cloud.google.com/go/firestore"
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/trend-ai/TrendAI_mobile_backend/conf"
	"github.com/trend-ai/TrendAI_mobile_backend/models"
	"github.com/trend-ai/TrendAI_mobile_backend/services/authentications"
	"github.com/trend-ai/TrendAI_mobile_backend/services/databases"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
	"time"
)

type AuthController struct {
	beego.Controller
}

// @Title Login
// @Description Login API
// @Param	access_token	body	string	true	"Twitter access token"
// @Param	access_token_secret	body	string	true	"Twitter access token secret"
// @Success 200 {object} models.AuthenticationResponse
// @Failure 400 {object} models.ResponseWithError
// @Failure 401 {object} models.ResponseWithError
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
		o.Ctx.Output.SetStatus(http.StatusBadRequest)
		res = models.NewResponseWithError("unauthorized", "Couldn't parse your request")
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
		o.Ctx.Output.SetStatus(http.StatusUnauthorized)
		res = models.NewResponseWithError("unauthorized", "Could't validate your credentials")
		return
	}

	ctx := databases.Context
	userCollection := models.GetUserCollection()

	// Get internal user which matched with twitter email
	var user models.User
	userDoc, err := userCollection.Where("email", "==", twitterUser.Email).Documents(ctx).Next()
	var userRef *firestore.DocumentRef
	if err != nil {
		// Create new user document
		userRef = userCollection.NewDoc()
		user = models.User{
			Name:      twitterUser.Name,
			Email:     twitterUser.Email,
			CreatedAt: time.Now().UTC(),
		}
	} else {
		userRef = userDoc.Ref
		if err = userDoc.DataTo(&user); err != nil {
			logs.Error("Parse user data failure:", err)
			o.Ctx.Output.SetStatus(http.StatusUnauthorized)
			res = models.NewResponseWithError("unauthorized", "Could't validate your credentials")
			return
		}
		// Re-sync twitter data
		user.Name = twitterUser.Name
	}

	// Current external user
	externalUser := models.ExternalUser{
		AppId:           conf.Get().TwitterAppId,
		UserId:          strconv.FormatInt(twitterUser.ID, 10),
		CreatedAt:       time.Now().UTC(),
		LastConnectedAt: time.Now().UTC(),
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
	_, err = userRef.Set(ctx, user)

	// If saving fail, then response error
	if err != nil {
		logs.Error("Couldn't update user", err)
		o.Ctx.Output.SetStatus(http.StatusUnauthorized)
		res = models.NewResponseWithError("unauthorized", "Unauthorized")
		return
	}

	// Get user ID
	user.Id = userRef.ID

	// Generate authentication to for current user
	authenticationToken, err := authentications.GenerateAuthenticationTokenByUser(user)
	if err != nil {
		logs.Error("Couldn't generate authentication token", err)
		o.Ctx.Output.SetStatus(http.StatusUnauthorized)
		res = models.NewResponseWithError("unauthorized", "Unauthorized")
		return
	}

	// Respond authentication data
	res = user.ToAuthenticationResponse(*authenticationToken)
}
