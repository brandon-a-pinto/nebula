package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id,omitempty"`
	Email       string    `json:"email" validate:"required"`
	Username    string    `json:"username" validate:"required"`
	DisplayName string    `json:"displayName" validate:"required"`
	Password    string    `json:"password,omitempty" validate:"required"`
	IsAdmin     bool      `json:"isAdmin" validate:"required"`
	CreatedAt   time.Time `json:"createdAt" validate:"required"`
	UpdatedAt   time.Time `json:"updatedAt" validate:"required"`
}

type CreateUserParams struct {
	Email                string `json:"email"`
	Username             string `json:"username"`
	DisplayName          string `json:"displayName"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"passwordConfirmation"`
}
