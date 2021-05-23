package graph_test

import (
	mock_authToken "github.com/wheatandcat/memoir-backend/client/authToken/mocks"
	mock_timegen "github.com/wheatandcat/memoir-backend/client/timegen/mocks"
	mock_uuidgen "github.com/wheatandcat/memoir-backend/client/uuidgen/mocks"
	"github.com/wheatandcat/memoir-backend/graph"
	"github.com/wheatandcat/memoir-backend/repository"
)

func newGraph() graph.Graph {
	client := &graph.Client{
		UUID:      &mock_uuidgen.UUID{},
		Time:      &mock_timegen.Time{},
		AuthToken: &mock_authToken.AuthToken{},
	}

	app := &graph.Application{
		UserRepository: &repository.UserRepositoryInterfaceMock{},
		ItemRepository: &repository.ItemRepositoryInterfaceMock{},
	}

	g := graph.Graph{
		UserID: "test",
		Client: client,
		App:    app,
	}

	return g
}
