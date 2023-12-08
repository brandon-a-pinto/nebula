package usecase

import (
	"context"
	"encoding/json"

	"github.com/brandon-a-pinto/nebula/broker-service/internal/domain/contract"
	"github.com/brandon-a-pinto/nebula/broker-service/internal/domain/dto"
)

type CreateLogUsecase struct {
	LogEvent contract.ILogEvent
}

func NewCreateLogUsecase(logEvent contract.ILogEvent) *CreateLogUsecase {
	return &CreateLogUsecase{
		LogEvent: logEvent,
	}
}

func (u *CreateLogUsecase) Exec(ctx context.Context, input dto.CreateLogInput) error {
	jsonValue, err := json.Marshal(input)
	if err != nil {
		return err
	}

	err = u.LogEvent.Create(string(jsonValue), "log."+input.Type)
	if err != nil {
		return err
	}

	return nil
}
