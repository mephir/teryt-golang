package parser

import (
	"encoding/xml"
	"fmt"
	"os"
)

type Parser interface {
	Clean() error
}

type XmlParser struct {
	decoder    *xml.Decoder
	fileHandle *os.File
}

func CreateParser(path string) (*XmlParser, error) {
	fh, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}

	decoder := xml.NewDecoder(fh)

	return &XmlParser{decoder: decoder, fileHandle: fh}, nil
}

func (p *XmlParser) Clean() error {
	return p.fileHandle.Close()
}
