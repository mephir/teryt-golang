package main

import (
	"path/filepath"

	"github.com/mephir/teryt-golang/internal/teryt"
)

func main() {
	basePath, err := filepath.Abs("_zrzuty/")
	if err != nil {
		panic(err)
	}

	paths := []string{
		filepath.Join(basePath, "TERC_Adresowy_2025-04-01.xml"),
		filepath.Join(basePath, "WMRODZ_2025-04-01.xml"),
		filepath.Join(basePath, "SIMC_Adresowy_2025-04-01.xml"),
		filepath.Join(basePath, "ULIC_Adresowy_2025-04-01.xml"),
	}

	teryt := teryt.NewInstance()
	teryt.LoadFromFiles(paths...)

	// writer := writer.NewWriter("gob", nil)
	// writer.Write(teryt, "teryt.gob")
}
