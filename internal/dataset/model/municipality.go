package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	terytUuid "github.com/mephir/teryt-golang/internal/uuid"
)

type Municipality struct {
	VoivodeshipId uint
	CountyId      uint
	Id            uint
	Name          string
	UnitType      string
	AsOf          time.Time
	Type          MunicipalityType
}

func (m Municipality) Identifier() uint {
	return m.Type.Id + m.Id*10 + m.CountyId*1000 + m.VoivodeshipId*100000
}

func (m Municipality) ToString() string {
	return fmt.Sprintf("%s %s", m.UnitType, m.Name)
}

func (m Municipality) Uuid() uuid.UUID {
	data := terytUuid.UuidData{
		VoivodeshipId:      uint8(m.VoivodeshipId),
		CountyId:           uint8(m.CountyId),
		MunicipalityId:     uint8(m.Id),
		MunicipalityTypeId: uint8(m.Type.Id),
		AsOf:               m.AsOf,
		Name:               m.ToString(),
	}

	id, err := data.Encode()
	if err != nil {
		panic(fmt.Sprintf("failed to encode UUID: %v", err))
	}

	return id
}

func (m Municipality) TerytId() string {
	return fmt.Sprintf("%02d%02d%02d%01d", m.VoivodeshipId, m.CountyId, m.Id, m.Type.Id)
}

func (m Municipality) GetCountyIdentifier() uint {
	return m.VoivodeshipId*100 + m.CountyId
}
