package firestorage

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/storage"
	"github.com/pt-suzuki/auto_transcription/config"
	firebase2 "github.com/pt-suzuki/auto_transcription/infrastructure/firebase"
	"log"
)

func GetLocalClient() *storage.Client {
	appConfig := config.NewConfig()
	conf := &firebase.Config{
		StorageBucket: appConfig.StorageBucket,
	}
	app := firebase2.GetLocalApp(conf)

	ctx := context.Background()
	fireStoreClient, err := app.Storage(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return fireStoreClient
}
