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
	"gopkg.in/go-playground/assert.v1"

	moq_repository "github.com/wheatandcat/memoir-backend/repository/moq"
)

func TestCreatePushToken(t *testing.T) {
	ctx := context.Background()

	client := &graph.Client{
		UUID: &uuidgen.UUID{},
		Time: &timegen.Time{},
	}

	date := client.Time.ParseInLocation("2020-01-01T00:00:00")

	pushToken := &model.PushToken{
		UserID:    "test",
		Token:     "token",
		DeviceID:  "deviceID",
		CreatedAt: date,
		UpdatedAt: date,
	}

	g := newGraph()

	pushTokenRepositoryMock := &moq_repository.PushTokenRepositoryInterfaceMock{
		CreateFunc: func(ctx context.Context, f *firestore.Client, userID string, i *model.PushToken) error {
			return nil
		},
	}

	g.App.PushTokenRepository = pushTokenRepositoryMock

	tests := []struct {
		name   string
		param  *model.NewPushToken
		result *model.PushToken
	}{
		{
			name: "Pushトークンを作成する",
			param: &model.NewPushToken{
				Token:    "token",
				DeviceID: "deviceID",
			},
			result: pushToken,
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			r, _ := g.CreatePushToken(ctx, td.param)
			diff := cmp.Diff(r, td.result)
			if diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			} else {
				assert.Equal(t, diff, "")
			}
		})
	}

}
