package repository

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/pkg/errors"
	"github.com/wheatandcat/memoir-backend/graph/model"
)

type ItemRepositoryInterface interface {
	Create(ctx context.Context, f *firestore.Client, userID string, i *model.Item) error
	Update(ctx context.Context, f *firestore.Client, userID string, i *model.UpdateItem, updatedAt time.Time) error
	Delete(ctx context.Context, f *firestore.Client, userID string, i *model.DeleteItem) error
	GetItem(ctx context.Context, f *firestore.Client, userID string, id string) (*model.Item, error)
	GetItemsInDate(ctx context.Context, f *firestore.Client, userID string, date time.Time) ([]*model.Item, error)
	GetItemsInPeriod(ctx context.Context, f *firestore.Client, userID string, stertDate time.Time, endDate time.Time, first int, cursor ItemsInPeriodCursor) ([]*model.Item, error)
	GetItemUserMultipleInPeriod(ctx context.Context, f *firestore.Client, userID []string, stertDate time.Time, endDate time.Time, first int, cursor ItemsInPeriodCursor) ([]*model.Item, error)
}

// ItemKey is item key
type ItemKey struct {
	UserID string
}

// ItemRepository is repository for item
type ItemRepository struct {
}

// NewItemRepository is Create new ItemRepository
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

	return errors.WithStack(err)
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

	return errors.WithStack(err)
}

// Delete アイテムを削除する
func (re *ItemRepository) Delete(ctx context.Context, f *firestore.Client, userID string, i *model.DeleteItem) error {
	_, err := getItemCollection(f, userID).Doc(i.ID).Delete(ctx)
	return err
}

// GetItem アイテムを取得する
func (re *ItemRepository) GetItem(ctx context.Context, f *firestore.Client, userID string, id string) (*model.Item, error) {
	var i *model.Item

	ds, err := getItemCollection(f, userID).Doc(id).Get(ctx)
	if err != nil {
		return i, errors.WithStack(err)
	}

	ds.DataTo(&i)

	return i, nil
}

// GetItemsInDate 日付でアイテムを取得する
func (re *ItemRepository) GetItemsInDate(ctx context.Context, f *firestore.Client, userID string, date time.Time) ([]*model.Item, error) {
	var items []*model.Item

	matchItem := getItemCollection(f, userID).Where("Date", "==", date).OrderBy("CreatedAt", firestore.Desc).Documents(ctx)
	docs, err := matchItem.GetAll()
	if err != nil {
		return nil, errors.WithStack(err)
	}

	for _, doc := range docs {
		var item *model.Item
		doc.DataTo(&item)

		items = append(items, item)
	}

	return items, nil
}

type ItemsInPeriodCursor struct {
	ID     string
	UserID string
}

// GetItemsInPeriod 期間でアイテムを取得する
func (re *ItemRepository) GetItemsInPeriod(ctx context.Context, f *firestore.Client, userID string, startDate time.Time, endDate time.Time, first int, cursor ItemsInPeriodCursor) ([]*model.Item, error) {
	var items []*model.Item
	query := getItemCollection(f, userID).Where("Date", ">=", startDate).Where("Date", "<=", endDate).OrderBy("Date", firestore.Asc).OrderBy("CreatedAt", firestore.Asc).OrderBy("ID", firestore.Asc)

	if cursor.ID != "" {
		ds, err := getItemCollection(f, cursor.UserID).Doc(cursor.ID).Get(ctx)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		query = query.StartAfter(ds)
	}

	matchItem := query.Limit(first).Documents(ctx)
	docs, err := matchItem.GetAll()

	if err != nil {
		return nil, errors.WithStack(err)
	}

	for _, doc := range docs {
		var item *model.Item
		doc.DataTo(&item)

		items = append(items, item)
	}

	return items, nil
}

// GetItemUserMultipleInPeriod 期間でアイテムを取得する
func (re *ItemRepository) GetItemUserMultipleInPeriod(ctx context.Context, f *firestore.Client, userID []string, startDate time.Time, endDate time.Time, first int, cursor ItemsInPeriodCursor) ([]*model.Item, error) {
	var items []*model.Item

	query := f.CollectionGroup("items").Where("UserID", "in", userID).Where("Date", ">=", startDate).Where("Date", "<=", endDate).OrderBy("Date", firestore.Asc).OrderBy("CreatedAt", firestore.Asc)

	if cursor.ID != "" {
		ds, err := getItemCollection(f, cursor.UserID).Doc(cursor.ID).Get(ctx)
		if err != nil {
			return nil, errors.WithStack(err)
		}

		query = query.StartAfter(ds)
	}

	matchItem := query.Limit(first).Documents(ctx)
	docs, err := matchItem.GetAll()

	if err != nil {
		return nil, errors.WithStack(err)
	}

	for _, doc := range docs {
		var item *model.Item
		doc.DataTo(&item)

		items = append(items, item)
	}

	return items, nil
}
