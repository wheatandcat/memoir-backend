package graph

import (
	"context"
	"strings"

	"google.golang.org/grpc/codes"

	"github.com/wheatandcat/memoir-backend/client/task"
	"github.com/wheatandcat/memoir-backend/graph/model"
	"github.com/wheatandcat/memoir-backend/repository"
	ce "github.com/wheatandcat/memoir-backend/usecase/custom_error"
)

// CreateRelationshipRequest 共有の招待リクエストを作成する
func (g *Graph) CreateRelationshipRequest(ctx context.Context, input model.NewRelationshipRequest) (*model.RelationshipRequest, error) {
	if !g.Client.AuthToken.Valid(ctx) {
		return nil, ce.CustomError(ce.NewInvalidAuthError("invalid authorization"))
	}

	i, err := g.App.InviteRepository.Find(ctx, g.FirestoreClient, input.Code)
	if err != nil {
		return nil, ce.CustomError(err)
	}
	if i.UserID == "" {
		return nil, ce.CustomError(ce.NewNotFoundError("招待コードが見つかりません"))
	}

	if i.UserID == g.UserID {
		return nil, ce.CustomError(ce.NewRequestError(ce.CodeMyInviteCode, "自身の招待コードです"))
	}

	uuid := g.Client.UUID.Get()

	rr := &model.RelationshipRequest{
		ID:         uuid,
		FollowerID: g.UserID,
		FollowedID: i.UserID,
		Status:     repository.RelationshipRequestStatusRequest,
		CreatedAt:  g.Client.Time.Now(),
		UpdatedAt:  g.Client.Time.Now(),
	}

	data, err := g.App.RelationshipRequestRepository.Find(ctx, g.FirestoreClient, rr)
	if GrpcErrorStatusCode(err) != codes.NotFound {
		if data.Status == repository.RelationshipRequestStatusRequest {
			return nil, ce.CustomError(ce.NewAlreadyExists("既に招待リクエスト済みです"))
		}
	}

	u, err := g.App.UserRepository.FindByUID(ctx, g.FirestoreClient, i.UserID)
	if err != nil {
		return nil, ce.CustomError(err)
	}

	tokens := g.App.PushTokenRepository.GetTokens(ctx, g.FirestoreClient, i.UserID)

	if len(tokens) > 0 {
		me, err := g.App.UserRepository.FindByUID(ctx, g.FirestoreClient, g.UserID)
		if err != nil {
			return nil, ce.CustomError(err)
		}

		r := task.NotificationRequest{
			Token:     tokens,
			Title:     me.DisplayName + "さんから共有メンバーの申請が届いています",
			Body:      me.DisplayName + "さんから共有メンバーの申請が届いています",
			URLScheme: "MyPage",
		}

		if _, err = g.Client.Task.PushNotification(r); err != nil {
			return nil, ce.CustomError(err)
		}
	}

	if err = g.App.RelationshipRequestRepository.Create(ctx, g.FirestoreClient, rr); err != nil {
		return nil, ce.CustomError(err)
	}

	rr.User = u

	return rr, nil
}

