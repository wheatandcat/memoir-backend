package graph

import (
	"context"

	"github.com/wheatandcat/memoir-backend/graph/model"
)

// PushToken トークン作成
func (g *Graph) CreatePushToken(ctx context.Context, input *model.NewPushToken) (*model.PushToken, error) {
	i := &model.PushToken{
		UserID:    g.UserID,
		Token:     input.Token,
		DeviceID:  input.DeviceID,
		CreatedAt: g.Client.Time.Now(),
		UpdatedAt: g.Client.Time.Now(),
	}

	if err := g.App.PushTokenRepository.Create(ctx, g.FirestoreClient, g.UserID, i); err != nil {
		return nil, err
	}

	return i, nil
}
