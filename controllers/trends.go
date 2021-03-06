package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/trend-ai/TrendAI_mobile_backend/models"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

// Operations about Users
type TrendsController struct {
	beego.Controller
}

// @Title Get
// @Description Get trending topics by location
// @Failure 200 {object} []models.TrendingTopicResponse
// @Failure 400 {object} models.ResponseWithError
// @Failure 500 {object} models.ResponseWithError
// @router / [get]
func (o *TrendsController) Get() {
	var res models.Response
	defer func() {
		o.Data["json"] = res
		o.ServeJSON()
	}()

	// Get current user from request
	user := o.Ctx.Input.GetData("user").(models.User)

	// Create new twitter client from current user
	client := user.NewTwitterClient()

	// Default WOEID
	var woeId int64 = models.DefaultWoeID

	// Get lat, lng from request
	var lat, lng float64
	lat, errLat := o.GetFloat("lat")
	lng, errLng := o.GetFloat("lng")

	// If request has location
	if errLat == nil && errLng == nil {
		// Find WOEID by latitude, longitude
		locations, _, err := client.Trends.Closest(&twitter.ClosestParams{
			Lat:  lat,
			Long: lng,
		})

		// Check if location is invalid
		if err != nil || locations == nil || len(locations) <= 0 {
			if err != nil {
				logs.Error("Couldn't get WOEID by location", err)
			}
			o.Ctx.Output.SetStatus(http.StatusBadRequest)
			res = models.NewResponseWithError("invalid_location", "Invalid location.")
			return
		}

		// Get closest location
		location := locations[0]
		woeId = location.WOEID

		// Update WOE location to our database
		woeLocation := models.ToWoeLocation(location)
		if err = woeLocation.Sync(); err != nil {
			logs.Error("Couldn't update WOE location", err)
		}

		// Save location to current user's data if they haven't location yet
		if user.Location == nil {
			user.Location = &models.Location{
				Lat:   lat,
				Lng:   lng,
				Woeid: woeId,
			}
			// Save user's data
			if err := models.GetUserCollection().UpdateId(user.Id, user); err != nil {
				logs.Error("Update user's location failed", err)
			}
		}
	}

	// Get trending topic for this location
	trends, _, err := client.Trends.Place(woeId, nil)
	if err != nil {
		logs.Error("Couldn't get trending data", err)
		o.Ctx.Output.SetStatus(http.StatusBadRequest)
		res = models.NewResponseWithError("get_failed", "Couldn't get trending data.")
		return
	}

	// Init response data
	topicsResponse := make([]models.TrendingTopicResponse, 0)

	// Check if trending topics available for this location
	if trends != nil && len(trends) > 0 {
		// Get the closest trending
		trendsList := trends[0]

		// Get topic collection
		topicCollection := models.GetTrendingTopicCollection()

		// Current time
		now := time.Now().UTC()

		// Loop all trending topics to update to database
		for _, trend := range trendsList.Trends {
			var topic models.TrendingTopic
			if err := topicCollection.Find(bson.M{"name": trend.Name}).One(&topic); err != nil {
				// Topic doesn't exists, create it
				topic = models.TrendingTopic{
					Id:        bson.NewObjectId(),
					Name:      trend.Name,
					Volume:    trend.TweetVolume,
					UpdatedAt: now,
					CreatedAt: now,
				}
				if woeId != models.DefaultWoeID {
					topic.Woeids = []int64{woeId}
				}
				if err := topicCollection.Insert(&topic); err != nil {
					logs.Error("Couldn't insert new topic", err)
					o.Ctx.Output.SetStatus(http.StatusInternalServerError)
					res = models.NewResponseWithError("get_failed", "Couldn't get trending data.")
					return
				}
			} else {
				// If topic already exists, update it
				topic.Volume = trend.TweetVolume
				topic.UpdatedAt = now
				if woeId != models.DefaultWoeID {
					// Find woeid in this topic, if woeid doesn't exists in list, add it
					found := false
					for _, v := range topic.Woeids {
						if v == woeId {
							found = true
							break
						}
					}
					if !found {
						topic.Woeids = append(topic.Woeids, woeId)
					}
				}
				if err := topicCollection.UpdateId(topic.Id, topic); err != nil {
					logs.Error("Couldn't update topic", err)
					o.Ctx.Output.SetStatus(http.StatusInternalServerError)
					res = models.NewResponseWithError("get_failed", "Couldn't get trending data.")
					return
				}
			}

			// Append data to response object
			topicsResponse = append(topicsResponse, topic.ToResponse())
		}
	}

	// Set response data
	res = topicsResponse
}
