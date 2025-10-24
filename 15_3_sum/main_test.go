package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_threeSum(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "test_1",
			args: args{nums: []int{-1, 0, 1, 2, -1, -4}},
			want: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		{
			name: "test_2",
			args: args{nums: []int{0, 1, 1}},
			want: [][]int{},
		},
		{
			name: "test_3",
			args: args{nums: []int{0, 0, 0}},
			want: [][]int{{0, 0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.args.nums)
			require.Len(t, got, len(tt.want))
		})
	}
}
