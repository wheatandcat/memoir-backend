package graph

import (
	"context"
	"fmt"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/wheatandcat/memoir-backend/auth"
	"github.com/wheatandcat/memoir-backend/client/timegen"
	"github.com/wheatandcat/memoir-backend/client/uuidgen"
)

// Graph Graph struct
type Graph struct {
	UserID          string
	FirebaseUID     string
	FirestoreClient *firestore.Client
	App             *Application
	Client          *Client
}

type Client struct {
	UUID uuidgen.UUIDGenerator
	Time timegen.TimeGenerator
}

// NewGraph Graphを作成
func NewGraph(ctx context.Context, app *Application, f *firestore.Client) (*Graph, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("Access denied")
	}

	if user.FirebaseUID != "" {
		u, err := app.UserRepository.FindByFirebaseUID(ctx, f, user.FirebaseUID)
		if err != nil {
			return nil, fmt.Errorf("Firebase Auth Invalid")
		}

		user.ID = u.ID
	}

	log.Println("UserID:" + user.ID)

	if user.ID == "" {
		return nil, fmt.Errorf("UserID Invalid")
	}

	return NewGraphWithSetUserID(app, f, user.ID), nil
}

// NewGraphWithSetUserID Graphを作成（ログイン前）
func NewGraphWithSetUserID(app *Application, f *firestore.Client, uid string) *Graph {
	client := &Client{
		UUID: &uuidgen.UUID{},
		Time: &timegen.Time{},
	}

	return &Graph{
		UserID:          uid,
		FirestoreClient: f,
		App:             app,
		Client:          client,
	}
}
