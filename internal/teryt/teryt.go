package teryt

import (
	"fmt"

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
	Streets           *collection.Collection[model.Street]
}

func NewInstance() *Teryt {
	return &Teryt{
		Voivodeships:      collection.NewCollection[model.Voivodeship](),
		Counties:          collection.NewCollection[model.County](),
		Municipalities:    collection.NewCollection[model.Municipality](),
		MunicipalityTypes: model.GetMunicipalityTypesCollection(),
		LocalityTypes:     collection.NewCollection[model.LocalityType](),
		Localities:        collection.NewCollection[model.Locality](),
		Streets:           collection.NewCollection[model.Street](),
	}
}

func (t *Teryt) LoadFromFiles(paths ...string) error {
	var parsers []parser.Parser

	for _, path := range paths {
		parser, err := parser.Open(path)
		if err != nil {
			return err
		}

		parsers = append(parsers, parser)
	}

	if err := validateParserSet(&parsers); err != nil {
		return err
	}

	handler := newParsingHandler(t)
	for _, parser := range parsers {
		handler.Add(parser)
	}

	handler.Wait()

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
