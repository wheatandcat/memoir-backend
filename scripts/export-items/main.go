package main

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func main() {
	ctx := context.Background()
	opt := option.WithCredentialsFile("../../serviceAccount.json")
	config := &firebase.Config{ProjectID: os.Getenv("FIREBASE_PROJECT_ID")}
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	userID := "test_id"

	firestore, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalf("error getting Firestore client: %v\n", err)
	}

	items, err := firestore.Collection("users/" + userID + "/items").Documents(ctx).GetAll()
	if err != nil {
		log.Fatalf("error getting items: %v\n", err)
	}

	for _, item := range items {
		fmt.Println(item.Data()["Title"])
	}
}
