package datastruct

import "github.com/mephir/teryt-golang/internal/dataset/model"

type Ulic struct {
	Woj          uint                   `xml:"WOJ"`      // Województwo, Voivodeship
	Pow          uint                   `xml:"POW"`      // Powiat, County
	Gmi          uint                   `xml:"GMI"`      // Gmina, Municipality
	Rodz         model.MunicipalityType `xml:"RODZ_GMI"` // Rodzaj jednostki, Unit type
	Sym          uint                   `xml:"SYM"`      // Identyfikator miejścowości, Locality identifier
	SymUl        uint                   `xml:"SYM_UL"`   // Identyfikator ulicy, Street identifier
	Type         string                 `xml:"CECHA"`    // Określenie rodzaju ulicy, Street type determination
	SortableName string                 `xml:"NAZWA_1"`  // Nazwa uzywana do sortowania alfabetycznego, Name used for alphabetical sorting
	NamePrefix   string                 `xml:"NAZWA_2"`  // Pozostajaca czesc nazwy, Remaining part of the name
	AsOf         AsOf                   `xml:"STAN_NA"`  // Data aktualizacji, Update date
}

func (u Ulic) ToModel() (model.Model, error) {
	return &model.Street{
		VoivodeshipId:      u.Woj,
		CountyId:           u.Pow,
		MunicipalityId:     u.Gmi,
		MunicipalityTypeId: u.Rodz.Id,
		Id:                 u.SymUl,
		LocalityId:         u.Sym,
		SortableName:       u.SortableName,
		NamePrefix:         u.NamePrefix,
		Type:               u.Type,
		AsOf:               u.AsOf.Time,
	}, nil
}
