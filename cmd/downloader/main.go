package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/mephir/teryt-golang/internal/dataset"
	"github.com/mephir/teryt-golang/internal/downloader"
)

type datasets dataset.Collection

func (d *datasets) String() string {
	return fmt.Sprint(*d)
}

func (d *datasets) Set(value string) error {
	if strings.ToLower(value) == "all" {
		*d = dataset.DefaultDatasets
		return nil
	}

	value = strings.ToUpper(value)
	for _, ds := range strings.Split(value, ",") {
		data := strings.Split(ds, "-")
		if len(data) == 1 {
			if data[0] == "WMRODZ" {
				*d = append(*d, dataset.Dataset{Name: data[0], Variant: ""})
			} else {
				*d = append(*d, dataset.Dataset{Name: data[0], Variant: "U"})
			}
		} else if len(data) == 2 {
			*d = append(*d, dataset.Dataset{Name: data[0], Variant: data[1]})
		} else {
			return fmt.Errorf("invalid datasets")
		}
	}

	for _, ds := range *d {
		if err := ds.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func getDownloadDirectory(outputDirectory string, extract bool) string {
	if extract {
		path, err := os.MkdirTemp("", "teryt")
		if err != nil {
			panic(err)
		}

		defer os.RemoveAll(path)

		return path
	}

	return outputDirectory
}

var datasetsFlag datasets
var listFlag bool
var allFlag bool
var outputDirFlag string
var extractFlag bool
var timeoutFlag int
var userAgentFlag string
var dateFlag string
var xmlOnlyFlag bool
var csvOnlyFlag bool

func init() {
	flag.Var(&datasetsFlag, "datasets", "Comma separated list of datasets")
	flag.Var(&datasetsFlag, "d", "Dataset(shorthand) can be used multiple times")
	flag.BoolVar(&listFlag, "list", false, "List available datasets")
	flag.BoolVar(&listFlag, "l", false, "List available datasets(shorthand)")
	flag.BoolVar(&allFlag, "all", false, "All datasets")
	flag.StringVar(&outputDirFlag, "output-dir", ".", "Output directory")
	flag.StringVar(&outputDirFlag, "o", ".", "Output directory(shorthand)")
	flag.BoolVar(&extractFlag, "extract", false, "Extract downloaded files")
	flag.BoolVar(&extractFlag, "e", false, "Extract downloaded files(shorthand)")
	flag.IntVar(&timeoutFlag, "timeout", 10, "HTTP client timeout in seconds")
	flag.StringVar(&userAgentFlag, "user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:136.0) Gecko/20100101 Firefox/136.0", "HTTP client User agent")
	flag.StringVar(&dateFlag, "date", "today", "Date of the dataset format: YYYY-MM-DD")
	flag.BoolVar(&xmlOnlyFlag, "xml-only", false, "Extract only XML files")
	flag.BoolVar(&csvOnlyFlag, "csv-only", false, "Extract only CSV files")
}

func main() {
	flag.Parse()

	if listFlag {
		fmt.Println("Available datasets:")
		for _, ds := range dataset.DefaultDatasets {
			fmt.Printf("%s - %s\n", strings.ToLower(ds.Id()), ds.ToString())
		}
		os.Exit(0)
	}

	if allFlag {
		flag.Set("datasets", "all")
	}

	if (len(datasetsFlag) == 0) && !listFlag && !allFlag {
		flag.Usage()
		fmt.Fprintln(os.Stderr, "No datasets provided")
		os.Exit(2)
	}

	if (xmlOnlyFlag || csvOnlyFlag) && !extractFlag {
		fmt.Fprintln(os.Stderr, "xml-only and csv-only flags require extract flag")
		os.Exit(2)
	}

	if xmlOnlyFlag && csvOnlyFlag {
		fmt.Fprintln(os.Stderr, "xml-only and csv-only flags are mutually exclusive")
		os.Exit(2)
	}

	var date time.Time
	if dateFlag == "today" {
		date = time.Now().Local()
	} else {
		var err error
		date, err = time.Parse("2006-01-02", dateFlag)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid date format")
			os.Exit(2)
		}
	}

	for _, ds := range datasetsFlag {
		downloadDirectory := getDownloadDirectory(outputDirFlag, extractFlag)
		filepath := filepath.Join(downloadDirectory, ds.ToFilename(date))
		fmt.Printf("Downloading %s...\n", ds.ToString())
		err := downloader.DownloadDataset(ds, filepath)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if extractFlag {
			fmt.Printf("Extracting %s...\n", ds.ToFilename(date))

			if xmlOnlyFlag {
				downloader.ExtractFiles(filepath, outputDirFlag, ".xml")
			} else if csvOnlyFlag {
				downloader.ExtractFiles(filepath, outputDirFlag, ".csv")
			} else {
				downloader.ExtractFiles(filepath, outputDirFlag, "")
			}
		}
	}
}
