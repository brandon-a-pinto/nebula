package entity

import (
	"time"

	"github.com/brandon-a-pinto/nebula/user-service/internal/domain/dto"
	"github.com/brandon-a-pinto/nebula/user-service/internal/domain/validation"
	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID
	Email       string
	Username    string
	DisplayName string
	Password    string
	IsAdmin     bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewUser(input dto.CreateUserInput) (*User, error) {
	user := &User{
		Email:       input.Email,
		Username:    input.Username,
		DisplayName: input.DisplayName,
		Password:    input.Password,
		IsAdmin:     false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := validation.CreateUserValidation(input); err != nil {
		return nil, err
	}

	return user, nil
}
