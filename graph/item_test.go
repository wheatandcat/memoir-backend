package graph_test

import (
	"context"
	"testing"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/google/go-cmp/cmp"
	"github.com/wheatandcat/memoir-backend/client/timegen"
	"github.com/wheatandcat/memoir-backend/client/uuidgen"
	"github.com/wheatandcat/memoir-backend/graph"
	"github.com/wheatandcat/memoir-backend/graph/model"
	"github.com/wheatandcat/memoir-backend/repository"
	"gopkg.in/go-playground/assert.v1"
)

func TestGetItemsInDate(t *testing.T) {
	ctx := context.Background()

	client := &graph.Client{
		UUID: &uuidgen.UUID{},
		Time: &timegen.Time{},
	}

	date := client.Time.ParseInLocation("2019-01-01T00:00:00")

	items := []*model.Item{{
		ID:         "test1",
		CategoryID: 1,
		Title:      "test-title",
		Date:       date,
		CreatedAt:  date,
		UpdatedAt:  date,
	}}

	g := newGraph()

	itemRepositoryMock := &repository.ItemRepositoryInterfaceMock{
		GetItemsInDateFunc: func(ctx context.Context, f *firestore.Client, userID string, date time.Time) ([]*model.Item, error) {
			return items, nil
		},
	}

	g.App.ItemRepository = itemRepositoryMock

	tests := []struct {
		name   string
		param  time.Time
		result []*model.Item
	}{
		{
			name:   "日付でアイテムを取得する",
			param:  date,
			result: items,
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			r, _ := g.GetItemsInDate(ctx, td.param)
			diff := cmp.Diff(r, td.result)
			if diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			} else {
				assert.Equal(t, diff, "")
			}
		})
	}
}
