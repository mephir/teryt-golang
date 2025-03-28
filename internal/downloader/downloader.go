package downloader

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"time"

	"github.com/mephir/teryt-golang/internal/dataset"

	"github.com/carlmjohnson/requests"
	"github.com/gabriel-vasile/mimetype"
)

type viewstate struct {
	__VIEWSTATE          string
	__VIEWSTATEGENERATOR string
}

func (v *viewstate) Fetch() error {
	client := *http.DefaultClient
	client.CheckRedirect = requests.NoFollow
	client.Timeout = 10 * time.Second

	var body string
	err := requests.
		URL("https://eteryt.stat.gov.pl/eTeryt/rejestr_teryt/udostepnianie_danych/baza_teryt/uzytkownicy_indywidualni/pobieranie/pliki_pelne.aspx?contrast=default").
		Client(&client).
		UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:136.0) Gecko/20100101 Firefox/136.0").
		ToString(&body).
		Fetch(context.Background())

	if err != nil {
		return err
	}

	regex := regexp.MustCompile(`id="__VIEWSTATE" value="(.+?)"`)
	matches := regex.FindStringSubmatch(body)
	if len(matches) > 0 {
		v.__VIEWSTATE = matches[1]
	} else {
		return fmt.Errorf("no __VIEWSTATE found")
	}

	regex = regexp.MustCompile(`id="__VIEWSTATEGENERATOR" value="(.+?)"`)
	matches = regex.FindStringSubmatch(body)
	if len(matches) > 0 {
		v.__VIEWSTATEGENERATOR = matches[1]
	} else {
		return fmt.Errorf("no __VIEWSTATEGENERATOR found")
	}

	return nil
}

func tbdate(date time.Time) string {
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

	year, month, day := date.Date()

	return fmt.Sprintf("%d %s %d", day, months[int(month)], year)
}

func DownloadDataset(dataset dataset.Dataset, outputPath string) error {
	state := viewstate{}
	err := state.Fetch()
	if err != nil {
		return err
	}

	client := *http.DefaultClient
	client.CheckRedirect = requests.NoFollow
	client.Timeout = 10 * time.Second

	form := url.Values{}
	form.Add("__VIEWSTATE", state.__VIEWSTATE)
	form.Add("__VIEWSTATEGENERATOR", state.__VIEWSTATEGENERATOR)
	form.Add("__EVENTTARGET", dataset.ToTarget())
	form.Add("ctl00$body$TBData", tbdate(time.Now().Local()))

	err = requests.
		URL("https://eteryt.stat.gov.pl/eTeryt/rejestr_teryt/udostepnianie_danych/baza_teryt/uzytkownicy_indywidualni/pobieranie/pliki_pelne.aspx?contrast=default").
		Client(&client).
		UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:136.0) Gecko/20100101 Firefox/136.0").
		BodyForm(form).
		ToFile(outputPath).
		Fetch(context.Background())

	if err != nil {
		return err
	}

	// this check is necessary due page does not always return clear http status code when download fails
	// so we need to check if file is valid zip
	mimetype.SetLimit(32) // zip requires only 8 bytes to detect, but we need to read more to detect other types
	mime, err := mimetype.DetectFile(outputPath)
	if err != nil {
		return fmt.Errorf("could not detect mimetype: %w", err)
	}

	if !mime.Is("application/zip") {
		os.Remove(outputPath)
		return fmt.Errorf("invalid mimetype: %s", mime.String())
	}

	return nil
}
