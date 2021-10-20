package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/pkg/errors"
	ce "github.com/wheatandcat/memoir-backend/usecase/custom_error"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//go:generate moq -out=moq/common.go -pkg=moqs . CommonRepositoryInterface

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

	return ce.CustomError(err)
}

func GrpcErrorStatusCode(err error) codes.Code {
	return status.Code(errors.Cause(err))
}
