package auth

import (
	"context"
	"fmt"

	"cloud.google.com/go/firestore"
	"github.com/wheatandcat/memoir-backend/graph/model"
	"github.com/wheatandcat/memoir-backend/repository"
	ce "github.com/wheatandcat/memoir-backend/usecase/custom_error"
	"google.golang.org/grpc/codes"
)

// CreateUser ユーザー作成
func (uci *useCaseImpl) CreateAuthUser(ctx context.Context, f *firestore.Client, input *model.NewAuthUser, u *repository.User, mu *model.AuthUser) error {
	aref := f.Collection("auth").Doc(u.FirebaseUID)
	uref := f.Collection("users").Doc(u.ID)
	err := f.RunTransaction(ctx, func(ctx context.Context, tx *firestore.Transaction) error {
		doc, err := tx.Get(aref)
		fmt.Printf("%+v", doc)

		if repository.GrpcErrorStatusCode(err) == codes.InvalidArgument || repository.GrpcErrorStatusCode(err) == codes.NotFound {
			// 既にユーザー作成済みの場合は更新しないで完了
			err := tx.Set(aref, mu)
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
