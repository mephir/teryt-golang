package datastruct

import (
	"encoding/xml"
	"time"
)

type AsOf struct {
	time.Time
}

func (s *AsOf) UnmarshalXML(decoder *xml.Decoder, start xml.StartElement) error {
	var dateString string
	if err := decoder.DecodeElement(&dateString, &start); err != nil {
		return err
	}

	date, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return err
	}

	s.Time = date

	return nil
}
