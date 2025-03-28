package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/mephir/teryt-golang/internal/dataset"
	"github.com/mephir/teryt-golang/internal/downloader"
)

func extract_file(dataset dataset.Dataset) error {
	archive, err := zip.OpenReader(dataset.ToFilename(time.Now().Local()))
	if err != nil {
		return err
	}
	defer archive.Close()

	for _, file := range archive.File {
		fmt.Printf("Extracting %s\n", file.Name)

		if file.FileInfo().IsDir() {
			continue
		}

		src, err := file.Open()
		if err != nil {
			return err
		}

		dstFile, err := os.OpenFile(file.Name, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}

		_, err = io.Copy(dstFile, src)
		if err != nil {
			return err
		}

		src.Close()
		dstFile.Close()

	}

	return nil
}

func main() {
	dataset := dataset.Dataset{Name: "SIMC", Variant: "A"}

	fmt.Println("Downloading file")
	err := downloader.DownloadDataset(dataset, dataset.ToFilename(time.Now().Local()))
	if err != nil {
		panic(err)
	}

	fmt.Println("Extracting file")
	err = extract_file(dataset)
	if err != nil {
		panic(err)
	}

	fmt.Println("Done")
}
