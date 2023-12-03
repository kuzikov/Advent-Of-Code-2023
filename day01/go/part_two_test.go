package main

import (
	"testing"
)

func Test_starTwo(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "hlpqrdh3",
			args: args{"hlpqrdh3"},
			want: 33,
		},
		{
			name: "324pzonenine",
			args: args{"324pzonenine"},
			want: 39,
		}, {
			name: "sevenineightwone",
			args: args{"sevenineightwone"},
			want: 71,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := starTwo(tt.args.str); got != tt.want {
				t.Errorf("starTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
