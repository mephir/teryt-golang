package datastruct

import "github.com/mephir/teryt-golang/internal/dataset/model"

type Simc struct {
	Woj    uint                   `xml:"WOJ"`      // Wojewodztwo, Voivodeship
	Pow    uint                   `xml:"POW"`      // Powiat, County
	Gmi    uint                   `xml:"GMI"`      // Gmina, Municipality
	Rodz   model.MunicipalityType `xml:"RODZ_GMI"` // Rodzaj gminy, Municipality type
	Rm     uint                   `xml:"RM"`       // Rodzaj miejscowosci, Locality type
	Mz     bool                   `xml:"MZ"`       // Wystepowanie nazwy zwyczajowej, Common name presence
	Name   string                 `xml:"NAZWA"`    // Nazwa, Name
	Sym    uint                   `xml:"SYM"`      // Identyfikator miejscowosci, Locality identifier
	Sympod uint                   `xml:"SYMPOD"`   // Identyfikator miejscowosci podstawowej, Main locality identifier
	AsOf   AsOf                   `xml:"STAN_NA"`  // Data aktualizacji, Update date
}

func (s Simc) ToModel() (model.Model, error) {
	return &model.Locality{
		Id:               s.Sym,
		TypeId:           s.Rm,
		CommonName:       s.Mz,
		Name:             s.Name,
		ParentLocalityId: s.Sympod,
		MunicipalityId:   s.GetMunicipalityIdentifier(),
		AsOf:             s.AsOf.Time,
	}, nil
}

func (s Simc) GetMunicipalityIdentifier() uint {
	return s.Woj*100000 + s.Pow*1000 + s.Gmi*10 + s.Rodz.Id
}
