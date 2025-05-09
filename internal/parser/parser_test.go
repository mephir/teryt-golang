package parser

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/mephir/teryt-golang/internal/dataset/datastruct"
)

var sourcePath string

func TestMain(m *testing.M) {
	source, err := filepath.Abs("./data")
	if err != nil {
		panic(err)
	}
	sourcePath = source

	code := m.Run()

	os.Exit(code)
}

func TestParser_Fetch(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		wantErr  bool
		wantType reflect.Type
	}{
		{"WMRODZ", filepath.Join(sourcePath, "WMRODZ_2025-04-01.xml"), false, reflect.TypeOf(&datastruct.Wmrodz{})},
		{"TERC", filepath.Join(sourcePath, "TERC_Adresowy_2025-04-01.xml"), false, reflect.TypeOf(&datastruct.Terc{})},
		{"SIMC", filepath.Join(sourcePath, "SIMC_Adresowy_2025-04-01.xml"), false, reflect.TypeOf(&datastruct.Simc{})},
		{"ULIC", filepath.Join(sourcePath, "ULIC_Adresowy_2025-04-01.xml"), false, reflect.TypeOf(&datastruct.Ulic{})},
		{"SIMC_S", filepath.Join(sourcePath, "SIMC_Statystyczny_2025-04-01.xml"), false, reflect.TypeOf(&datastruct.SimcS{})},
		{"WMRODZ_without_proper_filename", filepath.Join(sourcePath, "wmrodz_without_proper_name.xml"), false, reflect.TypeOf(&datastruct.Wmrodz{})},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser, err := Open(tt.filename)
			if err != nil {
				t.Fatalf("could not create parser: %v", err)
				return
			}
			defer parser.Close()

			item, err := parser.Fetch()
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.Fetch() error = %v, wantErr %v", err, tt.wantErr)
			}

			if reflect.TypeOf(item) != tt.wantType {
				t.Errorf("Parser.Fetch() = %v, want %v", reflect.TypeOf(item), tt.wantType)
			}
		})
	}
}

func TestParser_FetchAll(t *testing.T) {
	tests := []struct {
		name       string
		filename   string
		wantErr    bool
		wantType   reflect.Type
		wantLength uint8
	}{
		{"WMRODZ", filepath.Join(sourcePath, "WMRODZ_2025-04-01.xml"), false, reflect.TypeOf(&datastruct.Wmrodz{}), 12},
		{"TERC", filepath.Join(sourcePath, "TERC_Adresowy_2025-04-01.xml"), false, reflect.TypeOf(&datastruct.Terc{}), 1},
		{"SIMC", filepath.Join(sourcePath, "SIMC_Adresowy_2025-04-01.xml"), false, reflect.TypeOf(&datastruct.Simc{}), 1},
		{"ULIC", filepath.Join(sourcePath, "ULIC_Adresowy_2025-04-01.xml"), false, reflect.TypeOf(&datastruct.Ulic{}), 1},
		{"SIMC_S", filepath.Join(sourcePath, "SIMC_Statystyczny_2025-04-01.xml"), false, reflect.TypeOf(&datastruct.SimcS{}), 1},
		{"WMRODZ_without_proper_filename", filepath.Join(sourcePath, "wmrodz_without_proper_name.xml"), false, reflect.TypeOf(&datastruct.Wmrodz{}), 12},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser, err := Open(tt.filename)
			if err != nil {
				t.Fatalf("could not create parser: %v", err)
				return
			}
			defer parser.Close()

			items, err := parser.FetchAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.FetchAll() error = %v, wantErr %v", err, tt.wantErr)
			}

			if reflect.TypeOf(items[0]) != tt.wantType {
				t.Errorf("Parser.FetchAll() = %v, want %v", reflect.TypeOf(items[0]), tt.wantType)
			}

			if uint8(len(items)) != tt.wantLength {
				t.Errorf("Parser.FetchAll() = %v, want %v", len(items), tt.wantLength)
			}
		})
	}
}
