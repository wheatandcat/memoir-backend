package graph

import (
	"context"

	"github.com/wheatandcat/memoir-backend/graph/model"
)

// CreateUser ユーザー作成
func (g *Graph) CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	u := &model.User{
		ID:        input.ID,
		CreatedAt: g.Client.Time.Now(),
		UpdatedAt: g.Client.Time.Now(),
	}

	if err := g.App.UserRepository.Create(ctx, g.FirestoreClient, u); err != nil {
		return nil, err
	}

	return u, nil

}

// GetUser ユーザー取得
func (g *Graph) GetUser(ctx context.Context) (*model.User, error) {
	u, err := g.App.UserRepository.FindByUID(ctx, g.FirestoreClient, g.UserID)
	if err != nil {
		return nil, err
	}

	return u, nil

}
