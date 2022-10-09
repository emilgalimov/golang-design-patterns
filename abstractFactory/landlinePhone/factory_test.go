package landlinePhone

import (
	"github.com/emilgalimov/golang-design-patterns/abstractFactory/phoneInterfaces"
	"reflect"
	"testing"
)

func TestFactory_CreatePrinter(t *testing.T) {
	tests := []struct {
		name string
		want phoneInterfaces.Printer
	}{
		{
			"It returns printer",
			&Printer{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Factory{}
			if got := f.CreatePrinter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreatePrinter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFactory_CreateProcessor(t *testing.T) {
	tests := []struct {
		name string
		want phoneInterfaces.Processor
	}{
		{
			"It returns processor",
			&Processor{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Factory{}
			if got := f.CreateProcessor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateProcessor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFactory_CreateSource(t *testing.T) {
	tests := []struct {
		name string
		want phoneInterfaces.Source
	}{
		{
			"It returns source",
			&Source{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := Factory{}
			if got := f.CreateSource(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateSource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewFactory(t *testing.T) {
	tests := []struct {
		name string
		want *Factory
	}{
		{
			"creates factory",
			&Factory{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFactory(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFactory() = %v, want %v", got, tt.want)
			}
		})
	}
}
