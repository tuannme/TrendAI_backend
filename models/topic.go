package models

import (
	"github.com/trend-ai/TrendAI_mobile_backend/services/databases"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var topicCollection *mgo.Collection

func init() {
	topicCollection = databases.GetMongoCollection("topics")
}

func GetTopicCollection() *mgo.Collection {
	return topicCollection
}

type Topic struct {
	Id     bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string        `json:"name" bson:"name"`
	Volume int64         `json:"volume" bson:"volume"`
}

type TopicResponse struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Volume int64  `json:"volume"`
}

func (t *Topic) ToResponse() TopicResponse {
	return TopicResponse{
		Id:     t.Id.Hex(),
		Name:   t.Name,
		Volume: t.Volume,
	}
}
