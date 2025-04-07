package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	terytUuid "github.com/mephir/teryt-golang/internal/uuid"
)

type Locality struct {
	VoivodeshipId      uint
	CountyId           uint
	MunicipalityId     uint
	MunicipalityTypeId uint
	Id                 uint
	ParentLocalityId   uint
	TypeId             uint
	CommonName         bool
	Name               string
	AsOf               time.Time
}

func (l Locality) Identifier() uint {
	return l.Id
}

func (l Locality) String() string {
	return l.Name
}

func (l Locality) TerytId() string {
	return fmt.Sprintf("%07d", l.Id)
}

func (l Locality) Uuid() uuid.UUID {
	data := terytUuid.UuidData{
		VoivodeshipId:      uint8(l.VoivodeshipId),
		CountyId:           uint8(l.CountyId),
		MunicipalityId:     uint8(l.MunicipalityId),
		MunicipalityTypeId: uint8(l.MunicipalityTypeId),
		LocalityId:         func() *uint32 { id := uint32(l.Id); return &id }(),
		AsOf:               l.AsOf,
		Name:               l.String(),
	}

	id, err := data.Encode()
	if err != nil {
		panic(fmt.Sprintf("failed to encode UUID: %v", err))
	}

	return id
}
