package repository

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"
	"google.golang.org/grpc/codes"

	"github.com/wheatandcat/memoir-backend/graph/model"
	ce "github.com/wheatandcat/memoir-backend/usecase/custom_error"
)

const (
	RelationshipRequestStatusRequest = 1
	RelationshipRequestStatusOK      = 2
	RelationshipRequestStatusNG      = 3
)

//go:generate moq -out=moq/relationship_request.go -pkg=moqs . RelationshipRequestInterface

type RelationshipRequestInterface interface {
	Create(ctx context.Context, f *firestore.Client, i *model.RelationshipRequest) error
	Update(ctx context.Context, f *firestore.Client, batch *firestore.BulkWriter, i *model.RelationshipRequest) error
	DeleteByFollowedID(ctx context.Context, f *firestore.Client, batch *firestore.BulkWriter, userID string) error
	DeleteByFollowerID(ctx context.Context, f *firestore.Client, batch *firestore.BulkWriter, userID string) error
	Find(ctx context.Context, f *firestore.Client, i *model.RelationshipRequest) (*model.RelationshipRequest, error)
	FindByFollowedID(ctx context.Context, f *firestore.Client, userID string, first int, cursor RelationshipRequestCursor) ([]*model.RelationshipRequest, error)
}

type RelationshipRequestRepository struct {
}

func NewRelationshipRequestRepository() RelationshipRequestInterface {
	return &RelationshipRequestRepository{}
}

type RelationshipRequestCursor struct {
	FollowerID string
	FollowedID string
}

type RelationshipRequestData struct {
	ID         string
	FollowerID string
	FollowedID string
	Status     int
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

// Create 作成する
func (re *RelationshipRequestRepository) Create(ctx context.Context, f *firestore.Client, i *model.RelationshipRequest) error {
	rrd := RelationshipRequestData{
		ID:         i.ID,
		FollowerID: i.FollowerID,
		FollowedID: i.FollowedID,
		Status:     i.Status,
		CreatedAt:  i.CreatedAt,
		UpdatedAt:  i.UpdatedAt,
	}

	_, err := f.Collection("relationshipRequests").Doc(i.FollowerID+"_"+i.FollowedID).Set(ctx, rrd)
	return ce.CustomError(err)
}

// Update 更新する
func (re *RelationshipRequestRepository) Update(ctx context.Context, f *firestore.Client, batch *firestore.BulkWriter, i *model.RelationshipRequest) error {
	var u []firestore.Update
	if i.Status != 0 {
		u = append(u, firestore.Update{Path: "Status", Value: i.Status})
	}
	u = append(u, firestore.Update{Path: "UpdatedAt", Value: i.UpdatedAt})

	ref := f.Collection("relationshipRequests").Doc(i.FollowerID + "_" + i.FollowedID)
	j, err := batch.Update(ref, u)
	if err != nil {
		return ce.CustomError(err)
	}
	if j == nil {
		return ce.CustomError(fmt.Errorf("BulkWriter: got nil"))
	}
	if _, err := j.Results(); err != nil {
		return ce.CustomError(err)
	}
	return nil
}

// DeleteByFollowedID ユーザーIDから削除する
func (re *RelationshipRequestRepository) DeleteByFollowedID(ctx context.Context, f *firestore.Client, batch *firestore.BulkWriter, userID string) error {
	matchItem := f.Collection("relationshipRequests").Where("FollowedID", "==", userID).OrderBy("CreatedAt", firestore.Desc).Documents(ctx)
	docs, err := matchItem.GetAll()

	if err != nil {
		return ce.CustomError(err)
	}

	if len(docs) == 0 {
		return nil
	}

	for _, doc := range docs {
		j, err := batch.Delete(doc.Ref)
		if err != nil {
			return ce.CustomError(err)
		}
		if j == nil {
			return ce.CustomError(fmt.Errorf("BulkWriter: got nil"))
		}
		if _, err := j.Results(); err != nil {
			return ce.CustomError(err)
		}
	}

	return nil
}

// DeleteByFollowerID ユーザーIDから削除する
func (re *RelationshipRequestRepository) DeleteByFollowerID(ctx context.Context, f *firestore.Client, batch *firestore.BulkWriter, userID string) error {
	matchItem := f.Collection("relationshipRequests").Where("FollowerID", "==", userID).OrderBy("CreatedAt", firestore.Desc).Documents(ctx)
	docs, err := matchItem.GetAll()

	if err != nil {
		return ce.CustomError(err)
	}

	if len(docs) == 0 {
		return nil
	}

	for _, doc := range docs {
		j, err := batch.Delete(doc.Ref)
		if err != nil {
			return ce.CustomError(err)
		}
		if j == nil {
			return ce.CustomError(fmt.Errorf("BulkWriter: got nil"))
		}
		if _, err := j.Results(); err != nil {
			return ce.CustomError(err)
		}
	}

	return nil
}

// Find 取得する
func (re *RelationshipRequestRepository) Find(ctx context.Context, f *firestore.Client, i *model.RelationshipRequest) (*model.RelationshipRequest, error) {
	var rr *model.RelationshipRequest
	ds, err := f.Collection("relationshipRequests").Doc(i.FollowerID + "_" + i.FollowedID).Get(ctx)
	if err != nil {
		if GrpcErrorStatusCode(err) == codes.InvalidArgument || GrpcErrorStatusCode(err) == codes.NotFound {
			return &model.RelationshipRequest{}, nil
		}

		return i, ce.CustomError(err)
	}

	err = ds.DataTo(&rr)

	return rr, ce.CustomError(err)
}

// FindByFollowedID ページングで取得する
func (re *RelationshipRequestRepository) FindByFollowedID(ctx context.Context, f *firestore.Client, userID string, first int, cursor RelationshipRequestCursor) ([]*model.RelationshipRequest, error) {
	query := f.Collection("relationshipRequests").Where("FollowedID", "==", userID).Where("Status", "==", RelationshipRequestStatusRequest).OrderBy("CreatedAt", firestore.Desc)

	if cursor.FollowerID != "" {
		ds, err := f.Collection("relationshipRequests").Doc(cursor.FollowerID + "_" + cursor.FollowedID).Get(ctx)
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

	items := make([]*model.RelationshipRequest, len(docs))

	for i, doc := range docs {
		var item *model.RelationshipRequest
		if err = doc.DataTo(&item); err != nil {
			return items, ce.CustomError(err)
		}
		items[i] = item
	}

	return items, nil
}
