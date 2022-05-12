package graph_test

import (
	"errors"
	"testing"

	mock_authToken "github.com/wheatandcat/memoir-backend/client/authToken/mocks"
	mock_task "github.com/wheatandcat/memoir-backend/client/task/mocks"
	mock_timegen "github.com/wheatandcat/memoir-backend/client/timegen/mocks"
	mock_uuidgen "github.com/wheatandcat/memoir-backend/client/uuidgen/mocks"
	"github.com/wheatandcat/memoir-backend/graph"
	ce "github.com/wheatandcat/memoir-backend/usecase/custom_error"
	"gopkg.in/go-playground/assert.v1"

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
		AuthRepository:                &moq_repository.AuthRepositoryInterfaceMock{},

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

func checkErrorCode(t *testing.T, err1 error, err2 error) {
	code1 := ce.CodeDefault
	code2 := ce.CodeDefault

	var re1 ce.RequestError
	if errors.As(err1, &re1) {
		code1 = re1.Code
	}
	var re2 ce.RequestError
	if errors.As(err2, &re2) {
		code2 = re2.Code
	}

	assert.Equal(t, code1, code2)

}
