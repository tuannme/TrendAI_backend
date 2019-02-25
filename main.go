package main

import (
	"github.com/astaxie/beego"
	_ "github.com/trend-ai/TrendAI_mobile_backend/conf"
	_ "github.com/trend-ai/TrendAI_mobile_backend/routers"
	"github.com/trend-ai/TrendAI_mobile_backend/services/databases"
)

func main() {
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
