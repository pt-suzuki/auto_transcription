package firebase_admin

import (
	"context"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"log"
)

func GetFirebaseAdminClient() *auth.Client {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}
	client, err := app.Auth(context.Background())
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}
	return client
}
