package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Voivodeship struct { // Wojewodztwo
	Id       uint
	Name     string
	UnitType string
	AsOf     time.Time
	Counties []*County
}

func (v Voivodeship) Identifier() uint {
	return v.Id
}

func (v Voivodeship) Uuid() uuid.UUID {
	return uuid.NewSHA1(uuid.Nil, []byte(v.ToString()))
}

func (v Voivodeship) ToString() string {
	return fmt.Sprintf("%s %s", v.UnitType, v.Name)
}

func (v Voivodeship) TerytId() string {
	return fmt.Sprintf("%02d", v.Id)
}
