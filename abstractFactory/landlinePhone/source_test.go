package landlinePhone

import (
	"reflect"
	"testing"
)

func TestNewSource(t *testing.T) {
	tests := []struct {
		name string
		want *Source
	}{
		{
			"get row message",
			&Source{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSource(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSource_GetRaw(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			"get row message",
			"landline-phone-message",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Source{}
			if got := s.GetRaw(); got != tt.want {
				t.Errorf("GetRaw() = %v, want %v", got, tt.want)
			}
		})
	}
}
