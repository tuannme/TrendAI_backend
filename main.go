package main

import (
	"github.com/astaxie/beego"
	_ "github.com/trend-ai/TrendAI_mobile_backend/conf"
	_ "github.com/trend-ai/TrendAI_mobile_backend/routers"
	"github.com/trend-ai/TrendAI_mobile_backend/services/databases"
)

func main() {
	// Get firestore client
	firestoreClient := databases.GetFirestoreClient()

	// Close firestore connection when main function stopped
	defer func() {
		_ = firestoreClient.Close()
	}()

	// Beego config
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// Run application
	beego.Run()
}
