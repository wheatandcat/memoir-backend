package graph

import "cloud.google.com/go/firestore"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver Resolver type
type Resolver struct {
	FirestoreClient *firestore.Client
	App             *Application
}
