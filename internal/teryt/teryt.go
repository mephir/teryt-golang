package teryt

import (
	"github.com/mephir/teryt-golang/internal/collection"
	"github.com/mephir/teryt-golang/internal/dataset/model"
)

type Teryt struct {
	Voivodeships      *collection.Collection[model.Voivodeship]
	Counties          *collection.Collection[model.County]
	Municipalities    *collection.Collection[model.Municipality]
	MunicipalityTypes *collection.Collection[model.MunicipalityType]
	LocalityTypes     *collection.Collection[model.LocalityType]
}

func NewInstance() *Teryt {
	return &Teryt{
		Voivodeships:   collection.NewCollection[model.Voivodeship](),
		Counties:       collection.NewCollection[model.County](),
		Municipalities: collection.NewCollection[model.Municipality](),
		// MunicipalityTypes: model.GetMunicipalityTypesCollection(),
		// LocalityTypes: collection.NewCollection[model.LocalityType](),
	}
}
