package repository

import (
	"context"

	"cloud.google.com/go/firestore"
	"github.com/wheatandcat/memoir-backend/graph/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type InviteRepositoryInterface interface {
	Create(ctx context.Context, f *firestore.Client, batch *firestore.WriteBatch, i *model.Invite)
	Delete(ctx context.Context, f *firestore.Client, batch *firestore.WriteBatch, code string)
	Commit(ctx context.Context, batch *firestore.WriteBatch) error
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

// Commit コミットする
func (re *InviteRepository) Commit(ctx context.Context, batch *firestore.WriteBatch) error {
	_, err := batch.Commit(ctx)

	return err
}

// FindByUserID ユーザーIDから取得する
func (re *InviteRepository) FindByUserID(ctx context.Context, f *firestore.Client, userID string) (*model.Invite, error) {
	matchItem := f.Collection("invites").Where("UserID", "==", userID).OrderBy("CreatedAt", firestore.Desc).Limit(1).Documents(ctx)
	docs, err := matchItem.GetAll()

	if err != nil {
		return nil, err
	}

	if len(docs) == 0 {
		return &model.Invite{}, nil
	}

	var invite *model.Invite
	docs[0].DataTo(&invite)

	return invite, nil
}

// Find 取得する
func (re *InviteRepository) Find(ctx context.Context, f *firestore.Client, code string) (*model.Invite, error) {
	var i *model.Invite
	ds, err := f.Collection("invites").Doc(code).Get(ctx)

	if err != nil {
		if status.Code(err) == codes.InvalidArgument || status.Code(err) == codes.NotFound {
			return &model.Invite{}, nil
		}

		return i, err
	}

	ds.DataTo(&i)

	return i, err
}
