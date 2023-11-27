package factory

import (
	"github.com/brandon-a-pinto/nebula/broker-service/configs"
	"github.com/brandon-a-pinto/nebula/broker-service/internal/application/usecase"
)

func CreateUserFactory() *usecase.CreateUserUsecase {
	config := configs.LoadConfig()
	return usecase.NewCreateUserUsecase(config.UserGRPCPort)
}
