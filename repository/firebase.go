package repository

import (
	"context"
	"os"

	firebase "firebase.google.com/go"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

// FirebaseApp Firebase App
func FirebaseApp(ctx context.Context) (*firebase.App, error) {
	credentials, err := google.CredentialsFromJSON(ctx, []byte(os.Getenv("FIREBASE_KEYFILE_JSON")))
	if err != nil {
		return nil, err
	}
	opt := option.WithCredentials(credentials)

	return firebase.NewApp(ctx, nil, opt)
}
