package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Locality struct {
	Id               uint
	Type             *LocalityType
	TypeId           uint
	CommonName       bool
	Name             string
	ParentLocalityId uint
	MunicipalityId   uint //Identifier of the municipality
	Municipality     *Municipality
	AsOf             time.Time
}

func (l Locality) Identifier() uint {
	return l.Type.Id
}

func (l Locality) ToString() string {
	return l.Name
}

func (l Locality) TerytId() string {
	return fmt.Sprintf("%07d", l.Id)
}

func (l Locality) Uuid() uuid.UUID {
	return uuid.NewSHA1(l.Municipality.Uuid(), []byte(l.ToString()))
}
