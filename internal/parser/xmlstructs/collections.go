package xmlstructs

import "encoding/xml"

type Rodz struct {
	Id   uint8
	Name string
}

var RodzCollection = map[uint8]Rodz{
	1: {1, "gmina miejska"},
	2: {2, "gmina wiejska"},
	3: {3, "gmina miejsko-wiejska"},
	4: {4, "miasto w gminie miejsko-wiejskiej"},
	5: {5, "obszar wiejski w gminie miejsko-wiejskiej"},
	8: {8, "dzielnica m. st. Warszawy"},
	9: {9, "delegatura"},
	0: {0, "brak lub nieznany"},
}

func (r *Rodz) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var rodzId uint8
	if err := d.DecodeElement(&rodzId, &start); err != nil {
		return err
	}

	*r = RodzCollection[rodzId]

	return nil
}
