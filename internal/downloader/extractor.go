package downloader

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
)

func ExtractAllFiles(archivePath string, destination string) error {
	archive, err := zip.OpenReader(archivePath)
	if err != nil {
		return err
	}
	defer archive.Close()

	for _, file := range archive.File {
		if file.FileInfo().IsDir() {
			continue
		}

		src, err := file.Open()
		if err != nil {
			return err
		}

		dstFile, err := os.OpenFile(fmt.Sprintf("%s/%s", destination, file.Name), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
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
