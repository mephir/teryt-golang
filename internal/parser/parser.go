package parser

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"reflect"

	"github.com/mephir/teryt-golang/internal/dataset"
	"github.com/mephir/teryt-golang/internal/parser/xmlstructs"
)

type Parser interface {
	Close()
	Fetch() (any, error)
	FetchAll() (any, error)
}

type XmlParser struct {
	decoder    *xml.Decoder
	fileHandle *os.File
	structType reflect.Type
	Dataset    dataset.Dataset
}

func Open(path string) (*XmlParser, error) {
	dataset, err := dataset.Determine(path)
	if err != nil {
		return nil, fmt.Errorf("could not determine dataset: %w", err)
	}

	fh, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	var structType reflect.Type
	switch dataset.Name {
	case "SIMC":
		if dataset.Variant == "S" {
			structType = reflect.TypeOf(xmlstructs.SimcS{})
		} else {
			structType = reflect.TypeOf(xmlstructs.Simc{})
		}
	case "ULIC":
		structType = reflect.TypeOf(xmlstructs.Ulic{})
	case "TERC":
		structType = reflect.TypeOf(xmlstructs.Terc{})
	case "WMRODZ":
		structType = reflect.TypeOf(xmlstructs.Wmrodz{})
	}

	return &XmlParser{
		decoder:    nil,
		fileHandle: fh,
		structType: structType,
		Dataset:    *dataset,
	}, nil
}

func (p *XmlParser) Close() error {
	return p.fileHandle.Close()
}

func (p *XmlParser) Fetch() (any, error) {
	decoder := p.getDecoder()

	for {
		token, err := decoder.Token()
		if err != nil {
			return nil, err
		}

		switch t := token.(type) {
		case xml.StartElement:
			if t.Name.Local == "row" {
				item := reflect.New(p.structType).Interface()
				if err := decoder.DecodeElement(&item, &t); err != nil {
					return nil, fmt.Errorf("could not decode element: %w", err)
				}

				return item, nil
			}
		}
	}
}

func (p *XmlParser) FetchAll() ([]any, error) {
	var result []any
	decoder := p.newDecoder()

	for {
		token, err := decoder.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		switch t := token.(type) {
		case xml.StartElement:
			if t.Name.Local == "row" {
				item := reflect.New(p.structType).Interface()
				if err := decoder.DecodeElement(&item, &t); err != nil {
					return nil, fmt.Errorf("could not decode element: %w", err)
				}

				result = append(result, item)
			}
		}
	}

	return result, nil
}

func (p *XmlParser) getDecoder() *xml.Decoder {
	if p.decoder == nil {
		p.decoder = p.newDecoder()
	}

	return p.decoder
}

func (p *XmlParser) newDecoder() *xml.Decoder {
	return xml.NewDecoder(p.fileHandle)
}
