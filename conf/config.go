package conf

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	AppName               string
	AppKey                string
	MongoDatabase         string
	MongoUri              string
	TwitterAppId          string
	TwitterConsumerKey    string
	TwitterConsumerSecret string
}

var config Config

// Init config variables
func init() {
	initEnv()
	config.AppKey = os.Getenv("APP_NAME")
	config.AppKey = os.Getenv("APP_KEY")
	config.MongoDatabase = os.Getenv("MONGO_DATABASE")
	config.MongoUri = os.Getenv("MONGO_URI")
	config.TwitterAppId = os.Getenv("TWITTER_APP_ID")
	config.TwitterConsumerKey = os.Getenv("TWITTER_CONSUMER_KEY")
	config.TwitterConsumerSecret = os.Getenv("TWITTER_CONSUMER_SECRET")
}

// Init load environment variables from .env
func initEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

// Get all configs
func Get() Config {
	return config
}