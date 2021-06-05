package graph

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/99designs/gqlgen/graphql"
	"github.com/wheatandcat/memoir-backend/auth"
	"github.com/wheatandcat/memoir-backend/client/authToken"
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
	UUID      uuidgen.UUIDGenerator
	Time      timegen.TimeGenerator
	AuthToken authToken.AuthTokenClient
}

// NewGraph Graphを作成
func NewGraph(ctx context.Context, app *Application, f *firestore.Client) (*Graph, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return nil, fmt.Errorf("Access denied")
	}

	if user.FirebaseUID == "" {
		// Firebase認証無し
		u, err := app.UserRepository.FindDatabaseDataByUID(ctx, f, user.ID)
		if err != nil {
			return nil, fmt.Errorf("User Invalid")
		}
		if u.FirebaseUID != "" {
			return nil, fmt.Errorf("Need to Firebase Auth")
		}
	} else {
		// Firebase認証有り
		u, err := app.UserRepository.FindByFirebaseUID(ctx, f, user.FirebaseUID)
		if err != nil {
			return nil, fmt.Errorf("Firebase Auth Invalid")
		}

		user.ID = u.ID
	}

	if user.ID == "" {
		return nil, fmt.Errorf("UserID Invalid")
	}

	return NewGraphWithSetUserID(app, f, user.ID), nil
}

// NewGraphWithSetUserID Graphを作成（ログイン前）
func NewGraphWithSetUserID(app *Application, f *firestore.Client, uid string) *Graph {
	client := &Client{
		UUID:      &uuidgen.UUID{},
		Time:      &timegen.Time{},
		AuthToken: &authToken.AuthToken{},
	}

	return &Graph{
		UserID:          uid,
		FirestoreClient: f,
		App:             app,
		Client:          client,
	}
}

func GetNestCollectFields(ctx context.Context, cfs []graphql.CollectedField, columnName string) []graphql.CollectedField {
	for _, cf := range cfs {
		if cf.Name == columnName {
			return graphql.CollectFields(graphql.GetOperationContext(ctx), cf.Selections, nil)
		}
	}

	return []graphql.CollectedField{}
}

func GetNestCollectFieldArgumentValue(ctx context.Context, cfs []graphql.CollectedField, columnName string, argumentName string) string {
	for _, cf := range cfs {
		if cf.Name == columnName {
			for _, arg := range cf.Arguments {
				if arg.Name == argumentName {
					return arg.Value.String()
				}
			}

		}
	}

	return ""
}
