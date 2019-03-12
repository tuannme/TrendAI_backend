package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/jasonlvhit/gocron"
	_ "github.com/trend-ai/TrendAI_mobile_backend/conf"
	_ "github.com/trend-ai/TrendAI_mobile_backend/routers"
	"github.com/trend-ai/TrendAI_mobile_backend/services/cronservice"
	"github.com/trend-ai/TrendAI_mobile_backend/services/databases"
)

// Function contains all cron jobs necessary for our application
func cronJobs() {
	gocron.Every(1).Days().Do(cronservice.SyncWoeLocation)
	<-gocron.Start()
}

// Function to run at start of application
func runAtStart() {
	cronservice.SyncWoeLocation()
}

func main() {
	logs.Debug("App initiated: ", beego.BConfig.AppName)
	logs.Debug("App run mode: ", beego.BConfig.RunMode)

	// Start some functions before our application start
	go runAtStart()

	// Start schedule cron jobs
	go cronJobs()

	// Close database connection when stop main function
	defer databases.GetMongoSession().Close()

	// Run application
	beego.Run()
}
