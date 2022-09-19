package firestorage

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/storage"
	"github.com/pt-suzuki/auto_transcription/config"
	firebase2 "github.com/pt-suzuki/auto_transcription/infrastructure/firebase"
	"log"
)

func GetClient() *storage.Client {
	appConfig := config.NewConfig()
	conf := &firebase.Config{
		ProjectID:     appConfig.GcpProjectId,
		StorageBucket: appConfig.StorageBucket,
	}
	app := firebase2.GetApp(conf)

	ctx := context.Background()
	storageClient, err := app.Storage(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return storageClient
}
