package datastruct

import (
	"encoding/xml"
	"strconv"
	"strings"

	"github.com/mephir/teryt-golang/internal/dataset/model"
)

type Terc struct {
	Woj      uint                    `xml:"WOJ"`       // Województwo, Voivodeship
	Pow      *uint                   `xml:"POW"`       // Powiat, County
	Gmi      *uint                   `xml:"GMI"`       // Gmina, Municipality
	Rodz     *model.MunicipalityType `xml:"RODZ"`      // Rodzaj jednostki, Unit type
	Name     string                  `xml:"NAZWA"`     // Nazwa, Name
	UnitType string                  `xml:"NAZWA_DOD"` // Określenie jednostki, Unit determination
	AsOf     AsOf                    `xml:"STAN_NA"`   // Data aktualizacji, Update date
}

func (t *Terc) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type TAlias Terc
	aux := &struct {
		Pow *string `xml:"POW"`
		Gmi *string `xml:"GMI"`

		*TAlias
	}{
		TAlias: (*TAlias)(t),
	}

	if err := d.DecodeElement(aux, &start); err != nil {
		return err
	}

	if aux.Pow != nil && *aux.Pow != "" {
		val, err := strconv.ParseUint(*aux.Pow, 10, 64)
		if err != nil {
			return err
		}
		t.Pow = new(uint)
		*t.Pow = uint(val)
	} else {
		t.Pow = nil
	}

	if aux.Gmi != nil && *aux.Gmi != "" {
		val, err := strconv.ParseUint(*aux.Gmi, 10, 64)
		if err != nil {
			return err
		}
		t.Gmi = new(uint)
		*t.Gmi = uint(val)
	} else {
		t.Gmi = nil
	}

	return nil
}

func (t Terc) ToModel() (model.Model, error) {
	if t.IsVoivodeship() {
		return &model.Voivodeship{
			Id:       t.Woj,
			Name:     strings.ToLower(t.Name),
			UnitType: t.UnitType,
			AsOf:     t.AsOf.Time,
		}, nil
	}

	if t.IsCounty() {
		return &model.County{
			VoivodeshipId: t.Woj,
			Id:            *t.Pow,
			Name:          t.Name,
			UnitType:      t.UnitType,
			AsOf:          t.AsOf.Time,
		}, nil
	}

	return &model.Municipality{
		VoivodeshipId: t.Woj,
		CountyId:      *t.Pow,
		Id:            *t.Gmi,
		Name:          t.Name,
		UnitType:      t.UnitType,
		AsOf:          t.AsOf.Time,
		Type:          *t.Rodz,
	}, nil
}

func (t Terc) IsVoivodeship() bool {
	return t.Pow == nil && t.Gmi == nil
}

func (t Terc) IsCounty() bool {
	return t.Pow != nil && t.Gmi == nil
}

func (t Terc) IsMunicipality() bool {
	return t.Pow != nil && t.Gmi != nil
}
