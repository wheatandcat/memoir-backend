package graph

import (
	"context"
	"fmt"
	"strings"

	"github.com/wheatandcat/memoir-backend/graph/model"
	"github.com/wheatandcat/memoir-backend/repository"
)

// CreateRelationshipRequest 共有の招待リクエストを作成する
func (g *Graph) CreateRelationshipRequest(ctx context.Context, input *model.NewRelationshipRequest) (*model.RelationshipRequest, error) {
	if !g.Client.AuthToken.Valid(ctx) {
		return nil, fmt.Errorf("Invalid Authorization")
	}

	i, err := g.App.InviteRepository.Find(ctx, g.FirestoreClient, input.Code)
	if err != nil {
		return nil, err
	}
	if i.UserID == "" {
		return nil, fmt.Errorf("存在しない招待コードです")
	}

	rr := &model.RelationshipRequest{
		FollowerID: g.UserID,
		FollowedID: i.UserID,
		Status:     repository.RelationshipRequestStatusRequest,
		CreatedAt:  g.Client.Time.Now(),
		UpdatedAt:  g.Client.Time.Now(),
	}

	if err = g.App.RelationshipRequestRepository.Create(ctx, g.FirestoreClient, rr); err != nil {
		return nil, err
	}

	return rr, nil
}

// GetRelationshipRequests 共有の招待リクエストを取得する
func (g *Graph) GetRelationshipRequests(ctx context.Context, input *model.InputRelationshipRequests) (*model.RelationshipRequests, error) {
	t := g.Client.Time
	if !g.Client.AuthToken.Valid(ctx) {
		return nil, fmt.Errorf("Invalid Authorization")
	}

	cursor := repository.RelationshipRequestCursor{
		FollowerID: "",
		FollowedID: "",
	}

	cursorDate := strings.Split(*input.After, "/")
	if len(cursorDate) > 1 {
		cursor = repository.RelationshipRequestCursor{
			FollowerID: cursorDate[0],
			FollowedID: cursorDate[1],
		}
	}

	items, err := g.App.RelationshipRequestRepository.FindByFollowedID(ctx, g.FirestoreClient, g.UserID, input.First, cursor)
	if err != nil {
		return nil, err
	}

	var rres []*model.RelationshipRequestEdge
	for index, i := range items {
		items[index].CreatedAt = t.Location(i.CreatedAt)
		items[index].UpdatedAt = t.Location(i.UpdatedAt)

		rres = append(rres, &model.RelationshipRequestEdge{
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
	ibp := &model.RelationshipRequests{
		Edges:    rres,
		PageInfo: pi,
	}

	return ibp, nil
}
