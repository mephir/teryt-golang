package datastruct

import "github.com/mephir/teryt-golang/internal/dataset/model"

type SimcS struct {
	Woj       uint       `xml:"WOJ"`      // Wojewodztwo, Voivodeship
	Pow       uint16     `xml:"POW"`      // Powiat, County
	Gmi       uint16     `xml:"GMI"`      // Gmina, Municipality
	Rodz      model.Rodz `xml:"RODZ_GMI"` // Rodzaj gminy, Municipality type
	Rm        uint8      `xml:"RM"`       // Rodzaj miejscowosci, Locality type
	Mz        bool       `xml:"MZ"`       // Wystepowanie nazwy zwyczajowej, Common name presence
	Nmst      uint8      `xml:"NMST"`     // Numer miejscowosci statystycznej w ramach gminy, Statistical locality number within the municipality
	Nmsk      uint16     `xml:"NMSK"`     // numer miejscowości składowej w ramach miejscowości statystycznej, Number of a component locality within a statistical locality
	Symbm     uint8      `xml:"SYMBM"`    // Okreslenie miejscowosci, Locality determination
	Symstat   uint32     `xml:"SYMSTAT"`  // Identyfikator miejscowości statystycznej, do której należy dana miejscowość, Identifier of the statistical locality to which the given locality belongs
	Nazwa     string     `xml:"NAZWA"`    // Nazwa, Name
	Sym       uint32     `xml:"SYM"`      // Identyfikator miejscowosci, Locality identifier
	Sympod    uint32     `xml:"SYMPOD"`   // Identyfikator miejscowosci podstawowej, Main locality identifier
	UpdatedAt AsOf       `xml:"STAN_NA"`  // Data aktualizacji, Update date
}
