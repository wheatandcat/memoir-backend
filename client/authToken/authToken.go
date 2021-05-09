package authToken

import (
	"context"

	"github.com/wheatandcat/memoir-backend/auth"
)

// AuthTokenClient 認証トークン
type AuthTokenClient interface {
	Get(ctx context.Context) string
	Valid(ctx context.Context) bool
}

// AuthToken has generating method.
type AuthToken struct {
}

// Get トークンを取得
func (*AuthToken) Get(ctx context.Context) string {
	raw, _ := ctx.Value(auth.UserCtxKey).(*auth.User)

	return raw.FirebaseUID
}

// Valid トークンが存在するか判定
func (*AuthToken) Valid(ctx context.Context) bool {
	raw, _ := ctx.Value(auth.UserCtxKey).(*auth.User)

	return raw.FirebaseUID != ""
}
