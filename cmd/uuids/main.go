package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/mephir/teryt-golang/internal/teryt"
	"github.com/mr-tron/base58"
)

type inputFiles []string

type WriteJob struct {
	Offset int64
	Data   []byte
}

var output string
var input inputFiles
var separator string
var useBase58 bool
var useShortString bool
var useFullString bool
var consumersPoolSize uint

// var producersPoolSize uint

func init() {
	flag.StringVar(&output, "output", "", "Output file")
	flag.StringVar(&output, "o", "", "Output file (short)")
	flag.Var(&input, "input", "Input XML files to process (can be specified multiple times)")
	flag.Var(&input, "i", "Input XML files to process (can be specified multiple times) (short)")
	flag.StringVar(&separator, "separator", "", "Separator for uuids")
	flag.BoolVar(&useBase58, "base58", false, "Use base58 format for uuids")
	flag.BoolVar(&useShortString, "short", false, "Use short string format for uuids")
	flag.BoolVar(&useFullString, "full", false, "Use full string format for uuids")
	flag.UintVar(&consumersPoolSize, "consumers", 5, "Number of consumers (file writers)")
	// flag.UintVar(&producersPoolSize, "producers", 3, "Number of producers (data providers)")
}

func main() {
	flag.Parse()

	trueCounter := 0
	length := 16 // length in bytes to move offset while saving

	if useBase58 {
		trueCounter++
		length = 22
	}
	if useShortString {
		trueCounter++
		length = 32
	}
	if useFullString {
		trueCounter++
		length = 36
	}
	if trueCounter > 1 {
		panic("Only one of --base58, --short, or --full can be used at a time")
	}

	if separator != "" {
		length += len(separator)
	}

	basePath, err := filepath.Abs("_zrzuty/")
	if err != nil {
		panic(err)
	}

	paths := []string{
		filepath.Join(basePath, "TERC_Adresowy_2025-04-01.xml"),
		filepath.Join(basePath, "SIMC_Adresowy_2025-04-01.xml"),
		filepath.Join(basePath, "ULIC_Adresowy_2025-04-01.xml"),
	}

	teryt := teryt.NewInstance()
	teryt.LoadFromFiles(paths...)

	fh, err := os.Create(output) //os.OpenFile(output, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	defer fh.Close()

	var wgWriters sync.WaitGroup
	// var wgProducers sync.WaitGroup
	var sequence int64 = 0
	writeJobChan := make(chan *WriteJob, 100)

	for i := range int(consumersPoolSize) {
		wgWriters.Add(1)

		go writer(i, fh, writeJobChan, &wgWriters)
	}

	start := time.Now()

	for item := range teryt.Voivodeships.Iterator() {
		var data []byte
		id := item.Uuid()

		if useFullString {
			data = []byte(id.String())
		} else if useShortString {
			data = []byte(strings.ReplaceAll(id.String(), "-", ""))
		} else if useBase58 {
			data = []byte(base58.Encode([]byte(id.String())))
		} else {
			data = id[:]
		}

		if separator != "" {
			data = append(data, []byte(separator)...)
		}

		job := &WriteJob{
			Data:   []byte(data),
			Offset: sequence * int64(length),
		}
		sequence++

		writeJobChan <- job
	}

	for item := range teryt.Counties.Iterator() {
		var data []byte
		id := item.Uuid()

		if useFullString {
			data = []byte(id.String())
		} else if useShortString {
			data = []byte(strings.ReplaceAll(id.String(), "-", ""))
		} else if useBase58 {
			data = []byte(base58.Encode([]byte(id.String())))
		} else {
			data = id[:]
		}

		if separator != "" {
			data = append(data, []byte(separator)...)
		}

		job := &WriteJob{
			Data:   []byte(data),
			Offset: sequence * int64(length),
		}
		sequence++

		writeJobChan <- job
	}

	for item := range teryt.Municipalities.Iterator() {
		var data []byte
		id := item.Uuid()

		if useFullString {
			data = []byte(id.String())
		} else if useShortString {
			data = []byte(strings.ReplaceAll(id.String(), "-", ""))
		} else if useBase58 {
			data = []byte(base58.Encode([]byte(id.String())))
		} else {
			data = id[:]
		}

		if separator != "" {
			data = append(data, []byte(separator)...)
		}

		job := &WriteJob{
			Data:   []byte(data),
			Offset: sequence * int64(length),
		}
		sequence++

		writeJobChan <- job
	}

	for item := range teryt.Localities.Iterator() {
		var data []byte
		id := item.Uuid()

		if useFullString {
			data = []byte(id.String())
		} else if useShortString {
			data = []byte(strings.ReplaceAll(id.String(), "-", ""))
		} else if useBase58 {
			data = []byte(base58.Encode([]byte(id.String())))
		} else {
			data = id[:]
		}

		if separator != "" {
			data = append(data, []byte(separator)...)
		}

		job := &WriteJob{
			Data:   []byte(data),
			Offset: sequence * int64(length),
		}
		sequence++

		writeJobChan <- job
	}

	for item := range teryt.Streets.Iterator() {
		var data []byte
		id := item.Uuid()

		if useFullString {
			data = []byte(id.String())
		} else if useShortString {
			data = []byte(strings.ReplaceAll(id.String(), "-", ""))
		} else if useBase58 {
			data = []byte(base58.Encode([]byte(id.String())))
		} else {
			data = id[:]
		}

		if separator != "" {
			data = append(data, []byte(separator)...)
		}

		job := &WriteJob{
			Data:   []byte(data),
			Offset: sequence * int64(length),
		}
		sequence++

		writeJobChan <- job
	}

	close(writeJobChan)
	wgWriters.Wait()

	elapsed := time.Since(start)
	fmt.Printf("Elapsed time: %s\n", elapsed)
}

func writer(id int, fh *os.File, channel <-chan *WriteJob, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range channel {
		_, err := fh.WriteAt((*job).Data, (*job).Offset)
		if err != nil {
			fmt.Printf("[W%d] Error writing at %d: %v\n", id, job.Offset, err)
			continue
		}
		// fmt.Printf("[W%d] Wrote at offset %d\n", id, job.Offset)
	}
}

func (i *inputFiles) Set(value string) error {
	if value == "" {
		return fmt.Errorf("input file cannot be empty")
	}

	if _, err := os.Stat(value); err != nil {
		return err
	}

	*i = append(*i, value)
	return nil
}

func (i *inputFiles) String() string {
	return fmt.Sprint(*i)
}
