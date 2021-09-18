package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

	return errors.WithStack(err)
}

func GrpcErrorStatusCode(err error) codes.Code {
	return status.Code(errors.Cause(err))
}
