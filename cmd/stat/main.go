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
	}

	teryt := teryt.NewInstance()
	teryt.LoadFromFiles(paths...)

	fmt.Printf("Voivodeships: %d\n", teryt.Voivodeships.Count())
	fmt.Printf("Counties: %d\n", teryt.Counties.Count())
	fmt.Printf("Municipalities: %d\n", teryt.Municipalities.Count())
	fmt.Printf("LocalityTypes: %d\n", teryt.LocalityTypes.Count())

	// for v := range teryt.Voivodeships.Iterator() {
	// 	fmt.Printf("%s %s\n", v.UnitType, v.Name)
	// 	for _, c := range v.Counties {
	// 		fmt.Printf("\t%s %s\n", c.UnitType, c.Name)
	// 		for _, m := range c.Municipalities {
	// 			fmt.Printf("\t\t%s %s\n", m.UnitType, m.Name)
	// 		}
	// 	}
	// }
}
