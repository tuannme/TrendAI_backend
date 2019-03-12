package cronservice

import (
	"github.com/astaxie/beego/logs"
	"github.com/trend-ai/TrendAI_mobile_backend/conf"
	"github.com/trend-ai/TrendAI_mobile_backend/models"
)

// Sync WOE locations which have trending to our database
func SyncWoeLocation() {
	logs.Info("SyncWoeLocation is running...")
	certs := models.TwitterCredentials{
		AccessToken:       conf.Get().TwitterServerAccessToken,
		AccessTokenSecret: conf.Get().TwitterServerAccessTokenSecret,
	}
	client := certs.NewTwitterClient()
	locations, _, err := client.Trends.Available()
	if err != nil {
		logs.Error("Couldn't get trends available", err)
		return
	}
	synced := 0
	for _, location := range locations {
		woeLocation := models.ToWoeLocation(location)
		woeLocation.TrendsAvailable = true
		err := woeLocation.Sync()
		if err != nil {
			logs.Error("Couldn't sync WOE location", err)
			return
		}
		synced++
	}
	logs.Info(synced, "locations have been synced")
}
