package graph

import (
	"context"
	"time"

	"github.com/wheatandcat/memoir-backend/graph/model"
)

// CreateItem アイテム作成
func (g *Graph) CreateItem(ctx context.Context, input *model.NewItem) (*model.Item, error) {
	i := &model.Item{
		ID:         g.Client.UUID.Get(),
		Title:      input.Title,
		Date:       input.Date,
		CategoryID: input.CategoryID,
		Like:       input.Like,
		Dislike:    input.Dislike,
		CreatedAt:  g.Client.Time.Now(),
		UpdatedAt:  g.Client.Time.Now(),
	}

	if err := g.App.ItemRepository.Create(ctx, g.FirestoreClient, g.UserID, i); err != nil {
		return nil, err
	}

	return i, nil
}

// UpdateItem アイテム更新
func (g *Graph) UpdateItem(ctx context.Context, input *model.UpdateItem) (*model.Item, error) {
	i := &model.Item{
		ID:        input.ID,
		UpdatedAt: g.Client.Time.Now(),
	}

	if err := g.App.ItemRepository.Update(ctx, g.FirestoreClient, g.UserID, input, i.UpdatedAt); err != nil {
		return nil, err
	}

	return i, nil
}

// DeleteItem アイテム削除
func (g *Graph) DeleteItem(ctx context.Context, input *model.DeleteItem) (*model.Item, error) {
	i := &model.Item{
		ID: input.ID,
	}

	if err := g.App.ItemRepository.Delete(ctx, g.FirestoreClient, g.UserID, input); err != nil {
		return nil, err
	}

	return i, nil
}

// GetItem アイテム取得
func (g *Graph) GetItem(ctx context.Context, id string) (*model.Item, error) {
	i, err := g.App.ItemRepository.GetItem(ctx, g.FirestoreClient, g.UserID, id)
	if err != nil {
		return nil, err
	}

	t := g.Client.Time

	i.Date = t.Location(i.Date)
	i.CreatedAt = t.Location(i.CreatedAt)
	i.UpdatedAt = t.Location(i.UpdatedAt)

	return i, nil
}

// GetItemsByDate 日付でアイテムを取得
func (g *Graph) GetItemsByDate(ctx context.Context, date time.Time) ([]*model.Item, error) {
	items, err := g.App.ItemRepository.GetItemsByDate(ctx, g.FirestoreClient, g.UserID, date)
	if err != nil {
		return nil, err
	}

	t := g.Client.Time

	for index, i := range items {
		items[index].Date = t.Location(i.Date)
		items[index].CreatedAt = t.Location(i.CreatedAt)
		items[index].UpdatedAt = t.Location(i.UpdatedAt)
	}

	return items, nil
}
