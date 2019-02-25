package databases

import (
	"cloud.google.com/go/firestore"
	"context"
	"firebase.google.com/go"
	"github.com/astaxie/beego/logs"
	"google.golang.org/api/option"
)

var firestoreClient *firestore.Client
var Context = context.Background()

func init() {
	var err error
	firestoreClient, err = InitFirestore("conf/service_account.json")
	if err != nil {
		panic(err)
	}
}

// Init firestore service
func InitFirestore(serviceAccountFilename string) (*firestore.Client, error) {
	logs.Debug("Firestore service initiated!")

	// Use a service account
	sa := option.WithCredentialsFile(serviceAccountFilename)
	app, err := firebase.NewApp(Context, nil, sa)

	//conf := &firebase.Config{ProjectID: "ascendant-line-229014"}
	//app, err := firebase.NewApp(Context, conf)

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
