package graph

import (
	"context"
	"strings"

	"cloud.google.com/go/firestore"
	"github.com/wheatandcat/memoir-backend/graph/model"
	ce "github.com/wheatandcat/memoir-backend/usecase/custom_error"
)

// CreateInvite 招待を作成する
func (g *Graph) CreateInvite(ctx context.Context) (*model.Invite, error) {
	if !g.Client.AuthToken.Valid(ctx) {
		return nil, ce.CustomError(ce.NewInvalidAuthError("invalid authorization"))
	}

	i, err := g.App.InviteRepository.FindByUserID(ctx, g.FirestoreClient, g.UserID)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	if i.UserID == g.UserID {
		return nil, ce.CustomError(ce.NewRequestError(ce.CodeMyInviteCode, "自身の招待コードです"))
	}

	uuid := g.Client.UUID.Get()
	code := strings.ToUpper(uuid[0:8])

	i = &model.Invite{
		Code:      code,
		UserID:    g.UserID,
		CreatedAt: g.Client.Time.Now(),
		UpdatedAt: g.Client.Time.Now(),
	}

	err = g.FirestoreClient.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		if err := g.App.InviteRepository.Create(ctx, g.FirestoreClient, tx, i); err != nil {
			return ce.CustomError(err)
		}

		return nil
	})
	if err != nil {
		return nil, ce.CustomError(err)
	}

	return i, nil
}

// UpdateInvite 招待を更新する
func (g *Graph) UpdateInvite(ctx context.Context) (*model.Invite, error) {
	if !g.Client.AuthToken.Valid(ctx) {
		return nil, ce.CustomError(ce.NewInvalidAuthError("invalid authorization"))
	}

	i, err := g.App.InviteRepository.FindByUserID(ctx, g.FirestoreClient, g.UserID)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	uuid := g.Client.UUID.Get()
	code := strings.ToUpper(uuid[0:8])

	err = g.FirestoreClient.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		if err := g.App.InviteRepository.Delete(ctx, g.FirestoreClient, tx, i.Code); err != nil {
			return err
		}

		i.Code = code
		i.UpdatedAt = g.Client.Time.Now()
		if err := g.App.InviteRepository.Create(ctx, g.FirestoreClient, tx, i); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return nil, ce.CustomError(err)
	}

	return i, nil
}

// GetInviteByUseID ユーザーIDから招待を取得する
func (g *Graph) GetInviteByUseID(ctx context.Context) (*model.Invite, error) {
	if !g.Client.AuthToken.Valid(ctx) {
		return nil, ce.CustomError(ce.NewInvalidAuthError("invalid authorization"))
	}

	i, err := g.App.InviteRepository.FindByUserID(ctx, g.FirestoreClient, g.UserID)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	return i, nil
}

// GetInviteByCode コードから招待を取得する
func (g *Graph) GetInviteByCode(ctx context.Context, code string) (*model.User, error) {
	if !g.Client.AuthToken.Valid(ctx) {
		return nil, ce.CustomError(ce.NewInvalidAuthError("invalid authorization"))
	}

	i, err := g.App.InviteRepository.Find(ctx, g.FirestoreClient, code)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	if i.UserID == "" {
		return nil, ce.CustomError(ce.NewNotFoundError("招待コードが見つかりません"))
	}

	u, err := g.App.UserRepository.FindByUID(ctx, g.FirestoreClient, i.UserID)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	return u, nil
}
