package tests

import (
	"testing"

	"github.com/mephir/teryt-golang/internal/dataset"
)

func TestDataset_Validate(t *testing.T) {
	type fields struct {
		Name    string
		Variant string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"Valid", fields{"SIMC", "S"}, false},
		{"Invalid name", fields{"INVALID", "A"}, true},
		{"Invalid variant", fields{"SIMC", "INVALID"}, true},
		{"Invalid variant", fields{"WMRODZ", "A"}, true},
		{"WMRODZ proper variant", fields{"WMRODZ", ""}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &dataset.Dataset{
				Name:    tt.fields.Name,
				Variant: tt.fields.Variant,
			}
			if err := d.Validate(); (err != nil) != tt.wantErr {
				t.Errorf("Dataset.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDataset_ToTarget(t *testing.T) {
	type fields struct {
		Name    string
		Variant string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"WMRODZ", fields{"WMRODZ", ""}, "ctl00$body$BRodzMiejPobierz"},
		{"SIMC", fields{"SIMC", "A"}, "ctl00$body$BSIMCAdresowyPobierz"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &dataset.Dataset{
				Name:    tt.fields.Name,
				Variant: tt.fields.Variant,
			}
			if got := d.ToTarget(); got != tt.want {
				t.Errorf("Dataset.ToTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
