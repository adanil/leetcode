package main

import "testing"

func Test_trailingZeroes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "case_1",
			args: args{n: 3},
			want: 0,
		},
		{
			name: "case_2",
			args: args{n: 5},
			want: 1,
		},
		{
			name: "case_3",
			args: args{n: 0},
			want: 0,
		},
		{
			name: "case_4",
			args: args{n: 8},
			want: 1,
		},
		{
			name: "case_5",
			args: args{n: 10},
			want: 2,
		},
		{
			name: "case_6",
			args: args{n: 30},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trailingZeroes(tt.args.n); got != tt.want {
				t.Errorf("trailingZeroes() = %v, want %v", got, tt.want)
			}
		})
	}
}
