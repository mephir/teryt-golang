package model

import "time"

type Wmrodz struct {
	Id        uint
	Name      string
	UpdatedAt time.Time
}

func (w Wmrodz) Identifier() uint {
	return w.Id
}
