package mobilePhone

import (
	"github.com/emilgalimov/golang-design-patterns/abstractFactory"
)

type Factory struct {
}

func (f *Factory) CreateSource() abstractFactory.Source {
	//TODO implement me
	panic("implement me")
}

func (f *Factory) CreateProcessor() abstractFactory.Processor {
	//TODO implement me
	panic("implement me")
}

func (f *Factory) CreatePrinter() abstractFactory.Printer {
	//TODO implement me
	panic("implement me")
}

func NewFactory() *Factory {
	return &Factory{}
}
