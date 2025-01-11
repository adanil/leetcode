package main

import (
	"slices"
	"testing"
)

func Test_merge(t *testing.T) {
	type args struct {
		nums1 []int
		m     int
		nums2 []int
		n     int
	}
	tests := []struct {
		name   string
		args   args
		result []int
	}{
		{
			name: "case_1",
			args: args{
				nums1: []int{4, 13, 18, 0, 0, 0, 0, 0, 0},
				m:     3,
				nums2: []int{1, 5, 10, 12, 14, 27},
				n:     6,
			},
			result: []int{1, 4, 5, 10, 12, 13, 14, 18, 27},
		},
		{
			name: "case_2",
			args: args{
				nums1: []int{0, 0, 0},
				m:     0,
				nums2: []int{1, 2, 3},
				n:     3,
			},
			result: []int{1, 2, 3},
		},
		{
			name: "case_3",
			args: args{
				nums1: []int{1},
				m:     1,
				nums2: []int{},
				n:     0,
			},
			result: []int{1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.nums1, tt.args.m, tt.args.nums2, tt.args.n)
			if eq := slices.Equal(tt.args.nums1, tt.result); !eq {
				t.Fatalf("Actual and expected slices aren't equal\nexpected: %v\ngot: %v",
					tt.result, tt.args.nums1)
			}
		})
	}
}
