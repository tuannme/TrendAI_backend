package models

import (
	"github.com/trend-ai/TrendAI_mobile_backend/services/databases"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var trendingTopicCollection *mgo.Collection

func init() {
	trendingTopicCollection = databases.GetMongoCollection("trendingTopics")
}

func GetTrendingTopicCollection() *mgo.Collection {
	return trendingTopicCollection
}

type TrendingTopic struct {
	Id        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name      string        `json:"name" bson:"name"`
	Volume    int64         `json:"volume" bson:"volume"`
	Woeids    []int64       `json:"woeids,omitempty" bson:"woeids,omitempty"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updated_at,omitempty"`
	CreatedAt time.Time     `json:"created_at" bson:"created_at,omitempty"`
}

type TrendingTopicResponse struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Volume int64  `json:"volume"`
}

func (t *TrendingTopic) ToResponse() TrendingTopicResponse {
	return TrendingTopicResponse{
		Id:     t.Id.Hex(),
		Name:   t.Name,
		Volume: t.Volume,
	}
}
