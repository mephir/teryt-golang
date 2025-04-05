package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Street struct {
	Id           uint
	LocalityId   uint
	SortableName string
	NamePrefix   string
	Type         string
	AsOf         time.Time
}

func (s Street) Identifier() uint {
	return s.Id + s.LocalityId*100000
}

func (s Street) Uuid() uuid.UUID {
	return uuid.New()
}

func (s Street) ToString() string {
	var parts []string

	if s.Type != "" {
		parts = append(parts, s.Type)
	}
	if s.NamePrefix != "" {
		parts = append(parts, s.NamePrefix)
	}
	parts = append(parts, s.SortableName)

	return strings.Join(parts, " ")
}

func (s Street) TerytId() string {
	return fmt.Sprintf("%05d", s.Id)
}
