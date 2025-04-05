package teryt

import (
	"fmt"
	"io"
	"reflect"
	"sync"

	"github.com/mephir/teryt-golang/internal/dataset/datastruct"
	"github.com/mephir/teryt-golang/internal/dataset/model"
	"github.com/mephir/teryt-golang/internal/parser"
)

type parsingHandler struct {
	teryt *Teryt

	wg *sync.WaitGroup
}

func newParsingHandler(teryt *Teryt) *parsingHandler {
	return &parsingHandler{
		teryt: teryt,
		wg:    &sync.WaitGroup{},
	}
}

func (p *parsingHandler) Add(parser parser.Parser) {
	var buffer int = 1000
	channel := reflect.MakeChan(reflect.ChanOf(reflect.BothDir, reflect.PointerTo(parser.GetStructType())), buffer).Interface()

	switch ch := channel.(type) {
	case chan *datastruct.Terc:
		p.wg.Add(1)
		go p.tercProcessChannel(ch)
	case chan *datastruct.Wmrodz:
		p.wg.Add(1)
		go p.wmrodzProcessChannel(ch)
	case chan *datastruct.Simc:
		p.wg.Add(1)
		go p.simcProcessChannel(ch)
	case chan *datastruct.Ulic:
		p.wg.Add(1)
		go p.ulicProcessChannel(ch)
	default:
		panic(fmt.Errorf("invalid channel type: %T", ch))
	}

	go p.sendDataToChannel(channel, parser)
}

func (p *parsingHandler) ulicProcessChannel(channel <-chan *datastruct.Ulic) {
	defer p.wg.Done()

	for data := range channel {
		if m, err := data.ToModel(); err != nil {
			panic(fmt.Errorf("could not convert data to model: %w", err))
		} else {
			switch m := m.(type) {
			case *model.Street:
				if err := p.teryt.Streets.Add(m); err != nil {
					panic(fmt.Errorf("could not add street to collection: %w", err))
				}
			default:
				panic(fmt.Errorf("invalid type: %T", m))
			}
		}
	}
}

func (p *parsingHandler) simcProcessChannel(channel <-chan *datastruct.Simc) {
	defer p.wg.Done()

	for data := range channel {
		if m, err := data.ToModel(); err != nil {
			panic(fmt.Errorf("could not convert data to model: %w", err))
		} else {
			switch m := m.(type) {
			case *model.Locality:
				p.teryt.Localities.Add(m)
			default:
				panic(fmt.Errorf("invalid type: %T", m))
			}
		}
	}
}

func (p *parsingHandler) wmrodzProcessChannel(channel <-chan *datastruct.Wmrodz) {
	defer p.wg.Done()

	for data := range channel {
		if m, err := data.ToModel(); err != nil {
			panic(fmt.Errorf("could not convert data to model: %w", err))
		} else {
			switch m := m.(type) {
			case *model.LocalityType:
				p.teryt.LocalityTypes.Add(m)
			default:
				panic(fmt.Errorf("invalid type: %T", m))
			}
		}
	}
}

func (p *parsingHandler) tercProcessChannel(channel <-chan *datastruct.Terc) {
	defer p.wg.Done()

	for data := range channel {
		if m, err := data.ToModel(); err != nil {
			panic(fmt.Errorf("could not convert data to model: %w", err))
		} else {
			switch m := m.(type) {
			case *model.Voivodeship:
				p.teryt.Voivodeships.Add(m)
			case *model.County:
				p.teryt.Counties.Add(m)
			case *model.Municipality:
				p.teryt.Municipalities.Add(m)
			default:
				panic(fmt.Errorf("invalid type: %T", m))
			}
		}
	}
}

func (p *parsingHandler) sendDataToChannel(channel any, parser parser.Parser) {
	defer reflect.ValueOf(channel).Close()
	defer parser.Close()

	ch := reflect.ValueOf(channel).Interface()

	for {
		data, err := parser.Fetch()
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(fmt.Errorf("could not fetch data from %s: %w", parser.GetDataset().Name, err))
		}

		reflect.ValueOf(ch).Send(reflect.ValueOf(data))
	}
}

func (p *parsingHandler) Wait() {
	p.wg.Wait()
}
