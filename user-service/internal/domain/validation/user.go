package validation

import (
	"errors"

	"github.com/brandon-a-pinto/nebula/user-service/internal/domain/dto"
)

func CreateUserValidation(input dto.CreateUserInput) error {
	if input.Email == "" {
		return errors.New("email is required")
	}

	if input.Username == "" {
		return errors.New("username is required")
	}

	if input.DisplayName == "" {
		return errors.New("display_name is required")
	}

	if input.Password == "" {
		return errors.New("password is required")
	}
	if len(input.Password) < 8 {
		return errors.New("password must be at least 8 characters")
	}

	return nil
}
