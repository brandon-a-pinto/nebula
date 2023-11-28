package entity

import (
	"time"

	"github.com/brandon-a-pinto/nebula/logger-service/internal/domain/dto"
)

type Log struct {
	Type      string    `bson:"type"`
	Msg       string    `bson:"msg"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}

func NewLog(input dto.CreateLogInput) *Log {
	return &Log{
		Msg:       input.Msg,
		Type:      input.Type,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
