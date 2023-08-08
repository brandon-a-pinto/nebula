package models

import "time"

type Request struct {
	Name string `json:"name"`
	Data any    `json:"data"`
}

type Log struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	Data      any       `bson:"data" json:"data"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}
