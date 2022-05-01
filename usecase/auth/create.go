package auth

import (
	"context"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/wheatandcat/memoir-backend/graph/model"
	"github.com/wheatandcat/memoir-backend/repository"
	ce "github.com/wheatandcat/memoir-backend/usecase/custom_error"
	"google.golang.org/grpc/codes"
)

type Auth struct {
	UserID    string
	CreatedAt time.Time
}

// CreateAuthUser 認証ユーザー作成
func (uci *useCaseImpl) CreateAuthUser(ctx context.Context, f *firestore.Client, input *model.NewAuthUser, u *repository.User, mu *model.AuthUser) error {
	aref := f.Collection("auth").Doc(u.FirebaseUID)
	uref := f.Collection("users").Doc(u.ID)
	err := f.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		_, err := tx.Get(aref)

		if repository.GrpcErrorStatusCode(err) == codes.InvalidArgument || repository.GrpcErrorStatusCode(err) == codes.NotFound {
			// 既にユーザー作成済みの場合は更新しないで完了
			a := Auth{
				UserID:    mu.ID,
				CreatedAt: mu.CreatedAt,
			}
			err := tx.Set(aref, a)
			if err != nil {
				return ce.CustomError(err)
			}

			if input.IsNewUser {
				u.CreatedAt = mu.CreatedAt

				if err := tx.Set(uref, u); err != nil {
					return ce.CustomError(err)
				}
			} else {
				u.UpdatedAt = mu.UpdatedAt
				var uu []firestore.Update
				uu = append(uu, firestore.Update{Path: "FirebaseUID", Value: u.FirebaseUID})

				if err = tx.Update(uref, uu); err != nil {
					return ce.CustomError(err)
				}
			}

			return nil
		}

		return nil
	})

	return err
}
