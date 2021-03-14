package repository

import (
	"context"

	"cloud.google.com/go/firestore"

	"github.com/wheatandcat/memoir-backend/graph/model"
)

// UserRepositoryInterface is repository interface
type UserRepositoryInterface interface {
	Create(ctx context.Context, f *firestore.Client, u *model.User) error
	FindByUID(ctx context.Context, f *firestore.Client, uid string) (*model.User, error)
}

// UserRepository is repository for user
type UserRepository struct {
}

// NewUserRepository is Create new UserRepository
func NewUserRepository() UserRepositoryInterface {
	return &UserRepository{}
}

// Create ユーザーを作成する
func (re *UserRepository) Create(ctx context.Context, f *firestore.Client, u *model.User) error {
	_, err := f.Collection("users").Doc(u.ID).Set(ctx, u)

	return err
}

// FindByUID ユーザーIDから取得する
func (re *UserRepository) FindByUID(ctx context.Context, f *firestore.Client, uid string) (*model.User, error) {
	var u *model.User
	ds, err := f.Collection("users").Doc(uid).Get(ctx)
	if err != nil {
		return u, err
	}

	ds.DataTo(&u)

	return u, nil
}
