package exporter

type Writer interface {
	WriteItem(item any) error
}
