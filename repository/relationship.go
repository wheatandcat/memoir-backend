package repository

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"

	ce "github.com/wheatandcat/memoir-backend/usecase/custom_error"

	"github.com/wheatandcat/memoir-backend/graph/model"
)

//go:generate moq -out=moq/relationship.go -pkg=moqs . RelationshipInterface

type RelationshipInterface interface {
	Create(ctx context.Context, f *firestore.Client, tx *firestore.Transaction, i *model.Relationship) error
	Delete(ctx context.Context, f *firestore.Client, tx *firestore.Transaction, i *model.Relationship) error
	FindByFollowedID(ctx context.Context, f *firestore.Client, userID string, first int, cursor RelationshipCursor) ([]*model.Relationship, error)
	ExistByFollowedID(ctx context.Context, f *firestore.Client, userID string) (bool, error)
}

type RelationshipRepository struct {
}

func NewRelationshipRepository() RelationshipInterface {
	return &RelationshipRepository{}
}

type RelationshipCursor struct {
	FollowerID string
	FollowedID string
}

type RelationshipData struct {
	ID         string
	FollowerID string
	FollowedID string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// Create 作成する
func (re *RelationshipRepository) Create(ctx context.Context, f *firestore.Client, tx *firestore.Transaction, i *model.Relationship) error {
	rrd := RelationshipData{
		ID:         i.ID,
		FollowerID: i.FollowerID,
		FollowedID: i.FollowedID,
		CreatedAt:  i.CreatedAt,
		UpdatedAt:  i.UpdatedAt,
	}

	ref := f.Collection("relationships").Doc(i.FollowerID + "_" + i.FollowedID)
	err := tx.Set(ref, rrd)
	if err != nil {
		return ce.CustomError(err)
	}
	return nil
}

// Delete 削除する
func (re *RelationshipRepository) Delete(ctx context.Context, f *firestore.Client, tx *firestore.Transaction, i *model.Relationship) error {
	ref := f.Collection("relationships").Doc(i.FollowerID + "_" + i.FollowedID)
	err := tx.Delete(ref)
	if err != nil {
		return ce.CustomError(err)
	}
	return nil
}

// Find 取得する
func (re *RelationshipRepository) Find(ctx context.Context, f *firestore.Client, i *model.Relationship) (*model.Relationship, error) {
	var rr *model.Relationship
	ds, err := f.Collection("relationships").Doc(i.FollowerID + "_" + i.FollowedID).Get(ctx)
	if err != nil {
		return i, ce.CustomError(err)
	}

	if err := ds.DataTo(&rr); err != nil {
		return i, ce.CustomError(err)
	}

	return rr, nil
}

// FindByFollowedID ページングで取得する
func (re *RelationshipRepository) FindByFollowedID(ctx context.Context, f *firestore.Client, userID string, first int, cursor RelationshipCursor) ([]*model.Relationship, error) {
	query := f.Collection("relationships").Where("FollowedID", "==", userID).OrderBy("CreatedAt", firestore.Desc)

	if cursor.FollowerID != "" {
		ds, err := f.Collection("relationships").Doc(cursor.FollowerID + "_" + cursor.FollowedID).Get(ctx)
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

	items := make([]*model.Relationship, len(docs))
	for i, doc := range docs {
		var item *model.Relationship
		if err = doc.DataTo(&item); err != nil {
			return items, ce.CustomError(err)
		}

		items[i] = item
	}

	return items, nil
}

// ExistByFollowedID フォローしているか確認
func (re *RelationshipRepository) ExistByFollowedID(ctx context.Context, f *firestore.Client, userID string) (bool, error) {
	query := f.Collection("relationships").Where("FollowedID", "==", userID).OrderBy("CreatedAt", firestore.Desc)
	matchItem := query.Documents(ctx)
	docs, err := matchItem.GetAll()

	if err != nil {
		return false, ce.CustomError(err)
	}

	if len(docs) > 0 {
		return true, nil
	}

	return false, nil
}
