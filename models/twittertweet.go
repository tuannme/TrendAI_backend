package models

import (
	"github.com/trend-ai/TrendAI_mobile_backend/models/tweetentities"
	"github.com/trend-ai/TrendAI_mobile_backend/services/databases"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var tweetCollection *mgo.Collection

func init() {
	tweetCollection = databases.GetMongoCollection("twitterTweets")
}

func GetTweetCollection() *mgo.Collection {
	return tweetCollection
}

type Tweet struct {
	Id               bson.ObjectId         `json:"id,omitempty" bson:"_id,omitempty"`
	TwitterId        int64                 `json:"twitter_id" bson:"twitter_id"`
	TwitterUserId    bson.ObjectId         `json:"twitter_user_id" bson:"twitter_user_id"`
	Text             string                `json:"text" bson:"text"`
	FullText         string                `json:"full_text" bson:"full_text"`
	DisplayTextRange tweetentities.Indices `json:"display_text_range" bson:"display_text_range"`
	Truncated        bool                  `json:"truncated" bson:"truncated"`
	Entities         *TweetEntities        `json:"entities,omitempty" bson:"entities,omitempty"`
	Lang             string                `json:"lang" bson:"lang"`
	FavoriteCount    int                   `json:"favorite_count" bson:"favorite_count"`
	QuoteCount       int                   `json:"quote_count" bson:"quote_count"`
	ReplyCount       int                   `json:"reply_count" bson:"reply_count"`
	RetweetCount     int                   `json:"retweet_count" bson:"retweet_count"`
	CreatedAt        time.Time             `json:"created_at" bson:"created_at"`
}

type TweetEntities struct {
	Hashtags     []tweetentities.HashtagEntity `json:"hashtags" bson:"hashtags"`
	Media        []tweetentities.MediaEntity   `json:"media" bson:"media"`
	Urls         []tweetentities.URLEntity     `json:"urls" bson:"urls"`
	UserMentions []tweetentities.MentionEntity `json:"user_mentions" bson:"user_mentions"`
}
