package models

import (
	"errors"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/trend-ai/TrendAI_mobile_backend/services/databases"
	"github.com/trend-ai/TrendAI_mobile_backend/services/twitterservice"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"time"
)

var userCollection *mgo.Collection

func init() {
	userCollection = databases.GetMongoCollection("users")
}

func GetUserCollection() *mgo.Collection {
	return userCollection
}

const (
	UserGenderMale   = 0
	UserGenderFemale = 1
)

var UserGenders = map[int]string{
	UserGenderMale:   "male",
	UserGenderFemale: "female",
}

type User struct {
	Id                 bson.ObjectId       `json:"id,omitempty" bson:"_id,omitempty"`
	Name               string              `json:"name" bson:"name"`
	Email              string              `json:"email" bson:"email"`
	Gender             int                 `json:"gender" bson:"gender"`
	Dob                time.Time           `json:"dob" bson:"dob"`
	Education          string              `json:"education" bson:"education"`
	Location           *Location           `json:"location,omitempty" bson:"location,omitempty"`
	InterestCategories []bson.ObjectId     `json:"interest_categories" bson:"interest_categories"`
	FavouritesCount    int                 `json:"favourites_count" bson:"favourites_count"`
	FollowersCount     int                 `json:"followers_count" bson:"followers_count"`
	FriendsCount       int                 `json:"friends_count" bson:"friends_count"`
	StatusesCount      int                 `json:"statuses_count" bson:"statuses_count"`
	TweetStat          TweetStat           `json:"tweet_stat" bson:"tweet_stat,omitempty"`
	ExternalUsers      []ExternalUser      `json:"external_users" bson:"external_users,omitempty"`
	TwitterCredentials *TwitterCredentials `json:"twitter_credentials,omitempty" bson:"twitter_credentials,omitempty"`
	UpdatedAt          time.Time           `json:"updated_at" bson:"updated_at,omitempty"`
	CreatedAt          time.Time           `json:"created_at" bson:"created_at,omitempty"`
}

type Location struct {
	Lat   float64 `json:"lat" bson:"lat"`
	Lng   float64 `json:"lng" bson:"lng"`
	Woeid int64   `json:"woeid,omitempty" bson:"woeid,omitempty"`
}

type TweetStat struct {
	ReplyCount    int `json:"reply_count" bson:"reply_count"`
	RetweetCount  int `json:"retweet_count" bson:"retweet_count"`
	FavoriteCount int `json:"favorite_count" bson:"favorite_count"`
}

type ExternalUser struct {
	AppId           string    `json:"app_id" bson:"app_id"`
	UserId          string    `json:"user_id" bson:"user_id"`
	Username        string    `json:"username" bson:"username"`
	LastConnectedAt time.Time `json:"last_connected_at" bson:"last_connected_at"`
	CreatedAt       time.Time `json:"created_at" bson:"created_at"`
}

type TwitterCredentials struct {
	AccessToken       string `json:"access_token" json:"access_token"`
	AccessTokenSecret string `json:"access_token_secret" bson:"access_token_secret"`
}

type UserResponse struct {
	Id        string    `json:"id,omitempty"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Gender    string    `json:"gender"`
	Dob       time.Time `json:"dob"`
	Education string    `json:"education"`
	CreatedAt time.Time `json:"created_at"`
}

// Get response data for current user
func (u *User) ToResponse() UserResponse {
	gender, _ := UserGenderToStr(u.Gender)
	return UserResponse{
		Id:        u.Id.Hex(),
		Name:      u.Name,
		Email:     u.Email,
		Gender:    gender,
		Dob:       u.Dob,
		Education: u.Education,
		CreatedAt: u.CreatedAt,
	}
}

// Create new twitter client from twitter credentials
func (t *TwitterCredentials) NewTwitterClient() *twitter.Client {
	return twitterservice.NewTwitterClient(t.AccessToken, t.AccessTokenSecret)
}

// Create new twitter client from current user's credentials
func (u *User) NewTwitterClient() *twitter.Client {
	return u.TwitterCredentials.NewTwitterClient()
}

// Convert gender in int to string
func UserGenderToStr(gender int) (string, error) {
	val, exists := UserGenders[gender]
	if !exists {
		return "", errors.New("gender doesn't exists")
	}
	return val, nil
}

// Convert gender in string to int
func UserGenderToInt(gender string) (int, error) {
	for k, v := range UserGenders {
		if v == gender {
			return k, nil
		}
	}
	return 0, errors.New("gender doesn't exists")
}
