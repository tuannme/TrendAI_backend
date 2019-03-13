package twitterservice

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"github.com/trend-ai/TrendAI_mobile_backend/conf"
)

var config *oauth1.Config

type TwitterClient twitter.Client

func init() {
	// http.Client will automatically authorize Requests
	config = oauth1.NewConfig(conf.Get().TwitterConsumerKey, conf.Get().TwitterConsumerSecret)
}

func NewTwitterClient(accessToken string, accessTokenSecret string) *TwitterClient {
	token := oauth1.NewToken(accessToken, accessTokenSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	// twitter client
	client := TwitterClient(*twitter.NewClient(httpClient))
	return &client
}
