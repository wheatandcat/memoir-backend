package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/wheatandcat/memoir-backend/graph/generated"
	"github.com/wheatandcat/memoir-backend/graph/model"
	ce "github.com/wheatandcat/memoir-backend/usecase/custom_error"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	g := NewGraphWithSetUserID(ctx, r.App, r.FirestoreClient, input.ID, "")
	result, err := g.CreateUser(ctx, &input)

	if err != nil {
		return nil, ce.CustomError(err)
	}

	return result, nil
}

// CreateAuthUser is the resolver for the createAuthUser field.
func (r *mutationResolver) CreateAuthUser(ctx context.Context, input model.NewAuthUser) (*model.AuthUser, error) {
	g := NewGraphWithSetUserID(ctx, r.App, r.FirestoreClient, "", "")
	result, err := g.CreateAuthUser(ctx, &input)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	return result, nil
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUser) (*model.User, error) {
	if err := input.Validate(); err != nil {
		return nil, ce.CustomError(err)
	}

	g, err := NewGraph(ctx, r.App, r.FirestoreClient)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	result, err := g.UpdateUser(ctx, &input)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	return result, nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context) (*model.User, error) {
	g, err := NewGraph(ctx, r.App, r.FirestoreClient)
	if err != nil {
		return nil, ce.CustomError(err)
	}
	result, err := g.DeleteUser(ctx)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	return result, nil
}

// CreateItem is the resolver for the createItem field.
func (r *mutationResolver) CreateItem(ctx context.Context, input model.NewItem) (*model.Item, error) {
	if err := input.Validate(); err != nil {
		return nil, ce.CustomError(err)
	}

	g, err := NewGraph(ctx, r.App, r.FirestoreClient)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	result, err := g.CreateItem(ctx, &input)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	return result, nil
}

// UpdateItem is the resolver for the updateItem field.
func (r *mutationResolver) UpdateItem(ctx context.Context, input model.UpdateItem) (*model.Item, error) {
	if err := input.Validate(); err != nil {
		return nil, ce.CustomError(err)
	}

	g, err := NewGraph(ctx, r.App, r.FirestoreClient)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	result, err := g.UpdateItem(ctx, &input)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	return result, nil
}

// DeleteItem is the resolver for the deleteItem field.
func (r *mutationResolver) DeleteItem(ctx context.Context, input model.DeleteItem) (*model.Item, error) {
	g, err := NewGraph(ctx, r.App, r.FirestoreClient)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	result, err := g.DeleteItem(ctx, &input)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	return result, nil
}

// CreateInvite is the resolver for the createInvite field.
func (r *mutationResolver) CreateInvite(ctx context.Context) (*model.Invite, error) {
	g, err := NewGraph(ctx, r.App, r.FirestoreClient)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	result, err := g.CreateInvite(ctx)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	return result, nil
}

// UpdateInvite is the resolver for the updateInvite field.
func (r *mutationResolver) UpdateInvite(ctx context.Context) (*model.Invite, error) {
	g, err := NewGraph(ctx, r.App, r.FirestoreClient)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	result, err := g.UpdateInvite(ctx)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	return result, nil
}

// CreateRelationshipRequest is the resolver for the createRelationshipRequest field.
func (r *mutationResolver) CreateRelationshipRequest(ctx context.Context, input model.NewRelationshipRequest) (*model.RelationshipRequest, error) {
	g, err := NewGraph(ctx, r.App, r.FirestoreClient)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	result, err := g.CreateRelationshipRequest(ctx, input)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	return result, nil
}

// AcceptRelationshipRequest is the resolver for the acceptRelationshipRequest field.
func (r *mutationResolver) AcceptRelationshipRequest(ctx context.Context, followedID string) (*model.RelationshipRequest, error) {
	g, err := NewGraph(ctx, r.App, r.FirestoreClient)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	result, err := g.AcceptRelationshipRequest(ctx, followedID)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	return result, nil
}

// NgRelationshipRequest is the resolver for the ngRelationshipRequest field.
func (r *mutationResolver) NgRelationshipRequest(ctx context.Context, followedID string) (*model.RelationshipRequest, error) {
	g, err := NewGraph(ctx, r.App, r.FirestoreClient)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	result, err := g.NgRelationshipRequest(ctx, followedID)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	return result, nil
}

