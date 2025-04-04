package main

import (
	"fmt"
	"io"
	"path/filepath"
	"reflect"
	"sync"

	"github.com/mephir/teryt-golang/internal/dataset/datastruct"
	"github.com/mephir/teryt-golang/internal/dataset/model"
	"github.com/mephir/teryt-golang/internal/parser"
	"github.com/mephir/teryt-golang/internal/teryt"
)

func main() {
	path, err := filepath.Abs("_zrzuty/TERC_Adresowy_2025-04-01.xml")
	if err != nil {
		panic(err)
	}

	teryt := teryt.NewInstance()

	parser, err := parser.Open(path)
	if err != nil {
		panic(err)
	}
	defer parser.Close()

	orphanedCounties := make(map[uint][]*model.County)
	orphanedMunicipalities := make(map[uint][]*model.Municipality)
	var mutexes = map[string]*sync.Mutex{
		"counties":       {},
		"municipalities": {},
	}

	var wg sync.WaitGroup

	voivodeshipsChan := make(chan *model.Voivodeship)
	countiesChan := make(chan *model.County)
	municipalitiesChan := make(chan *model.Municipality)

	wg.Add(3)
	go processVoivodeships(teryt, voivodeshipsChan, &wg, mutexes, orphanedCounties)
	go processCounties(teryt, countiesChan, &wg, mutexes, orphanedCounties, orphanedMunicipalities)
	go processMunicipalities(teryt, municipalitiesChan, &wg, mutexes, orphanedMunicipalities)

	// // handle parsing
	for {
		data, err := parser.Fetch()
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		if (reflect.TypeOf(data) != reflect.TypeOf(&datastruct.Terc{})) {
			panic("Invalid type")
		}

		m, err := data.ToModel()
		if err != nil {
			panic(err)
		}

		switch m := m.(type) {
		case *model.Voivodeship:
			voivodeshipsChan <- m
		case *model.County:
			countiesChan <- m
		case *model.Municipality:
			municipalitiesChan <- m
		}
	}

	close(voivodeshipsChan)
	close(countiesChan)
	close(municipalitiesChan)
	wg.Wait()

	total := teryt.Voivodeships.Count() + teryt.Counties.Count() + teryt.Municipalities.Count()

	fmt.Printf("Total: %d\n", total)
	fmt.Printf("Voivodeships: %d\n", teryt.Voivodeships.Count())
	fmt.Printf("Counties: %d\n", teryt.Counties.Count())
	fmt.Printf("Municipalities: %d\n", teryt.Municipalities.Count())

	for v := range teryt.Voivodeships.Iterator() {
		fmt.Printf("%s %s\n", v.UnitType, v.Name)
		// for _, c := range v.Counties {
		// 	fmt.Printf("\t%s %s\n", c.UnitType, c.Name)
		// 	for _, m := range c.Municipalities {
		// 		fmt.Printf("\t\t%s %s\n", m.UnitType, m.Name)
		// 	}
		// }
	}
}

func processVoivodeships(
	teryt *teryt.Teryt,
	voivodeshipsChan chan *model.Voivodeship,
	wg *sync.WaitGroup,
	mutexes map[string]*sync.Mutex,
	orphanedCounties map[uint][]*model.County,
) {
	defer wg.Done()

	for v := range voivodeshipsChan {
		teryt.Voivodeships.Add(v)
		mutexes["counties"].Lock()
		if cs, exists := orphanedCounties[v.Identifier()]; exists {
			v.Counties = append(v.Counties, cs...)

			for _, county := range cs {
				county.Voivodeship = v
			}

			delete(orphanedCounties, v.Identifier())
		}
		mutexes["counties"].Unlock()
	}
}

func processCounties(
	teryt *teryt.Teryt,
	countiesChan chan *model.County,
	wg *sync.WaitGroup,
	mutexes map[string]*sync.Mutex,
	orphanedCounties map[uint][]*model.County,
	orphanedMunicipalities map[uint][]*model.Municipality,
) {
	defer wg.Done()

	for c := range countiesChan {
		teryt.Counties.Add(c)
		if v := teryt.Voivodeships.Get(c.VoivodeshipId); v != nil {
			v.Counties = append(v.Counties, c)
			c.Voivodeship = v
		} else {
			mutexes["counties"].Lock()
			orphanedCounties[c.VoivodeshipId] = append(orphanedCounties[c.VoivodeshipId], c)
			mutexes["counties"].Unlock()
		}

		mutexes["municipalities"].Lock()
		if ms, exists := orphanedMunicipalities[c.Identifier()]; exists {
			c.Municipalities = append(c.Municipalities, ms...)

			for _, municipality := range ms {
				municipality.County = c
			}

			delete(orphanedMunicipalities, c.Identifier())
		}
		mutexes["municipalities"].Unlock()
	}
}

func processMunicipalities(
	teryt *teryt.Teryt,
	municipalitiesChan chan *model.Municipality,
	wg *sync.WaitGroup,
	mutexes map[string]*sync.Mutex,
	orphanedMunicipalities map[uint][]*model.Municipality,
) {
	defer wg.Done()

	for m := range municipalitiesChan {
		teryt.Municipalities.Add(m)
		if c := teryt.Counties.Get(m.CountyId + 100*m.VoivodeshipId); c != nil {
			c.Municipalities = append(c.Municipalities, m)
			m.County = c
		} else {
			mutexes["municipalities"].Lock()
			orphanedMunicipalities[m.CountyId+100*m.VoivodeshipId] = append(orphanedMunicipalities[m.CountyId], m)
			mutexes["municipalities"].Unlock()
		}
	}
}
