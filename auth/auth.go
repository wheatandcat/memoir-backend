package auth

import (
	"context"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// User ユーザータイプ
type User struct {
	ID          string
	FirebaseUID string
}

// NotLoginMiddleware ログイン前の時のMiddleware
func NotLoginMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			uid := r.Header.Get("Userid")
			if uid == "" {
				next.ServeHTTP(w, r)
				return
			}

			u := &User{
				ID: uid,
			}

			ctx := context.WithValue(r.Context(), userCtxKey, u)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// FirebaseLoginMiddleware ログイン後の時のMiddleware
func FirebaseLoginMiddleware(app *firebase.App) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			client, err := app.Auth(r.Context())
			if err != nil {
				http.Error(w, "Firebase not initialize", http.StatusBadRequest)
				return
			}

			auth := r.Header.Get("Authorization")
			if auth == "" {
				next.ServeHTTP(w, r)
				return
			}

			idToken := strings.Replace(auth, "Bearer ", "", 1)
			token, err := client.VerifyIDToken(r.Context(), idToken)
			if err != nil {
				http.Error(w, "Invalid token", http.StatusForbidden)
				return
			}

			u := &User{
				FirebaseUID: token.UID,
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
