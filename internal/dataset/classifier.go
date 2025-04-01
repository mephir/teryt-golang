package dataset

import (
	"encoding/xml"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Determine(path string) (*Dataset, error) {
	dataset, _ := DetermineByFilename(filepath.Base(path))
	if dataset == nil {
		dataset = &Dataset{}
	}

	fh, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %w", err)
	}
	defer fh.Close()

	decoder := xml.NewDecoder(fh)
	dsByContent, err := DetermineByContent(decoder)
	if err != nil {
		return nil, fmt.Errorf("could not determine dataset by content: %w", err)
	}

	if dataset.Name != dsByContent.Name {
		dataset.Name = dsByContent.Name
	}

	if !dataset.Date.Equal(dsByContent.Date) {
		dataset.Date = dsByContent.Date
	}

	if dsByContent.Variant != "" && dataset.Variant != dsByContent.Variant {
		dataset.Variant = dsByContent.Variant
	}

	return dataset, nil
}

func DetermineByFilename(filename string) (*Dataset, error) {
	var date time.Time

	filename = filename[:len(filename)-len(filepath.Ext(filename))] // remove extension
	parts := strings.Split(filename, "_")

	if len(parts) < 2 || len(parts) > 3 {
		return nil, fmt.Errorf("unexpected filename format: %s", filename)
	}

	date, err := time.Parse("2006-01-02", parts[len(parts)-1])
	if err != nil {
		return nil, fmt.Errorf("could not parse date: %w", err)
	}

	if len(parts) == 2 && strings.ToUpper(parts[0]) != "WMRODZ" {
		return nil, fmt.Errorf("unexpected filename format: %s", filename)
	}

	dataset := &Dataset{}
	dataset.Name = strings.ToUpper(parts[0])
	dataset.Date = date
	if dataset.Name != "WMRODZ" {
		dataset.Variant = strings.ToUpper(string([]rune(parts[1])[0]))
	}

	err = dataset.Validate()
	if err != nil {
		return nil, fmt.Errorf("could not validate dataset: %w", err)
	}

	return dataset, nil
}

func DetermineByContent(decoder *xml.Decoder) (*Dataset, error) {
	var name string
	var date time.Time
	var variant string

	for {
		token, err := decoder.Token()
		if err != nil {
			if name == "" {
				return nil, fmt.Errorf("unexpected end of file, could not determine dataset name")
			}
			break
		}

		switch t := token.(type) {
		case xml.StartElement:
			if t.Name.Local == "catalog" {
				for _, attr := range t.Attr {
					if attr.Name.Local == "name" {
						name = attr.Value
					}

					if attr.Name.Local == "date" {
						date, err = time.Parse("2006-01-02", attr.Value)
						if err != nil {
							return nil, fmt.Errorf("could not parse date: %w", err)
						}
					}
				}
			}

			if name == "SIMC" && !date.IsZero() && t.Name.Local == "row" {
				childs, err := fetchChildren(decoder)
				if err != nil {
					return nil, fmt.Errorf("could not fetch children: %w", err)
				}

				// nasty hack to determine variant/name, just amount of children nodes of row element, should compare fields
				if len(childs) == 3 {
					name = "WMRODZ"
				} else if len(childs) == 14 {
					variant = "S" // Statystyczny
				} else if len(childs) != 10 {
					return nil, fmt.Errorf("unexpected amount of children nodes: %d", len(childs))
				}

				break
			}
		}
	}

	return &Dataset{Name: name, Variant: variant, Date: date}, nil
}

func fetchChildren(decoder *xml.Decoder) ([]string, error) {
	var children []string

	for {
		childToken, err := decoder.Token()
		if err != nil {
			return nil, err
		}

		if childEnd, ok := childToken.(xml.EndElement); ok && childEnd.Name.Local == "row" {
			break
		}

		if childStart, ok := childToken.(xml.StartElement); ok {
			var content string
			err := decoder.DecodeElement(&content, &childStart)
			if err != nil {
				return nil, err
			}

			children = append(children, childStart.Name.Local)
		}
	}

	return children, nil
}
