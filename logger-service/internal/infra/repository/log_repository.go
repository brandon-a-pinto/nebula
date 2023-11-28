package repository

import (
	"context"

	"github.com/brandon-a-pinto/nebula/logger-service/internal/domain/entity"
	"github.com/brandon-a-pinto/nebula/logger-service/pkg/infra/database"
)

type LogRepository struct{}

func NewLogRepository() *LogRepository {
	return &LogRepository{}
}

func (r *LogRepository) Save(ctx context.Context, log *entity.Log) error {
	_, err := database.MI.Collection("logs").InsertOne(ctx, log)
	if err != nil {
		return err
	}

	return nil
}
