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

func (d *Dataset) Validate() error {
	datasets := map[string][]string{
		"TERC":   {"U", "A"},
		"SIMC":   {"U", "A", "S"},
		"ULIC":   {"U", "A"},
		"WMRODZ": {""},
	}

	if _, ok := datasets[d.Name]; !ok {
		return fmt.Errorf("invalid dataset name")
	}

	for _, variant := range datasets[d.Name] {
		if variant == d.Variant {
			return nil
		}
	}

	return fmt.Errorf("invalid dataset variant")
}

func (d *Dataset) VariantName() string {
	variants := map[string]string{
		"U": "Urzedowy",
		"A": "Adresowy",
		"S": "Statystyczny",
	}

	return variants[d.Variant]
}

func (d *Dataset) ToString() string {
	if d.Name == "WMRODZ" {
		return d.Name
	}

	return fmt.Sprintf("%s %s", d.Name, d.VariantName())
}

func (d *Dataset) ToFilename() string {
	return strings.ReplaceAll(fmt.Sprintf("%s_%s.zip", d.ToString(), time.Now().Local().Format("2006-01-02")), " ", "_")
}

func (d *Dataset) ToTarget() string {
	if d.Name == "WMRODZ" {
		return "ctl00$body$BRodzMiejPobierz"
	}

	return fmt.Sprintf("ctl00$body$B%s%sPobierz", d.Name, d.VariantName())
}
