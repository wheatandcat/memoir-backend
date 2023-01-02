package repository

import (
	"context"
	"fmt"
	"time"

	"cloud.google.com/go/firestore"

	"github.com/wheatandcat/memoir-backend/graph/model"
	ce "github.com/wheatandcat/memoir-backend/usecase/custom_error"
)

//go:generate moq -out=moq/user.go -pkg=moqs . UserRepositoryInterface

type UserRepositoryInterface interface {
	Create(ctx context.Context, f *firestore.Client, u *model.User) error
	Update(ctx context.Context, f *firestore.Client, u *model.User) error
	UpdateFirebaseUID(ctx context.Context, f *firestore.Client, user *User) error
	Delete(ctx context.Context, f *firestore.Client, batch *firestore.BulkWriter, uid string) error
	FindByUID(ctx context.Context, f *firestore.Client, uid string) (*model.User, error)
	FindDatabaseDataByUID(ctx context.Context, f *firestore.Client, uid string) (*User, error)
	FindByFirebaseUID(ctx context.Context, f *firestore.Client, fUID string) (*model.User, error)
	ExistByFirebaseUID(ctx context.Context, f *firestore.Client, fUID string) (bool, error)
	FindInUID(ctx context.Context, f *firestore.Client, uid []string) ([]*model.User, error)
}

type User struct {
	ID          string
	FirebaseUID string
	DisplayName string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserRepository struct {
}

func NewUserRepository() UserRepositoryInterface {
	return &UserRepository{}
}

// Create ユーザーを作成する
func (re *UserRepository) Create(ctx context.Context, f *firestore.Client, u *model.User) error {
	_, err := f.Collection("users").Doc(u.ID).Set(ctx, u)

	return ce.CustomError(err)
}

// Update ユーザーを更新する
func (re *UserRepository) Update(ctx context.Context, f *firestore.Client, u *model.User) error {
	var fu []firestore.Update
	if u.DisplayName != "" {
		fu = append(fu, firestore.Update{Path: "DisplayName", Value: u.DisplayName})
	}
	if u.Image != "" {
		fu = append(fu, firestore.Update{Path: "Image", Value: u.Image})
	}

	fu = append(fu, firestore.Update{Path: "UpdatedAt", Value: u.UpdatedAt})

	_, err := f.Collection("users").Doc(u.ID).Update(ctx, fu)

	return ce.CustomError(err)
}

// Delete ユーザーを削除する
func (re *UserRepository) Delete(ctx context.Context, f *firestore.Client, batch *firestore.BulkWriter, uid string) error {
	ref := f.Collection("users").Doc(uid)
	j, err := batch.Delete(ref)
	if err != nil {
		return ce.CustomError(err)
	}
	if j == nil {
		return ce.CustomError(fmt.Errorf("BulkWriter: got nil"))
	}
	if _, err := j.Results(); err != nil {
		return ce.CustomError(err)
	}

	matchPushToken := ref.Collection("pushToken").Documents(ctx)
	docPushTokens, err := matchPushToken.GetAll()
	if err != nil {
		return ce.CustomError(err)
	}
	for _, doc := range docPushTokens {
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

	matchItems := ref.Collection("items").Documents(ctx)
	docItems, err := matchItems.GetAll()
	if err != nil {
		return ce.CustomError(err)
	}
	for _, doc := range docItems {
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

// UpdateFirebaseUID ユーザーFirebaseUIを更新する
func (re *UserRepository) UpdateFirebaseUID(ctx context.Context, f *firestore.Client, user *User) error {
	var u []firestore.Update
	u = append(u, firestore.Update{Path: "FirebaseUID", Value: user.FirebaseUID})
	u = append(u, firestore.Update{Path: "UpdatedAt", Value: user.UpdatedAt})

	_, err := f.Collection("users").Doc(user.ID).Update(ctx, u)

	return ce.CustomError(err)
}

// FindByUID ユーザーIDから取得する
func (re *UserRepository) FindByUID(ctx context.Context, f *firestore.Client, uid string) (*model.User, error) {
	var u *model.User
	ds, err := f.Collection("users").Doc(uid).Get(ctx)
	if err != nil {
		return u, ce.CustomError(err)
	}

	if err = ds.DataTo(&u); err != nil {
		return u, ce.CustomError(err)
	}

	return u, nil
}

// FindDatabaseDataByUID ユーザーIDからデータベースのデータを取得する
func (re *UserRepository) FindDatabaseDataByUID(ctx context.Context, f *firestore.Client, uid string) (*User, error) {
	var u *User
	ds, err := f.Collection("users").Doc(uid).Get(ctx)
	if err != nil {
		return u, ce.CustomError(err)
	}

	if err = ds.DataTo(&u); err != nil {
		return u, ce.CustomError(err)
	}

	return u, nil
}

// FindByFirebaseUID FirebaseユーザーIDから取得する
func (re *UserRepository) FindByFirebaseUID(ctx context.Context, f *firestore.Client, fUID string) (*model.User, error) {
	var u *model.User
	matchItem := f.Collection("users").Where("FirebaseUID", "==", fUID).OrderBy("CreatedAt", firestore.Asc).Limit(1).Documents(ctx)
	docs, err := matchItem.GetAll()
	if err != nil {
		return nil, ce.CustomError(err)
	}

	if len(docs) == 0 {
		return nil, ce.CustomError(ce.NewNotFoundError("ユーザーが存在しません"))
	}

	if err = docs[0].DataTo(&u); err != nil {
		return u, ce.CustomError(err)
	}

	return u, nil
}

// ExistByFirebaseUID FirebaseユーザーIDが存在するか取得する
func (re *UserRepository) ExistByFirebaseUID(ctx context.Context, f *firestore.Client, fUID string) (bool, error) {
	matchItem := f.Collection("users").Where("FirebaseUID", "==", fUID).OrderBy("CreatedAt", firestore.Asc).Limit(1).Documents(ctx)
	docs, err := matchItem.GetAll()
	if err != nil {
		return false, ce.CustomError(err)
	}

	return len(docs) > 0, nil

}

// FindInUID ユーザーIDリストから取得する
func (re *UserRepository) FindInUID(ctx context.Context, f *firestore.Client, uid []string) ([]*model.User, error) {
	matchItem := f.Collection("users").Where("ID", "in", uid).OrderBy("CreatedAt", firestore.Asc).Limit(10).Documents(ctx)
	docs, err := matchItem.GetAll()
	if err != nil {
		return nil, ce.CustomError(err)
	}

	us := make([]*model.User, len(docs))
	for i, doc := range docs {
		var u *model.User
		if err = doc.DataTo(&u); err != nil {
			return us, ce.CustomError(err)
		}
		us[i] = u
	}

	return us, nil
}
