package tests

import (
	"encoding/xml"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/mephir/teryt-golang/internal/dataset"
)

type fields struct {
	Name    string
	Variant string
	Date    time.Time
}

func TestClassifier_Determine(t *testing.T) {
	var sourcePath, err = filepath.Abs("./data")
	if err != nil {
		t.Fatalf("could not determine absolute path: %v", err)
	}

	tests := []struct {
		name    string
		want    fields
		wantErr bool
		path    string
	}{
		{"Empty file", fields{Name: "", Variant: "", Date: time.Time{}}, true, filepath.Join(sourcePath, "empty.xml")},
		{"SIMC Adresowy", fields{Name: "SIMC", Variant: "A", Date: time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)}, false, filepath.Join(sourcePath, "SIMC_Adresowy_2025-04-01.xml")},
		{"SIMC Statystyczny", fields{Name: "SIMC", Variant: "S", Date: time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)}, false, filepath.Join(sourcePath, "SIMC_Statystyczny_2025-04-01.xml")},
		{"WMRODZ", fields{Name: "WMRODZ", Variant: "", Date: time.Date(2013, time.February, 28, 0, 0, 0, 0, time.UTC)}, false, filepath.Join(sourcePath, "WMRODZ_2025-04-01.xml")},
		{"WMRODZ non default name", fields{Name: "WMRODZ", Variant: "", Date: time.Date(2013, time.February, 28, 0, 0, 0, 0, time.UTC)}, false, filepath.Join(sourcePath, "wmrodz_without_proper_name.xml")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dataset.Determine(tt.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("Classifier.DetermineByContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && got.Name != tt.want.Name {
				t.Errorf("Classifier.DetermineByContent().Name = %v, want %v", got.Name, tt.want.Name)
			}
			if got != nil && !got.Date.Equal(tt.want.Date) {
				t.Errorf("Classifier.DetermineByContent().Date = %v, want %v", got.Date, tt.want.Date)
			}
			if got != nil && got.Variant != tt.want.Variant {
				t.Errorf("Classifier.DetermineByContent().Variant = %v, want %v", got.Variant, tt.want.Variant)
			}
		})
	}
}

func TestClassifier_DetermineByFilename(t *testing.T) {
	tests := []struct {
		name     string
		want     fields
		wantErr  bool
		filename string
	}{
		{"Empty file", fields{Name: "", Variant: "", Date: time.Time{}}, true, "empty.xml"},
		{"SIMC Adresowy", fields{Name: "SIMC", Variant: "A", Date: time.Date(2025, time.April, 1, 0, 0, 0, 0, time.UTC)}, false, "SIMC_Adresowy_2025-04-01.xml"},
		{"SIMC Urzedowy", fields{Name: "SIMC", Variant: "U", Date: time.Date(2025, time.April, 1, 0, 0, 0, 0, time.UTC)}, false, "SIMC_Urzedowy_2025-04-01.xml"},
		{"SIMC Statystyczny", fields{Name: "SIMC", Variant: "S", Date: time.Date(2025, time.April, 1, 0, 0, 0, 0, time.UTC)}, false, "SIMC_Statystyczny_2025-04-01.xml"},
		{"TERC Adresowy", fields{Name: "TERC", Variant: "A", Date: time.Date(2025, time.April, 1, 0, 0, 0, 0, time.UTC)}, false, "TERC_Adresowy_2025-04-01.xml"},
		{"TERC Urzedowy", fields{Name: "TERC", Variant: "U", Date: time.Date(2025, time.April, 1, 0, 0, 0, 0, time.UTC)}, false, "TERC_Urzedowy_2025-04-01.xml"},
		{"ULIC Adresowy", fields{Name: "ULIC", Variant: "A", Date: time.Date(2025, time.April, 1, 0, 0, 0, 0, time.UTC)}, false, "ULIC_Adresowy_2025-04-01.xml"},
		{"ULIC Urzedowy", fields{Name: "ULIC", Variant: "U", Date: time.Date(2025, time.April, 1, 0, 0, 0, 0, time.UTC)}, false, "ULIC_Urzedowy_2025-04-01.xml"},
		{"WMRODZ", fields{Name: "WMRODZ", Variant: "", Date: time.Date(2025, time.April, 1, 0, 0, 0, 0, time.UTC)}, false, "WMRODZ_2025-04-01.xml"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dataset.DetermineByFilename(tt.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("Classifier.DetermineByContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && got.Name != tt.want.Name {
				t.Errorf("Classifier.DetermineByContent().Name = %v, want %v", got.Name, tt.want.Name)
			}
			if got != nil && !got.Date.Equal(tt.want.Date) {
				t.Errorf("Classifier.DetermineByContent().Date = %v, want %v", got.Date, tt.want.Date)
			}
			if got != nil && got.Variant != tt.want.Variant {
				t.Errorf("Classifier.DetermineByContent().Variant = %v, want %v", got.Variant, tt.want.Variant)
			}
		})
	}

}

