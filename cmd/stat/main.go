package main

import (
	"fmt"
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

	fmt.Printf("Voivodeships: %d\n", teryt.Voivodeships.Count())
	fmt.Printf("Counties: %d\n", teryt.Counties.Count())
	fmt.Printf("Municipalities: %d\n", teryt.Municipalities.Count())
	fmt.Printf("MunicipalityTypes: %d\n", teryt.MunicipalityTypes.Count())
	fmt.Printf("Localities: %d\n", teryt.Localities.Count())
	fmt.Printf("LocalityTypes: %d\n", teryt.LocalityTypes.Count())
	fmt.Printf("Streets: %d\n", teryt.Streets.Count())
}
