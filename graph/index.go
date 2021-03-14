package graph

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/wheatandcat/memoir-backend/auth"
	"github.com/wheatandcat/memoir-backend/client/timegen"
	"github.com/wheatandcat/memoir-backend/client/uuidgen"
)

// Graph Graph struct
type Graph struct {
	UserID          string
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