func TestClassifier_DetermineByContent(t *testing.T) {
	var basePath, err = filepath.Abs("./data")
	if err != nil {
		t.Fatalf("could not determine absolute path: %v", err)
	}

	tests := []struct {
		name    string
		want    fields
		wantErr bool
		xml     string
	}{
		{"Empty file", fields{Name: "", Variant: "", Date: time.Time{}}, true, filepath.Join(basePath, "empty.xml")},
		{"SIMC Adresowy", fields{Name: "SIMC", Variant: "", Date: time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)}, false, filepath.Join(basePath, "SIMC_Adresowy_2025-04-01.xml")},
		{"SIMC Urzedowy", fields{Name: "SIMC", Variant: "", Date: time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)}, false, filepath.Join(basePath, "SIMC_Urzedowy_2025-04-01.xml")},
		{"SIMC Statystyczny", fields{Name: "SIMC", Variant: "S", Date: time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)}, false, filepath.Join(basePath, "SIMC_Statystyczny_2025-04-01.xml")},
		{"TERC Adresowy", fields{Name: "TERC", Variant: "", Date: time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)}, false, filepath.Join(basePath, "TERC_Adresowy_2025-04-01.xml")},
		{"TERC Urzedowy", fields{Name: "TERC", Variant: "", Date: time.Date(2025, time.January, 1, 0, 0, 0, 0, time.UTC)}, false, filepath.Join(basePath, "TERC_Urzedowy_2025-04-01.xml")},
		{"ULIC Adresowy", fields{Name: "ULIC", Variant: "", Date: time.Date(2025, time.March, 31, 0, 0, 0, 0, time.UTC)}, false, filepath.Join(basePath, "ULIC_Adresowy_2025-04-01.xml")},
		{"ULIC Urzedowy", fields{Name: "ULIC", Variant: "", Date: time.Date(2025, time.March, 31, 0, 0, 0, 0, time.UTC)}, false, filepath.Join(basePath, "ULIC_Urzedowy_2025-04-01.xml")},
		{"WMRODZ", fields{Name: "WMRODZ", Variant: "", Date: time.Date(2013, time.February, 28, 0, 0, 0, 0, time.UTC)}, false, filepath.Join(basePath, "WMRODZ_2025-04-01.xml")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fh, err := os.Open(tt.xml)
			if err != nil {
				t.Fatalf("could not open file: %v", err)
			}
			defer fh.Close()

			reader := xml.NewDecoder(fh)
			got, err := dataset.DetermineByContent(reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("Classifier.DetermineByContent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != nil && got.Name != tt.want.Name {
				t.Errorf("Classifier.DetermineByContent().Name = %v, want %v", got.Name, tt.want.Name)
			}
			if got != nil && !got.Date.Equal(tt.want.Date) {
				t.Errorf("Classifier.DetermineByContent().Date = %v, want %v", got.Date, tt.want.Date)
			}
			if got != nil && got.Variant != tt.want.Variant {
				t.Errorf("Classifier.DetermineByContent().Variant = %v, want %v", got.Variant, tt.want.Variant)
			}
		})
	}
}
