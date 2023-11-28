package factory

import (
	"github.com/brandon-a-pinto/nebula/logger-service/internal/application/usecase"
	"github.com/brandon-a-pinto/nebula/logger-service/internal/infra/repository"
)

func CreateLogFactory() *usecase.CreateLogUsecase {
	return usecase.NewCreateLogUsecase(repository.NewLogRepository())
}
