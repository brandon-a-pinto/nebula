package repository

import (
	"context"

	"github.com/brandon-a-pinto/nebula/user-service/internal/domain/entity"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) Save(ctx context.Context, user *entity.User) error {
	return nil
}
