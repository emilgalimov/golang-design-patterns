package landlinePhone

import (
	"github.com/emilgalimov/golang-design-patterns/abstractFactory/phoneInterfaces"
)

type Factory struct {
}

func NewFactory() *Factory {
	return &Factory{}
}

func (f Factory) CreateSource() phoneInterfaces.Source {
	return NewSource()
}

func (f Factory) CreateProcessor() phoneInterfaces.Processor {
	return NewProcessor()
}

func (f Factory) CreatePrinter() phoneInterfaces.Printer {
	return NewPrinter()
}
