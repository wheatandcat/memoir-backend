package graph_test

import (
	"errors"

	mock_authToken "github.com/wheatandcat/memoir-backend/client/authToken/mocks"
	mock_task "github.com/wheatandcat/memoir-backend/client/task/mocks"
	mock_timegen "github.com/wheatandcat/memoir-backend/client/timegen/mocks"
	mock_uuidgen "github.com/wheatandcat/memoir-backend/client/uuidgen/mocks"
	"github.com/wheatandcat/memoir-backend/graph"
	ce "github.com/wheatandcat/memoir-backend/usecase/custom_error"

	moq_repository "github.com/wheatandcat/memoir-backend/repository/moq"
	moq_usecase_auth "github.com/wheatandcat/memoir-backend/usecase/auth/moq"
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

		AuthUseCase: &moq_usecase_auth.UseCaseMock{},
	}

	g := graph.Graph{
		UserID: "test",
		Client: client,
		App:    app,
	}

	return g
}

func getErrorCode(err error) string {
	errorCode := ce.CodeDefault
	var re ce.RequestError
	if errors.As(err, &re) {
		errorCode = re.Code
	}

	return errorCode
}
