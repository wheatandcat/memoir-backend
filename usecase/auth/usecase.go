package auth

//go:generate moq -out=moq/auth.go -pkg=moq . UseCase

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/wheatandcat/memoir-backend/graph/model"
	"github.com/wheatandcat/memoir-backend/repository"
)

type UseCase interface {
	CreateAuthUser(ctx context.Context, f *firestore.Client, input *model.NewAuthUser, u *repository.User, mu *model.AuthUser) error
}

type useCaseImpl struct {
}

func New() UseCase {
	return &useCaseImpl{}
}
