package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/pkg/errors"
	"github.com/wheatandcat/memoir-backend/graph/model"
	"google.golang.org/grpc/codes"
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

	return errors.WithStack(err)
}

func (re *PushTokenRepository) GetItems(ctx context.Context, f *firestore.Client, userID string) ([]*model.PushToken, error) {
	var items []*model.PushToken

	matchItem := getPushTokenCollection(f, userID).Documents(ctx)
	docs, err := matchItem.GetAll()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	for _, doc := range docs {
		var item *model.PushToken
		doc.DataTo(&item)

		items = append(items, item)
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
