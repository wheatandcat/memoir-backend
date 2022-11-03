package auth

import (
	"context"

	"cloud.google.com/go/firestore"

	"github.com/wheatandcat/memoir-backend/repository"
	ce "github.com/wheatandcat/memoir-backend/usecase/custom_error"
)

// ユーザーを削除
func (uci *useCaseImpl) DeleteAuthUser(ctx context.Context, f *firestore.Client, uid string) error {
	app, err := repository.FirebaseApp(ctx)
	if err != nil {
		return ce.CustomError(err)
	}
	client, err := app.Auth(ctx)
	if err != nil {
		return ce.CustomError(err)
	}
	if err := client.DeleteUser(ctx, uid); err != nil {
		return ce.CustomError(err)
	}

	return nil
}
