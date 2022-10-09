package mobilePhone

import (
	"reflect"
	"testing"
)

func Test_printer_Print(t *testing.T) {
	type args struct {
		message string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"it prints Hello Moto",
			args{
				message: "Hello Moto",
			},
			"mobile phone says: Hello Moto",
		},
		{
			"it prints Nokia3310 message",
			args{
				message: "Nokia3310 message",
			},
			"mobile phone says: Nokia3310 message",
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

func TestNewPrinter(t *testing.T) {
	tests := []struct {
		name string
		want *Printer
	}{
		{
			"It returns printer",
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
