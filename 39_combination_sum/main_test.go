package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_combinationSum(t *testing.T) {
	type args struct {
		candidates []int
		target     int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "case_1",
			args: args{
				candidates: []int{2, 3, 6, 7},
				target:     7,
			},
			want: [][]int{{2, 2, 3}, {7}},
		},
		{
			name: "case_2",
			args: args{
				candidates: []int{2, 3, 5},
				target:     8,
			},
			want: [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}},
		},
		{
			name: "case_3",
			args: args{
				candidates: []int{2},
				target:     1,
			},
			want: [][]int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum(tt.args.candidates, tt.args.target)
			require.Equal(t, len(tt.want), len(got))
			for _, exp := range tt.want {
				found := false
				for _, act := range got {
					if ElementsMatch(exp, act) {
						found = true
						break
					}
				}
				if !found {
					t.Fatalf("slice %v not found", exp)
				}
			}
		})
	}
}

func ElementsMatch(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	sort.Slice(a, func(i, j int) bool { return a[i] < a[j] })
	sort.Slice(b, func(i, j int) bool { return b[i] < b[j] })

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
