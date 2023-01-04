package graph_test

import (
	"context"
	"testing"

	"cloud.google.com/go/firestore"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/go-playground/assert.v1"

	"github.com/wheatandcat/memoir-backend/client/timegen"
	"github.com/wheatandcat/memoir-backend/client/uuidgen"
	"github.com/wheatandcat/memoir-backend/graph"
	"github.com/wheatandcat/memoir-backend/graph/model"
	"github.com/wheatandcat/memoir-backend/repository"

	moq_repository "github.com/wheatandcat/memoir-backend/repository/moq"
)

func TestCreateRelationshipRequest(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	client := &graph.Client{
		UUID: &uuidgen.UUID{},
		Time: &timegen.Time{},
	}

	date := client.Time.ParseInLocation("2020-01-01T00:00:00")

	u := &model.User{
		ID:          "test2",
		DisplayName: "test",
		CreatedAt:   date,
		UpdatedAt:   date,
	}

	rr := &model.RelationshipRequest{
		ID:         "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		FollowerID: "test",
		FollowedID: "test2",
		Status:     repository.RelationshipRequestStatusRequest,
		CreatedAt:  date,
		UpdatedAt:  date,
		User:       u,
	}

	invite := &model.Invite{
		UserID:    "test2",
		Code:      "ABCDEFGH",
		CreatedAt: date,
		UpdatedAt: date,
	}

	g := newGraph(ctx)

	inviteRepositoryMock := &moq_repository.InviteRepositoryInterfaceMock{
		FindFunc: func(ctx context.Context, f *firestore.Client, code string) (*model.Invite, error) {
			return invite, nil
		},
	}

	relationshipRequestRepositoryMock := &moq_repository.RelationshipRequestInterfaceMock{
		FindFunc: func(ctx context.Context, f *firestore.Client, i *model.RelationshipRequest) (*model.RelationshipRequest, error) {
			return nil, status.Errorf(codes.NotFound, "%q not found", "")
		},
		CreateFunc: func(ctx context.Context, f *firestore.Client, i *model.RelationshipRequest) error {
			return nil
		},
	}

	userRepositoryMock := &moq_repository.UserRepositoryInterfaceMock{
		FindByUIDFunc: func(ctx context.Context, f *firestore.Client, uid string) (*model.User, error) {
			return u, nil
		},
	}

	pushTokenRepositoryMock := &moq_repository.PushTokenRepositoryInterfaceMock{
		GetTokensFunc: func(ctx context.Context, f *firestore.Client, uid string) []string {
			return []string{}
		},
	}

	g.App.InviteRepository = inviteRepositoryMock
	g.App.RelationshipRequestRepository = relationshipRequestRepositoryMock
	g.App.UserRepository = userRepositoryMock
	g.App.PushTokenRepository = pushTokenRepositoryMock

	tests := []struct {
		name   string
		param  model.NewRelationshipRequest
		result *model.RelationshipRequest
	}{
		{
			name: "招待を申請をリクエストする",
			param: model.NewRelationshipRequest{
				Code: "ABCDEFGH",
			},
			result: rr,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r, _ := g.CreateRelationshipRequest(ctx, tt.param)
			diff := cmp.Diff(r, tt.result)
			if diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			} else {
				assert.Equal(t, diff, "")
			}
		})
	}
}

func TestAcceptRelationshipRequest(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	client := &graph.Client{
		UUID: &uuidgen.UUID{},
		Time: &timegen.Time{},
	}

	date := client.Time.ParseInLocation("2020-01-01T00:00:00")

	rr := &model.RelationshipRequest{
		FollowerID: "test",
		FollowedID: "test",
		Status:     repository.RelationshipRequestStatusOK,
		UpdatedAt:  date,
	}
	u := &model.User{
		ID:          "test2",
		DisplayName: "test",
		CreatedAt:   date,
		UpdatedAt:   date,
	}

	g := newGraph(ctx)

	relationshipRequestRepositoryMock := &moq_repository.RelationshipRequestInterfaceMock{
		UpdateFunc: func(ctx context.Context, f *firestore.Client, batch *firestore.Transaction, i *model.RelationshipRequest) error {
			return nil
		},
		FindFunc: func(ctx context.Context, f *firestore.Client, i *model.RelationshipRequest) (*model.RelationshipRequest, error) {
			return &model.RelationshipRequest{}, nil
		},
	}
	relationshipRepositoryMock := &moq_repository.RelationshipInterfaceMock{
		CreateFunc: func(ctx context.Context, f *firestore.Client, batch *firestore.Transaction, i *model.Relationship) error {
			return nil
		},
	}
	pushTokenRepositoryMock := &moq_repository.PushTokenRepositoryInterfaceMock{
		GetTokensFunc: func(ctx context.Context, f *firestore.Client, uid string) []string {
			return []string{}
		},
	}
	userRepositoryMock := &moq_repository.UserRepositoryInterfaceMock{
		FindByUIDFunc: func(ctx context.Context, f *firestore.Client, uid string) (*model.User, error) {
			return u, nil
		},
	}

	g.App.RelationshipRequestRepository = relationshipRequestRepositoryMock
	g.App.RelationshipRepository = relationshipRepositoryMock
	g.App.UserRepository = userRepositoryMock
	g.App.PushTokenRepository = pushTokenRepositoryMock

	tests := []struct {
		name   string
		param  string
		result *model.RelationshipRequest
	}{
		{
			name:   "招待リクエストを承諾する",
			param:  "test",
			result: rr,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r, _ := g.AcceptRelationshipRequest(ctx, tt.param)
			diff := cmp.Diff(r, tt.result)
			if diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			} else {
				assert.Equal(t, diff, "")
			}
		})
	}
}

