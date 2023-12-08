package factory

import (
	"github.com/brandon-a-pinto/nebula/broker-service/internal/application/usecase"
	"github.com/brandon-a-pinto/nebula/broker-service/internal/infra/event"
)

func CreateLogFactory() *usecase.CreateLogUsecase {
	return usecase.NewCreateLogUsecase(event.NewLogEvent())
}
