package mobilePhone

import (
	"reflect"
	"testing"
)

func Test_source_GetRaw(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			"It returns row message",
			"mobile_phone_message",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Source{}
			if got := s.GetRaw(); got != tt.want {
				t.Errorf("GetRaw() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSource(t *testing.T) {
	tests := []struct {
		name string
		want *Source
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSource(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSource() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewSource1(t *testing.T) {
	tests := []struct {
		name string
		want *Source
	}{
		{
			"It returns source",
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
