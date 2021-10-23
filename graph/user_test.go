package graph_test

import (
	"context"
	"testing"

	"cloud.google.com/go/firestore"
	"github.com/google/go-cmp/cmp"
	"github.com/wheatandcat/memoir-backend/auth"
	"github.com/wheatandcat/memoir-backend/client/timegen"
	"github.com/wheatandcat/memoir-backend/client/uuidgen"
	"github.com/wheatandcat/memoir-backend/graph"
	"github.com/wheatandcat/memoir-backend/graph/model"
	"github.com/wheatandcat/memoir-backend/repository"
	"gopkg.in/go-playground/assert.v1"

	moq_repository "github.com/wheatandcat/memoir-backend/repository/moq"
)

type contextKey struct {
	name string
}

func TestUpdateUser(t *testing.T) {

	u := &auth.User{
		ID:          "test",
		FirebaseUID: "test",
	}

	ctx := context.WithValue(context.Background(), &contextKey{"user"}, u)

	client := &graph.Client{
		UUID: &uuidgen.UUID{},
		Time: &timegen.Time{},
	}

	g := newGraph()

	userRepositoryMock := &moq_repository.UserRepositoryInterfaceMock{
		UpdateFunc: func(ctx context.Context, f *firestore.Client, u *model.User) error {
			return nil
		},
	}
	g.App.UserRepository = userRepositoryMock

	tests := []struct {
		name   string
		param  *model.UpdateUser
		result *model.User
	}{
		{
			name: "ユーザーを更新する",
			param: &model.UpdateUser{
				DisplayName: "test-name",
			},
			result: &model.User{
				ID:          "test",
				DisplayName: "test-name",
				UpdatedAt:   client.Time.ParseInLocation("2020-01-01T00:00:00"),
			},
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			r, _ := g.UpdateUser(ctx, td.param)
			diff := cmp.Diff(r, td.result)
			if diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			} else {
				assert.Equal(t, diff, "")
			}
		})
	}

}

func TestCreateUser(t *testing.T) {
	type __ = context.Context
	type ___ = *firestore.Client

	u := &auth.User{
		ID:          "test",
		FirebaseUID: "test",
	}

	ctx := context.WithValue(context.Background(), &contextKey{"user"}, u)

	client := &graph.Client{
		UUID: &uuidgen.UUID{},
		Time: &timegen.Time{},
	}

	g := newGraph()

	userRepositoryMock := &moq_repository.UserRepositoryInterfaceMock{
		CreateFunc: func(_ __, _ ___, u *model.User) error {
			return nil
		},
	}
	g.App.UserRepository = userRepositoryMock

	tests := []struct {
		name   string
		param  *model.NewUser
		result *model.User
	}{
		{
			name: "ユーザーを更新する",
			param: &model.NewUser{
				ID: "test",
			},
			result: &model.User{
				ID:        "test",
				CreatedAt: client.Time.ParseInLocation("2020-01-01T00:00:00"),
				UpdatedAt: client.Time.ParseInLocation("2020-01-01T00:00:00"),
			},
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			r, _ := g.CreateUser(ctx, td.param)
			diff := cmp.Diff(r, td.result)
			if diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			} else {
				assert.Equal(t, diff, "")
			}
		})
	}

}

func TestCreateAuthUser(t *testing.T) {
	type __ = context.Context
	type ___ = *firestore.Client

	u := &auth.User{
		ID:          "test",
		FirebaseUID: "test",
	}

	ctx := context.WithValue(context.Background(), &contextKey{"user"}, u)

	g := newGraph()

	tests := []struct {
		name   string
		param  *model.NewAuthUser
		mock   func()
		result *model.AuthUser
	}{
		{
			name: "認証済みユーザーを作成 かつ 既にユーザーは作成済み",
			param: &model.NewAuthUser{
				ID:        "test",
				IsNewUser: false,
			},
			mock: func() {
				userRepositoryMock := &moq_repository.UserRepositoryInterfaceMock{
					ExistByFirebaseUIDFunc: func(_ __, _ ___, fUID string) (bool, error) {
						return true, nil
					},
				}
				g.App.UserRepository = userRepositoryMock
			},
			result: &model.AuthUser{
				ID: "test",
			},
		},
		{
			name: "認証済みユーザーを作成 かつ 既にユーザーは未作成",
			param: &model.NewAuthUser{
				ID:        "test",
				IsNewUser: false,
			},
			mock: func() {
				userRepositoryMock := &moq_repository.UserRepositoryInterfaceMock{
					ExistByFirebaseUIDFunc: func(_ __, _ ___, fUID string) (bool, error) {
						return false, nil
					},
					CreateFunc: func(_ __, _ ___, u *model.User) error {
						return nil
					},
					UpdateFirebaseUIDFunc: func(_ __, _ ___, user *repository.User) error {
						return nil
					},
				}
				g.App.UserRepository = userRepositoryMock
			},
			result: &model.AuthUser{
				ID: "test",
			},
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			td.mock()
			r, _ := g.CreateAuthUser(ctx, td.param)
			diff := cmp.Diff(r, td.result)
			if diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			} else {
				assert.Equal(t, diff, "")
			}
		})
	}

}
