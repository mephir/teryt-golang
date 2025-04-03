package datastruct

import "github.com/mephir/teryt-golang/internal/dataset/model"

type Wmrodz struct {
	Id        uint   `xml:"RM"`      // Identyfikator, Identifier
	Name      string `xml:"NAZWA"`   // Nazwa, Name
	UpdatedAt AsOf   `xml:"STAN_NA"` // Data aktualizacji, Update date
}

func (w Wmrodz) ToModel() model.Wmrodz {
	return model.Wmrodz{
		Id:        w.Id,
		Name:      w.Name,
		UpdatedAt: w.UpdatedAt.Time,
	}
}
