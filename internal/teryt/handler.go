package teryt

import (
	"sync"

	"github.com/mephir/teryt-golang/internal/dataset/model"
	"github.com/mephir/teryt-golang/internal/parser"
)

type parsingHandler struct {
	teryt *Teryt

	// channels and routines handling
	voivodeshipsChan   chan *model.Voivodeship
	countiesChan       chan *model.County
	municipalitiesChan chan *model.Municipality
	localityTypesChan  chan *model.LocalityType
	localitiesChan     chan *model.Locality
	wg                 sync.WaitGroup

	// maps to handle orphans
	tercOrphansCounty                sync.Map
	tercOrphansMunicipality          sync.Map
	simcOrphanedLocalityMunicipality sync.Map
	simcOrphanedLocalityType         sync.Map
}

func (p *parsingHandler) Close() {
	close(p.voivodeshipsChan)
	close(p.countiesChan)
	close(p.municipalitiesChan)
	close(p.localityTypesChan)
	close(p.localitiesChan)
	p.wg.Wait()
}

func (p *parsingHandler) parseSimc(parser parser.Parser) {
	p.wg.Add(1)
	go p.processLocalities()

	for {
		data, err := parser.Fetch()
		if err != nil {
			break
		}

		m, err := data.ToModel()
		if err != nil {
			panic(err)
		}

		switch m := m.(type) {
		case *model.Locality:
			p.localitiesChan <- m
		default:
			panic("Invalid type")
		}
	}
}

func (p *parsingHandler) parseWmrodz(parser parser.Parser) {
	p.wg.Add(1)
	go p.processLocalityTypes()

	for {
		data, err := parser.Fetch()
		if err != nil {
			break
		}

		m, err := data.ToModel()
		if err != nil {
			panic(err)
		}

		switch m := m.(type) {
		case *model.LocalityType:
			p.localityTypesChan <- m
		default:
			panic("Invalid type")
		}
	}
}

func (p *parsingHandler) parseTerc(parser parser.Parser) {
	p.wg.Add(3)
	go p.processVoivodeships()
	go p.processCounties()
	go p.processMunicipalities()

	for {
		data, err := parser.Fetch()
		if err != nil {
			break
		}

		m, err := data.ToModel()
		if err != nil {
			panic(err)
		}

		switch m := m.(type) {
		case *model.Voivodeship:
			p.voivodeshipsChan <- m
		case *model.County:
			p.countiesChan <- m
		case *model.Municipality:
			p.municipalitiesChan <- m
		default:
			panic("Invalid type")
		}
	}
}

func (p *parsingHandler) processVoivodeships() {
	defer p.wg.Done()

	for voivodeship := range p.voivodeshipsChan {
		p.teryt.Voivodeships.Add(voivodeship)

		if cs, ok := p.tercOrphansCounty.Load(voivodeship.Id); ok {
			for _, county := range cs.([]*model.County) {
				county.Voivodeship = voivodeship
			}
			p.tercOrphansCounty.Delete(voivodeship.Id)
		}
	}
}

func (p *parsingHandler) processCounties() {
	defer p.wg.Done()
	for county := range p.countiesChan {
		p.teryt.Counties.Add(county)

		if v := p.teryt.Voivodeships.Get(county.VoivodeshipId); v != nil {
			county.Voivodeship = v
			v.Counties = append(v.Counties, county)
		} else {
			if orphans, ok := p.tercOrphansCounty.Load(county.VoivodeshipId); ok {
				orphans = append(orphans.([]*model.County), county)
				p.tercOrphansCounty.Store(county.VoivodeshipId, orphans)
			} else {
				p.tercOrphansCounty.Store(county.VoivodeshipId, []*model.County{county})
			}
		}

		if ms, ok := p.tercOrphansMunicipality.Load(county.Identifier()); ok {
			for _, municipality := range ms.([]*model.Municipality) {
				municipality.County = county
			}
			p.tercOrphansMunicipality.Delete(county.Identifier())
		}
	}
}

func (p *parsingHandler) processMunicipalities() {
	defer p.wg.Done()

	for m := range p.municipalitiesChan {
		p.teryt.Municipalities.Add(m)

		if c := p.teryt.Counties.Get(m.GetCountyIdentifier()); c != nil {
			m.County = c
			c.Municipalities = append(c.Municipalities, m)
		} else {
			if orphans, ok := p.tercOrphansMunicipality.Load(m.GetCountyIdentifier()); ok {
				orphans = append(orphans.([]*model.Municipality), m)
				p.tercOrphansMunicipality.Store(m.GetCountyIdentifier(), orphans)
			} else {
				p.tercOrphansMunicipality.Store(m.GetCountyIdentifier(), []*model.Municipality{m})
			}
		}
	}
}

func (p *parsingHandler) processLocalityTypes() {
	defer p.wg.Done()

	for m := range p.localityTypesChan {
		p.teryt.LocalityTypes.Add(m)
	}
}

func (p *parsingHandler) processLocalities() {
	defer p.wg.Done()

	for l := range p.localitiesChan {
		p.teryt.Localities.Add(l)

		// handle municipality relation
		if m := p.teryt.Municipalities.Get(l.MunicipalityId); m != nil {
			l.Municipality = m
			m.Localities = append(m.Localities, l)
		} else {
			if orphans, ok := p.simcOrphanedLocalityMunicipality.Load(l.MunicipalityId); ok {
				orphans = append(orphans.([]*model.Locality), l)
				p.simcOrphanedLocalityMunicipality.Store(l.MunicipalityId, orphans)
			} else {
				p.simcOrphanedLocalityMunicipality.Store(l.MunicipalityId, []*model.Locality{l})
			}
		}
	}
}
