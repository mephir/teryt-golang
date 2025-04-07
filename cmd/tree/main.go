package main

import (
	"github.com/google/uuid"
	"github.com/mephir/teryt-golang/internal/tree/avltree"
)

func main() {
	// basePath, err := filepath.Abs("_zrzuty/")
	// if err != nil {
	// 	panic(err)
	// }

	// paths := []string{
	// 	filepath.Join(basePath, "TERC_Adresowy_2025-04-01.xml"),
	// 	// filepath.Join(basePath, "WMRODZ_2025-04-01.xml"),
	// 	// filepath.Join(basePath, "SIMC_Adresowy_2025-04-01.xml"),
	// 	// filepath.Join(basePath, "ULIC_Adresowy_2025-04-01.xml"),
	// }

	// teryt := teryt.NewInstance()
	// teryt.LoadFromFiles(paths...)

	tree := &avltree.AvlTree[uuid.UUID]{}
	for range 10 {
		tree.Insert(uuid.New())
	}
	// tree.Insert(0)
	// tree.Insert(1)
	// tree.Insert(2)
	// tree.Insert(3)
	// tree.Insert(4)
	// tree.Insert(5)
	// tree.Insert(6)
	// tree.Insert(7)

	tree.Print(true)
}
