package entity

import (
	"github.com/brandon-a-pinto/nebula/broker-service/internal/domain/dto"
	"github.com/brandon-a-pinto/nebula/broker-service/internal/domain/validation"
)

type Log struct {
	Msg  string
	Type string
}

func NewLog(input dto.CreateLogInput) (*Log, error) {
	log := &Log{
		Type: input.Type,
		Msg:  input.Msg,
	}

	err := validation.CreateLogValidation(input)
	if err != nil {
		return nil, err
	}

	return log, nil
}
