package repository

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"

	firestorepb "cloud.google.com/go/firestore/apiv1/firestorepb"
	"github.com/wheatandcat/memoir-backend/graph/model"
	ce "github.com/wheatandcat/memoir-backend/usecase/custom_error"
)

//go:generate moq -out=moq/item.go -pkg=moqs . ItemRepositoryInterface

type ItemRepositoryInterface interface {
	Create(ctx context.Context, f *firestore.Client, userID string, i *model.Item) error
	Update(ctx context.Context, f *firestore.Client, userID string, i *model.UpdateItem, updatedAt time.Time) error
	Delete(ctx context.Context, f *firestore.Client, userID string, i *model.DeleteItem) error
	GetItem(ctx context.Context, f *firestore.Client, userID string, id string) (*model.Item, error)
	GetItemsInDate(ctx context.Context, f *firestore.Client, userID string, date time.Time) ([]*model.Item, error)
	GetItemsInPeriod(ctx context.Context, f *firestore.Client, userID string, stertDate time.Time, endDate time.Time, first int, cursor ItemsInPeriodCursor) ([]*model.Item, error)
	GetItemUserMultipleInPeriod(ctx context.Context, f *firestore.Client, sip SearchItemParam, first int, cursor ItemsInPeriodCursor) ([]*model.Item, error)
	GetCountUserMultipleInPeriod(ctx context.Context, f *firestore.Client, sip SearchItemParam) (int, error)
}

type ItemKey struct {
	UserID string
}

type ItemRepository struct {
}

func NewItemRepository() ItemRepositoryInterface {
	return &ItemRepository{}
}

// getItemCollection アイテムのコレクションを取得する
func getItemCollection(f *firestore.Client, userID string) *firestore.CollectionRef {
	return f.Collection("users/" + userID + "/items")
}

// Create アイテムを作成する
func (re *ItemRepository) Create(ctx context.Context, f *firestore.Client, userID string, i *model.Item) error {
	_, err := getItemCollection(f, userID).Doc(i.ID).Set(ctx, i)

	return ce.CustomError(err)
}

// Update アイテムを更新する
func (re *ItemRepository) Update(ctx context.Context, f *firestore.Client, userID string, i *model.UpdateItem, updatedAt time.Time) error {
	var u []firestore.Update
	if i.Title != nil {
		u = append(u, firestore.Update{Path: "Title", Value: i.Title})
	}
	if i.Date != nil {
		u = append(u, firestore.Update{Path: "Date", Value: i.Date})
	}
	if i.CategoryID != nil {
		u = append(u, firestore.Update{Path: "CategoryID", Value: i.CategoryID})
	}
	if i.Like != nil {
		u = append(u, firestore.Update{Path: "Like", Value: i.Like})
	}
	if i.Dislike != nil {
		u = append(u, firestore.Update{Path: "Dislike", Value: i.Dislike})
	}
	u = append(u, firestore.Update{Path: "UpdatedAt", Value: updatedAt})

	_, err := getItemCollection(f, userID).Doc(i.ID).Update(ctx, u)

	return ce.CustomError(err)
}

// Delete アイテムを削除する
func (re *ItemRepository) Delete(ctx context.Context, f *firestore.Client, userID string, i *model.DeleteItem) error {
	_, err := getItemCollection(f, userID).Doc(i.ID).Delete(ctx)
	return ce.CustomError(err)
}

// GetItem アイテムを取得する
func (re *ItemRepository) GetItem(ctx context.Context, f *firestore.Client, userID string, id string) (*model.Item, error) {
	var i *model.Item

	ds, err := getItemCollection(f, userID).Doc(id).Get(ctx)
	if err != nil {
		return i, ce.CustomError(err)
	}

	if err = ds.DataTo(&i); err != nil {
		return nil, ce.CustomError(err)
	}

	return i, nil
}

// GetItemsInDate 日付でアイテムを取得する
func (re *ItemRepository) GetItemsInDate(ctx context.Context, f *firestore.Client, userID string, date time.Time) ([]*model.Item, error) {
	matchItem := getItemCollection(f, userID).Where("Date", "==", date).OrderBy("CreatedAt", firestore.Desc).Documents(ctx)
	docs, err := matchItem.GetAll()
	if err != nil {
		return nil, ce.CustomError(err)
	}

	items := make([]*model.Item, len(docs))
	for i, doc := range docs {
		var item *model.Item
		if err = doc.DataTo(&item); err != nil {
			return items, ce.CustomError(err)
		}

		items[i] = item
	}

	return items, nil
}

