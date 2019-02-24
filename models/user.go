package models

import (
	"github.com/trend-ai/TrendAI_mobile_backend/services/databases"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var userCollection *mgo.Collection

func init() {
	userCollection = databases.GetMongoCollection("user")
}

func GetUserCollection() *mgo.Collection {
	return userCollection
}

type User struct {
	Id            bson.ObjectId  `json:"id,omitempty" bson:"_id,omitempty"`
	Name          string         `json:"name" bson:"name"`
	Email         string         `json:"email" bson:"email"`
	ExternalUsers []ExternalUser `json:"external_users" bson:"external_users"`
	CreatedAt     time.Time      `json:"created_at" bson:"created_at"`
}

type ExternalUser struct {
	AppId           string    `json:"app_id" bson:"app_id"`
	UserId          string    `json:"user_id" bson:"user_id"`
	LastConnectedAt time.Time `json:"last_connected_at" bson:"last_connected_at"`
	CreatedAt       time.Time `json:"created_at" bson:"created_at"`
}
