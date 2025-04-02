package xmlstructs

type Terc struct {
	Woj       uint8  `xml:"WOJ"`       // Województwo, Voivodeship
	Pow       uint8  `xml:"POW"`       // Powiat, County
	Gmi       uint8  `xml:"GMI"`       // Gmina, Municipality
	Rodz      Rodz   `xml:"RODZ"`      // Rodzaj jednostki, Unit type
	Nazwa     string `xml:"NAZWA"`     // Nazwa, Name
	NazwaDod  string `xml:"NAZWA_DOD"` // określenie jednostki, Unit determination
	UpdatedAt StanNa `xml:"STAN_NA"`   // Data aktualizacji, Update date
}
