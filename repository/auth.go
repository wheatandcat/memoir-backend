package repository

import (
	"context"

	"cloud.google.com/go/firestore"
)

//go:generate moq -out=moq/auth.go -pkg=moqs . AuthRepositoryInterface

type AuthRepositoryInterface interface {
	Delete(ctx context.Context, f *firestore.Client, batch *firestore.WriteBatch, uid string)
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
func (re *AuthRepository) Delete(ctx context.Context, f *firestore.Client, batch *firestore.WriteBatch, uid string) {
	ref := f.Collection("auth").Doc(uid)
	batch.Delete(ref)
}
