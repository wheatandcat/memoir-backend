package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/wheatandcat/memoir-backend/graph/generated"
	"github.com/wheatandcat/memoir-backend/graph/model"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	g := NewGraphWithSetUserID(r.App, r.FirestoreClient, input.ID)
	result, err := g.CreateUser(ctx, &input)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *mutationResolver) CreateItem(ctx context.Context, input model.NewItem) (*model.Item, error) {
	g, err := NewGraph(ctx, r.App, r.FirestoreClient)
	if err != nil {
		return nil, err
	}

	result, err := g.CreateItem(ctx, &input)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *mutationResolver) UpdateItem(ctx context.Context, input model.UpdateItem) (*model.Item, error) {
	g, err := NewGraph(ctx, r.App, r.FirestoreClient)
	if err != nil {
		return nil, err
	}

	result, err := g.UpdateItem(ctx, &input)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *mutationResolver) DeleteItem(ctx context.Context, input model.DeleteItem) (*model.Item, error) {
	g, err := NewGraph(ctx, r.App, r.FirestoreClient)
	if err != nil {
		return nil, err
	}

	result, err := g.DeleteItem(ctx, &input)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *queryResolver) User(ctx context.Context) (*model.User, error) {
	g, err := NewGraph(ctx, r.App, r.FirestoreClient)
	if err != nil {
		return nil, err
	}

	result, err := g.GetUser(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *queryResolver) Item(ctx context.Context, id string) (*model.Item, error) {
	g, err := NewGraph(ctx, r.App, r.FirestoreClient)
	if err != nil {
		return nil, err
	}

	result, err := g.GetItem(ctx, id)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *queryResolver) ItemsByDate(ctx context.Context, date time.Time) ([]*model.Item, error) {
	g, err := NewGraph(ctx, r.App, r.FirestoreClient)
	if err != nil {
		return nil, err
	}

	result, err := g.GetItemsByDate(ctx, date)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
