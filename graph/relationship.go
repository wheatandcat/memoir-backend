package graph

import (
	"context"
	"fmt"
	"strings"

	"github.com/wheatandcat/memoir-backend/graph/model"
	"github.com/wheatandcat/memoir-backend/repository"
)

// DeleteRelationship 共有ユーザーを解除する
func (g *Graph) DeleteRelationship(ctx context.Context, followedID string) (*model.Relationship, error) {
	if !g.Client.AuthToken.Valid(ctx) {
		return nil, fmt.Errorf("Invalid Authorization")
	}

	batch := g.FirestoreClient.Batch()

	r1 := &model.Relationship{
		FollowerID: g.UserID,
		FollowedID: followedID,
	}
	r2 := &model.Relationship{
		FollowerID: followedID,
		FollowedID: g.UserID,
	}

	g.App.RelationshipRepository.Delete(ctx, g.FirestoreClient, batch, r1)
	g.App.RelationshipRepository.Delete(ctx, g.FirestoreClient, batch, r2)

	if err := g.App.CommonRepository.Commit(ctx, batch); err != nil {
		return nil, err
	}

	return r1, nil
}

// GetRelationships 共有の招待リクエストを取得する
func (g *Graph) GetRelationships(ctx context.Context, input model.InputRelationships, userSkip bool) (*model.Relationships, error) {
	t := g.Client.Time
	if !g.Client.AuthToken.Valid(ctx) {
		return nil, fmt.Errorf("Invalid Authorization")
	}

	cursor := repository.RelationshipCursor{
		FollowerID: "",
		FollowedID: "",
	}

	cursorDate := strings.Split(*input.After, "/")
	if len(cursorDate) > 1 {
		cursor = repository.RelationshipCursor{
			FollowerID: cursorDate[0],
			FollowedID: cursorDate[1],
		}
	}

	items, err := g.App.RelationshipRepository.FindByFollowedID(ctx, g.FirestoreClient, g.UserID, input.First, cursor)
	if err != nil {
		return nil, err
	}

	userID := []string{}
	for _, i := range items {
		userID = append(userID, i.FollowerID)
	}
	users := []*model.User{}

	if !userSkip {
		users, err = g.App.UserRepository.FindInUID(ctx, g.FirestoreClient, userID)
		if err != nil {
			return nil, err
		}

	}

	var rres []*model.RelationshipEdge

	for index, i := range items {
		items[index].CreatedAt = t.Location(i.CreatedAt)
		items[index].UpdatedAt = t.Location(i.UpdatedAt)

		user := &model.User{}

		for _, u := range users {

			if u.ID == i.FollowerID {
				user = u
			}

		}

		items[index].User = user

		rres = append(rres, &model.RelationshipEdge{
			Node:   items[index],
			Cursor: i.FollowedID + "/" + i.FollowerID,
		})

	}

	pi := &model.PageInfo{
		HasNextPage: false,
		EndCursor:   "",
	}

	if len(rres) > 0 {
		pi.HasNextPage = input.First == len(items)
		pi.EndCursor = rres[len(items)-1].Cursor
	}
	ibp := &model.Relationships{
		Edges:    rres,
		PageInfo: pi,
	}

	return ibp, nil
}
