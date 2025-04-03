package main

import (
	"fmt"
	"io"
	"path/filepath"
	"reflect"
	"sync"

	"github.com/mephir/teryt-golang/internal/collection"
	"github.com/mephir/teryt-golang/internal/dataset/datastruct"
	"github.com/mephir/teryt-golang/internal/dataset/model"
	"github.com/mephir/teryt-golang/internal/parser"
)

func main() {
	path, err := filepath.Abs("_zrzuty/TERC_Adresowy_2025-04-01.xml")
	if err != nil {
		panic(err)
	}

	parser, err := parser.Open(path)
	if err != nil {
		panic(err)
	}
	defer parser.Close()
	voivodeships := collection.NewCollection[model.Voivodeship]()
	counties := collection.NewCollection[model.County]()
	municipalisties := collection.NewCollection[model.Municipality]()

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
	go processVoivodeships(voivodeshipsChan, orphanedCounties, &wg, mutexes, voivodeships)
	go processCounties(countiesChan, voivodeships, orphanedCounties, &wg, mutexes, counties, orphanedMunicipalities)
	go processMunicipalities(municipalitiesChan, counties, orphanedMunicipalities, &wg, mutexes, municipalisties)

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
			// if m.Type.Id >= 5 { // prawdziwe gminy w ludzkiej glowie
			// 	continue
			// }
			municipalitiesChan <- m
		}
	}

	close(voivodeshipsChan)
	close(countiesChan)
	close(municipalitiesChan)
	wg.Wait()

	total := voivodeships.Count() + counties.Count() + municipalisties.Count()

	fmt.Printf("Total: %d\n", total)
	fmt.Printf("Voivodeships: %d\n", voivodeships.Count())
	fmt.Printf("Counties: %d\n", counties.Count())
	fmt.Printf("Municipalities: %d\n", municipalisties.Count())

	for _, v := range voivodeships.Items {
		fmt.Printf("%s %s\n", v.UnitType, v.Name)
		for _, c := range v.Counties {
			fmt.Printf("\t%s %s\n", c.UnitType, c.Name)
			for _, m := range c.Municipalities {
				fmt.Printf("\t\t%s %s\n", m.UnitType, m.Name)
			}
		}
	}
}

func processVoivodeships(
	voivodeshipsChan chan *model.Voivodeship,
	orphanedCounties map[uint][]*model.County,
	wg *sync.WaitGroup,
	mutexes map[string]*sync.Mutex,
	voivodeships *collection.Collection[model.Voivodeship],
) {
	defer wg.Done()

	for v := range voivodeshipsChan {
		voivodeships.Add(v)

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
	countiesChan chan *model.County,
	voivodeships *collection.Collection[model.Voivodeship],
	orphanedCounties map[uint][]*model.County,
	wg *sync.WaitGroup,
	mutexes map[string]*sync.Mutex,
	counties *collection.Collection[model.County],
	orphanedMunicipalities map[uint][]*model.Municipality,
) {
	defer wg.Done()

	for c := range countiesChan {
		counties.Add(c)
		if v := voivodeships.Get(c.VoivodeshipId); v != nil {
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
	municipalitiesChan chan *model.Municipality,
	counties *collection.Collection[model.County],
	orphanedMunicipalities map[uint][]*model.Municipality,
	wg *sync.WaitGroup,
	mutexes map[string]*sync.Mutex,
	municipalisties *collection.Collection[model.Municipality],
) {
	defer wg.Done()

	for m := range municipalitiesChan {
		municipalisties.Add(m)
		if c := counties.Get(m.CountyId + 100*m.VoivodeshipId); c != nil {
			c.Municipalities = append(c.Municipalities, m)
			m.County = c
		} else {
			mutexes["municipalities"].Lock()
			orphanedMunicipalities[m.CountyId+100*m.VoivodeshipId] = append(orphanedMunicipalities[m.CountyId], m)
			mutexes["municipalities"].Unlock()
		}
	}
}
