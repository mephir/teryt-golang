package dataset

import (
	"fmt"
	"strings"
	"time"
)

type Dataset struct {
	Name    string
	Variant string
}

var Datasets = []Dataset{
	{"TERC", "U"},
	{"TERC", "A"},
	{"SIMC", "U"},
	{"SIMC", "A"},
	{"SIMC", "S"},
	{"ULIC", "U"},
	{"ULIC", "A"},
	{"WMRODZ", ""},
}

var variants = map[string]string{
	"U": "Urzedowy",
	"A": "Adresowy",
	"S": "Statystyczny",
}

func (d *Dataset) Validate() error {
	for _, dataset := range Datasets {
		if d.Name == dataset.Name && d.Variant == dataset.Variant {
			return nil
		}
	}

	return fmt.Errorf("invalid dataset")
}

func (d *Dataset) VariantName() string {
	return variants[d.Variant]
}

func (d *Dataset) ToString() string {
	if d.Name == "WMRODZ" {
		return d.Name
	}

	return fmt.Sprintf("%s %s", d.Name, d.VariantName())
}

func (d *Dataset) ToFilename(date time.Time) string {
	return fmt.Sprintf("%s_%s.zip", strings.ReplaceAll(d.ToString(), " ", "_"), date.Format("2006-01-02"))
}

func (d *Dataset) ToTarget() string {
	if d.Name == "WMRODZ" {
		return "ctl00$body$BRodzMiejPobierz"
	}

	return fmt.Sprintf("ctl00$body$B%s%sPobierz", d.Name, d.VariantName())
}
