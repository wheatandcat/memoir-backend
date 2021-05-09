package mock_authToken

import (
	"context"
)

// AuthToken has generating method.
type AuthToken struct {
}

// Get トークンを取得する.
func (*AuthToken) Get(ctx context.Context) string {
	return "test-token"
}

// Is トークンを取得する.
func (*AuthToken) Valid(ctx context.Context) bool {
	return true
}