// AcceptRelationshipRequest 招待リクエストを承諾する
func (g *Graph) AcceptRelationshipRequest(ctx context.Context, followedID string) (*model.RelationshipRequest, error) {
	if !g.Client.AuthToken.Valid(ctx) {
		return nil, ce.CustomError(ce.NewInvalidAuthError("invalid authorization"))
	}
	rr1 := &model.RelationshipRequest{
		FollowerID: followedID,
		FollowedID: g.UserID,
		Status:     repository.RelationshipRequestStatusOK,
		UpdatedAt:  g.Client.Time.Now(),
	}
	rr2 := &model.RelationshipRequest{
		FollowerID: g.UserID,
		FollowedID: followedID,
		Status:     repository.RelationshipRequestStatusOK,
		UpdatedAt:  g.Client.Time.Now(),
	}

	isFollowedRequest := false
	rr2Data, _ := g.App.RelationshipRequestRepository.Find(ctx, g.FirestoreClient, rr2)
	if rr2Data.ID != "" {
		isFollowedRequest = true
	}

	batch := g.FirestoreClient.BulkWriter(ctx)
	if err := g.App.RelationshipRequestRepository.Update(ctx, g.FirestoreClient, batch, rr1); err != nil {
		return nil, ce.CustomError(err)
	}

	if isFollowedRequest {
		// 相手側もリクエストしていた場合はstatusを更新
		if err := g.App.RelationshipRequestRepository.Update(ctx, g.FirestoreClient, batch, rr2); err != nil {
			return nil, ce.CustomError(err)
		}
	}

	r1 := &model.Relationship{
		ID:         g.Client.UUID.Get(),
		FollowerID: g.UserID,
		FollowedID: followedID,
		CreatedAt:  g.Client.Time.Now(),
		UpdatedAt:  g.Client.Time.Now(),
	}
	r2 := &model.Relationship{
		ID:         g.Client.UUID.Get(),
		FollowerID: followedID,
		FollowedID: g.UserID,
		CreatedAt:  g.Client.Time.Now(),
		UpdatedAt:  g.Client.Time.Now(),
	}
	if err := g.App.RelationshipRepository.Create(ctx, g.FirestoreClient, batch, r1); err != nil {
		return nil, ce.CustomError(err)
	}
	if err := g.App.RelationshipRepository.Create(ctx, g.FirestoreClient, batch, r2); err != nil {
		return nil, ce.CustomError(err)
	}

	g.App.CommonRepository.Commit(ctx, batch)

	tokens := g.App.PushTokenRepository.GetTokens(ctx, g.FirestoreClient, followedID)

	if len(tokens) > 0 {
		u, err := g.App.UserRepository.FindByUID(ctx, g.FirestoreClient, g.UserID)
		if err != nil {
			return nil, ce.CustomError(err)
		}

		r := task.NotificationRequest{
			Token:     tokens,
			Title:     u.DisplayName + "さんと共有メンバーになりました",
			Body:      u.DisplayName + "さんと共有メンバーになりました",
			URLScheme: "MyPage",
		}

		if _, err = g.Client.Task.PushNotification(r); err != nil {
			return nil, ce.CustomError(err)
		}
	}

	return rr1, nil
}

// ngRelationshipRequest 招待リクエストを拒否する
func (g *Graph) NgRelationshipRequest(ctx context.Context, followedID string) (*model.RelationshipRequest, error) {
	if !g.Client.AuthToken.Valid(ctx) {
		return nil, ce.CustomError(ce.NewInvalidAuthError("invalid authorization"))
	}
	rr := &model.RelationshipRequest{
		FollowerID: followedID,
		FollowedID: g.UserID,
		Status:     repository.RelationshipRequestStatusNG,
		UpdatedAt:  g.Client.Time.Now(),
	}

	batch := g.FirestoreClient.BulkWriter(ctx)
	if err := g.App.RelationshipRequestRepository.Update(ctx, g.FirestoreClient, batch, rr); err != nil {
		return nil, ce.CustomError(err)
	}

	g.App.CommonRepository.Commit(ctx, batch)
	return rr, nil
}

// GetRelationshipRequests 共有の招待リクエストを取得する
func (g *Graph) GetRelationshipRequests(ctx context.Context, input model.InputRelationshipRequests, userSkip bool) (*model.RelationshipRequests, error) {
	t := g.Client.Time
	if !g.Client.AuthToken.Valid(ctx) {
		return nil, ce.CustomError(ce.NewInvalidAuthError("invalid authorization"))
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
		return nil, ce.CustomError(err)
	}

	userID := []string{}
	for _, i := range items {
		userID = append(userID, i.FollowerID)
	}
	users := []*model.User{}

	if !userSkip && len(userID) > 0 {
		users, err = g.App.UserRepository.FindInUID(ctx, g.FirestoreClient, userID)
		if err != nil {
			return nil, ce.CustomError(err)
		}

	}

	rres := make([]*model.RelationshipRequestEdge, len(items))

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

		rres[index] = &model.RelationshipRequestEdge{
			Node:   items[index],
			Cursor: i.FollowedID + "/" + i.FollowerID,
		}

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
