package graph

import (
	"context"
	"strings"
	"time"

	"github.com/wheatandcat/memoir-backend/graph/model"
	"github.com/wheatandcat/memoir-backend/repository"
	ce "github.com/wheatandcat/memoir-backend/usecase/custom_error"
	"google.golang.org/grpc/codes"
)

// CreateItem アイテム作成
func (g *Graph) CreateItem(ctx context.Context, input *model.NewItem) (*model.Item, error) {
	i := &model.Item{
		ID:         g.Client.UUID.Get(),
		UserID:     g.UserID,
		Title:      input.Title,
		Date:       input.Date,
		CategoryID: input.CategoryID,
		Like:       input.Like,
		Dislike:    input.Dislike,
		CreatedAt:  g.Client.Time.Now(),
		UpdatedAt:  g.Client.Time.Now(),
	}

	if err := g.App.ItemRepository.Create(ctx, g.FirestoreClient, g.UserID, i); err != nil {
		return nil, ce.CustomError(err)
	}

	return i, nil
}

// UpdateItem アイテム更新
func (g *Graph) UpdateItem(ctx context.Context, input *model.UpdateItem) (*model.Item, error) {
	i := &model.Item{
		ID:        input.ID,
		UserID:    g.UserID,
		Date:      *input.Date,
		UpdatedAt: g.Client.Time.Now(),
	}

	if err := g.App.ItemRepository.Update(ctx, g.FirestoreClient, g.UserID, input, i.UpdatedAt); err != nil {
		return nil, ce.CustomError(err)
	}

	return i, nil
}

// DeleteItem アイテム削除
func (g *Graph) DeleteItem(ctx context.Context, input *model.DeleteItem) (*model.Item, error) {
	i := &model.Item{
		ID: input.ID,
	}

	if err := g.App.ItemRepository.Delete(ctx, g.FirestoreClient, g.UserID, input); err != nil {
		return nil, ce.CustomError(err)
	}

	return i, nil
}

// GetItem アイテム取得
func (g *Graph) GetItem(ctx context.Context, id string) (*model.Item, error) {
	i, err := g.App.ItemRepository.GetItem(ctx, g.FirestoreClient, g.UserID, id)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	t := g.Client.Time

	i.Date = t.Location(i.Date)
	i.CreatedAt = t.Location(i.CreatedAt)
	i.UpdatedAt = t.Location(i.UpdatedAt)

	return i, nil
}

// GetItemsInDate 日付でアイテムを取得
func (g *Graph) GetItemsInDate(ctx context.Context, date time.Time) ([]*model.Item, error) {
	items, err := g.App.ItemRepository.GetItemsInDate(ctx, g.FirestoreClient, g.UserID, date)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	t := g.Client.Time

	for index, i := range items {
		items[index].Date = t.Location(i.Date)
		items[index].CreatedAt = t.Location(i.CreatedAt)
		items[index].UpdatedAt = t.Location(i.UpdatedAt)
	}

	return items, nil
}

// GetItemsInPeriod 期間でアイテムを取得
func (g *Graph) GetItemsInPeriod(ctx context.Context, input model.InputItemsInPeriod) (*model.ItemsInPeriod, error) {
	t := g.Client.Time
	userID := []string{g.UserID}

	rrs, err := g.App.RelationshipRepository.FindByFollowedID(ctx, g.FirestoreClient, g.UserID, 5, repository.RelationshipCursor{
		FollowerID: "",
		FollowedID: "",
	})
	if err == nil {
		for _, rr := range rrs {
			userID = append(userID, rr.FollowerID)
		}
	} else {
		if GrpcErrorStatusCode(err) != codes.NotFound {
			return nil, ce.CustomError(err)
		}
	}

	cursor := repository.ItemsInPeriodCursor{
		ID:     "",
		UserID: "",
	}
	cursorData := strings.Split(*input.After, "/")
	if len(cursorData) > 1 {
		cursor = repository.ItemsInPeriodCursor{
			UserID: cursorData[0],
			ID:     cursorData[1],
		}
	}

	if len(input.UserIDList) > 0 {
		tUserID := []string{}
		tUserIDList := []string{}
		for _, id := range input.UserIDList {
			tUserIDList = append(tUserIDList, *id)
		}
		// UserIDListの設定がある場合は、UserIDをフィルタリングする
		for _, id := range userID {
			if Contains(tUserIDList, id) {
				tUserID = append(tUserID, id)
			}
		}
		userID = tUserID
	}

	dislike := false
	if input.Dislike != nil {
		dislike = *input.Dislike
	}
	like := false
	if input.Like != nil {
		like = *input.Like
	}
	ci := 0
	if input.CategoryID != nil {
		ci = *input.CategoryID
	}

	sip := repository.SearchItemParam{
		UserID:     userID,
		StartDate:  input.StartDate,
		EndDate:    input.EndDate,
		Like:       like,
		Dislike:    dislike,
		CategoryID: ci,
	}

	items, err := g.App.ItemRepository.GetItemUserMultipleInPeriod(ctx, g.FirestoreClient, sip, input.First, cursor)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	ibpes := make([]*model.ItemsInPeriodEdge, len(items))
	for index, i := range items {
		items[index].Date = t.Location(i.Date)
		items[index].CreatedAt = t.Location(i.CreatedAt)
		items[index].UpdatedAt = t.Location(i.UpdatedAt)

		ibpes[index] = &model.ItemsInPeriodEdge{
			Node:   items[index],
			Cursor: i.UserID + "/" + i.ID,
		}
	}

	pi := &model.PageInfo{
		HasNextPage: false,
		EndCursor:   "",
	}
	if len(ibpes) > 0 {
		pi.HasNextPage = input.First == len(items)
		pi.EndCursor = ibpes[len(items)-1].Cursor
	}
	ibp := &model.ItemsInPeriod{
		Edges:    ibpes,
		PageInfo: pi,
	}

	return ibp, nil
}
