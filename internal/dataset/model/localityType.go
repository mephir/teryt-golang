package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type LocalityType struct {
	Id        uint
	Name      string
	UpdatedAt time.Time
}

func (w LocalityType) Identifier() uint {
	return w.Id
}

func (w LocalityType) Uuid() uuid.UUID {
	return uuid.NewSHA1(uuid.Nil, []byte(w.Name))
}

func (w LocalityType) ToString() string {
	return w.Name
}

func (w LocalityType) TerytId() string {
	return fmt.Sprintf("%02d", w.Id)
}