func TestNgRelationshipRequest(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	client := &graph.Client{
		UUID: &uuidgen.UUID{},
		Time: &timegen.Time{},
	}

	date := client.Time.ParseInLocation("2020-01-01T00:00:00")

	rr := &model.RelationshipRequest{
		FollowerID: "test",
		FollowedID: "test",
		Status:     repository.RelationshipRequestStatusNG,
		UpdatedAt:  date,
	}

	g := newGraph(ctx)

	relationshipRequestRepositoryMock := &moq_repository.RelationshipRequestInterfaceMock{
		UpdateFunc: func(ctx context.Context, f *firestore.Client, batch *firestore.Transaction, i *model.RelationshipRequest) error {
			return nil
		},
	}

	g.App.RelationshipRequestRepository = relationshipRequestRepositoryMock

	tests := []struct {
		name   string
		param  string
		result *model.RelationshipRequest
	}{
		{
			name:   "招待リクエストを拒否する",
			param:  "test",
			result: rr,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r, _ := g.NgRelationshipRequest(ctx, tt.param)
			diff := cmp.Diff(r, tt.result)
			if diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			} else {
				assert.Equal(t, diff, "")
			}
		})
	}
}

func TestGetRelationshipRequests(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	client := &graph.Client{
		UUID: &uuidgen.UUID{},
		Time: &timegen.Time{},
	}

	date := client.Time.ParseInLocation("2020-01-01T00:00:00")

	user := &model.User{
		ID:          "test",
		DisplayName: "",
		Image:       "",
		CreatedAt:   date,
		UpdatedAt:   date,
	}

	rr := []*model.RelationshipRequest{{
		ID:         "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		FollowerID: "test",
		FollowedID: "test",
		Status:     repository.RelationshipRequestStatusRequest,
		CreatedAt:  date,
		UpdatedAt:  date,
		User:       user,
	}}

	rres := []*model.RelationshipRequestEdge{{
		Node: &model.RelationshipRequest{
			ID:         "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			FollowerID: "test",
			FollowedID: "test",
			Status:     repository.RelationshipRequestStatusRequest,
			CreatedAt:  date,
			UpdatedAt:  date,
			User:       user,
		},
		Cursor: "test/test",
	}}

	rrs := &model.RelationshipRequests{
		PageInfo: &model.PageInfo{
			EndCursor:   "test/test",
			HasNextPage: false,
		},
		Edges: rres,
	}

	users := []*model.User{{
		ID:          "test",
		DisplayName: "",
		Image:       "",
		CreatedAt:   date,
		UpdatedAt:   date,
	}}

	g := newGraph(ctx)

	relationshipRequestRepositoryMock := &moq_repository.RelationshipRequestInterfaceMock{
		FindByFollowedIDFunc: func(ctx context.Context, f *firestore.Client, userID string, first int, cursor repository.RelationshipRequestCursor) ([]*model.RelationshipRequest, error) {
			return rr, nil
		},
	}

	userRepositoryInterfaceMock := &moq_repository.UserRepositoryInterfaceMock{
		FindInUIDFunc: func(ctx context.Context, f *firestore.Client, uid []string) ([]*model.User, error) {
			return users, nil
		},
	}

	g.App.UserRepository = userRepositoryInterfaceMock
	g.App.RelationshipRequestRepository = relationshipRequestRepositoryMock

	after := ""

	tests := []struct {
		name     string
		param    model.InputRelationshipRequests
		userSkip bool
		result   *model.RelationshipRequests
	}{
		{
			name: "申請されている招待を取得する",
			param: model.InputRelationshipRequests{
				First: 5,
				After: &after,
			},
			userSkip: false,
			result:   rrs,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r, _ := g.GetRelationshipRequests(ctx, tt.param, tt.userSkip)
			diff := cmp.Diff(r, tt.result)
			if diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			} else {
				assert.Equal(t, diff, "")
			}
		})
	}
}
