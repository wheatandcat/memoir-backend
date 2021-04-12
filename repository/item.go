package repository

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/wheatandcat/memoir-backend/graph/model"
)

// UserRepositoryInterface is repository interface
type ItemRepositoryInterface interface {
	Create(ctx context.Context, f *firestore.Client, userID string, i *model.Item) error
	Update(ctx context.Context, f *firestore.Client, userID string, i *model.UpdateItem, updatedAt time.Time) error
	Delete(ctx context.Context, f *firestore.Client, userID string, i *model.DeleteItem) error
	GetItem(ctx context.Context, f *firestore.Client, userID string, id string) (*model.Item, error)
	GetItemsByDate(ctx context.Context, f *firestore.Client, userID string, date time.Time) ([]*model.Item, error)
	GetItemsByPeriod(ctx context.Context, f *firestore.Client, userID string, stertDate time.Time, endDate time.Time, first int, cursor ItemsByPeriodCursor) ([]*model.Item, error)
	GetItemUserMultipleInPeriod(ctx context.Context, f *firestore.Client, userID string, stertDate time.Time, endDate time.Time, first int, cursor ItemsByPeriodCursor) ([]*model.Item, error)
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

	return err
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

	return err
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
		return i, err
	}

	ds.DataTo(&i)

	return i, nil
}

// GetItemsByDate 日付でアイテムを取得する
func (re *ItemRepository) GetItemsByDate(ctx context.Context, f *firestore.Client, userID string, date time.Time) ([]*model.Item, error) {
	var items []*model.Item

	matchItem := getItemCollection(f, userID).Where("Date", "==", date).OrderBy("CreatedAt", firestore.Desc).Documents(ctx)
	docs, err := matchItem.GetAll()
	if err != nil {
		return nil, err
	}

	for _, doc := range docs {
		var item *model.Item
		doc.DataTo(&item)

		items = append(items, item)
	}

	return items, nil
}

type ItemsByPeriodCursor struct {
	Date      time.Time
	CreatedAt time.Time
	ID        string
}

// GetItemsByPeriod 期間でアイテムを取得する
func (re *ItemRepository) GetItemsByPeriod(ctx context.Context, f *firestore.Client, userID string, startDate time.Time, endDate time.Time, first int, cursor ItemsByPeriodCursor) ([]*model.Item, error) {
	var items []*model.Item
	query := getItemCollection(f, userID).Where("Date", ">=", startDate).Where("Date", "<=", endDate).OrderBy("Date", firestore.Asc).OrderBy("CreatedAt", firestore.Asc).OrderBy("ID", firestore.Asc)

	log.Println(cursor.ID)
	if cursor.ID != "" {
		query = query.StartAfter(cursor.Date, cursor.CreatedAt, cursor.ID)
	}

	matchItem := query.Limit(first).Documents(ctx)
	docs, err := matchItem.GetAll()

	if err != nil {
		return nil, err
	}

	for _, doc := range docs {
		var item *model.Item
		doc.DataTo(&item)

		items = append(items, item)
	}

	return items, nil
}

// GetItemUserMultipleInPeriod 期間でアイテムを取得する
func (re *ItemRepository) GetItemUserMultipleInPeriod(ctx context.Context, f *firestore.Client, userID string, startDate time.Time, endDate time.Time, first int, cursor ItemsByPeriodCursor) ([]*model.Item, error) {
	var items []*model.Item
	query := f.CollectionGroup("items").Where("UserID", "in", []string{userID}).Where("Date", ">=", startDate).Where("Date", "<=", endDate).OrderBy("Date", firestore.Asc).OrderBy("CreatedAt", firestore.Asc).OrderBy("ID", firestore.Asc)

	if cursor.ID != "" {
		query = query.StartAfter(cursor.Date, cursor.CreatedAt, cursor.ID)
	}

	matchItem := query.Limit(first).Documents(ctx)
	docs, err := matchItem.GetAll()

	if err != nil {
		return nil, err
	}

	for _, doc := range docs {
		var item *model.Item
		doc.DataTo(&item)

		items = append(items, item)
	}

	return items, nil
}
