package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/wheatandcat/memoir-backend/graph/model"
	ce "github.com/wheatandcat/memoir-backend/usecase/custom_error"
	"google.golang.org/grpc/codes"
)

//go:generate moq -out=moq/invite.go -pkg=moqs . InviteRepositoryInterface

type InviteRepositoryInterface interface {
	Create(ctx context.Context, f *firestore.Client, batch *firestore.WriteBatch, i *model.Invite)
	Delete(ctx context.Context, f *firestore.Client, batch *firestore.WriteBatch, code string)
	DeleteByUserID(ctx context.Context, f *firestore.Client, batch *firestore.WriteBatch, userID string) error
	Find(ctx context.Context, f *firestore.Client, code string) (*model.Invite, error)
	FindByUserID(ctx context.Context, f *firestore.Client, userID string) (*model.Invite, error)
}

type InviteRepository struct {
}

func NewInviteRepository() InviteRepositoryInterface {
	return &InviteRepository{}
}

// Create 招待を作成する
func (re *InviteRepository) Create(ctx context.Context, f *firestore.Client, batch *firestore.WriteBatch, i *model.Invite) {
	ref := f.Collection("invites").Doc(i.Code)
	batch.Set(ref, i)
}

// Delete アイテムを削除する
func (re *InviteRepository) Delete(ctx context.Context, f *firestore.Client, batch *firestore.WriteBatch, code string) {
	ref := f.Collection("invites").Doc(code)
	batch.Delete(ref)
}

// FindByUserID ユーザーIDから取得する
func (re *InviteRepository) FindByUserID(ctx context.Context, f *firestore.Client, userID string) (*model.Invite, error) {
	matchItem := f.Collection("invites").Where("UserID", "==", userID).OrderBy("CreatedAt", firestore.Desc).Limit(1).Documents(ctx)
	docs, err := matchItem.GetAll()

	if err != nil {
		return nil, ce.CustomError(err)
	}

	if len(docs) == 0 {
		return &model.Invite{}, nil
	}

	var invite *model.Invite

	if err = docs[0].DataTo(&invite); err != nil {
		return nil, ce.CustomError(err)
	}

	return invite, nil
}

// Find 取得する
func (re *InviteRepository) Find(ctx context.Context, f *firestore.Client, code string) (*model.Invite, error) {
	var i *model.Invite
	ds, err := f.Collection("invites").Doc(code).Get(ctx)

	if err != nil {
		if GrpcErrorStatusCode(err) == codes.InvalidArgument || GrpcErrorStatusCode(err) == codes.NotFound {
			return &model.Invite{}, nil
		}

		return i, ce.CustomError(err)
	}

	if err = ds.DataTo(&i); err != nil {
		return nil, ce.CustomError(err)
	}

	return i, ce.CustomError(err)
}

// DeleteByUserID ユーザーIDから削除する
func (re *InviteRepository) DeleteByUserID(ctx context.Context, f *firestore.Client, batch *firestore.WriteBatch, userID string) error {
	matchItem := f.Collection("invites").Where("UserID", "==", userID).OrderBy("CreatedAt", firestore.Desc).Documents(ctx)
	docs, err := matchItem.GetAll()

	if err != nil {
		return ce.CustomError(err)
	}

	if len(docs) == 0 {
		return nil
	}

	for _, doc := range docs {
		batch.Delete(doc.Ref)
	}

	return nil
}
