package uuid

import (
	"crypto/sha1"
	"time"

	"github.com/google/uuid"
)

type UuidData struct {
	VoivodeshipId      uint8
	CountyId           uint8
	MunicipalityId     uint8
	MunicipalityTypeId uint8
	AsOf               time.Time
	LocalityId         *uint32
	StreetId           *uint32
	Name               string // used for hash to fill missing bits
}

func (u *UuidData) Encode() (uuid.UUID, error) {
	var raw [16]byte

	teryt := (uint32(u.VoivodeshipId) << 18) | (uint32(u.CountyId) << 11) | (uint32(u.MunicipalityId) << 4) | uint32(u.MunicipalityTypeId)
	date := uint32(u.AsOf.Year())<<9 | uint32(u.AsOf.YearDay())
	hashShifted := shiftLeft4BitsKeepFirstNibble(sha1.Sum([]byte(u.Name))) // Fill remaining UUID bits with the hashed name

	// Fill the UUID with the TERYT
	raw[0] = byte(teryt >> 16)
	raw[1] = byte(teryt >> 8)
	raw[2] = byte(teryt)

	if u.LocalityId != nil {
		raw[3] = byte((*u.LocalityId >> 16) & 0xFF)
		raw[4] = byte((*u.LocalityId >> 8) & 0xFF)
		raw[5] = byte(*u.LocalityId & 0xFF)
	}

	raw[6] = hashShifted[0] | 0x80 // Set version (8), shifted hash contains always 0 as first 4 bits
	raw[7] = hashShifted[1]
	raw[8] = (hashShifted[2] & 0x3F) | 0x80 // Set variant (IETF), ensure the first 2 bits are 10
	raw[9] = hashShifted[3]

	// Fill the UUID with the date
	raw[13] = byte(date >> 16)
	raw[14] = byte(date >> 8)
	raw[15] = byte(date)

	if u.StreetId != nil {
		raw[10] = byte((*u.StreetId >> 16) & 0xFF)
		raw[11] = byte((*u.StreetId >> 8) & 0xFF)
		raw[12] = byte(*u.StreetId & 0xFF)
	} else {
		copy(raw[10:13], hashShifted[4:])
	}

	var id uuid.UUID
	if err := id.UnmarshalBinary(raw[:]); err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func shiftRight4Bits(input [21]byte) [20]byte {
	var output [20]byte

	output[0] = (input[0] << 4) | (input[1] >> 4)

	for i := 1; i < 19; i++ {
		output[i] = (input[i] << 4) | (input[i+1] >> 4)
	}

	output[19] = (input[19] << 4) | (input[20] >> 4)

	return output
}

func shiftLeft4BitsKeepFirstNibble(input [20]byte) [21]byte {
	var output [21]byte

	output[0] = (input[0] & 0xF0) >> 4

	for i := 0; i < len(input)-1; i++ {
		output[i+1] = (input[i] << 4) | (input[i+1] >> 4)
	}

	output[20] = input[len(input)-1] << 4

	return output
}
