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
	Type          Rodz
	CountyId      uint
	County        *County
	VoivodeshipId uint
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
