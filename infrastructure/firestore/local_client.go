package firestore

import (
	"cloud.google.com/go/firestore"
	"context"
	firebase2 "github.com/pt-suzuki/auto_transcription/infrastructure/firebase"
	"log"
)

func GetLocalClient() *firestore.Client {
	app := firebase2.GetLocalApp(nil)

	ctx := context.Background()
	fireStoreClient, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	return fireStoreClient
}
