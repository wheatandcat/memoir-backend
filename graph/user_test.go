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
	moq_usecase_auth "github.com/wheatandcat/memoir-backend/usecase/auth/moq"
	ce "github.com/wheatandcat/memoir-backend/usecase/custom_error"
)

type contextKey struct {
	name string
}

type __ = context.Context
type ___ = *firestore.Client
type ____ = *firestore.WriteBatch

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
		UpdateFunc: func(_ __, _ ___, u *model.User) error {
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
				authUseCaseMock := &moq_usecase_auth.UseCaseMock{
					CreateAuthUserFunc: func(_ __, _ ___, input *model.NewAuthUser, u *repository.User, mu *model.AuthUser) (string, error) {
						return "", nil
					},
				}
				g.App.AuthUseCase = authUseCaseMock

			},
			result: &model.AuthUser{
				ID:        "test",
				CreatedAt: client.Time.ParseInLocation("2020-01-01T00:00:00"),
				UpdatedAt: client.Time.ParseInLocation("2020-01-01T00:00:00"),
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
				authUseCaseMock := &moq_usecase_auth.UseCaseMock{
					CreateAuthUserFunc: func(_ __, _ ___, input *model.NewAuthUser, u *repository.User, mu *model.AuthUser) (string, error) {
						return "", nil
					},
				}
				g.App.AuthUseCase = authUseCaseMock
			},
			result: &model.AuthUser{
				ID:        "test",
				CreatedAt: client.Time.ParseInLocation("2020-01-01T00:00:00"),
				UpdatedAt: client.Time.ParseInLocation("2020-01-01T00:00:00"),
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

func TestGetUser(t *testing.T) {
	ctx := context.Background()

	g := newGraph()

	u := &model.User{
		ID: "test",
	}

	userRepositoryMock := &moq_repository.UserRepositoryInterfaceMock{
		FindByUIDFunc: func(_ __, _ ___, _ string) (*model.User, error) {
			return u, nil
		},
	}
	g.App.UserRepository = userRepositoryMock

	tests := []struct {
		name   string
		result *model.User
	}{
		{
			name: "ユーザーを取得する",
			result: &model.User{
				ID: "test",
			},
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			r, _ := g.GetUser(ctx)
			diff := cmp.Diff(r, td.result)
			if diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			} else {
				assert.Equal(t, diff, "")
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	ctx := context.Background()

	g := newGraph()

	relationshipMock := &moq_repository.RelationshipInterfaceMock{
		ExistByFollowedIDFunc: func(_ __, _ ___, _ string) (bool, error) {
			return false, nil
		},
	}

	userRepositoryMock := &moq_repository.UserRepositoryInterfaceMock{
		DeleteFunc: func(_ __, _ ___, _ ____, uid string) {
			return
		},
	}

	inviteRepositoryMock := &moq_repository.InviteRepositoryInterfaceMock{
		DeleteByUserIDFunc: func(_ __, _ ___, _ ____, uid string) error {
			return nil
		},
	}

	relationshipRequestMock := &moq_repository.RelationshipRequestInterfaceMock{
		DeleteByFollowedIDFunc: func(_ __, _ ___, _ ____, uid string) error {
			return nil
		},
		DeleteByFollowerIDFunc: func(_ __, _ ___, _ ____, uid string) error {
			return nil
		},
	}

	authMock := &moq_usecase_auth.UseCaseMock{
		DeleteAuthUserFunc: func(_ __, _ ___, uid string) error {
			return nil
		},
	}

	commonRepositoryMock := &moq_repository.CommonRepositoryInterfaceMock{
		CommitFunc: func(_ __, _ ____) error {
			return nil
		},
	}

	g.App.RelationshipRepository = relationshipMock
	g.App.UserRepository = userRepositoryMock
	g.App.InviteRepository = inviteRepositoryMock
	g.App.RelationshipRequestRepository = relationshipRequestMock
	g.App.AuthUseCase = authMock
	g.App.CommonRepository = commonRepositoryMock

	tests := []struct {
		name   string
		mock   func()
		result *model.User
		err    error
	}{
		{
			name: "ユーザーを削除する",
			mock: func() {},
			result: &model.User{
				ID: "test",
			},
			err: nil,
		},
		{
			name: "共有メンバーが存在する",
			mock: func() {
				relationshipMock = &moq_repository.RelationshipInterfaceMock{
					ExistByFollowedIDFunc: func(_ __, _ ___, _ string) (bool, error) {
						return true, nil
					},
				}
				g.App.RelationshipRepository = relationshipMock
			},
			result: nil,
			err:    ce.NewRequestError(ce.HasRelationshipByDeleteUser, "共有メンバーが設定されています"),
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			td.mock()
			r, err := g.DeleteUser(ctx)
			diff := cmp.Diff(r, td.result)
			if diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			} else {
				assert.Equal(t, diff, "")
			}
			if td.err != nil {
				checkErrorCode(t, err, td.err)
			}
		})
	}
}
