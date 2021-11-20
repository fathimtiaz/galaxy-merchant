package galaxymerchant

import (
	"testing"
)

func Test_isDictionaryLineType(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{line: "this is L"},
			want: true,
		},
		{
			args: args{line: "this i L"},
			want: false,
		},
		{
			name: "true",
			args: args{line: "this is H"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isDictionaryLineType(tt.args.line); got != tt.want {
				t.Errorf("isDictionaryLineType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isPriceLineType(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{line: "Credits"},
			want: false,
		},
		{
			name: "no is",
			args: args{line: "a b ia 1 Credits"},
			want: false,
		},
		{
			name: "minimal",
			args: args{line: "a b is 8 Credits"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPriceLineType(tt.args.line); got != tt.want {
				t.Errorf("isPriceLineType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isQueryLineType(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			args: args{line: "?"},
			want: false,
		},
		{
			name: "no is",
			args: args{line: "how much ia a ?"},
			want: false,
		},
		{
			name: "minimal",
			args: args{line: "how much is a ?"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isQueryLineType(tt.args.line); got != tt.want {
				t.Errorf("isQueryLineType() = %v, want %v", got, tt.want)
			}
		})
	}
}
