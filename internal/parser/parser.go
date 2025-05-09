package parser

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
	"reflect"

	"github.com/mephir/teryt-golang/internal/dataset"
	"github.com/mephir/teryt-golang/internal/dataset/datastruct"
)

type Parser interface {
	Close() error
	Fetch() (datastruct.Datastruct, error)
	FetchAll() ([]datastruct.Datastruct, error)
	GetDataset() *dataset.Dataset
	GetStructType() reflect.Type
}

type XmlParser struct {
	Dataset dataset.Dataset

	decoder    *xml.Decoder
	fileHandle *os.File
	structType reflect.Type
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
			structType = reflect.TypeOf(datastruct.SimcS{})
		} else {
			structType = reflect.TypeOf(datastruct.Simc{})
		}
	case "ULIC":
		structType = reflect.TypeOf(datastruct.Ulic{})
	case "TERC":
		structType = reflect.TypeOf(datastruct.Terc{})
	case "WMRODZ":
		structType = reflect.TypeOf(datastruct.Wmrodz{})
	}

	return &XmlParser{
		Dataset:    *dataset,
		decoder:    nil,
		fileHandle: fh,
		structType: structType,
	}, nil
}

func (p *XmlParser) Close() error {
	return p.fileHandle.Close()
}

func (p *XmlParser) Fetch() (datastruct.Datastruct, error) {
	decoder := p.getDecoder()

	for {
		token, err := decoder.Token()
		if err != nil {
			return nil, err
		}

		switch t := token.(type) {
		case xml.StartElement:
			if t.Name.Local == "row" {
				item := reflect.New(p.structType).Interface().(datastruct.Datastruct)
				if err := decoder.DecodeElement(&item, &t); err != nil {
					return nil, err
				}

				return item, nil
			}
		}
	}
}

func (p *XmlParser) FetchAll() ([]datastruct.Datastruct, error) {
	var result []datastruct.Datastruct
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
				item := reflect.New(p.structType).Interface().(datastruct.Datastruct)
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

func (p *XmlParser) GetDataset() *dataset.Dataset {
	return &p.Dataset
}

func (p *XmlParser) GetStructType() reflect.Type {
	return p.structType
}
