package graph_test

import (
	mock_authToken "github.com/wheatandcat/memoir-backend/client/authToken/mocks"
	mock_task "github.com/wheatandcat/memoir-backend/client/task/mocks"
	mock_timegen "github.com/wheatandcat/memoir-backend/client/timegen/mocks"
	mock_uuidgen "github.com/wheatandcat/memoir-backend/client/uuidgen/mocks"
	"github.com/wheatandcat/memoir-backend/graph"

	moq_repository "github.com/wheatandcat/memoir-backend/repository/moq"
)

func newGraph() graph.Graph {
	client := &graph.Client{
		UUID:      &mock_uuidgen.UUID{},
		Time:      &mock_timegen.Time{},
		AuthToken: &mock_authToken.AuthToken{},
		Task:      mock_task.NewNotificationTask(),
	}

	app := &graph.Application{
		UserRepository:                &moq_repository.UserRepositoryInterfaceMock{},
		ItemRepository:                &moq_repository.ItemRepositoryInterfaceMock{},
		InviteRepository:              &moq_repository.InviteRepositoryInterfaceMock{},
		RelationshipRequestRepository: &moq_repository.RelationshipRequestInterfaceMock{},
		RelationshipRepository:        &moq_repository.RelationshipInterfaceMock{},
		PushTokenRepository:           &moq_repository.PushTokenRepositoryInterfaceMock{},
		CommonRepository:              &moq_repository.CommonRepositoryInterfaceMock{},
	}

	g := graph.Graph{
		UserID: "test",
		Client: client,
		App:    app,
	}

	return g
}
