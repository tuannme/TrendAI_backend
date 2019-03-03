package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	_ "github.com/trend-ai/TrendAI_mobile_backend/conf"
	_ "github.com/trend-ai/TrendAI_mobile_backend/routers"
	"github.com/trend-ai/TrendAI_mobile_backend/services/databases"
)

func main() {
	logs.Debug("App initiated: ", beego.BConfig.AppName)
	logs.Debug("App run mode: ", beego.BConfig.RunMode)

	// Close database connection when stop main function
	defer databases.GetMongoSession().Close()

	// Beego config
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// Run application
	beego.Run()
}
