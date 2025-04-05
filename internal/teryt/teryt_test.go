package teryt

import (
	"testing"

	"github.com/mephir/teryt-golang/internal/dataset"
	"github.com/mephir/teryt-golang/internal/parser"
)

func TestTeryt_sortParsers(t *testing.T) {
	tests := []struct {
		name    string
		parsers []parser.Parser
		want    []string
	}{
		{
			name: "Test single Dataset",
			parsers: []parser.Parser{
				&parser.XmlParser{Dataset: dataset.Dataset{Name: "SIMC"}},
			},
			want: []string{"SIMC"},
		},
		{
			name: "Already sorted",
			parsers: []parser.Parser{
				&parser.XmlParser{Dataset: dataset.Dataset{Name: "TERC"}},
				&parser.XmlParser{Dataset: dataset.Dataset{Name: "WMRODZ"}},
				&parser.XmlParser{Dataset: dataset.Dataset{Name: "SIMC"}},
				&parser.XmlParser{Dataset: dataset.Dataset{Name: "ULIC"}},
			},
			want: []string{"TERC", "WMRODZ", "SIMC", "ULIC"},
		},
		{
			name: "Unsorted",
			parsers: []parser.Parser{
				&parser.XmlParser{Dataset: dataset.Dataset{Name: "SIMC"}},
				&parser.XmlParser{Dataset: dataset.Dataset{Name: "TERC"}},
				&parser.XmlParser{Dataset: dataset.Dataset{Name: "ULIC"}},
				&parser.XmlParser{Dataset: dataset.Dataset{Name: "WMRODZ"}},
			},
			want: []string{"TERC", "WMRODZ", "SIMC", "ULIC"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sortParsers(&tt.parsers)
			var keys []string
			for _, parser := range tt.parsers {
				keys = append(keys, parser.GetDataset().Name)
			}

			for i, v := range tt.want {
				if keys[i] != v {
					t.Fatalf("Expected %v, got %v", tt.want, keys)
				}
			}
		})
	}
}

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
