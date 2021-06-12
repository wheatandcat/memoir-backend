package graph

import "github.com/wheatandcat/memoir-backend/repository"

// Application is app interface
type Application struct {
	UserRepository                repository.UserRepositoryInterface
	ItemRepository                repository.ItemRepositoryInterface
	InviteRepository              repository.InviteRepositoryInterface
	RelationshipRequestRepository repository.RelationshipRequestInterface
	RelationshipRepository        repository.RelationshipInterface
	CommonRepository              repository.CommonRepositoryInterface
}

// NewApplication アプリケーションを作成する
func NewApplication() *Application {
	return &Application{
		UserRepository:                repository.NewUserRepository(),
		ItemRepository:                repository.NewItemRepository(),
		InviteRepository:              repository.NewInviteRepository(),
		RelationshipRequestRepository: repository.NewRelationshipRequestRepository(),
		RelationshipRepository:        repository.NewRelationshipRepository(),
		CommonRepository:              repository.NewCommonRepository(),
	}
}
