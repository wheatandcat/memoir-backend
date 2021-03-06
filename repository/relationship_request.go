package repository

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/wheatandcat/memoir-backend/graph/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	RelationshipRequestStatusRequest = 1
	RelationshipRequestStatusOK      = 2
	RelationshipRequestStatusNG      = 3
)

type RelationshipRequestInterface interface {
	Create(ctx context.Context, f *firestore.Client, i *model.RelationshipRequest) error
	Update(ctx context.Context, f *firestore.Client, batch *firestore.WriteBatch, i *model.RelationshipRequest)
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
	return err
}

// Update 更新する
func (re *RelationshipRequestRepository) Update(ctx context.Context, f *firestore.Client, batch *firestore.WriteBatch, i *model.RelationshipRequest) {
	var u []firestore.Update
	if i.Status != 0 {
		u = append(u, firestore.Update{Path: "Status", Value: i.Status})
	}
	u = append(u, firestore.Update{Path: "UpdatedAt", Value: i.UpdatedAt})

	ref := f.Collection("relationshipRequests").Doc(i.FollowerID + "_" + i.FollowedID)
	batch.Update(ref, u)
}

// Find 取得する
func (re *RelationshipRequestRepository) Find(ctx context.Context, f *firestore.Client, i *model.RelationshipRequest) (*model.RelationshipRequest, error) {
	var rr *model.RelationshipRequest
	ds, err := f.Collection("relationshipRequests").Doc(i.FollowerID + "_" + i.FollowedID).Get(ctx)
	if err != nil {
		if status.Code(err) == codes.InvalidArgument || status.Code(err) == codes.NotFound {
			return &model.RelationshipRequest{}, nil
		}

		return i, err
	}

	ds.DataTo(&rr)

	return rr, err
}

// FindByFollowedID ページングで取得する
func (re *RelationshipRequestRepository) FindByFollowedID(ctx context.Context, f *firestore.Client, userID string, first int, cursor RelationshipRequestCursor) ([]*model.RelationshipRequest, error) {
	var items []*model.RelationshipRequest
	query := f.Collection("relationshipRequests").Where("FollowedID", "==", userID).Where("Status", "==", RelationshipRequestStatusRequest).OrderBy("CreatedAt", firestore.Desc)

	if cursor.FollowerID != "" {
		ds, err := f.Collection("relationshipRequests").Doc(cursor.FollowerID + "_" + cursor.FollowedID).Get(ctx)
		if err != nil {
			return nil, err
		}

		query = query.StartAfter(ds)
	}

	matchItem := query.Limit(first).Documents(ctx)
	docs, err := matchItem.GetAll()

	if err != nil {
		return nil, err
	}

	for _, doc := range docs {
		var item *model.RelationshipRequest
		doc.DataTo(&item)

		items = append(items, item)
	}

	return items, nil
}
