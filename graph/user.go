package graph

import (
	"context"

	"github.com/wheatandcat/memoir-backend/graph/model"
	"github.com/wheatandcat/memoir-backend/repository"
	ce "github.com/wheatandcat/memoir-backend/usecase/custom_error"
)

// CreateUser ユーザー作成
func (g *Graph) CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	u := &model.User{
		ID:        input.ID,
		CreatedAt: g.Client.Time.Now(),
		UpdatedAt: g.Client.Time.Now(),
	}

	if err := g.App.UserRepository.Create(ctx, g.FirestoreClient, u); err != nil {
		return nil, ce.CustomError(err)
	}

	return u, nil
}

// CreateAuthUser 認証済みユーザーを作成
func (g *Graph) CreateAuthUser(ctx context.Context, input *model.NewAuthUser) (*model.AuthUser, error) {
	if !g.Client.AuthToken.Valid(ctx) {
		return nil, ce.CustomError(ce.NewInvalidAuthError("invalid authorization"))
	}

	u := &repository.User{
		ID:          input.ID,
		FirebaseUID: g.Client.AuthToken.Get(ctx),
		UpdatedAt:   g.Client.Time.Now(),
	}

	mu := &model.AuthUser{
		ID:        input.ID,
		CreatedAt: g.Client.Time.Now(),
		UpdatedAt: g.Client.Time.Now(),
	}

	exist, err := g.App.UserRepository.ExistByFirebaseUID(ctx, g.FirestoreClient, u.FirebaseUID)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	if exist {
		// 既にユーザー作成済みの場合は更新しないで完了
		return mu, nil
	}

	displayName, err := g.App.AuthUseCase.CreateAuthUser(ctx, g.FirestoreClient, input, u, mu)
	mu.DisplayName = displayName

	if err != nil {
		exist, err := g.App.UserRepository.ExistByFirebaseUID(ctx, g.FirestoreClient, u.FirebaseUID)
		if err != nil {
			return nil, ce.CustomError(err)
		}

		if exist {
			// 既にユーザー作成済みの場合は更新しないで完了
			return mu, nil
		}

		return nil, ce.CustomError(err)
	}

	return mu, nil
}

// GetUser ユーザー取得
func (g *Graph) GetUser(ctx context.Context) (*model.User, error) {
	u, err := g.App.UserRepository.FindByUID(ctx, g.FirestoreClient, g.UserID)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	return u, nil

}

// ExistAuthUser 認証ユーザーが存在するか判定する
func (g *Graph) ExistAuthUser(ctx context.Context) (*model.ExistAuthUser, error) {
	rau := &model.ExistAuthUser{}

	exist, err := g.App.UserRepository.ExistByFirebaseUID(ctx, g.FirestoreClient, g.Client.AuthToken.Get(ctx))
	if err != nil {
		return nil, ce.CustomError(err)
	}

	rau.Exist = exist

	return rau, nil
}

// UpdateUser ユーザーを更新
func (g *Graph) UpdateUser(ctx context.Context, input *model.UpdateUser) (*model.User, error) {
	u := &model.User{
		ID:          g.UserID,
		DisplayName: input.DisplayName,
		Image:       input.Image,
		UpdatedAt:   g.Client.Time.Now(),
	}

	err := g.App.UserRepository.Update(ctx, g.FirestoreClient, u)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	return u, nil
}
