package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Municipality struct { // Gmina
	Id            uint
	Name          string
	UnitType      string
	AsOf          time.Time
	Type          MunicipalityType
	CountyId      uint
	County        *County
	VoivodeshipId uint
	Localities    []*Locality
}

func (m Municipality) Identifier() uint {
	return m.Type.Id + m.Id*10 + m.CountyId*1000 + m.VoivodeshipId*100000
}

func (m Municipality) ToString() string {
	return fmt.Sprintf("%s %s", m.UnitType, m.Name)
}

func (m Municipality) Uuid() uuid.UUID {
	return uuid.NewSHA1(m.County.Uuid(), []byte(m.ToString()))
}

func (m Municipality) TerytId() string {
	return fmt.Sprintf("%02d%02d%02d%01d", m.VoivodeshipId, m.CountyId, m.Id, m.Type.Id)
}

func (m Municipality) GetCountyIdentifier() uint {
	if m.County == nil {
		return m.VoivodeshipId*100 + m.CountyId
	}

	return m.County.Identifier()
}
