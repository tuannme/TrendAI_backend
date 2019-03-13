package twitterservice

import (
	"errors"
	"github.com/dghubble/go-twitter/twitter"
)

// Get Woe location by latitude and longitude
func (c *TwitterClient) GetWoeLocationByLatLng(lat, lng float64) (*twitter.Location, error) {
	// Find WOEID by latitude, longitude
	locations, _, err := c.Trends.Closest(&twitter.ClosestParams{
		Lat:  lat,
		Long: lng,
	})

	// Check if location is invalid
	if err != nil || locations == nil || len(locations) <= 0 {
		if err != nil {
			return nil, err
		}
		return nil, errors.New("invalid location latitude, longitude")
	}

	// Get closest location
	location := locations[0]
	return &location, nil
}
