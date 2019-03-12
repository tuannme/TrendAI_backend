package models

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/trend-ai/TrendAI_mobile_backend/services/databases"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var woeLocationCollection *mgo.Collection

func init() {
	woeLocationCollection = databases.GetMongoCollection("woeLocations")
}

func GetWoeLocationCollection() *mgo.Collection {
	return woeLocationCollection
}

const DefaultWoeID = 1

type WoeLocation struct {
	Id              bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Woeid           int64         `json:"woeid" bson:"woeid"`
	Name            string        `json:"name,omitempty" bson:"name,omitempty"`
	Country         string        `json:"country,omitempty" bson:"country,omitempty"`
	CountryCode     string        `json:"country_code,omitempty" bson:"country_code,omitempty"`
	ParentId        int           `json:"parent_id,omitempty" bson:"parent_id,omitempty"`
	PlaceType       PlaceType     `json:"place_type,omitempty" bson:"place_type,omitempty"`
	TrendsAvailable bool          `json:"trends_available,omitempty" bson:"trends_available,omitempty"`
}

type PlaceType struct {
	Code int    `json:"code" bson:"code"`
	Name string `json:"name" bson:"name"`
}

// Convert twitter location to our WOE location
func ToWoeLocation(location twitter.Location) WoeLocation {
	return WoeLocation{
		Woeid:       location.WOEID,
		Name:        location.Name,
		Country:     location.Country,
		CountryCode: location.CountryCode,
		ParentId:    location.ParentID,
		PlaceType: PlaceType{
			Code: location.PlaceType.Code,
			Name: location.PlaceType.Name,
		},
	}
}

// Sync WOE location with our database
func (woe *WoeLocation) Sync() error {
	var woeLocation WoeLocation

	// Find WOEID in our database
	err := woeLocationCollection.Find(bson.M{"woeid": woe.Woeid}).One(&woeLocation)

	// If WOE doesn't exists on our database, insert it
	if err != nil {
		woe.Id = bson.NewObjectId()
		err := woeLocationCollection.Insert(&woe)
		if err != nil {
			return err
		}
		return nil
	}

	// If WOE already exists on our database, update it
	if woeLocation.TrendsAvailable {
		woe.TrendsAvailable = true
	}
	if err = woeLocationCollection.UpdateId(woeLocation.Id, woe); err != nil {
		return err
	}
	return nil
}
