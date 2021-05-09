package graph

import (
	"context"
	"fmt"

	"github.com/wheatandcat/memoir-backend/graph/model"
	"github.com/wheatandcat/memoir-backend/repository"
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

// CreateAuthUser 認証済みユーザーを作成
func (g *Graph) CreateAuthUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	if !g.Client.AuthToken.Valid(ctx) {
		return nil, fmt.Errorf("Invalid Authorization")
	}

	u := &repository.User{
		ID:          input.ID,
		FirebaseUID: g.Client.AuthToken.Get(ctx),
		UpdatedAt:   g.Client.Time.Now(),
	}

	mu := &model.User{
		ID: input.ID,
	}

	exist, err := g.App.UserRepository.ExistByFirebaseUID(ctx, g.FirestoreClient, u.FirebaseUID)
	if err != nil {
		return nil, err
	}

	if exist == true {
		// 既にユーザー作成済みの場合は更新しないで完了
		return mu, nil
	}

	if err = g.App.UserRepository.UpdateFirebaseUID(ctx, g.FirestoreClient, u); err != nil {
		return nil, err
	}

	return mu, nil
}

// GetUser ユーザー取得
func (g *Graph) GetUser(ctx context.Context) (*model.User, error) {
	u, err := g.App.UserRepository.FindByUID(ctx, g.FirestoreClient, g.UserID)
	if err != nil {
		return nil, err
	}

	return u, nil

}

// UpdateUser ユーザーを更新
func (g *Graph) UpdateUser(ctx context.Context, input *model.UpdateUser) (*model.User, error) {

	if !g.Client.AuthToken.Valid(ctx) {
		return nil, fmt.Errorf("Invalid Authorization")
	}

	u := &model.User{
		ID:          g.UserID,
		DisplayName: input.DisplayName,
		UpdatedAt:   g.Client.Time.Now(),
	}

	err := g.App.UserRepository.Update(ctx, g.FirestoreClient, u)
	if err != nil {
		return nil, err
	}

	return u, nil
}
