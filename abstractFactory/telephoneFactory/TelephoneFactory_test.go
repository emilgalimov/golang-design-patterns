package telephoneFactory

import (
	"github.com/emilgalimov/golang-design-patterns/abstractFactory/landlinePhone"
	"github.com/emilgalimov/golang-design-patterns/abstractFactory/mobilePhone"
	"reflect"
	"testing"
)

func TestGetTelephoneFactory(t *testing.T) {
	type args struct {
		phoneType int
	}
	tests := []struct {
		name    string
		args    args
		want    PhoneFactory
		wantErr bool
	}{
		{
			name:    "it returns mobile phone",
			args:    args{phoneType: 0},
			want:    &mobilePhone.Factory{},
			wantErr: false,
		},
		{
			name:    "it returns landline phone",
			args:    args{phoneType: 1},
			want:    &landlinePhone.Factory{},
			wantErr: false,
		},
		{
			name:    "it returns error when unknown type",
			args:    args{phoneType: 2},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetTelephoneFactory(tt.args.phoneType)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTelephoneFactory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTelephoneFactory() got = %v, want %v", got, tt.want)
			}
		})
	}
}
