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
	"gopkg.in/go-playground/assert.v1"
)

func TestDeleteRelationship(t *testing.T) {
	ctx := context.Background()

	rr := &model.Relationship{
		FollowerID: "test",
		FollowedID: "test",
	}

	g := newGraph()

	relationshipRepositoryMock := &repository.RelationshipInterfaceMock{
		DeleteFunc: func(ctx context.Context, f *firestore.Client, batch *firestore.WriteBatch, i *model.Relationship) {
		},
	}
	commonRepositoryMock := &repository.CommonRepositoryInterfaceMock{
		CommitFunc: func(ctx context.Context, batch *firestore.WriteBatch) error {
			return nil
		},
	}

	g.App.RelationshipRepository = relationshipRepositoryMock
	g.App.CommonRepository = commonRepositoryMock

	tests := []struct {
		name   string
		param  string
		result *model.Relationship
	}{
		{
			name:   "共有ユーザーを解除する",
			param:  "test",
			result: rr,
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			r, _ := g.DeleteRelationship(ctx, td.param)
			diff := cmp.Diff(r, td.result)
			if diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			} else {
				assert.Equal(t, diff, "")
			}
		})
	}

}

func TestGetRelationships(t *testing.T) {
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

	rr := []*model.Relationship{{
		ID:         "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		FollowerID: "test",
		FollowedID: "test",
		CreatedAt:  date,
		UpdatedAt:  date,
		User:       user,
	}}

	rres := []*model.RelationshipEdge{{
		Node: &model.Relationship{
			ID:         "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			FollowerID: "test",
			FollowedID: "test",
			CreatedAt:  date,
			UpdatedAt:  date,
			User:       user,
		},
		Cursor: "test/test",
	}}

	rrs := &model.Relationships{
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

	relationshipRepositoryMock := &repository.RelationshipInterfaceMock{
		FindByFollowedIDFunc: func(ctx context.Context, f *firestore.Client, userID string, first int, cursor repository.RelationshipCursor) ([]*model.Relationship, error) {
			return rr, nil
		},
	}

	userRepositoryInterfaceMock := &repository.UserRepositoryInterfaceMock{
		FindInUIDFunc: func(ctx context.Context, f *firestore.Client, uid []string) ([]*model.User, error) {
			return users, nil
		},
	}

	g.App.UserRepository = userRepositoryInterfaceMock
	g.App.RelationshipRepository = relationshipRepositoryMock

	after := ""

	tests := []struct {
		name     string
		param    model.InputRelationships
		userSkip bool
		result   *model.Relationships
	}{
		{
			name: "共有メンバーを取得する",
			param: model.InputRelationships{
				First: 5,
				After: &after,
			},
			userSkip: false,
			result:   rrs,
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			r, _ := g.GetRelationships(ctx, td.param, td.userSkip)
			diff := cmp.Diff(r, td.result)
			if diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			} else {
				assert.Equal(t, diff, "")
			}
		})
	}
}
