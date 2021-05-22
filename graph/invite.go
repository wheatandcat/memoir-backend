package graph

import (
	"context"

	"github.com/wheatandcat/memoir-backend/graph/model"
)

// CreateInvite 招待を作成する
func (g *Graph) CreateInvite(ctx context.Context) (*model.Invite, error) {
	uuid := g.Client.UUID.Get()
	code := uuid[0:6]

	i := &model.Invite{
		Code:      code,
		UserID:    g.UserID,
		CreatedAt: g.Client.Time.Now(),
		UpdatedAt: g.Client.Time.Now(),
	}

	batch := g.FirestoreClient.Batch()
	g.App.InviteRepository.Create(ctx, g.FirestoreClient, batch, i)

	if _, err := batch.Commit(ctx); err != nil {
		return nil, err
	}

	return i, nil
}

// UpdateInvite 招待を更新する
func (g *Graph) UpdateInvite(ctx context.Context) (*model.Invite, error) {

	i, err := g.App.InviteRepository.FindByUserID(ctx, g.FirestoreClient, g.UserID)
	if err != nil {
		return nil, err
	}

	batch := g.FirestoreClient.Batch()
	g.App.InviteRepository.Delete(ctx, g.FirestoreClient, batch, i.Code)
	g.App.InviteRepository.Create(ctx, g.FirestoreClient, batch, i)

	if _, err := batch.Commit(ctx); err != nil {
		return nil, err
	}

	return g.CreateInvite(ctx)
}
