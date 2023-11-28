package usecase

import (
	"context"

	"github.com/brandon-a-pinto/nebula/logger-service/internal/domain/contract"
	"github.com/brandon-a-pinto/nebula/logger-service/internal/domain/dto"
	"github.com/brandon-a-pinto/nebula/logger-service/internal/domain/entity"
)

type CreateLogUsecase struct {
	LogRepository contract.ILogRepository
}

func NewCreateLogUsecase(logRepository contract.ILogRepository) *CreateLogUsecase {
	return &CreateLogUsecase{
		LogRepository: logRepository,
	}
}

func (u *CreateLogUsecase) Exec(ctx context.Context, input dto.CreateLogInput) error {
	log := entity.NewLog(input)

	err := u.LogRepository.Save(ctx, log)
	if err != nil {
		return err
	}

	return nil
}
