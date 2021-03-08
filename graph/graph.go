package graph

import "github.com/wheatandcat/memoir-backend/repository"

// Application is app interface
type Application struct {
	UserRepository repository.UserRepositoryInterface
}

// NewApplication アプリケーションを作成する
func NewApplication() *Application {
	return &Application{
		UserRepository: repository.NewUserRepository(),
	}
}
