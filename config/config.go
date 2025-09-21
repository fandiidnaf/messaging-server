package config

import (
	"context"
	"log"

	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func InitFirebaseApp() *firebase.App {
	opt := option.WithCredentialsFile("config/firebase/service-account-firebase.json")

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("failed to init firebase: %v", err)
	}

	return app
}
