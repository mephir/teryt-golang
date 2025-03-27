package main

import (
	"archive/zip"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"time"

	"github.com/mephir/teryt-golang/internal/dataset"

	"github.com/carlmjohnson/requests"
)

func retrieve_viewstate(ctx context.Context) (map[string]string, error) {
	output := make(map[string]string)
	client := *http.DefaultClient
	client.CheckRedirect = requests.NoFollow

	var body string
	err := requests.
		URL("https://eteryt.stat.gov.pl/eTeryt/rejestr_teryt/udostepnianie_danych/baza_teryt/uzytkownicy_indywidualni/pobieranie/pliki_pelne.aspx?contrast=default").
		Client(&client).
		UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:136.0) Gecko/20100101 Firefox/136.0").
		ToString(&body).
		Fetch(ctx)

	if err != nil {
		return nil, err
	}

	regex := regexp.MustCompile(`id="__VIEWSTATE" value="(.+?)"`)
	matches := regex.FindStringSubmatch(body)
	if len(matches) > 0 {
		output["__VIEWSTATE"] = matches[1]
	} else {
		return nil, fmt.Errorf("no __VIEWSTATE found")
	}

	regex = regexp.MustCompile(`id="__VIEWSTATEGENERATOR" value="(.+?)"`)
	matches = regex.FindStringSubmatch(body)
	if len(matches) > 0 {
		output["__VIEWSTATEGENERATOR"] = matches[1]
	} else {
		return nil, fmt.Errorf("no __VIEWSTATEGENERATOR found")
	}

	return output, nil
}

func param_tbdata() string {
	// those data has to be in polish unfortunately(in proper grammatical case), since those are names ofmonths i decided to hard code them instead using some library
	months := map[int]string{
		1:  "stycznia",
		2:  "lutego",
		3:  "marca",
		4:  "kwietnia",
		5:  "maja",
		6:  "czerwca",
		7:  "lipca",
		8:  "sierpnia",
		9:  "września",
		10: "października",
		11: "listopada",
		12: "grudnia",
	}

	now := time.Now().Local()
	year, month, day := now.Date()

	return fmt.Sprintf("%d %s %d", day, months[int(month)], year)
}

func download_file(ctx context.Context, viewstate map[string]string, dataset dataset.Dataset) error {
	client := *http.DefaultClient
	client.CheckRedirect = requests.NoFollow

	params := url.Values{}
	params.Add("__VIEWSTATE", viewstate["__VIEWSTATE"])
	params.Add("__VIEWSTATEGENERATOR", viewstate["__VIEWSTATEGENERATOR"])
	params.Add("__EVENTTARGET", dataset.ToTarget())
	params.Add("ctl00$body$TBData", param_tbdata())

	err := requests.
		URL("https://eteryt.stat.gov.pl/eTeryt/rejestr_teryt/udostepnianie_danych/baza_teryt/uzytkownicy_indywidualni/pobieranie/pliki_pelne.aspx?contrast=default").
		Client(&client).
		UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:136.0) Gecko/20100101 Firefox/136.0").
		BodyForm(params).
		ToFile(dataset.ToFilename()).
		Fetch(ctx)

	if err != nil {
		return err
	}

	return nil
}

func extract_file(dataset dataset.Dataset) error {
	archive, err := zip.OpenReader(dataset.ToFilename())
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
	fmt.Println("Retrieving viewstate")

	viewstate, err := retrieve_viewstate(context.Background())
	if err != nil {
		panic(err)
	}

	dataset := dataset.Dataset{Name: "SIMC", Variant: "A"}

	fmt.Println("Downloading file")
	err = download_file(context.Background(), viewstate, dataset)
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
