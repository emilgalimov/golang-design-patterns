package landlinePhone

import (
	"reflect"
	"testing"
)

func TestNewPrinter(t *testing.T) {
	tests := []struct {
		name string
		want *Printer
	}{
		{
			"creates printer",
			&Printer{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPrinter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPrinter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrinter_Print(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"creates printer",
			args{
				"Hello world",
			},
			"landline phone says:Hello world",
		},
		{
			"creates printer",
			args{
				"Goodbye world",
			},
			"landline phone says:Goodbye world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := Printer{}
			if got := pr.Print(tt.args.message); got != tt.want {
				t.Errorf("Print() = %v, want %v", got, tt.want)
			}
		})
	}
}
