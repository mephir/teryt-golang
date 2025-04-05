package datastruct

import "github.com/mephir/teryt-golang/internal/dataset/model"

type Ulic struct {
	Woj           uint8                  `xml:"WOJ"`      // Województwo, Voivodeship
	Pow           uint8                  `xml:"POW"`      // Powiat, County
	Gmi           uint8                  `xml:"GMI"`      // Gmina, Municipality
	Rodz          model.MunicipalityType `xml:"RODZ_GMI"` // Rodzaj jednostki, Unit type
	Sym           uint32                 `xml:"SYM"`      // Identyfikator miejścowości, Locality identifier
	SymUl         uint32                 `xml:"SYM_UL"`   // Identyfikator ulicy, Street identifier
	Type          string                 `xml:"CECHA"`    // Określenie rodzaju ulicy, Street type determination
	Name          string                 `xml:"NAZWA_1"`  // Nazwa uzywana do sortowania alfabetycznego, Name used for alphabetical sorting
	NameRemaining string                 `xml:"NAZWA_2"`  // Pozostajaca czesc nazwy, Remaining part of the name
}
