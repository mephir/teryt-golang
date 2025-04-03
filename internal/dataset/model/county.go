package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type County struct { // Gmina
	Id             uint
	Name           string
	UnitType       string
	AsOf           time.Time
	Voivodeship    *Voivodeship
	VoivodeshipId  uint
	Municipalities []*Municipality
}

func (c County) Identifier() uint {
	return c.VoivodeshipId*100 + c.Id
}

func (c County) Uuid() uuid.UUID {
	return uuid.NewSHA1(c.Voivodeship.Uuid(), []byte(c.ToString()))
}

func (c County) ToString() string {
	return fmt.Sprintf("%s %s", c.UnitType, c.Name)
}
