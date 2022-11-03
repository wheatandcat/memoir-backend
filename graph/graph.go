package graph

import (
	"go.opentelemetry.io/otel/trace"

	"github.com/wheatandcat/memoir-backend/repository"
	"github.com/wheatandcat/memoir-backend/usecase/auth"
)

type Application struct {
	UserRepository                repository.UserRepositoryInterface
	ItemRepository                repository.ItemRepositoryInterface
	InviteRepository              repository.InviteRepositoryInterface
	RelationshipRequestRepository repository.RelationshipRequestInterface
	RelationshipRepository        repository.RelationshipInterface
	PushTokenRepository           repository.PushTokenRepositoryInterface
	CommonRepository              repository.CommonRepositoryInterface
	AuthRepository                repository.AuthRepositoryInterface

	AuthUseCase auth.UseCase
	TraceClient trace.Tracer
}

// NewApplication アプリケーションを作成する
func NewApplication(tr trace.Tracer) *Application {
	return &Application{
		UserRepository:                repository.NewUserRepository(),
		ItemRepository:                repository.NewItemRepository(),
		InviteRepository:              repository.NewInviteRepository(),
		RelationshipRequestRepository: repository.NewRelationshipRequestRepository(),
		RelationshipRepository:        repository.NewRelationshipRepository(),
		PushTokenRepository:           repository.NewPushTokenRepository(),
		CommonRepository:              repository.NewCommonRepository(),
		AuthRepository:                repository.NewAuthRepository(),

		AuthUseCase: auth.New(),
		TraceClient: tr,
	}
}
