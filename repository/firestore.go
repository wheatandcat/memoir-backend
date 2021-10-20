package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	ce "github.com/wheatandcat/memoir-backend/usecase/custom_error"
)

// FirestoreClient Firestore Client
func FirestoreClient(ctx context.Context) (*firestore.Client, error) {
	app, err := FirebaseApp(ctx)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	return app.Firestore(ctx)
}
