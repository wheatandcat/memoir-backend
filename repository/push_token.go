package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"

	"github.com/wheatandcat/memoir-backend/graph/model"
	ce "github.com/wheatandcat/memoir-backend/usecase/custom_error"
)

//go:generate moq -out=moq/push_token.go -pkg=moqs . PushTokenRepositoryInterface

type PushTokenRepositoryInterface interface {
	Create(ctx context.Context, f *firestore.Client, userID string, i *model.PushToken) error
	GetItems(ctx context.Context, f *firestore.Client, userID string) ([]*model.PushToken, error)
	GetTokens(ctx context.Context, f *firestore.Client, userID string) []string
}

type PushTokenRepository struct {
}

// NewPushTokenRepository is Create new PushTokenRepository
func NewPushTokenRepository() PushTokenRepositoryInterface {
	return &PushTokenRepository{}
}

func getPushTokenCollection(f *firestore.Client, userID string) *firestore.CollectionRef {
	return f.Collection("users/" + userID + "/pushToken")
}

func (re *PushTokenRepository) Create(ctx context.Context, f *firestore.Client, userID string, i *model.PushToken) error {
	_, err := getPushTokenCollection(f, userID).Doc(i.DeviceID).Set(ctx, i)

	return ce.CustomError(err)
}

func (re *PushTokenRepository) GetItems(ctx context.Context, f *firestore.Client, userID string) ([]*model.PushToken, error) {
	matchItem := getPushTokenCollection(f, userID).Documents(ctx)
	docs, err := matchItem.GetAll()
	if err != nil {
		return nil, ce.CustomError(err)
	}

	items := make([]*model.PushToken, len(docs))
	for i, doc := range docs {
		var item *model.PushToken
		if err = doc.DataTo(&item); err != nil {
			return items, ce.CustomError(err)
		}
		items[i] = item
	}

	return items, nil
}

func (re *PushTokenRepository) GetTokens(ctx context.Context, f *firestore.Client, userID string) []string {
	tokens := []string{}
	items, err := re.GetItems(ctx, f, userID)
	if GrpcErrorStatusCode(err) == codes.InvalidArgument || GrpcErrorStatusCode(err) == codes.NotFound {
		return tokens
	}

	for _, item := range items {
		tokens = append(tokens, item.Token)
	}

	return tokens

}
