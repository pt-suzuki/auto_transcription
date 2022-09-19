package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase "firebase.google.com/go"
	"github.com/pt-suzuki/auto_transcription/config"
	firebase2 "github.com/pt-suzuki/auto_transcription/infrastructure/firebase"
	"log"
)

func GetClient() *firestore.Client {
	appConfig := config.NewConfig()
	conf := &firebase.Config{
		ProjectID: appConfig.GcpProjectId,
	}
	app := firebase2.GetApp(conf)

	ctx := context.Background()
	fireStoreClient, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return fireStoreClient
}
