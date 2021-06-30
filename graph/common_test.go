package graph_test

import (
	mock_authToken "github.com/wheatandcat/memoir-backend/client/authToken/mocks"
	mock_task "github.com/wheatandcat/memoir-backend/client/task/mocks"
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
		Task:      mock_task.NewNotificationTask(),
	}

	app := &graph.Application{
		UserRepository:                &repository.UserRepositoryInterfaceMock{},
		ItemRepository:                &repository.ItemRepositoryInterfaceMock{},
		InviteRepository:              &repository.InviteRepositoryInterfaceMock{},
		RelationshipRequestRepository: &repository.RelationshipRequestInterfaceMock{},
		RelationshipRepository:        &repository.RelationshipInterfaceMock{},
		PushTokenRepository:           &repository.PushTokenRepositoryInterfaceMock{},
		CommonRepository:              &repository.CommonRepositoryInterfaceMock{},
	}

	g := graph.Graph{
		UserID: "test",
		Client: client,
		App:    app,
	}

	return g
}
