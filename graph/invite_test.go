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

func TestCreateInvite(t *testing.T) {
	ctx := context.Background()

	client := &graph.Client{
		UUID: &uuidgen.UUID{},
		Time: &timegen.Time{},
	}

	date := client.Time.ParseInLocation("2020-01-01T00:00:00")

	invite := &model.Invite{
		UserID:    "test",
		Code:      "ABCDEFGH",
		CreatedAt: date,
		UpdatedAt: date,
	}

	g := newGraph()

	inviteRepositoryMock := &repository.InviteRepositoryInterfaceMock{
		CreateFunc: func(ctx context.Context, f *firestore.Client, batch *firestore.WriteBatch, i *model.Invite) {
			return
		},
		FindByUserIDFunc: func(ctx context.Context, f *firestore.Client, userID string) (*model.Invite, error) {
			return &model.Invite{}, nil
		},
		CommitFunc: func(ctx context.Context, batch *firestore.WriteBatch) error {
			return nil
		},
	}

	g.App.InviteRepository = inviteRepositoryMock

	tests := []struct {
		name   string
		result *model.Invite
	}{
		{
			name:   "招待コードを作成する",
			result: invite,
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			r, _ := g.CreateInvite(ctx)
			diff := cmp.Diff(r, td.result)
			if diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			} else {
				assert.Equal(t, diff, "")
			}
		})
	}
}

func TestUpdateInvite(t *testing.T) {
	ctx := context.Background()

	client := &graph.Client{
		UUID: &uuidgen.UUID{},
		Time: &timegen.Time{},
	}

	date := client.Time.ParseInLocation("2020-01-01T00:00:00")

	invite := &model.Invite{
		UserID:    "test",
		Code:      "ABCDEFGH",
		CreatedAt: date,
		UpdatedAt: date,
	}

	g := newGraph()

	inviteRepositoryMock := &repository.InviteRepositoryInterfaceMock{
		CreateFunc: func(ctx context.Context, f *firestore.Client, batch *firestore.WriteBatch, i *model.Invite) {
			return
		},
		DeleteFunc: func(ctx context.Context, f *firestore.Client, batch *firestore.WriteBatch, code string) {
			return
		},
		FindByUserIDFunc: func(ctx context.Context, f *firestore.Client, userID string) (*model.Invite, error) {
			return &model.Invite{
				UserID:    "test",
				CreatedAt: date,
			}, nil
		},
		CommitFunc: func(ctx context.Context, batch *firestore.WriteBatch) error {
			return nil
		},
	}

	g.App.InviteRepository = inviteRepositoryMock

	tests := []struct {
		name   string
		result *model.Invite
	}{
		{
			name:   "招待コードを更新する",
			result: invite,
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			r, _ := g.UpdateInvite(ctx)
			diff := cmp.Diff(r, td.result)
			if diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			} else {
				assert.Equal(t, diff, "")
			}
		})
	}
}

func TestGetInviteByUseID(t *testing.T) {
	ctx := context.Background()

	client := &graph.Client{
		UUID: &uuidgen.UUID{},
		Time: &timegen.Time{},
	}

	date := client.Time.ParseInLocation("2020-01-01T00:00:00")

	invite := &model.Invite{
		UserID:    "test",
		Code:      "ABCDEFGH",
		CreatedAt: date,
		UpdatedAt: date,
	}

	g := newGraph()

	inviteRepositoryMock := &repository.InviteRepositoryInterfaceMock{
		FindByUserIDFunc: func(ctx context.Context, f *firestore.Client, userID string) (*model.Invite, error) {
			return invite, nil
		},
	}

	g.App.InviteRepository = inviteRepositoryMock

	tests := []struct {
		name   string
		result *model.Invite
	}{
		{
			name:   "招待コードを取得する",
			result: invite,
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			r, _ := g.GetInviteByUseID(ctx)
			diff := cmp.Diff(r, td.result)
			if diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			} else {
				assert.Equal(t, diff, "")
			}
		})
	}
}

func TestGetInviteByCode(t *testing.T) {
	ctx := context.Background()

	client := &graph.Client{
		UUID: &uuidgen.UUID{},
		Time: &timegen.Time{},
	}

	date := client.Time.ParseInLocation("2020-01-01T00:00:00")

	invite := &model.Invite{
		UserID:    "test",
		Code:      "ABCDEFGH",
		CreatedAt: date,
		UpdatedAt: date,
	}
	user := &model.User{
		ID:        "test",
		CreatedAt: date,
		UpdatedAt: date,
	}

	g := newGraph()

	inviteRepositoryMock := &repository.InviteRepositoryInterfaceMock{
		FindFunc: func(ctx context.Context, f *firestore.Client, code string) (*model.Invite, error) {
			return invite, nil
		},
	}
	userRepositoryMock := &repository.UserRepositoryInterfaceMock{
		FindByUIDFunc: func(ctx context.Context, f *firestore.Client, uid string) (*model.User, error) {
			return user, nil
		},
	}

	g.App.InviteRepository = inviteRepositoryMock
	g.App.UserRepository = userRepositoryMock

	tests := []struct {
		name   string
		param  string
		result *model.User
	}{
		{
			name:   "コードから招待を取得する",
			param:  "ABCDEFGH",
			result: user,
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			r, _ := g.GetInviteByCode(ctx, td.param)
			diff := cmp.Diff(r, td.result)
			if diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			} else {
				assert.Equal(t, diff, "")
			}
		})
	}
}
