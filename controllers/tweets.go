package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/trend-ai/TrendAI_mobile_backend/models"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

// Operations about Users
type TweetsController struct {
	beego.Controller
}

// @Title Get
// @Description Get trending topics by location
// @Failure 200 {object} []models.TopicResponse
// @Failure 400 {object} models.ResponseWithError
// @Failure 500 {object} models.ResponseWithError
// @router / [get]
func (o *TweetsController) Get() {
	var res models.Response
	defer func() {
		o.Data["json"] = res
		o.ServeJSON()
	}()

	// Get topic ID from request
	topicIdStr := o.GetString("topic_id")
	if !bson.IsObjectIdHex(topicIdStr) {
		o.Ctx.Output.SetStatus(http.StatusBadRequest)
		res = models.NewResponseWithError("topic_not_found", "Topic doesn't exists.")
		return
	}

	var topicId = bson.ObjectIdHex(o.GetString("topic_id"))
	var topicCollection = models.GetTrendingTopicCollection()
	var topic models.TrendingTopic

	// Get topic in our database
	if err := topicCollection.FindId(topicId).One(&topic); err != nil {
		o.Ctx.Output.SetStatus(http.StatusBadRequest)
		res = models.NewResponseWithError("topic_not_found", "Topic doesn't exists.")
		return
	}

	// Get current user from request
	user := o.Ctx.Input.GetData("user").(models.User)

	// Create new twitter client from current user
	client := user.NewTwitterClient()

	includeEntities := true
	search, _, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query:           topic.Name,
		TweetMode:       "extended",
		IncludeEntities: &includeEntities,
		Count:           100,
	})
	if err != nil {
		logs.Error("Couldn't search tweets", err)
		o.Ctx.Output.SetStatus(http.StatusInternalServerError)
		res = models.NewResponseWithError("system_error", "Couldn't get tweets by topic.")
		return
	}
	for _, tweet := range search.Statuses {
		logs.Info(tweet)
	}
}