type ItemsInPeriodCursor struct {
	ID     string
	UserID string
}

// GetItemsInPeriod 期間でアイテムを取得する
func (re *ItemRepository) GetItemsInPeriod(ctx context.Context, f *firestore.Client, userID string, startDate time.Time, endDate time.Time, first int, cursor ItemsInPeriodCursor) ([]*model.Item, error) {
	query := getItemCollection(f, userID).Where("Date", ">=", startDate).Where("Date", "<=", endDate).OrderBy("Date", firestore.Asc).OrderBy("CreatedAt", firestore.Asc).OrderBy("ID", firestore.Asc)

	if cursor.ID != "" {
		ds, err := getItemCollection(f, cursor.UserID).Doc(cursor.ID).Get(ctx)
		if err != nil {
			return nil, ce.CustomError(err)
		}

		query = query.StartAfter(ds)
	}

	matchItem := query.Limit(first).Documents(ctx)
	docs, err := matchItem.GetAll()

	if err != nil {
		return nil, ce.CustomError(err)
	}

	items := make([]*model.Item, len(docs))

	for i, doc := range docs {
		var item *model.Item
		if err = doc.DataTo(&item); err != nil {
			return items, ce.CustomError(err)
		}

		items[i] = item
	}

	return items, nil
}

type SearchItemParam struct {
	UserID     []string
	StartDate  time.Time
	EndDate    time.Time
	Like       bool
	Dislike    bool
	CategoryID int
}

func getQueryUserMultipleInPeriod(f *firestore.Client, sip SearchItemParam) firestore.Query {
	query := f.CollectionGroup("items").Where("UserID", "in", sip.UserID).Where("Date", ">=", sip.StartDate).Where("Date", "<=", sip.EndDate)

	if sip.Like && !sip.Dislike {
		query = query.Where("Like", "==", true).Where("Dislike", "==", false)
	} else if !sip.Like && sip.Dislike {
		query = query.Where("Like", "==", false).Where("Dislike", "==", true)
	}

	if sip.CategoryID != 0 {
		query = query.Where("CategoryID", "==", sip.CategoryID)
	}

	query = query.OrderBy("Date", firestore.Asc).OrderBy("CreatedAt", firestore.Asc)

	return query
}

// GetItemUserMultipleInPeriod 期間でアイテムを取得する
func (re *ItemRepository) GetItemUserMultipleInPeriod(ctx context.Context, f *firestore.Client, sip SearchItemParam, first int, cursor ItemsInPeriodCursor) ([]*model.Item, error) {
	query := getQueryUserMultipleInPeriod(f, sip)

	if cursor.ID != "" {
		ds, err := getItemCollection(f, cursor.UserID).Doc(cursor.ID).Get(ctx)
		if err != nil {
			return nil, ce.CustomError(err)
		}

		query = query.StartAfter(ds)
	}

	matchItem := query.Limit(first).Documents(ctx)
	docs, err := matchItem.GetAll()

	if err != nil {
		return nil, ce.CustomError(err)
	}

	items := make([]*model.Item, len(docs))

	for i, doc := range docs {
		var item *model.Item
		if err = doc.DataTo(&item); err != nil {
			return items, ce.CustomError(err)
		}
		items[i] = item
	}

	return items, nil
}

// GetCountUserMultipleInPeriod 期間でアイテムの総数を取得する
func (re *ItemRepository) GetCountUserMultipleInPeriod(ctx context.Context, f *firestore.Client, sip SearchItemParam) (int, error) {
	query := getQueryUserMultipleInPeriod(f, sip)

	alias := "items"
	ar, err := query.NewAggregationQuery().WithCount(alias).Get(ctx)
	if err != nil {
		return 0, ce.CustomError(err)
	}

	count, ok := ar[alias]
	if !ok {
		return 0, nil
	}

	ce := count.(*firestorepb.Value)

	return int(ce.GetIntegerValue()), nil
}
