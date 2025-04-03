package model

import (
	"github.com/google/uuid"
)

type Model interface {
	Identifier() uint
	Uuid() uuid.UUID
	ToString() string
}
