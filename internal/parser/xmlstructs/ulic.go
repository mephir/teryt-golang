package xmlstructs

type Ulic struct {
	Woj            uint8  `xml:"WOJ"`      // Województwo, Voivodeship
	Pow            uint8  `xml:"POW"`      // Powiat, County
	Gmi            uint8  `xml:"GMI"`      // Gmina, Municipality
	Rodz           Rodz   `xml:"RODZ_GMI"` // Rodzaj jednostki, Unit type
	Sym            uint32 `xml:"SYM"`      // Identyfikator miejścowości, Locality identifier
	SymUl          uint32 `xml:"SYM_UL"`   // Identyfikator ulicy, Street identifier
	Cecha          string `xml:"CECHA"`    // Określenie rodzaju ulicy, Street type determination
	Nazwa          string `xml:"NAZWA_1"`  // Nazwa uzywana do sortowania alfabetycznego, Name used for alphabetical sorting
	NazwaPozostala string `xml:"NAZWA_2"`  // Pozostajaca czesc nazwy, Remaining part of the name
}
