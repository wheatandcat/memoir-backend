package repository

import (
	"context"

	"cloud.google.com/go/firestore"
)

type CommonRepositoryInterface interface {
	Commit(ctx context.Context, batch *firestore.WriteBatch) error
}

type CommonRepository struct {
}

func NewCommonRepository() CommonRepositoryInterface {
	return &CommonRepository{}
}

// Commit コミットする
func (re *CommonRepository) Commit(ctx context.Context, batch *firestore.WriteBatch) error {
	_, err := batch.Commit(ctx)

	return err
}
