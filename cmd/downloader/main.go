package main

import (
	"fmt"
	"time"

	"github.com/mephir/teryt-golang/internal/dataset"
	"github.com/mephir/teryt-golang/internal/downloader"
)

func main() {
	dataset := dataset.Dataset{Name: "SIMC", Variant: "A"}
	filename := dataset.ToFilename(time.Now().Local())

	fmt.Println("Downloading file")
	err := downloader.DownloadDataset(dataset, filename)
	if err != nil {
		panic(err)
	}

	fmt.Println("Extracting files")
	err = downloader.ExtractAllFiles(filename, ".")
	if err != nil {
		panic(err)
	}

	fmt.Println("Done")
}
