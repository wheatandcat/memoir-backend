package graph_test

import (
	"context"
	"testing"

	"cloud.google.com/go/firestore"
	"github.com/google/go-cmp/cmp"
	"gopkg.in/go-playground/assert.v1"

	"github.com/wheatandcat/memoir-backend/client/timegen"
	"github.com/wheatandcat/memoir-backend/client/uuidgen"
	"github.com/wheatandcat/memoir-backend/graph"
	"github.com/wheatandcat/memoir-backend/graph/model"
	ce "github.com/wheatandcat/memoir-backend/usecase/custom_error"

	moq_repository "github.com/wheatandcat/memoir-backend/repository/moq"
)

func TestCreateInvite(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	client := &graph.Client{
		UUID: &uuidgen.UUID{},
		Time: &timegen.Time{},
	}

	date := client.Time.ParseInLocation("2020-01-01T00:00:00")

	type appParam struct {
		Invite *model.Invite
	}

	appFunc := func(param appParam) graph.Graph {
		g := newGraph(ctx)

		inviteRepositoryMock := &moq_repository.InviteRepositoryInterfaceMock{
			CreateFunc: func(ctx context.Context, f *firestore.Client, batch *firestore.Transaction, i *model.Invite) error {
				return nil
			},
			FindByUserIDFunc: func(ctx context.Context, f *firestore.Client, userID string) (*model.Invite, error) {
				return param.Invite, nil
			},
		}

		g.App.InviteRepository = inviteRepositoryMock

		return g
	}

	tests := []struct {
		name    string
		param   appParam
		result  *model.Invite
		errCode string
	}{
		{
			name: "招待コードを作成する",
			param: appParam{
				Invite: &model.Invite{},
			},
			result: &model.Invite{
				UserID:    "test",
				Code:      "ABCDEFGH",
				CreatedAt: date,
				UpdatedAt: date,
			},
			errCode: "",
		},
		{
			name: "エラー:自身の招待コードです",
			param: appParam{
				Invite: &model.Invite{
					UserID: "test",
				},
			},
			result:  nil,
			errCode: ce.CodeMyInviteCode,
		},
	}

	for _, td := range tests {
		td := td
		t.Run(td.name, func(t *testing.T) {
			g := appFunc(td.param)
			r, err := g.CreateInvite(ctx)
			if td.result != nil {
				diff := cmp.Diff(r, td.result)
				if diff != "" {
					t.Errorf("differs: (-got +want)\n%s", diff)
				} else {
					assert.Equal(t, diff, "")
				}
			} else {
				diff := cmp.Diff(getErrorCode(err), td.errCode)
				if diff != "" {
					t.Errorf("differs: (-got +want)\n%s", diff)
				} else {
					assert.Equal(t, diff, "")
				}
			}
		})
	}
}

func TestUpdateInvite(t *testing.T) {
	t.Parallel()
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

	g := newGraph(ctx)

	inviteRepositoryMock := &moq_repository.InviteRepositoryInterfaceMock{
		CreateFunc: func(ctx context.Context, f *firestore.Client, batch *firestore.Transaction, i *model.Invite) error {
			return nil
		},
		DeleteFunc: func(ctx context.Context, f *firestore.Client, batch *firestore.Transaction, code string) error {
			return nil
		},
		FindByUserIDFunc: func(ctx context.Context, f *firestore.Client, userID string) (*model.Invite, error) {
			return &model.Invite{
				UserID:    "test",
				CreatedAt: date,
			}, nil
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

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r, _ := g.UpdateInvite(ctx)
			diff := cmp.Diff(r, tt.result)
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

	g := newGraph(ctx)

	inviteRepositoryMock := &moq_repository.InviteRepositoryInterfaceMock{
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

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			r, _ := g.GetInviteByUseID(ctx)
			diff := cmp.Diff(r, tt.result)
			if diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			} else {
				assert.Equal(t, diff, "")
			}
		})
	}
}

func TestGetInviteByCode(t *testing.T) {
	t.Parallel()
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

	g := newGraph(ctx)

	inviteRepositoryMock := &moq_repository.InviteRepositoryInterfaceMock{
		FindFunc: func(ctx context.Context, f *firestore.Client, code string) (*model.Invite, error) {
			return invite, nil
		},
	}
	userRepositoryMock := &moq_repository.UserRepositoryInterfaceMock{
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

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r, _ := g.GetInviteByCode(ctx, tt.param)
			diff := cmp.Diff(r, tt.result)
			if diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			} else {
				assert.Equal(t, diff, "")
			}
		})
	}
}
