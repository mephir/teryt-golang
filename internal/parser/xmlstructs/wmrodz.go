package xmlstructs

type Wmrodz struct {
	Id        uint8  `xml:"RM"`      // Identyfikator, Identifier
	Nazwa     string `xml:"NAZWA"`   // Nazwa, Name
	UpdatedAt StanNa `xml:"STAN_NA"` // Data aktualizacji, Update date
}
