package databases

import (
	"github.com/astaxie/beego/logs"
	"github.com/trend-ai/TrendAI_mobile_backend/conf"
	"gopkg.in/mgo.v2"
)

var session *mgo.Session
var db *mgo.Database

// Init mongo service
func init() {
	logs.Debug("Mongo service initiated!")
	// Connect to mongodb
	var err error
	session, err = mgo.Dial(conf.Get().MongoUri)
	if err != nil {
		panic(err)
	}

	// Select database
	db = session.DB(conf.Get().MongoDatabase)
}

// Get current mongo session
func GetMongoSession() *mgo.Session {
	return session
}

// Get current mongo database connection
func GetMongoDatabase() *mgo.Database {
	return db
}

// Get mongo collection by collection name
func GetMongoCollection(collection string) *mgo.Collection {
	return db.C(collection)
}
