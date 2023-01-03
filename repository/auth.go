package repository

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	ce "github.com/wheatandcat/memoir-backend/usecase/custom_error"
)

//go:generate moq -out=moq/auth.go -pkg=moqs . AuthRepositoryInterface

type AuthRepositoryInterface interface {
	Delete(ctx context.Context, f *firestore.Client, batch *firestore.BulkWriter, uid string) error
}

type Auth struct {
	ID string
}

type AuthRepository struct {
}

func NewAuthRepository() AuthRepositoryInterface {
	return &AuthRepository{}
}

// Delete Authを削除する
func (re *AuthRepository) Delete(ctx context.Context, f *firestore.Client, batch *firestore.BulkWriter, uid string) error {
	ref := f.Collection("auth").Doc(uid)
	j, err := batch.Delete(ref)
	if err != nil {
		return ce.CustomError(err)
	}
	if j == nil {
		return ce.CustomError(fmt.Errorf("BulkWriter: got nil"))
	}
	if _, err := j.Results(); err != nil {
		return ce.CustomError(err)
	}
	return nil
}
