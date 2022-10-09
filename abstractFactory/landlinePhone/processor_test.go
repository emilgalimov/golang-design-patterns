package landlinePhone

import (
	"reflect"
	"testing"
)

func TestNewProcessor(t *testing.T) {
	tests := []struct {
		name string
		want *Processor
	}{
		{
			"create processor",
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

func TestProcessor_ProcessRaw(t *testing.T) {
	type args struct {
		rawMessage string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"process row message",
			args{
				"hello-world",
			},
			"hello world",
		},
		{
			"process row message",
			args{
				"goodbye-world-2",
			},
			"goodbye world 2",
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
