package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	ce "github.com/wheatandcat/memoir-backend/usecase/custom_error"
)

//go:generate moq -out=moq/auth.go -pkg=moqs . AuthRepositoryInterface

type AuthRepositoryInterface interface {
	Delete(ctx context.Context, f *firestore.Client, tx *firestore.Transaction, uid string) error
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
func (re *AuthRepository) Delete(ctx context.Context, f *firestore.Client, tx *firestore.Transaction, uid string) error {
	ref := f.Collection("auth").Doc(uid)
	err := tx.Delete(ref)
	if err != nil {
		return ce.CustomError(err)
	}
	return nil
}
