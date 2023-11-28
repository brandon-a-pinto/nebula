package contract

import (
	"context"

	"github.com/brandon-a-pinto/nebula/logger-service/internal/domain/entity"
)

type ILogRepository interface {
	Save(ctx context.Context, log *entity.Log) error
}
