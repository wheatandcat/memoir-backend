package graph

import (
	"context"
	"fmt"

	"github.com/wheatandcat/memoir-backend/graph/model"
)

// CreateInvite 招待を作成する
func (g *Graph) CreateInvite(ctx context.Context) (*model.Invite, error) {
	if !g.Client.AuthToken.Valid(ctx) {
		return nil, fmt.Errorf("Invalid Authorization")
	}

	i, err := g.App.InviteRepository.FindByUserID(ctx, g.FirestoreClient, g.UserID)
	if err != nil {
		return nil, err
	}
	if i.Code != "" {
		return nil, fmt.Errorf("既に招待コードは作成済みです")
	}

	uuid := g.Client.UUID.Get()
	code := uuid[0:8]

	i = &model.Invite{
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
	if !g.Client.AuthToken.Valid(ctx) {
		return nil, fmt.Errorf("Invalid Authorization")
	}

	i, err := g.App.InviteRepository.FindByUserID(ctx, g.FirestoreClient, g.UserID)
	if err != nil {
		return nil, err
	}

	uuid := g.Client.UUID.Get()
	code := uuid[0:8]

	batch := g.FirestoreClient.Batch()
	g.App.InviteRepository.Delete(ctx, g.FirestoreClient, batch, i.Code)

	i.Code = code
	i.UpdatedAt = g.Client.Time.Now()
	g.App.InviteRepository.Create(ctx, g.FirestoreClient, batch, i)

	if _, err := batch.Commit(ctx); err != nil {
		return nil, err
	}

	return i, nil
}

// GetInviteByUseID ユーザーIDから招待を取得する
func (g *Graph) GetInviteByUseID(ctx context.Context) (*model.Invite, error) {
	if !g.Client.AuthToken.Valid(ctx) {
		return nil, fmt.Errorf("Invalid Authorization")
	}

	i, err := g.App.InviteRepository.FindByUserID(ctx, g.FirestoreClient, g.UserID)
	if err != nil {
		return nil, err
	}

	return i, nil
}

// GetInviteByCode コードから招待を取得する
func (g *Graph) GetInviteByCode(ctx context.Context, code string) (*model.User, error) {
	if !g.Client.AuthToken.Valid(ctx) {
		return nil, fmt.Errorf("Invalid Authorization")
	}

	i, err := g.App.InviteRepository.Find(ctx, g.FirestoreClient, code)
	if err != nil {
		return nil, err
	}

	if i.UserID == "" {
		return nil, fmt.Errorf("招待コードが見つかりません")
	}

	u, err := g.App.UserRepository.FindByUID(ctx, g.FirestoreClient, i.UserID)
	if err != nil {
		return nil, err
	}

	return u, nil
}