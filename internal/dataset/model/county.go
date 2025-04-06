package model

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	terytUuid "github.com/mephir/teryt-golang/internal/uuid"
)

type County struct {
	VoivodeshipId uint
	Id            uint
	Name          string
	UnitType      string
	AsOf          time.Time
}

func (c County) Identifier() uint {
	return c.VoivodeshipId*100 + c.Id
}

func (c County) Uuid() uuid.UUID {
	data := terytUuid.UuidData{
		VoivodeshipId:      uint8(c.VoivodeshipId),
		CountyId:           uint8(c.Id),
		MunicipalityId:     0,
		MunicipalityTypeId: 0,
		AsOf:               c.AsOf,
		Name:               c.ToString(),
	}

	id, err := data.Encode()
	if err != nil {
		panic(fmt.Sprintf("failed to encode UUID: %v", err))
	}

	return id
}

func (c County) ToString() string {
	return fmt.Sprintf("%s %s", c.UnitType, c.Name)
}

func (c County) TerytId() string {
	return fmt.Sprintf("%02d%02d", c.VoivodeshipId, c.Id)
}
