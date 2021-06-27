package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/wheatandcat/memoir-backend/graph/model"
)

type PushTokenRepositoryInterface interface {
	Create(ctx context.Context, f *firestore.Client, userID string, i *model.PushToken) error
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

	return err
}
