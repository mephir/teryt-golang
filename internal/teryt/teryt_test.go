package teryt

import (
	"testing"

	"github.com/mephir/teryt-golang/internal/dataset"
	"github.com/mephir/teryt-golang/internal/parser"
)

func TestTeryt_validatePaarserSet(t *testing.T) {
	tests := []struct {
		name    string
		parsers []parser.Parser
		wantErr bool
	}{
		{
			name: "Valid parsers",
			parsers: []parser.Parser{
				&parser.XmlParser{Dataset: dataset.Dataset{Name: "TERC"}},
				&parser.XmlParser{Dataset: dataset.Dataset{Name: "WMRODZ"}},
				&parser.XmlParser{Dataset: dataset.Dataset{Name: "SIMC"}},
				&parser.XmlParser{Dataset: dataset.Dataset{Name: "ULIC"}},
			},
			wantErr: false,
		},
		{
			name: "Duplicate parsers",
			parsers: []parser.Parser{
				&parser.XmlParser{Dataset: dataset.Dataset{Name: "TERC"}},
				&parser.XmlParser{Dataset: dataset.Dataset{Name: "WMRODZ"}},
				&parser.XmlParser{Dataset: dataset.Dataset{Name: "TERC"}},
				&parser.XmlParser{Dataset: dataset.Dataset{Name: "ULIC"}},
				&parser.XmlParser{Dataset: dataset.Dataset{Name: "SIMC"}},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateParserSet(&tt.parsers); (err != nil) != tt.wantErr {
				t.Errorf("validateParserSet() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
