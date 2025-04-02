package xmlstructs

type Simc struct {
	Woj       uint8  `xml:"WOJ"`      // Wojewodztwo, Voivodeship
	Pow       uint16 `xml:"POW"`      // Powiat, County
	Gmi       uint16 `xml:"GMI"`      // Gmina, Municipality
	Rodz      Rodz   `xml:"RODZ_GMI"` // Rodzaj gminy, Municipality type
	Rm        uint8  `xml:"RM"`       // Rodzaj miejscowosci, Locality type
	Mz        bool   `xml:"MZ"`       // Wystepowanie nazwy zwyczjowej, Common name presence
	Nazwa     string `xml:"NAZWA"`    // Nazwa, Name
	Sym       uint32 `xml:"SYM"`      // Identyfikator miejscowosci, Locality identifier
	Sympod    uint32 `xml:"SYMPOD"`   // Identyfikator miejscowosci podstawowej, Main locality identifier
	UpdatedAt StanNa `xml:"STAN_NA"`  // Data aktualizacji, Update date
}
