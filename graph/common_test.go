package graph_test

import (
	"github.com/wheatandcat/memoir-backend/client/timegen"
	"github.com/wheatandcat/memoir-backend/client/uuidgen"
	"github.com/wheatandcat/memoir-backend/graph"
	"github.com/wheatandcat/memoir-backend/repository"
)

func newGraph() graph.Graph {
	client := &graph.Client{
		UUID: &uuidgen.UUID{},
		Time: &timegen.Time{},
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
