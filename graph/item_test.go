package graph_test

import (
	"context"
	"testing"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/google/go-cmp/cmp"
	"gopkg.in/go-playground/assert.v1"

	"github.com/wheatandcat/memoir-backend/client/timegen"
	"github.com/wheatandcat/memoir-backend/client/uuidgen"
	"github.com/wheatandcat/memoir-backend/graph"
	"github.com/wheatandcat/memoir-backend/graph/model"
	"github.com/wheatandcat/memoir-backend/repository"

	moq_repository "github.com/wheatandcat/memoir-backend/repository/moq"
)

func TestGetItemsInDate(t *testing.T) {
	t.Parallel()
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

	itemRepositoryMock := &moq_repository.ItemRepositoryInterfaceMock{
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

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r, _ := g.GetItemsInDate(ctx, tt.param)
			diff := cmp.Diff(r, tt.result)
			if diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			} else {
				assert.Equal(t, diff, "")
			}
		})
	}
}

func TestGetItemsInPeriod(t *testing.T) {
	t.Parallel()
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
		UserID:     "test-user",
		Date:       date,
		CreatedAt:  date,
		UpdatedAt:  date,
	}}

	rr := []*model.Relationship{{
		ID:         "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		FollowerID: "test",
		FollowedID: "test",
		CreatedAt:  date,
		UpdatedAt:  date,
	}}

	g := newGraph()

	itemRepositoryMock := &moq_repository.ItemRepositoryInterfaceMock{
		GetItemUserMultipleInPeriodFunc: func(_ context.Context, _ *firestore.Client, _ repository.SearchItemParam, _ int, _ repository.ItemsInPeriodCursor) ([]*model.Item, error) {
			return items, nil
		},
		GetCountUserMultipleInPeriodFunc: func(ctx context.Context, f *firestore.Client, sip repository.SearchItemParam) (int, error) {
			return len(items), nil
		},
	}

	relationshipRepositoryMock := &moq_repository.RelationshipInterfaceMock{
		FindByFollowedIDFunc: func(ctx context.Context, f *firestore.Client, userID string, first int, cursor repository.RelationshipCursor) ([]*model.Relationship, error) {
			return rr, nil
		},
	}

	g.App.ItemRepository = itemRepositoryMock
	g.App.RelationshipRepository = relationshipRepositoryMock
	after := ""

	iipe := []*model.ItemsInPeriodEdge{{
		Cursor: "test-user/test1",
		Node: &model.Item{
			ID:         "test1",
			CategoryID: 1,
			Title:      "test-title",
			UserID:     "test-user",
			Date:       date,
			CreatedAt:  date,
			UpdatedAt:  date,
		},
	}}

	result := &model.ItemsInPeriod{
		PageInfo: &model.PageInfo{
			EndCursor:   "test-user/test1",
			HasNextPage: false,
		},
		Edges:      iipe,
		TotalCount: len(iipe),
	}

	like := true
	dislike := true
	categoryID := 1

	tests := []struct {
		name   string
		param  model.InputItemsInPeriod
		result *model.ItemsInPeriod
	}{
		{
			name: "日付でアイテムを取得する",
			param: model.InputItemsInPeriod{
				After:     &after,
				First:     10,
				StartDate: date,
				EndDate:   date,
			},
			result: result,
		},
		{
			name: "日付、Like、Dislike、CategoryIDを指定してアイテムを取得する",
			param: model.InputItemsInPeriod{
				After:      &after,
				First:      10,
				StartDate:  date,
				EndDate:    date,
				Like:       &like,
				Dislike:    &dislike,
				CategoryID: &categoryID,
			},
			result: result,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r, _ := g.GetItemsInPeriod(ctx, tt.param)
			diff := cmp.Diff(r, tt.result)
			if diff != "" {
				t.Errorf("differs: (-got +want)\n%s", diff)
			} else {
				assert.Equal(t, diff, "")
			}
		})
	}
}
