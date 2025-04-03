package model

import (
	"encoding/xml"
	"fmt"

	"github.com/google/uuid"
)

type Rodz struct {
	Id   uint
	Name string
}

var rodzCollection = []Rodz{
	{0, "brak lub nieznany"},
	{1, "gmina miejska"},
	{2, "gmina wiejska"},
	{3, "gmina miejsko-wiejska"},
	{4, "miasto w gminie miejsko-wiejskiej"},
	{5, "obszar wiejski w gminie miejsko-wiejskiej"},
	{8, "dzielnica m. st. Warszawy"},
	{9, "delegatura"},
}

func (r Rodz) Identifier() uint {
	return r.Id
}

func (r Rodz) Uuid() uuid.UUID {
	return uuid.NewSHA1(uuid.Nil, []byte(r.Name))
}

func (r Rodz) ToString() string {
	return r.Name
}

func RodzGet(id uint) (Rodz, error) {
	for _, rodz := range rodzCollection {
		if rodz.Identifier() == id {
			return rodz, nil
		}
	}

	return rodzCollection[0], fmt.Errorf("Rodz with id %d not found", id)
}

func (r *Rodz) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var rodzId uint
	if err := d.DecodeElement(&rodzId, &start); err != nil {
		return err
	}

	rodz, err := RodzGet(rodzId)
	if err != nil {
		return err
	}

	*r = rodz

	return nil
}
