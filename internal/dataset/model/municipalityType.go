package model

import (
	"encoding/xml"
	"fmt"

	"github.com/google/uuid"
	"github.com/mephir/teryt-golang/internal/collection"
)

type MunicipalityType struct {
	Id   uint
	Name string
}

func (r MunicipalityType) Identifier() uint {
	return r.Id
}

func (r MunicipalityType) Uuid() uuid.UUID {
	return uuid.NewSHA1(uuid.Nil, []byte(r.Name))
}

func (r MunicipalityType) ToString() string {
	return r.Name
}

func (r MunicipalityType) TerytId() string {
	return fmt.Sprintf("%02d", r.Id)
}

func (r *MunicipalityType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var typeId uint
	if err := d.DecodeElement(&typeId, &start); err != nil {
		return err
	}

	rodz := (*GetMunicipalityTypesCollection()).Get(typeId)
	if rodz == nil {
		return fmt.Errorf("unknown municipality type %d", typeId)
	}

	*r = *rodz

	return nil
}

func GetMunicipalityTypesCollection() *collection.Collection[MunicipalityType] {
	typesCollection := collection.NewCollection[MunicipalityType]()
	types := []*MunicipalityType{
		{0, "brak lub nieznany"},
		{1, "gmina miejska"},
		{2, "gmina wiejska"},
		{3, "gmina miejsko-wiejska"},
		{4, "miasto w gminie miejsko-wiejskiej"},
		{5, "obszar wiejski w gminie miejsko-wiejskiej"},
		{8, "dzielnica m. st. Warszawy"},
		{9, "delegatura"},
	}

	typesCollection.Add(types...)

	return typesCollection
}
