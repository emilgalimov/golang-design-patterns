package mobilePhone

import (
	"reflect"
	"testing"
)

func Test_processor_ProcessRaw(t *testing.T) {
	type args struct {
		rawMessage string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"It Decodes digital_message_is",
			args{
				rawMessage: "digital_message_is",
			},
			"digital message is",
		},
		{
			"It Decodes hello_world",
			args{
				rawMessage: "hello_world",
			},
			"hello world",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pr := Processor{}
			if got := pr.ProcessRaw(tt.args.rawMessage); got != tt.want {
				t.Errorf("ProcessRaw() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewProcessor(t *testing.T) {
	tests := []struct {
		name string
		want *Processor
	}{
		{
			"It returns processor",
			&Processor{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProcessor(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProcessor() = %v, want %v", got, tt.want)
			}
		})
	}
}
