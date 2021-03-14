package repository

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/wheatandcat/memoir-backend/graph/model"
)

// UserRepositoryInterface is repository interface
type ItemRepositoryInterface interface {
	Create(ctx context.Context, f *firestore.Client, userID string, i *model.Item) error
	GetItem(ctx context.Context, f *firestore.Client, userID string, id string) (*model.Item, error)
	GetItemsByDate(ctx context.Context, f *firestore.Client, userID string, date time.Time) ([]*model.Item, error)
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

	matchItem := f.CollectionGroup("items").Where("Date", "==", date).OrderBy("ID", firestore.Asc).Documents(ctx)
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
