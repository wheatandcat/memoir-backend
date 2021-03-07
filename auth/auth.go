package auth

import (
	"context"
	"net/http"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// User ユーザータイプ
type User struct {
	ID string
}

// NotLoginMiddleware ログイン前の時のMiddleware
func NotLoginMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			uid := r.Header.Get("Userid")

			u := &User{
				ID: uid,
			}

			ctx := context.WithValue(r.Context(), userCtxKey, u)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *User {
	raw, _ := ctx.Value(userCtxKey).(*User)
	return raw
}
