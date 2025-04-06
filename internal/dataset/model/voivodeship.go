package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	terytUuid "github.com/mephir/teryt-golang/internal/uuid"
)

type Voivodeship struct { // Wojewodztwo
	Id       uint
	Name     string
	UnitType string
	AsOf     time.Time
}

func (v Voivodeship) Identifier() uint {
	return v.Id
}

func (v Voivodeship) Uuid() uuid.UUID {
	data := terytUuid.UuidData{
		VoivodeshipId:      uint8(v.Id),
		CountyId:           0,
		MunicipalityId:     0,
		MunicipalityTypeId: 0,
		AsOf:               v.AsOf,
		Name:               v.ToString(),
	}

	id, err := data.Encode()
	if err != nil {
		panic(fmt.Sprintf("failed to encode UUID: %v", err))
	}

	return id
}

func (v Voivodeship) ToString() string {
	return fmt.Sprintf("%s %s", v.UnitType, v.Name)
}

func (v Voivodeship) TerytId() string {
	return fmt.Sprintf("%02d", v.Id)
}
