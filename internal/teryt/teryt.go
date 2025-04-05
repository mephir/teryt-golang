package teryt

import (
	"fmt"
	"sort"

	"github.com/mephir/teryt-golang/internal/collection"
	"github.com/mephir/teryt-golang/internal/dataset/model"
	"github.com/mephir/teryt-golang/internal/parser"
)

type Teryt struct {
	Voivodeships      *collection.Collection[model.Voivodeship]
	Counties          *collection.Collection[model.County]
	Municipalities    *collection.Collection[model.Municipality]
	MunicipalityTypes *collection.Collection[model.MunicipalityType]
	LocalityTypes     *collection.Collection[model.LocalityType]
	Localities        *collection.Collection[model.Locality]
}

func NewInstance() *Teryt {
	return &Teryt{
		Voivodeships:      collection.NewCollection[model.Voivodeship](),
		Counties:          collection.NewCollection[model.County](),
		Municipalities:    collection.NewCollection[model.Municipality](),
		MunicipalityTypes: model.GetMunicipalityTypesCollection(),
		LocalityTypes:     collection.NewCollection[model.LocalityType](),
		Localities:        collection.NewCollection[model.Locality](),
	}
}

func (t *Teryt) LoadFromFiles(paths ...string) error {
	var parsers []parser.Parser

	for _, path := range paths {
		parser, err := parser.Open(path)
		if err != nil {
			return err
		}
		defer parser.Close()

		parsers = append(parsers, parser)
	}

	if len(parsers) > 1 {
		sortParsers(&parsers)
	}

	if err := validateParserSet(&parsers); err != nil {
		return err
	}

	handler := t.newParsingHandler()

	for _, parser := range parsers {
		if err := t.handleParsing(handler, parser); err != nil {
			return err
		}
	}
	handler.Close()

	return nil
}

func (t *Teryt) handleParsing(h *parsingHandler, p parser.Parser) error {
	if p == nil {
		return fmt.Errorf("parser is nil")
	}
	switch p.GetDataset().Name {
	case "TERC":
		h.parseTerc(p)
	case "WMRODZ":
		h.parseWmrodz(p)
	// case "SIMC":
	// 	return t.parseSimc(parser)
	// case "ULIC":
	// 	return t.parseUlic(parser)
	default:
		return fmt.Errorf("unknown dataset: %s", p.GetDataset().Name)
	}

	return nil
}

func validateParserSet(parsers *[]parser.Parser) error {
	exists := make(map[string]bool)

	for _, parser := range *parsers {
		if exists[parser.GetDataset().Name] {
			return fmt.Errorf("duplicate dataset found: %s", parser.GetDataset().Name)
		}
		exists[parser.GetDataset().Name] = true
	}

	return nil
}

// sort parsers in order to eliminate unecessary orphans, when parsing multiple datasets at once
func sortParsers(parsers *[]parser.Parser) {
	if len(*parsers) < 2 {
		return
	}

	weights := map[string]int{
		"TERC":   0,
		"WMRODZ": 1,
		"SIMC":   2,
		"ULIC":   3,
	}

	sort.Slice(*parsers, func(i, j int) bool {
		return weights[(*parsers)[i].GetDataset().Name] < weights[(*parsers)[j].GetDataset().Name]
	})
}

func (teryt *Teryt) newParsingHandler() *parsingHandler {
	return &parsingHandler{
		teryt: teryt,

		voivodeshipsChan:   make(chan *model.Voivodeship),
		countiesChan:       make(chan *model.County),
		municipalitiesChan: make(chan *model.Municipality),
		localityTypesChan:  make(chan *model.LocalityType),
		localitiesChan:     make(chan *model.Locality),
	}
}
