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
	if len(os.Args) < 3 {
		log.Fatal("Usage: go run main.go <export.json file path> <userID>")
	}

	filePath := os.Args[1]
	userID := os.Args[2]
	ctx := context.Background()

	items, err := loadItemsFromJSON(filePath)
	if err != nil {
		log.Fatalf("Failed to load items from JSON: %v", err)
	}

	client, err := createFirestoreClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create Firestore client: %v", err)
	}
	defer client.Close()

	if err := importItems(ctx, client, userID, items); err != nil {
		log.Fatalf("Failed to import items: %v", err)
	}

	fmt.Printf("Successfully imported %d items for user %s\n", len(items), userID)
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

func loadItemsFromJSON(filePath string) ([]*model.Item, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %w", err)
	}

	var items []*model.Item
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %w", err)
	}

	return items, nil
}

func importItems(ctx context.Context, client *firestore.Client, userID string, items []*model.Item) error {
	batch := client.Batch()
	count := 0

	for _, item := range items {
		item.UserID = userID

		docRef := client.Collection("users/" + userID + "/items").Doc(item.ID)
		batch.Set(docRef, item)
		count++

		if count%500 == 0 {
			if _, err := batch.Commit(ctx); err != nil {
				return fmt.Errorf("error committing batch: %w", err)
			}
			batch = client.Batch()
		}
	}

	if count%500 != 0 {
		if _, err := batch.Commit(ctx); err != nil {
			return fmt.Errorf("error committing final batch: %w", err)
		}
	}

	return nil
}
