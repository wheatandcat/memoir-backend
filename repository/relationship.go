package repository

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/wheatandcat/memoir-backend/graph/model"
	ce "github.com/wheatandcat/memoir-backend/usecase/custom_error"
)

//go:generate moq -out=moq/relationship.go -pkg=moqs . RelationshipInterface

type RelationshipInterface interface {
	Create(ctx context.Context, f *firestore.Client, batch *firestore.WriteBatch, i *model.Relationship)
	Delete(ctx context.Context, f *firestore.Client, batch *firestore.WriteBatch, i *model.Relationship)
	FindByFollowedID(ctx context.Context, f *firestore.Client, userID string, first int, cursor RelationshipCursor) ([]*model.Relationship, error)
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
func (re *RelationshipRepository) Create(ctx context.Context, f *firestore.Client, batch *firestore.WriteBatch, i *model.Relationship) {
	rrd := RelationshipData{
		ID:         i.ID,
		FollowerID: i.FollowerID,
		FollowedID: i.FollowedID,
		CreatedAt:  i.CreatedAt,
		UpdatedAt:  i.UpdatedAt,
	}

	ref := f.Collection("relationships").Doc(i.FollowerID + "_" + i.FollowedID)
	batch.Set(ref, rrd)
}

// Delete 削除する
func (re *RelationshipRepository) Delete(ctx context.Context, f *firestore.Client, batch *firestore.WriteBatch, i *model.Relationship) {
	ref := f.Collection("relationships").Doc(i.FollowerID + "_" + i.FollowedID)
	batch.Delete(ref)
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
