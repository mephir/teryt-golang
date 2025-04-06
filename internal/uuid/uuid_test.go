package uuid

import (
	"testing"
	"time"
)

func TestUuid_Encoding(t *testing.T) {
	tests := []struct {
		name     string
		data     UuidData
		expected string
	}{
		{
			name: "Test Voivodeship",
			data: UuidData{
				VoivodeshipId:      32,
				CountyId:           0,
				MunicipalityId:     0,
				MunicipalityTypeId: 0,
				AsOf:               time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
				StreetId:           nil,
				Name:               "województwo zachodniopomorskie",
			},
			expected: "80000000-0000-84a2-ae3a-1411d30fd201",
		},
		{
			name: "Test County",
			data: UuidData{
				VoivodeshipId:      32,
				CountyId:           9,
				MunicipalityId:     0,
				MunicipalityTypeId: 0,
				AsOf:               time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
				Name:               "powiat koszaliński",
			},
			expected: "80480000-0000-875b-89f6-1afafe0fd201",
		},
		{
			name: "Test Municipality",
			data: UuidData{
				VoivodeshipId:      32,
				CountyId:           9,
				MunicipalityId:     5,
				MunicipalityTypeId: 4,
				AsOf:               time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
				Name:               "miasto Mielno",
			},
			expected: "80485400-0000-8232-aad1-4121110fd201",
		},
		{
			name: "Test Municipality 2",
			data: UuidData{
				VoivodeshipId:      32,
				CountyId:           9,
				MunicipalityId:     5,
				MunicipalityTypeId: 5,
				AsOf:               time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
				Name:               "obszar wiejski Mielno",
			},
			expected: "80485500-0000-89e6-91f9-32ce460fd201",
		},
		{
			name: "Test Locality",
			data: UuidData{
				VoivodeshipId:      32,
				CountyId:           9,
				MunicipalityId:     5,
				MunicipalityTypeId: 4,
				LocalityId:         func() *uint32 { v := uint32(308353); return &v }(),
				AsOf:               time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
				Name:               "Mielno",
			},
			expected: "80485404-b481-8de0-b688-ecf65f0fd201",
		},
		{
			name: "Test Street",
			data: UuidData{
				VoivodeshipId:      32,
				CountyId:           9,
				MunicipalityId:     5,
				MunicipalityTypeId: 5,
				LocalityId:         func() *uint32 { v := uint32(308353); return &v }(),
				StreetId:           func() *uint32 { v := uint32(19016); return &v }(),
				AsOf:               time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
				Name:               "Róży Wiatrów",
			},
			expected: "80485504-b481-85f8-a5bf-004a480fd201",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u, err := tt.data.Encode()
			if err != nil {
				t.Fatalf("failed to encode UUID: %v", err)
			}

			if u.String() != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, u.String())
			}
		})
	}
}
