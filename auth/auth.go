package auth

import (
	"context"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	ce "github.com/wheatandcat/memoir-backend/usecase/custom_error"
	"go.uber.org/zap"
)

var UserCtxKey = &contextKey{"user"}

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

			zap.L().Info("useInfo", zap.String("ID", u.ID))

			ctx := context.WithValue(r.Context(), UserCtxKey, u)
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
				e := ce.CustomErrorWrap(err, "Firebase not initialize:")
				http.Error(w, e.Error(), http.StatusBadRequest)
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
				e := ce.CustomError(err)
				http.Error(w, e.Error(), http.StatusForbidden)
				return
			}

			u := &User{
				FirebaseUID: token.UID,
			}

			zap.L().Info("useInfo", zap.String("FirebaseUID", u.FirebaseUID))

			ctx := context.WithValue(r.Context(), UserCtxKey, u)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *User {
	raw, _ := ctx.Value(UserCtxKey).(*User)
	return raw
}
