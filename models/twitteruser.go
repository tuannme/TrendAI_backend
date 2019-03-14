package models

import (
	"github.com/dghubble/go-twitter/twitter"
	"gopkg.in/mgo.v2/bson"
	"time"
)

type TwitterUser struct {
	Id              bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	TwitterId       int64         `json:"twitter_id" bson:"twitter_id"`
	Email           string        `json:"email" bson:"email"`
	Name            string        `json:"name" bson:"name"`
	ScreenName      string        `json:"screen_name" bson:"screen_name"`
	Description     string        `json:"description" bson:"description"`
	Lang            string        `json:"lang" bson:"lang"`
	Location        string        `json:"location" bson:"location"`
	Timezone        string        `json:"time_zone" bson:"time_zone"`
	StatusesCount   int           `json:"statuses_count" bson:"statuses_count"`
	FavouritesCount int           `json:"favourites_count" bson:"favourites_count"`
	FriendsCount    int           `json:"friends_count" bson:"friends_count"`
	CreatedAt       time.Time     `json:"created_at" bson:"created_at"`
}

type ExternalTwitterUser twitter.User

// Get user data from twitter user's data
func (t *ExternalTwitterUser) ToUser() User {
	return User{
		Name:            t.Name,
		Email:           t.Email,
		FavouritesCount: t.FavouritesCount,
		FollowersCount:  t.FollowersCount,
		FriendsCount:    t.FriendsCount,
		StatusesCount:   t.StatusesCount,
	}
}

// Sync new twitter data
func (u *User) SyncTwitterData(t *ExternalTwitterUser) {
	u.Name = t.Name
	u.FavouritesCount = t.FavouritesCount
	u.FollowersCount = t.FollowersCount
	u.StatusesCount = t.StatusesCount
}
