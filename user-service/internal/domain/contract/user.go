package contract

import (
	"context"

	"github.com/brandon-a-pinto/nebula/user-service/internal/domain/entity"
)

type IUserRepository interface {
	Save(ctx context.Context, user *entity.User) error
}
