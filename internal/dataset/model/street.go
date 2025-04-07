package model

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	terytUuid "github.com/mephir/teryt-golang/internal/uuid"
)

type Street struct {
	VoivodeshipId      uint
	CountyId           uint
	MunicipalityId     uint
	MunicipalityTypeId uint
	LocalityId         uint
	Id                 uint
	SortableName       string
	NamePrefix         string
	Type               string
	AsOf               time.Time
}

func (s Street) Identifier() uint {
	return s.Id + s.LocalityId*100000
}

func (s Street) Uuid() uuid.UUID {
	data := terytUuid.UuidData{
		VoivodeshipId:      uint8(s.VoivodeshipId),
		CountyId:           uint8(s.CountyId),
		MunicipalityId:     uint8(s.MunicipalityId),
		MunicipalityTypeId: uint8(s.MunicipalityTypeId),
		LocalityId:         func() *uint32 { id := uint32(s.LocalityId); return &id }(),
		StreetId:           func() *uint32 { id := uint32(s.Id); return &id }(),
		AsOf:               s.AsOf,
		Name:               s.String(),
	}

	id, err := data.Encode()
	if err != nil {
		panic(fmt.Sprintf("failed to encode UUID: %v", err))
	}
	return id
}

func (s Street) String() string {
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
	return fmt.Sprintf("%07d%05d", s.LocalityId, s.Id)
}
