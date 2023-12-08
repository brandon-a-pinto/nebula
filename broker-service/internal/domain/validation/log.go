package validation

import (
	"errors"

	"github.com/brandon-a-pinto/nebula/broker-service/internal/domain/dto"
)

func CreateLogValidation(input dto.CreateLogInput) error {
	if input.Type == "" {
		return errors.New("type is required")
	}

	if input.Type != "INFO" && input.Type != "WARN" && input.Type != "ERROR" {
		return errors.New("type is invalid")
	}

	if input.Msg == "" {
		return errors.New("msg is required")
	}

	return nil
}
