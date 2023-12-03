package main

import "testing"

func Test_partOne(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "only nums",
			args: args{"8274"},
			want: 84,
		}, {
			name: "mixed",
			args: args{"four82nine74eou"},
			want: 84,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partOne(tt.args.str); got != tt.want {
				t.Errorf("partOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
