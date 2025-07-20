package graph

import (
	"context"

	"github.com/samber/lo"
	"github.com/wheatandcat/memoir-backend/graph/model"
	ce "github.com/wheatandcat/memoir-backend/usecase/custom_error"
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

	//既に同じデータが存在する場合は重複して登録しない
	items, err := g.App.PushTokenRepository.GetItems(ctx, g.FirestoreClient, g.UserID)
	if err != nil {
		return nil, ce.CustomError(err)
	}
	_, ok := lo.Find(items, func(item *model.PushToken) bool {
		return item.Token == input.Token && item.DeviceID == input.DeviceID
	})
	if ok {
		return i, nil
	}

	if err := g.App.PushTokenRepository.Create(ctx, g.FirestoreClient, g.UserID, i); err != nil {
		return nil, ce.CustomError(err)
	}

	return i, nil
}
