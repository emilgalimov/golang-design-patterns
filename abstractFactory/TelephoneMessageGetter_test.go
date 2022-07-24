package abstractFactory

import "testing"

func Test_writeMessage1(t *testing.T) {
	type args struct {
		factory PhoneFactory
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "it writes message",
			args: args{
				factory:
			},
		},
	},
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := writeMessage(tt.args.factory); got != tt.want {
				t.Errorf("writeMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
