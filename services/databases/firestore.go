package databases

import (
	"cloud.google.com/go/firestore"
	"context"
	"firebase.google.com/go"
	"github.com/astaxie/beego/logs"
	"github.com/trend-ai/TrendAI_mobile_backend/conf"
	"google.golang.org/api/option"
	"os"
)

var serviceAccountFile = "conf/service_account.json"
var firestoreClient *firestore.Client
var Context = context.Background()

func init() {
	var err error
	firestoreClient, err = InitFirestore(serviceAccountFile)
	if err != nil {
		panic(err)
	}
}

// Init firestore service
func InitFirestore(serviceAccountFilename string) (*firestore.Client, error) {
	logs.Debug("Firestore service initiated!")

	var err error
	var app *firebase.App

	if _, err := os.Stat(serviceAccountFilename); os.IsNotExist(err) {
		// Use project ID if service account doesn't exists
		logs.Debug("Used GCP project ID")
		appConf := &firebase.Config{ProjectID: conf.Get().GoogleCloudProject}
		app, err = firebase.NewApp(Context, appConf)
	} else {
		// Use service account
		logs.Debug("Used GCP service account")
		sa := option.WithCredentialsFile(serviceAccountFilename)
		app, err = firebase.NewApp(Context, nil, sa)
	}

	if err != nil {
		return nil, err
	}

	client, err := app.Firestore(Context)
	if err != nil {
		return nil, err
	}

	firestoreClient = client
	return client, nil
}

// Get current firestore client
func GetFirestoreClient() *firestore.Client {
	return firestoreClient
}

// Get firestore collection by collection name
func GetFirestoreCollection(collection string) *firestore.CollectionRef {
	return firestoreClient.Collection(collection)
}
