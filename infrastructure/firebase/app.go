package firebase

import (
	"context"
	firebase "firebase.google.com/go"
	"log"
)

func GetApp(conf *firebase.Config) *firebase.App {
	ctx := context.Background()
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Fatalln(err)
	}

	return app
}

func GetLocalApp(conf *firebase.Config) *firebase.App {
	ctx := context.Background()
	sa := GetServiceAccount()
	app, err := firebase.NewApp(ctx, conf, sa)
	if err != nil {
		log.Fatalln(err)
	}

	return app
}
