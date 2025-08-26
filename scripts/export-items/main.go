package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/wheatandcat/memoir-backend/graph/model"
	"google.golang.org/api/option"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: go run main.go <userID>")
	}

	userID := os.Args[1]
	ctx := context.Background()

	client, err := createFirestoreClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}
	defer client.Close()

	items, err := getAllItems(ctx, client, userID)
	if err != nil {
		log.Fatalf("Failed to get items: %v", err)
	}

	if err := exportToJSON(items, userID); err != nil {
		log.Fatalf("Failed to export to JSON: %v", err)
	}

	fmt.Printf("Successfully exported %d items for user %s to %s_items.json\n", len(items), userID, userID)
}

func createFirestoreClient(ctx context.Context) (*firestore.Client, error) {
	opt := option.WithCredentialsFile("../../serviceAccount.json")
	config := &firebase.Config{ProjectID: os.Getenv("FIREBASE_PROJECT_ID")}
	app, err := firebase.NewApp(ctx, config, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %w", err)
	}

	client, err := app.Firestore(ctx)
	if err != nil {
		return nil, fmt.Errorf("error getting Firestore client: %w", err)
	}

	return client, nil
}

func getAllItems(ctx context.Context, client *firestore.Client, userID string) ([]*model.Item, error) {
	docs, err := client.Collection("users/" + userID + "/items").Documents(ctx).GetAll()
	if err != nil {
		return nil, fmt.Errorf("error getting items: %w", err)
	}

	items := make([]*model.Item, 0, len(docs))
	for _, doc := range docs {
		var item model.Item
		if err := doc.DataTo(&item); err != nil {
			log.Printf("Warning: Failed to parse item %s: %v", doc.Ref.ID, err)
			continue
		}
		items = append(items, &item)
	}

	return items, nil
}

func exportToJSON(items []*model.Item, userID string) error {
	filename := "export.json"

	jsonData, err := json.MarshalIndent(items, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling to JSON: %w", err)
	}

	if err := os.WriteFile(filename, jsonData, 0600); err != nil {
		return fmt.Errorf("error writing file: %w", err)
	}

	return nil
}
