package factory

import (
	"github.com/brandon-a-pinto/nebula/user-service/internal/application/usecase"
	"github.com/brandon-a-pinto/nebula/user-service/internal/infra/cryptography"
	"github.com/brandon-a-pinto/nebula/user-service/internal/infra/repository"
)

func CreateUserFactory() *usecase.CreateUserUsecase {
	return &usecase.CreateUserUsecase{
		UserRepository: repository.NewUserRepository(),
		BcryptAdapter:  cryptography.NewBcryptAdapter(),
	}
}