// DeleteRelationship is the resolver for the deleteRelationship field.
func (r *mutationResolver) DeleteRelationship(ctx context.Context, followedID string) (*model.Relationship, error) {
	g, err := NewGraph(ctx, r.App, r.FirestoreClient)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	result, err := g.DeleteRelationship(ctx, followedID)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	return result, nil
}

// CreatePushToken is the resolver for the createPushToken field.
func (r *mutationResolver) CreatePushToken(ctx context.Context, input model.NewPushToken) (*model.PushToken, error) {
	g, err := NewGraph(ctx, r.App, r.FirestoreClient)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	result, err := g.CreatePushToken(ctx, &input)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	return result, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	g, err := NewGraph(ctx, r.App, r.FirestoreClient)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	result, err := g.GetUser(ctx)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	return result, nil
}

// ExistAuthUser is the resolver for the existAuthUser field.
func (r *queryResolver) ExistAuthUser(ctx context.Context) (*model.ExistAuthUser, error) {
	g := NewGraphWithSetUserID(ctx, r.App, r.FirestoreClient, "", "")

	result, err := g.ExistAuthUser(ctx)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	return result, nil
}

// Item is the resolver for the item field.
func (r *queryResolver) Item(ctx context.Context, id string) (*model.Item, error) {
	g, err := NewGraph(ctx, r.App, r.FirestoreClient)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	result, err := g.GetItem(ctx, id)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	return result, nil
}

// ItemsByDate is the resolver for the itemsByDate field.
func (r *queryResolver) ItemsByDate(ctx context.Context, date time.Time) ([]*model.Item, error) {
	g, err := NewGraph(ctx, r.App, r.FirestoreClient)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	result, err := g.GetItemsInDate(ctx, date)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	return result, nil
}

// ItemsInDate is the resolver for the itemsInDate field.
func (r *queryResolver) ItemsInDate(ctx context.Context, date time.Time) ([]*model.Item, error) {
	g, err := NewGraph(ctx, r.App, r.FirestoreClient)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	result, err := g.GetItemsInDate(ctx, date)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	return result, nil
}

// ItemsInPeriod is the resolver for the itemsInPeriod field.
func (r *queryResolver) ItemsInPeriod(ctx context.Context, input model.InputItemsInPeriod) (*model.ItemsInPeriod, error) {
	g, err := NewGraph(ctx, r.App, r.FirestoreClient)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	result, err := g.GetItemsInPeriod(ctx, input)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	return result, nil
}

// Invite is the resolver for the invite field.
func (r *queryResolver) Invite(ctx context.Context) (*model.Invite, error) {
	g, err := NewGraph(ctx, r.App, r.FirestoreClient)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	result, err := g.GetInviteByUseID(ctx)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	return result, nil
}

// InviteByCode is the resolver for the inviteByCode field.
func (r *queryResolver) InviteByCode(ctx context.Context, code string) (*model.User, error) {
	g, err := NewGraph(ctx, r.App, r.FirestoreClient)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	result, err := g.GetInviteByCode(ctx, code)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	return result, nil
}

// RelationshipRequests is the resolver for the relationshipRequests field.
func (r *queryResolver) RelationshipRequests(ctx context.Context, input model.InputRelationshipRequests) (*model.RelationshipRequests, error) {
	g, err := NewGraph(ctx, r.App, r.FirestoreClient)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	userSkip := true

	octx := graphql.GetOperationContext(ctx)
	switch skip := octx.Variables["skip"].(type) {
	case bool:
		userSkip = skip
	}

	result, err := g.GetRelationshipRequests(ctx, input, userSkip)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	return result, nil
}

// Relationships is the resolver for the relationships field.
func (r *queryResolver) Relationships(ctx context.Context, input model.InputRelationships) (*model.Relationships, error) {
	g, err := NewGraph(ctx, r.App, r.FirestoreClient)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	userSkip := true

	octx := graphql.GetOperationContext(ctx)
	switch skip := octx.Variables["skip"].(type) {
	case bool:
		userSkip = skip
	}

	result, err := g.GetRelationships(ctx, input, userSkip)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	return result, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
