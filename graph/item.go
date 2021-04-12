package graph

import (
	"context"
	"strings"
	"time"

	"github.com/wheatandcat/memoir-backend/graph/model"
	"github.com/wheatandcat/memoir-backend/repository"
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

// GetItemsByPeriod 期間でアイテムを取得
func (g *Graph) GetItemsByPeriod(ctx context.Context, input model.InputItemsByPeriod) (*model.ItemsByPeriod, error) {
	t := g.Client.Time

	cursor := repository.ItemsByPeriodCursor{
		Date:      time.Now(),
		CreatedAt: time.Now(),
		ID:        "",
	}
	cursorDate := strings.Split(*input.After, "/")
	if len(cursorDate) > 1 {
		cursor = repository.ItemsByPeriodCursor{
			Date:      t.ParseInLocationTimezone(cursorDate[0]),
			CreatedAt: t.ParseInLocationTimezone(cursorDate[1]),
			ID:        cursorDate[2],
		}
	}

	items, err := g.App.ItemRepository.GetItemsByPeriod(ctx, g.FirestoreClient, g.UserID, input.StartDate, input.EndDate, input.First, cursor)
	if err != nil {
		return nil, err
	}

	var ibpes []*model.ItemsByPeriodEdge
	for index, i := range items {
		items[index].Date = t.Location(i.Date)
		items[index].CreatedAt = t.Location(i.CreatedAt)
		items[index].UpdatedAt = t.Location(i.UpdatedAt)

		ibpes = append(ibpes, &model.ItemsByPeriodEdge{
			Node:   items[index],
			Cursor: i.Date.Format("2006-01-02T15:04:05+09:00") + "/" + i.CreatedAt.Format("2006-01-02T15:04:05+09:00") + "/" + i.ID,
		})
	}

	pi := &model.PageInfo{
		HasNextPage: false,
		EndCursor:   "",
	}
	if len(ibpes) > 0 {
		pi.HasNextPage = input.First == len(items)
		pi.EndCursor = ibpes[len(items)-1].Cursor
	}
	ibp := &model.ItemsByPeriod{
		Edges:    ibpes,
		PageInfo: pi,
	}

	return ibp, nil
}
