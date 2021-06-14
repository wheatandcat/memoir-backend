package graph_test

import (
	"context"
	"testing"

	"cloud.google.com/go/firestore"
	"github.com/google/go-cmp/cmp"
	"github.com/wheatandcat/memoir-backend/client/timegen"
	"github.com/wheatandcat/memoir-backend/client/uuidgen"
	"github.com/wheatandcat/memoir-backend/graph"
	"github.com/wheatandcat/memoir-backend/graph/model"
	"github.com/wheatandcat/memoir-backend/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/go-playground/assert.v1"
)

func TestCreateRelationshipRequest(t *testing.T) {
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

	g := newGraph()

	inviteRepositoryMock := &repository.InviteRepositoryInterfaceMock{
		FindFunc: func(ctx context.Context, f *firestore.Client, code string) (*model.Invite, error) {
			return invite, nil
		},
	}

	relationshipRequestRepositoryMock := &repository.RelationshipRequestInterfaceMock{
		FindFunc: func(ctx context.Context, f *firestore.Client, i *model.RelationshipRequest) (*model.RelationshipRequest, error) {
			return nil, status.Errorf(codes.NotFound, "%q not found", "")
		},
		CreateFunc: func(ctx context.Context, f *firestore.Client, i *model.RelationshipRequest) error {
			return nil
		},
	}

	userRepositoryMock := &repository.UserRepositoryInterfaceMock{
		FindByUIDFunc: func(ctx context.Context, f *firestore.Client, uid string) (*model.User, error) {
			return u, nil
		},
	}

	g.App.InviteRepository = inviteRepositoryMock
	g.App.RelationshipRequestRepository = relationshipRequestRepositoryMock
	g.App.UserRepository = userRepositoryMock

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

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			r, _ := g.CreateRelationshipRequest(ctx, td.param)
			diff := cmp.Diff(r, td.result)
			if diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			} else {
				assert.Equal(t, diff, "")
			}
		})
	}
}

func TestAcceptRelationshipRequest(t *testing.T) {
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

	g := newGraph()

	relationshipRequestRepositoryMock := &repository.RelationshipRequestInterfaceMock{
		UpdateFunc: func(ctx context.Context, f *firestore.Client, batch *firestore.WriteBatch, i *model.RelationshipRequest) {
		},
		FindFunc: func(ctx context.Context, f *firestore.Client, i *model.RelationshipRequest) (*model.RelationshipRequest, error) {
			return &model.RelationshipRequest{}, nil
		},
	}
	relationshipRepositoryMock := &repository.RelationshipInterfaceMock{
		CreateFunc: func(ctx context.Context, f *firestore.Client, batch *firestore.WriteBatch, i *model.Relationship) {
		},
	}
	commonRepositoryMock := &repository.CommonRepositoryInterfaceMock{
		CommitFunc: func(ctx context.Context, batch *firestore.WriteBatch) error {
			return nil
		},
	}

	g.App.RelationshipRequestRepository = relationshipRequestRepositoryMock
	g.App.RelationshipRepository = relationshipRepositoryMock
	g.App.CommonRepository = commonRepositoryMock

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

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			r, _ := g.AcceptRelationshipRequest(ctx, td.param)
			diff := cmp.Diff(r, td.result)
			if diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			} else {
				assert.Equal(t, diff, "")
			}
		})
	}
}

func TestNgRelationshipRequest(t *testing.T) {
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

	g := newGraph()

	relationshipRequestRepositoryMock := &repository.RelationshipRequestInterfaceMock{
		UpdateFunc: func(ctx context.Context, f *firestore.Client, batch *firestore.WriteBatch, i *model.RelationshipRequest) {
		},
	}
	commonRepositoryMock := &repository.CommonRepositoryInterfaceMock{
		CommitFunc: func(ctx context.Context, batch *firestore.WriteBatch) error {
			return nil
		},
	}

	g.App.RelationshipRequestRepository = relationshipRequestRepositoryMock
	g.App.CommonRepository = commonRepositoryMock

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

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			r, _ := g.NgRelationshipRequest(ctx, td.param)
			diff := cmp.Diff(r, td.result)
			if diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			} else {
				assert.Equal(t, diff, "")
			}
		})
	}
}

func TestGetRelationshipRequests(t *testing.T) {
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

	g := newGraph()

	relationshipRequestRepositoryMock := &repository.RelationshipRequestInterfaceMock{
		FindByFollowedIDFunc: func(ctx context.Context, f *firestore.Client, userID string, first int, cursor repository.RelationshipRequestCursor) ([]*model.RelationshipRequest, error) {
			return rr, nil
		},
	}

	userRepositoryInterfaceMock := &repository.UserRepositoryInterfaceMock{
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

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			r, _ := g.GetRelationshipRequests(ctx, td.param, td.userSkip)
			diff := cmp.Diff(r, td.result)
			if diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			} else {
				assert.Equal(t, diff, "")
			}
		})
	}
}
