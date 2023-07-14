package stringbyteconverter

import (
	"reflect"
	"testing"
)

func TestConverter_ToBytes(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Converts string to byte",
			args: args{s: "yuksel"},
			want: 6,
		},

		{
			name: "Converts empty string to byte",
			args: args{s: ""},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Converter{}
			if got := c.ToBytes(tt.args.s); !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("ToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConverter_ToString(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Converts byte array to string",
			args: args{b: []byte{121, 117, 107, 115, 101, 108}},
			want: "yuksel",
		},

		{
			name: "Converts nil byte array to string",
			args: args{b: nil},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Converter{}
			if got := c.ToString(tt.args.b); *got != tt.want {
				t.Errorf("ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
